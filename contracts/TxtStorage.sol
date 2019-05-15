/*
    File Name: TxtStorage.sol
    Author: Elio Decolli    (eliodecolli@gmail.com)
    Purpose: Serves as a metadata storage for a specified item available on a hybrid blockchain.
 */


 // Owner's address is the TxtStreamer address and the streamer's owner is the public wallet address of the user.
 // RREGULLO TRANSFERIMIN !!!  [X]


pragma solidity 0.5.8;
pragma experimental ABIEncoderV2;

import "./TxtStreamer.sol";
import "./Global.sol";

contract TxtStorage{
    
    struct Customer{
        address sender;
        bool pending;
        uint256 price_paid;
        TxtStreamer stream;
    }

    address private owner;
    string private hash_data;
    string private name;

    int256 private Rating;

    uint256 private total_price;

    event RequestData(address sender, uint rId);   // This waits for approval by the owner.
 
    event RetrieveData(address sender, uint rId);   // This retrives the DL link and code automatically.

    event ReputationGiven(bool negative, uint256 amount);


    uint256 private Price;

    TxtStreamer private Owner;
    
    mapping(address => Customer) private Customers;


    bool public IsAvailable;

    GlobalStorage m_global;


    // Tokens experiment
    
    uint256 private good_rep;
    uint256 private bad_rep;

    bool private open_access = false;

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

    function RequestAccess(address from, uint rId) public {
        if(open_access){
            emit RetrieveData(from, rId);
        } else {
            emit RequestData(from, rId);
        }
    }

    // End of tokens experiment

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

    function SetHashData(string memory hd) public {
        require(msg.sender == owner);
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

    function IsUserDone(address user) private view returns(bool){
         return Customers[user].pending;
    }

    ////
    // Leave a positive feedback. :)
    function IncreaseRep(uint256 rep) public {
        require(msg.sender != owner);
        require(IsUserDone(msg.sender));
        require(IsAvailable == true, "Current storage is not available anymore.");
        //require(multiplier == -1 || multiplier == 1);

        //TxtStreamer txtSr = TxtStreamer(txtStreamer);
        //Rating += (txtSr.GetCurrentReputation() / 20) * multiplier;

        m_global.IncrementTotalRep(rep);
        emit ReputationGiven(false, rep);
        return;
    }

    ////
    // Get the current reputation of this storage.
    function GetReputation() public view returns(int256){
        return Rating;
    }

    ////
    // Request data access for free to this storage.
    function GetDataFree() public {
        require(Rating < 0, "This document is not for free.");
        require(!IsUserDone(msg.sender), "You've already made a request for this document.");
        require(msg.sender != owner, "Invalid request source.");
        require(IsAvailable == true, "Current storage is not available anymore.");

        Customers[msg.sender].sender = msg.sender;
        Customers[msg.sender].pending = true;
        Customers[msg.sender].stream = TxtStreamer(msg.sender);
        emit RequestData(msg.sender, 1);
    }

    ////
    // Request data access to this storage.
    function GetData() public {
        //require(msg.value >= Price, "Price is too low.");
        require(!IsUserDone(msg.sender), "You've already made a request for this document.");
        require(msg.sender != owner);
        require(IsAvailable == true, "Current storage is not available anymore.");

        require(Rating > 0, "This document has a negative rating, therefore it's free.");

        emit RequestData(msg.sender, 1);
    }

    ////
    // Accept transaction and send the user the download link as well as the code to access it.
    function TransferData(address dude, uint256 code, string memory dlLink) public {
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(msg.sender == owner);
        require(!IsUserDone(dude));

        //owner.transfer(Customers[dude].price_paid);
        Customers[dude].pending = false;
        Customers[dude].stream.Tell(true, code, dlLink);
    }

    ////
    // Reject the current transfer request and tell the sender about this.
    function RejectTransfer(address dude) public {
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(msg.sender == owner);
        require(!IsUserDone(dude));

        //owner.transfer(Customers[dude].price_paid);
        Customers[dude].pending = false;
        Customers[dude].stream.Tell(false, 0, "");
    }

    ////
    // We don't deposit ether in smart contracts. This function logs the current transaction as "pending" and returns the address of the owner, who in turn accepts the Ether.
    function DepositEther(uint256 val) public returns(address payable){
        require(IsAvailable == true, "Current storage is not available anymore.");
        require(val == Price, "Specified amount is not valid.");
        require(msg.sender != owner, "Invalid request source.");

        Customers[msg.sender].price_paid = val;
        Customers[msg.sender].sender = msg.sender;
        Customers[msg.sender].pending = true;
        Customers[msg.sender].stream = TxtStreamer(msg.sender);
        return TracebackOwner();
    }

}