## middleware-cli

middleware-cli

### Synopsis

middleware-cli is Cli for management middleware DittoNetwork

### Install 
[You must install golang](https://go.dev/doc/install)
```bash
go build -o middleware-cli

```

### Usage

```bash
middleware-cli [Method] [Options] [General Options]
```

or
```bash
go run main.go [Method] [Options] [General Options]
```

### General options

```
  -h, --help                 help for middleware-cli
      --middleware string    middleware address
      --private-key string   private key to use
      --provider string      provider to use
```

### SEE ALSO

* [middleware-cli addCO](docs/middleware-cli_addCO.md)	 - Add collateral-oracle address
* [middleware-cli adminRole](docs/middleware-cli_adminRole.md)	 - Get default admin role
* [middleware-cli completion](docs/middleware-cli_completion.md)	 - Generate the autocompletion script for the specified shell
* [middleware-cli gendoc](docs/middleware-cli_gendoc.md)	 - Generate documentation
* [middleware-cli grantRole](docs/middleware-cli_grantRole.md)	 - Grant role
* [middleware-cli hasRole](docs/middleware-cli_hasRole.md)	 - Check role for address
* [middleware-cli isOpReg](docs/middleware-cli_isOpReg.md)	 - Check operator registered or not
* [middleware-cli minStake](docs/middleware-cli_minStake.md)	 - Get minimum stake value
* [middleware-cli operatorRole](docs/middleware-cli_operatorRole.md)	 - Get operator role
* [middleware-cli pauseOp](docs/middleware-cli_pauseOp.md)	 - Pause operator
* [middleware-cli pauseOpVault](docs/middleware-cli_pauseOpVault.md)	 - Pause operator vault
* [middleware-cli pauseShVault](docs/middleware-cli_pauseShVault.md)	 - Pause shared vault
* [middleware-cli regOp](docs/middleware-cli_regOp.md)	 - Register operator in middleware
* [middleware-cli regOpVault](docs/middleware-cli_regOpVault.md)	 - Register operator vault
* [middleware-cli regShVault](docs/middleware-cli_regShVault.md)	 - Register shared vault
* [middleware-cli removeCO](docs/middleware-cli_removeCO.md)	 - Remove collateral-oracle address
* [middleware-cli renounceRole](docs/middleware-cli_renounceRole.md)	 - Renounce role
* [middleware-cli revokeRole](docs/middleware-cli_revokeRole.md)	 - Revoke role
* [middleware-cli setMinStake](docs/middleware-cli_setMinStake.md)	 - Set minimum stake
* [middleware-cli unOp](docs/middleware-cli_unOp.md)	 - Unregister operator
* [middleware-cli unPauseOp](docs/middleware-cli_unPauseOp.md)	 - Unpause operator
* [middleware-cli unPauseOpVault](docs/middleware-cli_unPauseOpVault.md)	 - Unpause operator vault
* [middleware-cli unPauseShVault](docs/middleware-cli_unPauseShVault.md)	 - Unpause shared vault
* [middleware-cli unRegOpVault](docs/middleware-cli_unRegOpVault.md)	 - Unregister operator vault
* [middleware-cli unRegShVault](docs/middleware-cli_unRegShVault.md)	 - Unregister shared vault
* [middleware-cli updateKeyOp](docs/middleware-cli_updateKeyOp.md)	 - Update operator
* [middleware-cli valSet](docs/middleware-cli_valSet.md)	 - Get validator set
* [middleware-cli valSetUnbonding](docs/middleware-cli_valSetUnbonding.md)	 - Get validator set with unbonding time
