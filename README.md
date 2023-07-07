# Genesis Setup

## Group Genesis Config

Run with:

```
./create-group-with-policy.sh > group_genesis.json
```

I've tested with various members.json addresses and each time the first group
policy address is
`cosmos1afk9zr2hn2jsac63h4hm60vl9z3e5u69gndzf7c99cqge3vzwjzsfwkgpd`. This means
we can hardcode that address into the dualityd binary before we have our final
group genesis generated.

To update the genesis.json with the final group_genesis.json run the following.

```
group_genesis=$(dasel -f "group_genesis.json" '.')
dasel put object -f "old_genesis.json" -s "$group_genesis" '.app_state.group' > new_genesis.json
```

## Wasm Genesis Config

We're locking things down so that only the admin group defined by the process above has the permission
to upload and instantiate wasm contracts. [Docs here](https://github.com/CosmWasm/wasmd/tree/main#genesis-configuration).

```
wasm_genesis=$(dasel -f "wasm_genesis.json" '.')
dasel put object -f "old_genesis.json" -s "$wasm_genesis" '.app_state.wasm' > new_genesis.json
```