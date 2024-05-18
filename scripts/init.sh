#!/bin/bash

IF_INIT=1
IF_ADD_USERS=1
IF_COLLECT_GENTXS=1

CMD_HOME=./build/blogd
DENOM=alice
CONFIG_HOME=./private/.blogapp
CHAIN_ID=blog-1
MONIKER=blog
MIN_GAS_PRICE=0

declare -a USERS=("james" "alice")
declare -a USERS_ACCOUNT_TOKENS=("james" 100000000)
declare -a STAKING_USERS_ACCOUNT_TOKENS=("james" 70000000)
KEYRING=test

if [[ $IF_INIT -eq 1 ]]; then
  $CMD_HOME init "$MONIKER" \
    --chain-id "$CHAIN_ID" \
    --default-denom "$DENOM" \
    --home "$CONFIG_HOME"
fi

if [[ $IF_ADD_USERS -eq 1 ]]; then
  for ((i=0; i<${#USERS[@]}; i+=1)); do
    name=${USERS[i]}
    $CMD_HOME keys add $name \
        --home "$CONFIG_HOME" \
        --keyring-backend "$KEYRING"
  done
fi

if [[ $IF_COLLECT_GENTXS -eq 1 ]]; then
    for ((i=0; i<${#USERS_ACCOUNT_TOKENS[@]}; i+=2)); do
        name=${USERS_ACCOUNT_TOKENS[i]}
        token=${USERS_ACCOUNT_TOKENS[i+1]}
        $CMD_HOME genesis add-genesis-account $name "$token""$DENOM" \
            --home "$CONFIG_HOME" \
            --keyring-backend "$KEYRING"
    done
    for ((i=0; i<${#STAKING_USERS_ACCOUNT_TOKENS[@]}; i+=2)); do
        name=${STAKING_USERS_ACCOUNT_TOKENS[i]}
        token=${STAKING_USERS_ACCOUNT_TOKENS[i+1]}
        $CMD_HOME genesis gentx $name "$token""$DENOM" \
            --home "$CONFIG_HOME" \
            --keyring-backend "$KEYRING" \
            --chain-id "$CHAIN_ID"
    done

    $CMD_HOME genesis collect-gentxs \
    --home "$CONFIG_HOME"
fi

app_config_path="$CONFIG_HOME/config/app.toml"
if [ -f "$app_config_path" ]; then
  sed -i '' "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"$MIN_GAS_PRICE$DENOM\"/g" "$app_config_path"
fi