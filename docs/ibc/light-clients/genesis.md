<!--
order: 8
-->

# Genesis metadata

Learn how to implement the `ExportMetadata` interface {synopsis}

## Pre-requisite readings

- [Cosmos SDK module genesis](https://docs.cosmos.network/v0.47/building-modules/genesis) {prereq}

`ClientState` instances are provided their own isolated and namespaced client store upon initialisation. `ClientState` implementations may choose to store any amount of arbitrary metadata in order to verify counterparty consensus state and perform light client updates correctly.

The `ExportMetadata` method of the [`ClientState` interface](https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/core/exported/client.go#L47) provides light client modules with the ability to persist metadata in genesis exports.

```go
ExportMetadata(clientStore sdk.KVStore) []GenesisMetadata
```

`ExportMetadata` is provided the client store and returns an array of `GenesisMetadata`. For maximum flexibility, `GenesisMetadata` is defined as a simple interface containing two distinct `Key` and `Value` accessor methods.

```go
type GenesisMetadata interface {
  // return store key that contains metadata without clientID-prefix
  GetKey() []byte
  // returns metadata value
  GetValue() []byte
}
```

This allows `ClientState` instances to retrieve and export any number of key-value pairs which are maintained within the store in their raw `[]byte` form.

When a chain is started with a `genesis.json` file which contains `ClientState` metadata (for example, when performing manual upgrades using an exported `genesis.json`) the `02-client` submodule of core IBC will handle setting the key-value pairs within their respective client stores. [See `02-client` `InitGenesis`](https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/core/02-client/genesis.go#L18-L22).

Please refer to the [Tendermint light client implementation](https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/light-clients/07-tendermint/genesis.go#L12) for an example.
