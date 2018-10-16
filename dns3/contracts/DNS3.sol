/** Copyright 2018 Wolk Inc.
* This file is part of the DNS3 project for ETHSanFransisco.
*
* The DNS3 project is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The Plasmacash library is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
*
* You should have received a copy of the GNU Lesser General Public License
* along with the plasmacash library. If not, see <http://www.gnu.org/licenses/>.
*/

/**
 * @title  DNS3
 * @author Michael Chung (michael@wolk.com)
 * @dev Rough DNS3 with IPFS Multihash support.
 */

import "./SafeMath.sol";
pragma solidity ^0.4.25;

contract DNS3 {
    event DomainClaim(string Domain, bytes32 indexed DomainHash, address indexed Owner, uint256 Deposit);
    event DomainTransfer(bytes32 indexed DomainHash, address indexed originalOwner, address indexed newOwner, uint256 newDeposit);
    event DomainRelease(bytes32 indexed DomainHash, address indexed Owner, uint256 halfRefund);
    event DomainsUpdate(uint256 indexed blkNum, uint8 hashFunction,uint8 size, bytes32 digest);
    event ZoneUpdate(bytes32 indexed DomainHash, uint8 hashFunction,uint8 size, bytes32 digest);
    event ApprovedBuyer(bytes32 indexed DomainHash, address indexed approvedBuyer);

    struct Multihash {
        uint8 hashFunction;
        uint8 size;
        bytes32 digest;
    }

    using SafeMath for uint256;
    mapping(bytes32=>string) public DomainName;
    mapping(bytes32=>address) public OwnedDomain;
    mapping(bytes32=>address) public Buyer;
    mapping(bytes32=>uint256) DomainDeposit;
    mapping(bytes32=>Multihash) public ZoneHash;
    mapping(uint256=>Multihash) public PublishedDomains;


    uint256 public currentBlkNum;
    address public authority;

    modifier isAuthority() {
        require(msg.sender == authority);
        _;
    }

   constructor() public {
        authority = msg.sender;
        currentBlkNum = 0;
    }

    function registerDomain(string _domain) public payable returns (bool) {
        bytes32 domainHash = keccak256(abi.encodePacked(_domain));
        require(OwnedDomain[domainHash] == 0, 'Domain already registered');
        DomainName[domainHash] = _domain;
        OwnedDomain[domainHash] = msg.sender;
        emit DomainClaim(_domain, domainHash, msg.sender, msg.value);
        DomainDeposit[domainHash] = DomainDeposit[domainHash].add(msg.value);
        return true;
    }

    function approvedBuyer(bytes32 domainHash, address _canBuy) public returns (bool) {
        require(OwnedDomain[domainHash] == msg.sender, 'Unauthorized approval');
        Buyer[domainHash] = _canBuy;
        emit ApprovedBuyer(domainHash, _canBuy);
        return true;
    }

    function cancelBuyer(bytes32 domainHash) public returns (bool) {
        require(OwnedDomain[domainHash] == msg.sender, 'Unauthorized transfer');
        Buyer[domainHash] = 0x0;
        return true;
    }

    function acquireDomain(bytes32 domainHash) public payable returns (bool) {
        require(Buyer[domainHash] == msg.sender && DomainDeposit[domainHash].mul(4) >= msg.value, 'Invalid Purchase');
        address seller = OwnedDomain[domainHash];

        uint256 newDeposit = msg.value.div(4); //25% were kept in contract
        uint256 revenue = DomainDeposit[domainHash].add(msg.value).sub(newDeposit);

        DomainDeposit[domainHash] = 0;
        seller.transfer(revenue);

        OwnedDomain[domainHash] = msg.sender;
        DomainDeposit[domainHash] = newDeposit;

        delete Buyer[domainHash];
        emit DomainTransfer(domainHash, seller, msg.sender, newDeposit);
        return true;
    }

    function releaseDomain(bytes32 domainHash) public returns (bool) {
        require(OwnedDomain[domainHash] == msg.sender, 'Unauthorized transfer');
        address exiter = OwnedDomain[domainHash];
        uint256 halfRefund = DomainDeposit[domainHash].div(2);
        delete OwnedDomain[domainHash];
        exiter.transfer(halfRefund);
        authority.transfer(halfRefund);
        emit DomainRelease(domainHash, msg.sender, halfRefund);
        return true;
    }

    function submitZone(bytes ipfsHashByte, bytes32 domainHash) public returns (bool) {
        require(OwnedDomain[domainHash] == msg.sender, 'Unauthorized zoneUpdate');
        Multihash memory ipfsHash = _setMultihash(ipfsHashByte);
        ZoneHash[domainHash] = ipfsHash;
        emit ZoneUpdate(domainHash, ipfsHash.hashFunction, ipfsHash.size, ipfsHash.digest);
        return true;
    }


    function getZone(bytes32 domainHash) public view returns (bytes32 digest, uint8 hashFunction,uint8 size) {
        require(OwnedDomain[domainHash] != 0, "Domain not registered");
        Multihash memory ipfsHash = ZoneHash[domainHash];
        return (ipfsHash.digest, ipfsHash.hashFunction, ipfsHash.size);
    }

    function updateDomains(bytes ipfsHashByte, uint256 blkNum) public isAuthority {
         //require(PublishedDomain[blkNum] != 0, "Multihash already set");
         //require(currentBlkNum.add(1) != PublishedDomain[blkNum], "Multihash already set");
         Multihash memory ipfsHash = _setMultihash(ipfsHashByte);
         PublishedDomains[blkNum] = ipfsHash;
         emit DomainsUpdate(blkNum, ipfsHash.hashFunction, ipfsHash.size, ipfsHash.digest);
    }

    function _setMultihash(bytes ipfsHashByte) private pure returns (Multihash memory) {
        bytes32 digest;
        assembly { digest := mload(add(ipfsHashByte,34)) }
        Multihash memory ipfsHash;
        ipfsHash.hashFunction = uint8(ipfsHashByte[0]);
        ipfsHash.size = uint8(ipfsHashByte[1]);
        ipfsHash.digest = digest;
        return ipfsHash;
    }
}
