## 1. Initial nodes :

```bash
kava init node1 --chain-id=aizel_2024-12 --home=$KAVAHOME/node1

kava init node2 --chain-id=aizel_2024-12 --home=$KAVAHOME/node2

kava init node3 --chain-id=aizel_2024-12 --home=$KAVAHOME/node3
```
## 2.0 add keys :

### 2.0.1 For keys alice, only put in node1:
```bash 
# Replace <your-key-name> with a name for your key that you will remember, such as alice, bob, james
# kava keys add alice --home=$KAVAHOME/node1 --keyring-backend=file
# or import it from the exported private key 
kava keys import alice $KAVAHOME/alice.pem --home=$KAVAHOME/node1 --keyring-backend=file
# To see a list of wallets on your node
kava keys list --home=$KAVAHOME/node1 --keyring-backend=file
```

### 2.0.2 For keys bob, put in both node1 and node2:
```bash 
# Replace <your-key-name> with a name for your key that you will remember, such as alice, bob, james
# kava keys add bob --home=$KAVAHOME/node1 --keyring-backend=file
# or import it from the exported private key to node1
kava keys import bob $KAVAHOME/bob.pem --home=$KAVAHOME/node1 --keyring-backend=file
# import to node2
kava keys import bob $KAVAHOME/bob.pem --home=$KAVAHOME/node2 --keyring-backend=file
```

### 2.0.3 For keys james, put in both node1 and node3:
```bash 
# Replace <your-key-name> with a name for your key that you will remember, such as alice, bob, james
# kava keys add james --home=$KAVAHOME/node1 --keyring-backend=file
# or import it from the exported private key to node1
kava keys import james $KAVAHOME/james.pem --home=$KAVAHOME/node1 --keyring-backend=file
# import to node3
kava keys import james $KAVAHOME/james.pem --home=$KAVAHOME/node3 --keyring-backend=file
```

## 2.1 Add Genesis Accounts : Add initial funds to the accounts in the genesis file:

```bash
# pre fund ukava
kava add-genesis-account alice  1000000000ukava  --home=$KAVAHOME/node1 --keyring-backend=file

kava add-genesis-account bob  1000000000ukava  --home=$KAVAHOME/node1  --keyring-backend=file

kava add-genesis-account james  1000000000ukava  --home=$KAVAHOME/node1  --keyring-backend=file

```

## 2.2 stake and set validators :
### 2.2.1 Manually add balance and total_supply for stake , each account one by one
```bash
      "balances": [
        {
          "address": "kava1qqkdj2vvzdgq7e8yzmncjq5nfl0lhpy4rt6e2x",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        },
        {
          "address": "kava18p4et3um5aez9cnmxwld7yft6xu2xtvnv5jwqm",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        },
        {
          "address": "kava1m57dmeyqqa3ujv3zpy353em8fdram97e49789t",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "stake",
          "amount": "4123191823776"
        },
        {
          "denom": "ukava",
          "amount": "3000000000"
        }
      ]
```
### 2.2.2 Gentx for alice
```bash
# for alice
kava gentx alice 1374397274592stake --chain-id aizel_2024-12 --commission-rate 0.10 --commission-max-rate 0.20 --commission-max-change-rate 0.01 --min-self-delegation 1 --home=$KAVAHOME/node1   --keyring-backend=file
```

### 2.2.3 collect gents
```bash
kava collect-gentxs --home=$KAVAHOME/node1 --keyring-backend=file
```

## 2.3 pre-fund evm addresses :

### 2.3.1 using the following tool to convert evm address to bech32 format address :

```golang
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func evmToBech32(evmAddr, prefix string) (string, error) {
	// Remove the "0x" prefix and decode the hex address
	addressBytes, err := hex.DecodeString(evmAddr[2:])
	if err != nil {
		return "", err
	}

	// Convert to Bech32 format
	bech32Addr, err := bech32.ConvertAndEncode(prefix, addressBytes)
	if err != nil {
		return "", err
	}

	return bech32Addr, nil
}

func main() {
	evmAddr := "0xaaafB3972B05630fCceE866eC69CdADd9baC2771"
	prefix := "kava"

	bech32Addr, err := evmToBech32(evmAddr, prefix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bech32 Address:", bech32Addr)
}
```

### 2.3.2 add the converted evm address to genesis account and pre-fund with aphoton

```json
## add new account to accounts
{
          "@type": "/ethermint.types.v1.EthAccount",
          "base_account": {
            "address": "kava142hm89etq43sln8wsehvd8x6mkd6cfm3yugv6c",
            "pub_key": null,
            "account_number": "3",
            "sequence": "1"
          },
          "code_hash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
}
## add prefund to balances
{
          "address": "kava142hm89etq43sln8wsehvd8x6mkd6cfm3yugv6c",
          "coins": [
            {
              "denom": "aphoton",
              "amount": "100000000000000000000000000"
            }
          ]
}
## add new coins to total_supply
{
          "denom": "aphoton",
          "amount": "100000000000000000000000000"
}
```

### 2.3.3 update genesis eve account :

```bash
kava debug addr AAAFB3972B05630FCCEE866EC69CDADD9BAC2771  --home=$KAVAHOME/node1
```

### 2.3.4 update genesis file's consensus_params to follow :

```json
 "block": {
      "max_bytes": "142020096",
      "max_gas": "3000000000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1448576"
    },
```

### 2.3.5 Validate genesis file :

```bash
kava validate-genesis --home=$KAVAHOME/node1
```
### 2.3.6 Copy the final validated genesis file to other nodes's config directory
```bash
cp $KAVAHOME/node1/config/genesis.json $KAVAHOME/node2/config/
cp $KAVAHOME/node1/config/genesis.json $KAVAHOME/node3/config/
```

# 5. Start nodes :
```bash
nohup kava start --home=$KAVAHOME/node1 > $KAVAHOME/node1/node1.log 2>&1 &
nohup kava start --home=$KAVAHOME/node2 > $KAVAHOME/node2/node2.log 2>&1 &
nohup kava start --home=$KAVAHOME/node3 > $KAVAHOME/node3/node3.log 2>&1 &
```


