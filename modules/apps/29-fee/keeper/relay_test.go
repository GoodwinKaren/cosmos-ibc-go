package keeper_test

import (
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

func (suite *KeeperTestSuite) TestWriteAcknowledgementAsync() {
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				suite.chainB.GetSimApp().IBCFeeKeeper.SetRelayerAddressForAsyncAck(suite.chainB.GetContext(), channeltypes.NewPacketId(suite.path.EndpointA.ChannelID, suite.path.EndpointA.ChannelConfig.PortID, 1), suite.chainA.SenderAccount.GetAddress().String())
				suite.chainB.GetSimApp().IBCFeeKeeper.SetCounterpartyAddress(suite.chainB.GetContext(), suite.chainA.SenderAccount.GetAddress().String(), suite.chainB.SenderAccount.GetAddress().String())
			},
			true,
		},
		{
			"relayer address not set for async WriteAcknowledgement",
			func() {},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// open incentivized channel
			suite.coordinator.Setup(suite.path)

			// build packet
			timeoutTimestamp := ^uint64(0)
			packet := channeltypes.NewPacket(
				[]byte("packetData"),
				1,
				suite.path.EndpointA.ChannelConfig.PortID,
				suite.path.EndpointA.ChannelID,
				suite.path.EndpointB.ChannelConfig.PortID,
				suite.path.EndpointB.ChannelID,
				clienttypes.ZeroHeight(),
				timeoutTimestamp,
			)

			ack := channeltypes.NewResultAcknowledgement([]byte("success"))
			chanCap := suite.chainB.GetChannelCapability(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID)

			// malleate test case
			tc.malleate()

			err := suite.chainB.GetSimApp().IBCFeeKeeper.WriteAcknowledgement(suite.chainB.GetContext(), chanCap, packet, ack)

			if tc.expPass {
				suite.Require().NoError(err)
				_, found := suite.chainB.GetSimApp().IBCFeeKeeper.GetRelayerAddressForAsyncAck(suite.chainB.GetContext(), channeltypes.NewPacketId(suite.path.EndpointA.ChannelID, suite.path.EndpointA.ChannelConfig.PortID, 1))
				suite.Require().False(found)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestWriteAcknowledgementAsyncFeeDisabled() {
	// open incentivized channel
	suite.coordinator.Setup(suite.path)
	suite.chainB.GetSimApp().IBCFeeKeeper.DeleteFeeEnabled(suite.chainB.GetContext(), suite.path.EndpointB.ChannelConfig.PortID, "channel-0")

	// build packet
	timeoutTimestamp := ^uint64(0)
	packet := channeltypes.NewPacket(
		[]byte("packetData"),
		1,
		suite.path.EndpointA.ChannelConfig.PortID,
		suite.path.EndpointA.ChannelID,
		suite.path.EndpointB.ChannelConfig.PortID,
		suite.path.EndpointB.ChannelID,
		clienttypes.ZeroHeight(),
		timeoutTimestamp,
	)

	ack := channeltypes.NewResultAcknowledgement([]byte("success"))
	chanCap := suite.chainB.GetChannelCapability(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID)

	err := suite.chainB.GetSimApp().IBCFeeKeeper.WriteAcknowledgement(suite.chainB.GetContext(), chanCap, packet, ack)
	suite.Require().NoError(err)

	packetAck, _ := suite.chainB.GetSimApp().GetIBCKeeper().ChannelKeeper.GetPacketAcknowledgement(suite.chainB.GetContext(), packet.DestinationPort, packet.DestinationChannel, 1)
	suite.Require().Equal(packetAck, channeltypes.CommitAcknowledgement(ack.Acknowledgement()))
}
