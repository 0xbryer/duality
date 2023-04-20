[View code on GitHub](https://github.com/duality-labs/duality/incentives/keeper/gauge.go)

The `keeper` package contains the implementation of the `Keeper` struct, which is responsible for managing the state of the `duality` project. This file contains functions that allow for the creation, retrieval, and modification of gauges, which are used to distribute rewards to users based on certain conditions.

The `GetLastGaugeID` function retrieves the ID of the last used gauge from the key-value store. The `SetLastGaugeID` function sets the ID of the last used gauge to the provided ID.

The `getGaugesFromIterator` function iterates over everything in a gauge's iterator until it reaches the end and returns all gauges iterated over. The `setGaugeRefs` function sets the gauge references based on the gauge's status (upcoming, active, or finished).

The `setGauge` function sets the gauge inside the store. The `CreateGauge` function creates a gauge and sends coins to the gauge. The `AddToGaugeRewards` function adds coins to a gauge. The `GetGaugeByID` function retrieves a gauge from the gauge ID. The `GetGauges` function returns upcoming, active, and finished gauges. The `GetNotFinishedGauges` function returns both upcoming and active gauges. The `GetEpochInfo` function returns the EpochInfo struct given the context.

The `moveUpcomingGaugeToActiveGauge` function moves a gauge that has reached its start time from an upcoming to an active status. The `moveActiveGaugeToFinishedGauge` function moves a gauge that has completed its distribution from an active to a finished status. The `GetActiveGauges` function returns active gauges. The `GetUpcomingGauges` function returns upcoming gauges. The `GetFinishedGauges` function returns finished gauges. The `GetGaugesByPair` function returns gauges by pair.

Overall, this file provides the functionality to create, modify, and retrieve gauges, which are used to distribute rewards to users based on certain conditions. These functions are used in the larger `duality` project to manage the distribution of rewards to users.
## Questions: 
 1. What is the purpose of the `duality` project and how does this code file fit into the project?
- This code file is a part of the `keeper` package in the `duality` project. It contains functions related to managing gauges and rewards in the project.

2. What is the role of the `GetLastGaugeID` and `SetLastGaugeID` functions?
- `GetLastGaugeID` returns the last used gauge ID from the key-value store in the context. `SetLastGaugeID` sets the last used gauge ID to the provided ID in the key-value store in the context.

3. What is the purpose of the `moveUpcomingGaugeToActiveGauge` and `moveActiveGaugeToFinishedGauge` functions?
- `moveUpcomingGaugeToActiveGauge` moves a gauge that has reached its start time from an upcoming to an active status. `moveActiveGaugeToFinishedGauge` moves a gauge that has completed its distribution from an active to a finished status.