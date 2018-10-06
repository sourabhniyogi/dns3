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

pragma solidity ^0.4.25;

import "./SafeMath.sol";

contract DNS3 {
    event DomainClaim(string Domain, bytes32 indexed DomainHash, address indexed Owner, uint256 Deposit);
    event DomainTransfer(bytes32 indexed DomainHash, address indexed originalOwner, address indexed newOwner);
    event DomainRelease(bytes32 indexed DomainHash, address indexed Owner, uint256 halfRefund);
    event DomainUpdate(uint256 indexed blkNum, bytes32 digest, uint8 hashFunction,uint8 size);

    struct Multihash {
        bytes32 digest;
        uint8 hashFunction;
        uint8 size;
    }

    using SafeMath for uint256;
    mapping(bytes32=>address) OwnedDomain;
    mapping(bytes32=>uint256) DomainDeposit;
    mapping(uint256=>Multihash) public PublishedDomain;

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
        OwnedDomain[domainHash] = msg.sender;
        emit DomainClaim(_domain, domainHash, msg.sender, msg.value);
        DomainDeposit[domainHash] = DomainDeposit[domainHash].add(msg.value);
        return true;
    }


    function transferDomain(bytes32 domainHash, address newOwner) public returns (bool) {
        require(OwnedDomain[domainHash] == msg.sender && newOwner != 0x0, 'Unauthorized transfer');
        OwnedDomain[domainHash] = newOwner;
        emit DomainTransfer(domainHash, msg.sender, newOwner);
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

    function updateDomain(bytes ipfsHashByte, uint256 blkNum) public isAuthority {
         //require(PublishedDomain[blkNum] != 0, "Multihash already set");
         //require(currentBlkNum.add(1) != PublishedDomain[blkNum], "Multihash already set");
         uint16 hashPart;
         bytes32 digest;
         assembly {
             hashPart := div(mload(add(ipfsHashByte, 32)), exp(256, 30))
             digest := mload(add(ipfsHashByte, 2))
         }
         Multihash memory ipfsHash;
         ipfsHash.digest = digest;
         ipfsHash.hashFunction = uint8(hashPart / 256);
         ipfsHash.size = uint8(hashPart % 256);
         PublishedDomain[blkNum] = ipfsHash;
         emit DomainUpdate(blkNum, ipfsHash.digest, ipfsHash.hashFunction, ipfsHash.size);
    }
}
