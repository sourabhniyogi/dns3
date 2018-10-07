// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package dns3

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DNS3ABI is the input ABI used to generate the binding from.
const DNS3ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Buyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"OwnedDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"getZone\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"name\":\"_canBuy\",\"type\":\"address\"}],\"name\":\"approvedBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"cancelBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBlkNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"acquireDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"releaseDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_domain\",\"type\":\"string\"}],\"name\":\"registerDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PublishedDomains\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DomainName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZoneHash\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"},{\"name\":\"blkNum\",\"type\":\"uint256\"}],\"name\":\"updateDomains\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"},{\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"submitZone\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"Domain\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"DomainHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"Owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Deposit\",\"type\":\"uint256\"}],\"name\":\"DomainClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"DomainHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newDeposit\",\"type\":\"uint256\"}],\"name\":\"DomainTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"DomainHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"Owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"halfRefund\",\"type\":\"uint256\"}],\"name\":\"DomainRelease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"blkNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint8\"}],\"name\":\"DomainsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"DomainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint8\"}],\"name\":\"ZoneUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"DomainHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"approvedBuyer\",\"type\":\"address\"}],\"name\":\"ApprovedBuyer\",\"type\":\"event\"}]"

// DNS3 is an auto generated Go binding around an Ethereum contract.
type DNS3 struct {
	DNS3Caller     // Read-only binding to the contract
	DNS3Transactor // Write-only binding to the contract
	DNS3Filterer   // Log filterer for contract events
}

// DNS3Caller is an auto generated read-only Go binding around an Ethereum contract.
type DNS3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNS3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DNS3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNS3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNS3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNS3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNS3Session struct {
	Contract     *DNS3             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNS3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNS3CallerSession struct {
	Contract *DNS3Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DNS3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNS3TransactorSession struct {
	Contract     *DNS3Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNS3Raw is an auto generated low-level Go binding around an Ethereum contract.
type DNS3Raw struct {
	Contract *DNS3 // Generic contract binding to access the raw methods on
}

// DNS3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNS3CallerRaw struct {
	Contract *DNS3Caller // Generic read-only contract binding to access the raw methods on
}

// DNS3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNS3TransactorRaw struct {
	Contract *DNS3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDNS3 creates a new instance of DNS3, bound to a specific deployed contract.
func NewDNS3(address common.Address, backend bind.ContractBackend) (*DNS3, error) {
	contract, err := bindDNS3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNS3{DNS3Caller: DNS3Caller{contract: contract}, DNS3Transactor: DNS3Transactor{contract: contract}, DNS3Filterer: DNS3Filterer{contract: contract}}, nil
}

// NewDNS3Caller creates a new read-only instance of DNS3, bound to a specific deployed contract.
func NewDNS3Caller(address common.Address, caller bind.ContractCaller) (*DNS3Caller, error) {
	contract, err := bindDNS3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNS3Caller{contract: contract}, nil
}

// NewDNS3Transactor creates a new write-only instance of DNS3, bound to a specific deployed contract.
func NewDNS3Transactor(address common.Address, transactor bind.ContractTransactor) (*DNS3Transactor, error) {
	contract, err := bindDNS3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNS3Transactor{contract: contract}, nil
}

// NewDNS3Filterer creates a new log filterer instance of DNS3, bound to a specific deployed contract.
func NewDNS3Filterer(address common.Address, filterer bind.ContractFilterer) (*DNS3Filterer, error) {
	contract, err := bindDNS3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNS3Filterer{contract: contract}, nil
}

// bindDNS3 binds a generic wrapper to an already deployed contract.
func bindDNS3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNS3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNS3 *DNS3Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNS3.Contract.DNS3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNS3 *DNS3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNS3.Contract.DNS3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNS3 *DNS3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNS3.Contract.DNS3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNS3 *DNS3CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNS3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNS3 *DNS3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNS3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNS3 *DNS3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNS3.Contract.contract.Transact(opts, method, params...)
}

// Buyer is a free data retrieval call binding the contract method 0x0a5c9f49.
//
// Solidity: function Buyer( bytes32) constant returns(address)
func (_DNS3 *DNS3Caller) Buyer(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNS3.contract.Call(opts, out, "Buyer", arg0)
	return *ret0, err
}

// Buyer is a free data retrieval call binding the contract method 0x0a5c9f49.
//
// Solidity: function Buyer( bytes32) constant returns(address)
func (_DNS3 *DNS3Session) Buyer(arg0 [32]byte) (common.Address, error) {
	return _DNS3.Contract.Buyer(&_DNS3.CallOpts, arg0)
}

// Buyer is a free data retrieval call binding the contract method 0x0a5c9f49.
//
// Solidity: function Buyer( bytes32) constant returns(address)
func (_DNS3 *DNS3CallerSession) Buyer(arg0 [32]byte) (common.Address, error) {
	return _DNS3.Contract.Buyer(&_DNS3.CallOpts, arg0)
}

// DomainName is a free data retrieval call binding the contract method 0xac97ec57.
//
// Solidity: function DomainName( bytes32) constant returns(string)
func (_DNS3 *DNS3Caller) DomainName(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DNS3.contract.Call(opts, out, "DomainName", arg0)
	return *ret0, err
}

// DomainName is a free data retrieval call binding the contract method 0xac97ec57.
//
// Solidity: function DomainName( bytes32) constant returns(string)
func (_DNS3 *DNS3Session) DomainName(arg0 [32]byte) (string, error) {
	return _DNS3.Contract.DomainName(&_DNS3.CallOpts, arg0)
}

// DomainName is a free data retrieval call binding the contract method 0xac97ec57.
//
// Solidity: function DomainName( bytes32) constant returns(string)
func (_DNS3 *DNS3CallerSession) DomainName(arg0 [32]byte) (string, error) {
	return _DNS3.Contract.DomainName(&_DNS3.CallOpts, arg0)
}

// OwnedDomain is a free data retrieval call binding the contract method 0x3dbaf06b.
//
// Solidity: function OwnedDomain( bytes32) constant returns(address)
func (_DNS3 *DNS3Caller) OwnedDomain(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNS3.contract.Call(opts, out, "OwnedDomain", arg0)
	return *ret0, err
}

// OwnedDomain is a free data retrieval call binding the contract method 0x3dbaf06b.
//
// Solidity: function OwnedDomain( bytes32) constant returns(address)
func (_DNS3 *DNS3Session) OwnedDomain(arg0 [32]byte) (common.Address, error) {
	return _DNS3.Contract.OwnedDomain(&_DNS3.CallOpts, arg0)
}

// OwnedDomain is a free data retrieval call binding the contract method 0x3dbaf06b.
//
// Solidity: function OwnedDomain( bytes32) constant returns(address)
func (_DNS3 *DNS3CallerSession) OwnedDomain(arg0 [32]byte) (common.Address, error) {
	return _DNS3.Contract.OwnedDomain(&_DNS3.CallOpts, arg0)
}

// PublishedDomains is a free data retrieval call binding the contract method 0xa7eee31d.
//
// Solidity: function PublishedDomains( uint256) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Caller) PublishedDomains(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _DNS3.contract.Call(opts, out, "PublishedDomains", arg0)
	return *ret, err
}

// PublishedDomains is a free data retrieval call binding the contract method 0xa7eee31d.
//
// Solidity: function PublishedDomains( uint256) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Session) PublishedDomains(arg0 *big.Int) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.PublishedDomains(&_DNS3.CallOpts, arg0)
}

// PublishedDomains is a free data retrieval call binding the contract method 0xa7eee31d.
//
// Solidity: function PublishedDomains( uint256) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3CallerSession) PublishedDomains(arg0 *big.Int) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.PublishedDomains(&_DNS3.CallOpts, arg0)
}

// ZoneHash is a free data retrieval call binding the contract method 0xb0083f94.
//
// Solidity: function ZoneHash( bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Caller) ZoneHash(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _DNS3.contract.Call(opts, out, "ZoneHash", arg0)
	return *ret, err
}

// ZoneHash is a free data retrieval call binding the contract method 0xb0083f94.
//
// Solidity: function ZoneHash( bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Session) ZoneHash(arg0 [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.ZoneHash(&_DNS3.CallOpts, arg0)
}

// ZoneHash is a free data retrieval call binding the contract method 0xb0083f94.
//
// Solidity: function ZoneHash( bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3CallerSession) ZoneHash(arg0 [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.ZoneHash(&_DNS3.CallOpts, arg0)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DNS3 *DNS3Caller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNS3.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DNS3 *DNS3Session) Authority() (common.Address, error) {
	return _DNS3.Contract.Authority(&_DNS3.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DNS3 *DNS3CallerSession) Authority() (common.Address, error) {
	return _DNS3.Contract.Authority(&_DNS3.CallOpts)
}

// CurrentBlkNum is a free data retrieval call binding the contract method 0x8c1cdd5d.
//
// Solidity: function currentBlkNum() constant returns(uint256)
func (_DNS3 *DNS3Caller) CurrentBlkNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DNS3.contract.Call(opts, out, "currentBlkNum")
	return *ret0, err
}

// CurrentBlkNum is a free data retrieval call binding the contract method 0x8c1cdd5d.
//
// Solidity: function currentBlkNum() constant returns(uint256)
func (_DNS3 *DNS3Session) CurrentBlkNum() (*big.Int, error) {
	return _DNS3.Contract.CurrentBlkNum(&_DNS3.CallOpts)
}

// CurrentBlkNum is a free data retrieval call binding the contract method 0x8c1cdd5d.
//
// Solidity: function currentBlkNum() constant returns(uint256)
func (_DNS3 *DNS3CallerSession) CurrentBlkNum() (*big.Int, error) {
	return _DNS3.Contract.CurrentBlkNum(&_DNS3.CallOpts)
}

// GetZone is a free data retrieval call binding the contract method 0x4b42e8d5.
//
// Solidity: function getZone(domainHash bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Caller) GetZone(opts *bind.CallOpts, domainHash [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _DNS3.contract.Call(opts, out, "getZone", domainHash)
	return *ret, err
}

// GetZone is a free data retrieval call binding the contract method 0x4b42e8d5.
//
// Solidity: function getZone(domainHash bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Session) GetZone(domainHash [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.GetZone(&_DNS3.CallOpts, domainHash)
}

// GetZone is a free data retrieval call binding the contract method 0x4b42e8d5.
//
// Solidity: function getZone(domainHash bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3CallerSession) GetZone(domainHash [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _DNS3.Contract.GetZone(&_DNS3.CallOpts, domainHash)
}

// AcquireDomain is a paid mutator transaction binding the contract method 0x8d520a3e.
//
// Solidity: function acquireDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Transactor) AcquireDomain(opts *bind.TransactOpts, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "acquireDomain", domainHash)
}

// AcquireDomain is a paid mutator transaction binding the contract method 0x8d520a3e.
//
// Solidity: function acquireDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Session) AcquireDomain(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.AcquireDomain(&_DNS3.TransactOpts, domainHash)
}

// AcquireDomain is a paid mutator transaction binding the contract method 0x8d520a3e.
//
// Solidity: function acquireDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3TransactorSession) AcquireDomain(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.AcquireDomain(&_DNS3.TransactOpts, domainHash)
}

// ApprovedBuyer is a paid mutator transaction binding the contract method 0x54623f1a.
//
// Solidity: function approvedBuyer(domainHash bytes32, _canBuy address) returns(bool)
func (_DNS3 *DNS3Transactor) ApprovedBuyer(opts *bind.TransactOpts, domainHash [32]byte, _canBuy common.Address) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "approvedBuyer", domainHash, _canBuy)
}

// ApprovedBuyer is a paid mutator transaction binding the contract method 0x54623f1a.
//
// Solidity: function approvedBuyer(domainHash bytes32, _canBuy address) returns(bool)
func (_DNS3 *DNS3Session) ApprovedBuyer(domainHash [32]byte, _canBuy common.Address) (*types.Transaction, error) {
	return _DNS3.Contract.ApprovedBuyer(&_DNS3.TransactOpts, domainHash, _canBuy)
}

// ApprovedBuyer is a paid mutator transaction binding the contract method 0x54623f1a.
//
// Solidity: function approvedBuyer(domainHash bytes32, _canBuy address) returns(bool)
func (_DNS3 *DNS3TransactorSession) ApprovedBuyer(domainHash [32]byte, _canBuy common.Address) (*types.Transaction, error) {
	return _DNS3.Contract.ApprovedBuyer(&_DNS3.TransactOpts, domainHash, _canBuy)
}

// CancelBuyer is a paid mutator transaction binding the contract method 0x7a0c87db.
//
// Solidity: function cancelBuyer(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Transactor) CancelBuyer(opts *bind.TransactOpts, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "cancelBuyer", domainHash)
}

// CancelBuyer is a paid mutator transaction binding the contract method 0x7a0c87db.
//
// Solidity: function cancelBuyer(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Session) CancelBuyer(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.CancelBuyer(&_DNS3.TransactOpts, domainHash)
}

// CancelBuyer is a paid mutator transaction binding the contract method 0x7a0c87db.
//
// Solidity: function cancelBuyer(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3TransactorSession) CancelBuyer(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.CancelBuyer(&_DNS3.TransactOpts, domainHash)
}

// RegisterDomain is a paid mutator transaction binding the contract method 0xa5cad08a.
//
// Solidity: function registerDomain(_domain string) returns(bool)
func (_DNS3 *DNS3Transactor) RegisterDomain(opts *bind.TransactOpts, _domain string) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "registerDomain", _domain)
}

// RegisterDomain is a paid mutator transaction binding the contract method 0xa5cad08a.
//
// Solidity: function registerDomain(_domain string) returns(bool)
func (_DNS3 *DNS3Session) RegisterDomain(_domain string) (*types.Transaction, error) {
	return _DNS3.Contract.RegisterDomain(&_DNS3.TransactOpts, _domain)
}

// RegisterDomain is a paid mutator transaction binding the contract method 0xa5cad08a.
//
// Solidity: function registerDomain(_domain string) returns(bool)
func (_DNS3 *DNS3TransactorSession) RegisterDomain(_domain string) (*types.Transaction, error) {
	return _DNS3.Contract.RegisterDomain(&_DNS3.TransactOpts, _domain)
}

// ReleaseDomain is a paid mutator transaction binding the contract method 0x9e60fe19.
//
// Solidity: function releaseDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Transactor) ReleaseDomain(opts *bind.TransactOpts, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "releaseDomain", domainHash)
}

// ReleaseDomain is a paid mutator transaction binding the contract method 0x9e60fe19.
//
// Solidity: function releaseDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Session) ReleaseDomain(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.ReleaseDomain(&_DNS3.TransactOpts, domainHash)
}

// ReleaseDomain is a paid mutator transaction binding the contract method 0x9e60fe19.
//
// Solidity: function releaseDomain(domainHash bytes32) returns(bool)
func (_DNS3 *DNS3TransactorSession) ReleaseDomain(domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.ReleaseDomain(&_DNS3.TransactOpts, domainHash)
}

// SubmitZone is a paid mutator transaction binding the contract method 0xc6b2199b.
//
// Solidity: function submitZone(digest bytes32, hashFunction uint8, size uint8, domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Transactor) SubmitZone(opts *bind.TransactOpts, digest [32]byte, hashFunction uint8, size uint8, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "submitZone", digest, hashFunction, size, domainHash)
}

// SubmitZone is a paid mutator transaction binding the contract method 0xc6b2199b.
//
// Solidity: function submitZone(digest bytes32, hashFunction uint8, size uint8, domainHash bytes32) returns(bool)
func (_DNS3 *DNS3Session) SubmitZone(digest [32]byte, hashFunction uint8, size uint8, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.SubmitZone(&_DNS3.TransactOpts, digest, hashFunction, size, domainHash)
}

// SubmitZone is a paid mutator transaction binding the contract method 0xc6b2199b.
//
// Solidity: function submitZone(digest bytes32, hashFunction uint8, size uint8, domainHash bytes32) returns(bool)
func (_DNS3 *DNS3TransactorSession) SubmitZone(digest [32]byte, hashFunction uint8, size uint8, domainHash [32]byte) (*types.Transaction, error) {
	return _DNS3.Contract.SubmitZone(&_DNS3.TransactOpts, digest, hashFunction, size, domainHash)
}

// UpdateDomains is a paid mutator transaction binding the contract method 0xb25037f2.
//
// Solidity: function updateDomains(digest bytes32, hashFunction uint8, size uint8, blkNum uint256) returns()
func (_DNS3 *DNS3Transactor) UpdateDomains(opts *bind.TransactOpts, digest [32]byte, hashFunction uint8, size uint8, blkNum *big.Int) (*types.Transaction, error) {
	return _DNS3.contract.Transact(opts, "updateDomains", digest, hashFunction, size, blkNum)
}

// UpdateDomains is a paid mutator transaction binding the contract method 0xb25037f2.
//
// Solidity: function updateDomains(digest bytes32, hashFunction uint8, size uint8, blkNum uint256) returns()
func (_DNS3 *DNS3Session) UpdateDomains(digest [32]byte, hashFunction uint8, size uint8, blkNum *big.Int) (*types.Transaction, error) {
	return _DNS3.Contract.UpdateDomains(&_DNS3.TransactOpts, digest, hashFunction, size, blkNum)
}

// UpdateDomains is a paid mutator transaction binding the contract method 0xb25037f2.
//
// Solidity: function updateDomains(digest bytes32, hashFunction uint8, size uint8, blkNum uint256) returns()
func (_DNS3 *DNS3TransactorSession) UpdateDomains(digest [32]byte, hashFunction uint8, size uint8, blkNum *big.Int) (*types.Transaction, error) {
	return _DNS3.Contract.UpdateDomains(&_DNS3.TransactOpts, digest, hashFunction, size, blkNum)
}

// DNS3ApprovedBuyerIterator is returned from FilterApprovedBuyer and is used to iterate over the raw logs and unpacked data for ApprovedBuyer events raised by the DNS3 contract.
type DNS3ApprovedBuyerIterator struct {
	Event *DNS3ApprovedBuyer // Event containing the contract specifics and raw log

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
func (it *DNS3ApprovedBuyerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3ApprovedBuyer)
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
		it.Event = new(DNS3ApprovedBuyer)
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
func (it *DNS3ApprovedBuyerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3ApprovedBuyerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3ApprovedBuyer represents a ApprovedBuyer event raised by the DNS3 contract.
type DNS3ApprovedBuyer struct {
	DomainHash    [32]byte
	ApprovedBuyer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApprovedBuyer is a free log retrieval operation binding the contract event 0x8ff5938703189b358f138254b13e02f879f688d7870358359d06a91ab44d655e.
//
// Solidity: e ApprovedBuyer(DomainHash indexed bytes32, approvedBuyer indexed address)
func (_DNS3 *DNS3Filterer) FilterApprovedBuyer(opts *bind.FilterOpts, DomainHash [][32]byte, approvedBuyer []common.Address) (*DNS3ApprovedBuyerIterator, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var approvedBuyerRule []interface{}
	for _, approvedBuyerItem := range approvedBuyer {
		approvedBuyerRule = append(approvedBuyerRule, approvedBuyerItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "ApprovedBuyer", DomainHashRule, approvedBuyerRule)
	if err != nil {
		return nil, err
	}
	return &DNS3ApprovedBuyerIterator{contract: _DNS3.contract, event: "ApprovedBuyer", logs: logs, sub: sub}, nil
}

// WatchApprovedBuyer is a free log subscription operation binding the contract event 0x8ff5938703189b358f138254b13e02f879f688d7870358359d06a91ab44d655e.
//
// Solidity: e ApprovedBuyer(DomainHash indexed bytes32, approvedBuyer indexed address)
func (_DNS3 *DNS3Filterer) WatchApprovedBuyer(opts *bind.WatchOpts, sink chan<- *DNS3ApprovedBuyer, DomainHash [][32]byte, approvedBuyer []common.Address) (event.Subscription, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var approvedBuyerRule []interface{}
	for _, approvedBuyerItem := range approvedBuyer {
		approvedBuyerRule = append(approvedBuyerRule, approvedBuyerItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "ApprovedBuyer", DomainHashRule, approvedBuyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3ApprovedBuyer)
				if err := _DNS3.contract.UnpackLog(event, "ApprovedBuyer", log); err != nil {
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

// DNS3DomainClaimIterator is returned from FilterDomainClaim and is used to iterate over the raw logs and unpacked data for DomainClaim events raised by the DNS3 contract.
type DNS3DomainClaimIterator struct {
	Event *DNS3DomainClaim // Event containing the contract specifics and raw log

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
func (it *DNS3DomainClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3DomainClaim)
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
		it.Event = new(DNS3DomainClaim)
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
func (it *DNS3DomainClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3DomainClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3DomainClaim represents a DomainClaim event raised by the DNS3 contract.
type DNS3DomainClaim struct {
	Domain     string
	DomainHash [32]byte
	Owner      common.Address
	Deposit    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainClaim is a free log retrieval operation binding the contract event 0x846d545d61d5e79431cd6c42d63c0f43e3139996333fdc601d7a392792f5483a.
//
// Solidity: e DomainClaim(Domain string, DomainHash indexed bytes32, Owner indexed address, Deposit uint256)
func (_DNS3 *DNS3Filterer) FilterDomainClaim(opts *bind.FilterOpts, DomainHash [][32]byte, Owner []common.Address) (*DNS3DomainClaimIterator, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var OwnerRule []interface{}
	for _, OwnerItem := range Owner {
		OwnerRule = append(OwnerRule, OwnerItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "DomainClaim", DomainHashRule, OwnerRule)
	if err != nil {
		return nil, err
	}
	return &DNS3DomainClaimIterator{contract: _DNS3.contract, event: "DomainClaim", logs: logs, sub: sub}, nil
}

// WatchDomainClaim is a free log subscription operation binding the contract event 0x846d545d61d5e79431cd6c42d63c0f43e3139996333fdc601d7a392792f5483a.
//
// Solidity: e DomainClaim(Domain string, DomainHash indexed bytes32, Owner indexed address, Deposit uint256)
func (_DNS3 *DNS3Filterer) WatchDomainClaim(opts *bind.WatchOpts, sink chan<- *DNS3DomainClaim, DomainHash [][32]byte, Owner []common.Address) (event.Subscription, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var OwnerRule []interface{}
	for _, OwnerItem := range Owner {
		OwnerRule = append(OwnerRule, OwnerItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "DomainClaim", DomainHashRule, OwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3DomainClaim)
				if err := _DNS3.contract.UnpackLog(event, "DomainClaim", log); err != nil {
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

// DNS3DomainReleaseIterator is returned from FilterDomainRelease and is used to iterate over the raw logs and unpacked data for DomainRelease events raised by the DNS3 contract.
type DNS3DomainReleaseIterator struct {
	Event *DNS3DomainRelease // Event containing the contract specifics and raw log

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
func (it *DNS3DomainReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3DomainRelease)
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
		it.Event = new(DNS3DomainRelease)
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
func (it *DNS3DomainReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3DomainReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3DomainRelease represents a DomainRelease event raised by the DNS3 contract.
type DNS3DomainRelease struct {
	DomainHash [32]byte
	Owner      common.Address
	HalfRefund *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainRelease is a free log retrieval operation binding the contract event 0x2ec3a970f8978b642b90cc4e5befa7c677d7635100a2edb8fcfb01e2c56d308c.
//
// Solidity: e DomainRelease(DomainHash indexed bytes32, Owner indexed address, halfRefund uint256)
func (_DNS3 *DNS3Filterer) FilterDomainRelease(opts *bind.FilterOpts, DomainHash [][32]byte, Owner []common.Address) (*DNS3DomainReleaseIterator, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var OwnerRule []interface{}
	for _, OwnerItem := range Owner {
		OwnerRule = append(OwnerRule, OwnerItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "DomainRelease", DomainHashRule, OwnerRule)
	if err != nil {
		return nil, err
	}
	return &DNS3DomainReleaseIterator{contract: _DNS3.contract, event: "DomainRelease", logs: logs, sub: sub}, nil
}

// WatchDomainRelease is a free log subscription operation binding the contract event 0x2ec3a970f8978b642b90cc4e5befa7c677d7635100a2edb8fcfb01e2c56d308c.
//
// Solidity: e DomainRelease(DomainHash indexed bytes32, Owner indexed address, halfRefund uint256)
func (_DNS3 *DNS3Filterer) WatchDomainRelease(opts *bind.WatchOpts, sink chan<- *DNS3DomainRelease, DomainHash [][32]byte, Owner []common.Address) (event.Subscription, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var OwnerRule []interface{}
	for _, OwnerItem := range Owner {
		OwnerRule = append(OwnerRule, OwnerItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "DomainRelease", DomainHashRule, OwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3DomainRelease)
				if err := _DNS3.contract.UnpackLog(event, "DomainRelease", log); err != nil {
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

// DNS3DomainTransferIterator is returned from FilterDomainTransfer and is used to iterate over the raw logs and unpacked data for DomainTransfer events raised by the DNS3 contract.
type DNS3DomainTransferIterator struct {
	Event *DNS3DomainTransfer // Event containing the contract specifics and raw log

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
func (it *DNS3DomainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3DomainTransfer)
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
		it.Event = new(DNS3DomainTransfer)
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
func (it *DNS3DomainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3DomainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3DomainTransfer represents a DomainTransfer event raised by the DNS3 contract.
type DNS3DomainTransfer struct {
	DomainHash    [32]byte
	OriginalOwner common.Address
	NewOwner      common.Address
	NewDeposit    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDomainTransfer is a free log retrieval operation binding the contract event 0x4bb7496f8fade8e3b5312d785e6bc7b61b94160334dbf39d63140ea47010b4e8.
//
// Solidity: e DomainTransfer(DomainHash indexed bytes32, originalOwner indexed address, newOwner indexed address, newDeposit uint256)
func (_DNS3 *DNS3Filterer) FilterDomainTransfer(opts *bind.FilterOpts, DomainHash [][32]byte, originalOwner []common.Address, newOwner []common.Address) (*DNS3DomainTransferIterator, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var originalOwnerRule []interface{}
	for _, originalOwnerItem := range originalOwner {
		originalOwnerRule = append(originalOwnerRule, originalOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "DomainTransfer", DomainHashRule, originalOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DNS3DomainTransferIterator{contract: _DNS3.contract, event: "DomainTransfer", logs: logs, sub: sub}, nil
}

// WatchDomainTransfer is a free log subscription operation binding the contract event 0x4bb7496f8fade8e3b5312d785e6bc7b61b94160334dbf39d63140ea47010b4e8.
//
// Solidity: e DomainTransfer(DomainHash indexed bytes32, originalOwner indexed address, newOwner indexed address, newDeposit uint256)
func (_DNS3 *DNS3Filterer) WatchDomainTransfer(opts *bind.WatchOpts, sink chan<- *DNS3DomainTransfer, DomainHash [][32]byte, originalOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}
	var originalOwnerRule []interface{}
	for _, originalOwnerItem := range originalOwner {
		originalOwnerRule = append(originalOwnerRule, originalOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "DomainTransfer", DomainHashRule, originalOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3DomainTransfer)
				if err := _DNS3.contract.UnpackLog(event, "DomainTransfer", log); err != nil {
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

// DNS3DomainsUpdateIterator is returned from FilterDomainsUpdate and is used to iterate over the raw logs and unpacked data for DomainsUpdate events raised by the DNS3 contract.
type DNS3DomainsUpdateIterator struct {
	Event *DNS3DomainsUpdate // Event containing the contract specifics and raw log

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
func (it *DNS3DomainsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3DomainsUpdate)
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
		it.Event = new(DNS3DomainsUpdate)
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
func (it *DNS3DomainsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3DomainsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3DomainsUpdate represents a DomainsUpdate event raised by the DNS3 contract.
type DNS3DomainsUpdate struct {
	BlkNum       *big.Int
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDomainsUpdate is a free log retrieval operation binding the contract event 0x3b44489944e9b19b2dfe9dead7d6f267ed5450bd7250051518352a2b276e38ce.
//
// Solidity: e DomainsUpdate(blkNum indexed uint256, digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Filterer) FilterDomainsUpdate(opts *bind.FilterOpts, blkNum []*big.Int) (*DNS3DomainsUpdateIterator, error) {

	var blkNumRule []interface{}
	for _, blkNumItem := range blkNum {
		blkNumRule = append(blkNumRule, blkNumItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "DomainsUpdate", blkNumRule)
	if err != nil {
		return nil, err
	}
	return &DNS3DomainsUpdateIterator{contract: _DNS3.contract, event: "DomainsUpdate", logs: logs, sub: sub}, nil
}

// WatchDomainsUpdate is a free log subscription operation binding the contract event 0x3b44489944e9b19b2dfe9dead7d6f267ed5450bd7250051518352a2b276e38ce.
//
// Solidity: e DomainsUpdate(blkNum indexed uint256, digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Filterer) WatchDomainsUpdate(opts *bind.WatchOpts, sink chan<- *DNS3DomainsUpdate, blkNum []*big.Int) (event.Subscription, error) {

	var blkNumRule []interface{}
	for _, blkNumItem := range blkNum {
		blkNumRule = append(blkNumRule, blkNumItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "DomainsUpdate", blkNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3DomainsUpdate)
				if err := _DNS3.contract.UnpackLog(event, "DomainsUpdate", log); err != nil {
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

// DNS3ZoneUpdateIterator is returned from FilterZoneUpdate and is used to iterate over the raw logs and unpacked data for ZoneUpdate events raised by the DNS3 contract.
type DNS3ZoneUpdateIterator struct {
	Event *DNS3ZoneUpdate // Event containing the contract specifics and raw log

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
func (it *DNS3ZoneUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNS3ZoneUpdate)
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
		it.Event = new(DNS3ZoneUpdate)
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
func (it *DNS3ZoneUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNS3ZoneUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNS3ZoneUpdate represents a ZoneUpdate event raised by the DNS3 contract.
type DNS3ZoneUpdate struct {
	DomainHash   [32]byte
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterZoneUpdate is a free log retrieval operation binding the contract event 0x75b7002be71f5325c43af0ef07f9bd3a409fc9f01e52c84e97c29f2dc82b30a3.
//
// Solidity: e ZoneUpdate(DomainHash indexed bytes32, digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Filterer) FilterZoneUpdate(opts *bind.FilterOpts, DomainHash [][32]byte) (*DNS3ZoneUpdateIterator, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}

	logs, sub, err := _DNS3.contract.FilterLogs(opts, "ZoneUpdate", DomainHashRule)
	if err != nil {
		return nil, err
	}
	return &DNS3ZoneUpdateIterator{contract: _DNS3.contract, event: "ZoneUpdate", logs: logs, sub: sub}, nil
}

// WatchZoneUpdate is a free log subscription operation binding the contract event 0x75b7002be71f5325c43af0ef07f9bd3a409fc9f01e52c84e97c29f2dc82b30a3.
//
// Solidity: e ZoneUpdate(DomainHash indexed bytes32, digest bytes32, hashFunction uint8, size uint8)
func (_DNS3 *DNS3Filterer) WatchZoneUpdate(opts *bind.WatchOpts, sink chan<- *DNS3ZoneUpdate, DomainHash [][32]byte) (event.Subscription, error) {

	var DomainHashRule []interface{}
	for _, DomainHashItem := range DomainHash {
		DomainHashRule = append(DomainHashRule, DomainHashItem)
	}

	logs, sub, err := _DNS3.contract.WatchLogs(opts, "ZoneUpdate", DomainHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNS3ZoneUpdate)
				if err := _DNS3.contract.UnpackLog(event, "ZoneUpdate", log); err != nil {
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
