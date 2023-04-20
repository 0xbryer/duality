[View code on GitHub](https://github.com/duality-labs/duality/incentives/keeper/stake.go)

The `keeper` package contains the implementation of the `Keeper` struct, which is responsible for managing the state of the `incentives` module in the `duality` project. The `Keeper` struct provides several methods for interacting with the state of the module, including `GetLastStakeID`, `SetLastStakeID`, `Stake`, `Unstake`, `setStake`, `deleteStake`, `GetStakeByID`, `GetStakesByQueryCondition`, `GetStakes`, `getStakesByAccount`, and `CreateStake`.

The `GetLastStakeID` method retrieves the ID used last time from the key-value store associated with the `incentives` module. The `SetLastStakeID` method saves the ID used by the last stake to the key-value store. The `Stake` method is an internal utility that stakes coins and sets corresponding states. This method is only called by either of the two possible entry points to stake tokens: `CreateStake` and `AddTokensToStakeByID`. The `Unstake` method is used to unstake tokens and remove the corresponding stake object from the state. The `setStake` method is a utility to store a stake object into the store. The `deleteStake` method removes the stake object from the state. The `GetStakeByID` method returns a stake from a stake ID. The `GetStakesByQueryCondition` method returns the period locks associated with an account based on a query condition. The `GetStakes` method returns all the period locks in the state. The `getStakesByAccount` method returns the period locks associated with an account. The `CreateStake` method creates a new stake object and stores it in the state.

Overall, the `keeper` package provides a set of methods for managing the state of the `incentives` module in the `duality` project. These methods can be used to stake and unstake tokens, retrieve stake objects from the state, and manage the period locks associated with an account.
## Questions: 
 1. What is the purpose of the `Stake` function and when is it called?
- The `Stake` function is an internal utility to stake coins and set corresponding states. It is called by either of the two possible entry points to stake tokens: `CreateStake` and `AddTokensToStakeByID`.

2. What is the difference between `GetStakes` and `GetStakesByAccount` functions?
- The `GetStakes` function returns all period locks, while the `GetStakesByAccount` function returns the period locks associated with a specific account.

3. What is the purpose of the `setStake` function?
- The `setStake` function is a utility to store a stake object into the store. It takes a `Keeper` and a `Stake` object as input, and stores the object in the key-value store associated with the `Keeper`.