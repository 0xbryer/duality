[View code on GitHub](https://github.com/duality-labs/duality/epochs/client/cli/query.go)

The code in this file is part of the duality project and provides a set of CLI query commands for the epochs module. The `GetQueryCmd` function returns a `cobra.Command` object that can be used to execute the CLI commands. The `GetCmdEpochInfos` and `GetCmdCurrentEpoch` functions return `osmocli.QueryDescriptor` objects that describe the CLI commands for querying epoch information.

The `GetQueryCmd` function first calls `osmocli.QueryIndexCmd` to create a new `cobra.Command` object with the name of the epochs module. It then adds two query commands to the command object using `osmocli.AddQueryCmd`. The first command queries running epoch information using the `EpochInfos` function from the `types` package. The second command queries the current epoch by specified identifier using the `QueryCurrentEpochRequest` function from the `types` package.

The `GetCmdEpochInfos` and `GetCmdCurrentEpoch` functions both return `osmocli.QueryDescriptor` objects that describe the CLI commands for querying epoch information. The `GetCmdEpochInfos` function returns a descriptor for querying running epoch information, while the `GetCmdCurrentEpoch` function returns a descriptor for querying the current epoch by specified identifier.

Overall, this code provides a set of CLI commands for querying epoch information in the duality project. These commands can be used by developers and users to retrieve information about the current epoch and running epoch information. For example, a developer may use these commands to debug issues related to epoch transitions or to monitor the progress of the current epoch.
## Questions: 
 1. What is the purpose of the `duality-labs/duality/osmoutils/osmocli` package?
- The `duality-labs/duality/osmoutils/osmocli` package is used to create query commands for the `duality` module's CLI.

2. What is the difference between `GetCmdEpochInfos` and `GetCmdCurrentEpoch`?
- `GetCmdEpochInfos` is used to query running epoch information, while `GetCmdCurrentEpoch` is used to query the current epoch by a specified identifier.

3. What is the role of `osmocli.AddQueryCmd` in `GetQueryCmd`?
- `osmocli.AddQueryCmd` is used to add query commands to the `cmd` object returned by `GetQueryCmd`, using the `types.NewQueryClient` function and either `GetCmdEpochInfos` or `GetCmdCurrentEpoch` as arguments.