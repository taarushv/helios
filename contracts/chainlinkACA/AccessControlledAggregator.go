// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlinkACA

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChainlinkACAABI is the input ABI used to generate the binding from.
const ChainlinkACAABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_minSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_maxSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AddedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AvailableFundsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdateRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"OraclePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemovedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"RequesterPermissionsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"paymentAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"minSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"maxSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"restartDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timeout\",\"type\":\"uint32\"}],\"name\":\"RoundDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"submission\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"round\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SubmissionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocatedFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removed\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_added\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_addedAdmins\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"}],\"name\":\"changeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_queriedRoundId\",\"type\":\"uint32\"}],\"name\":\"oracleRoundState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_eligibleToSubmit\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_roundId\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"_latestSubmission\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"_startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_availableFunds\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_oracleCount\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restartDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_delay\",\"type\":\"uint32\"}],\"name\":\"setRequesterPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newValidator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_submission\",\"type\":\"int256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAvailableFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"}],\"name\":\"updateFutureRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"withdrawablePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChainlinkACA is an auto generated Go binding around an Ethereum contract.
type ChainlinkACA struct {
	ChainlinkACACaller     // Read-only binding to the contract
	ChainlinkACATransactor // Write-only binding to the contract
	ChainlinkACAFilterer   // Log filterer for contract events
}

// ChainlinkACACaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkACACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkACATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkACATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkACAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkACAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkACASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkACASession struct {
	Contract     *ChainlinkACA     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkACACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkACACallerSession struct {
	Contract *ChainlinkACACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChainlinkACATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkACATransactorSession struct {
	Contract     *ChainlinkACATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChainlinkACARaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkACARaw struct {
	Contract *ChainlinkACA // Generic contract binding to access the raw methods on
}

// ChainlinkACACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkACACallerRaw struct {
	Contract *ChainlinkACACaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkACATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkACATransactorRaw struct {
	Contract *ChainlinkACATransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkACA creates a new instance of ChainlinkACA, bound to a specific deployed contract.
func NewChainlinkACA(address common.Address, backend bind.ContractBackend) (*ChainlinkACA, error) {
	contract, err := bindChainlinkACA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACA{ChainlinkACACaller: ChainlinkACACaller{contract: contract}, ChainlinkACATransactor: ChainlinkACATransactor{contract: contract}, ChainlinkACAFilterer: ChainlinkACAFilterer{contract: contract}}, nil
}

// NewChainlinkACACaller creates a new read-only instance of ChainlinkACA, bound to a specific deployed contract.
func NewChainlinkACACaller(address common.Address, caller bind.ContractCaller) (*ChainlinkACACaller, error) {
	contract, err := bindChainlinkACA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACACaller{contract: contract}, nil
}

// NewChainlinkACATransactor creates a new write-only instance of ChainlinkACA, bound to a specific deployed contract.
func NewChainlinkACATransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkACATransactor, error) {
	contract, err := bindChainlinkACA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACATransactor{contract: contract}, nil
}

// NewChainlinkACAFilterer creates a new log filterer instance of ChainlinkACA, bound to a specific deployed contract.
func NewChainlinkACAFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkACAFilterer, error) {
	contract, err := bindChainlinkACA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAFilterer{contract: contract}, nil
}

// bindChainlinkACA binds a generic wrapper to an already deployed contract.
func bindChainlinkACA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkACAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkACA *ChainlinkACARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkACA.Contract.ChainlinkACACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkACA *ChainlinkACARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.ChainlinkACATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkACA *ChainlinkACARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.ChainlinkACATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkACA *ChainlinkACACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkACA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkACA *ChainlinkACATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkACA *ChainlinkACATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.contract.Transact(opts, method, params...)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACaller) AllocatedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "allocatedFunds")
	return *ret0, err
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACASession) AllocatedFunds() (*big.Int, error) {
	return _ChainlinkACA.Contract.AllocatedFunds(&_ChainlinkACA.CallOpts)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACallerSession) AllocatedFunds() (*big.Int, error) {
	return _ChainlinkACA.Contract.AllocatedFunds(&_ChainlinkACA.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACaller) AvailableFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "availableFunds")
	return *ret0, err
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACASession) AvailableFunds() (*big.Int, error) {
	return _ChainlinkACA.Contract.AvailableFunds(&_ChainlinkACA.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACallerSession) AvailableFunds() (*big.Int, error) {
	return _ChainlinkACA.Contract.AvailableFunds(&_ChainlinkACA.CallOpts)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_ChainlinkACA *ChainlinkACACaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "checkEnabled")
	return *ret0, err
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_ChainlinkACA *ChainlinkACASession) CheckEnabled() (bool, error) {
	return _ChainlinkACA.Contract.CheckEnabled(&_ChainlinkACA.CallOpts)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_ChainlinkACA *ChainlinkACACallerSession) CheckEnabled() (bool, error) {
	return _ChainlinkACA.Contract.CheckEnabled(&_ChainlinkACA.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkACA *ChainlinkACACaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkACA *ChainlinkACASession) Decimals() (uint8, error) {
	return _ChainlinkACA.Contract.Decimals(&_ChainlinkACA.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkACA *ChainlinkACACallerSession) Decimals() (uint8, error) {
	return _ChainlinkACA.Contract.Decimals(&_ChainlinkACA.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkACA *ChainlinkACACaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkACA *ChainlinkACASession) Description() (string, error) {
	return _ChainlinkACA.Contract.Description(&_ChainlinkACA.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkACA *ChainlinkACACallerSession) Description() (string, error) {
	return _ChainlinkACA.Contract.Description(&_ChainlinkACA.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_ChainlinkACA *ChainlinkACACaller) GetAdmin(opts *bind.CallOpts, _oracle common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "getAdmin", _oracle)
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_ChainlinkACA *ChainlinkACASession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _ChainlinkACA.Contract.GetAdmin(&_ChainlinkACA.CallOpts, _oracle)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_ChainlinkACA *ChainlinkACACallerSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _ChainlinkACA.Contract.GetAdmin(&_ChainlinkACA.CallOpts, _oracle)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkACA *ChainlinkACACaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "getAnswer", _roundId)
	return *ret0, err
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkACA *ChainlinkACASession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkACA.Contract.GetAnswer(&_ChainlinkACA.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkACA *ChainlinkACACallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkACA.Contract.GetAnswer(&_ChainlinkACA.CallOpts, _roundId)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ChainlinkACA *ChainlinkACACaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "getOracles")
	return *ret0, err
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ChainlinkACA *ChainlinkACASession) GetOracles() ([]common.Address, error) {
	return _ChainlinkACA.Contract.GetOracles(&_ChainlinkACA.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ChainlinkACA *ChainlinkACACallerSession) GetOracles() ([]common.Address, error) {
	return _ChainlinkACA.Contract.GetOracles(&_ChainlinkACA.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACACaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkACA.contract.Call(opts, out, "getRoundData", _roundId)
	return *ret, err
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACASession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkACA.Contract.GetRoundData(&_ChainlinkACA.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACACallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkACA.Contract.GetRoundData(&_ChainlinkACA.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkACA *ChainlinkACACaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "getTimestamp", _roundId)
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkACA *ChainlinkACASession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkACA.Contract.GetTimestamp(&_ChainlinkACA.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkACA *ChainlinkACACallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkACA.Contract.GetTimestamp(&_ChainlinkACA.CallOpts, _roundId)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_ChainlinkACA *ChainlinkACACaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "hasAccess", _user, _calldata)
	return *ret0, err
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_ChainlinkACA *ChainlinkACASession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _ChainlinkACA.Contract.HasAccess(&_ChainlinkACA.CallOpts, _user, _calldata)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_ChainlinkACA *ChainlinkACACallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _ChainlinkACA.Contract.HasAccess(&_ChainlinkACA.CallOpts, _user, _calldata)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkACA *ChainlinkACACaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "latestAnswer")
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkACA *ChainlinkACASession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestAnswer(&_ChainlinkACA.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkACA *ChainlinkACACallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestAnswer(&_ChainlinkACA.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "latestRound")
	return *ret0, err
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkACA *ChainlinkACASession) LatestRound() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestRound(&_ChainlinkACA.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACallerSession) LatestRound() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestRound(&_ChainlinkACA.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACACaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkACA.contract.Call(opts, out, "latestRoundData")
	return *ret, err
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACASession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkACA.Contract.LatestRoundData(&_ChainlinkACA.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkACA *ChainlinkACACallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkACA.Contract.LatestRoundData(&_ChainlinkACA.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "latestTimestamp")
	return *ret0, err
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkACA *ChainlinkACASession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestTimestamp(&_ChainlinkACA.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACallerSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkACA.Contract.LatestTimestamp(&_ChainlinkACA.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_ChainlinkACA *ChainlinkACACaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "linkToken")
	return *ret0, err
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_ChainlinkACA *ChainlinkACASession) LinkToken() (common.Address, error) {
	return _ChainlinkACA.Contract.LinkToken(&_ChainlinkACA.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_ChainlinkACA *ChainlinkACACallerSession) LinkToken() (common.Address, error) {
	return _ChainlinkACA.Contract.LinkToken(&_ChainlinkACA.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACaller) MaxSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "maxSubmissionCount")
	return *ret0, err
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACASession) MaxSubmissionCount() (uint32, error) {
	return _ChainlinkACA.Contract.MaxSubmissionCount(&_ChainlinkACA.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACallerSession) MaxSubmissionCount() (uint32, error) {
	return _ChainlinkACA.Contract.MaxSubmissionCount(&_ChainlinkACA.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACACaller) MaxSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "maxSubmissionValue")
	return *ret0, err
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACASession) MaxSubmissionValue() (*big.Int, error) {
	return _ChainlinkACA.Contract.MaxSubmissionValue(&_ChainlinkACA.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACACallerSession) MaxSubmissionValue() (*big.Int, error) {
	return _ChainlinkACA.Contract.MaxSubmissionValue(&_ChainlinkACA.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACaller) MinSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "minSubmissionCount")
	return *ret0, err
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACASession) MinSubmissionCount() (uint32, error) {
	return _ChainlinkACA.Contract.MinSubmissionCount(&_ChainlinkACA.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACallerSession) MinSubmissionCount() (uint32, error) {
	return _ChainlinkACA.Contract.MinSubmissionCount(&_ChainlinkACA.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACACaller) MinSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "minSubmissionValue")
	return *ret0, err
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACASession) MinSubmissionValue() (*big.Int, error) {
	return _ChainlinkACA.Contract.MinSubmissionValue(&_ChainlinkACA.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_ChainlinkACA *ChainlinkACACallerSession) MinSubmissionValue() (*big.Int, error) {
	return _ChainlinkACA.Contract.MinSubmissionValue(&_ChainlinkACA.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ChainlinkACA *ChainlinkACACaller) OracleCount(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "oracleCount")
	return *ret0, err
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ChainlinkACA *ChainlinkACASession) OracleCount() (uint8, error) {
	return _ChainlinkACA.Contract.OracleCount(&_ChainlinkACA.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ChainlinkACA *ChainlinkACACallerSession) OracleCount() (uint8, error) {
	return _ChainlinkACA.Contract.OracleCount(&_ChainlinkACA.CallOpts)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_ChainlinkACA *ChainlinkACACaller) OracleRoundState(opts *bind.CallOpts, _oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	ret := new(struct {
		EligibleToSubmit bool
		RoundId          uint32
		LatestSubmission *big.Int
		StartedAt        uint64
		Timeout          uint64
		AvailableFunds   *big.Int
		OracleCount      uint8
		PaymentAmount    *big.Int
	})
	out := ret
	err := _ChainlinkACA.contract.Call(opts, out, "oracleRoundState", _oracle, _queriedRoundId)
	return *ret, err
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_ChainlinkACA *ChainlinkACASession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _ChainlinkACA.Contract.OracleRoundState(&_ChainlinkACA.CallOpts, _oracle, _queriedRoundId)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_ChainlinkACA *ChainlinkACACallerSession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _ChainlinkACA.Contract.OracleRoundState(&_ChainlinkACA.CallOpts, _oracle, _queriedRoundId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkACA *ChainlinkACACaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkACA *ChainlinkACASession) Owner() (common.Address, error) {
	return _ChainlinkACA.Contract.Owner(&_ChainlinkACA.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkACA *ChainlinkACACallerSession) Owner() (common.Address, error) {
	return _ChainlinkACA.Contract.Owner(&_ChainlinkACA.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "paymentAmount")
	return *ret0, err
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkACA *ChainlinkACASession) PaymentAmount() (*big.Int, error) {
	return _ChainlinkACA.Contract.PaymentAmount(&_ChainlinkACA.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkACA *ChainlinkACACallerSession) PaymentAmount() (*big.Int, error) {
	return _ChainlinkACA.Contract.PaymentAmount(&_ChainlinkACA.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACaller) RestartDelay(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "restartDelay")
	return *ret0, err
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_ChainlinkACA *ChainlinkACASession) RestartDelay() (uint32, error) {
	return _ChainlinkACA.Contract.RestartDelay(&_ChainlinkACA.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACallerSession) RestartDelay() (uint32, error) {
	return _ChainlinkACA.Contract.RestartDelay(&_ChainlinkACA.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACaller) Timeout(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_ChainlinkACA *ChainlinkACASession) Timeout() (uint32, error) {
	return _ChainlinkACA.Contract.Timeout(&_ChainlinkACA.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_ChainlinkACA *ChainlinkACACallerSession) Timeout() (uint32, error) {
	return _ChainlinkACA.Contract.Timeout(&_ChainlinkACA.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_ChainlinkACA *ChainlinkACACaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "validator")
	return *ret0, err
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_ChainlinkACA *ChainlinkACASession) Validator() (common.Address, error) {
	return _ChainlinkACA.Contract.Validator(&_ChainlinkACA.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_ChainlinkACA *ChainlinkACACallerSession) Validator() (common.Address, error) {
	return _ChainlinkACA.Contract.Validator(&_ChainlinkACA.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkACA *ChainlinkACASession) Version() (*big.Int, error) {
	return _ChainlinkACA.Contract.Version(&_ChainlinkACA.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkACA *ChainlinkACACallerSession) Version() (*big.Int, error) {
	return _ChainlinkACA.Contract.Version(&_ChainlinkACA.CallOpts)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_ChainlinkACA *ChainlinkACACaller) WithdrawablePayment(opts *bind.CallOpts, _oracle common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkACA.contract.Call(opts, out, "withdrawablePayment", _oracle)
	return *ret0, err
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_ChainlinkACA *ChainlinkACASession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _ChainlinkACA.Contract.WithdrawablePayment(&_ChainlinkACA.CallOpts, _oracle)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_ChainlinkACA *ChainlinkACACallerSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _ChainlinkACA.Contract.WithdrawablePayment(&_ChainlinkACA.CallOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_ChainlinkACA *ChainlinkACATransactor) AcceptAdmin(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "acceptAdmin", _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_ChainlinkACA *ChainlinkACASession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AcceptAdmin(&_ChainlinkACA.TransactOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AcceptAdmin(&_ChainlinkACA.TransactOpts, _oracle)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkACA *ChainlinkACATransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkACA *ChainlinkACASession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AcceptOwnership(&_ChainlinkACA.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AcceptOwnership(&_ChainlinkACA.TransactOpts)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACATransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "addAccess", _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACASession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AddAccess(&_ChainlinkACA.TransactOpts, _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.AddAccess(&_ChainlinkACA.TransactOpts, _user)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_ChainlinkACA *ChainlinkACATransactor) ChangeOracles(opts *bind.TransactOpts, _removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "changeOracles", _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_ChainlinkACA *ChainlinkACASession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.ChangeOracles(&_ChainlinkACA.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.ChangeOracles(&_ChainlinkACA.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACATransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "disableAccessCheck")
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACASession) DisableAccessCheck() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.DisableAccessCheck(&_ChainlinkACA.TransactOpts)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.DisableAccessCheck(&_ChainlinkACA.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACATransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "enableAccessCheck")
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACASession) EnableAccessCheck() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.EnableAccessCheck(&_ChainlinkACA.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.EnableAccessCheck(&_ChainlinkACA.TransactOpts)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_ChainlinkACA *ChainlinkACATransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "onTokenTransfer", arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_ChainlinkACA *ChainlinkACASession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.OnTokenTransfer(&_ChainlinkACA.TransactOpts, arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.OnTokenTransfer(&_ChainlinkACA.TransactOpts, arg0, arg1, _data)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACATransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "removeAccess", _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACASession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.RemoveAccess(&_ChainlinkACA.TransactOpts, _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.RemoveAccess(&_ChainlinkACA.TransactOpts, _user)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_ChainlinkACA *ChainlinkACATransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_ChainlinkACA *ChainlinkACASession) RequestNewRound() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.RequestNewRound(&_ChainlinkACA.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_ChainlinkACA *ChainlinkACATransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.RequestNewRound(&_ChainlinkACA.TransactOpts)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_ChainlinkACA *ChainlinkACATransactor) SetRequesterPermissions(opts *bind.TransactOpts, _requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "setRequesterPermissions", _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_ChainlinkACA *ChainlinkACASession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.SetRequesterPermissions(&_ChainlinkACA.TransactOpts, _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.SetRequesterPermissions(&_ChainlinkACA.TransactOpts, _requester, _authorized, _delay)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_ChainlinkACA *ChainlinkACATransactor) SetValidator(opts *bind.TransactOpts, _newValidator common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "setValidator", _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_ChainlinkACA *ChainlinkACASession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.SetValidator(&_ChainlinkACA.TransactOpts, _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.SetValidator(&_ChainlinkACA.TransactOpts, _newValidator)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_ChainlinkACA *ChainlinkACATransactor) Submit(opts *bind.TransactOpts, _roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "submit", _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_ChainlinkACA *ChainlinkACASession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.Submit(&_ChainlinkACA.TransactOpts, _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.Submit(&_ChainlinkACA.TransactOpts, _roundId, _submission)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_ChainlinkACA *ChainlinkACATransactor) TransferAdmin(opts *bind.TransactOpts, _oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "transferAdmin", _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_ChainlinkACA *ChainlinkACASession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.TransferAdmin(&_ChainlinkACA.TransactOpts, _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.TransferAdmin(&_ChainlinkACA.TransactOpts, _oracle, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkACA *ChainlinkACATransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkACA *ChainlinkACASession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.TransferOwnership(&_ChainlinkACA.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.TransferOwnership(&_ChainlinkACA.TransactOpts, _to)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_ChainlinkACA *ChainlinkACATransactor) UpdateAvailableFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "updateAvailableFunds")
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_ChainlinkACA *ChainlinkACASession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.UpdateAvailableFunds(&_ChainlinkACA.TransactOpts)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _ChainlinkACA.Contract.UpdateAvailableFunds(&_ChainlinkACA.TransactOpts)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_ChainlinkACA *ChainlinkACATransactor) UpdateFutureRounds(opts *bind.TransactOpts, _paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "updateFutureRounds", _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_ChainlinkACA *ChainlinkACASession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.UpdateFutureRounds(&_ChainlinkACA.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.UpdateFutureRounds(&_ChainlinkACA.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACATransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACASession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.WithdrawFunds(&_ChainlinkACA.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.WithdrawFunds(&_ChainlinkACA.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACATransactor) WithdrawPayment(opts *bind.TransactOpts, _oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.contract.Transact(opts, "withdrawPayment", _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACASession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.WithdrawPayment(&_ChainlinkACA.TransactOpts, _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_ChainlinkACA *ChainlinkACATransactorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkACA.Contract.WithdrawPayment(&_ChainlinkACA.TransactOpts, _oracle, _recipient, _amount)
}

// ChainlinkACAAddedAccessIterator is returned from FilterAddedAccess and is used to iterate over the raw logs and unpacked data for AddedAccess events raised by the ChainlinkACA contract.
type ChainlinkACAAddedAccessIterator struct {
	Event *ChainlinkACAAddedAccess // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAAddedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAAddedAccess)
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
		it.Event = new(ChainlinkACAAddedAccess)
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
func (it *ChainlinkACAAddedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAAddedAccess represents a AddedAccess event raised by the ChainlinkACA contract.
type ChainlinkACAAddedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedAccess is a free log retrieval operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*ChainlinkACAAddedAccessIterator, error) {

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAAddedAccessIterator{contract: _ChainlinkACA.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

// WatchAddedAccess is a free log subscription operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ChainlinkACAAddedAccess) (event.Subscription, error) {

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAAddedAccess)
				if err := _ChainlinkACA.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

// ParseAddedAccess is a log parse operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseAddedAccess(log types.Log) (*ChainlinkACAAddedAccess, error) {
	event := new(ChainlinkACAAddedAccess)
	if err := _ChainlinkACA.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the ChainlinkACA contract.
type ChainlinkACAAnswerUpdatedIterator struct {
	Event *ChainlinkACAAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAAnswerUpdated)
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
		it.Event = new(ChainlinkACAAnswerUpdated)
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
func (it *ChainlinkACAAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAAnswerUpdated represents a AnswerUpdated event raised by the ChainlinkACA contract.
type ChainlinkACAAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ChainlinkACAAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAAnswerUpdatedIterator{contract: _ChainlinkACA.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACAAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAAnswerUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseAnswerUpdated(log types.Log) (*ChainlinkACAAnswerUpdated, error) {
	event := new(ChainlinkACAAnswerUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAAvailableFundsUpdatedIterator is returned from FilterAvailableFundsUpdated and is used to iterate over the raw logs and unpacked data for AvailableFundsUpdated events raised by the ChainlinkACA contract.
type ChainlinkACAAvailableFundsUpdatedIterator struct {
	Event *ChainlinkACAAvailableFundsUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAAvailableFundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAAvailableFundsUpdated)
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
		it.Event = new(ChainlinkACAAvailableFundsUpdated)
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
func (it *ChainlinkACAAvailableFundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAAvailableFundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAAvailableFundsUpdated represents a AvailableFundsUpdated event raised by the ChainlinkACA contract.
type ChainlinkACAAvailableFundsUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAvailableFundsUpdated is a free log retrieval operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterAvailableFundsUpdated(opts *bind.FilterOpts, amount []*big.Int) (*ChainlinkACAAvailableFundsUpdatedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAAvailableFundsUpdatedIterator{contract: _ChainlinkACA.contract, event: "AvailableFundsUpdated", logs: logs, sub: sub}, nil
}

// WatchAvailableFundsUpdated is a free log subscription operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchAvailableFundsUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACAAvailableFundsUpdated, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAAvailableFundsUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
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

// ParseAvailableFundsUpdated is a log parse operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseAvailableFundsUpdated(log types.Log) (*ChainlinkACAAvailableFundsUpdated, error) {
	event := new(ChainlinkACAAvailableFundsUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACACheckAccessDisabledIterator is returned from FilterCheckAccessDisabled and is used to iterate over the raw logs and unpacked data for CheckAccessDisabled events raised by the ChainlinkACA contract.
type ChainlinkACACheckAccessDisabledIterator struct {
	Event *ChainlinkACACheckAccessDisabled // Event containing the contract specifics and raw log

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
func (it *ChainlinkACACheckAccessDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACACheckAccessDisabled)
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
		it.Event = new(ChainlinkACACheckAccessDisabled)
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
func (it *ChainlinkACACheckAccessDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACACheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACACheckAccessDisabled represents a CheckAccessDisabled event raised by the ChainlinkACA contract.
type ChainlinkACACheckAccessDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessDisabled is a free log retrieval operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_ChainlinkACA *ChainlinkACAFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ChainlinkACACheckAccessDisabledIterator, error) {

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &ChainlinkACACheckAccessDisabledIterator{contract: _ChainlinkACA.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessDisabled is a free log subscription operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_ChainlinkACA *ChainlinkACAFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ChainlinkACACheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACACheckAccessDisabled)
				if err := _ChainlinkACA.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

// ParseCheckAccessDisabled is a log parse operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_ChainlinkACA *ChainlinkACAFilterer) ParseCheckAccessDisabled(log types.Log) (*ChainlinkACACheckAccessDisabled, error) {
	event := new(ChainlinkACACheckAccessDisabled)
	if err := _ChainlinkACA.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACACheckAccessEnabledIterator is returned from FilterCheckAccessEnabled and is used to iterate over the raw logs and unpacked data for CheckAccessEnabled events raised by the ChainlinkACA contract.
type ChainlinkACACheckAccessEnabledIterator struct {
	Event *ChainlinkACACheckAccessEnabled // Event containing the contract specifics and raw log

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
func (it *ChainlinkACACheckAccessEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACACheckAccessEnabled)
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
		it.Event = new(ChainlinkACACheckAccessEnabled)
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
func (it *ChainlinkACACheckAccessEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACACheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACACheckAccessEnabled represents a CheckAccessEnabled event raised by the ChainlinkACA contract.
type ChainlinkACACheckAccessEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessEnabled is a free log retrieval operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_ChainlinkACA *ChainlinkACAFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ChainlinkACACheckAccessEnabledIterator, error) {

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &ChainlinkACACheckAccessEnabledIterator{contract: _ChainlinkACA.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessEnabled is a free log subscription operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_ChainlinkACA *ChainlinkACAFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ChainlinkACACheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACACheckAccessEnabled)
				if err := _ChainlinkACA.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

// ParseCheckAccessEnabled is a log parse operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_ChainlinkACA *ChainlinkACAFilterer) ParseCheckAccessEnabled(log types.Log) (*ChainlinkACACheckAccessEnabled, error) {
	event := new(ChainlinkACACheckAccessEnabled)
	if err := _ChainlinkACA.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACANewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the ChainlinkACA contract.
type ChainlinkACANewRoundIterator struct {
	Event *ChainlinkACANewRound // Event containing the contract specifics and raw log

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
func (it *ChainlinkACANewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACANewRound)
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
		it.Event = new(ChainlinkACANewRound)
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
func (it *ChainlinkACANewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACANewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACANewRound represents a NewRound event raised by the ChainlinkACA contract.
type ChainlinkACANewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ChainlinkACANewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACANewRoundIterator{contract: _ChainlinkACA.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ChainlinkACANewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACANewRound)
				if err := _ChainlinkACA.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseNewRound(log types.Log) (*ChainlinkACANewRound, error) {
	event := new(ChainlinkACANewRound)
	if err := _ChainlinkACA.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAOracleAdminUpdateRequestedIterator is returned from FilterOracleAdminUpdateRequested and is used to iterate over the raw logs and unpacked data for OracleAdminUpdateRequested events raised by the ChainlinkACA contract.
type ChainlinkACAOracleAdminUpdateRequestedIterator struct {
	Event *ChainlinkACAOracleAdminUpdateRequested // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAOracleAdminUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAOracleAdminUpdateRequested)
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
		it.Event = new(ChainlinkACAOracleAdminUpdateRequested)
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
func (it *ChainlinkACAOracleAdminUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAOracleAdminUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAOracleAdminUpdateRequested represents a OracleAdminUpdateRequested event raised by the ChainlinkACA contract.
type ChainlinkACAOracleAdminUpdateRequested struct {
	Oracle   common.Address
	Admin    common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdateRequested is a free log retrieval operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterOracleAdminUpdateRequested(opts *bind.FilterOpts, oracle []common.Address) (*ChainlinkACAOracleAdminUpdateRequestedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAOracleAdminUpdateRequestedIterator{contract: _ChainlinkACA.contract, event: "OracleAdminUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdateRequested is a free log subscription operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchOracleAdminUpdateRequested(opts *bind.WatchOpts, sink chan<- *ChainlinkACAOracleAdminUpdateRequested, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAOracleAdminUpdateRequested)
				if err := _ChainlinkACA.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
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

// ParseOracleAdminUpdateRequested is a log parse operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseOracleAdminUpdateRequested(log types.Log) (*ChainlinkACAOracleAdminUpdateRequested, error) {
	event := new(ChainlinkACAOracleAdminUpdateRequested)
	if err := _ChainlinkACA.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAOracleAdminUpdatedIterator is returned from FilterOracleAdminUpdated and is used to iterate over the raw logs and unpacked data for OracleAdminUpdated events raised by the ChainlinkACA contract.
type ChainlinkACAOracleAdminUpdatedIterator struct {
	Event *ChainlinkACAOracleAdminUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAOracleAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAOracleAdminUpdated)
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
		it.Event = new(ChainlinkACAOracleAdminUpdated)
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
func (it *ChainlinkACAOracleAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAOracleAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAOracleAdminUpdated represents a OracleAdminUpdated event raised by the ChainlinkACA contract.
type ChainlinkACAOracleAdminUpdated struct {
	Oracle   common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdated is a free log retrieval operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterOracleAdminUpdated(opts *bind.FilterOpts, oracle []common.Address, newAdmin []common.Address) (*ChainlinkACAOracleAdminUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAOracleAdminUpdatedIterator{contract: _ChainlinkACA.contract, event: "OracleAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdated is a free log subscription operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchOracleAdminUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACAOracleAdminUpdated, oracle []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAOracleAdminUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
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

// ParseOracleAdminUpdated is a log parse operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseOracleAdminUpdated(log types.Log) (*ChainlinkACAOracleAdminUpdated, error) {
	event := new(ChainlinkACAOracleAdminUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAOraclePermissionsUpdatedIterator is returned from FilterOraclePermissionsUpdated and is used to iterate over the raw logs and unpacked data for OraclePermissionsUpdated events raised by the ChainlinkACA contract.
type ChainlinkACAOraclePermissionsUpdatedIterator struct {
	Event *ChainlinkACAOraclePermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAOraclePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAOraclePermissionsUpdated)
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
		it.Event = new(ChainlinkACAOraclePermissionsUpdated)
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
func (it *ChainlinkACAOraclePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAOraclePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAOraclePermissionsUpdated represents a OraclePermissionsUpdated event raised by the ChainlinkACA contract.
type ChainlinkACAOraclePermissionsUpdated struct {
	Oracle      common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePermissionsUpdated is a free log retrieval operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterOraclePermissionsUpdated(opts *bind.FilterOpts, oracle []common.Address, whitelisted []bool) (*ChainlinkACAOraclePermissionsUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAOraclePermissionsUpdatedIterator{contract: _ChainlinkACA.contract, event: "OraclePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchOraclePermissionsUpdated is a free log subscription operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchOraclePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACAOraclePermissionsUpdated, oracle []common.Address, whitelisted []bool) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAOraclePermissionsUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
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

// ParseOraclePermissionsUpdated is a log parse operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseOraclePermissionsUpdated(log types.Log) (*ChainlinkACAOraclePermissionsUpdated, error) {
	event := new(ChainlinkACAOraclePermissionsUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the ChainlinkACA contract.
type ChainlinkACAOwnershipTransferRequestedIterator struct {
	Event *ChainlinkACAOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAOwnershipTransferRequested)
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
		it.Event = new(ChainlinkACAOwnershipTransferRequested)
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
func (it *ChainlinkACAOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the ChainlinkACA contract.
type ChainlinkACAOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkACAOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAOwnershipTransferRequestedIterator{contract: _ChainlinkACA.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ChainlinkACAOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAOwnershipTransferRequested)
				if err := _ChainlinkACA.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseOwnershipTransferRequested(log types.Log) (*ChainlinkACAOwnershipTransferRequested, error) {
	event := new(ChainlinkACAOwnershipTransferRequested)
	if err := _ChainlinkACA.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkACA contract.
type ChainlinkACAOwnershipTransferredIterator struct {
	Event *ChainlinkACAOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAOwnershipTransferred)
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
		it.Event = new(ChainlinkACAOwnershipTransferred)
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
func (it *ChainlinkACAOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkACA contract.
type ChainlinkACAOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkACAOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAOwnershipTransferredIterator{contract: _ChainlinkACA.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkACAOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAOwnershipTransferred)
				if err := _ChainlinkACA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkACAOwnershipTransferred, error) {
	event := new(ChainlinkACAOwnershipTransferred)
	if err := _ChainlinkACA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACARemovedAccessIterator is returned from FilterRemovedAccess and is used to iterate over the raw logs and unpacked data for RemovedAccess events raised by the ChainlinkACA contract.
type ChainlinkACARemovedAccessIterator struct {
	Event *ChainlinkACARemovedAccess // Event containing the contract specifics and raw log

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
func (it *ChainlinkACARemovedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACARemovedAccess)
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
		it.Event = new(ChainlinkACARemovedAccess)
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
func (it *ChainlinkACARemovedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACARemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACARemovedAccess represents a RemovedAccess event raised by the ChainlinkACA contract.
type ChainlinkACARemovedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedAccess is a free log retrieval operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*ChainlinkACARemovedAccessIterator, error) {

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &ChainlinkACARemovedAccessIterator{contract: _ChainlinkACA.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

// WatchRemovedAccess is a free log subscription operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ChainlinkACARemovedAccess) (event.Subscription, error) {

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACARemovedAccess)
				if err := _ChainlinkACA.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

// ParseRemovedAccess is a log parse operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseRemovedAccess(log types.Log) (*ChainlinkACARemovedAccess, error) {
	event := new(ChainlinkACARemovedAccess)
	if err := _ChainlinkACA.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACARequesterPermissionsSetIterator is returned from FilterRequesterPermissionsSet and is used to iterate over the raw logs and unpacked data for RequesterPermissionsSet events raised by the ChainlinkACA contract.
type ChainlinkACARequesterPermissionsSetIterator struct {
	Event *ChainlinkACARequesterPermissionsSet // Event containing the contract specifics and raw log

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
func (it *ChainlinkACARequesterPermissionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACARequesterPermissionsSet)
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
		it.Event = new(ChainlinkACARequesterPermissionsSet)
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
func (it *ChainlinkACARequesterPermissionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACARequesterPermissionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACARequesterPermissionsSet represents a RequesterPermissionsSet event raised by the ChainlinkACA contract.
type ChainlinkACARequesterPermissionsSet struct {
	Requester  common.Address
	Authorized bool
	Delay      uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequesterPermissionsSet is a free log retrieval operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterRequesterPermissionsSet(opts *bind.FilterOpts, requester []common.Address) (*ChainlinkACARequesterPermissionsSetIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACARequesterPermissionsSetIterator{contract: _ChainlinkACA.contract, event: "RequesterPermissionsSet", logs: logs, sub: sub}, nil
}

// WatchRequesterPermissionsSet is a free log subscription operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchRequesterPermissionsSet(opts *bind.WatchOpts, sink chan<- *ChainlinkACARequesterPermissionsSet, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACARequesterPermissionsSet)
				if err := _ChainlinkACA.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
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

// ParseRequesterPermissionsSet is a log parse operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseRequesterPermissionsSet(log types.Log) (*ChainlinkACARequesterPermissionsSet, error) {
	event := new(ChainlinkACARequesterPermissionsSet)
	if err := _ChainlinkACA.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACARoundDetailsUpdatedIterator is returned from FilterRoundDetailsUpdated and is used to iterate over the raw logs and unpacked data for RoundDetailsUpdated events raised by the ChainlinkACA contract.
type ChainlinkACARoundDetailsUpdatedIterator struct {
	Event *ChainlinkACARoundDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACARoundDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACARoundDetailsUpdated)
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
		it.Event = new(ChainlinkACARoundDetailsUpdated)
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
func (it *ChainlinkACARoundDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACARoundDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACARoundDetailsUpdated represents a RoundDetailsUpdated event raised by the ChainlinkACA contract.
type ChainlinkACARoundDetailsUpdated struct {
	PaymentAmount      *big.Int
	MinSubmissionCount uint32
	MaxSubmissionCount uint32
	RestartDelay       uint32
	Timeout            uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRoundDetailsUpdated is a free log retrieval operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterRoundDetailsUpdated(opts *bind.FilterOpts, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (*ChainlinkACARoundDetailsUpdatedIterator, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACARoundDetailsUpdatedIterator{contract: _ChainlinkACA.contract, event: "RoundDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundDetailsUpdated is a free log subscription operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchRoundDetailsUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACARoundDetailsUpdated, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (event.Subscription, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACARoundDetailsUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
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

// ParseRoundDetailsUpdated is a log parse operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseRoundDetailsUpdated(log types.Log) (*ChainlinkACARoundDetailsUpdated, error) {
	event := new(ChainlinkACARoundDetailsUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACASubmissionReceivedIterator is returned from FilterSubmissionReceived and is used to iterate over the raw logs and unpacked data for SubmissionReceived events raised by the ChainlinkACA contract.
type ChainlinkACASubmissionReceivedIterator struct {
	Event *ChainlinkACASubmissionReceived // Event containing the contract specifics and raw log

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
func (it *ChainlinkACASubmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACASubmissionReceived)
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
		it.Event = new(ChainlinkACASubmissionReceived)
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
func (it *ChainlinkACASubmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACASubmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACASubmissionReceived represents a SubmissionReceived event raised by the ChainlinkACA contract.
type ChainlinkACASubmissionReceived struct {
	Submission *big.Int
	Round      uint32
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmissionReceived is a free log retrieval operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterSubmissionReceived(opts *bind.FilterOpts, submission []*big.Int, round []uint32, oracle []common.Address) (*ChainlinkACASubmissionReceivedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACASubmissionReceivedIterator{contract: _ChainlinkACA.contract, event: "SubmissionReceived", logs: logs, sub: sub}, nil
}

// WatchSubmissionReceived is a free log subscription operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchSubmissionReceived(opts *bind.WatchOpts, sink chan<- *ChainlinkACASubmissionReceived, submission []*big.Int, round []uint32, oracle []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACASubmissionReceived)
				if err := _ChainlinkACA.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
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

// ParseSubmissionReceived is a log parse operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseSubmissionReceived(log types.Log) (*ChainlinkACASubmissionReceived, error) {
	event := new(ChainlinkACASubmissionReceived)
	if err := _ChainlinkACA.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkACAValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the ChainlinkACA contract.
type ChainlinkACAValidatorUpdatedIterator struct {
	Event *ChainlinkACAValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkACAValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkACAValidatorUpdated)
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
		it.Event = new(ChainlinkACAValidatorUpdated)
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
func (it *ChainlinkACAValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkACAValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkACAValidatorUpdated represents a ValidatorUpdated event raised by the ChainlinkACA contract.
type ChainlinkACAValidatorUpdated struct {
	Previous common.Address
	Current  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_ChainlinkACA *ChainlinkACAFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*ChainlinkACAValidatorUpdatedIterator, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ChainlinkACA.contract.FilterLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkACAValidatorUpdatedIterator{contract: _ChainlinkACA.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_ChainlinkACA *ChainlinkACAFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkACAValidatorUpdated, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ChainlinkACA.contract.WatchLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkACAValidatorUpdated)
				if err := _ChainlinkACA.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
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

// ParseValidatorUpdated is a log parse operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_ChainlinkACA *ChainlinkACAFilterer) ParseValidatorUpdated(log types.Log) (*ChainlinkACAValidatorUpdated, error) {
	event := new(ChainlinkACAValidatorUpdated)
	if err := _ChainlinkACA.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
