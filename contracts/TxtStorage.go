// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TxtStorageABI is the input ABI used to generate the binding from.
const TxtStorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ReputationGiven\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"AddRep\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"GetData\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"GetDataFree\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RequestData\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"hd\",\"type\":\"string\"}],\"name\":\"SetHashData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dude\",\"type\":\"address\"}],\"name\":\"TransferData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"hash_\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DoneCustomers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetHashData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetReputation\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PendingCustomers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"toBytes\",\"outputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TxtStorageBin is the compiled bytecode used for deploying new contracts.
const TxtStorageBin = `{
	"linkReferences": {},
	"object": "608060405234801561001057600080fd5b50604051610b46380380610b4683398101604090815281516020808401519284015160008054600160a060020a031916331790559184018051909492909201916100609160029190860190610087565b5060078290556000600455805161007e906001906020840190610087565b50505050610122565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c857805160ff19168380011785556100f5565b828001600101855582156100f5579182015b828111156100f55782518255916020019190600101906100da565b50610101929150610105565b5090565b61011f91905b80821115610101576000815560010161010b565b90565b610a15806101316000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631401bd9581146100b357806320ccbe681461013d578063485e36c814610160578063593b79fe1461019457806376b8e528146101b55780637b295c51146101bd57806384b82a8a146101d55780638b2cb74b146101ed578063b674f79a14610214578063e27e91601461026d578063fe704a9d14610282575b600080fd5b3480156100bf57600080fd5b506100c8610297565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101025781810151838201526020016100ea565b50505050905090810190601f16801561012f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014957600080fd5b5061015e600160a060020a036004351661032d565b005b34801561016c57600080fd5b506101786004356103c8565b60408051600160a060020a039092168252519081900360200190f35b3480156101a057600080fd5b506100c8600160a060020a03600435166103f0565b61015e61041b565b3480156101c957600080fd5b5061015e60043561062d565b3480156101e157600080fd5b50610178600435610698565b3480156101f957600080fd5b506102026106a6565b60408051918252519081900360200190f35b34801561022057600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015e9436949293602493928401919081908401838280828437509497506106ac9650505050505050565b34801561027957600080fd5b5061015e6106d8565b34801561028e57600080fd5b506100c8610862565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156103225780601f106102f757610100808354040283529160200191610322565b820191906000526020600020905b81548152906001019060200180831161030557829003601f168201915b505050505090505b90565b600054600160a060020a0316331461034457600080fd5b61034d816108c0565b1561035757600080fd5b60008054600454604051600160a060020a039092169281156108fc029290818181858888f19350505050158015610392573d6000803e3d6000fd5b50600160a060020a03166000908152600860205260409020805474ff000000000000000000000000000000000000000019169055565b60068054829081106103d657fe5b600091825260209091200154600160a060020a0316905081565b6040805174140000000000000000000000000000000000000000909218601483015260348201905290565b600754341015610475576040805160e560020a62461bcd02815260206004820152601160248201527f507269636520697320746f6f206c6f772e000000000000000000000000000000604482015290519081900360640190fd5b61047e336108c0565b156104f9576040805160e560020a62461bcd02815260206004820152603060248201527f596f7527766520616c7265616479206d6164652061207265717565737420666f60448201527f72207468697320646f63756d656e742e00000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a031633141561051157600080fd5b600354600012610591576040805160e560020a62461bcd02815260206004820152603960248201527f5468697320646f63756d656e74206861732061206e656761746976652072617460448201527f696e672c207468657265666f7265206974277320667265652e00000000000000606482015290519081900360840190fd5b33600081815260086020908152604091829020805474ff00000000000000000000000000000000000000001973ffffffffffffffffffffffffffffffffffffffff19909116851716740100000000000000000000000000000000000000001790556004805434019055815192835290517f70b912fa029fe2496dbbf0edd4740f19ff757041675de8572cdd263fa630ed939281900390910190a1565b600054600160a060020a031633141561064557600080fd5b61064e336108c0565b151561065957600080fd5b60038054820190556040805133815290517fe645ec368e9f33f328b76bbcf9d56855e7bfe7d5892c74e354aadfc80685257a9181900360200190a15b50565b60058054829081106103d657fe5b60035490565b600054600160a060020a03163314156100ae5780516106d2906001906020840190610951565b50610695565b600354600013610732576040805160e560020a62461bcd02815260206004820152601e60248201527f5468697320646f63756d656e74206973206e6f7420666f7220667265652e0000604482015290519081900360640190fd5b61073b336108c0565b156107b6576040805160e560020a62461bcd02815260206004820152603060248201527f596f7527766520616c7265616479206d6164652061207265717565737420666f60448201527f72207468697320646f63756d656e742e00000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a03163314156107ce57600080fd5b33600081815260086020908152604091829020805474ff00000000000000000000000000000000000000001973ffffffffffffffffffffffffffffffffffffffff1990911685171674010000000000000000000000000000000000000000179055815192835290517f70b912fa029fe2496dbbf0edd4740f19ff757041675de8572cdd263fa630ed939281900390910190a1565b60028054604080516020601f60001961010060018716150201909416859004938401819004810282018101909252828152606093909290918301828280156103225780601f106102f757610100808354040283529160200191610322565b60006108cb82610914565b15156108d95750600061090f565b50600160a060020a03811660009081526008602052604090205474010000000000000000000000000000000000000000900460ff165b919050565b600160a060020a03808216600090815260086020526040812054909161093a91166103f0565b5115156109495750600061090f565b506001919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061099257805160ff19168380011785556109bf565b828001600101855582156109bf579182015b828111156109bf5782518255916020019190600101906109a4565b506109cb9291506109cf565b5090565b61032a91905b808211156109cb57600081556001016109d55600a165627a7a72305820ee4762095c2411800322547234431db8105f0d374425f505bec9f1b91a1609660029",
	"opcodes": "PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 MLOAD PUSH2 0xB46 CODESIZE SUB DUP1 PUSH2 0xB46 DUP4 CODECOPY DUP2 ADD PUSH1 0x40 SWAP1 DUP2 MSTORE DUP2 MLOAD PUSH1 0x20 DUP1 DUP5 ADD MLOAD SWAP3 DUP5 ADD MLOAD PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB NOT AND CALLER OR SWAP1 SSTORE SWAP2 DUP5 ADD DUP1 MLOAD SWAP1 SWAP5 SWAP3 SWAP1 SWAP3 ADD SWAP2 PUSH2 0x60 SWAP2 PUSH1 0x2 SWAP2 SWAP1 DUP7 ADD SWAP1 PUSH2 0x87 JUMP JUMPDEST POP PUSH1 0x7 DUP3 SWAP1 SSTORE PUSH1 0x0 PUSH1 0x4 SSTORE DUP1 MLOAD PUSH2 0x7E SWAP1 PUSH1 0x1 SWAP1 PUSH1 0x20 DUP5 ADD SWAP1 PUSH2 0x87 JUMP JUMPDEST POP POP POP POP PUSH2 0x122 JUMP JUMPDEST DUP3 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x1F ADD PUSH1 0x20 SWAP1 DIV DUP2 ADD SWAP3 DUP3 PUSH1 0x1F LT PUSH2 0xC8 JUMPI DUP1 MLOAD PUSH1 0xFF NOT AND DUP4 DUP1 ADD OR DUP6 SSTORE PUSH2 0xF5 JUMP JUMPDEST DUP3 DUP1 ADD PUSH1 0x1 ADD DUP6 SSTORE DUP3 ISZERO PUSH2 0xF5 JUMPI SWAP2 DUP3 ADD JUMPDEST DUP3 DUP2 GT ISZERO PUSH2 0xF5 JUMPI DUP3 MLOAD DUP3 SSTORE SWAP2 PUSH1 0x20 ADD SWAP2 SWAP1 PUSH1 0x1 ADD SWAP1 PUSH2 0xDA JUMP JUMPDEST POP PUSH2 0x101 SWAP3 SWAP2 POP PUSH2 0x105 JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST PUSH2 0x11F SWAP2 SWAP1 JUMPDEST DUP1 DUP3 GT ISZERO PUSH2 0x101 JUMPI PUSH1 0x0 DUP2 SSTORE PUSH1 0x1 ADD PUSH2 0x10B JUMP JUMPDEST SWAP1 JUMP JUMPDEST PUSH2 0xA15 DUP1 PUSH2 0x131 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0xAE JUMPI PUSH4 0xFFFFFFFF PUSH29 0x100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 CALLDATALOAD DIV AND PUSH4 0x1401BD95 DUP2 EQ PUSH2 0xB3 JUMPI DUP1 PUSH4 0x20CCBE68 EQ PUSH2 0x13D JUMPI DUP1 PUSH4 0x485E36C8 EQ PUSH2 0x160 JUMPI DUP1 PUSH4 0x593B79FE EQ PUSH2 0x194 JUMPI DUP1 PUSH4 0x76B8E528 EQ PUSH2 0x1B5 JUMPI DUP1 PUSH4 0x7B295C51 EQ PUSH2 0x1BD JUMPI DUP1 PUSH4 0x84B82A8A EQ PUSH2 0x1D5 JUMPI DUP1 PUSH4 0x8B2CB74B EQ PUSH2 0x1ED JUMPI DUP1 PUSH4 0xB674F79A EQ PUSH2 0x214 JUMPI DUP1 PUSH4 0xE27E9160 EQ PUSH2 0x26D JUMPI DUP1 PUSH4 0xFE704A9D EQ PUSH2 0x282 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0xBF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0xC8 PUSH2 0x297 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD PUSH1 0x20 DUP1 DUP3 MSTORE DUP4 MLOAD DUP2 DUP4 ADD MSTORE DUP4 MLOAD SWAP2 SWAP3 DUP4 SWAP3 SWAP1 DUP4 ADD SWAP2 DUP6 ADD SWAP1 DUP1 DUP4 DUP4 PUSH1 0x0 JUMPDEST DUP4 DUP2 LT ISZERO PUSH2 0x102 JUMPI DUP2 DUP2 ADD MLOAD DUP4 DUP3 ADD MSTORE PUSH1 0x20 ADD PUSH2 0xEA JUMP JUMPDEST POP POP POP POP SWAP1 POP SWAP1 DUP2 ADD SWAP1 PUSH1 0x1F AND DUP1 ISZERO PUSH2 0x12F JUMPI DUP1 DUP3 SUB DUP1 MLOAD PUSH1 0x1 DUP4 PUSH1 0x20 SUB PUSH2 0x100 EXP SUB NOT AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP JUMPDEST POP SWAP3 POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x149 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x15E PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB PUSH1 0x4 CALLDATALOAD AND PUSH2 0x32D JUMP JUMPDEST STOP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x16C JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x178 PUSH1 0x4 CALLDATALOAD PUSH2 0x3C8 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB SWAP1 SWAP3 AND DUP3 MSTORE MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x20 ADD SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1A0 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0xC8 PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB PUSH1 0x4 CALLDATALOAD AND PUSH2 0x3F0 JUMP JUMPDEST PUSH2 0x15E PUSH2 0x41B JUMP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1C9 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x15E PUSH1 0x4 CALLDATALOAD PUSH2 0x62D JUMP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1E1 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x178 PUSH1 0x4 CALLDATALOAD PUSH2 0x698 JUMP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1F9 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x202 PUSH2 0x6A6 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP2 DUP3 MSTORE MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x20 ADD SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x220 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 DUP1 MLOAD PUSH1 0x20 PUSH1 0x4 DUP1 CALLDATALOAD DUP1 DUP3 ADD CALLDATALOAD PUSH1 0x1F DUP2 ADD DUP5 SWAP1 DIV DUP5 MUL DUP6 ADD DUP5 ADD SWAP1 SWAP6 MSTORE DUP5 DUP5 MSTORE PUSH2 0x15E SWAP5 CALLDATASIZE SWAP5 SWAP3 SWAP4 PUSH1 0x24 SWAP4 SWAP3 DUP5 ADD SWAP2 SWAP1 DUP2 SWAP1 DUP5 ADD DUP4 DUP3 DUP1 DUP3 DUP5 CALLDATACOPY POP SWAP5 SWAP8 POP PUSH2 0x6AC SWAP7 POP POP POP POP POP POP POP JUMP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x279 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x15E PUSH2 0x6D8 JUMP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x28E JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0xC8 PUSH2 0x862 JUMP JUMPDEST PUSH1 0x1 DUP1 SLOAD PUSH1 0x40 DUP1 MLOAD PUSH1 0x20 PUSH1 0x1F PUSH1 0x2 PUSH1 0x0 NOT PUSH2 0x100 DUP8 DUP10 AND ISZERO MUL ADD SWAP1 SWAP6 AND SWAP5 SWAP1 SWAP5 DIV SWAP4 DUP5 ADD DUP2 SWAP1 DIV DUP2 MUL DUP3 ADD DUP2 ADD SWAP1 SWAP3 MSTORE DUP3 DUP2 MSTORE PUSH1 0x60 SWAP4 SWAP1 SWAP3 SWAP1 SWAP2 DUP4 ADD DUP3 DUP3 DUP1 ISZERO PUSH2 0x322 JUMPI DUP1 PUSH1 0x1F LT PUSH2 0x2F7 JUMPI PUSH2 0x100 DUP1 DUP4 SLOAD DIV MUL DUP4 MSTORE SWAP2 PUSH1 0x20 ADD SWAP2 PUSH2 0x322 JUMP JUMPDEST DUP3 ADD SWAP2 SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 JUMPDEST DUP2 SLOAD DUP2 MSTORE SWAP1 PUSH1 0x1 ADD SWAP1 PUSH1 0x20 ADD DUP1 DUP4 GT PUSH2 0x305 JUMPI DUP3 SWAP1 SUB PUSH1 0x1F AND DUP3 ADD SWAP2 JUMPDEST POP POP POP POP POP SWAP1 POP JUMPDEST SWAP1 JUMP JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND CALLER EQ PUSH2 0x344 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x34D DUP2 PUSH2 0x8C0 JUMP JUMPDEST ISZERO PUSH2 0x357 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x4 SLOAD PUSH1 0x40 MLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB SWAP1 SWAP3 AND SWAP3 DUP2 ISZERO PUSH2 0x8FC MUL SWAP3 SWAP1 DUP2 DUP2 DUP2 DUP6 DUP9 DUP9 CALL SWAP4 POP POP POP POP ISZERO DUP1 ISZERO PUSH2 0x392 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x8 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 DUP1 SLOAD PUSH21 0xFF0000000000000000000000000000000000000000 NOT AND SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x6 DUP1 SLOAD DUP3 SWAP1 DUP2 LT PUSH2 0x3D6 JUMPI INVALID JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 SWAP1 SWAP2 KECCAK256 ADD SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND SWAP1 POP DUP2 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD PUSH21 0x140000000000000000000000000000000000000000 SWAP1 SWAP3 XOR PUSH1 0x14 DUP4 ADD MSTORE PUSH1 0x34 DUP3 ADD SWAP1 MSTORE SWAP1 JUMP JUMPDEST PUSH1 0x7 SLOAD CALLVALUE LT ISZERO PUSH2 0x475 JUMPI PUSH1 0x40 DUP1 MLOAD PUSH1 0xE5 PUSH1 0x2 EXP PUSH3 0x461BCD MUL DUP2 MSTORE PUSH1 0x20 PUSH1 0x4 DUP3 ADD MSTORE PUSH1 0x11 PUSH1 0x24 DUP3 ADD MSTORE PUSH32 0x507269636520697320746F6F206C6F772E000000000000000000000000000000 PUSH1 0x44 DUP3 ADD MSTORE SWAP1 MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x64 ADD SWAP1 REVERT JUMPDEST PUSH2 0x47E CALLER PUSH2 0x8C0 JUMP JUMPDEST ISZERO PUSH2 0x4F9 JUMPI PUSH1 0x40 DUP1 MLOAD PUSH1 0xE5 PUSH1 0x2 EXP PUSH3 0x461BCD MUL DUP2 MSTORE PUSH1 0x20 PUSH1 0x4 DUP3 ADD MSTORE PUSH1 0x30 PUSH1 0x24 DUP3 ADD MSTORE PUSH32 0x596F7527766520616C7265616479206D6164652061207265717565737420666F PUSH1 0x44 DUP3 ADD MSTORE PUSH32 0x72207468697320646F63756D656E742E00000000000000000000000000000000 PUSH1 0x64 DUP3 ADD MSTORE SWAP1 MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x84 ADD SWAP1 REVERT JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND CALLER EQ ISZERO PUSH2 0x511 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x3 SLOAD PUSH1 0x0 SLT PUSH2 0x591 JUMPI PUSH1 0x40 DUP1 MLOAD PUSH1 0xE5 PUSH1 0x2 EXP PUSH3 0x461BCD MUL DUP2 MSTORE PUSH1 0x20 PUSH1 0x4 DUP3 ADD MSTORE PUSH1 0x39 PUSH1 0x24 DUP3 ADD MSTORE PUSH32 0x5468697320646F63756D656E74206861732061206E6567617469766520726174 PUSH1 0x44 DUP3 ADD MSTORE PUSH32 0x696E672C207468657265666F7265206974277320667265652E00000000000000 PUSH1 0x64 DUP3 ADD MSTORE SWAP1 MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x84 ADD SWAP1 REVERT JUMPDEST CALLER PUSH1 0x0 DUP2 DUP2 MSTORE PUSH1 0x8 PUSH1 0x20 SWAP1 DUP2 MSTORE PUSH1 0x40 SWAP2 DUP3 SWAP1 KECCAK256 DUP1 SLOAD PUSH21 0xFF0000000000000000000000000000000000000000 NOT PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF NOT SWAP1 SWAP2 AND DUP6 OR AND PUSH21 0x10000000000000000000000000000000000000000 OR SWAP1 SSTORE PUSH1 0x4 DUP1 SLOAD CALLVALUE ADD SWAP1 SSTORE DUP2 MLOAD SWAP3 DUP4 MSTORE SWAP1 MLOAD PUSH32 0x70B912FA029FE2496DBBF0EDD4740F19FF757041675DE8572CDD263FA630ED93 SWAP3 DUP2 SWAP1 SUB SWAP1 SWAP2 ADD SWAP1 LOG1 JUMP JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND CALLER EQ ISZERO PUSH2 0x645 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x64E CALLER PUSH2 0x8C0 JUMP JUMPDEST ISZERO ISZERO PUSH2 0x659 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x3 DUP1 SLOAD DUP3 ADD SWAP1 SSTORE PUSH1 0x40 DUP1 MLOAD CALLER DUP2 MSTORE SWAP1 MLOAD PUSH32 0xE645EC368E9F33F328B76BBCF9D56855E7BFE7D5892C74E354AADFC80685257A SWAP2 DUP2 SWAP1 SUB PUSH1 0x20 ADD SWAP1 LOG1 JUMPDEST POP JUMP JUMPDEST PUSH1 0x5 DUP1 SLOAD DUP3 SWAP1 DUP2 LT PUSH2 0x3D6 JUMPI INVALID JUMPDEST PUSH1 0x3 SLOAD SWAP1 JUMP JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND CALLER EQ ISZERO PUSH2 0xAE JUMPI DUP1 MLOAD PUSH2 0x6D2 SWAP1 PUSH1 0x1 SWAP1 PUSH1 0x20 DUP5 ADD SWAP1 PUSH2 0x951 JUMP JUMPDEST POP PUSH2 0x695 JUMP JUMPDEST PUSH1 0x3 SLOAD PUSH1 0x0 SGT PUSH2 0x732 JUMPI PUSH1 0x40 DUP1 MLOAD PUSH1 0xE5 PUSH1 0x2 EXP PUSH3 0x461BCD MUL DUP2 MSTORE PUSH1 0x20 PUSH1 0x4 DUP3 ADD MSTORE PUSH1 0x1E PUSH1 0x24 DUP3 ADD MSTORE PUSH32 0x5468697320646F63756D656E74206973206E6F7420666F7220667265652E0000 PUSH1 0x44 DUP3 ADD MSTORE SWAP1 MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x64 ADD SWAP1 REVERT JUMPDEST PUSH2 0x73B CALLER PUSH2 0x8C0 JUMP JUMPDEST ISZERO PUSH2 0x7B6 JUMPI PUSH1 0x40 DUP1 MLOAD PUSH1 0xE5 PUSH1 0x2 EXP PUSH3 0x461BCD MUL DUP2 MSTORE PUSH1 0x20 PUSH1 0x4 DUP3 ADD MSTORE PUSH1 0x30 PUSH1 0x24 DUP3 ADD MSTORE PUSH32 0x596F7527766520616C7265616479206D6164652061207265717565737420666F PUSH1 0x44 DUP3 ADD MSTORE PUSH32 0x72207468697320646F63756D656E742E00000000000000000000000000000000 PUSH1 0x64 DUP3 ADD MSTORE SWAP1 MLOAD SWAP1 DUP2 SWAP1 SUB PUSH1 0x84 ADD SWAP1 REVERT JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB AND CALLER EQ ISZERO PUSH2 0x7CE JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST CALLER PUSH1 0x0 DUP2 DUP2 MSTORE PUSH1 0x8 PUSH1 0x20 SWAP1 DUP2 MSTORE PUSH1 0x40 SWAP2 DUP3 SWAP1 KECCAK256 DUP1 SLOAD PUSH21 0xFF0000000000000000000000000000000000000000 NOT PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF NOT SWAP1 SWAP2 AND DUP6 OR AND PUSH21 0x10000000000000000000000000000000000000000 OR SWAP1 SSTORE DUP2 MLOAD SWAP3 DUP4 MSTORE SWAP1 MLOAD PUSH32 0x70B912FA029FE2496DBBF0EDD4740F19FF757041675DE8572CDD263FA630ED93 SWAP3 DUP2 SWAP1 SUB SWAP1 SWAP2 ADD SWAP1 LOG1 JUMP JUMPDEST PUSH1 0x2 DUP1 SLOAD PUSH1 0x40 DUP1 MLOAD PUSH1 0x20 PUSH1 0x1F PUSH1 0x0 NOT PUSH2 0x100 PUSH1 0x1 DUP8 AND ISZERO MUL ADD SWAP1 SWAP5 AND DUP6 SWAP1 DIV SWAP4 DUP5 ADD DUP2 SWAP1 DIV DUP2 MUL DUP3 ADD DUP2 ADD SWAP1 SWAP3 MSTORE DUP3 DUP2 MSTORE PUSH1 0x60 SWAP4 SWAP1 SWAP3 SWAP1 SWAP2 DUP4 ADD DUP3 DUP3 DUP1 ISZERO PUSH2 0x322 JUMPI DUP1 PUSH1 0x1F LT PUSH2 0x2F7 JUMPI PUSH2 0x100 DUP1 DUP4 SLOAD DIV MUL DUP4 MSTORE SWAP2 PUSH1 0x20 ADD SWAP2 PUSH2 0x322 JUMP JUMPDEST PUSH1 0x0 PUSH2 0x8CB DUP3 PUSH2 0x914 JUMP JUMPDEST ISZERO ISZERO PUSH2 0x8D9 JUMPI POP PUSH1 0x0 PUSH2 0x90F JUMP JUMPDEST POP PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB DUP2 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x8 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SLOAD PUSH21 0x10000000000000000000000000000000000000000 SWAP1 DIV PUSH1 0xFF AND JUMPDEST SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x1 PUSH1 0xA0 PUSH1 0x2 EXP SUB DUP1 DUP3 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x8 PUSH1 0x20 MSTORE PUSH1 0x40 DUP2 KECCAK256 SLOAD SWAP1 SWAP2 PUSH2 0x93A SWAP2 AND PUSH2 0x3F0 JUMP JUMPDEST MLOAD ISZERO ISZERO PUSH2 0x949 JUMPI POP PUSH1 0x0 PUSH2 0x90F JUMP JUMPDEST POP PUSH1 0x1 SWAP2 SWAP1 POP JUMP JUMPDEST DUP3 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x1F ADD PUSH1 0x20 SWAP1 DIV DUP2 ADD SWAP3 DUP3 PUSH1 0x1F LT PUSH2 0x992 JUMPI DUP1 MLOAD PUSH1 0xFF NOT AND DUP4 DUP1 ADD OR DUP6 SSTORE PUSH2 0x9BF JUMP JUMPDEST DUP3 DUP1 ADD PUSH1 0x1 ADD DUP6 SSTORE DUP3 ISZERO PUSH2 0x9BF JUMPI SWAP2 DUP3 ADD JUMPDEST DUP3 DUP2 GT ISZERO PUSH2 0x9BF JUMPI DUP3 MLOAD DUP3 SSTORE SWAP2 PUSH1 0x20 ADD SWAP2 SWAP1 PUSH1 0x1 ADD SWAP1 PUSH2 0x9A4 JUMP JUMPDEST POP PUSH2 0x9CB SWAP3 SWAP2 POP PUSH2 0x9CF JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST PUSH2 0x32A SWAP2 SWAP1 JUMPDEST DUP1 DUP3 GT ISZERO PUSH2 0x9CB JUMPI PUSH1 0x0 DUP2 SSTORE PUSH1 0x1 ADD PUSH2 0x9D5 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0xee 0x47 PUSH3 0x95C24 GT DUP1 SUB 0x22 SLOAD PUSH19 0x34431DB8105F0D374425F505BEC9F1B91A1609 PUSH7 0x290000000000 ",
	"sourceMap": "28:3265:0:-;;;546:199;8:9:-1;5:2;;;30:1;27;20:12;5:2;546:199:0;;;;;;;;;;;;;;;;;;;;;;;;;;;618:5;:18;;-1:-1:-1;;;;;;618:18:0;626:10;618:18;;;546:199;;;647:12;;546:199;;;;;;;647:12;;:4;;:12;;;;;:::i;:::-;-1:-1:-1;670:5:0;:13;;;708:1;694:11;:15;720:17;;;;:9;;:17;;;;;:::i;:::-;;546:199;;;28:3265;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;28:3265:0;;;-1:-1:-1;28:3265:0;:::i;:::-;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;"
}`

// DeployTxtStorage deploys a new Ethereum contract, binding an instance of TxtStorage to it.
func DeployTxtStorage(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, price *big.Int, hash_ string) (common.Address, *types.Transaction, *TxtStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(TxtStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TxtStorageBin), backend, name_, price, hash_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TxtStorage{TxtStorageCaller: TxtStorageCaller{contract: contract}, TxtStorageTransactor: TxtStorageTransactor{contract: contract}, TxtStorageFilterer: TxtStorageFilterer{contract: contract}}, nil
}

// TxtStorage is an auto generated Go binding around an Ethereum contract.
type TxtStorage struct {
	TxtStorageCaller     // Read-only binding to the contract
	TxtStorageTransactor // Write-only binding to the contract
	TxtStorageFilterer   // Log filterer for contract events
}

// TxtStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxtStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxtStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxtStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxtStorageSession struct {
	Contract     *TxtStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxtStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxtStorageCallerSession struct {
	Contract *TxtStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TxtStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxtStorageTransactorSession struct {
	Contract     *TxtStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TxtStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxtStorageRaw struct {
	Contract *TxtStorage // Generic contract binding to access the raw methods on
}

// TxtStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxtStorageCallerRaw struct {
	Contract *TxtStorageCaller // Generic read-only contract binding to access the raw methods on
}

// TxtStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxtStorageTransactorRaw struct {
	Contract *TxtStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxtStorage creates a new instance of TxtStorage, bound to a specific deployed contract.
func NewTxtStorage(address common.Address, backend bind.ContractBackend) (*TxtStorage, error) {
	contract, err := bindTxtStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxtStorage{TxtStorageCaller: TxtStorageCaller{contract: contract}, TxtStorageTransactor: TxtStorageTransactor{contract: contract}, TxtStorageFilterer: TxtStorageFilterer{contract: contract}}, nil
}

// NewTxtStorageCaller creates a new read-only instance of TxtStorage, bound to a specific deployed contract.
func NewTxtStorageCaller(address common.Address, caller bind.ContractCaller) (*TxtStorageCaller, error) {
	contract, err := bindTxtStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxtStorageCaller{contract: contract}, nil
}

// NewTxtStorageTransactor creates a new write-only instance of TxtStorage, bound to a specific deployed contract.
func NewTxtStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*TxtStorageTransactor, error) {
	contract, err := bindTxtStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxtStorageTransactor{contract: contract}, nil
}

// NewTxtStorageFilterer creates a new log filterer instance of TxtStorage, bound to a specific deployed contract.
func NewTxtStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*TxtStorageFilterer, error) {
	contract, err := bindTxtStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxtStorageFilterer{contract: contract}, nil
}

// bindTxtStorage binds a generic wrapper to an already deployed contract.
func bindTxtStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxtStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxtStorage *TxtStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxtStorage.Contract.TxtStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxtStorage *TxtStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStorage.Contract.TxtStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxtStorage *TxtStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxtStorage.Contract.TxtStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxtStorage *TxtStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxtStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxtStorage *TxtStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxtStorage *TxtStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxtStorage.Contract.contract.Transact(opts, method, params...)
}

// DoneCustomers is a free data retrieval call binding the contract method 0x485e36c8.
//
// Solidity: function DoneCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageCaller) DoneCustomers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "DoneCustomers", arg0)
	return *ret0, err
}

// DoneCustomers is a free data retrieval call binding the contract method 0x485e36c8.
//
// Solidity: function DoneCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageSession) DoneCustomers(arg0 *big.Int) (common.Address, error) {
	return _TxtStorage.Contract.DoneCustomers(&_TxtStorage.CallOpts, arg0)
}

// DoneCustomers is a free data retrieval call binding the contract method 0x485e36c8.
//
// Solidity: function DoneCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageCallerSession) DoneCustomers(arg0 *big.Int) (common.Address, error) {
	return _TxtStorage.Contract.DoneCustomers(&_TxtStorage.CallOpts, arg0)
}

// GetHashData is a free data retrieval call binding the contract method 0x1401bd95.
//
// Solidity: function GetHashData() constant returns(string)
func (_TxtStorage *TxtStorageCaller) GetHashData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "GetHashData")
	return *ret0, err
}

// GetHashData is a free data retrieval call binding the contract method 0x1401bd95.
//
// Solidity: function GetHashData() constant returns(string)
func (_TxtStorage *TxtStorageSession) GetHashData() (string, error) {
	return _TxtStorage.Contract.GetHashData(&_TxtStorage.CallOpts)
}

// GetHashData is a free data retrieval call binding the contract method 0x1401bd95.
//
// Solidity: function GetHashData() constant returns(string)
func (_TxtStorage *TxtStorageCallerSession) GetHashData() (string, error) {
	return _TxtStorage.Contract.GetHashData(&_TxtStorage.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0xfe704a9d.
//
// Solidity: function GetName() constant returns(string)
func (_TxtStorage *TxtStorageCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "GetName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0xfe704a9d.
//
// Solidity: function GetName() constant returns(string)
func (_TxtStorage *TxtStorageSession) GetName() (string, error) {
	return _TxtStorage.Contract.GetName(&_TxtStorage.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0xfe704a9d.
//
// Solidity: function GetName() constant returns(string)
func (_TxtStorage *TxtStorageCallerSession) GetName() (string, error) {
	return _TxtStorage.Contract.GetName(&_TxtStorage.CallOpts)
}

// GetReputation is a free data retrieval call binding the contract method 0x8b2cb74b.
//
// Solidity: function GetReputation() constant returns(int256)
func (_TxtStorage *TxtStorageCaller) GetReputation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "GetReputation")
	return *ret0, err
}

// GetReputation is a free data retrieval call binding the contract method 0x8b2cb74b.
//
// Solidity: function GetReputation() constant returns(int256)
func (_TxtStorage *TxtStorageSession) GetReputation() (*big.Int, error) {
	return _TxtStorage.Contract.GetReputation(&_TxtStorage.CallOpts)
}

// GetReputation is a free data retrieval call binding the contract method 0x8b2cb74b.
//
// Solidity: function GetReputation() constant returns(int256)
func (_TxtStorage *TxtStorageCallerSession) GetReputation() (*big.Int, error) {
	return _TxtStorage.Contract.GetReputation(&_TxtStorage.CallOpts)
}

// PendingCustomers is a free data retrieval call binding the contract method 0x84b82a8a.
//
// Solidity: function PendingCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageCaller) PendingCustomers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "PendingCustomers", arg0)
	return *ret0, err
}

// PendingCustomers is a free data retrieval call binding the contract method 0x84b82a8a.
//
// Solidity: function PendingCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageSession) PendingCustomers(arg0 *big.Int) (common.Address, error) {
	return _TxtStorage.Contract.PendingCustomers(&_TxtStorage.CallOpts, arg0)
}

// PendingCustomers is a free data retrieval call binding the contract method 0x84b82a8a.
//
// Solidity: function PendingCustomers( uint256) constant returns(address)
func (_TxtStorage *TxtStorageCallerSession) PendingCustomers(arg0 *big.Int) (common.Address, error) {
	return _TxtStorage.Contract.PendingCustomers(&_TxtStorage.CallOpts, arg0)
}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(a address) constant returns(b bytes)
func (_TxtStorage *TxtStorageCaller) ToBytes(opts *bind.CallOpts, a common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TxtStorage.contract.Call(opts, out, "toBytes", a)
	return *ret0, err
}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(a address) constant returns(b bytes)
func (_TxtStorage *TxtStorageSession) ToBytes(a common.Address) ([]byte, error) {
	return _TxtStorage.Contract.ToBytes(&_TxtStorage.CallOpts, a)
}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(a address) constant returns(b bytes)
func (_TxtStorage *TxtStorageCallerSession) ToBytes(a common.Address) ([]byte, error) {
	return _TxtStorage.Contract.ToBytes(&_TxtStorage.CallOpts, a)
}

// AddRep is a paid mutator transaction binding the contract method 0x7b295c51.
//
// Solidity: function AddRep(val int256) returns()
func (_TxtStorage *TxtStorageTransactor) AddRep(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _TxtStorage.contract.Transact(opts, "AddRep", val)
}

// AddRep is a paid mutator transaction binding the contract method 0x7b295c51.
//
// Solidity: function AddRep(val int256) returns()
func (_TxtStorage *TxtStorageSession) AddRep(val *big.Int) (*types.Transaction, error) {
	return _TxtStorage.Contract.AddRep(&_TxtStorage.TransactOpts, val)
}

// AddRep is a paid mutator transaction binding the contract method 0x7b295c51.
//
// Solidity: function AddRep(val int256) returns()
func (_TxtStorage *TxtStorageTransactorSession) AddRep(val *big.Int) (*types.Transaction, error) {
	return _TxtStorage.Contract.AddRep(&_TxtStorage.TransactOpts, val)
}

// GetData is a paid mutator transaction binding the contract method 0x76b8e528.
//
// Solidity: function GetData() returns()
func (_TxtStorage *TxtStorageTransactor) GetData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStorage.contract.Transact(opts, "GetData")
}

// GetData is a paid mutator transaction binding the contract method 0x76b8e528.
//
// Solidity: function GetData() returns()
func (_TxtStorage *TxtStorageSession) GetData() (*types.Transaction, error) {
	return _TxtStorage.Contract.GetData(&_TxtStorage.TransactOpts)
}

// GetData is a paid mutator transaction binding the contract method 0x76b8e528.
//
// Solidity: function GetData() returns()
func (_TxtStorage *TxtStorageTransactorSession) GetData() (*types.Transaction, error) {
	return _TxtStorage.Contract.GetData(&_TxtStorage.TransactOpts)
}

// GetDataFree is a paid mutator transaction binding the contract method 0xe27e9160.
//
// Solidity: function GetDataFree() returns()
func (_TxtStorage *TxtStorageTransactor) GetDataFree(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStorage.contract.Transact(opts, "GetDataFree")
}

// GetDataFree is a paid mutator transaction binding the contract method 0xe27e9160.
//
// Solidity: function GetDataFree() returns()
func (_TxtStorage *TxtStorageSession) GetDataFree() (*types.Transaction, error) {
	return _TxtStorage.Contract.GetDataFree(&_TxtStorage.TransactOpts)
}

// GetDataFree is a paid mutator transaction binding the contract method 0xe27e9160.
//
// Solidity: function GetDataFree() returns()
func (_TxtStorage *TxtStorageTransactorSession) GetDataFree() (*types.Transaction, error) {
	return _TxtStorage.Contract.GetDataFree(&_TxtStorage.TransactOpts)
}

// SetHashData is a paid mutator transaction binding the contract method 0xb674f79a.
//
// Solidity: function SetHashData(hd string) returns()
func (_TxtStorage *TxtStorageTransactor) SetHashData(opts *bind.TransactOpts, hd string) (*types.Transaction, error) {
	return _TxtStorage.contract.Transact(opts, "SetHashData", hd)
}

// SetHashData is a paid mutator transaction binding the contract method 0xb674f79a.
//
// Solidity: function SetHashData(hd string) returns()
func (_TxtStorage *TxtStorageSession) SetHashData(hd string) (*types.Transaction, error) {
	return _TxtStorage.Contract.SetHashData(&_TxtStorage.TransactOpts, hd)
}

// SetHashData is a paid mutator transaction binding the contract method 0xb674f79a.
//
// Solidity: function SetHashData(hd string) returns()
func (_TxtStorage *TxtStorageTransactorSession) SetHashData(hd string) (*types.Transaction, error) {
	return _TxtStorage.Contract.SetHashData(&_TxtStorage.TransactOpts, hd)
}

// TransferData is a paid mutator transaction binding the contract method 0x20ccbe68.
//
// Solidity: function TransferData(dude address) returns()
func (_TxtStorage *TxtStorageTransactor) TransferData(opts *bind.TransactOpts, dude common.Address) (*types.Transaction, error) {
	return _TxtStorage.contract.Transact(opts, "TransferData", dude)
}

// TransferData is a paid mutator transaction binding the contract method 0x20ccbe68.
//
// Solidity: function TransferData(dude address) returns()
func (_TxtStorage *TxtStorageSession) TransferData(dude common.Address) (*types.Transaction, error) {
	return _TxtStorage.Contract.TransferData(&_TxtStorage.TransactOpts, dude)
}

// TransferData is a paid mutator transaction binding the contract method 0x20ccbe68.
//
// Solidity: function TransferData(dude address) returns()
func (_TxtStorage *TxtStorageTransactorSession) TransferData(dude common.Address) (*types.Transaction, error) {
	return _TxtStorage.Contract.TransferData(&_TxtStorage.TransactOpts, dude)
}

// TxtStorageReputationGivenIterator is returned from FilterReputationGiven and is used to iterate over the raw logs and unpacked data for ReputationGiven events raised by the TxtStorage contract.
type TxtStorageReputationGivenIterator struct {
	Event *TxtStorageReputationGiven // Event containing the contract specifics and raw log

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
func (it *TxtStorageReputationGivenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxtStorageReputationGiven)
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
		it.Event = new(TxtStorageReputationGiven)
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
func (it *TxtStorageReputationGivenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxtStorageReputationGivenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxtStorageReputationGiven represents a ReputationGiven event raised by the TxtStorage contract.
type TxtStorageReputationGiven struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReputationGiven is a free log retrieval operation binding the contract event 0xe645ec368e9f33f328b76bbcf9d56855e7bfe7d5892c74e354aadfc80685257a.
//
// Solidity: e ReputationGiven(sender address)
func (_TxtStorage *TxtStorageFilterer) FilterReputationGiven(opts *bind.FilterOpts) (*TxtStorageReputationGivenIterator, error) {

	logs, sub, err := _TxtStorage.contract.FilterLogs(opts, "ReputationGiven")
	if err != nil {
		return nil, err
	}
	return &TxtStorageReputationGivenIterator{contract: _TxtStorage.contract, event: "ReputationGiven", logs: logs, sub: sub}, nil
}

// WatchReputationGiven is a free log subscription operation binding the contract event 0xe645ec368e9f33f328b76bbcf9d56855e7bfe7d5892c74e354aadfc80685257a.
//
// Solidity: e ReputationGiven(sender address)
func (_TxtStorage *TxtStorageFilterer) WatchReputationGiven(opts *bind.WatchOpts, sink chan<- *TxtStorageReputationGiven) (event.Subscription, error) {

	logs, sub, err := _TxtStorage.contract.WatchLogs(opts, "ReputationGiven")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxtStorageReputationGiven)
				if err := _TxtStorage.contract.UnpackLog(event, "ReputationGiven", log); err != nil {
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

// TxtStorageRequestDataIterator is returned from FilterRequestData and is used to iterate over the raw logs and unpacked data for RequestData events raised by the TxtStorage contract.
type TxtStorageRequestDataIterator struct {
	Event *TxtStorageRequestData // Event containing the contract specifics and raw log

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
func (it *TxtStorageRequestDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxtStorageRequestData)
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
		it.Event = new(TxtStorageRequestData)
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
func (it *TxtStorageRequestDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxtStorageRequestDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxtStorageRequestData represents a RequestData event raised by the TxtStorage contract.
type TxtStorageRequestData struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestData is a free log retrieval operation binding the contract event 0x70b912fa029fe2496dbbf0edd4740f19ff757041675de8572cdd263fa630ed93.
//
// Solidity: e RequestData(sender address)
func (_TxtStorage *TxtStorageFilterer) FilterRequestData(opts *bind.FilterOpts) (*TxtStorageRequestDataIterator, error) {

	logs, sub, err := _TxtStorage.contract.FilterLogs(opts, "RequestData")
	if err != nil {
		return nil, err
	}
	return &TxtStorageRequestDataIterator{contract: _TxtStorage.contract, event: "RequestData", logs: logs, sub: sub}, nil
}

// WatchRequestData is a free log subscription operation binding the contract event 0x70b912fa029fe2496dbbf0edd4740f19ff757041675de8572cdd263fa630ed93.
//
// Solidity: e RequestData(sender address)
func (_TxtStorage *TxtStorageFilterer) WatchRequestData(opts *bind.WatchOpts, sink chan<- *TxtStorageRequestData) (event.Subscription, error) {

	logs, sub, err := _TxtStorage.contract.WatchLogs(opts, "RequestData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxtStorageRequestData)
				if err := _TxtStorage.contract.UnpackLog(event, "RequestData", log); err != nil {
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
