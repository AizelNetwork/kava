# 1. Initial nodes :

```bash
kava init node1 --chain-id=aizel_2024-12 --home=$KAVAHOME/node1

kava init node2 --chain-id=aizel_2024-12 --home=$KAVAHOME/node2

kava init node3 --chain-id=aizel_2024-12 --home=$KAVAHOME/node3
```

# 2.0 Add Genesis Accounts : Add initial funds to the accounts in the genesis file:

```bash
kava add-genesis-account alice  1000000000ukava  --home=$KAVAHOME/node1

kava add-genesis-account bob  1000000000ukava  --home=$KAVAHOME/node1

kava add-genesis-account james  1000000000ukava  --home=$KAVAHOME/node1

kava add-genesis-account $(kava keys show node1 -a --keyring-backend test --home=$KAVAHOME/node1) 1000000000ukava --home=$KAVAHOME/node1
kava add-genesis-account $(kava keys show node2 -a --keyring-backend test --home=$KAVAHOME/node2) 1000000000ukava --home=$KAVAHOME/node2
kava add-genesis-account $(kava keys show node3 -a --keyring-backend test --home=$KAVAHOME/node3) 1000000000ukava --home=$KAVAHOME/node3
```

# 2.1 create gentx :
```bash
kava add-genesis-account alice  1374397274592stake  --home=$KAVAHOME/node1

kava gentx alice 1374397274592stake --chain-id aizel_2024-12 --commission-rate 0.10 --commission-max-rate 0.20 --commission-max-change-rate 0.01 --min-self-delegation 1 --home=$KAVAHOME/node1
```

# 2.2 collect gents
```bash
kava collect-gentxs --home=$KAVAHOME/node1
```

# 3 pre-fund evm addresses :

# 3.1 import evm address :
```bash
kava keys unsafe-import-eth-key emv1 [pk]
```

# 3.2 add genesis account

```bash
kava add-genesis-account evm1 100000000000000000000aphoton --home=$KAVAHOME/node1
```

# 3.3 update genesis eve account :

```bash
kava debug addr AAAFB3972B05630FCCEE866EC69CDADD9BAC2771
```

# 3.3.1 replace the evm account to follow :

```json
{
  "@type": "/ethermint.types.v1.EthAccount",
  "base_account": {
    "address": "kava1qqxdxk59gwh05shfn0fhyzccxk6t5xxfxscw4c",
    "pub_key": null,
    "account_number": "0",
    "sequence": "1"
  },
  "code_hash": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```


#  3.3.2 update the balances' account




# 4. Validate genesis file :

```bash
kava validate-genesis --home=$KAVAHOME/node1
```


# 5. Start nodes :
```bash
nohup kava start --home=$KAVAHOME/node1 > $KAVAHOME/node1/node1.log 2>&1 &
nohup kava start --home=$KAVAHOME/node2 > $KAVAHOME/node2/node2.log 2>&1 &
nohup kava start --home=$KAVAHOME/node3 > $KAVAHOME/node3/node3.log 2>&1 &
```




# 6. Keys :

```bash
# Replace <your-key-name> with a name for your key that you will remember
kava keys add <your-key-name>
# To see a list of wallets on your node
kava keys list
```

