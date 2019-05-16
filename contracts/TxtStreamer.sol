/*
    File Name: TxtStreamer.sol
    Author: Elio Decolli    (eliodecolli@gmail.com)
    Purpose: Representation of an individual on the chain.
    Last Updated: 5/16/2019
 */

pragma solidity 0.5.8;
pragma experimental ABIEncoderV2;

import "./Global.sol";
import "./TxtStorage.sol";


contract TxtStreamer{

    address payable private owner;
    string private _name;

    uint256 private _reputation;

    GlobalStorage m_global;


    event TxtResponse(bool mResponse, uint256 code, string dlLink);
    event UserReputationChanged(uint256 val, address from);

    mapping(address => string) private DeployedStorages;
    uint256 private totalStorages;

    constructor(address global) public{
        owner = msg.sender;
        _reputation = 0;
        totalStorages = 0;

        if(global != address(0)){
            m_global = GlobalStorage(global);
        }
    }

    function GetOwner() public view returns(address payable){
        return owner;
    }

    function Tell(bool resp, uint256 code, string memory dlLink) public {
        require(msg.sender != owner);

        emit TxtResponse(resp, code, dlLink);
    }

    function GetCurrentReputation() public view returns(uint256){
        return _reputation;
    }

    function SubReputation(uint256 val) public payable{
        require(msg.sender != owner);
        require(msg.value >= 10000000000000, "Minimal value in order to give reputation is 1e13 wei.");

        _reputation -= val;
        emit UserReputationChanged(val, msg.sender);
    }

    function GiveReputation(uint256 val) public payable{
        require(msg.sender != owner);
        require(msg.value >= 10000000000000, "Minimal value in order to give reputation is 1e13 wei.");

        _reputation += val;
        emit UserReputationChanged(val, msg.sender);
    }

    function GetAllowance() private view returns(bool){
        if(_reputation > m_global.GetMedian()){
            return true;
        }
        else{
            if(_reputation >= totalStorages){
                return true;
            }
            else{
                return false;
            }
        }
    }

    function AddStorage(uint256 price, string memory fileHash, string memory name) public returns(address){
        require(msg.sender == owner, "You don't have sufficient priviledges to perform this operation.");
        //require(place == address(0), "Invalid address specified.");
        require(bytes(name).length > 0, "Invalid name specified.");
        require(bytes(fileHash).length > 0, "Invalid file hash specified.");
        //require(GetAllowance(), "You don't have sufficient reputation in order to upload a new material.");
        
        TxtStorage m_storage = new TxtStorage(name, price, fileHash, address(m_global), 0xf5F3686F4Ef910d5aE03D55Fe2Bf93BB0a6EE5Ba);
        address place = address(m_storage);

        DeployedStorages[place] = name;
        totalStorages++;
        _reputation += 1;
        
        m_global.UploadContent(name, price, fileHash);
        return place;
    }

}