/*
    File Name: TxtStorage.sol
    Author: Elio Decolli    (eliodecolli@gmail.com)
    Purpose: Serves as a metadata storage for a specified item available on a hybrid blockchain.
    Last Updated: 5/16/2019
 */


 // Owner's address is the TxtStreamer address and the streamer's owner is the public wallet address of the user.
 // RREGULLO TRANSFERIMIN !!!  [X] ?


pragma solidity 0.5.8;
pragma experimental ABIEncoderV2;

import "./TxtStreamer.sol";
import "./Global.sol";

contract TxtStorage{

    address private owner;
    string private hash_data;
    string private name;
    int256 private Rating;
    uint256 private total_price;
    uint256 private Price;
    TxtStreamer private Owner;
    bool public IsAvailable;
    GlobalStorage m_global;
    uint256 private good_rep;
    uint256 private bad_rep;

    bool private open_access = false;

    mapping(address => bool) private AllowedCustomers;
    mapping(address => uint256) private CustomersCodes;

    modifier mustBeAllowed(){
        require(AllowedCustomers[msg.sender] == true, "The specified address is not allowed to access the storage.");
        _;
    }

    modifier ownerOnly(){
        require(msg.sender == owner, "Invalid access credentials.");
        _;
    }

    event RequestData(address sender);   // This waits for approval by the owner.
    event RetrieveData(address sender);   // This retrives the DL link and code automatically.
    event ReputationGiven(bool negative, uint256 amount);
    event RequestGenerateNewLink(address sender, uint256 code);

    constructor(string memory name_, uint256 price, string memory hash_, address gl, address payable _owner) public{
        owner = _owner;
        name = name_;
        Price = price;
        total_price = 0;
        hash_data = hash_;
        IsAvailable = true;

        Owner = TxtStreamer(owner);
        m_global = GlobalStorage(gl);
    }


    // Tokens experiment

    function GetRep() public view returns(uint256, uint256){
        return (good_rep, bad_rep);
    }

    function PlusGood(uint256 val) public {
        require(msg.sender == address(m_global));
        good_rep += val;
        bad_rep -= val;

        emit ReputationGiven(false, val);
    }

    function PlusBad(uint256 val) public{
        require(msg.sender == address(m_global));
        good_rep -= val;
        bad_rep += val;

        emit ReputationGiven(true, val);
    }

    // End of tokens experiment

    function SetAccessType(bool access) public ownerOnly{
        open_access = access;
    }

    function TracebackOwner() private view returns(address payable){
        return Owner.GetOwner();
    }

    function GetGlobal() public view returns(GlobalStorage){
        return m_global;
    }

    function SetStatus(bool status) public{
        require(msg.sender == owner, "You don't have sufficient priviledges to complete this operation.");
        IsAvailable = status;
    }

    function GetHashData() public view returns (string memory) {
        return hash_data;
    }

    function GetName() public view returns (string memory) {
        return name;
    }

    function SetHashData(string memory hd) public ownerOnly {
        require(IsAvailable == true, "Current storage is not available anymore.");

        hash_data = hd;
    }
    

    // Ty Stackoverflow :)
    function toBytes(address a) private view returns (bytes memory b){
        assembly {
            let m := mload(0x40)
            mstore(add(m, 20), xor(0x140000000000000000000000000000000000000000, a))
            mstore(0x40, add(m, 52))
            b := m
        }
    }

    ////
    // Request data access to this storage.
    function ReqeustDataAccess() public {
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(Rating > 0, "This document has a negative rating, therefore it's free.");

        if(!open_access)
        {
            emit RequestData(msg.sender);
        }
        else
        {
            emit RetrieveData(msg.sender);   // This must redirect the GoServer to call TransferData() or RejectTransfer()!
        }
    }

    ////
    // Accept transaction and send the user the download link as well as the code to access it.
    // It's recommended to use this function only once, to allow the user to access the data.
    function TransferData(address dude, uint256 code, string memory dlLink) public {
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(msg.sender == owner);

        AllowedCustomers[dude] = true;   // Here we are bypassing the hash-check to make sure what kind of contract is trying to call our code.
        CustomersCodes[dude] = code;
        TxtStreamer xStreamer = TxtStreamer(dude);
        xStreamer.Tell(true, code, dlLink);
    }

    function GenerateNewLink() public mustBeAllowed {
        emit RequestGenerateNewLink(msg.sender, CustomersCodes[msg.sender]);
    }

    ////
    // Reject the current transfer request and tell the sender about this.
    function RejectTransfer(address dude) public {
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(msg.sender == owner);

        TxtStreamer xStreamer = TxtStreamer(dude);
        xStreamer.Tell(false, 0, "");
    }

    ////
    // We don't deposit ether in smart contracts. This function logs the current transaction as "pending" and returns the address of the owner, who in turn accepts the Ether.
    function DepositEther(uint256 val) public returns(address payable){
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(val == Price, "Specified amount is not valid.");
        require(msg.sender != owner, "Invalid request source.");

        return TracebackOwner();
    }

}