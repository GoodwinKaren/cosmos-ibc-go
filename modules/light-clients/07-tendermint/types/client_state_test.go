package types_test

import (
	"time"

	ics23 "github.com/confio/ics23/go"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	ibctesting "github.com/cosmos/ibc-go/testing"
	ibcmock "github.com/cosmos/ibc-go/testing/mock"
)

const (
	testClientID     = "clientidone"
	testConnectionID = "connectionid"
	testPortID       = "testportid"
	testChannelID    = "testchannelid"
	testSequence     = 1
	longChainID      = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

var (
	invalidProof = []byte("invalid proof")
)

func (suite *TendermintTestSuite) TestValidate() {
	testCases := []struct {
		name        string
		clientState *types.ClientState
		expPass     bool
	}{
		{
			name:        "valid client",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     true,
		},
		{
			name:        "valid client with nil upgrade path",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), nil, false, false),
			expPass:     true,
		},
		{
			name:        "invalid chainID",
			clientState: types.NewClientState("  ", types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid chainID - chainID is above maximum character length",
			clientState: types.NewClientState(longChainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid trust level",
			clientState: types.NewClientState(chainID, types.Fraction{Numerator: 0, Denominator: 1}, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid trusting period",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, 0, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid unbonding period",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, 0, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid max clock drift",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, 0, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "invalid height",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, clienttypes.ZeroHeight(), commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "trusting period not less than unbonding period",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "proof specs is nil",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, nil, upgradePath, false, false),
			expPass:     false,
		},
		{
			name:        "proof specs contains nil",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, []*ics23.ProofSpec{ics23.TendermintSpec, nil}, upgradePath, false, false),
			expPass:     false,
		},
	}

	for _, tc := range testCases {
		err := tc.clientState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

func (suite *TendermintTestSuite) TestInitialize() {

	testCases := []struct {
		name           string
		consensusState exported.ConsensusState
		expPass        bool
	}{
		{
			name:           "valid consensus",
			consensusState: &types.ConsensusState{},
			expPass:        true,
		},
		{
			name:           "invalid consensus: consensus state is solomachine consensus",
			consensusState: ibctesting.NewSolomachine(suite.T(), suite.chainA.Codec, "solomachine", "", 2).ConsensusState(),
			expPass:        false,
		},
	}

	path := ibctesting.NewPath(suite.chainA, suite.chainB)
	err := path.EndpointA.CreateClient()
	suite.Require().NoError(err)

	clientState := suite.chainA.GetClientState(path.EndpointA.ClientID)
	store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

	for _, tc := range testCases {
		err := clientState.Initialize(suite.chainA.GetContext(), suite.chainA.Codec, store, tc.consensusState)
		if tc.expPass {
			suite.Require().NoError(err, "valid case returned an error")
		} else {
			suite.Require().Error(err, "invalid case didn't return an error")
		}
	}
}

func (suite *TendermintTestSuite) TestVerifyClientConsensusState() {
	testCases := []struct {
		name           string
		clientState    *types.ClientState
		consensusState *types.ConsensusState
		prefix         commitmenttypes.MerklePrefix
		proof          []byte
		expPass        bool
	}{
		// FIXME: uncomment
		// {
		// 	name:        "successful verification",
		// 	clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height,  commitmenttypes.GetSDKSpecs()),
		// 	consensusState: types.ConsensusState{
		// 		Root: commitmenttypes.NewMerkleRoot(suite.header.Header.GetAppHash()),
		// 	},
		// 	prefix:  commitmenttypes.NewMerklePrefix([]byte("ibc")),
		// 	expPass: true,
		// },
		{
			name:        "ApplyPrefix failed",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			consensusState: &types.ConsensusState{
				Root: commitmenttypes.NewMerkleRoot(suite.header.Header.GetAppHash()),
			},
			prefix:  commitmenttypes.MerklePrefix{},
			expPass: false,
		},
		{
			name:        "latest client height < height",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			consensusState: &types.ConsensusState{
				Root: commitmenttypes.NewMerkleRoot(suite.header.Header.GetAppHash()),
			},
			prefix:  commitmenttypes.NewMerklePrefix([]byte("ibc")),
			expPass: false,
		},
		{
			name:        "client is frozen",
			clientState: &types.ClientState{LatestHeight: height, FrozenHeight: clienttypes.NewHeight(height.RevisionNumber, height.RevisionHeight-1)},
			consensusState: &types.ConsensusState{
				Root: commitmenttypes.NewMerkleRoot(suite.header.Header.GetAppHash()),
			},
			prefix:  commitmenttypes.NewMerklePrefix([]byte("ibc")),
			expPass: false,
		},
		{
			name:        "proof verification failed",
			clientState: types.NewClientState(chainID, types.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false),
			consensusState: &types.ConsensusState{
				Root:               commitmenttypes.NewMerkleRoot(suite.header.Header.GetAppHash()),
				NextValidatorsHash: suite.valsHash,
			},
			prefix:  commitmenttypes.NewMerklePrefix([]byte("ibc")),
			proof:   []byte{},
			expPass: false,
		},
	}

	for i, tc := range testCases {
		tc := tc

		err := tc.clientState.VerifyClientConsensusState(
			nil, suite.cdc, height, "chainA", tc.clientState.LatestHeight, tc.prefix, tc.proof, tc.consensusState,
		)

		if tc.expPass {
			suite.Require().NoError(err, "valid test case %d failed: %s", i, tc.name)
		} else {
			suite.Require().Error(err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}

// test verification of the connection on chainB being represented in the
// light client on chainA
func (suite *TendermintTestSuite) TestVerifyConnectionState() {
	var (
		clientState *types.ClientState
		proof       []byte
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupNew(path)
			connection := path.EndpointB.GetConnection()

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make connection proof
			connectionKey := host.ConnectionKey(path.EndpointB.ConnectionID)
			proof, proofHeight = suite.chainB.QueryProof(connectionKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			err := clientState.VerifyConnectionState(
				store, suite.chainA.Codec, proofHeight, &prefix, proof, path.EndpointB.ConnectionID, connection,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// test verification of the channel on chainB being represented in the light
// client on chainA
func (suite *TendermintTestSuite) TestVerifyChannelState() {
	var (
		clientState *types.ClientState
		proof       []byte
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupNew(path)
			channel := path.EndpointB.GetChannel()

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make channel proof
			channelKey := host.ChannelKey(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			proof, proofHeight = suite.chainB.QueryProof(channelKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			err := clientState.VerifyChannelState(
				store, suite.chainA.Codec, proofHeight, &prefix, proof,
				path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, channel,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// test verification of the packet commitment on chainB being represented
// in the light client on chainA. A send from chainB to chainA is simulated.
func (suite *TendermintTestSuite) TestVerifyPacketCommitment() {
	var (
		clientState *types.ClientState
		proof       []byte
		delayPeriod uint64
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			name: "delay period has passed",
			malleate: func() {
				delayPeriod = uint64(time.Second.Nanoseconds())
			},
			expPass: true,
		},
		{
			name: "delay period has not passed",
			malleate: func() {
				delayPeriod = uint64(time.Hour.Nanoseconds())
			},
			expPass: false,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupNew(path)
			packet := channeltypes.NewPacket(ibctesting.MockPacketData, 1, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, clienttypes.NewHeight(0, 100), 0)
			err := suite.coordinator.SendPacket(suite.chainB, suite.chainA, packet, path.EndpointA.ClientID)
			suite.Require().NoError(err)

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make packet commitment proof
			packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
			proof, proofHeight = suite.chainB.QueryProof(packetKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			currentTime := uint64(suite.chainA.GetContext().BlockTime().UnixNano())
			commitment := channeltypes.CommitPacket(suite.chainA.App.IBCKeeper.Codec(), packet)
			err = clientState.VerifyPacketCommitment(
				store, suite.chainA.Codec, proofHeight, currentTime, delayPeriod, &prefix, proof,
				packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence(), commitment,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// test verification of the acknowledgement on chainB being represented
// in the light client on chainA. A send and ack from chainA to chainB
// is simulated.
func (suite *TendermintTestSuite) TestVerifyPacketAcknowledgement() {
	var (
		clientState *types.ClientState
		proof       []byte
		delayPeriod uint64
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			name: "delay period has passed",
			malleate: func() {
				delayPeriod = uint64(time.Second.Nanoseconds())
			},
			expPass: true,
		},
		{
			name: "delay period has not passed",
			malleate: func() {
				delayPeriod = uint64(time.Hour.Nanoseconds())
			},
			expPass: false,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupNew(path)
			packet := channeltypes.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(0, 100), 0)

			// send packet
			err := suite.coordinator.SendPacket(suite.chainA, suite.chainB, packet, path.EndpointB.ClientID)
			suite.Require().NoError(err)

			// write receipt and ack
			err = suite.coordinator.RecvPacket(suite.chainA, suite.chainB, path.EndpointA.ClientID, packet)
			suite.Require().NoError(err)

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make packet acknowledgement proof
			acknowledgementKey := host.PacketAcknowledgementKey(packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence())
			proof, proofHeight = suite.chainB.QueryProof(acknowledgementKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			currentTime := uint64(suite.chainA.GetContext().BlockTime().UnixNano())
			err = clientState.VerifyPacketAcknowledgement(
				store, suite.chainA.Codec, proofHeight, currentTime, delayPeriod, &prefix, proof,
				packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence(), ibcmock.MockAcknowledgement.Acknowledgement(),
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// test verification of the absent acknowledgement on chainB being represented
// in the light client on chainA. A send from chainB to chainA is simulated, but
// no receive.
func (suite *TendermintTestSuite) TestVerifyPacketReceiptAbsence() {
	var (
		clientState *types.ClientState
		proof       []byte
		delayPeriod uint64
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			name: "delay period has passed",
			malleate: func() {
				delayPeriod = uint64(time.Second.Nanoseconds())
			},
			expPass: true,
		},
		{
			name: "delay period has not passed",
			malleate: func() {
				delayPeriod = uint64(time.Hour.Nanoseconds())
			},
			expPass: false,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupNew(path)
			packet := channeltypes.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(0, 100), 0)

			// send packet, but no recv
			err := suite.coordinator.SendPacket(suite.chainA, suite.chainB, packet, path.EndpointB.ClientID)
			suite.Require().NoError(err)

			// need to update chainA's client representing chainB to prove missing ack
			suite.coordinator.UpdateClient(suite.chainA, suite.chainB, path.EndpointA.ClientID, exported.Tendermint)

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make packet receipt absence proof
			receiptKey := host.PacketReceiptKey(packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence())
			proof, proofHeight = suite.chainB.QueryProof(receiptKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			currentTime := uint64(suite.chainA.GetContext().BlockTime().UnixNano())
			err = clientState.VerifyPacketReceiptAbsence(
				store, suite.chainA.Codec, proofHeight, currentTime, delayPeriod, &prefix, proof,
				packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence(),
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// test verification of the next receive sequence on chainB being represented
// in the light client on chainA. A send and receive from chainB to chainA is
// simulated.
func (suite *TendermintTestSuite) TestVerifyNextSeqRecv() {
	var (
		clientState *types.ClientState
		proof       []byte
		delayPeriod uint64
		proofHeight exported.Height
		prefix      commitmenttypes.MerklePrefix
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful verification", func() {}, true,
		},
		{
			name: "delay period has passed",
			malleate: func() {
				delayPeriod = uint64(time.Second.Nanoseconds())
			},
			expPass: true,
		},
		{
			name: "delay period has not passed",
			malleate: func() {
				delayPeriod = uint64(time.Hour.Nanoseconds())
			},
			expPass: false,
		},
		{
			"ApplyPrefix failed", func() {
				prefix = commitmenttypes.MerklePrefix{}
			}, false,
		},
		{
			"latest client height < height", func() {
				proofHeight = clientState.LatestHeight.Increment()
			}, false,
		},
		{
			"client is frozen", func() {
				clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			}, false,
		},
		{
			"proof verification failed", func() {
				proof = invalidProof
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			// setup testing conditions
			path := ibctesting.NewPath(suite.chainA, suite.chainB)
			path.SetChannelOrdered()
			suite.coordinator.SetupNew(path)
			packet := channeltypes.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(0, 100), 0)

			// send packet
			err := path.EndpointA.SendPacket(packet)
			suite.Require().NoError(err)

			// next seq recv incremented
			err = path.EndpointB.RecvPacket(packet)
			suite.Require().NoError(err)

			var ok bool
			clientStateI := suite.chainA.GetClientState(path.EndpointA.ClientID)
			clientState, ok = clientStateI.(*types.ClientState)
			suite.Require().True(ok)

			prefix = suite.chainB.GetPrefix()

			// make next seq recv proof
			nextSeqRecvKey := host.NextSequenceRecvKey(packet.GetDestPort(), packet.GetDestChannel())
			proof, proofHeight = suite.chainB.QueryProof(nextSeqRecvKey)

			tc.malleate() // make changes as necessary

			store := suite.chainA.App.IBCKeeper.ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)

			currentTime := uint64(suite.chainA.GetContext().BlockTime().UnixNano())
			err = clientState.VerifyNextSequenceRecv(
				store, suite.chainA.Codec, proofHeight, currentTime, delayPeriod, &prefix, proof,
				packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence()+1,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
