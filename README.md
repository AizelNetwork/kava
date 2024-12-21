<p align="center">
  <img src="./kava-logo.svg" width="300">
</p>

<div align="center">

[![version](https://img.shields.io/github/tag/kava-labs/kava.svg)](https://github.com/kava-labs/kava/releases/latest)
[![CircleCI](https://circleci.com/gh/Kava-Labs/kava/tree/master.svg?style=shield)](https://circleci.com/gh/Kava-Labs/kava/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kava-labs/kava)](https://goreportcard.com/report/github.com/kava-labs/kava)
[![API Reference](https://godoc.org/github.com/Kava-Labs/kava?status.svg)](https://godoc.org/github.com/Kava-Labs/kava)
[![GitHub](https://img.shields.io/github/license/kava-labs/kava.svg)](https://github.com/Kava-Labs/kava/blob/master/LICENSE.md)
[![Twitter Follow](https://img.shields.io/twitter/follow/KAVA_CHAIN.svg?label=Follow&style=social)](https://twitter.com/KAVA_CHAIN)
[![Discord Chat](https://img.shields.io/discord/704389840614981673.svg)](https://discord.com/invite/kQzh3Uv)

</div>

<div align="center">

### [Telegram](https://t.me/kavalabs) | [Medium](https://medium.com/@kava_labs) | [Discord](https://discord.gg/JJYnuCx)

</div>

Reference implementation of Kava, a blockchain for cross-chain DeFi. Built using the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk).

## Mainnet

The current recommended version of the software for mainnet is [v0.26.2](https://github.com/Kava-Labs/kava/releases/tag/v0.26.2) The `master` branch of this repository often contains considerable development work since the last mainnet release and is __not__ runnable on mainnet.

### Installation and Setup
For detailed instructions see [the Kava docs](https://docs.kava.io/docs/nodes-and-validators/validator-node).

```bash
git checkout v0.26.2
make install
```

End-to-end tests of Kava use a tool for generating networks with different configurations: [kvtool](https://github.com/Kava-Labs/kvtool).
This is included as a git submodule at [`tests/e2e/kvtool`](tests/e2e/kvtool/).
When first cloning the repository, if you intend to run the e2e integration tests, you must also
clone the submodules:
```bash
git clone --recurse-submodules https://github.com/Kava-Labs/kava.git
```

Or, if you have already cloned the repo: `git submodule update --init`

## Testnet

For further information on joining the testnet, head over to the [testnet repo](https://github.com/Kava-Labs/kava-testnets).

## Docs

Kava protocol and client documentation can be found in the [Kava docs](https://docs.kava.io).

If you have technical questions or concerns, ask a developer or community member in the [Kava discord](https://discord.com/invite/kQzh3Uv).

## Aizel Chain

Once you have the binary installed from your customized Kava-based codebase, the next steps to run your own chain are similar to setting up a standard Cosmos-SDK-based network, with adjustments made to reflect your custom modifications. Here’s a general outline:

1. **Initialize Your Chain**:  
   Use the CLI binary you built (e.g. `kava`, or whatever custom name you gave it) to initialize your chain. This step creates the initial configuration files and a genesis file tailored to your custom chain’s parameters. For example:
   ```bash
   kava init aizel-mainnet --chain-id aizel_2222-10
   ```
   - **Moniker**: A human-readable identifier for your node.
   - **Chain-ID**: A unique identifier for your custom chain. Make sure this is unique and not one used by existing networks.

2. **Customize the Genesis File**:  
   After initialization, a `genesis.json` file is placed in your `.kava/config` directory (or the corresponding `.appname/config` if you renamed it). You can:
   - Pre-fund accounts by adding balances directly into the `genesis.json`.
   - Adjust staking parameters, governance rules, EVM configurations, and block time parameters.
   - If you’ve replaced the native token with a custom token, ensure the denomination and any special parameters are updated here.

3. **Adjust Configuration Files**:  
   - **`config.toml`**: Configure networking (persistent peers, seeds), consensus, and logging.
   - **`app.toml`**: Adjust application-level parameters such as:
     - Minimum gas prices.
     - EVM RPC addresses and ports.
     - Database backend (goleveldb or rocksdb).
   - **`client.toml`**: Set up client parameters if necessary.

   Since you’re building a custom chain, these values can be tailored to your requirements. You might remove or simplify certain parameters that were relevant to mainnet Kava but not for your project.

4. **Add Keys and Accounts**:  
   Create keys for accounts that will hold your pre-funded tokens:
   ```bash
   ./kava keys add <key-name>
   ```
   After creating keys, you can edit the genesis file to allocate tokens to these accounts. If you want certain accounts funded at genesis, add them to the `accounts` and `balances` sections in `genesis.json`.

5. **Validate Your Genesis**:  
   Run:
   ```bash
   ./kava validate-genesis
   ```
   This ensures your `genesis.json` is properly structured and all parameters are valid.

6. **Start Your Node**:  
   Once configuration and genesis customizations are complete, start the node:
   ```bash
   ./kava start
   ```
   The node should begin producing blocks if it’s the only validator, or start waiting for other validators if you plan a multi-validator test network.

7. **Add More Validators (Optional)**:  
   If you want a multi-node test network, repeat similar steps on other machines (or multiple terminals), copy the same `genesis.json` to each node’s config directory, and update the `persistent_peers` in `config.toml` so they can connect. Each validator should have its own keys and be added to the genesis or submit a create-validator transaction after the network starts running (if you allow it).

8. **Testing and Iteration**:  
   Make transactions, test governance proposals, and interact with the EVM endpoints if you enabled them. If something needs to change, you can reset and re-initialize the chain with adjusted parameters until you’re satisfied.


## Security

If you find a security issue, please report it to security [at] kavalabs.io. Depending on the verification and severity, a bug bounty may be available.

## License

Copyright © Kava Labs, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
