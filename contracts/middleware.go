// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package middleware

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IValidatorDataValidatorData is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataValidatorData struct {
	Operator  common.Address
	Key       [32]byte
	VaultData []IValidatorDataVaultDataSet
}

// IValidatorDataValidatorDataShort is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataValidatorDataShort struct {
	Operator common.Address
	Key      [32]byte
	Stake    *big.Int
}

// IValidatorDataVaultDataSet is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataVaultDataSet struct {
	Vault           common.Address
	Stake           *big.Int
	PowerExpiresAt  *big.Int
	UnbondingTime   *big.Int
	UnbondingAmount *big.Int
}

// MiddlewareMetaData contains all meta data concerning the Middleware contract.
var MiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BaseMiddleware_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EpochCapture_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyManager256_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Operators_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OzAccessControl_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SharedVaults_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCaptureTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpochDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpochStart\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRole\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPower\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.ValidatorDataShort[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSetWithUnbonding\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.ValidatorData[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vaultData\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.VaultDataSet[]\",\"components\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"powerExpiresAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"unbondingTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"unbondingAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultOracle\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashingWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"vaultRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorNetOptin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyWasActiveAt\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"key_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorByKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinstake\",\"inputs\":[{\"name\":\"minStake_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeToPower\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"power\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralOracleAdded\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralOracleRemoved\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InstantSlash\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeChanged\",\"inputs\":[{\"name\":\"oldMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SelectorRoleSet\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VetoSlash\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUseForVaults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUsedOtherOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CollateralWithOracleAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddressCollateralOrOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Enabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImmutablePeriodNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InactiveVaultSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDisabledKeysReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonVetoSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorSpecificVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSupportedCollateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotOptedIn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorStakeIsSmall\",\"inputs\":[{\"name\":\"minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooOldTimestampSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownSlasherType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultEpochTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultNotInitialized\",\"inputs\":[]}]",
}

// MiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use MiddlewareMetaData.ABI instead.
var MiddlewareABI = MiddlewareMetaData.ABI

// Middleware is an auto generated Go binding around an Ethereum contract.
type Middleware struct {
	MiddlewareCaller     // Read-only binding to the contract
	MiddlewareTransactor // Write-only binding to the contract
	MiddlewareFilterer   // Log filterer for contract events
}

// MiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiddlewareSession struct {
	Contract     *Middleware       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiddlewareCallerSession struct {
	Contract *MiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiddlewareTransactorSession struct {
	Contract     *MiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiddlewareRaw struct {
	Contract *Middleware // Generic contract binding to access the raw methods on
}

// MiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiddlewareCallerRaw struct {
	Contract *MiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// MiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiddlewareTransactorRaw struct {
	Contract *MiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiddleware creates a new instance of Middleware, bound to a specific deployed contract.
func NewMiddleware(address common.Address, backend bind.ContractBackend) (*Middleware, error) {
	contract, err := bindMiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Middleware{MiddlewareCaller: MiddlewareCaller{contract: contract}, MiddlewareTransactor: MiddlewareTransactor{contract: contract}, MiddlewareFilterer: MiddlewareFilterer{contract: contract}}, nil
}

// NewMiddlewareCaller creates a new read-only instance of Middleware, bound to a specific deployed contract.
func NewMiddlewareCaller(address common.Address, caller bind.ContractCaller) (*MiddlewareCaller, error) {
	contract, err := bindMiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiddlewareCaller{contract: contract}, nil
}

// NewMiddlewareTransactor creates a new write-only instance of Middleware, bound to a specific deployed contract.
func NewMiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*MiddlewareTransactor, error) {
	contract, err := bindMiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiddlewareTransactor{contract: contract}, nil
}

// NewMiddlewareFilterer creates a new log filterer instance of Middleware, bound to a specific deployed contract.
func NewMiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*MiddlewareFilterer, error) {
	contract, err := bindMiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiddlewareFilterer{contract: contract}, nil
}

// bindMiddleware binds a generic wrapper to an already deployed contract.
func bindMiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Middleware *MiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Middleware.Contract.MiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Middleware *MiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Middleware.Contract.MiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Middleware *MiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Middleware.Contract.MiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Middleware *MiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Middleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Middleware *MiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Middleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Middleware *MiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Middleware.Contract.contract.Transact(opts, method, params...)
}

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) BaseMiddlewareVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "BaseMiddleware_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) BaseMiddlewareVERSION() (uint64, error) {
	return _Middleware.Contract.BaseMiddlewareVERSION(&_Middleware.CallOpts)
}

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) BaseMiddlewareVERSION() (uint64, error) {
	return _Middleware.Contract.BaseMiddlewareVERSION(&_Middleware.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Middleware.Contract.DEFAULTADMINROLE(&_Middleware.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Middleware.Contract.DEFAULTADMINROLE(&_Middleware.CallOpts)
}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) EpochCaptureVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "EpochCapture_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) EpochCaptureVERSION() (uint64, error) {
	return _Middleware.Contract.EpochCaptureVERSION(&_Middleware.CallOpts)
}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) EpochCaptureVERSION() (uint64, error) {
	return _Middleware.Contract.EpochCaptureVERSION(&_Middleware.CallOpts)
}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) KeyManager256VERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "KeyManager256_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) KeyManager256VERSION() (uint64, error) {
	return _Middleware.Contract.KeyManager256VERSION(&_Middleware.CallOpts)
}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) KeyManager256VERSION() (uint64, error) {
	return _Middleware.Contract.KeyManager256VERSION(&_Middleware.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareSession) OPERATORROLE() ([32]byte, error) {
	return _Middleware.Contract.OPERATORROLE(&_Middleware.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Middleware *MiddlewareCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Middleware.Contract.OPERATORROLE(&_Middleware.CallOpts)
}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) OperatorsVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "Operators_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) OperatorsVERSION() (uint64, error) {
	return _Middleware.Contract.OperatorsVERSION(&_Middleware.CallOpts)
}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) OperatorsVERSION() (uint64, error) {
	return _Middleware.Contract.OperatorsVERSION(&_Middleware.CallOpts)
}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) OzAccessControlVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "OzAccessControl_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) OzAccessControlVERSION() (uint64, error) {
	return _Middleware.Contract.OzAccessControlVERSION(&_Middleware.CallOpts)
}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) OzAccessControlVERSION() (uint64, error) {
	return _Middleware.Contract.OzAccessControlVERSION(&_Middleware.CallOpts)
}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCaller) SharedVaultsVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "SharedVaults_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Middleware *MiddlewareSession) SharedVaultsVERSION() (uint64, error) {
	return _Middleware.Contract.SharedVaultsVERSION(&_Middleware.CallOpts)
}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Middleware *MiddlewareCallerSession) SharedVaultsVERSION() (uint64, error) {
	return _Middleware.Contract.SharedVaultsVERSION(&_Middleware.CallOpts)
}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Middleware *MiddlewareCaller) CollateralOracle(opts *bind.CallOpts, collateral common.Address) (common.Address, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "collateralOracle", collateral)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Middleware *MiddlewareSession) CollateralOracle(collateral common.Address) (common.Address, error) {
	return _Middleware.Contract.CollateralOracle(&_Middleware.CallOpts, collateral)
}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Middleware *MiddlewareCallerSession) CollateralOracle(collateral common.Address) (common.Address, error) {
	return _Middleware.Contract.CollateralOracle(&_Middleware.CallOpts, collateral)
}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Middleware *MiddlewareCaller) GetCaptureTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getCaptureTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Middleware *MiddlewareSession) GetCaptureTimestamp() (*big.Int, error) {
	return _Middleware.Contract.GetCaptureTimestamp(&_Middleware.CallOpts)
}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Middleware *MiddlewareCallerSession) GetCaptureTimestamp() (*big.Int, error) {
	return _Middleware.Contract.GetCaptureTimestamp(&_Middleware.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Middleware *MiddlewareCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Middleware *MiddlewareSession) GetCurrentEpoch() (*big.Int, error) {
	return _Middleware.Contract.GetCurrentEpoch(&_Middleware.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Middleware *MiddlewareCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Middleware.Contract.GetCurrentEpoch(&_Middleware.CallOpts)
}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Middleware *MiddlewareCaller) GetEpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getEpochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Middleware *MiddlewareSession) GetEpochDuration() (*big.Int, error) {
	return _Middleware.Contract.GetEpochDuration(&_Middleware.CallOpts)
}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Middleware *MiddlewareCallerSession) GetEpochDuration() (*big.Int, error) {
	return _Middleware.Contract.GetEpochDuration(&_Middleware.CallOpts)
}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Middleware *MiddlewareCaller) GetEpochStart(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getEpochStart", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Middleware *MiddlewareSession) GetEpochStart(epoch *big.Int) (*big.Int, error) {
	return _Middleware.Contract.GetEpochStart(&_Middleware.CallOpts, epoch)
}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Middleware *MiddlewareCallerSession) GetEpochStart(epoch *big.Int) (*big.Int, error) {
	return _Middleware.Contract.GetEpochStart(&_Middleware.CallOpts, epoch)
}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Middleware *MiddlewareCaller) GetRole(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getRole", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Middleware *MiddlewareSession) GetRole(selector [4]byte) ([32]byte, error) {
	return _Middleware.Contract.GetRole(&_Middleware.CallOpts, selector)
}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Middleware *MiddlewareCallerSession) GetRole(selector [4]byte) ([32]byte, error) {
	return _Middleware.Contract.GetRole(&_Middleware.CallOpts, selector)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Middleware *MiddlewareCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Middleware *MiddlewareSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Middleware.Contract.GetRoleAdmin(&_Middleware.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Middleware *MiddlewareCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Middleware.Contract.GetRoleAdmin(&_Middleware.CallOpts, role)
}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Middleware *MiddlewareCaller) GetTotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getTotalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Middleware *MiddlewareSession) GetTotalPower() (*big.Int, error) {
	return _Middleware.Contract.GetTotalPower(&_Middleware.CallOpts)
}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Middleware *MiddlewareCallerSession) GetTotalPower() (*big.Int, error) {
	return _Middleware.Contract.GetTotalPower(&_Middleware.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Middleware *MiddlewareCaller) GetValidatorSet(opts *bind.CallOpts) ([]IValidatorDataValidatorDataShort, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getValidatorSet")

	if err != nil {
		return *new([]IValidatorDataValidatorDataShort), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorDataValidatorDataShort)).(*[]IValidatorDataValidatorDataShort)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Middleware *MiddlewareSession) GetValidatorSet() ([]IValidatorDataValidatorDataShort, error) {
	return _Middleware.Contract.GetValidatorSet(&_Middleware.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Middleware *MiddlewareCallerSession) GetValidatorSet() ([]IValidatorDataValidatorDataShort, error) {
	return _Middleware.Contract.GetValidatorSet(&_Middleware.CallOpts)
}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Middleware *MiddlewareCaller) GetValidatorSetWithUnbonding(opts *bind.CallOpts) ([]IValidatorDataValidatorData, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getValidatorSetWithUnbonding")

	if err != nil {
		return *new([]IValidatorDataValidatorData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorDataValidatorData)).(*[]IValidatorDataValidatorData)

	return out0, err

}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Middleware *MiddlewareSession) GetValidatorSetWithUnbonding() ([]IValidatorDataValidatorData, error) {
	return _Middleware.Contract.GetValidatorSetWithUnbonding(&_Middleware.CallOpts)
}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Middleware *MiddlewareCallerSession) GetValidatorSetWithUnbonding() ([]IValidatorDataValidatorData, error) {
	return _Middleware.Contract.GetValidatorSetWithUnbonding(&_Middleware.CallOpts)
}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Middleware *MiddlewareCaller) GetVaultOracle(opts *bind.CallOpts, vault common.Address) (common.Address, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "getVaultOracle", vault)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Middleware *MiddlewareSession) GetVaultOracle(vault common.Address) (common.Address, error) {
	return _Middleware.Contract.GetVaultOracle(&_Middleware.CallOpts, vault)
}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Middleware *MiddlewareCallerSession) GetVaultOracle(vault common.Address) (common.Address, error) {
	return _Middleware.Contract.GetVaultOracle(&_Middleware.CallOpts, vault)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Middleware *MiddlewareCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Middleware *MiddlewareSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Middleware.Contract.HasRole(&_Middleware.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Middleware *MiddlewareCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Middleware.Contract.HasRole(&_Middleware.CallOpts, role, account)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_Middleware *MiddlewareCaller) IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "isOperatorRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_Middleware *MiddlewareSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _Middleware.Contract.IsOperatorRegistered(&_Middleware.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_Middleware *MiddlewareCallerSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _Middleware.Contract.IsOperatorRegistered(&_Middleware.CallOpts, operator)
}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Middleware *MiddlewareCaller) KeyWasActiveAt(opts *bind.CallOpts, timestamp *big.Int, key_ []byte) (bool, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "keyWasActiveAt", timestamp, key_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Middleware *MiddlewareSession) KeyWasActiveAt(timestamp *big.Int, key_ []byte) (bool, error) {
	return _Middleware.Contract.KeyWasActiveAt(&_Middleware.CallOpts, timestamp, key_)
}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Middleware *MiddlewareCallerSession) KeyWasActiveAt(timestamp *big.Int, key_ []byte) (bool, error) {
	return _Middleware.Contract.KeyWasActiveAt(&_Middleware.CallOpts, timestamp, key_)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Middleware *MiddlewareCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Middleware *MiddlewareSession) MinStake() (*big.Int, error) {
	return _Middleware.Contract.MinStake(&_Middleware.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Middleware *MiddlewareCallerSession) MinStake() (*big.Int, error) {
	return _Middleware.Contract.MinStake(&_Middleware.CallOpts)
}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Middleware *MiddlewareCaller) OperatorByKey(opts *bind.CallOpts, key []byte) (common.Address, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "operatorByKey", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Middleware *MiddlewareSession) OperatorByKey(key []byte) (common.Address, error) {
	return _Middleware.Contract.OperatorByKey(&_Middleware.CallOpts, key)
}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Middleware *MiddlewareCallerSession) OperatorByKey(key []byte) (common.Address, error) {
	return _Middleware.Contract.OperatorByKey(&_Middleware.CallOpts, key)
}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Middleware *MiddlewareCaller) OperatorKey(opts *bind.CallOpts, operator common.Address) ([]byte, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "operatorKey", operator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Middleware *MiddlewareSession) OperatorKey(operator common.Address) ([]byte, error) {
	return _Middleware.Contract.OperatorKey(&_Middleware.CallOpts, operator)
}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Middleware *MiddlewareCallerSession) OperatorKey(operator common.Address) ([]byte, error) {
	return _Middleware.Contract.OperatorKey(&_Middleware.CallOpts, operator)
}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Middleware *MiddlewareCaller) StakeToPower(opts *bind.CallOpts, vault common.Address, stake *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Middleware.contract.Call(opts, &out, "stakeToPower", vault, stake)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Middleware *MiddlewareSession) StakeToPower(vault common.Address, stake *big.Int) (*big.Int, error) {
	return _Middleware.Contract.StakeToPower(&_Middleware.CallOpts, vault, stake)
}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Middleware *MiddlewareCallerSession) StakeToPower(vault common.Address, stake *big.Int) (*big.Int, error) {
	return _Middleware.Contract.StakeToPower(&_Middleware.CallOpts, vault, stake)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Middleware *MiddlewareTransactor) AddCollateralOracle(opts *bind.TransactOpts, collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "addCollateralOracle", collateral, oracle)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Middleware *MiddlewareSession) AddCollateralOracle(collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.AddCollateralOracle(&_Middleware.TransactOpts, collateral, oracle)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Middleware *MiddlewareTransactorSession) AddCollateralOracle(collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.AddCollateralOracle(&_Middleware.TransactOpts, collateral, oracle)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.GrantRole(&_Middleware.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.GrantRole(&_Middleware.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Middleware *MiddlewareTransactor) Initialize(opts *bind.TransactOpts, network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "initialize", network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Middleware *MiddlewareSession) Initialize(network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Middleware.Contract.Initialize(&_Middleware.TransactOpts, network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Middleware *MiddlewareTransactorSession) Initialize(network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Middleware.Contract.Initialize(&_Middleware.TransactOpts, network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Middleware *MiddlewareTransactor) PauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "pauseOperator", operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Middleware *MiddlewareSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseOperator(&_Middleware.TransactOpts, operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Middleware *MiddlewareTransactorSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseOperator(&_Middleware.TransactOpts, operator)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactor) PauseOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "pauseOperatorVault", operator, vault)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareSession) PauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactorSession) PauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactor) PauseSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "pauseSharedVault", sharedVault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareSession) PauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactorSession) PauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.PauseSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Middleware *MiddlewareTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "registerOperator", operator, key, vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Middleware *MiddlewareSession) RegisterOperator(operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterOperator(&_Middleware.TransactOpts, operator, key, vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Middleware *MiddlewareTransactorSession) RegisterOperator(operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterOperator(&_Middleware.TransactOpts, operator, key, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactor) RegisterOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "registerOperatorVault", operator, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareSession) RegisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactorSession) RegisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactor) RegisterSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "registerSharedVault", sharedVault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareSession) RegisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactorSession) RegisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RegisterSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Middleware *MiddlewareTransactor) RemoveCollateralOracle(opts *bind.TransactOpts, collateral common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "removeCollateralOracle", collateral)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Middleware *MiddlewareSession) RemoveCollateralOracle(collateral common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RemoveCollateralOracle(&_Middleware.TransactOpts, collateral)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Middleware *MiddlewareTransactorSession) RemoveCollateralOracle(collateral common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RemoveCollateralOracle(&_Middleware.TransactOpts, collateral)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Middleware *MiddlewareTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Middleware *MiddlewareSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RenounceRole(&_Middleware.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Middleware *MiddlewareTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RenounceRole(&_Middleware.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RevokeRole(&_Middleware.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Middleware *MiddlewareTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.RevokeRole(&_Middleware.TransactOpts, role, account)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Middleware *MiddlewareTransactor) SetMinstake(opts *bind.TransactOpts, minStake_ *big.Int) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "setMinstake", minStake_)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Middleware *MiddlewareSession) SetMinstake(minStake_ *big.Int) (*types.Transaction, error) {
	return _Middleware.Contract.SetMinstake(&_Middleware.TransactOpts, minStake_)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Middleware *MiddlewareTransactorSession) SetMinstake(minStake_ *big.Int) (*types.Transaction, error) {
	return _Middleware.Contract.SetMinstake(&_Middleware.TransactOpts, minStake_)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Middleware *MiddlewareTransactor) UnpauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unpauseOperator", operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Middleware *MiddlewareSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseOperator(&_Middleware.TransactOpts, operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Middleware *MiddlewareTransactorSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseOperator(&_Middleware.TransactOpts, operator)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactor) UnpauseOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unpauseOperatorVault", operator, vault)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareSession) UnpauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactorSession) UnpauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactor) UnpauseSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unpauseSharedVault", sharedVault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareSession) UnpauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactorSession) UnpauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnpauseSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Middleware *MiddlewareTransactor) UnregisterOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unregisterOperator", operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Middleware *MiddlewareSession) UnregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterOperator(&_Middleware.TransactOpts, operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Middleware *MiddlewareTransactorSession) UnregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterOperator(&_Middleware.TransactOpts, operator)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactor) UnregisterOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unregisterOperatorVault", operator, vault)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareSession) UnregisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Middleware *MiddlewareTransactorSession) UnregisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterOperatorVault(&_Middleware.TransactOpts, operator, vault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactor) UnregisterSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "unregisterSharedVault", sharedVault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareSession) UnregisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Middleware *MiddlewareTransactorSession) UnregisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Middleware.Contract.UnregisterSharedVault(&_Middleware.TransactOpts, sharedVault)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Middleware *MiddlewareTransactor) UpdateOperatorKey(opts *bind.TransactOpts, operator common.Address, key []byte) (*types.Transaction, error) {
	return _Middleware.contract.Transact(opts, "updateOperatorKey", operator, key)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Middleware *MiddlewareSession) UpdateOperatorKey(operator common.Address, key []byte) (*types.Transaction, error) {
	return _Middleware.Contract.UpdateOperatorKey(&_Middleware.TransactOpts, operator, key)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Middleware *MiddlewareTransactorSession) UpdateOperatorKey(operator common.Address, key []byte) (*types.Transaction, error) {
	return _Middleware.Contract.UpdateOperatorKey(&_Middleware.TransactOpts, operator, key)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Middleware *MiddlewareTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Middleware.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Middleware *MiddlewareSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Middleware.Contract.Fallback(&_Middleware.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Middleware *MiddlewareTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Middleware.Contract.Fallback(&_Middleware.TransactOpts, calldata)
}

// MiddlewareCollateralOracleAddedIterator is returned from FilterCollateralOracleAdded and is used to iterate over the raw logs and unpacked data for CollateralOracleAdded events raised by the Middleware contract.
type MiddlewareCollateralOracleAddedIterator struct {
	Event *MiddlewareCollateralOracleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareCollateralOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareCollateralOracleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareCollateralOracleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareCollateralOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareCollateralOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareCollateralOracleAdded represents a CollateralOracleAdded event raised by the Middleware contract.
type MiddlewareCollateralOracleAdded struct {
	Collateral common.Address
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralOracleAdded is a free log retrieval operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) FilterCollateralOracleAdded(opts *bind.FilterOpts, collateral []common.Address, oracle []common.Address) (*MiddlewareCollateralOracleAddedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "CollateralOracleAdded", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareCollateralOracleAddedIterator{contract: _Middleware.contract, event: "CollateralOracleAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralOracleAdded is a free log subscription operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) WatchCollateralOracleAdded(opts *bind.WatchOpts, sink chan<- *MiddlewareCollateralOracleAdded, collateral []common.Address, oracle []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "CollateralOracleAdded", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareCollateralOracleAdded)
				if err := _Middleware.contract.UnpackLog(event, "CollateralOracleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralOracleAdded is a log parse operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) ParseCollateralOracleAdded(log types.Log) (*MiddlewareCollateralOracleAdded, error) {
	event := new(MiddlewareCollateralOracleAdded)
	if err := _Middleware.contract.UnpackLog(event, "CollateralOracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareCollateralOracleRemovedIterator is returned from FilterCollateralOracleRemoved and is used to iterate over the raw logs and unpacked data for CollateralOracleRemoved events raised by the Middleware contract.
type MiddlewareCollateralOracleRemovedIterator struct {
	Event *MiddlewareCollateralOracleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareCollateralOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareCollateralOracleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareCollateralOracleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareCollateralOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareCollateralOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareCollateralOracleRemoved represents a CollateralOracleRemoved event raised by the Middleware contract.
type MiddlewareCollateralOracleRemoved struct {
	Collateral common.Address
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralOracleRemoved is a free log retrieval operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) FilterCollateralOracleRemoved(opts *bind.FilterOpts, collateral []common.Address, oracle []common.Address) (*MiddlewareCollateralOracleRemovedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "CollateralOracleRemoved", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareCollateralOracleRemovedIterator{contract: _Middleware.contract, event: "CollateralOracleRemoved", logs: logs, sub: sub}, nil
}

// WatchCollateralOracleRemoved is a free log subscription operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) WatchCollateralOracleRemoved(opts *bind.WatchOpts, sink chan<- *MiddlewareCollateralOracleRemoved, collateral []common.Address, oracle []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "CollateralOracleRemoved", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareCollateralOracleRemoved)
				if err := _Middleware.contract.UnpackLog(event, "CollateralOracleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralOracleRemoved is a log parse operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Middleware *MiddlewareFilterer) ParseCollateralOracleRemoved(log types.Log) (*MiddlewareCollateralOracleRemoved, error) {
	event := new(MiddlewareCollateralOracleRemoved)
	if err := _Middleware.contract.UnpackLog(event, "CollateralOracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Middleware contract.
type MiddlewareInitializedIterator struct {
	Event *MiddlewareInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareInitialized represents a Initialized event raised by the Middleware contract.
type MiddlewareInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Middleware *MiddlewareFilterer) FilterInitialized(opts *bind.FilterOpts) (*MiddlewareInitializedIterator, error) {

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MiddlewareInitializedIterator{contract: _Middleware.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Middleware *MiddlewareFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MiddlewareInitialized) (event.Subscription, error) {

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareInitialized)
				if err := _Middleware.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Middleware *MiddlewareFilterer) ParseInitialized(log types.Log) (*MiddlewareInitialized, error) {
	event := new(MiddlewareInitialized)
	if err := _Middleware.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareInstantSlashIterator is returned from FilterInstantSlash and is used to iterate over the raw logs and unpacked data for InstantSlash events raised by the Middleware contract.
type MiddlewareInstantSlashIterator struct {
	Event *MiddlewareInstantSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareInstantSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareInstantSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareInstantSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareInstantSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareInstantSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareInstantSlash represents a InstantSlash event raised by the Middleware contract.
type MiddlewareInstantSlash struct {
	Vault      common.Address
	Subnetwork [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInstantSlash is a free log retrieval operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Middleware *MiddlewareFilterer) FilterInstantSlash(opts *bind.FilterOpts) (*MiddlewareInstantSlashIterator, error) {

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "InstantSlash")
	if err != nil {
		return nil, err
	}
	return &MiddlewareInstantSlashIterator{contract: _Middleware.contract, event: "InstantSlash", logs: logs, sub: sub}, nil
}

// WatchInstantSlash is a free log subscription operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Middleware *MiddlewareFilterer) WatchInstantSlash(opts *bind.WatchOpts, sink chan<- *MiddlewareInstantSlash) (event.Subscription, error) {

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "InstantSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareInstantSlash)
				if err := _Middleware.contract.UnpackLog(event, "InstantSlash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInstantSlash is a log parse operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Middleware *MiddlewareFilterer) ParseInstantSlash(log types.Log) (*MiddlewareInstantSlash, error) {
	event := new(MiddlewareInstantSlash)
	if err := _Middleware.contract.UnpackLog(event, "InstantSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareMinStakeChangedIterator is returned from FilterMinStakeChanged and is used to iterate over the raw logs and unpacked data for MinStakeChanged events raised by the Middleware contract.
type MiddlewareMinStakeChangedIterator struct {
	Event *MiddlewareMinStakeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareMinStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareMinStakeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareMinStakeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareMinStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareMinStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareMinStakeChanged represents a MinStakeChanged event raised by the Middleware contract.
type MiddlewareMinStakeChanged struct {
	OldMinstake *big.Int
	NewMinstake *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinStakeChanged is a free log retrieval operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Middleware *MiddlewareFilterer) FilterMinStakeChanged(opts *bind.FilterOpts) (*MiddlewareMinStakeChangedIterator, error) {

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return &MiddlewareMinStakeChangedIterator{contract: _Middleware.contract, event: "MinStakeChanged", logs: logs, sub: sub}, nil
}

// WatchMinStakeChanged is a free log subscription operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Middleware *MiddlewareFilterer) WatchMinStakeChanged(opts *bind.WatchOpts, sink chan<- *MiddlewareMinStakeChanged) (event.Subscription, error) {

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareMinStakeChanged)
				if err := _Middleware.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinStakeChanged is a log parse operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Middleware *MiddlewareFilterer) ParseMinStakeChanged(log types.Log) (*MiddlewareMinStakeChanged, error) {
	event := new(MiddlewareMinStakeChanged)
	if err := _Middleware.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Middleware contract.
type MiddlewareRoleAdminChangedIterator struct {
	Event *MiddlewareRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareRoleAdminChanged represents a RoleAdminChanged event raised by the Middleware contract.
type MiddlewareRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Middleware *MiddlewareFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MiddlewareRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareRoleAdminChangedIterator{contract: _Middleware.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Middleware *MiddlewareFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MiddlewareRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareRoleAdminChanged)
				if err := _Middleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Middleware *MiddlewareFilterer) ParseRoleAdminChanged(log types.Log) (*MiddlewareRoleAdminChanged, error) {
	event := new(MiddlewareRoleAdminChanged)
	if err := _Middleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Middleware contract.
type MiddlewareRoleGrantedIterator struct {
	Event *MiddlewareRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareRoleGranted represents a RoleGranted event raised by the Middleware contract.
type MiddlewareRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MiddlewareRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareRoleGrantedIterator{contract: _Middleware.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MiddlewareRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareRoleGranted)
				if err := _Middleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) ParseRoleGranted(log types.Log) (*MiddlewareRoleGranted, error) {
	event := new(MiddlewareRoleGranted)
	if err := _Middleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Middleware contract.
type MiddlewareRoleRevokedIterator struct {
	Event *MiddlewareRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareRoleRevoked represents a RoleRevoked event raised by the Middleware contract.
type MiddlewareRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MiddlewareRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareRoleRevokedIterator{contract: _Middleware.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MiddlewareRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareRoleRevoked)
				if err := _Middleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Middleware *MiddlewareFilterer) ParseRoleRevoked(log types.Log) (*MiddlewareRoleRevoked, error) {
	event := new(MiddlewareRoleRevoked)
	if err := _Middleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareSelectorRoleSetIterator is returned from FilterSelectorRoleSet and is used to iterate over the raw logs and unpacked data for SelectorRoleSet events raised by the Middleware contract.
type MiddlewareSelectorRoleSetIterator struct {
	Event *MiddlewareSelectorRoleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareSelectorRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareSelectorRoleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareSelectorRoleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareSelectorRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareSelectorRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareSelectorRoleSet represents a SelectorRoleSet event raised by the Middleware contract.
type MiddlewareSelectorRoleSet struct {
	Selector [4]byte
	Role     [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSelectorRoleSet is a free log retrieval operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Middleware *MiddlewareFilterer) FilterSelectorRoleSet(opts *bind.FilterOpts, selector [][4]byte, role [][32]byte) (*MiddlewareSelectorRoleSetIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "SelectorRoleSet", selectorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &MiddlewareSelectorRoleSetIterator{contract: _Middleware.contract, event: "SelectorRoleSet", logs: logs, sub: sub}, nil
}

// WatchSelectorRoleSet is a free log subscription operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Middleware *MiddlewareFilterer) WatchSelectorRoleSet(opts *bind.WatchOpts, sink chan<- *MiddlewareSelectorRoleSet, selector [][4]byte, role [][32]byte) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "SelectorRoleSet", selectorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareSelectorRoleSet)
				if err := _Middleware.contract.UnpackLog(event, "SelectorRoleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSelectorRoleSet is a log parse operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Middleware *MiddlewareFilterer) ParseSelectorRoleSet(log types.Log) (*MiddlewareSelectorRoleSet, error) {
	event := new(MiddlewareSelectorRoleSet)
	if err := _Middleware.contract.UnpackLog(event, "SelectorRoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiddlewareVetoSlashIterator is returned from FilterVetoSlash and is used to iterate over the raw logs and unpacked data for VetoSlash events raised by the Middleware contract.
type MiddlewareVetoSlashIterator struct {
	Event *MiddlewareVetoSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiddlewareVetoSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiddlewareVetoSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MiddlewareVetoSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MiddlewareVetoSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiddlewareVetoSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiddlewareVetoSlash represents a VetoSlash event raised by the Middleware contract.
type MiddlewareVetoSlash struct {
	Vault      common.Address
	Subnetwork [32]byte
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVetoSlash is a free log retrieval operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Middleware *MiddlewareFilterer) FilterVetoSlash(opts *bind.FilterOpts) (*MiddlewareVetoSlashIterator, error) {

	logs, sub, err := _Middleware.contract.FilterLogs(opts, "VetoSlash")
	if err != nil {
		return nil, err
	}
	return &MiddlewareVetoSlashIterator{contract: _Middleware.contract, event: "VetoSlash", logs: logs, sub: sub}, nil
}

// WatchVetoSlash is a free log subscription operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Middleware *MiddlewareFilterer) WatchVetoSlash(opts *bind.WatchOpts, sink chan<- *MiddlewareVetoSlash) (event.Subscription, error) {

	logs, sub, err := _Middleware.contract.WatchLogs(opts, "VetoSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiddlewareVetoSlash)
				if err := _Middleware.contract.UnpackLog(event, "VetoSlash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVetoSlash is a log parse operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Middleware *MiddlewareFilterer) ParseVetoSlash(log types.Log) (*MiddlewareVetoSlash, error) {
	event := new(MiddlewareVetoSlash)
	if err := _Middleware.contract.UnpackLog(event, "VetoSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
