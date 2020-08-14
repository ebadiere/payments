/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalHermesStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_channelImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_parentAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channelAddress\",\"type\":\"address\"}],\"name\":\"ConsumerChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newURL\",\"type\":\"bytes\"}],\"name\":\"HermesURLUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hermesOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ur\",\"type\":\"bytes\"}],\"name\":\"RegisteredHermes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"}],\"name\":\"RegisteredIdentity\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hermeses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"function()viewexternalreturns(uint256)\",\"name\":\"stake\",\"type\":\"function\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalHermesStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hermesStake\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_hermesFee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minChannelStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxChannelStake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_url\",\"type\":\"bytes\"}],\"name\":\"registerHermes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesOperator\",\"type\":\"address\"}],\"name\":\"getHermesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"getHermesURL\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_url\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateHermsURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"getProxyCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChannelImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHermesImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_parentAddress\",\"type\":\"address\"}],\"name\":\"hasParentRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"isHermes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405234801561001057600080fd5b506040516125c43803806125c4833981810160405260c081101561003357600080fd5b508051602082015160408301516060840151608085015160a090950151600483905593949293919290916001600160a01b03861661007057600080fd5b600280546001600160a01b0319166001600160a01b0388811691909117909155851661009b57600080fd5b600380546001600160a01b039687166001600160a01b03199182161790915560058054948716948216949094179093556006805492861692841692909217909155600780549190941691161790915550506124c9806100fb6000396000f3fe60806040526004361061014f5760003560e01c8063acc831d0116100b6578063df8de3e71161006f578063df8de3e714610821578063e325253714610854578063e617aaac14610887578063f2fde38b146108c2578063f58c5b6e146108f5578063fc0c546a1461090a576101a1565b8063acc831d0146105a3578063bf1eb88a146105d6578063c3c5a54714610609578063cdd596e01461063c578063cf10c9691461066f578063d5929fe314610744576101a1565b80636931b550116101085780636931b55014610492578063715018a6146104a75780637c671a21146104bc5780638da5cb5b146104d15780639936a87b146104e6578063ab867213146104fb576101a1565b806303fb422f146101a6578063238e130a14610278578063332ff7bd146102ad57806366330f56146102f457806366cf58751461043a578063692058c214610461576101a1565b366101a1576040805162461bcd60e51b815260206004820152601d60248201527f52656a656374696e672074782077697468206574686572732073656e74000000604482015290519081900360640190fd5b600080fd5b3480156101b257600080fd5b506101d9600480360360208110156101c957600080fd5b50356001600160a01b031661091f565b60405180856001600160a01b03168152602001848463ffffffff169060201b1760401b815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561023a578181015183820152602001610222565b50505050905090810190601f1680156102675780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561028457600080fd5b506102ab6004803603602081101561029b57600080fd5b50356001600160a01b03166109e3565b005b3480156102b957600080fd5b506102e0600480360360208110156102d057600080fd5b50356001600160a01b0316610ab3565b604080519115158252519081900360200190f35b34801561030057600080fd5b506102ab6004803603606081101561031757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561034157600080fd5b82018360208201111561035357600080fd5b803590602001918460018302840111600160201b8311171561037457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103c657600080fd5b8201836020820111156103d857600080fd5b803590602001918460018302840111600160201b831117156103f957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ac5945050505050565b34801561044657600080fd5b5061044f610d08565b60408051918252519081900360200190f35b34801561046d57600080fd5b50610476610d0e565b604080516001600160a01b039092168252519081900360200190f35b34801561049e57600080fd5b506102ab610d1d565b3480156104b357600080fd5b506102ab610d6e565b3480156104c857600080fd5b50610476610e19565b3480156104dd57600080fd5b50610476610e28565b3480156104f257600080fd5b50610476610e37565b34801561050757600080fd5b5061052e6004803603602081101561051e57600080fd5b50356001600160a01b0316610e46565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610568578181015183820152602001610550565b50505050905090810190601f1680156105955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105af57600080fd5b50610476600480360360208110156105c657600080fd5b50356001600160a01b0316610ec5565b3480156105e257600080fd5b5061052e600480360360208110156105f957600080fd5b50356001600160a01b0316610efc565b34801561061557600080fd5b506102e06004803603602081101561062c57600080fd5b50356001600160a01b0316610fad565b34801561064857600080fd5b506102e06004803603602081101561065f57600080fd5b50356001600160a01b0316611072565b34801561067b57600080fd5b506102ab600480360360a081101561069257600080fd5b6001600160a01b038235811692602081013592604082013592606083013516919081019060a081016080820135600160201b8111156106d057600080fd5b8201836020820111156106e257600080fd5b803590602001918460018302840111600160201b8311171561070357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061114a945050505050565b34801561075057600080fd5b506102ab600480360360c081101561076757600080fd5b6001600160a01b038235169160208101359161ffff604083013516916060810135916080820135919081019060c0810160a0820135600160201b8111156107ad57600080fd5b8201836020820111156107bf57600080fd5b803590602001918460018302840111600160201b831117156107e057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116a9945050505050565b34801561082d57600080fd5b506102ab6004803603602081101561084457600080fd5b50356001600160a01b0316611a48565b34801561086057600080fd5b506102ab6004803603602081101561087757600080fd5b50356001600160a01b0316611bae565b34801561089357600080fd5b50610476600480360360408110156108aa57600080fd5b506001600160a01b0381358116916020013516611d1a565b3480156108ce57600080fd5b506102ab600480360360208110156108e557600080fd5b50356001600160a01b0316611d81565b34801561090157600080fd5b50610476611e82565b34801561091657600080fd5b50610476611e91565b60086020908152600091825260409182902080546001808301546002808501805488516101009582161595909502600019011691909104601f81018790048702840187019097528683526001600160a01b039384169682871c9094169563ffffffff909216949390918301828280156109d95780601f106109ae576101008083540402835291602001916109d9565b820191906000526020600020905b8154815290600101906020018083116109bc57829003601f168201915b5050505050905084565b6000546001600160a01b0316331480610a0557506000546001600160a01b0316155b610a44576040805162461bcd60e51b81526020600482018190526024820152600080516020612424833981519152604482015290519081900360640190fd5b6001600160a01b038116610a5757600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03811615155b919050565b610ace83611ea0565b610b1f576040805162461bcd60e51b815260206004820181905260248201527f70726f7669646564206865726d65732068617320746f20626520616374697665604482015290519081900360640190fd5b6000610bc98230868660405160200180846001600160a01b031660601b8152601401836001600160a01b031660601b815260140182805190602001908083835b60208310610b7e5780518252601f199092019160209182019101610b5f565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120611fc590919063ffffffff16565b6001600160a01b03808616600090815260086020526040902054919250808316911614610c2f576040805162461bcd60e51b815260206004820152600f60248201526e77726f6e67207369676e617475726560881b604482015290519081900360640190fd5b6001600160a01b03841660009081526008602090815260409091208451610c5e9260029092019186019061226b565b50836001600160a01b03167fd8c638c85547b8717e0d5ca292cff6dbe8fc02fa6e6863a047971c39511643c7846040518080602001828103825283818151815260200191508051906020019080838360005b83811015610cc8578181015183820152602001610cb0565b50505050905090810190601f168015610cf55780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b60045481565b6003546001600160a01b031681565b6001546001600160a01b0316610d3257600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610d6b573d6000803e3d6000fd5b50565b6000546001600160a01b0316331480610d9057506000546001600160a01b0316155b610dcf576040805162461bcd60e51b81526020600482018190526024820152600080516020612424833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6005546001600160a01b031690565b6000546001600160a01b031690565b6006546001600160a01b031690565b6060806040518060600160405280603781526020016123cb603791399050606083901b60005b60148160ff161015610ebc57818160ff1660148110610e8757fe5b1a60f81b838260140160ff1681518110610e9d57fe5b60200101906001600160f81b031916908160001a905350600101610e6c565b50909392505050565b600080610ed8610ed3610e37565b610e46565b80516020909101209050610ef56001600160a01b038416826121b0565b9392505050565b6001600160a01b0381166000908152600860209081526040918290206002908101805484516001821615610100026000190190911692909204601f81018490048402830184019094528382526060939192909190830182828015610fa15780601f10610f7657610100808354040283529160200191610fa1565b820191906000526020600020905b815481529060010190602001808311610f8457829003601f168201915b50505050509050919050565b600754600090610fc5906001600160a01b0316610ab3565b801561104657506007546040805163c3c5a54760e01b81526001600160a01b0385811660048301529151919092169163c3c5a547916024808301926020929190829003018186803b15801561101957600080fd5b505afa15801561102d573d6000803e3d6000fd5b505050506040513d602081101561104357600080fd5b50515b1561105357506001610ac0565b506001600160a01b031660009081526009602052604090205460ff1690565b60075460009061108a906001600160a01b0316610ab3565b801561110b575060075460408051632a33ddbd60e01b81526001600160a01b03858116600483015291519190921691632a33ddbd916024808301926020929190829003018186803b1580156110de57600080fd5b505afa1580156110f2573d6000803e3d6000fd5b505050506040513d602081101561110857600080fd5b50515b1561111857506001610ac0565b6001600160a01b038083166000908152600860205260408120549091169061113f82610ec5565b3b1515949350505050565b61115385611ea0565b6111a4576040805162461bcd60e51b815260206004820152601e60248201527f70726f766964656420686173206861766520746f206265206163746976650000604482015290519081900360640190fd5b6040805130606090811b6020808401919091526001600160601b031989831b8116603485015260488401899052606884018890529186901b90911660888301528251607c818403018152609c90920190925280519101206000906112089083611fc5565b90506001600160a01b038116611257576040805162461bcd60e51b815260206004820152600f60248201526e77726f6e67207369676e617475726560881b604482015290519081900360640190fd5b600061126386866121f7565b6002549091506001600160a01b03166370a08231611281848a611d1a565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156112be57600080fd5b505afa1580156112d2573d6000803e3d6000fd5b505050506040513d60208110156112e857600080fd5b50518111156113285760405162461bcd60e51b815260040180806020018281038252602a815260200180612444602a913960400191505060405180910390fd5b604080516001600160601b0319606085811b82166020808501919091528b821b9092166034840152835160288185030181526048909301909352815191012090611373610ed3610e19565b905060006113818383612251565b60025460035460408051637b809f7b60e11b81526001600160a01b039384166004820152918316602483015288831660448301528d8316606483015260848201889052519293509083169163f7013ef69160a48082019260009290919082900301818387803b1580156113f357600080fd5b505af1158015611407573d6000803e3d6000fd5b5050505060008911801561142357506001600160a01b03871615155b15611566576002546040805163095ea7b360e01b81526001600160a01b038d81166004830152602482018d90529151919092169163095ea7b39160448083019260209291908290030181600087803b15801561147e57600080fd5b505af1158015611492573d6000803e3d6000fd5b505050506040513d60208110156114a857600080fd5b50516114e55760405162461bcd60e51b815260040180806020018281038252602d815260200180612325602d913960400191505060405180910390fd5b896001600160a01b0316630a798f2486898c6040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b031681526020018281526020019350505050600060405180830381600087803b15801561154d57600080fd5b505af1158015611561573d6000803e3d6000fd5b505050505b87156115ed576002546040805163a9059cbb60e01b8152336004820152602481018b905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b1580156115c057600080fd5b505af11580156115d4573d6000803e3d6000fd5b505050506040513d60208110156115ea57600080fd5b50505b896001600160a01b0316856001600160a01b03167f2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b8360405180826001600160a01b0316815260200191505060405180910390a361164a85610fad565b61169d576001600160a01b03808616600081815260096020526040808220805460ff1916600117905551928d16927fefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef9190a35b50505050505050505050565b6001600160a01b038616611704576040805162461bcd60e51b815260206004820152601e60248201527f6f70657261746f722063616e2774206265207a65726f20616464726573730000604482015290519081900360640190fd5b6004548510156117455760405162461bcd60e51b81526004018080602001828103825260328152602001806123776032913960400191505060405180910390fd5b600061175087610ec5565b905061175b81611072565b156117ad576040805162461bcd60e51b815260206004820152601960248201527f6865726d657320616c7265616479207265676973746572656400000000000000604482015290519081900360640190fd5b60006117cc886001600160a01b03166117c7610ed3610e37565b612251565b600254604080516323b872dd60e01b81523360048201526001600160a01b038085166024830152604482018c905291519394509116916323b872dd916064808201926020929091908290030181600087803b15801561182a57600080fd5b505af115801561183e573d6000803e3d6000fd5b505050506040513d602081101561185457600080fd5b5050600254604080516335d0375160e21b81526001600160a01b0392831660048201528a8316602482015261ffff89166044820152606481018890526084810187905290519183169163d740dd449160a48082019260009290919082900301818387803b1580156118c457600080fd5b505af11580156118d8573d6000803e3d6000fd5b5050604080516060810182526001600160a01b038c8116825263fc0e3d90602087811b640100000000600160c01b0390811692909217851b8185019081528486018b8152898516600090815260088452879020865181546001600160a01b031916961695909517855590516001850180546001600160c01b03191663ffffffff9290981c9182169190941617959095179091559251805192955090935061198692600285019291019061226b565b50905050806001600160a01b03167ff06d60cc2f463635fd237ad87f1d007af54840b82e7e4561707b1be63d91c260898560405180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611a035781810151838201526020016119eb565b50505050905090810190601f168015611a305780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505050505050565b6001546001600160a01b0316611a5d57600080fd5b6002546001600160a01b0382811691161415611aaa5760405162461bcd60e51b81526004018080602001828103825260258152602001806123526025913960400191505060405180910390fd5b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611af957600080fd5b505afa158015611b0d573d6000803e3d6000fd5b505050506040513d6020811015611b2357600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b158015611b7e57600080fd5b505af1158015611b92573d6000803e3d6000fd5b505050506040513d6020811015611ba857600080fd5b50505050565b6000546001600160a01b0316331480611bd057506000546001600160a01b0316155b611c0f576040805162461bcd60e51b81526020600482018190526024820152600080516020612424833981519152604482015290519081900360640190fd5b600254604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015611c5a57600080fd5b505afa158015611c6e573d6000803e3d6000fd5b505050506040513d6020811015611c8457600080fd5b5051905080611cc45760405162461bcd60e51b815260040180806020018281038252602681526020018061246e6026913960400191505060405180910390fd5b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015611b7e57600080fd5b600080611d28610ed3610e19565b8051602091820120604080516001600160601b0319606089811b82168387015288901b1660348201528151602881830301815260489091019091528051920191909120909150611d7881836121b0565b95945050505050565b6000546001600160a01b0316331480611da357506000546001600160a01b0316155b611de2576040805162461bcd60e51b81526020600482018190526024820152600080516020612424833981519152604482015290519081900360640190fd5b6001600160a01b038116611e275760405162461bcd60e51b81526004018080602001828103825260268152602001806122ff6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031690565b6002546001600160a01b031681565b600754600090611eb8906001600160a01b0316610ab3565b8015611f39575060075460408051631a3d9a5960e01b81526001600160a01b03858116600483015291519190921691631a3d9a59916024808301926020929190829003018186803b158015611f0c57600080fd5b505afa158015611f20573d6000803e3d6000fd5b505050506040513d6020811015611f3657600080fd5b50515b15611f4657506001610ac0565b6000826001600160a01b0316634e69d5606040518163ffffffff1660e01b815260040160206040518083038186803b158015611f8157600080fd5b505afa158015611f95573d6000803e3d6000fd5b505050506040513d6020811015611fab57600080fd5b505190506000816003811115611fbd57fe5b149392505050565b6000815160411461201d576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561208e5760405162461bcd60e51b81526004018080602001828103825260228152602001806123a96022913960400191505060405180910390fd5b8060ff16601b141580156120a657508060ff16601c14155b156120e25760405162461bcd60e51b81526004018080602001828103825260228152602001806124026022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561213e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166121a6576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b604080516001600160f81b03196020808301919091523060601b60218301526035820194909452605580820193909352815180820390930183526075019052805191012090565b600082820183811015610ef5576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600080838351602085016000f59050803b610ef557600080fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106122ac57805160ff19168380011785556122d9565b828001600101855582156122d9579182015b828111156122d95782518255916020019190600101906122be565b506122e59291506122e9565b5090565b5b808211156122e557600081556001016122ea56fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736865726d65732073686f756c642067657420617070726f76616c20746f207472616e7366657220746f6b656e736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265646865726d6573206861766520746f207374616b65206174206c65617374206d696e696d616c207374616b6520616d6f756e7445434453413a20696e76616c6964207369676e6174757265202773272076616c75653d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf345434453413a20696e76616c6964207369676e6174757265202776272076616c75654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726e6f7420656e6f756768742066756e647320696e206368616e6e656c20746f20636f7665722066656573636f6c6c6563746564206665652063616e6e6f74206265206c657373207468616e207a65726fa26469706673582212203499cfdd5cf7516ab3ed7cd63a869158319f2c473bde81bfe0db3bb32904f8f664736f6c634300060c0033"

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _minimalHermesStake *big.Int, _channelImplementation common.Address, _hermesImplementation common.Address, _parentAddress common.Address) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, _tokenAddress, _dexAddress, _minimalHermesStake, _channelImplementation, _hermesImplementation, _parentAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Registry *RegistryCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "dex")
	return *ret0, err
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Registry *RegistrySession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Registry *RegistryCallerSession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
func (_Registry *RegistryCaller) GetChannelAddress(opts *bind.CallOpts, _identity common.Address, _hermesId common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelAddress", _identity, _hermesId)
	return *ret0, err
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
func (_Registry *RegistrySession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
func (_Registry *RegistryCallerSession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() view returns(address)
func (_Registry *RegistryCaller) GetChannelImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelImplementation")
	return *ret0, err
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() view returns(address)
func (_Registry *RegistrySession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() view returns(address)
func (_Registry *RegistryCallerSession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_Registry *RegistryCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_Registry *RegistrySession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_Registry *RegistryCallerSession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistryCaller) GetHermesAddress(opts *bind.CallOpts, _hermesOperator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getHermesAddress", _hermesOperator)
	return *ret0, err
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistrySession) GetHermesAddress(_hermesOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetHermesAddress(&_Registry.CallOpts, _hermesOperator)
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistryCallerSession) GetHermesAddress(_hermesOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetHermesAddress(&_Registry.CallOpts, _hermesOperator)
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistryCaller) GetHermesImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getHermesImplementation")
	return *ret0, err
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistrySession) GetHermesImplementation() (common.Address, error) {
	return _Registry.Contract.GetHermesImplementation(&_Registry.CallOpts)
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistryCallerSession) GetHermesImplementation() (common.Address, error) {
	return _Registry.Contract.GetHermesImplementation(&_Registry.CallOpts)
}

// GetHermesURL is a free data retrieval call binding the contract method 0xbf1eb88a.
//
// Solidity: function getHermesURL(address _hermesId) view returns(bytes)
func (_Registry *RegistryCaller) GetHermesURL(opts *bind.CallOpts, _hermesId common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getHermesURL", _hermesId)
	return *ret0, err
}

// GetHermesURL is a free data retrieval call binding the contract method 0xbf1eb88a.
//
// Solidity: function getHermesURL(address _hermesId) view returns(bytes)
func (_Registry *RegistrySession) GetHermesURL(_hermesId common.Address) ([]byte, error) {
	return _Registry.Contract.GetHermesURL(&_Registry.CallOpts, _hermesId)
}

// GetHermesURL is a free data retrieval call binding the contract method 0xbf1eb88a.
//
// Solidity: function getHermesURL(address _hermesId) view returns(bytes)
func (_Registry *RegistryCallerSession) GetHermesURL(_hermesId common.Address) ([]byte, error) {
	return _Registry.Contract.GetHermesURL(&_Registry.CallOpts, _hermesId)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
func (_Registry *RegistryCaller) GetProxyCode(opts *bind.CallOpts, _implementation common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getProxyCode", _implementation)
	return *ret0, err
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
func (_Registry *RegistrySession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
func (_Registry *RegistryCallerSession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// HasParentRegistry is a free data retrieval call binding the contract method 0x332ff7bd.
//
// Solidity: function hasParentRegistry(address _parentAddress) pure returns(bool)
func (_Registry *RegistryCaller) HasParentRegistry(opts *bind.CallOpts, _parentAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "hasParentRegistry", _parentAddress)
	return *ret0, err
}

// HasParentRegistry is a free data retrieval call binding the contract method 0x332ff7bd.
//
// Solidity: function hasParentRegistry(address _parentAddress) pure returns(bool)
func (_Registry *RegistrySession) HasParentRegistry(_parentAddress common.Address) (bool, error) {
	return _Registry.Contract.HasParentRegistry(&_Registry.CallOpts, _parentAddress)
}

// HasParentRegistry is a free data retrieval call binding the contract method 0x332ff7bd.
//
// Solidity: function hasParentRegistry(address _parentAddress) pure returns(bool)
func (_Registry *RegistryCallerSession) HasParentRegistry(_parentAddress common.Address) (bool, error) {
	return _Registry.Contract.HasParentRegistry(&_Registry.CallOpts, _parentAddress)
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake, bytes url)
func (_Registry *RegistryCaller) Hermeses(opts *bind.CallOpts, arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
	Url      []byte
}, error) {
	ret := new(struct {
		Operator common.Address
		Stake    [24]byte
		Url      []byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "hermeses", arg0)
	return *ret, err
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake, bytes url)
func (_Registry *RegistrySession) Hermeses(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
	Url      []byte
}, error) {
	return _Registry.Contract.Hermeses(&_Registry.CallOpts, arg0)
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake, bytes url)
func (_Registry *RegistryCallerSession) Hermeses(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
	Url      []byte
}, error) {
	return _Registry.Contract.Hermeses(&_Registry.CallOpts, arg0)
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistryCaller) IsHermes(opts *bind.CallOpts, _hermesId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isHermes", _hermesId)
	return *ret0, err
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistrySession) IsHermes(_hermesId common.Address) (bool, error) {
	return _Registry.Contract.IsHermes(&_Registry.CallOpts, _hermesId)
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistryCallerSession) IsHermes(_hermesId common.Address) (bool, error) {
	return _Registry.Contract.IsHermes(&_Registry.CallOpts, _hermesId)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identity) view returns(bool)
func (_Registry *RegistryCaller) IsRegistered(opts *bind.CallOpts, _identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isRegistered", _identity)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identity) view returns(bool)
func (_Registry *RegistrySession) IsRegistered(_identity common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identity) view returns(bool)
func (_Registry *RegistryCallerSession) IsRegistered(_identity common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identity)
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistryCaller) MinimalHermesStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "minimalHermesStake")
	return *ret0, err
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistrySession) MinimalHermesStake() (*big.Int, error) {
	return _Registry.Contract.MinimalHermesStake(&_Registry.CallOpts)
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistryCallerSession) MinimalHermesStake() (*big.Int, error) {
	return _Registry.Contract.MinimalHermesStake(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Registry *RegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Registry *RegistrySession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Registry *RegistryCallerSession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistrySession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistrySession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// RegisterHermes is a paid mutator transaction binding the contract method 0xd5929fe3.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _hermesStake, uint16 _hermesFee, uint256 _minChannelStake, uint256 _maxChannelStake, bytes _url) returns()
func (_Registry *RegistryTransactor) RegisterHermes(opts *bind.TransactOpts, _hermesOperator common.Address, _hermesStake *big.Int, _hermesFee uint16, _minChannelStake *big.Int, _maxChannelStake *big.Int, _url []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerHermes", _hermesOperator, _hermesStake, _hermesFee, _minChannelStake, _maxChannelStake, _url)
}

// RegisterHermes is a paid mutator transaction binding the contract method 0xd5929fe3.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _hermesStake, uint16 _hermesFee, uint256 _minChannelStake, uint256 _maxChannelStake, bytes _url) returns()
func (_Registry *RegistrySession) RegisterHermes(_hermesOperator common.Address, _hermesStake *big.Int, _hermesFee uint16, _minChannelStake *big.Int, _maxChannelStake *big.Int, _url []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterHermes(&_Registry.TransactOpts, _hermesOperator, _hermesStake, _hermesFee, _minChannelStake, _maxChannelStake, _url)
}

// RegisterHermes is a paid mutator transaction binding the contract method 0xd5929fe3.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _hermesStake, uint16 _hermesFee, uint256 _minChannelStake, uint256 _maxChannelStake, bytes _url) returns()
func (_Registry *RegistryTransactorSession) RegisterHermes(_hermesOperator common.Address, _hermesStake *big.Int, _hermesFee uint16, _minChannelStake *big.Int, _maxChannelStake *big.Int, _url []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterHermes(&_Registry.TransactOpts, _hermesOperator, _hermesStake, _hermesFee, _minChannelStake, _maxChannelStake, _url)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerIdentity", _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistrySession) RegisterIdentity(_hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactorSession) RegisterIdentity(_hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistrySession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferCollectedFeeTo", _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistrySession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactorSession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UpdateHermsURL is a paid mutator transaction binding the contract method 0x66330f56.
//
// Solidity: function updateHermsURL(address _hermesId, bytes _url, bytes _signature) returns()
func (_Registry *RegistryTransactor) UpdateHermsURL(opts *bind.TransactOpts, _hermesId common.Address, _url []byte, _signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateHermsURL", _hermesId, _url, _signature)
}

// UpdateHermsURL is a paid mutator transaction binding the contract method 0x66330f56.
//
// Solidity: function updateHermsURL(address _hermesId, bytes _url, bytes _signature) returns()
func (_Registry *RegistrySession) UpdateHermsURL(_hermesId common.Address, _url []byte, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateHermsURL(&_Registry.TransactOpts, _hermesId, _url, _signature)
}

// UpdateHermsURL is a paid mutator transaction binding the contract method 0x66330f56.
//
// Solidity: function updateHermsURL(address _hermesId, bytes _url, bytes _signature) returns()
func (_Registry *RegistryTransactorSession) UpdateHermsURL(_hermesId common.Address, _url []byte, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateHermsURL(&_Registry.TransactOpts, _hermesId, _url, _signature)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistrySession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactorSession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
}

// RegistryConsumerChannelCreatedIterator is returned from FilterConsumerChannelCreated and is used to iterate over the raw logs and unpacked data for ConsumerChannelCreated events raised by the Registry contract.
type RegistryConsumerChannelCreatedIterator struct {
	Event *RegistryConsumerChannelCreated // Event containing the contract specifics and raw log

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
func (it *RegistryConsumerChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryConsumerChannelCreated)
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
		it.Event = new(RegistryConsumerChannelCreated)
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
func (it *RegistryConsumerChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryConsumerChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryConsumerChannelCreated represents a ConsumerChannelCreated event raised by the Registry contract.
type RegistryConsumerChannelCreated struct {
	IdentityHash   common.Address
	HermesId       common.Address
	ChannelAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConsumerChannelCreated is a free log retrieval operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
func (_Registry *RegistryFilterer) FilterConsumerChannelCreated(opts *bind.FilterOpts, identityHash []common.Address, hermesId []common.Address) (*RegistryConsumerChannelCreatedIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ConsumerChannelCreated", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryConsumerChannelCreatedIterator{contract: _Registry.contract, event: "ConsumerChannelCreated", logs: logs, sub: sub}, nil
}

// WatchConsumerChannelCreated is a free log subscription operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
func (_Registry *RegistryFilterer) WatchConsumerChannelCreated(opts *bind.WatchOpts, sink chan<- *RegistryConsumerChannelCreated, identityHash []common.Address, hermesId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ConsumerChannelCreated", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryConsumerChannelCreated)
				if err := _Registry.contract.UnpackLog(event, "ConsumerChannelCreated", log); err != nil {
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

// ParseConsumerChannelCreated is a log parse operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
func (_Registry *RegistryFilterer) ParseConsumerChannelCreated(log types.Log) (*RegistryConsumerChannelCreated, error) {
	event := new(RegistryConsumerChannelCreated)
	if err := _Registry.contract.UnpackLog(event, "ConsumerChannelCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the Registry contract.
type RegistryDestinationChangedIterator struct {
	Event *RegistryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *RegistryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDestinationChanged)
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
		it.Event = new(RegistryDestinationChanged)
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
func (it *RegistryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDestinationChanged represents a DestinationChanged event raised by the Registry contract.
type RegistryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*RegistryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryDestinationChangedIterator{contract: _Registry.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *RegistryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDestinationChanged)
				if err := _Registry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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

// ParseDestinationChanged is a log parse operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) ParseDestinationChanged(log types.Log) (*RegistryDestinationChanged, error) {
	event := new(RegistryDestinationChanged)
	if err := _Registry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryHermesURLUpdatedIterator is returned from FilterHermesURLUpdated and is used to iterate over the raw logs and unpacked data for HermesURLUpdated events raised by the Registry contract.
type RegistryHermesURLUpdatedIterator struct {
	Event *RegistryHermesURLUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryHermesURLUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryHermesURLUpdated)
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
		it.Event = new(RegistryHermesURLUpdated)
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
func (it *RegistryHermesURLUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryHermesURLUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryHermesURLUpdated represents a HermesURLUpdated event raised by the Registry contract.
type RegistryHermesURLUpdated struct {
	HermesId common.Address
	NewURL   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterHermesURLUpdated is a free log retrieval operation binding the contract event 0xd8c638c85547b8717e0d5ca292cff6dbe8fc02fa6e6863a047971c39511643c7.
//
// Solidity: event HermesURLUpdated(address indexed hermesId, bytes newURL)
func (_Registry *RegistryFilterer) FilterHermesURLUpdated(opts *bind.FilterOpts, hermesId []common.Address) (*RegistryHermesURLUpdatedIterator, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "HermesURLUpdated", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryHermesURLUpdatedIterator{contract: _Registry.contract, event: "HermesURLUpdated", logs: logs, sub: sub}, nil
}

// WatchHermesURLUpdated is a free log subscription operation binding the contract event 0xd8c638c85547b8717e0d5ca292cff6dbe8fc02fa6e6863a047971c39511643c7.
//
// Solidity: event HermesURLUpdated(address indexed hermesId, bytes newURL)
func (_Registry *RegistryFilterer) WatchHermesURLUpdated(opts *bind.WatchOpts, sink chan<- *RegistryHermesURLUpdated, hermesId []common.Address) (event.Subscription, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "HermesURLUpdated", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryHermesURLUpdated)
				if err := _Registry.contract.UnpackLog(event, "HermesURLUpdated", log); err != nil {
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

// ParseHermesURLUpdated is a log parse operation binding the contract event 0xd8c638c85547b8717e0d5ca292cff6dbe8fc02fa6e6863a047971c39511643c7.
//
// Solidity: event HermesURLUpdated(address indexed hermesId, bytes newURL)
func (_Registry *RegistryFilterer) ParseHermesURLUpdated(log types.Log) (*RegistryHermesURLUpdated, error) {
	event := new(RegistryHermesURLUpdated)
	if err := _Registry.contract.UnpackLog(event, "HermesURLUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRegisteredHermesIterator is returned from FilterRegisteredHermes and is used to iterate over the raw logs and unpacked data for RegisteredHermes events raised by the Registry contract.
type RegistryRegisteredHermesIterator struct {
	Event *RegistryRegisteredHermes // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredHermesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredHermes)
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
		it.Event = new(RegistryRegisteredHermes)
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
func (it *RegistryRegisteredHermesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredHermesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredHermes represents a RegisteredHermes event raised by the Registry contract.
type RegistryRegisteredHermes struct {
	HermesId       common.Address
	HermesOperator common.Address
	Ur             []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegisteredHermes is a free log retrieval operation binding the contract event 0xf06d60cc2f463635fd237ad87f1d007af54840b82e7e4561707b1be63d91c260.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator, bytes ur)
func (_Registry *RegistryFilterer) FilterRegisteredHermes(opts *bind.FilterOpts, hermesId []common.Address) (*RegistryRegisteredHermesIterator, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredHermes", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredHermesIterator{contract: _Registry.contract, event: "RegisteredHermes", logs: logs, sub: sub}, nil
}

// WatchRegisteredHermes is a free log subscription operation binding the contract event 0xf06d60cc2f463635fd237ad87f1d007af54840b82e7e4561707b1be63d91c260.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator, bytes ur)
func (_Registry *RegistryFilterer) WatchRegisteredHermes(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredHermes, hermesId []common.Address) (event.Subscription, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredHermes", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredHermes)
				if err := _Registry.contract.UnpackLog(event, "RegisteredHermes", log); err != nil {
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

// ParseRegisteredHermes is a log parse operation binding the contract event 0xf06d60cc2f463635fd237ad87f1d007af54840b82e7e4561707b1be63d91c260.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator, bytes ur)
func (_Registry *RegistryFilterer) ParseRegisteredHermes(log types.Log) (*RegistryRegisteredHermes, error) {
	event := new(RegistryRegisteredHermes)
	if err := _Registry.contract.UnpackLog(event, "RegisteredHermes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRegisteredIdentityIterator is returned from FilterRegisteredIdentity and is used to iterate over the raw logs and unpacked data for RegisteredIdentity events raised by the Registry contract.
type RegistryRegisteredIdentityIterator struct {
	Event *RegistryRegisteredIdentity // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredIdentity)
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
		it.Event = new(RegistryRegisteredIdentity)
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
func (it *RegistryRegisteredIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredIdentity represents a RegisteredIdentity event raised by the Registry contract.
type RegistryRegisteredIdentity struct {
	IdentityHash common.Address
	HermesId     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredIdentity is a free log retrieval operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) FilterRegisteredIdentity(opts *bind.FilterOpts, identityHash []common.Address, hermesId []common.Address) (*RegistryRegisteredIdentityIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredIdentity", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIdentityIterator{contract: _Registry.contract, event: "RegisteredIdentity", logs: logs, sub: sub}, nil
}

// WatchRegisteredIdentity is a free log subscription operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) WatchRegisteredIdentity(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredIdentity, identityHash []common.Address, hermesId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredIdentity", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredIdentity)
				if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
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

// ParseRegisteredIdentity is a log parse operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) ParseRegisteredIdentity(log types.Log) (*RegistryRegisteredIdentity, error) {
	event := new(RegistryRegisteredIdentity)
	if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
		return nil, err
	}
	return event, nil
}
