pragma solidity ^0.4.15;

contract Whilelist {
    address public owner;
    mapping (address => bool) public listed;
    event AddWhitelist(address addr);
    event DelWhitelist(address addr);

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function Whilelist() {
        owner = msg.sender;
    }

    function Add(address addr) onlyOwner {
        listed[addr] = true;
        AddWhitelist(addr);
    }

    function Del(address addr) onlyOwner {
        delete(listed[addr]);
        DelWhitelist(addr);
    }

    function isListed(address addr) constant returns (bool) {
        return listed[addr];
    }
}
