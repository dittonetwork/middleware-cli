package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	middlewareContract "github.com/dittonetwork/middleware-cli/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra/doc"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	client      *ethclient.Client
	middlewareC *middlewareContract.Middleware
)

func main() {
	ctx := context.Background()

	rootCmd := &cobra.Command{
		Use:   "middleware-cli",
		Short: "middleware-cli",
		Long:  "middleware-cli is Cli for management middleware DittoNetwork",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			provider, _ := cmd.Flags().GetString("provider")
			middleware, _ := cmd.Flags().GetString("middleware")
			if provider == "" {
				log.Fatal("provider rpc-url is required")
			} else if middleware == "" {
				log.Fatal("middleware address is required")
			}

			var err error
			client, err = ethclient.DialContext(ctx, provider)
			if err != nil {
				log.Fatal("Connect not created: ", err)
			}
			defer client.Close()

			middlewareC, err = middlewareContract.NewMiddleware(common.HexToAddress(middleware), client)
			if err != nil {
				log.Fatal("New ", err)
			}
		},
	}
	rootCmd.PersistentFlags().String("provider", "", "provider to use")
	rootCmd.PersistentFlags().String("middleware", "", "middleware address")
	rootCmd.PersistentFlags().String("private-key", "", "private key to use")
	_ = rootCmd.MarkFlagRequired("provider")
	_ = rootCmd.MarkFlagRequired("middleware")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "gendoc",
		Short: "Generate documentation",
		Run: func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(rootCmd, "./")
			if err != nil {
				log.Fatalf("Error generating documentation: %v", err)
			}
			log.Println("Documentation generated successfully!")
		},
	})

	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Operator command

	// Register operator
	var regOp = &cobra.Command{
		Use:   "regOp",
		Short: "Register operator in middleware",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			key, _ := cmd.Flags().GetString("key")
			vault, _ := cmd.Flags().GetString("vault")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if key == "" {
				fmt.Println("Please provide key")
			} else {
				vaultAddress := common.Address{}
				if vault != "" {
					vaultAddress = common.HexToAddress(vault)
				}
				tx, err := middlewareC.RegisterOperator(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					[]byte(key),
					vaultAddress,
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	regOp.Flags().String("operator", "", "operator address")
	regOp.Flags().String("key", "", "key")
	regOp.Flags().String("vault", "", "vault address (optional)")
	_ = regOp.MarkFlagRequired("operator")
	_ = regOp.MarkFlagRequired("key")

	rootCmd.AddCommand(regOp)

	// Unregister operator
	var unOp = &cobra.Command{
		Use:   "unOp",
		Short: "Unregister operator",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else {
				tx, err := middlewareC.UnregisterOperator(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unOp.Flags().String("operator", "", "operator address")
	_ = unOp.MarkFlagRequired("operator")

	rootCmd.AddCommand(unOp)

	// PauseOperator
	var pauseOp = &cobra.Command{
		Use:   "pauseOp",
		Short: "Pause operator",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else {
				tx, err := middlewareC.PauseOperator(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	pauseOp.Flags().String("operator", "", "operator address")
	_ = pauseOp.MarkFlagRequired("operator")
	rootCmd.AddCommand(pauseOp)

	// UnpauseOperator
	var unPauseOp = &cobra.Command{
		Use:   "unPauseOp",
		Short: "Unpause operator",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else {
				tx, err := middlewareC.UnpauseOperator(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unPauseOp.Flags().String("operator", "", "operator address")
	_ = unPauseOp.MarkFlagRequired("operator")
	rootCmd.AddCommand(unPauseOp)

	// UpdateKeyOperator
	var updateKeyOp = &cobra.Command{
		Use:   "updateKeyOp",
		Short: "Update operator",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			key, _ := cmd.Flags().GetString("key")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if key == "" {
				fmt.Println("Please provide key")
			} else {
				tx, err := middlewareC.UpdateOperatorKey(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					[]byte(key),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	updateKeyOp.Flags().String("operator", "", "operator address")
	updateKeyOp.Flags().String("key", "", "key")
	_ = updateKeyOp.MarkFlagRequired("operator")
	_ = updateKeyOp.MarkFlagRequired("key")
	rootCmd.AddCommand(updateKeyOp)

	// Register operator vault
	var regOpVault = &cobra.Command{
		Use:   "regOpVault",
		Short: "Register operator vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			vault, _ := cmd.Flags().GetString("vault")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.RegisterOperatorVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	regOpVault.Flags().String("operator", "", "operator address")
	regOpVault.Flags().String("vault", "", "vault address")
	_ = regOpVault.MarkFlagRequired("operator")
	_ = regOpVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(regOpVault)

	// Unregister operator vault
	var unRegOpVault = &cobra.Command{
		Use:   "unRegOpVault",
		Short: "Unregister operator vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			vault, _ := cmd.Flags().GetString("vault")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.UnregisterOperatorVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unRegOpVault.Flags().String("operator", "", "operator address")
	unRegOpVault.Flags().String("vault", "", "vault address")
	_ = unRegOpVault.MarkFlagRequired("operator")
	_ = unRegOpVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(unRegOpVault)

	// Pause operator vault
	var pauseOpVault = &cobra.Command{
		Use:   "pauseOpVault",
		Short: "Pause operator vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			vault, _ := cmd.Flags().GetString("vault")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.PauseOperatorVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	pauseOpVault.Flags().String("operator", "", "operator address")
	pauseOpVault.Flags().String("vault", "", "vault address")
	_ = pauseOpVault.MarkFlagRequired("operator")
	_ = pauseOpVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(pauseOpVault)

	// Unpause operator vault
	var unPauseOpVault = &cobra.Command{
		Use:   "unPauseOpVault",
		Short: "Unpause operator vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			operator, _ := cmd.Flags().GetString("operator")
			vault, _ := cmd.Flags().GetString("vault")
			if operator == "" {
				fmt.Println("Please provide operator address")
			} else if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.UnpauseOperatorVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(operator),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unPauseOpVault.Flags().String("operator", "", "operator address")
	unPauseOpVault.Flags().String("vault", "", "vault address")
	_ = unPauseOpVault.MarkFlagRequired("operator")
	_ = unPauseOpVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(unPauseOpVault)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Admin command

	// Register shared vault
	var regShVault = &cobra.Command{
		Use:   "regShVault",
		Short: "Register shared vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			vault, _ := cmd.Flags().GetString("vault")
			if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.RegisterSharedVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	regShVault.Flags().String("vault", "", "vault address")
	_ = regShVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(regShVault)

	// Unregister shared vault
	var unRegShVault = &cobra.Command{
		Use:   "unRegShVault",
		Short: "Unregister shared vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			vault, _ := cmd.Flags().GetString("vault")
			if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.UnregisterSharedVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unRegShVault.Flags().String("vault", "", "vault address")
	_ = unRegShVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(unRegShVault)

	// Pause shared vault
	var pauseShVault = &cobra.Command{
		Use:   "pauseShVault",
		Short: "Pause shared vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			vault, _ := cmd.Flags().GetString("vault")
			if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.PauseSharedVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	pauseShVault.Flags().String("vault", "", "vault address")
	_ = pauseShVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(pauseShVault)

	// Unpause shared vault
	var unPauseShVault = &cobra.Command{
		Use:   "unPauseShVault",
		Short: "Unpause shared vault",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			vault, _ := cmd.Flags().GetString("vault")
			if vault == "" {
				fmt.Println("Please provide vault address")
			} else {
				tx, err := middlewareC.UnpauseSharedVault(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(vault),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	unPauseShVault.Flags().String("vault", "", "vault address")
	_ = unPauseShVault.MarkFlagRequired("vault")
	rootCmd.AddCommand(unPauseShVault)

	// Grant Role
	var grantRole = &cobra.Command{
		Use:   "grantRole",
		Short: "Grant role",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			role, _ := cmd.Flags().GetString("role")
			grantedAddress, _ := cmd.Flags().GetString("granted-address")
			if role == "" {
				fmt.Println("Please provide role")
			} else if grantedAddress == "" {
				fmt.Println("Please provide granted address")
			} else {
				tx, err := middlewareC.GrantRole(
					getTrasactionOpt(ctx, privateKey),
					stringTo32ArrayBytes(role),
					common.HexToAddress(grantedAddress),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	grantRole.Flags().String("role", "", "role name")
	grantRole.Flags().String("granted-address", "", "granted address")
	_ = grantRole.MarkFlagRequired("role")
	_ = grantRole.MarkFlagRequired("granted-address")
	rootCmd.AddCommand(grantRole)

	// Renounce Role
	var renounceRole = &cobra.Command{
		Use:   "renounceRole",
		Short: "Renounce role",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			role, _ := cmd.Flags().GetString("role")
			renouncedAddress, _ := cmd.Flags().GetString("renounced-address")
			if role == "" {
				fmt.Println("Please provide role")
			} else if renouncedAddress == "" {
				fmt.Println("Please provide renounced address")
			} else {
				tx, err := middlewareC.RenounceRole(
					getTrasactionOpt(ctx, privateKey),
					stringTo32ArrayBytes(role),
					common.HexToAddress(renouncedAddress),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	renounceRole.Flags().String("role", "", "role name")
	renounceRole.Flags().String("renounced-address", "", "renounced address")
	_ = renounceRole.MarkFlagRequired("role")
	_ = renounceRole.MarkFlagRequired("renounced-address")
	rootCmd.AddCommand(renounceRole)

	// Revoke Role
	var revokeRole = &cobra.Command{
		Use:   "revokeRole",
		Short: "Revoke role",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			role, _ := cmd.Flags().GetString("role")
			revokedAddress, _ := cmd.Flags().GetString("revoked-address")
			if role == "" {
				fmt.Println("Please provide role")
			} else if revokedAddress == "" {
				fmt.Println("Please provide revoked address")
			} else {
				tx, errIn := middlewareC.RevokeRole(
					getTrasactionOpt(ctx, privateKey),
					stringTo32ArrayBytes(role),
					common.HexToAddress(revokedAddress),
				)
				if errIn != nil {
					log.Fatal(errIn)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	revokeRole.Flags().String("role", "", "role name")
	revokeRole.Flags().String("revoked-address", "", "revoked address")
	_ = revokeRole.MarkFlagRequired("role")
	_ = revokeRole.MarkFlagRequired("revoked-address")
	rootCmd.AddCommand(revokeRole)

	//addCollateralOracle - add a collateral corresponding to oracle
	var addCollateralOracle = &cobra.Command{
		Use:   "addCO",
		Short: "Add collateral-oracle address",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			collateral, _ := cmd.Flags().GetString("collateral")
			oracle, _ := cmd.Flags().GetString("oracle")
			if collateral == "" {
				fmt.Println("Please provide collateral")
			} else if oracle == "" {
				fmt.Println("Please provide oracle")
			} else {
				tx, err := middlewareC.AddCollateralOracle(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(collateral),
					common.HexToAddress(oracle),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	addCollateralOracle.Flags().String("collateral", "", "collateral address")
	addCollateralOracle.Flags().String("oracle", "", "oracle address")
	_ = addCollateralOracle.MarkFlagRequired("collateral")
	_ = addCollateralOracle.MarkFlagRequired("oracle")
	rootCmd.AddCommand(addCollateralOracle)

	//removeCollateralOracle - revoke oracle's collateral
	var removeCollateralOracle = &cobra.Command{
		Use:   "removeCO",
		Short: "Remove collateral-oracle address",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			collateral, _ := cmd.Flags().GetString("collateral")
			if collateral == "" {
				fmt.Println("Please provide collateral")
			} else {
				tx, err := middlewareC.RemoveCollateralOracle(
					getTrasactionOpt(ctx, privateKey),
					common.HexToAddress(collateral),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	removeCollateralOracle.Flags().String("collateral", "", "collateral address")
	_ = removeCollateralOracle.MarkFlagRequired("collateral")
	rootCmd.AddCommand(removeCollateralOracle)

	// setMinStake - set minStake for operator logging into the Ditto Network(middleware)
	var setMinStake = &cobra.Command{
		Use:   "setMinStake",
		Short: "Set minimum stake",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := getAndCheckPK(cmd)
			minstake, _ := cmd.Flags().GetString("min-stake")
			if minstake == "" {
				fmt.Println("Please provide minimum stake")
			} else {
				bgMS := new(big.Int)
				bgMS, _ = bgMS.SetString(minstake, 10)

				tx, err := middlewareC.SetMinstake(
					getTrasactionOpt(ctx, privateKey),
					bgMS,
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Tx sent: ", tx.Hash().Hex())
			}
		},
	}
	setMinStake.Flags().String("min-stake", "", "minimum stake")
	_ = setMinStake.MarkFlagRequired("min-stake")
	rootCmd.AddCommand(setMinStake)
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Common command

	// Get DEFAULT_ADMIN_ROLE
	var defaultAdminRole = &cobra.Command{
		Use:   "adminRole",
		Short: "Get default admin role",
		Run: func(cmd *cobra.Command, args []string) {
			adminRole, errIn := middlewareC.DEFAULTADMINROLE(&bind.CallOpts{})
			if errIn != nil {
				log.Fatal(errIn)
			}

			fmt.Println("Operator role: ", common.BytesToHash(adminRole[:]).Hex())
		},
	}
	rootCmd.AddCommand(defaultAdminRole)

	// Get OPERATOR_ROLE
	var operatorRole = &cobra.Command{
		Use:   "operatorRole",
		Short: "Get operator role",
		Run: func(cmd *cobra.Command, args []string) {
			operatorRole, errIn := middlewareC.OPERATORROLE(&bind.CallOpts{})
			if errIn != nil {
				log.Fatal(errIn)
			}

			fmt.Println("Operator role: ", common.BytesToHash(operatorRole[:]).Hex())
		},
	}
	rootCmd.AddCommand(operatorRole)

	// has role
	var hasRole = &cobra.Command{
		Use:   "hasRole",
		Short: "Check role for address",
		Run: func(cmd *cobra.Command, args []string) {
			role, _ := cmd.Flags().GetString("role")
			checkAddress, _ := cmd.Flags().GetString("check-address")
			if role == "" {
				fmt.Println("Please provide role")
			} else if checkAddress == "" {
				fmt.Println("Please provide checked address")
			} else {

				hasRoleResult, err := middlewareC.HasRole(
					&bind.CallOpts{},
					stringTo32ArrayBytes(role),
					common.HexToAddress(checkAddress),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Address: ", checkAddress, " has role: ", role, " is: ", hasRoleResult)
			}
		},
	}
	hasRole.Flags().String("role", "", "role")
	hasRole.Flags().String("check-address", "", "checked address")
	_ = hasRole.MarkFlagRequired("role")
	_ = hasRole.MarkFlagRequired("check-address")
	rootCmd.AddCommand(hasRole)

	// isOperatorRegistered
	var isOperatorRegistered = &cobra.Command{
		Use:   "isOpReg",
		Short: "Check operator registered or not",
		Run: func(cmd *cobra.Command, args []string) {
			operator, _ := cmd.Flags().GetString("operator")
			if operator == "" {
				fmt.Println("Please provide operator")
			} else {
				fmt.Println("Operator: ", operator)
				isOperator, err := middlewareC.IsOperatorRegistered(
					&bind.CallOpts{},
					common.HexToAddress(operator),
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Operator: ", operator, " isOperator: ", isOperator)
			}
		},
	}
	isOperatorRegistered.Flags().String("operator", "", "operator address")
	_ = isOperatorRegistered.MarkFlagRequired("operator")
	rootCmd.AddCommand(isOperatorRegistered)

	// Get MinStake
	var minStake = &cobra.Command{
		Use:   "minStake",
		Short: "Get minimum stake value",
		Run: func(cmd *cobra.Command, args []string) {
			ms, err := middlewareC.MinStake(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Middleware minStake value: ", ms.String())
		},
	}
	rootCmd.AddCommand(minStake)

	// Get Validator Set
	var validatorSet = &cobra.Command{
		Use:   "valSet",
		Short: "Get validator set",
		Run: func(cmd *cobra.Command, args []string) {
			validatorData, errIn := middlewareC.GetValidatorSet(&bind.CallOpts{})
			if errIn != nil {
				log.Fatal(errIn)
			}

			for _, validator := range validatorData {
				fmt.Printf(
					"Operator address: %s\nOperator stake: %s\nOperator key: %s\n\n",
					validator.Operator.String(), validator.Stake.String(), validator.Key,
				)
			}
		},
	}
	rootCmd.AddCommand(validatorSet)

	//Get Validator Set with unbonding time
	var validatorSetUnbonding = &cobra.Command{
		Use:   "valSetUnbonding",
		Short: "Get validator set with unbonding time",
		Run: func(cmd *cobra.Command, args []string) {
			validatorData, errIn := middlewareC.GetValidatorSetWithUnbonding(&bind.CallOpts{})
			if errIn != nil {
				log.Fatal(errIn)
			}

			for _, validator := range validatorData {
				fmt.Printf(
					"Operator address: %s\nOperator key: %s\n",
					validator.Operator.String(), validator.Key,
				)
				for _, dataValidator := range validator.VaultData {
					fmt.Printf(
						"Vault address: %s\nStake: %s\nPowerExpiresAt: %s\nUnbondingAmount: %s\nUnbondingTime: %s\n\n",
						dataValidator.Vault.String(), dataValidator.Stake.String(), dataValidator.PowerExpiresAt.String(), dataValidator.UnbondingAmount.String(),
						dataValidator.UnbondingTime.String(),
					)
				}

				fmt.Println()
			}
		},
	}

	rootCmd.AddCommand(validatorSetUnbonding)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getAndCheckPK(cmd *cobra.Command) string {
	privateKey, _ := cmd.Flags().GetString("private-key")
	if privateKey == "" {
		log.Fatal("private-key is required")
	}

	return privateKey
}

func getTrasactionOpt(ctx context.Context, privateKey string) *bind.TransactOpts {
	privateK, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateK.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum network: %v", err)
	}

	tO, err := bind.NewKeyedTransactorWithChainID(privateK, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tO.Nonce = big.NewInt(int64(nonce))
	tO.Value = big.NewInt(0)
	tO.GasPrice = gasPrice

	return tO
}

func stringTo32ArrayBytes(role string) [32]byte {
	var roleRes [32]byte
	copy(roleRes[:], common.HexToHash(role).Bytes())

	return roleRes
}
