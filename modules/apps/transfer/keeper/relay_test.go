package keeper_test

import (
	"fmt"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
)

// test sending from chainA to chainB using both coin that orignate on
// chainA and coin that orignate on chainB
func (s *KeeperTestSuite) TestSendTransfer() {
	var (
		coin            sdk.Coin
		path            *ibctesting.Path
		sender          sdk.AccAddress
		timeoutHeight   clienttypes.Height
		memo            string
		expEscrowAmount sdkmath.Int // total amount in escrow for denom on receiving chain
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful transfer with native token",
			func() {
				expEscrowAmount = sdkmath.NewInt(100)
			}, true,
		},
		{
			"successful transfer from source chain with memo",
			func() {
				memo = "memo" //nolint:goconst
				expEscrowAmount = sdkmath.NewInt(100)
			}, true,
		},
		{
			"successful transfer with IBC token",
			func() {
				// send IBC token back to chainB
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount)
			}, true,
		},
		{
			"successful transfer with IBC token and memo",
			func() {
				// send IBC token back to chainB
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount)
				memo = "memo"
			}, true,
		},
		{
			"source channel not found",
			func() {
				// channel references wrong ID
				path.EndpointA.ChannelID = ibctesting.InvalidID
			}, false,
		},
		{
			"transfer failed - sender account is blocked",
			func() {
				sender = s.chainA.GetSimApp().AccountKeeper.GetModuleAddress(types.ModuleName)
			}, false,
		},
		{
			"send coin failed",
			func() {
				coin = sdk.NewCoin("randomdenom", sdkmath.NewInt(100))
			}, false,
		},
		{
			"failed to parse coin denom",
			func() {
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, "randomdenom", coin.Amount)
			}, false,
		},
		{
			"send from module account failed, insufficient balance",
			func() {
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount.Add(sdkmath.NewInt(1)))
			}, false,
		},
		{
			"channel capability not found",
			func() {
				capability := s.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)

				// Release channel capability
				s.chainA.GetSimApp().ScopedTransferKeeper.ReleaseCapability(s.chainA.GetContext(), capability) //nolint:errcheck // ignore error for testing
			}, false,
		},
		{
			"SendPacket fails, timeout height and timeout timestamp are zero",
			func() {
				timeoutHeight = clienttypes.ZeroHeight()
				expEscrowAmount = sdkmath.NewInt(100)
			}, false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset

			path = ibctesting.NewTransferPath(s.chainA, s.chainB)
			s.coordinator.Setup(path)

			coin = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100))
			sender = s.chainA.SenderAccount.GetAddress()
			memo = ""
			timeoutHeight = s.chainB.GetTimeoutHeight()
			expEscrowAmount = sdkmath.ZeroInt()

			// create IBC token on chainA
			transferMsg := types.NewMsgTransfer(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, coin, s.chainB.SenderAccount.GetAddress().String(), s.chainA.SenderAccount.GetAddress().String(), s.chainA.GetTimeoutHeight(), 0, "")
			result, err := s.chainB.SendMsgs(transferMsg)
			s.Require().NoError(err) // message committed

			packet, err := ibctesting.ParsePacketFromEvents(result.Events)
			s.Require().NoError(err)

			err = path.RelayPacket(packet)
			s.Require().NoError(err)

			tc.malleate()

			msg := types.NewMsgTransfer(
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				coin, sender.String(), s.chainB.SenderAccount.GetAddress().String(),
				timeoutHeight, 0, // only use timeout height
				memo,
			)

			res, err := s.chainA.GetSimApp().TransferKeeper.Transfer(sdk.WrapSDKContext(s.chainA.GetContext()), msg)

			// check total amount in escrow of sent token denom on sending chain
			amount := s.chainA.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainA.GetContext(), coin.GetDenom())
			s.Require().Equal(expEscrowAmount, amount.Amount)

			if tc.expPass {
				s.Require().NoError(err)
				s.Require().NotNil(res)
			} else {
				s.Require().Error(err)
				s.Require().Nil(res)
			}
		})
	}
}

func (s *KeeperTestSuite) TestSendTransferSetsTotalEscrowAmountForSourceIBCToken() {
	/*
		Given the following flow of tokens:

		chain A (channel 0) -> (channel-0) chain B (channel-1) -> (channel-1) chain A
		stake                  transfer/channel-0/stake           transfer/channel-1/transfer/channel-0/stake
		                                  ^
		                                  |
		                             SendTransfer

		This test will transfer vouchers of denom "transfer/channel-0/stake" from chain B
		to chain A over channel-1 to assert that total escrow amount is stored on chain B
		for vouchers of denom "transfer/channel-0/stake" because chain B acts as source
		in this case.

		Set up:
		- Two transfer channels between chain A and chain B (channel-0 and channel-1).
		- Tokens of native denom "stake" on chain A transferred to chain B over channel-0
		and vouchers minted with denom trace "tranfer/channel-0/stake".

		Execute:
		- Transfer vouchers of denom trace "tranfer/channel-0/stake" from chain B to chain A
		over channel-1.

		Assert:
		- The vouchers are not of a native denom (because they are of an IBC denom), but chain B
		is the source, then the value for total escrow amount should still be stored for the IBC
		denom that corresponds to the trace "tranfer/channel-0/stake".
	*/

	// set up
	// 2 transfer channels between chain A and chain B
	path1 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path1)

	path2 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path2)

	// create IBC token on chain B with denom trace "transfer/channel-0/stake"
	coin := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100))
	transferMsg := types.NewMsgTransfer(
		path1.EndpointA.ChannelConfig.PortID,
		path1.EndpointA.ChannelID,
		coin,
		s.chainA.SenderAccount.GetAddress().String(),
		s.chainB.SenderAccount.GetAddress().String(),
		s.chainB.GetTimeoutHeight(), 0, "",
	)
	result, err := s.chainA.SendMsgs(transferMsg)
	s.Require().NoError(err) // message committed

	packet, err := ibctesting.ParsePacketFromEvents(result.Events)
	s.Require().NoError(err)

	err = path1.RelayPacket(packet)
	s.Require().NoError(err)

	// execute
	trace := types.ParseDenomTrace(types.GetPrefixedDenom(path1.EndpointB.ChannelConfig.PortID, path1.EndpointB.ChannelID, sdk.DefaultBondDenom))
	coin = sdk.NewCoin(trace.IBCDenom(), sdkmath.NewInt(100))
	msg := types.NewMsgTransfer(
		path2.EndpointB.ChannelConfig.PortID,
		path2.EndpointB.ChannelID,
		coin,
		s.chainB.SenderAccount.GetAddress().String(),
		s.chainA.SenderAccount.GetAddress().String(),
		s.chainA.GetTimeoutHeight(), 0, "",
	)

	res, err := s.chainB.GetSimApp().TransferKeeper.Transfer(sdk.WrapSDKContext(s.chainB.GetContext()), msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	// check total amount in escrow of sent token on sending chain
	totalEscrow := s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.NewInt(100), totalEscrow.Amount)
}

// test receiving coin on chainB with coin that orignate on chainA and
// coin that originated on chainB (source). The bulk of the testing occurs
// in the test case for loop since setup is intensive for all cases. The
// malleate function allows for testing invalid cases.
func (s *KeeperTestSuite) TestOnRecvPacket() {
	var (
		trace           types.DenomTrace
		amount          sdkmath.Int
		receiver        string
		memo            string
		expEscrowAmount sdkmath.Int // total amount in escrow for denom on receiving chain
	)

	testCases := []struct {
		msg          string
		malleate     func()
		recvIsSource bool // the receiving chain is the source of the coin originally
		expPass      bool
	}{
		{
			"success receive on source chain",
			func() {}, true, true,
		},
		{
			"success receive on source chain of half the amount",
			func() {
				amount = sdkmath.NewInt(50)
				expEscrowAmount = sdkmath.NewInt(50)
			}, true, true,
		},
		{
			"success receive on source chain with memo",
			func() {
				memo = "memo"
			}, true, true,
		},
		{
			"success receive with coin from another chain as source",
			func() {}, false, true,
		},
		{
			"success receive with coin from another chain as source with memo",
			func() {
				memo = "memo"
			}, false, true,
		},
		{
			"empty coin",
			func() {
				trace = types.DenomTrace{}
				amount = sdkmath.ZeroInt()
				expEscrowAmount = sdkmath.NewInt(100)
			}, true, false,
		},
		{
			"invalid receiver address",
			func() {
				receiver = "gaia1scqhwpgsmr6vmztaa7suurfl52my6nd2kmrudl"
				expEscrowAmount = sdkmath.NewInt(100)
			}, true, false,
		},

		// onRecvPacket
		// - coin from chain chainA
		{
			"failure: mint zero coin",
			func() {
				amount = sdkmath.ZeroInt()
			}, false, false,
		},

		// - coin being sent back to original chain (chainB)
		{
			"tries to unescrow more tokens than allowed",
			func() {
				amount = sdkmath.NewInt(1000000)
				expEscrowAmount = sdkmath.NewInt(100)
			}, true, false,
		},

		// - coin being sent to module address on chainA
		{
			"failure: receive on module account",
			func() {
				receiver = s.chainA.GetSimApp().AccountKeeper.GetModuleAddress(types.ModuleName).String()
			}, false, false,
		},

		// - coin being sent back to original chain (chainB) to module address
		{
			"failure: receive on module account on source chain",
			func() {
				receiver = s.chainB.GetSimApp().AccountKeeper.GetModuleAddress(types.ModuleName).String()
				expEscrowAmount = sdkmath.NewInt(100)
			}, true, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			s.SetupTest() // reset

			path := ibctesting.NewTransferPath(s.chainA, s.chainB)
			s.coordinator.Setup(path)
			receiver = s.chainB.SenderAccount.GetAddress().String() // must be explicitly changed in malleate

			memo = ""                           // can be explicitly changed in malleate
			amount = sdkmath.NewInt(100)        // must be explicitly changed in malleate
			expEscrowAmount = sdkmath.ZeroInt() // total amount in escrow of voucher denom on receiving chain
			seq := uint64(1)

			if tc.recvIsSource {
				// send coin from chainB to chainA, receive them, acknowledge them, and send back to chainB
				coinFromBToA := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100))
				transferMsg := types.NewMsgTransfer(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, coinFromBToA, s.chainB.SenderAccount.GetAddress().String(), s.chainA.SenderAccount.GetAddress().String(), clienttypes.NewHeight(1, 110), 0, memo)
				res, err := s.chainB.SendMsgs(transferMsg)
				s.Require().NoError(err) // message committed

				packet, err := ibctesting.ParsePacketFromEvents(res.Events)
				s.Require().NoError(err)

				err = path.RelayPacket(packet)
				s.Require().NoError(err) // relay committed

				seq++

				// NOTE: trace must be explicitly changed in malleate to test invalid cases
				trace = types.ParseDenomTrace(types.GetPrefixedDenom(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.DefaultBondDenom))
			} else {
				trace = types.ParseDenomTrace(sdk.DefaultBondDenom)
			}

			// send coin from chainA to chainB
			coin := sdk.NewCoin(trace.IBCDenom(), amount)
			transferMsg := types.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin, s.chainA.SenderAccount.GetAddress().String(), receiver, clienttypes.NewHeight(1, 110), 0, memo)
			_, err := s.chainA.SendMsgs(transferMsg)
			s.Require().NoError(err) // message committed

			tc.malleate()

			data := types.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), s.chainA.SenderAccount.GetAddress().String(), receiver, memo)
			packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)

			err = s.chainB.GetSimApp().TransferKeeper.OnRecvPacket(s.chainB.GetContext(), packet, data)

			// check total amount in escrow of received token denom on receiving chain
			var (
				denom       string
				totalEscrow sdk.Coin
			)
			if tc.recvIsSource {
				denom = sdk.DefaultBondDenom
			} else {
				denom = trace.IBCDenom()
			}
			totalEscrow = s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), denom)
			s.Require().Equal(expEscrowAmount, totalEscrow.Amount)

			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestOnRecvPacketSetsTotalEscrowAmountForSourceIBCToken() {
	/*
		Given the following flow of tokens:

		chain A (channel 0) -> (channel-0) chain B (channel-1) -> (channel-1) chain A (channel-1)             -> (channel-1) chain B
		stake                  transfer/channel-0/stake           transfer/channel-1/transfer/channel-0/stake    transfer/channel-0/stake
		                                                                                                                   ^
		                                                                                                                   |
		                                                                                                              OnRecvPacket

		This test will assert that on receiving vouchers of denom "transfer/channel-0/stake"
		on chain B the total escrow amount is updated on because chain B acted as source
		when vouchers were transferred to chain A over channel-1.

		Setup:
		- Two transfer channels between chain A and chain B.
		- Vouchers of denom trace "transfer/channel-0/stake" on chain B are in escrow
		account for port ID transfer and channel ID channel-1.

		Execute:
		- Receive vouchers of denom trace "transfer/channel-0/stake" from chain A to chain B
		over channel-1.

		Assert:
		- The vouchers are not of a native denom (because they are of an IBC denom), but chain B
		is the source, then the value for total escrow amount should still be updated for the IBC
		denom that corresponds to the trace "tranfer/channel-0/stake" when the vouchers are
		received back on chain B.
	*/

	seq := uint64(1)
	amount := sdkmath.NewInt(100)
	timeout := s.chainA.GetTimeoutHeight()

	// setup
	// 2 transfer channels between chain A and chain B
	path1 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path1)

	path2 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path2)

	// denomTrace path: {transfer/channel-1/transfer/channel-0}
	denomTrace := types.DenomTrace{
		BaseDenom: sdk.DefaultBondDenom,
		Path:      fmt.Sprintf("%s/%s/%s/%s", path2.EndpointA.ChannelConfig.PortID, path2.EndpointA.ChannelID, path1.EndpointB.ChannelConfig.PortID, path1.EndpointB.ChannelID),
	}
	data := types.NewFungibleTokenPacketData(
		denomTrace.GetFullDenomPath(),
		amount.String(),
		s.chainA.SenderAccount.GetAddress().String(),
		s.chainB.SenderAccount.GetAddress().String(), "",
	)
	packet := channeltypes.NewPacket(
		data.GetBytes(),
		seq,
		path2.EndpointA.ChannelConfig.PortID,
		path2.EndpointA.ChannelID,
		path2.EndpointB.ChannelConfig.PortID,
		path2.EndpointB.ChannelID,
		timeout, 0,
	)

	// fund escrow account for transfer and channel-1 on chain B
	// denomTrace path: transfer/channel-0
	denomTrace = types.DenomTrace{
		BaseDenom: sdk.DefaultBondDenom,
		Path:      fmt.Sprintf("%s/%s", path1.EndpointB.ChannelConfig.PortID, path1.EndpointB.ChannelID),
	}
	escrowAddress := types.GetEscrowAddress(path2.EndpointB.ChannelConfig.PortID, path2.EndpointB.ChannelID)
	coin := sdk.NewCoin(denomTrace.IBCDenom(), amount)
	s.Require().NoError(
		banktestutil.FundAccount(
			s.chainB.GetSimApp().BankKeeper,
			s.chainB.GetContext(),
			escrowAddress,
			sdk.NewCoins(coin),
		),
	)

	s.chainB.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainB.GetContext(), coin)
	totalEscrowChainB := s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.NewInt(100), totalEscrowChainB.Amount)

	// execute onRecvPacket, when chaninB receives the source token the escrow amount should decrease
	err := s.chainB.GetSimApp().TransferKeeper.OnRecvPacket(s.chainB.GetContext(), packet, data)
	s.Require().NoError(err)

	// check total amount in escrow of sent token on reveiving chain
	totalEscrowChainB = s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.ZeroInt(), totalEscrowChainB.Amount)
}

// TestOnAcknowledgementPacket tests that successful acknowledgement is a no-op
// and failure acknowledment leads to refund when attempting to send from chainA
// to chainB. If sender is source then the denomination being refunded has no
// trace
func (s *KeeperTestSuite) TestOnAcknowledgementPacket() {
	var (
		successAck      = channeltypes.NewResultAcknowledgement([]byte{byte(1)})
		failedAck       = channeltypes.NewErrorAcknowledgement(fmt.Errorf("failed packet transfer"))
		trace           types.DenomTrace
		amount          sdkmath.Int
		path            *ibctesting.Path
		expEscrowAmount sdkmath.Int
	)

	testCases := []struct {
		msg      string
		ack      channeltypes.Acknowledgement
		malleate func()
		success  bool // success of ack
		expPass  bool
	}{
		{
			"success ack causes no-op",
			successAck,
			func() {
				trace = types.ParseDenomTrace(types.GetPrefixedDenom(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, sdk.DefaultBondDenom))
			}, true, true,
		},
		{
			"successful refund from source chain",
			failedAck,
			func() {
				escrow := types.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				trace = types.ParseDenomTrace(sdk.DefaultBondDenom)
				coin := sdk.NewCoin(sdk.DefaultBondDenom, amount)

				s.Require().NoError(banktestutil.FundAccount(s.chainA.GetSimApp().BankKeeper, s.chainA.GetContext(), escrow, sdk.NewCoins(coin)))

				// set escrow amount that would have been stored after successful execution of MsgTransfer
				s.chainA.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainA.GetContext(), sdk.NewCoin(sdk.DefaultBondDenom, amount))
			}, false, true,
		},
		{
			"unsuccessful refund from source",
			failedAck,
			func() {
				trace = types.ParseDenomTrace(sdk.DefaultBondDenom)

				// set escrow amount that would have been stored after successful execution of MsgTransfer
				s.chainA.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainA.GetContext(), sdk.NewCoin(sdk.DefaultBondDenom, amount))
				expEscrowAmount = sdkmath.NewInt(100)
			}, false, false,
		},
		{
			"successful refund with coin from external chain",
			failedAck,
			func() {
				escrow := types.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				trace = types.ParseDenomTrace(types.GetPrefixedDenom(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.DefaultBondDenom))
				coin := sdk.NewCoin(trace.IBCDenom(), amount)

				s.Require().NoError(banktestutil.FundAccount(s.chainA.GetSimApp().BankKeeper, s.chainA.GetContext(), escrow, sdk.NewCoins(coin)))
			}, false, true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			s.SetupTest() // reset
			path = ibctesting.NewTransferPath(s.chainA, s.chainB)
			s.coordinator.Setup(path)
			amount = sdkmath.NewInt(100) // must be explicitly changed
			expEscrowAmount = sdkmath.ZeroInt()

			tc.malleate()

			data := types.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), s.chainA.SenderAccount.GetAddress().String(), s.chainB.SenderAccount.GetAddress().String(), "")
			packet := channeltypes.NewPacket(data.GetBytes(), 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)
			preCoin := s.chainA.GetSimApp().BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccount.GetAddress(), trace.IBCDenom())

			err := s.chainA.GetSimApp().TransferKeeper.OnAcknowledgementPacket(s.chainA.GetContext(), packet, data, tc.ack)

			// check total amount in escrow of sent token denom on sending chain
			totalEscrow := s.chainA.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainA.GetContext(), trace.IBCDenom())
			s.Require().Equal(expEscrowAmount, totalEscrow.Amount)

			if tc.expPass {
				s.Require().NoError(err)
				postCoin := s.chainA.GetSimApp().BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccount.GetAddress(), trace.IBCDenom())
				deltaAmount := postCoin.Amount.Sub(preCoin.Amount)

				if tc.success {
					s.Require().Equal(int64(0), deltaAmount.Int64(), "successful ack changed balance")
				} else {
					s.Require().Equal(amount, deltaAmount, "failed ack did not trigger refund")
				}
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestOnAcknowledgementPacketSetsTotalEscrowAmountForSourceIBCToken() {
	/*
		This test is testing the following scenario. Given tokens travelling like this:

		chain A (channel 0) -> (channel-0) chain B (channel-1) -> (channel-1) chain A (channel-1)
		stake                  transfer/channel-0/stake           transfer/channel-1/transfer/channel-0/stake
		                                 ^
		                                 |
		                         OnAcknowledgePacket

		We want to assert that on failed acknowledgment of vouchers sent with denom trace
		"transfer/channel-0/stake" on chain B the total escrow amount is updated.

		Set up:
		- Two transfer channels between chain A and chain B.
		- Vouckers of denom "transfer/channel-0/stake" on chain B are in escrow
		account for port ID transfer and channel ID channel-1.

		Execute:
		- Acknowledge vouchers of denom trace "tranfer/channel-0/stake" sent from chain B
		to chain B over channel-1.

		Assert:
		- The vouchers are not of a native denom (because they are of an IBC denom), but chain B
		is the source, then the value for total escrow amount should still be updated for the IBC
		denom that corresponds to the trace "tranfer/channel-0/stake" when processing the failed
		acknowledgement.
	*/

	seq := uint64(1)
	amount := sdkmath.NewInt(100)
	ack := channeltypes.NewErrorAcknowledgement(fmt.Errorf("failed packet transfer"))

	// set up
	// 2 transfer channels between chain A and chain B
	path1 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path1)

	path2 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path2)

	// fund escrow account for transfer and channel-1 on chain B
	// denomTrace path = transfer/channel-0
	denomTrace := types.DenomTrace{
		BaseDenom: sdk.DefaultBondDenom,
		Path:      fmt.Sprintf("%s/%s", path1.EndpointB.ChannelConfig.PortID, path1.EndpointB.ChannelID),
	}
	escrowAddress := types.GetEscrowAddress(path2.EndpointB.ChannelConfig.PortID, path2.EndpointB.ChannelID)
	coin := sdk.NewCoin(denomTrace.IBCDenom(), amount)
	s.Require().NoError(
		banktestutil.FundAccount(
			s.chainB.GetSimApp().BankKeeper,
			s.chainB.GetContext(),
			escrowAddress,
			sdk.NewCoins(coin),
		),
	)

	data := types.NewFungibleTokenPacketData(
		denomTrace.GetFullDenomPath(),
		amount.String(),
		s.chainB.SenderAccount.GetAddress().String(),
		s.chainA.SenderAccount.GetAddress().String(), "",
	)
	packet := channeltypes.NewPacket(
		data.GetBytes(),
		seq,
		path2.EndpointB.ChannelConfig.PortID,
		path2.EndpointB.ChannelID,
		path2.EndpointA.ChannelConfig.PortID,
		path2.EndpointA.ChannelID,
		s.chainA.GetTimeoutHeight(), 0,
	)

	s.chainB.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainB.GetContext(), coin)
	totalEscrowChainB := s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.NewInt(100), totalEscrowChainB.Amount)

	err := s.chainB.GetSimApp().TransferKeeper.OnAcknowledgementPacket(s.chainB.GetContext(), packet, data, ack)
	s.Require().NoError(err)

	// check total amount in escrow of sent token on sending chain
	totalEscrowChainB = s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.ZeroInt(), totalEscrowChainB.Amount)
}

// TestOnTimeoutPacket test private refundPacket function since it is a simple
// wrapper over it. The actual timeout does not matter since IBC core logic
// is not being tested. The test is timing out a send from chainA to chainB
// so the refunds are occurring on chainA.
func (s *KeeperTestSuite) TestOnTimeoutPacket() {
	var (
		trace           types.DenomTrace
		path            *ibctesting.Path
		amount          sdkmath.Int
		sender          string
		expEscrowAmount sdkmath.Int
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"successful timeout from sender as source chain",
			func() {
				escrow := types.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				trace = types.ParseDenomTrace(sdk.DefaultBondDenom)
				coin := sdk.NewCoin(trace.IBCDenom(), amount)
				expEscrowAmount = sdkmath.ZeroInt()

				// funds the escrow account to have balance
				s.Require().NoError(banktestutil.FundAccount(s.chainA.GetSimApp().BankKeeper, s.chainA.GetContext(), escrow, sdk.NewCoins(coin)))
				// set escrow amount that would have been stored after successful execution of MsgTransfer
				s.chainA.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainA.GetContext(), coin)
			}, true,
		},
		{
			"successful timeout from external chain",
			func() {
				escrow := types.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				trace = types.ParseDenomTrace(types.GetPrefixedDenom(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.DefaultBondDenom))
				coin := sdk.NewCoin(trace.IBCDenom(), amount)
				expEscrowAmount = sdkmath.ZeroInt()

				// funds the escrow account to have balance
				s.Require().NoError(banktestutil.FundAccount(s.chainA.GetSimApp().BankKeeper, s.chainA.GetContext(), escrow, sdk.NewCoins(coin)))
			}, true,
		},
		{
			"no balance for coin denom",
			func() {
				trace = types.ParseDenomTrace("bitcoin")
				expEscrowAmount = amount

				// set escrow amount that would have been stored after successful execution of MsgTransfer
				s.chainA.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainA.GetContext(), sdk.NewCoin(trace.IBCDenom(), amount))
			}, false,
		},
		{
			"unescrow failed",
			func() {
				trace = types.ParseDenomTrace(sdk.DefaultBondDenom)
				expEscrowAmount = amount

				// set escrow amount that would have been stored after successful execution of MsgTransfer
				s.chainA.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainA.GetContext(), sdk.NewCoin(trace.IBCDenom(), amount))
			}, false,
		},
		{
			"mint failed",
			func() {
				trace = types.ParseDenomTrace(types.GetPrefixedDenom(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.DefaultBondDenom))
				amount = sdkmath.OneInt()
				sender = "invalid address"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			s.SetupTest() // reset

			path = ibctesting.NewTransferPath(s.chainA, s.chainB)
			s.coordinator.Setup(path)
			amount = sdkmath.NewInt(100) // must be explicitly changed
			sender = s.chainA.SenderAccount.GetAddress().String()
			expEscrowAmount = sdkmath.ZeroInt()

			tc.malleate()

			data := types.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), sender, s.chainB.SenderAccount.GetAddress().String(), "")
			packet := channeltypes.NewPacket(data.GetBytes(), 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)
			preCoin := s.chainA.GetSimApp().BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccount.GetAddress(), trace.IBCDenom())

			err := s.chainA.GetSimApp().TransferKeeper.OnTimeoutPacket(s.chainA.GetContext(), packet, data)

			postCoin := s.chainA.GetSimApp().BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccount.GetAddress(), trace.IBCDenom())
			deltaAmount := postCoin.Amount.Sub(preCoin.Amount)

			// check total amount in escrow of sent token denom on sending chain
			totalEscrow := s.chainA.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainA.GetContext(), trace.IBCDenom())
			s.Require().Equal(expEscrowAmount, totalEscrow.Amount)

			if tc.expPass {
				s.Require().NoError(err)
				s.Require().Equal(amount.Int64(), deltaAmount.Int64(), "successful timeout did not trigger refund")
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestOnTimeoutPacketSetsTotalEscrowAmountForSourceIBCToken() {
	/*
		Given the following flow of tokens:

		chain A (channel 0) -> (channel-0) chain B (channel-1) -> (channel-1) chain A (channel-1)
		stake                  transfer/channel-0/stake           transfer/channel-1/transfer/channel-0/stake
		                                 ^
		                                 |
		                           OnTimeoutPacket

		We want to assert that on timeout of vouchers sent with denom trace
		"transfer/channel-0/stake" on chain B the total escrow amount is updated.

		Set up:
		- Two transfer channels between chain A and chain B.
		- Vouckers of denom "transfer/channel-0/stake" on chain B are in escrow
		account for port ID transfer and channel ID channel-1.

		Execute:
		- Timeout vouchers of denom trace "tranfer/channel-0/stake" sent from chain B
		to chain B over channel-1.

		Assert:
		- The vouchers are not of a native denom (because they are of an IBC denom), but chain B
		is the source, then the value for total escrow amount should still be updated for the IBC
		denom that corresponds to the trace "tranfer/channel-0/stake" when processing the timeout.
	*/

	seq := uint64(1)
	amount := sdkmath.NewInt(100)

	// set up
	// 2 transfer channels between chain A and chain B
	path1 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path1)

	path2 := ibctesting.NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(path2)

	// fund escrow account for transfer and channel-1 on chain B
	denomTrace := types.DenomTrace{
		BaseDenom: sdk.DefaultBondDenom,
		Path:      fmt.Sprintf("%s/%s", path1.EndpointB.ChannelConfig.PortID, path1.EndpointB.ChannelID),
	}
	escrowAddress := types.GetEscrowAddress(path2.EndpointB.ChannelConfig.PortID, path2.EndpointB.ChannelID)
	coin := sdk.NewCoin(denomTrace.IBCDenom(), amount)
	s.Require().NoError(
		banktestutil.FundAccount(
			s.chainB.GetSimApp().BankKeeper,
			s.chainB.GetContext(),
			escrowAddress,
			sdk.NewCoins(coin),
		),
	)

	data := types.NewFungibleTokenPacketData(
		denomTrace.GetFullDenomPath(),
		amount.String(),
		s.chainB.SenderAccount.GetAddress().String(),
		s.chainA.SenderAccount.GetAddress().String(), "",
	)
	packet := channeltypes.NewPacket(
		data.GetBytes(),
		seq,
		path2.EndpointB.ChannelConfig.PortID,
		path2.EndpointB.ChannelID,
		path2.EndpointA.ChannelConfig.PortID,
		path2.EndpointA.ChannelID,
		s.chainA.GetTimeoutHeight(), 0,
	)

	s.chainB.GetSimApp().TransferKeeper.SetTotalEscrowForDenom(s.chainB.GetContext(), coin)
	totalEscrowChainB := s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.NewInt(100), totalEscrowChainB.Amount)

	err := s.chainB.GetSimApp().TransferKeeper.OnTimeoutPacket(s.chainB.GetContext(), packet, data)
	s.Require().NoError(err)

	// check total amount in escrow of sent token on sending chain
	totalEscrowChainB = s.chainB.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(s.chainB.GetContext(), coin.GetDenom())
	s.Require().Equal(sdkmath.ZeroInt(), totalEscrowChainB.Amount)
}
