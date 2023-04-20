[View code on GitHub](https://github.com/duality-labs/duality/keeper/core.go)

This code is part of the `keeper` package and provides core functionalities for handling various operations in a decentralized exchange (DEX) module, such as depositing, withdrawing, swapping, placing limit orders, canceling limit orders, and withdrawing filled limit orders.

The `DepositCore` function handles the core logic for depositing tokens into the DEX. It checks and initializes data structures (tick, pair), calculates shares based on the amount deposited, and sends funds to the module address.

The `WithdrawCore` function handles the core logic for withdrawing tokens from the DEX. It calculates the amount of reserve tokens to withdraw based on the specified number of shares to remove and sends the withdrawn tokens to the user's account.

The `SwapCore` function facilitates swapping tokens in the DEX. It takes input tokens and calculates the output tokens based on the specified pair (token0, token1) and sends the output tokens to the user's account.

The `MultiHopSwapCore` function allows users to perform multi-hop swaps, which involve swapping tokens through multiple routes to achieve the best price.

The `PlaceLimitOrderCore` function handles placing limit orders in the DEX. It initializes data structures if needed, calculates and stores information for a new limit order at a specific tick.

The `CancelLimitOrderCore` function handles canceling limit orders in the DEX. It removes the specified number of shares from a limit order and returns the respective amount in terms of the reserve to the user.

The `WithdrawFilledLimitOrderCore` function handles withdrawing filled liquidity from a limit order based on the amount wished to receive. It calculates and sends filled liquidity from the module to the user's account.

These functions are essential for the operation of a decentralized exchange and can be used in various scenarios within the larger project, such as trading, providing liquidity, and managing orders.
## Questions: 
 1. **Question**: What is the purpose of the `TruncateInt` function mentioned in the note at the beginning of the code, and why is it used in multiple places?
   
   **Answer**: The `TruncateInt` function is used to convert Decimals back into sdk.Ints. It is used in multiple places because it may create some accounting anomalies, but it is considered preferable to other alternatives. The full ADR can be found at the provided link: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

2. **Question**: What is the purpose of the `IsBehindEnemyLines` function and why are there TODO comments related to it?

   **Answer**: The `IsBehindEnemyLines` function checks if a deposit is being made "behind enemy lines", meaning it is being deposited in a tick that is not favorable for the depositor. The TODO comments indicate that there are plans to allow users to deposit "behind enemy lines" in the future, but it is currently not allowed.

3. **Question**: What is the purpose of the `MultiHopSwapCore` function and how does it differ from the `SwapCore` function?

   **Answer**: The `MultiHopSwapCore` function facilitates a multi-hop swap, which means it allows swapping between multiple pairs of tokens in a single transaction. This is different from the `SwapCore` function, which only facilitates a swap between a single pair of tokens. The `MultiHopSwapCore` function can be useful when there is no direct liquidity between two tokens, and a multi-hop swap can provide better rates.