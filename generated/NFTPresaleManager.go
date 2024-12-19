// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftsale

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

// NFTPresaleManagerReferrerSignatureHashInfo is an auto generated low-level Go binding around an user-defined struct.
type NFTPresaleManagerReferrerSignatureHashInfo struct {
	Referrer  common.Address
	CreatedAt *big.Int
}

// NFTPresaleManagerMetaData contains all meta data concerning the NFTPresaleManager contract.
var NFTPresaleManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxNFTPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxNFTPerWallet\",\"type\":\"uint256\"}],\"name\":\"MaxNFTPerWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"PaymentTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReferralCommissionPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"SaleRoundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REFERRAL_COMMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"acceptedPaymentTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addToWhitelistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"blackListSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"buyNFTWithReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_whitelistOnly\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_maxNFTPerWalletInRound\",\"type\":\"uint256\"}],\"name\":\"createSaleRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getRoundInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"whitelistOnly\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalMintedInRound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNFTPerWalletInRound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"walletMinted\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structNFTPresaleManager.ReferrerSignatureHashInfo\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"hashSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNFTPerWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"saleRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"whitelistOnly\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNFTPerWalletInRound\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRoundEnded\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNFTPerWallet\",\"type\":\"uint256\"}],\"name\":\"setMaxNFTPerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNFTPerWallet\",\"type\":\"uint256\"}],\"name\":\"setMaxNFTPerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_accepted\",\"type\":\"bool\"}],\"name\":\"setPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setRoundEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTPresaleManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTPresaleManagerMetaData.ABI instead.
var NFTPresaleManagerABI = NFTPresaleManagerMetaData.ABI

// NFTPresaleManager is an auto generated Go binding around an Ethereum contract.
type NFTPresaleManager struct {
	NFTPresaleManagerCaller     // Read-only binding to the contract
	NFTPresaleManagerTransactor // Write-only binding to the contract
	NFTPresaleManagerFilterer   // Log filterer for contract events
}

// NFTPresaleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTPresaleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTPresaleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTPresaleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTPresaleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTPresaleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTPresaleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTPresaleManagerSession struct {
	Contract     *NFTPresaleManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NFTPresaleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTPresaleManagerCallerSession struct {
	Contract *NFTPresaleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NFTPresaleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTPresaleManagerTransactorSession struct {
	Contract     *NFTPresaleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NFTPresaleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTPresaleManagerRaw struct {
	Contract *NFTPresaleManager // Generic contract binding to access the raw methods on
}

// NFTPresaleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTPresaleManagerCallerRaw struct {
	Contract *NFTPresaleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NFTPresaleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTPresaleManagerTransactorRaw struct {
	Contract *NFTPresaleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTPresaleManager creates a new instance of NFTPresaleManager, bound to a specific deployed contract.
func NewNFTPresaleManager(address common.Address, backend bind.ContractBackend) (*NFTPresaleManager, error) {
	contract, err := bindNFTPresaleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManager{NFTPresaleManagerCaller: NFTPresaleManagerCaller{contract: contract}, NFTPresaleManagerTransactor: NFTPresaleManagerTransactor{contract: contract}, NFTPresaleManagerFilterer: NFTPresaleManagerFilterer{contract: contract}}, nil
}

// NewNFTPresaleManagerCaller creates a new read-only instance of NFTPresaleManager, bound to a specific deployed contract.
func NewNFTPresaleManagerCaller(address common.Address, caller bind.ContractCaller) (*NFTPresaleManagerCaller, error) {
	contract, err := bindNFTPresaleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerCaller{contract: contract}, nil
}

// NewNFTPresaleManagerTransactor creates a new write-only instance of NFTPresaleManager, bound to a specific deployed contract.
func NewNFTPresaleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTPresaleManagerTransactor, error) {
	contract, err := bindNFTPresaleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerTransactor{contract: contract}, nil
}

// NewNFTPresaleManagerFilterer creates a new log filterer instance of NFTPresaleManager, bound to a specific deployed contract.
func NewNFTPresaleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTPresaleManagerFilterer, error) {
	contract, err := bindNFTPresaleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerFilterer{contract: contract}, nil
}

// bindNFTPresaleManager binds a generic wrapper to an already deployed contract.
func bindNFTPresaleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTPresaleManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTPresaleManager *NFTPresaleManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTPresaleManager.Contract.NFTPresaleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTPresaleManager *NFTPresaleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.NFTPresaleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTPresaleManager *NFTPresaleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.NFTPresaleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTPresaleManager *NFTPresaleManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTPresaleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTPresaleManager *NFTPresaleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTPresaleManager *NFTPresaleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.contract.Transact(opts, method, params...)
}

// REFERRALCOMMISSION is a free data retrieval call binding the contract method 0xa365f2a5.
//
// Solidity: function REFERRAL_COMMISSION() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCaller) REFERRALCOMMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "REFERRAL_COMMISSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFERRALCOMMISSION is a free data retrieval call binding the contract method 0xa365f2a5.
//
// Solidity: function REFERRAL_COMMISSION() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerSession) REFERRALCOMMISSION() (*big.Int, error) {
	return _NFTPresaleManager.Contract.REFERRALCOMMISSION(&_NFTPresaleManager.CallOpts)
}

// REFERRALCOMMISSION is a free data retrieval call binding the contract method 0xa365f2a5.
//
// Solidity: function REFERRAL_COMMISSION() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) REFERRALCOMMISSION() (*big.Int, error) {
	return _NFTPresaleManager.Contract.REFERRALCOMMISSION(&_NFTPresaleManager.CallOpts)
}

// AcceptedPaymentTokens is a free data retrieval call binding the contract method 0x484bb244.
//
// Solidity: function acceptedPaymentTokens(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCaller) AcceptedPaymentTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "acceptedPaymentTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptedPaymentTokens is a free data retrieval call binding the contract method 0x484bb244.
//
// Solidity: function acceptedPaymentTokens(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerSession) AcceptedPaymentTokens(arg0 common.Address) (bool, error) {
	return _NFTPresaleManager.Contract.AcceptedPaymentTokens(&_NFTPresaleManager.CallOpts, arg0)
}

// AcceptedPaymentTokens is a free data retrieval call binding the contract method 0x484bb244.
//
// Solidity: function acceptedPaymentTokens(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) AcceptedPaymentTokens(arg0 common.Address) (bool, error) {
	return _NFTPresaleManager.Contract.AcceptedPaymentTokens(&_NFTPresaleManager.CallOpts, arg0)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "adminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerSession) AdminAddress() (common.Address, error) {
	return _NFTPresaleManager.Contract.AdminAddress(&_NFTPresaleManager.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) AdminAddress() (common.Address, error) {
	return _NFTPresaleManager.Contract.AdminAddress(&_NFTPresaleManager.CallOpts)
}

// BlackListSignature is a free data retrieval call binding the contract method 0x12f81038.
//
// Solidity: function blackListSignature(bytes ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCaller) BlackListSignature(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "blackListSignature", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackListSignature is a free data retrieval call binding the contract method 0x12f81038.
//
// Solidity: function blackListSignature(bytes ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerSession) BlackListSignature(arg0 []byte) (bool, error) {
	return _NFTPresaleManager.Contract.BlackListSignature(&_NFTPresaleManager.CallOpts, arg0)
}

// BlackListSignature is a free data retrieval call binding the contract method 0x12f81038.
//
// Solidity: function blackListSignature(bytes ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) BlackListSignature(arg0 []byte) (bool, error) {
	return _NFTPresaleManager.Contract.BlackListSignature(&_NFTPresaleManager.CallOpts, arg0)
}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCaller) CurrentRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "currentRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerSession) CurrentRoundId() (*big.Int, error) {
	return _NFTPresaleManager.Contract.CurrentRoundId(&_NFTPresaleManager.CallOpts)
}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) CurrentRoundId() (*big.Int, error) {
	return _NFTPresaleManager.Contract.CurrentRoundId(&_NFTPresaleManager.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NFTPresaleManager *NFTPresaleManagerCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NFTPresaleManager *NFTPresaleManagerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NFTPresaleManager.Contract.Eip712Domain(&_NFTPresaleManager.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NFTPresaleManager.Contract.Eip712Domain(&_NFTPresaleManager.CallOpts)
}

// GetRoundInfo is a free data retrieval call binding the contract method 0x88c3ffb0.
//
// Solidity: function getRoundInfo(uint256 _roundId) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMintedInRound, uint256 maxNFTPerWalletInRound, uint256 walletMinted)
func (_NFTPresaleManager *NFTPresaleManagerCaller) GetRoundInfo(opts *bind.CallOpts, _roundId *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMintedInRound     *big.Int
	MaxNFTPerWalletInRound *big.Int
	WalletMinted           *big.Int
}, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "getRoundInfo", _roundId)

	outstruct := new(struct {
		Id                     *big.Int
		StartTime              *big.Int
		EndTime                *big.Int
		MaxSupply              *big.Int
		Price                  *big.Int
		WhitelistOnly          bool
		TotalMintedInRound     *big.Int
		MaxNFTPerWalletInRound *big.Int
		WalletMinted           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.WhitelistOnly = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TotalMintedInRound = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MaxNFTPerWalletInRound = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.WalletMinted = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundInfo is a free data retrieval call binding the contract method 0x88c3ffb0.
//
// Solidity: function getRoundInfo(uint256 _roundId) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMintedInRound, uint256 maxNFTPerWalletInRound, uint256 walletMinted)
func (_NFTPresaleManager *NFTPresaleManagerSession) GetRoundInfo(_roundId *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMintedInRound     *big.Int
	MaxNFTPerWalletInRound *big.Int
	WalletMinted           *big.Int
}, error) {
	return _NFTPresaleManager.Contract.GetRoundInfo(&_NFTPresaleManager.CallOpts, _roundId)
}

// GetRoundInfo is a free data retrieval call binding the contract method 0x88c3ffb0.
//
// Solidity: function getRoundInfo(uint256 _roundId) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMintedInRound, uint256 maxNFTPerWalletInRound, uint256 walletMinted)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) GetRoundInfo(_roundId *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMintedInRound     *big.Int
	MaxNFTPerWalletInRound *big.Int
	WalletMinted           *big.Int
}, error) {
	return _NFTPresaleManager.Contract.GetRoundInfo(&_NFTPresaleManager.CallOpts, _roundId)
}

// HashSignature is a free data retrieval call binding the contract method 0x9e115f75.
//
// Solidity: function hashSignature(address _userAddr, (address,uint256) data) view returns(bytes32)
func (_NFTPresaleManager *NFTPresaleManagerCaller) HashSignature(opts *bind.CallOpts, _userAddr common.Address, data NFTPresaleManagerReferrerSignatureHashInfo) ([32]byte, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "hashSignature", _userAddr, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashSignature is a free data retrieval call binding the contract method 0x9e115f75.
//
// Solidity: function hashSignature(address _userAddr, (address,uint256) data) view returns(bytes32)
func (_NFTPresaleManager *NFTPresaleManagerSession) HashSignature(_userAddr common.Address, data NFTPresaleManagerReferrerSignatureHashInfo) ([32]byte, error) {
	return _NFTPresaleManager.Contract.HashSignature(&_NFTPresaleManager.CallOpts, _userAddr, data)
}

// HashSignature is a free data retrieval call binding the contract method 0x9e115f75.
//
// Solidity: function hashSignature(address _userAddr, (address,uint256) data) view returns(bytes32)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) HashSignature(_userAddr common.Address, data NFTPresaleManagerReferrerSignatureHashInfo) ([32]byte, error) {
	return _NFTPresaleManager.Contract.HashSignature(&_NFTPresaleManager.CallOpts, _userAddr, data)
}

// MaxNFTPerWallet is a free data retrieval call binding the contract method 0xc7c5a8c0.
//
// Solidity: function maxNFTPerWallet() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCaller) MaxNFTPerWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "maxNFTPerWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNFTPerWallet is a free data retrieval call binding the contract method 0xc7c5a8c0.
//
// Solidity: function maxNFTPerWallet() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerSession) MaxNFTPerWallet() (*big.Int, error) {
	return _NFTPresaleManager.Contract.MaxNFTPerWallet(&_NFTPresaleManager.CallOpts)
}

// MaxNFTPerWallet is a free data retrieval call binding the contract method 0xc7c5a8c0.
//
// Solidity: function maxNFTPerWallet() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) MaxNFTPerWallet() (*big.Int, error) {
	return _NFTPresaleManager.Contract.MaxNFTPerWallet(&_NFTPresaleManager.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerSession) NftContract() (common.Address, error) {
	return _NFTPresaleManager.Contract.NftContract(&_NFTPresaleManager.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) NftContract() (common.Address, error) {
	return _NFTPresaleManager.Contract.NftContract(&_NFTPresaleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerSession) Owner() (common.Address, error) {
	return _NFTPresaleManager.Contract.Owner(&_NFTPresaleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) Owner() (common.Address, error) {
	return _NFTPresaleManager.Contract.Owner(&_NFTPresaleManager.CallOpts)
}

// SaleRounds is a free data retrieval call binding the contract method 0x45199ce1.
//
// Solidity: function saleRounds(uint256 ) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMinted, uint256 maxNFTPerWalletInRound, bool isRoundEnded)
func (_NFTPresaleManager *NFTPresaleManagerCaller) SaleRounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMinted            *big.Int
	MaxNFTPerWalletInRound *big.Int
	IsRoundEnded           bool
}, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "saleRounds", arg0)

	outstruct := new(struct {
		Id                     *big.Int
		StartTime              *big.Int
		EndTime                *big.Int
		MaxSupply              *big.Int
		Price                  *big.Int
		WhitelistOnly          bool
		TotalMinted            *big.Int
		MaxNFTPerWalletInRound *big.Int
		IsRoundEnded           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.WhitelistOnly = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TotalMinted = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MaxNFTPerWalletInRound = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsRoundEnded = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// SaleRounds is a free data retrieval call binding the contract method 0x45199ce1.
//
// Solidity: function saleRounds(uint256 ) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMinted, uint256 maxNFTPerWalletInRound, bool isRoundEnded)
func (_NFTPresaleManager *NFTPresaleManagerSession) SaleRounds(arg0 *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMinted            *big.Int
	MaxNFTPerWalletInRound *big.Int
	IsRoundEnded           bool
}, error) {
	return _NFTPresaleManager.Contract.SaleRounds(&_NFTPresaleManager.CallOpts, arg0)
}

// SaleRounds is a free data retrieval call binding the contract method 0x45199ce1.
//
// Solidity: function saleRounds(uint256 ) view returns(uint256 id, uint256 startTime, uint256 endTime, uint256 maxSupply, uint256 price, bool whitelistOnly, uint256 totalMinted, uint256 maxNFTPerWalletInRound, bool isRoundEnded)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) SaleRounds(arg0 *big.Int) (struct {
	Id                     *big.Int
	StartTime              *big.Int
	EndTime                *big.Int
	MaxSupply              *big.Int
	Price                  *big.Int
	WhitelistOnly          bool
	TotalMinted            *big.Int
	MaxNFTPerWalletInRound *big.Int
	IsRoundEnded           bool
}, error) {
	return _NFTPresaleManager.Contract.SaleRounds(&_NFTPresaleManager.CallOpts, arg0)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerSession) TotalMinted() (*big.Int, error) {
	return _NFTPresaleManager.Contract.TotalMinted(&_NFTPresaleManager.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) TotalMinted() (*big.Int, error) {
	return _NFTPresaleManager.Contract.TotalMinted(&_NFTPresaleManager.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NFTPresaleManager.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _NFTPresaleManager.Contract.Whitelisted(&_NFTPresaleManager.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_NFTPresaleManager *NFTPresaleManagerCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _NFTPresaleManager.Contract.Whitelisted(&_NFTPresaleManager.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) AddToWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "addToWhitelist", account)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) AddToWhitelist(account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.AddToWhitelist(&_NFTPresaleManager.TransactOpts, account)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) AddToWhitelist(account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.AddToWhitelist(&_NFTPresaleManager.TransactOpts, account)
}

// AddToWhitelistBatch is a paid mutator transaction binding the contract method 0xa6177139.
//
// Solidity: function addToWhitelistBatch(address[] accounts) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) AddToWhitelistBatch(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "addToWhitelistBatch", accounts)
}

// AddToWhitelistBatch is a paid mutator transaction binding the contract method 0xa6177139.
//
// Solidity: function addToWhitelistBatch(address[] accounts) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) AddToWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.AddToWhitelistBatch(&_NFTPresaleManager.TransactOpts, accounts)
}

// AddToWhitelistBatch is a paid mutator transaction binding the contract method 0xa6177139.
//
// Solidity: function addToWhitelistBatch(address[] accounts) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) AddToWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.AddToWhitelistBatch(&_NFTPresaleManager.TransactOpts, accounts)
}

// BuyNFTWithReferral is a paid mutator transaction binding the contract method 0x9635c40a.
//
// Solidity: function buyNFTWithReferral(uint256 _roundId, address _paymentToken, uint256 _quantity, bytes _signature, address _referrer, uint256 createdAt) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) BuyNFTWithReferral(opts *bind.TransactOpts, _roundId *big.Int, _paymentToken common.Address, _quantity *big.Int, _signature []byte, _referrer common.Address, createdAt *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "buyNFTWithReferral", _roundId, _paymentToken, _quantity, _signature, _referrer, createdAt)
}

// BuyNFTWithReferral is a paid mutator transaction binding the contract method 0x9635c40a.
//
// Solidity: function buyNFTWithReferral(uint256 _roundId, address _paymentToken, uint256 _quantity, bytes _signature, address _referrer, uint256 createdAt) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) BuyNFTWithReferral(_roundId *big.Int, _paymentToken common.Address, _quantity *big.Int, _signature []byte, _referrer common.Address, createdAt *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.BuyNFTWithReferral(&_NFTPresaleManager.TransactOpts, _roundId, _paymentToken, _quantity, _signature, _referrer, createdAt)
}

// BuyNFTWithReferral is a paid mutator transaction binding the contract method 0x9635c40a.
//
// Solidity: function buyNFTWithReferral(uint256 _roundId, address _paymentToken, uint256 _quantity, bytes _signature, address _referrer, uint256 createdAt) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) BuyNFTWithReferral(_roundId *big.Int, _paymentToken common.Address, _quantity *big.Int, _signature []byte, _referrer common.Address, createdAt *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.BuyNFTWithReferral(&_NFTPresaleManager.TransactOpts, _roundId, _paymentToken, _quantity, _signature, _referrer, createdAt)
}

// CreateSaleRound is a paid mutator transaction binding the contract method 0xe6b6a7c8.
//
// Solidity: function createSaleRound(uint256 _startTime, uint256 _endTime, uint256 _maxSupply, uint256 _price, bool _whitelistOnly, uint256 _maxNFTPerWalletInRound) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) CreateSaleRound(opts *bind.TransactOpts, _startTime *big.Int, _endTime *big.Int, _maxSupply *big.Int, _price *big.Int, _whitelistOnly bool, _maxNFTPerWalletInRound *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "createSaleRound", _startTime, _endTime, _maxSupply, _price, _whitelistOnly, _maxNFTPerWalletInRound)
}

// CreateSaleRound is a paid mutator transaction binding the contract method 0xe6b6a7c8.
//
// Solidity: function createSaleRound(uint256 _startTime, uint256 _endTime, uint256 _maxSupply, uint256 _price, bool _whitelistOnly, uint256 _maxNFTPerWalletInRound) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) CreateSaleRound(_startTime *big.Int, _endTime *big.Int, _maxSupply *big.Int, _price *big.Int, _whitelistOnly bool, _maxNFTPerWalletInRound *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.CreateSaleRound(&_NFTPresaleManager.TransactOpts, _startTime, _endTime, _maxSupply, _price, _whitelistOnly, _maxNFTPerWalletInRound)
}

// CreateSaleRound is a paid mutator transaction binding the contract method 0xe6b6a7c8.
//
// Solidity: function createSaleRound(uint256 _startTime, uint256 _endTime, uint256 _maxSupply, uint256 _price, bool _whitelistOnly, uint256 _maxNFTPerWalletInRound) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) CreateSaleRound(_startTime *big.Int, _endTime *big.Int, _maxSupply *big.Int, _price *big.Int, _whitelistOnly bool, _maxNFTPerWalletInRound *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.CreateSaleRound(&_NFTPresaleManager.TransactOpts, _startTime, _endTime, _maxSupply, _price, _whitelistOnly, _maxNFTPerWalletInRound)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "removeFromWhitelist", account)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) RemoveFromWhitelist(account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.RemoveFromWhitelist(&_NFTPresaleManager.TransactOpts, account)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) RemoveFromWhitelist(account common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.RemoveFromWhitelist(&_NFTPresaleManager.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.RenounceOwnership(&_NFTPresaleManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.RenounceOwnership(&_NFTPresaleManager.TransactOpts)
}

// SetMaxNFTPerWallet is a paid mutator transaction binding the contract method 0xb5a37260.
//
// Solidity: function setMaxNFTPerWallet(uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) SetMaxNFTPerWallet(opts *bind.TransactOpts, _maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "setMaxNFTPerWallet", _maxNFTPerWallet)
}

// SetMaxNFTPerWallet is a paid mutator transaction binding the contract method 0xb5a37260.
//
// Solidity: function setMaxNFTPerWallet(uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) SetMaxNFTPerWallet(_maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetMaxNFTPerWallet(&_NFTPresaleManager.TransactOpts, _maxNFTPerWallet)
}

// SetMaxNFTPerWallet is a paid mutator transaction binding the contract method 0xb5a37260.
//
// Solidity: function setMaxNFTPerWallet(uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) SetMaxNFTPerWallet(_maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetMaxNFTPerWallet(&_NFTPresaleManager.TransactOpts, _maxNFTPerWallet)
}

// SetMaxNFTPerWallet0 is a paid mutator transaction binding the contract method 0xf497994b.
//
// Solidity: function setMaxNFTPerWallet(uint256 _roundId, uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) SetMaxNFTPerWallet0(opts *bind.TransactOpts, _roundId *big.Int, _maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "setMaxNFTPerWallet0", _roundId, _maxNFTPerWallet)
}

// SetMaxNFTPerWallet0 is a paid mutator transaction binding the contract method 0xf497994b.
//
// Solidity: function setMaxNFTPerWallet(uint256 _roundId, uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) SetMaxNFTPerWallet0(_roundId *big.Int, _maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetMaxNFTPerWallet0(&_NFTPresaleManager.TransactOpts, _roundId, _maxNFTPerWallet)
}

// SetMaxNFTPerWallet0 is a paid mutator transaction binding the contract method 0xf497994b.
//
// Solidity: function setMaxNFTPerWallet(uint256 _roundId, uint256 _maxNFTPerWallet) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) SetMaxNFTPerWallet0(_roundId *big.Int, _maxNFTPerWallet *big.Int) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetMaxNFTPerWallet0(&_NFTPresaleManager.TransactOpts, _roundId, _maxNFTPerWallet)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x430884cf.
//
// Solidity: function setPaymentToken(address _token, bool _accepted) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) SetPaymentToken(opts *bind.TransactOpts, _token common.Address, _accepted bool) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "setPaymentToken", _token, _accepted)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x430884cf.
//
// Solidity: function setPaymentToken(address _token, bool _accepted) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) SetPaymentToken(_token common.Address, _accepted bool) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetPaymentToken(&_NFTPresaleManager.TransactOpts, _token, _accepted)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x430884cf.
//
// Solidity: function setPaymentToken(address _token, bool _accepted) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) SetPaymentToken(_token common.Address, _accepted bool) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetPaymentToken(&_NFTPresaleManager.TransactOpts, _token, _accepted)
}

// SetRoundEnd is a paid mutator transaction binding the contract method 0xd2e068fd.
//
// Solidity: function setRoundEnd(uint256 _roundId, bool status) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) SetRoundEnd(opts *bind.TransactOpts, _roundId *big.Int, status bool) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "setRoundEnd", _roundId, status)
}

// SetRoundEnd is a paid mutator transaction binding the contract method 0xd2e068fd.
//
// Solidity: function setRoundEnd(uint256 _roundId, bool status) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) SetRoundEnd(_roundId *big.Int, status bool) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetRoundEnd(&_NFTPresaleManager.TransactOpts, _roundId, status)
}

// SetRoundEnd is a paid mutator transaction binding the contract method 0xd2e068fd.
//
// Solidity: function setRoundEnd(uint256 _roundId, bool status) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) SetRoundEnd(_roundId *big.Int, status bool) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.SetRoundEnd(&_NFTPresaleManager.TransactOpts, _roundId, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.TransferOwnership(&_NFTPresaleManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.TransferOwnership(&_NFTPresaleManager.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _token) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactor) WithdrawFunds(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.contract.Transact(opts, "withdrawFunds", _token)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _token) returns()
func (_NFTPresaleManager *NFTPresaleManagerSession) WithdrawFunds(_token common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.WithdrawFunds(&_NFTPresaleManager.TransactOpts, _token)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _token) returns()
func (_NFTPresaleManager *NFTPresaleManagerTransactorSession) WithdrawFunds(_token common.Address) (*types.Transaction, error) {
	return _NFTPresaleManager.Contract.WithdrawFunds(&_NFTPresaleManager.TransactOpts, _token)
}

// NFTPresaleManagerEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the NFTPresaleManager contract.
type NFTPresaleManagerEIP712DomainChangedIterator struct {
	Event *NFTPresaleManagerEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerEIP712DomainChanged)
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
		it.Event = new(NFTPresaleManagerEIP712DomainChanged)
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
func (it *NFTPresaleManagerEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerEIP712DomainChanged represents a EIP712DomainChanged event raised by the NFTPresaleManager contract.
type NFTPresaleManagerEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*NFTPresaleManagerEIP712DomainChangedIterator, error) {

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerEIP712DomainChangedIterator{contract: _NFTPresaleManager.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerEIP712DomainChanged)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseEIP712DomainChanged(log types.Log) (*NFTPresaleManagerEIP712DomainChanged, error) {
	event := new(NFTPresaleManagerEIP712DomainChanged)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerMaxNFTPerWalletUpdatedIterator is returned from FilterMaxNFTPerWalletUpdated and is used to iterate over the raw logs and unpacked data for MaxNFTPerWalletUpdated events raised by the NFTPresaleManager contract.
type NFTPresaleManagerMaxNFTPerWalletUpdatedIterator struct {
	Event *NFTPresaleManagerMaxNFTPerWalletUpdated // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerMaxNFTPerWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerMaxNFTPerWalletUpdated)
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
		it.Event = new(NFTPresaleManagerMaxNFTPerWalletUpdated)
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
func (it *NFTPresaleManagerMaxNFTPerWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerMaxNFTPerWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerMaxNFTPerWalletUpdated represents a MaxNFTPerWalletUpdated event raised by the NFTPresaleManager contract.
type NFTPresaleManagerMaxNFTPerWalletUpdated struct {
	RoundId         *big.Int
	MaxNFTPerWallet *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMaxNFTPerWalletUpdated is a free log retrieval operation binding the contract event 0x9661a815c7db8b3f4e48871cf7d74e5c190a544bfc3eee36ab2451f2dc3d2503.
//
// Solidity: event MaxNFTPerWalletUpdated(uint256 roundId, uint256 maxNFTPerWallet)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterMaxNFTPerWalletUpdated(opts *bind.FilterOpts) (*NFTPresaleManagerMaxNFTPerWalletUpdatedIterator, error) {

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "MaxNFTPerWalletUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerMaxNFTPerWalletUpdatedIterator{contract: _NFTPresaleManager.contract, event: "MaxNFTPerWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxNFTPerWalletUpdated is a free log subscription operation binding the contract event 0x9661a815c7db8b3f4e48871cf7d74e5c190a544bfc3eee36ab2451f2dc3d2503.
//
// Solidity: event MaxNFTPerWalletUpdated(uint256 roundId, uint256 maxNFTPerWallet)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchMaxNFTPerWalletUpdated(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerMaxNFTPerWalletUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "MaxNFTPerWalletUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerMaxNFTPerWalletUpdated)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "MaxNFTPerWalletUpdated", log); err != nil {
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

// ParseMaxNFTPerWalletUpdated is a log parse operation binding the contract event 0x9661a815c7db8b3f4e48871cf7d74e5c190a544bfc3eee36ab2451f2dc3d2503.
//
// Solidity: event MaxNFTPerWalletUpdated(uint256 roundId, uint256 maxNFTPerWallet)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseMaxNFTPerWalletUpdated(log types.Log) (*NFTPresaleManagerMaxNFTPerWalletUpdated, error) {
	event := new(NFTPresaleManagerMaxNFTPerWalletUpdated)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "MaxNFTPerWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerNFTPurchasedIterator is returned from FilterNFTPurchased and is used to iterate over the raw logs and unpacked data for NFTPurchased events raised by the NFTPresaleManager contract.
type NFTPresaleManagerNFTPurchasedIterator struct {
	Event *NFTPresaleManagerNFTPurchased // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerNFTPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerNFTPurchased)
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
		it.Event = new(NFTPresaleManagerNFTPurchased)
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
func (it *NFTPresaleManagerNFTPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerNFTPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerNFTPurchased represents a NFTPurchased event raised by the NFTPresaleManager contract.
type NFTPresaleManagerNFTPurchased struct {
	Buyer   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTPurchased is a free log retrieval operation binding the contract event 0x7f8aed29394d49180b299699236f142fbd3a78f6d1464929dd17e87b73ebc8d2.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 tokenId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterNFTPurchased(opts *bind.FilterOpts, buyer []common.Address) (*NFTPresaleManagerNFTPurchasedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "NFTPurchased", buyerRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerNFTPurchasedIterator{contract: _NFTPresaleManager.contract, event: "NFTPurchased", logs: logs, sub: sub}, nil
}

// WatchNFTPurchased is a free log subscription operation binding the contract event 0x7f8aed29394d49180b299699236f142fbd3a78f6d1464929dd17e87b73ebc8d2.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 tokenId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchNFTPurchased(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerNFTPurchased, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "NFTPurchased", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerNFTPurchased)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "NFTPurchased", log); err != nil {
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

// ParseNFTPurchased is a log parse operation binding the contract event 0x7f8aed29394d49180b299699236f142fbd3a78f6d1464929dd17e87b73ebc8d2.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 tokenId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseNFTPurchased(log types.Log) (*NFTPresaleManagerNFTPurchased, error) {
	event := new(NFTPresaleManagerNFTPurchased)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "NFTPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTPresaleManager contract.
type NFTPresaleManagerOwnershipTransferredIterator struct {
	Event *NFTPresaleManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerOwnershipTransferred)
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
		it.Event = new(NFTPresaleManagerOwnershipTransferred)
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
func (it *NFTPresaleManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerOwnershipTransferred represents a OwnershipTransferred event raised by the NFTPresaleManager contract.
type NFTPresaleManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTPresaleManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerOwnershipTransferredIterator{contract: _NFTPresaleManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerOwnershipTransferred)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NFTPresaleManagerOwnershipTransferred, error) {
	event := new(NFTPresaleManagerOwnershipTransferred)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerPaymentTokenUpdatedIterator is returned from FilterPaymentTokenUpdated and is used to iterate over the raw logs and unpacked data for PaymentTokenUpdated events raised by the NFTPresaleManager contract.
type NFTPresaleManagerPaymentTokenUpdatedIterator struct {
	Event *NFTPresaleManagerPaymentTokenUpdated // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerPaymentTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerPaymentTokenUpdated)
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
		it.Event = new(NFTPresaleManagerPaymentTokenUpdated)
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
func (it *NFTPresaleManagerPaymentTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerPaymentTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerPaymentTokenUpdated represents a PaymentTokenUpdated event raised by the NFTPresaleManager contract.
type NFTPresaleManagerPaymentTokenUpdated struct {
	Token    common.Address
	Accepted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenUpdated is a free log retrieval operation binding the contract event 0x19c6e2fbd27c3204efb106f0081d2d7ffaf44da224ca04fbc0d8a262ba40f555.
//
// Solidity: event PaymentTokenUpdated(address indexed token, bool accepted)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterPaymentTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*NFTPresaleManagerPaymentTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "PaymentTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerPaymentTokenUpdatedIterator{contract: _NFTPresaleManager.contract, event: "PaymentTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenUpdated is a free log subscription operation binding the contract event 0x19c6e2fbd27c3204efb106f0081d2d7ffaf44da224ca04fbc0d8a262ba40f555.
//
// Solidity: event PaymentTokenUpdated(address indexed token, bool accepted)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchPaymentTokenUpdated(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerPaymentTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "PaymentTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerPaymentTokenUpdated)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "PaymentTokenUpdated", log); err != nil {
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

// ParsePaymentTokenUpdated is a log parse operation binding the contract event 0x19c6e2fbd27c3204efb106f0081d2d7ffaf44da224ca04fbc0d8a262ba40f555.
//
// Solidity: event PaymentTokenUpdated(address indexed token, bool accepted)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParsePaymentTokenUpdated(log types.Log) (*NFTPresaleManagerPaymentTokenUpdated, error) {
	event := new(NFTPresaleManagerPaymentTokenUpdated)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "PaymentTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerReferralCommissionPaidIterator is returned from FilterReferralCommissionPaid and is used to iterate over the raw logs and unpacked data for ReferralCommissionPaid events raised by the NFTPresaleManager contract.
type NFTPresaleManagerReferralCommissionPaidIterator struct {
	Event *NFTPresaleManagerReferralCommissionPaid // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerReferralCommissionPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerReferralCommissionPaid)
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
		it.Event = new(NFTPresaleManagerReferralCommissionPaid)
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
func (it *NFTPresaleManagerReferralCommissionPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerReferralCommissionPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerReferralCommissionPaid represents a ReferralCommissionPaid event raised by the NFTPresaleManager contract.
type NFTPresaleManagerReferralCommissionPaid struct {
	Referrer common.Address
	Buyer    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReferralCommissionPaid is a free log retrieval operation binding the contract event 0x86ddab457291316e0f5496737e5ca67c4037234c32c3be04c48ae96186893a7b.
//
// Solidity: event ReferralCommissionPaid(address indexed referrer, address indexed buyer, uint256 amount)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterReferralCommissionPaid(opts *bind.FilterOpts, referrer []common.Address, buyer []common.Address) (*NFTPresaleManagerReferralCommissionPaidIterator, error) {

	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "ReferralCommissionPaid", referrerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerReferralCommissionPaidIterator{contract: _NFTPresaleManager.contract, event: "ReferralCommissionPaid", logs: logs, sub: sub}, nil
}

// WatchReferralCommissionPaid is a free log subscription operation binding the contract event 0x86ddab457291316e0f5496737e5ca67c4037234c32c3be04c48ae96186893a7b.
//
// Solidity: event ReferralCommissionPaid(address indexed referrer, address indexed buyer, uint256 amount)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchReferralCommissionPaid(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerReferralCommissionPaid, referrer []common.Address, buyer []common.Address) (event.Subscription, error) {

	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "ReferralCommissionPaid", referrerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerReferralCommissionPaid)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "ReferralCommissionPaid", log); err != nil {
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

// ParseReferralCommissionPaid is a log parse operation binding the contract event 0x86ddab457291316e0f5496737e5ca67c4037234c32c3be04c48ae96186893a7b.
//
// Solidity: event ReferralCommissionPaid(address indexed referrer, address indexed buyer, uint256 amount)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseReferralCommissionPaid(log types.Log) (*NFTPresaleManagerReferralCommissionPaid, error) {
	event := new(NFTPresaleManagerReferralCommissionPaid)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "ReferralCommissionPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the NFTPresaleManager contract.
type NFTPresaleManagerRemovedFromWhitelistIterator struct {
	Event *NFTPresaleManagerRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerRemovedFromWhitelist)
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
		it.Event = new(NFTPresaleManagerRemovedFromWhitelist)
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
func (it *NFTPresaleManagerRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the NFTPresaleManager contract.
type NFTPresaleManagerRemovedFromWhitelist struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts, account []common.Address) (*NFTPresaleManagerRemovedFromWhitelistIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "RemovedFromWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerRemovedFromWhitelistIterator{contract: _NFTPresaleManager.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerRemovedFromWhitelist, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "RemovedFromWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerRemovedFromWhitelist)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
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

// ParseRemovedFromWhitelist is a log parse operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseRemovedFromWhitelist(log types.Log) (*NFTPresaleManagerRemovedFromWhitelist, error) {
	event := new(NFTPresaleManagerRemovedFromWhitelist)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerSaleRoundCreatedIterator is returned from FilterSaleRoundCreated and is used to iterate over the raw logs and unpacked data for SaleRoundCreated events raised by the NFTPresaleManager contract.
type NFTPresaleManagerSaleRoundCreatedIterator struct {
	Event *NFTPresaleManagerSaleRoundCreated // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerSaleRoundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerSaleRoundCreated)
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
		it.Event = new(NFTPresaleManagerSaleRoundCreated)
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
func (it *NFTPresaleManagerSaleRoundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerSaleRoundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerSaleRoundCreated represents a SaleRoundCreated event raised by the NFTPresaleManager contract.
type NFTPresaleManagerSaleRoundCreated struct {
	RoundId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleRoundCreated is a free log retrieval operation binding the contract event 0xa678db96b3cef44cebd44ad25b4a0518d5fe891d2b92341f31adc288b84a9137.
//
// Solidity: event SaleRoundCreated(uint256 indexed roundId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterSaleRoundCreated(opts *bind.FilterOpts, roundId []*big.Int) (*NFTPresaleManagerSaleRoundCreatedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "SaleRoundCreated", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerSaleRoundCreatedIterator{contract: _NFTPresaleManager.contract, event: "SaleRoundCreated", logs: logs, sub: sub}, nil
}

// WatchSaleRoundCreated is a free log subscription operation binding the contract event 0xa678db96b3cef44cebd44ad25b4a0518d5fe891d2b92341f31adc288b84a9137.
//
// Solidity: event SaleRoundCreated(uint256 indexed roundId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchSaleRoundCreated(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerSaleRoundCreated, roundId []*big.Int) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "SaleRoundCreated", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerSaleRoundCreated)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "SaleRoundCreated", log); err != nil {
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

// ParseSaleRoundCreated is a log parse operation binding the contract event 0xa678db96b3cef44cebd44ad25b4a0518d5fe891d2b92341f31adc288b84a9137.
//
// Solidity: event SaleRoundCreated(uint256 indexed roundId)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseSaleRoundCreated(log types.Log) (*NFTPresaleManagerSaleRoundCreated, error) {
	event := new(NFTPresaleManagerSaleRoundCreated)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "SaleRoundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPresaleManagerWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the NFTPresaleManager contract.
type NFTPresaleManagerWhitelistedIterator struct {
	Event *NFTPresaleManagerWhitelisted // Event containing the contract specifics and raw log

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
func (it *NFTPresaleManagerWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPresaleManagerWhitelisted)
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
		it.Event = new(NFTPresaleManagerWhitelisted)
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
func (it *NFTPresaleManagerWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPresaleManagerWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPresaleManagerWhitelisted represents a Whitelisted event raised by the NFTPresaleManager contract.
type NFTPresaleManagerWhitelisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) FilterWhitelisted(opts *bind.FilterOpts, account []common.Address) (*NFTPresaleManagerWhitelistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.FilterLogs(opts, "Whitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &NFTPresaleManagerWhitelistedIterator{contract: _NFTPresaleManager.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *NFTPresaleManagerWhitelisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFTPresaleManager.contract.WatchLogs(opts, "Whitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPresaleManagerWhitelisted)
				if err := _NFTPresaleManager.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed account)
func (_NFTPresaleManager *NFTPresaleManagerFilterer) ParseWhitelisted(log types.Log) (*NFTPresaleManagerWhitelisted, error) {
	event := new(NFTPresaleManagerWhitelisted)
	if err := _NFTPresaleManager.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
