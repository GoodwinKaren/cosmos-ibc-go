package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	genesistypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/genesis/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
)

var (
	// TestOwnerAddress defines a reusable bech32 address for testing purposes
	TestOwnerAddress = "cosmos17dtl0mjt3t77kpuhg2edqzjpszulwhgzuj9ljs"

	// TestPortID defines a reusable port identifier for testing purposes
	TestPortID, _ = icatypes.NewControllerPortID(TestOwnerAddress)

	// TestVersion defines a reusable interchainaccounts version string for testing purposes
	TestVersion = string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: ibctesting.FirstConnectionID,
		HostConnectionId:       ibctesting.FirstConnectionID,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	chainC *ibctesting.TestChain
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
	suite.chainC = suite.coordinator.GetChain(ibctesting.GetChainID(3))
}

func NewICAPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = icatypes.HostPortID
	path.EndpointB.ChannelConfig.PortID = icatypes.HostPortID
	path.EndpointA.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointA.ChannelConfig.Version = TestVersion
	path.EndpointB.ChannelConfig.Version = TestVersion

	return path
}

// SetupICAPath invokes the InterchainAccounts entrypoint and subsequent channel handshake handlers
func SetupICAPath(path *ibctesting.Path, owner string) error {
	if err := RegisterInterchainAccount(path.EndpointA, owner); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenTry(); err != nil {
		return err
	}

	if err := path.EndpointA.ChanOpenAck(); err != nil {
		return err
	}

	return path.EndpointB.ChanOpenConfirm()
}

// RegisterInterchainAccount is a helper function for starting the channel handshake
func RegisterInterchainAccount(endpoint *ibctesting.Endpoint, owner string) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelSequence := endpoint.Chain.App.GetIBCKeeper().ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := endpoint.Chain.GetSimApp().ICAControllerKeeper.RegisterInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, owner, TestVersion); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = channeltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID

	return nil
}

func TestKeeperTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestGetAllPorts() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	expectedPorts := []string{TestPortID}

	ports := suite.chainA.GetSimApp().ICAControllerKeeper.GetAllPorts(suite.chainA.GetContext())
	suite.Require().Len(ports, len(expectedPorts))
	suite.Require().Equal(expectedPorts, ports)
}

func (suite *KeeperTestSuite) TestGetInterchainAccountAddress() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	counterpartyPortID := path.EndpointA.ChannelConfig.PortID

	retrievedAddr, found := suite.chainA.GetSimApp().ICAControllerKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), ibctesting.FirstConnectionID, counterpartyPortID)
	suite.Require().True(found)
	suite.Require().NotEmpty(retrievedAddr)

	retrievedAddr, found = suite.chainA.GetSimApp().ICAControllerKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), "invalid conn", "invalid port")
	suite.Require().False(found)
	suite.Require().Empty(retrievedAddr)
}

func (suite *KeeperTestSuite) TestGetAllActiveChannels() {
	var (
		expectedChannelID = "test-channel"
		expectedPortID    = "test-port"
	)

	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	suite.chainA.GetSimApp().ICAControllerKeeper.SetActiveChannelID(suite.chainA.GetContext(), ibctesting.FirstConnectionID, expectedPortID, expectedChannelID)

	expectedChannels := []genesistypes.ActiveChannel{
		{
			ConnectionId:        ibctesting.FirstConnectionID,
			PortId:              TestPortID,
			ChannelId:           path.EndpointA.ChannelID,
			IsMiddlewareEnabled: true,
		},
		{
			ConnectionId:        ibctesting.FirstConnectionID,
			PortId:              expectedPortID,
			ChannelId:           expectedChannelID,
			IsMiddlewareEnabled: false,
		},
	}

	activeChannels := suite.chainA.GetSimApp().ICAControllerKeeper.GetAllActiveChannels(suite.chainA.GetContext())
	suite.Require().Len(activeChannels, len(expectedChannels))
	suite.Require().Equal(expectedChannels, activeChannels)
}

func (suite *KeeperTestSuite) TestGetAllInterchainAccounts() {
	var (
		expectedAccAddr = "test-acc-addr"
		expectedPortID  = "test-port"
	)

	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, TestOwnerAddress)
	suite.Require().NoError(err)

	interchainAccAddr, exists := suite.chainB.GetSimApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), path.EndpointB.ConnectionID, path.EndpointA.ChannelConfig.PortID)
	suite.Require().True(exists)

	suite.chainA.GetSimApp().ICAControllerKeeper.SetInterchainAccountAddress(suite.chainA.GetContext(), ibctesting.FirstConnectionID, expectedPortID, expectedAccAddr)

	expectedAccounts := []genesistypes.RegisteredInterchainAccount{
		{
			ConnectionId:   ibctesting.FirstConnectionID,
			PortId:         TestPortID,
			AccountAddress: interchainAccAddr,
		},
		{
			ConnectionId:   ibctesting.FirstConnectionID,
			PortId:         expectedPortID,
			AccountAddress: expectedAccAddr,
		},
	}

	interchainAccounts := suite.chainA.GetSimApp().ICAControllerKeeper.GetAllInterchainAccounts(suite.chainA.GetContext())
	suite.Require().Len(interchainAccounts, len(expectedAccounts))
	suite.Require().Equal(expectedAccounts, interchainAccounts)
}

func (suite *KeeperTestSuite) TestIsActiveChannel() {
	suite.SetupTest()

	path := NewICAPath(suite.chainA, suite.chainB)
	owner := TestOwnerAddress
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, owner)
	suite.Require().NoError(err)
	portID := path.EndpointA.ChannelConfig.PortID

	isActive := suite.chainA.GetSimApp().ICAControllerKeeper.IsActiveChannel(suite.chainA.GetContext(), ibctesting.FirstConnectionID, portID)
	suite.Require().Equal(isActive, true)
}

func (suite *KeeperTestSuite) TestSetInterchainAccountAddress() {
	var (
		expectedAccAddr = "test-acc-addr"
		expectedPortID  = "test-port"
	)

	suite.chainA.GetSimApp().ICAControllerKeeper.SetInterchainAccountAddress(suite.chainA.GetContext(), ibctesting.FirstConnectionID, expectedPortID, expectedAccAddr)

	retrievedAddr, found := suite.chainA.GetSimApp().ICAControllerKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), ibctesting.FirstConnectionID, expectedPortID)
	suite.Require().True(found)
	suite.Require().Equal(expectedAccAddr, retrievedAddr)
}

func (suite *KeeperTestSuite) TestSetAndGetParams() {
	testCases := []struct {
		name    string
		input   types.Params
		expPass bool
	}{
		// it is not possible to set invalid booleans
		{"success: set params false", types.NewParams(false), true},
		{"success: set params true", types.NewParams(true), true},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			ctx := suite.chainA.GetContext()
			if tc.expPass {
				suite.chainA.GetSimApp().ICAControllerKeeper.SetParams(ctx, tc.input)
				expected := tc.input
				p := suite.chainA.GetSimApp().ICAControllerKeeper.GetParams(ctx)
				suite.Require().Equal(expected, p)
			} else { // currently not possible to set invalid params
				suite.Require().Panics(func() {
					suite.chainA.GetSimApp().ICAControllerKeeper.SetParams(ctx, tc.input)
				})
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUnsetParams() {
	suite.SetupTest()

	ctx := suite.chainA.GetContext()
	store := suite.chainA.GetContext().KVStore(suite.chainA.GetSimApp().GetKey(types.SubModuleName))
	store.Delete([]byte(types.ParamsKey))

	suite.Require().Panics(func() {
		suite.chainA.GetSimApp().ICAControllerKeeper.GetParams(ctx)
	})
}

func (suite *KeeperTestSuite) TestGetAuthority() {
	suite.SetupTest()

	authority := suite.chainA.GetSimApp().ICAControllerKeeper.GetAuthority()
	expectedAuth := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	suite.Require().Equal(expectedAuth, authority)
}
