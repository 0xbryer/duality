package keeper_test

import (
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = suite.TestingSuite(nil)

func (suite *KeeperTestSuite) TestStakeLifecycle() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpec: depositSpec{
			addr:   addr0,
			token0: sdk.NewInt64Coin("TokenA", 10),
			token1: sdk.NewInt64Coin("TokenB", 10),
			tick:   0,
			fee:    1,
		},
		stakeTimeOffset: -24 * time.Hour,
	})

	retrievedStake, err := suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)

	// unstake the full amount
	suite.App.IncentivesKeeper.Unstake(suite.Ctx, stake, sdk.Coins{})
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 20)), balances)
	_, err = suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	// should be deleted
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestStakeUnstakePartial() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpec: depositSpec{
			addr:   addr0,
			token0: sdk.NewInt64Coin("TokenA", 10),
			token1: sdk.NewInt64Coin("TokenB", 10),
			tick:   0,
			fee:    1,
		},
		stakeTimeOffset: -24 * time.Hour,
	})

	retrievedStake, err := suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)

	// unstake the full amount
	_, err = suite.App.IncentivesKeeper.Unstake(suite.Ctx, stake, sdk.Coins{sdk.NewInt64Coin(suite.LPDenom, 9)})
	suite.Require().NoError(err)
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().ElementsMatch(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 9)), balances)
	// should still be accessible
	retrievedStake, err = suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)
	suite.Require().ElementsMatch(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom, 11)), retrievedStake.Coins)
	// fin.
}