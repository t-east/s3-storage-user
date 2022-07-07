// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// ContentContentLog is an auto generated low-level Go binding around an user-defined struct.
type ContentContentLog struct {
	Owner    common.Address
	Hash     []string
	Provider common.Address
	LogId    string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ContentLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"GetContentLog\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"Hash\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"internalType\":\"structContent.ContentLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"InitContentLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ListContentID\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ListContentLog\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"Hash\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"internalType\":\"structContent.ContentLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_hash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"SetContentLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610fce380380610fce83398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610f3b806100936000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063242d3b4b1461006757806366754e451461009057806390ca908a146100b25780639a4790f5146100c7578063a33a9072146100dc578063d2056ab9146100ef575b600080fd5b61007a610075366004610b4f565b610104565b6040516100879190610dfe565b60405180910390f35b6100a361009e366004610b4f565b6102e4565b60405161008793929190610d1b565b6100c56100c0366004610a54565b6103a8565b005b6100cf6104fc565b6040516100879190610da9565b6100c56100ea366004610b4f565b610719565b6100f76107a0565b6040516100879190610d47565b60408051608081018252600080825260606020830181905292820152818101919091526000826040516101379190610c63565b90815260408051918290036020908101832060808401835280546001600160a01b0316845260018101805484518185028101850190955280855291938584019390929060009084015b8282101561022c57838290600052602060002001805461019f90610e75565b80601f01602080910402602001604051908101604052809291908181526020018280546101cb90610e75565b80156102185780601f106101ed57610100808354040283529160200191610218565b820191906000526020600020905b8154815290600101906020018083116101fb57829003601f168201915b505050505081526020019060010190610180565b5050509082525060028201546001600160a01b0316602082015260038201805460409092019161025b90610e75565b80601f016020809104026020016040519081016040528092919081815260200182805461028790610e75565b80156102d45780601f106102a9576101008083540402835291602001916102d4565b820191906000526020600020905b8154815290600101906020018083116102b757829003601f168201915b5050505050815250509050919050565b8051808201602090810180516000825292820191909301209152805460028201546003830180546001600160a01b0393841694929093169261032590610e75565b80601f016020809104026020016040519081016040528092919081815260200182805461035190610e75565b801561039e5780601f106103735761010080835404028352916020019161039e565b820191906000526020600020905b81548152906001019060200180831161038157829003601f168201915b5050505050905083565b6001546001600160a01b031633146103f75760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b604482015260640160405180910390fd5b336000836040516104089190610c63565b908152602001604051809103902060020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508260008360405161044e9190610c63565b9081526020016040518091039020600101908051906020019061047292919061087c565b50806000836040516104849190610c63565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b0393909316929092179091556004805460018101825560009190915283516104f6927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201918501906108d9565b50505050565b600454606090819060005b818110156107115760006004828154811061052457610524610ed9565b9060005260206000200160405161053b9190610c7f565b90815260408051918290036020908101832060808401835280546001600160a01b0316845260018101805484518185028101850190955280855291938584019390929060009084015b828210156106305783829060005260206000200180546105a390610e75565b80601f01602080910402602001604051908101604052809291908181526020018280546105cf90610e75565b801561061c5780601f106105f15761010080835404028352916020019161061c565b820191906000526020600020905b8154815290600101906020018083116105ff57829003601f168201915b505050505081526020019060010190610584565b5050509082525060028201546001600160a01b0316602082015260038201805460409092019161065f90610e75565b80601f016020809104026020016040519081016040528092919081815260200182805461068b90610e75565b80156106d85780601f106106ad576101008083540402835291602001916106d8565b820191906000526020600020905b8154815290600101906020018083116106bb57829003601f168201915b5050505050815250508382815181106106f3576106f3610ed9565b6020026020010181905250808061070990610eb0565b915050610507565b509092915050565b3360008260405161072a9190610c63565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b03939093169290921790915560048054600181018255600091909152825161079c927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201918401906108d9565b5050565b6060806004805480602002602001604051908101604052809291908181526020016000905b828210156108715783829060005260206000200180546107e490610e75565b80601f016020809104026020016040519081016040528092919081815260200182805461081090610e75565b801561085d5780601f106108325761010080835404028352916020019161085d565b820191906000526020600020905b81548152906001019060200180831161084057829003601f168201915b5050505050815260200190600101906107c5565b509295945050505050565b8280548282559060005260206000209081019282156108c9579160200282015b828111156108c957825180516108b99184916020909101906108d9565b509160200191906001019061089c565b506108d5929150610959565b5090565b8280546108e590610e75565b90600052602060002090601f016020900481019282610907576000855561094d565b82601f1061092057805160ff191683800117855561094d565b8280016001018555821561094d579182015b8281111561094d578251825591602001919060010190610932565b506108d5929150610976565b808211156108d557600061096d828261098b565b50600101610959565b5b808211156108d55760008155600101610977565b50805461099790610e75565b6000825580601f106109a7575050565b601f0160209004906000526020600020908101906109c59190610976565b50565b80356001600160a01b03811681146109df57600080fd5b919050565b600082601f8301126109f557600080fd5b813567ffffffffffffffff811115610a0f57610a0f610eef565b610a22601f8201601f1916602001610e18565b818152846020838601011115610a3757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215610a6957600080fd5b833567ffffffffffffffff80821115610a8157600080fd5b818601915086601f830112610a9557600080fd5b8135602082821115610aa957610aa9610eef565b8160051b610ab8828201610e18565b8381528281019086840183880185018d1015610ad357600080fd5b600093505b85841015610b1257803587811115610aef57600080fd5b610afd8e87838c01016109e4565b84525060019390930192918401918401610ad8565b509850505087013592505080821115610b2a57600080fd5b50610b37868287016109e4565b925050610b46604085016109c8565b90509250925092565b600060208284031215610b6157600080fd5b813567ffffffffffffffff811115610b7857600080fd5b610b84848285016109e4565b949350505050565b60008151808452610ba4816020860160208601610e49565b601f01601f19169290920160200192915050565b60006080830160018060a01b03835116845260208084015160808287015282815180855260a08801915060a08160051b8901019450838301925060005b81811015610c2357609f19898703018352610c11868551610b8c565b95509284019291840191600101610bf5565b50505050506040830151610c4260408601826001600160a01b03169052565b5060608301518482036060860152610c5a8282610b8c565b95945050505050565b60008251610c75818460208701610e49565b9190910192915050565b600080835481600182811c915080831680610c9b57607f831692505b6020808410821415610cbb57634e487b7160e01b86526022600452602486fd5b818015610ccf5760018114610ce057610d0d565b60ff19861689528489019650610d0d565b60008a81526020902060005b86811015610d055781548b820152908501908301610cec565b505084890196505b509498975050505050505050565b6001600160a01b03848116825283166020820152606060408201819052600090610c5a90830184610b8c565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610d9c57603f19888603018452610d8a858351610b8c565b94509285019290850190600101610d6e565b5092979650505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610d9c57603f19888603018452610dec858351610bb8565b94509285019290850190600101610dd0565b602081526000610e116020830184610bb8565b9392505050565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e4157610e41610eef565b604052919050565b60005b83811015610e64578181015183820152602001610e4c565b838111156104f65750506000910152565b600181811c90821680610e8957607f821691505b60208210811415610eaa57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610ed257634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212207652b622a4f26e6092bd651820b354f12ab3d531b3a893650b94c5d82922aa6664736f6c63430008070033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _spa common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _spa)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsCaller) ContentLogs(opts *bind.CallOpts, arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ContentLogs", arg0)

	outstruct := new(struct {
		Owner    common.Address
		Provider common.Address
		LogId    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LogId = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsSession) ContentLogs(arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	return _Contracts.Contract.ContentLogs(&_Contracts.CallOpts, arg0)
}

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsCallerSession) ContentLogs(arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	return _Contracts.Contract.ContentLogs(&_Contracts.CallOpts, arg0)
}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsCaller) GetContentLog(opts *bind.CallOpts, _contentId string) (ContentContentLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetContentLog", _contentId)

	if err != nil {
		return *new(ContentContentLog), err
	}

	out0 := *abi.ConvertType(out[0], new(ContentContentLog)).(*ContentContentLog)

	return out0, err

}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsCallerSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsCaller) ListContentID(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListContentID")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsSession) ListContentID() ([]string, error) {
	return _Contracts.Contract.ListContentID(&_Contracts.CallOpts)
}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsCallerSession) ListContentID() ([]string, error) {
	return _Contracts.Contract.ListContentID(&_Contracts.CallOpts)
}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,string[],address,string)[])
func (_Contracts *ContractsCaller) ListContentLog(opts *bind.CallOpts) ([]ContentContentLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListContentLog")

	if err != nil {
		return *new([]ContentContentLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContentContentLog)).(*[]ContentContentLog)

	return out0, err

}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,string[],address,string)[])
func (_Contracts *ContractsSession) ListContentLog() ([]ContentContentLog, error) {
	return _Contracts.Contract.ListContentLog(&_Contracts.CallOpts)
}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,string[],address,string)[])
func (_Contracts *ContractsCallerSession) ListContentLog() ([]ContentContentLog, error) {
	return _Contracts.Contract.ListContentLog(&_Contracts.CallOpts)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsTransactor) InitContentLog(opts *bind.TransactOpts, _contentId string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "InitContentLog", _contentId)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsSession) InitContentLog(_contentId string) (*types.Transaction, error) {
	return _Contracts.Contract.InitContentLog(&_Contracts.TransactOpts, _contentId)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsTransactorSession) InitContentLog(_contentId string) (*types.Transaction, error) {
	return _Contracts.Contract.InitContentLog(&_Contracts.TransactOpts, _contentId)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) SetContentLog(opts *bind.TransactOpts, _hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetContentLog", _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsSession) SetContentLog(_hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) SetContentLog(_hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}
