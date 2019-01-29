/*
    File Name: Global.sol
    Author: Elio Decolli    (eliodecolli@gmail.com)
    Purpose: Representation of a global storage on the chain.
 */

pragma solidity 0.5.2;
pragma experimental ABIEncoderV2;

import "./TxtStorage.sol";
import "./TxtStreamer.sol";

contract GlobalStorage{

    uint256 private totalStorage;
    uint256 private median;
    uint256 private totalAccounts;

    uint256 private rep_median;
    uint256 private total_rep;

    uint256 private globalId;

    mapping(address => address[]) private supply;
    mapping(address => bool) private accounts_present;

    mapping(address => bool) private storage_present;

    event CountIncreased(address nContract);

    bytes32 dummy_streamer;

    constructor(uint256 gid) public{
        totalStorage = 0;
        median = 0;
        totalAccounts = 0;
        globalId = gid;

        TxtStreamer mDummy_streamer = new TxtStreamer(address(this));
        address m_addr = address(mDummy_streamer);
        assembly{
            dummy_streamer := extcodehash(m_addr)
        }
    }

    function IsStorage(address addr) private view returns(bool){
        return storage_present[addr];
    }

    function IncrementTotalRep(uint256 rep) public {
        require(IsStorage(msg.sender), "This function can be called only via a TxtStorage contract.");
        total_rep += rep;
    }

    function DecrementTotalRep(uint256 rep) public {
        require(IsStorage(msg.sender), "This function can be called only via a TxtStorage contract.");
        total_rep -= rep;
    }

    function GetGlobalId() public view returns(uint256){
        return globalId;
    }

    function IsAccountPresent(address acc) private view returns(bool){
        return accounts_present[acc];
    }

    function IsStreamer(address addr) public view returns(bool o_streamer){
        bytes32 o_data;
        assembly{
            o_data := extcodehash(addr)
        }
        o_streamer = (o_data == dummy_streamer);
    }

    function RegisterStreamer() public {
        require(IsStreamer(msg.sender), "Invalid source calling function RegisterStreamer().");
        require(!IsAccountPresent(msg.sender), "Streamer already registered.");

        accounts_present[msg.sender] = true;
    }

    /////
    // Deploys a new TxtStorage.
    function UploadContent(string memory name, uint256 price, string memory hash_) public {
        require(IsAccountPresent(msg.sender), "Caller has not been registered on this global.");

        TxtStorage stor = new TxtStorage(name, price, hash_, address(this), msg.sender);   // Or maybe just use DELEGATECALL in order to set Ownership to the caller ?
        address content = address(stor);
        supply[msg.sender].push(content);
        totalStorage++;
        storage_present[content] = true;

        emit CountIncreased(content);
    }

    function GetTotalStorage() public view returns(uint256){
        return totalStorage;
    }

    function GetMedian() public view returns(uint256){
        return totalStorage / totalAccounts;
    }
}