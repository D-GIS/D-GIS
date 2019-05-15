pragma solidity 0.6.0;

import "./Global.sol";

contract Token{

    uint256 private g_id;

    GlobalStorage private gs;

    uint256 private total_supply;

    event TokensTransferred(uint256 val, address from, address to);
    event SupplyChanged(uint256 nVal);

    mapping(address => uint256) accounts;

    constructor(address global) public{
        gs = GlobalStorage(global);
        g_id = gs.GetGlobalId();
    }

    function TransferTokens(address destination, uint256 value) public returns(bool){
        require(value <= accounts[msg.sender]);

        accounts[msg.sender] -= value;
        accounts[destination] += value;

        emit TokensTransferred(value, msg.sender, destination);
    }

}