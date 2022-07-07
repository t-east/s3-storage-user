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

// AuditAuditLog is an auto generated low-level Go binding around an user-defined struct.
type AuditAuditLog struct {
	Result bool
	Chal   uint32
	K1     string
	K2     string
	Myu    string
	Gamma  string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tpa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"AuditLogs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"K1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"K2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Gamma\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"GetAuditLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"K1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"K2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Gamma\",\"type\":\"string\"}],\"internalType\":\"structAudit.AuditLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_chal\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_k2\",\"type\":\"string\"}],\"name\":\"InitAuditLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_contentIDs\",\"type\":\"string[]\"}],\"name\":\"ListAuditLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"K1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"K2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Gamma\",\"type\":\"string\"}],\"internalType\":\"structAudit.AuditLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gamma\",\"type\":\"string\"}],\"name\":\"SetProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"}],\"name\":\"SetResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161133738038061133783398101604081905261002f9161008b565b600180546001600160a01b03199081166001600160a01b0394851617909155600380548216948416949094179093556002805490931691161790556100ce565b80516001600160a01b038116811461008657600080fd5b919050565b6000806000606084860312156100a057600080fd5b6100a98461006f565b92506100b76020850161006f565b91506100c56040850161006f565b90509250925092565b61125a806100dd6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e4ca57f1461006757806342eccaf71461007c57806346e8ecbb146100a557806368e9b6dc146100c557806389f18698146100ea578063ff164f94146100fd575b600080fd5b61007a610075366004610ed8565b610110565b005b61008f61008a366004610d4b565b6102ce565b60405161009c919061104d565b60405180910390f35b6100b86100b3366004610e13565b6105c2565b60405161009c919061111d565b6100d86100d3366004610e13565b61089a565b60405161009c969594939291906110af565b61007a6100f8366004610e50565b610b05565b61007a61010b366004610e13565b610bc1565b6002546001600160a01b031633146101615760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b60448201526064015b60405180910390fd5b826000856040516101729190611031565b908152602001604051809103902060000160016101000a81548163ffffffff021916908363ffffffff160217905550816000856040516101b29190611031565b908152602001604051809103902060010190805190602001906101d6929190610c42565b50806000856040516101e89190611031565b9081526020016040518091039020600201908051906020019061020c929190610c42565b506040518060200160405280600081525060008560405161022d9190611031565b90815260200160405180910390206004019080519060200190610251929190610c42565b50604051806020016040528060008152506000856040516102729190611031565b90815260200160405180910390206003019080519060200190610296929190610c42565b50600080856040516102a89190611031565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b8051606090819060005b818110156105b95760008582815181106102f4576102f46111f8565b60200260200101516040516103099190611031565b908152604080516020928190038301812060c082018352805460ff811615158352610100900463ffffffff169382019390935260018301805491939284019161035190611194565b80601f016020809104026020016040519081016040528092919081815260200182805461037d90611194565b80156103ca5780601f1061039f576101008083540402835291602001916103ca565b820191906000526020600020905b8154815290600101906020018083116103ad57829003601f168201915b505050505081526020016002820180546103e390611194565b80601f016020809104026020016040519081016040528092919081815260200182805461040f90611194565b801561045c5780601f106104315761010080835404028352916020019161045c565b820191906000526020600020905b81548152906001019060200180831161043f57829003601f168201915b5050505050815260200160038201805461047590611194565b80601f01602080910402602001604051908101604052809291908181526020018280546104a190611194565b80156104ee5780601f106104c3576101008083540402835291602001916104ee565b820191906000526020600020905b8154815290600101906020018083116104d157829003601f168201915b5050505050815260200160048201805461050790611194565b80601f016020809104026020016040519081016040528092919081815260200182805461053390611194565b80156105805780601f1061055557610100808354040283529160200191610580565b820191906000526020600020905b81548152906001019060200180831161056357829003601f168201915b50505050508152505083828151811061059b5761059b6111f8565b602002602001018190525080806105b1906111cf565b9150506102d8565b50909392505050565b6106036040518060c00160405280600015158152602001600063ffffffff168152602001606081526020016060815260200160608152602001606081525090565b6000826040516106139190611031565b908152604080516020928190038301812060c082018352805460ff811615158352610100900463ffffffff169382019390935260018301805491939284019161065b90611194565b80601f016020809104026020016040519081016040528092919081815260200182805461068790611194565b80156106d45780601f106106a9576101008083540402835291602001916106d4565b820191906000526020600020905b8154815290600101906020018083116106b757829003601f168201915b505050505081526020016002820180546106ed90611194565b80601f016020809104026020016040519081016040528092919081815260200182805461071990611194565b80156107665780601f1061073b57610100808354040283529160200191610766565b820191906000526020600020905b81548152906001019060200180831161074957829003601f168201915b5050505050815260200160038201805461077f90611194565b80601f01602080910402602001604051908101604052809291908181526020018280546107ab90611194565b80156107f85780601f106107cd576101008083540402835291602001916107f8565b820191906000526020600020905b8154815290600101906020018083116107db57829003601f168201915b5050505050815260200160048201805461081190611194565b80601f016020809104026020016040519081016040528092919081815260200182805461083d90611194565b801561088a5780601f1061085f5761010080835404028352916020019161088a565b820191906000526020600020905b81548152906001019060200180831161086d57829003601f168201915b5050505050815250509050919050565b80516020818301810180516000825292820191909301209152805460018201805460ff83169361010090930463ffffffff169291906108d890611194565b80601f016020809104026020016040519081016040528092919081815260200182805461090490611194565b80156109515780601f1061092657610100808354040283529160200191610951565b820191906000526020600020905b81548152906001019060200180831161093457829003601f168201915b50505050509080600201805461096690611194565b80601f016020809104026020016040519081016040528092919081815260200182805461099290611194565b80156109df5780601f106109b4576101008083540402835291602001916109df565b820191906000526020600020905b8154815290600101906020018083116109c257829003601f168201915b5050505050908060030180546109f490611194565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2090611194565b8015610a6d5780601f10610a4257610100808354040283529160200191610a6d565b820191906000526020600020905b815481529060010190602001808311610a5057829003601f168201915b505050505090806004018054610a8290611194565b80601f0160208091040260200160405190810160405280929190818152602001828054610aae90611194565b8015610afb5780601f10610ad057610100808354040283529160200191610afb565b820191906000526020600020905b815481529060010190602001808311610ade57829003601f168201915b5050505050905086565b6001546001600160a01b03163314610b505760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b6044820152606401610158565b80600084604051610b619190611031565b90815260200160405180910390206004019080519060200190610b85929190610c42565b5081600084604051610b979190611031565b90815260200160405180910390206003019080519060200190610bbb929190610c42565b50505050565b6002546001600160a01b03163314610c0d5760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b6044820152606401610158565b6001600082604051610c1f9190611031565b908152604051908190036020019020805491151560ff1990921691909117905550565b828054610c4e90611194565b90600052602060002090601f016020900481019282610c705760008555610cb6565b82601f10610c8957805160ff1916838001178555610cb6565b82800160010185558215610cb6579182015b82811115610cb6578251825591602001919060010190610c9b565b50610cc2929150610cc6565b5090565b5b80821115610cc25760008155600101610cc7565b600082601f830112610cec57600080fd5b813567ffffffffffffffff811115610d0657610d0661120e565b610d19601f8201601f1916602001611137565b818152846020838601011115610d2e57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020808385031215610d5e57600080fd5b823567ffffffffffffffff80821115610d7657600080fd5b818501915085601f830112610d8a57600080fd5b813581811115610d9c57610d9c61120e565b8060051b610dab858201611137565b8281528581019085870183870188018b1015610dc657600080fd5b60009350835b85811015610e0357813587811115610de2578586fd5b610df08d8b838c0101610cdb565b8552509288019290880190600101610dcc565b50909a9950505050505050505050565b600060208284031215610e2557600080fd5b813567ffffffffffffffff811115610e3c57600080fd5b610e4884828501610cdb565b949350505050565b600080600060608486031215610e6557600080fd5b833567ffffffffffffffff80821115610e7d57600080fd5b610e8987838801610cdb565b94506020860135915080821115610e9f57600080fd5b610eab87838801610cdb565b93506040860135915080821115610ec157600080fd5b50610ece86828701610cdb565b9150509250925092565b60008060008060808587031215610eee57600080fd5b843567ffffffffffffffff80821115610f0657600080fd5b610f1288838901610cdb565b95506020870135915063ffffffff82168214610f2d57600080fd5b90935060408601359080821115610f4357600080fd5b610f4f88838901610cdb565b93506060870135915080821115610f6557600080fd5b50610f7287828801610cdb565b91505092959194509250565b60008151808452610f96816020860160208601611168565b601f01601f19169290920160200192915050565b80511515825263ffffffff60208201511660208301526000604082015160c06040850152610fdb60c0850182610f7e565b905060608301518482036060860152610ff48282610f7e565b9150506080830151848203608086015261100e8282610f7e565b91505060a083015184820360a08601526110288282610f7e565b95945050505050565b60008251611043818460208701611168565b9190910192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156110a257603f19888603018452611090858351610faa565b94509285019290850190600101611074565b5092979650505050505050565b861515815263ffffffff8616602082015260c0604082015260006110d660c0830187610f7e565b82810360608401526110e88187610f7e565b905082810360808401526110fc8186610f7e565b905082810360a08401526111108185610f7e565b9998505050505050505050565b6020815260006111306020830184610faa565b9392505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156111605761116061120e565b604052919050565b60005b8381101561118357818101518382015260200161116b565b83811115610bbb5750506000910152565b600181811c908216806111a857607f821691505b602082108114156111c957634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156111f157634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212202215a385abdf00f9576e95b1a1799b2a89316cfd777bbee94e84cb5491b4d35664736f6c63430008070033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _user common.Address, _spa common.Address, _tpa common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _user, _spa, _tpa)
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

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, string K1, string K2, string Myu, string Gamma)
func (_Contracts *ContractsCaller) AuditLogs(opts *bind.CallOpts, arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "AuditLogs", arg0)

	outstruct := new(struct {
		Result bool
		Chal   uint32
		K1     string
		K2     string
		Myu    string
		Gamma  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Chal = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.K1 = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.K2 = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Myu = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Gamma = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, string K1, string K2, string Myu, string Gamma)
func (_Contracts *ContractsSession) AuditLogs(arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	return _Contracts.Contract.AuditLogs(&_Contracts.CallOpts, arg0)
}

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, string K1, string K2, string Myu, string Gamma)
func (_Contracts *ContractsCallerSession) AuditLogs(arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	return _Contracts.Contract.AuditLogs(&_Contracts.CallOpts, arg0)
}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,string,string,string,string))
func (_Contracts *ContractsCaller) GetAuditLog(opts *bind.CallOpts, _contentId string) (AuditAuditLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetAuditLog", _contentId)

	if err != nil {
		return *new(AuditAuditLog), err
	}

	out0 := *abi.ConvertType(out[0], new(AuditAuditLog)).(*AuditAuditLog)

	return out0, err

}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,string,string,string,string))
func (_Contracts *ContractsSession) GetAuditLog(_contentId string) (AuditAuditLog, error) {
	return _Contracts.Contract.GetAuditLog(&_Contracts.CallOpts, _contentId)
}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,string,string,string,string))
func (_Contracts *ContractsCallerSession) GetAuditLog(_contentId string) (AuditAuditLog, error) {
	return _Contracts.Contract.GetAuditLog(&_Contracts.CallOpts, _contentId)
}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,string,string,string,string)[])
func (_Contracts *ContractsCaller) ListAuditLog(opts *bind.CallOpts, _contentIDs []string) ([]AuditAuditLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListAuditLog", _contentIDs)

	if err != nil {
		return *new([]AuditAuditLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuditAuditLog)).(*[]AuditAuditLog)

	return out0, err

}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,string,string,string,string)[])
func (_Contracts *ContractsSession) ListAuditLog(_contentIDs []string) ([]AuditAuditLog, error) {
	return _Contracts.Contract.ListAuditLog(&_Contracts.CallOpts, _contentIDs)
}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,string,string,string,string)[])
func (_Contracts *ContractsCallerSession) ListAuditLog(_contentIDs []string) ([]AuditAuditLog, error) {
	return _Contracts.Contract.ListAuditLog(&_Contracts.CallOpts, _contentIDs)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0x0e4ca57f.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, string _k1, string _k2) returns()
func (_Contracts *ContractsTransactor) InitAuditLog(opts *bind.TransactOpts, _contentID string, _chal uint32, _k1 string, _k2 string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "InitAuditLog", _contentID, _chal, _k1, _k2)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0x0e4ca57f.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, string _k1, string _k2) returns()
func (_Contracts *ContractsSession) InitAuditLog(_contentID string, _chal uint32, _k1 string, _k2 string) (*types.Transaction, error) {
	return _Contracts.Contract.InitAuditLog(&_Contracts.TransactOpts, _contentID, _chal, _k1, _k2)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0x0e4ca57f.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, string _k1, string _k2) returns()
func (_Contracts *ContractsTransactorSession) InitAuditLog(_contentID string, _chal uint32, _k1 string, _k2 string) (*types.Transaction, error) {
	return _Contracts.Contract.InitAuditLog(&_Contracts.TransactOpts, _contentID, _chal, _k1, _k2)
}

// SetProof is a paid mutator transaction binding the contract method 0x89f18698.
//
// Solidity: function SetProof(string _contentID, string _myu, string _gamma) returns()
func (_Contracts *ContractsTransactor) SetProof(opts *bind.TransactOpts, _contentID string, _myu string, _gamma string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetProof", _contentID, _myu, _gamma)
}

// SetProof is a paid mutator transaction binding the contract method 0x89f18698.
//
// Solidity: function SetProof(string _contentID, string _myu, string _gamma) returns()
func (_Contracts *ContractsSession) SetProof(_contentID string, _myu string, _gamma string) (*types.Transaction, error) {
	return _Contracts.Contract.SetProof(&_Contracts.TransactOpts, _contentID, _myu, _gamma)
}

// SetProof is a paid mutator transaction binding the contract method 0x89f18698.
//
// Solidity: function SetProof(string _contentID, string _myu, string _gamma) returns()
func (_Contracts *ContractsTransactorSession) SetProof(_contentID string, _myu string, _gamma string) (*types.Transaction, error) {
	return _Contracts.Contract.SetProof(&_Contracts.TransactOpts, _contentID, _myu, _gamma)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsTransactor) SetResult(opts *bind.TransactOpts, _contentID string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetResult", _contentID)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsSession) SetResult(_contentID string) (*types.Transaction, error) {
	return _Contracts.Contract.SetResult(&_Contracts.TransactOpts, _contentID)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsTransactorSession) SetResult(_contentID string) (*types.Transaction, error) {
	return _Contracts.Contract.SetResult(&_Contracts.TransactOpts, _contentID)
}
