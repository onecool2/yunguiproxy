// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package YunguiContract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// YunguiContractABI is the input ABI used to generate the binding from.
const YunguiContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"business_districtArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"fridge_storage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"city_info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mapWriter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"business_district_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"city_info_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"building_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buildingArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"business_area_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"business_areaArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fridge_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ordersArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"city_infoArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fridge_storageArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"fridge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fridgeArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"addWriter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"business_area\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"business_district\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orders_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fridge_storage_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"building\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// YunguiContract is an auto generated Go binding around an Ethereum contract.
type YunguiContract struct {
	YunguiContractCaller     // Read-only binding to the contract
	YunguiContractTransactor // Write-only binding to the contract
	YunguiContractFilterer   // Log filterer for contract events
}

// YunguiContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type YunguiContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YunguiContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YunguiContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YunguiContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YunguiContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YunguiContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YunguiContractSession struct {
	Contract     *YunguiContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YunguiContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YunguiContractCallerSession struct {
	Contract *YunguiContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// YunguiContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YunguiContractTransactorSession struct {
	Contract     *YunguiContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// YunguiContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type YunguiContractRaw struct {
	Contract *YunguiContract // Generic contract binding to access the raw methods on
}

// YunguiContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YunguiContractCallerRaw struct {
	Contract *YunguiContractCaller // Generic read-only contract binding to access the raw methods on
}

// YunguiContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YunguiContractTransactorRaw struct {
	Contract *YunguiContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYunguiContract creates a new instance of YunguiContract, bound to a specific deployed contract.
func NewYunguiContract(address common.Address, backend bind.ContractBackend) (*YunguiContract, error) {
	contract, err := bindYunguiContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YunguiContract{YunguiContractCaller: YunguiContractCaller{contract: contract}, YunguiContractTransactor: YunguiContractTransactor{contract: contract}, YunguiContractFilterer: YunguiContractFilterer{contract: contract}}, nil
}

// NewYunguiContractCaller creates a new read-only instance of YunguiContract, bound to a specific deployed contract.
func NewYunguiContractCaller(address common.Address, caller bind.ContractCaller) (*YunguiContractCaller, error) {
	contract, err := bindYunguiContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YunguiContractCaller{contract: contract}, nil
}

// NewYunguiContractTransactor creates a new write-only instance of YunguiContract, bound to a specific deployed contract.
func NewYunguiContractTransactor(address common.Address, transactor bind.ContractTransactor) (*YunguiContractTransactor, error) {
	contract, err := bindYunguiContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YunguiContractTransactor{contract: contract}, nil
}

// NewYunguiContractFilterer creates a new log filterer instance of YunguiContract, bound to a specific deployed contract.
func NewYunguiContractFilterer(address common.Address, filterer bind.ContractFilterer) (*YunguiContractFilterer, error) {
	contract, err := bindYunguiContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YunguiContractFilterer{contract: contract}, nil
}

// bindYunguiContract binds a generic wrapper to an already deployed contract.
func bindYunguiContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YunguiContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YunguiContract *YunguiContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YunguiContract.Contract.YunguiContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YunguiContract *YunguiContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YunguiContract.Contract.YunguiContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YunguiContract *YunguiContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YunguiContract.Contract.YunguiContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YunguiContract *YunguiContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YunguiContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YunguiContract *YunguiContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YunguiContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YunguiContract *YunguiContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YunguiContract.Contract.contract.Transact(opts, method, params...)
}

// BuildingArray is a free data retrieval call binding the contract method 0x76e71ec3.
//
// Solidity: function buildingArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) BuildingArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "buildingArray", arg0)
	return *ret0, err
}

// BuildingArray is a free data retrieval call binding the contract method 0x76e71ec3.
//
// Solidity: function buildingArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) BuildingArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BuildingArray(&_YunguiContract.CallOpts, arg0)
}

// BuildingArray is a free data retrieval call binding the contract method 0x76e71ec3.
//
// Solidity: function buildingArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) BuildingArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BuildingArray(&_YunguiContract.CallOpts, arg0)
}

// BuildingIndex is a free data retrieval call binding the contract method 0x65c192b4.
//
// Solidity: function building_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) BuildingIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "building_index")
	return *ret0, err
}

// BuildingIndex is a free data retrieval call binding the contract method 0x65c192b4.
//
// Solidity: function building_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) BuildingIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BuildingIndex(&_YunguiContract.CallOpts)
}

// BuildingIndex is a free data retrieval call binding the contract method 0x65c192b4.
//
// Solidity: function building_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) BuildingIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BuildingIndex(&_YunguiContract.CallOpts)
}

// BusinessAreaArray is a free data retrieval call binding the contract method 0x8042267c.
//
// Solidity: function business_areaArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) BusinessAreaArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "business_areaArray", arg0)
	return *ret0, err
}

// BusinessAreaArray is a free data retrieval call binding the contract method 0x8042267c.
//
// Solidity: function business_areaArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) BusinessAreaArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BusinessAreaArray(&_YunguiContract.CallOpts, arg0)
}

// BusinessAreaArray is a free data retrieval call binding the contract method 0x8042267c.
//
// Solidity: function business_areaArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) BusinessAreaArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BusinessAreaArray(&_YunguiContract.CallOpts, arg0)
}

// BusinessAreaIndex is a free data retrieval call binding the contract method 0x7e942474.
//
// Solidity: function business_area_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) BusinessAreaIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "business_area_index")
	return *ret0, err
}

// BusinessAreaIndex is a free data retrieval call binding the contract method 0x7e942474.
//
// Solidity: function business_area_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) BusinessAreaIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BusinessAreaIndex(&_YunguiContract.CallOpts)
}

// BusinessAreaIndex is a free data retrieval call binding the contract method 0x7e942474.
//
// Solidity: function business_area_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) BusinessAreaIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BusinessAreaIndex(&_YunguiContract.CallOpts)
}

// BusinessDistrictArray is a free data retrieval call binding the contract method 0x0cfa3f92.
//
// Solidity: function business_districtArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) BusinessDistrictArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "business_districtArray", arg0)
	return *ret0, err
}

// BusinessDistrictArray is a free data retrieval call binding the contract method 0x0cfa3f92.
//
// Solidity: function business_districtArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) BusinessDistrictArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BusinessDistrictArray(&_YunguiContract.CallOpts, arg0)
}

// BusinessDistrictArray is a free data retrieval call binding the contract method 0x0cfa3f92.
//
// Solidity: function business_districtArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) BusinessDistrictArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.BusinessDistrictArray(&_YunguiContract.CallOpts, arg0)
}

// BusinessDistrictIndex is a free data retrieval call binding the contract method 0x5666c2b6.
//
// Solidity: function business_district_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) BusinessDistrictIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "business_district_index")
	return *ret0, err
}

// BusinessDistrictIndex is a free data retrieval call binding the contract method 0x5666c2b6.
//
// Solidity: function business_district_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) BusinessDistrictIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BusinessDistrictIndex(&_YunguiContract.CallOpts)
}

// BusinessDistrictIndex is a free data retrieval call binding the contract method 0x5666c2b6.
//
// Solidity: function business_district_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) BusinessDistrictIndex() (*big.Int, error) {
	return _YunguiContract.Contract.BusinessDistrictIndex(&_YunguiContract.CallOpts)
}

// CityInfoArray is a free data retrieval call binding the contract method 0xabc21f5d.
//
// Solidity: function city_infoArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) CityInfoArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "city_infoArray", arg0)
	return *ret0, err
}

// CityInfoArray is a free data retrieval call binding the contract method 0xabc21f5d.
//
// Solidity: function city_infoArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) CityInfoArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.CityInfoArray(&_YunguiContract.CallOpts, arg0)
}

// CityInfoArray is a free data retrieval call binding the contract method 0xabc21f5d.
//
// Solidity: function city_infoArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) CityInfoArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.CityInfoArray(&_YunguiContract.CallOpts, arg0)
}

// CityInfoIndex is a free data retrieval call binding the contract method 0x5cd3680a.
//
// Solidity: function city_info_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) CityInfoIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "city_info_index")
	return *ret0, err
}

// CityInfoIndex is a free data retrieval call binding the contract method 0x5cd3680a.
//
// Solidity: function city_info_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) CityInfoIndex() (*big.Int, error) {
	return _YunguiContract.Contract.CityInfoIndex(&_YunguiContract.CallOpts)
}

// CityInfoIndex is a free data retrieval call binding the contract method 0x5cd3680a.
//
// Solidity: function city_info_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) CityInfoIndex() (*big.Int, error) {
	return _YunguiContract.Contract.CityInfoIndex(&_YunguiContract.CallOpts)
}

// FridgeArray is a free data retrieval call binding the contract method 0xd73e570f.
//
// Solidity: function fridgeArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) FridgeArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "fridgeArray", arg0)
	return *ret0, err
}

// FridgeArray is a free data retrieval call binding the contract method 0xd73e570f.
//
// Solidity: function fridgeArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) FridgeArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.FridgeArray(&_YunguiContract.CallOpts, arg0)
}

// FridgeArray is a free data retrieval call binding the contract method 0xd73e570f.
//
// Solidity: function fridgeArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) FridgeArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.FridgeArray(&_YunguiContract.CallOpts, arg0)
}

// FridgeIndex is a free data retrieval call binding the contract method 0x95c01a3a.
//
// Solidity: function fridge_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) FridgeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "fridge_index")
	return *ret0, err
}

// FridgeIndex is a free data retrieval call binding the contract method 0x95c01a3a.
//
// Solidity: function fridge_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) FridgeIndex() (*big.Int, error) {
	return _YunguiContract.Contract.FridgeIndex(&_YunguiContract.CallOpts)
}

// FridgeIndex is a free data retrieval call binding the contract method 0x95c01a3a.
//
// Solidity: function fridge_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) FridgeIndex() (*big.Int, error) {
	return _YunguiContract.Contract.FridgeIndex(&_YunguiContract.CallOpts)
}

// FridgeStorageArray is a free data retrieval call binding the contract method 0xb1b75987.
//
// Solidity: function fridge_storageArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) FridgeStorageArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "fridge_storageArray", arg0)
	return *ret0, err
}

// FridgeStorageArray is a free data retrieval call binding the contract method 0xb1b75987.
//
// Solidity: function fridge_storageArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) FridgeStorageArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.FridgeStorageArray(&_YunguiContract.CallOpts, arg0)
}

// FridgeStorageArray is a free data retrieval call binding the contract method 0xb1b75987.
//
// Solidity: function fridge_storageArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) FridgeStorageArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.FridgeStorageArray(&_YunguiContract.CallOpts, arg0)
}

// FridgeStorageIndex is a free data retrieval call binding the contract method 0xf42dbc78.
//
// Solidity: function fridge_storage_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) FridgeStorageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "fridge_storage_index")
	return *ret0, err
}

// FridgeStorageIndex is a free data retrieval call binding the contract method 0xf42dbc78.
//
// Solidity: function fridge_storage_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) FridgeStorageIndex() (*big.Int, error) {
	return _YunguiContract.Contract.FridgeStorageIndex(&_YunguiContract.CallOpts)
}

// FridgeStorageIndex is a free data retrieval call binding the contract method 0xf42dbc78.
//
// Solidity: function fridge_storage_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) FridgeStorageIndex() (*big.Int, error) {
	return _YunguiContract.Contract.FridgeStorageIndex(&_YunguiContract.CallOpts)
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_YunguiContract *YunguiContractCaller) MapWriter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "mapWriter", arg0)
	return *ret0, err
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_YunguiContract *YunguiContractSession) MapWriter(arg0 common.Address) (bool, error) {
	return _YunguiContract.Contract.MapWriter(&_YunguiContract.CallOpts, arg0)
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_YunguiContract *YunguiContractCallerSession) MapWriter(arg0 common.Address) (bool, error) {
	return _YunguiContract.Contract.MapWriter(&_YunguiContract.CallOpts, arg0)
}

// OrdersArray is a free data retrieval call binding the contract method 0xa96c879e.
//
// Solidity: function ordersArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCaller) OrdersArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "ordersArray", arg0)
	return *ret0, err
}

// OrdersArray is a free data retrieval call binding the contract method 0xa96c879e.
//
// Solidity: function ordersArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractSession) OrdersArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.OrdersArray(&_YunguiContract.CallOpts, arg0)
}

// OrdersArray is a free data retrieval call binding the contract method 0xa96c879e.
//
// Solidity: function ordersArray(uint256 ) constant returns(string)
func (_YunguiContract *YunguiContractCallerSession) OrdersArray(arg0 *big.Int) (string, error) {
	return _YunguiContract.Contract.OrdersArray(&_YunguiContract.CallOpts, arg0)
}

// OrdersIndex is a free data retrieval call binding the contract method 0xed4d809e.
//
// Solidity: function orders_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCaller) OrdersIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "orders_index")
	return *ret0, err
}

// OrdersIndex is a free data retrieval call binding the contract method 0xed4d809e.
//
// Solidity: function orders_index() constant returns(uint256)
func (_YunguiContract *YunguiContractSession) OrdersIndex() (*big.Int, error) {
	return _YunguiContract.Contract.OrdersIndex(&_YunguiContract.CallOpts)
}

// OrdersIndex is a free data retrieval call binding the contract method 0xed4d809e.
//
// Solidity: function orders_index() constant returns(uint256)
func (_YunguiContract *YunguiContractCallerSession) OrdersIndex() (*big.Int, error) {
	return _YunguiContract.Contract.OrdersIndex(&_YunguiContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_YunguiContract *YunguiContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _YunguiContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_YunguiContract *YunguiContractSession) Owner() (common.Address, error) {
	return _YunguiContract.Contract.Owner(&_YunguiContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_YunguiContract *YunguiContractCallerSession) Owner() (common.Address, error) {
	return _YunguiContract.Contract.Owner(&_YunguiContract.CallOpts)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_YunguiContract *YunguiContractTransactor) AddWriter(opts *bind.TransactOpts, writer common.Address) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "addWriter", writer)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_YunguiContract *YunguiContractSession) AddWriter(writer common.Address) (*types.Transaction, error) {
	return _YunguiContract.Contract.AddWriter(&_YunguiContract.TransactOpts, writer)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_YunguiContract *YunguiContractTransactorSession) AddWriter(writer common.Address) (*types.Transaction, error) {
	return _YunguiContract.Contract.AddWriter(&_YunguiContract.TransactOpts, writer)
}

// Building is a paid mutator transaction binding the contract method 0xfb52929d.
//
// Solidity: function building(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) Building(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "building", ss)
}

// Building is a paid mutator transaction binding the contract method 0xfb52929d.
//
// Solidity: function building(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) Building(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Building(&_YunguiContract.TransactOpts, ss)
}

// Building is a paid mutator transaction binding the contract method 0xfb52929d.
//
// Solidity: function building(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) Building(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Building(&_YunguiContract.TransactOpts, ss)
}

// BusinessArea is a paid mutator transaction binding the contract method 0xe6b624ea.
//
// Solidity: function business_area(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) BusinessArea(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "business_area", ss)
}

// BusinessArea is a paid mutator transaction binding the contract method 0xe6b624ea.
//
// Solidity: function business_area(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) BusinessArea(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.BusinessArea(&_YunguiContract.TransactOpts, ss)
}

// BusinessArea is a paid mutator transaction binding the contract method 0xe6b624ea.
//
// Solidity: function business_area(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) BusinessArea(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.BusinessArea(&_YunguiContract.TransactOpts, ss)
}

// BusinessDistrict is a paid mutator transaction binding the contract method 0xed25b996.
//
// Solidity: function business_district(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) BusinessDistrict(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "business_district", ss)
}

// BusinessDistrict is a paid mutator transaction binding the contract method 0xed25b996.
//
// Solidity: function business_district(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) BusinessDistrict(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.BusinessDistrict(&_YunguiContract.TransactOpts, ss)
}

// BusinessDistrict is a paid mutator transaction binding the contract method 0xed25b996.
//
// Solidity: function business_district(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) BusinessDistrict(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.BusinessDistrict(&_YunguiContract.TransactOpts, ss)
}

// CityInfo is a paid mutator transaction binding the contract method 0x233a1a7b.
//
// Solidity: function city_info(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) CityInfo(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "city_info", ss)
}

// CityInfo is a paid mutator transaction binding the contract method 0x233a1a7b.
//
// Solidity: function city_info(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) CityInfo(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.CityInfo(&_YunguiContract.TransactOpts, ss)
}

// CityInfo is a paid mutator transaction binding the contract method 0x233a1a7b.
//
// Solidity: function city_info(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) CityInfo(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.CityInfo(&_YunguiContract.TransactOpts, ss)
}

// Fridge is a paid mutator transaction binding the contract method 0xc674e7f0.
//
// Solidity: function fridge(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) Fridge(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "fridge", ss)
}

// Fridge is a paid mutator transaction binding the contract method 0xc674e7f0.
//
// Solidity: function fridge(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) Fridge(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Fridge(&_YunguiContract.TransactOpts, ss)
}

// Fridge is a paid mutator transaction binding the contract method 0xc674e7f0.
//
// Solidity: function fridge(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) Fridge(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Fridge(&_YunguiContract.TransactOpts, ss)
}

// FridgeStorage is a paid mutator transaction binding the contract method 0x18c9480b.
//
// Solidity: function fridge_storage(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) FridgeStorage(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "fridge_storage", ss)
}

// FridgeStorage is a paid mutator transaction binding the contract method 0x18c9480b.
//
// Solidity: function fridge_storage(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) FridgeStorage(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.FridgeStorage(&_YunguiContract.TransactOpts, ss)
}

// FridgeStorage is a paid mutator transaction binding the contract method 0x18c9480b.
//
// Solidity: function fridge_storage(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) FridgeStorage(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.FridgeStorage(&_YunguiContract.TransactOpts, ss)
}

// Orders is a paid mutator transaction binding the contract method 0x1a948947.
//
// Solidity: function orders(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactor) Orders(opts *bind.TransactOpts, ss string) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "orders", ss)
}

// Orders is a paid mutator transaction binding the contract method 0x1a948947.
//
// Solidity: function orders(string ss) returns(uint256)
func (_YunguiContract *YunguiContractSession) Orders(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Orders(&_YunguiContract.TransactOpts, ss)
}

// Orders is a paid mutator transaction binding the contract method 0x1a948947.
//
// Solidity: function orders(string ss) returns(uint256)
func (_YunguiContract *YunguiContractTransactorSession) Orders(ss string) (*types.Transaction, error) {
	return _YunguiContract.Contract.Orders(&_YunguiContract.TransactOpts, ss)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YunguiContract *YunguiContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YunguiContract *YunguiContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _YunguiContract.Contract.RenounceOwnership(&_YunguiContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YunguiContract *YunguiContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YunguiContract.Contract.RenounceOwnership(&_YunguiContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_YunguiContract *YunguiContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _YunguiContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_YunguiContract *YunguiContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _YunguiContract.Contract.TransferOwnership(&_YunguiContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_YunguiContract *YunguiContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _YunguiContract.Contract.TransferOwnership(&_YunguiContract.TransactOpts, _newOwner)
}

// YunguiContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the YunguiContract contract.
type YunguiContractOwnershipRenouncedIterator struct {
	Event *YunguiContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *YunguiContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YunguiContractOwnershipRenounced)
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
		it.Event = new(YunguiContractOwnershipRenounced)
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
func (it *YunguiContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YunguiContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YunguiContractOwnershipRenounced represents a OwnershipRenounced event raised by the YunguiContract contract.
type YunguiContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_YunguiContract *YunguiContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*YunguiContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _YunguiContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YunguiContractOwnershipRenouncedIterator{contract: _YunguiContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_YunguiContract *YunguiContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *YunguiContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _YunguiContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YunguiContractOwnershipRenounced)
				if err := _YunguiContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// YunguiContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YunguiContract contract.
type YunguiContractOwnershipTransferredIterator struct {
	Event *YunguiContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YunguiContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YunguiContractOwnershipTransferred)
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
		it.Event = new(YunguiContractOwnershipTransferred)
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
func (it *YunguiContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YunguiContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YunguiContractOwnershipTransferred represents a OwnershipTransferred event raised by the YunguiContract contract.
type YunguiContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YunguiContract *YunguiContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YunguiContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YunguiContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YunguiContractOwnershipTransferredIterator{contract: _YunguiContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YunguiContract *YunguiContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YunguiContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YunguiContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YunguiContractOwnershipTransferred)
				if err := _YunguiContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
