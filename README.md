# DNS3.0: Trustless DNS, or "7.7.7.7 (Layer 2 Plasma Chain) | 7.7.6.6 (Layer 1 DNS Blockchain) & Handshake > 8.8.8.8 (Google DNS) & Godaddy & ICANN"

This project aims to explore working prototypes of parts of a DNS 3.0 stack, aiming for trustless decentralized DNS systems integrating with Handshake's decentralized TLD.  

DNS is managed with _Zone_ files like this one, typically by a web master updating DNS records in a website UI in Godaddy / Cloudflare / Google / ... DNS Manager, so that when users visit a name-friendly url like "http://dns3.wolk.com" their browser will map them to an IP `104.154.155.233`

```
;; wolk.com zone entries
;; CNAME Records
dns3.wolk.com.   3600    IN  A   104.154.155.233
smt.wolk.com. 1   IN  A   35.224.4.165
blog.wolk.com.  1   IN  A   52.4.38.70
```

Behind the scenes, when users browsers ask for "dns3.wolk.com" there is:
 1. a user's _trusted_ DNS provider (e.g. 8.8.8.8 run by Google, or your local monopolist ISP) that figures out the name server for `wolk.com` and asks that name server to look up the value for `dns3`
 2. a name server for "wolk.com" hosted by a _trusted_ Zone file database editor
 3. a root "Top level domain" (or "TLD") server (e.g. ".com, .edu, .org"), which have historically been managed with a dozen or so organizations (see http://www.root-servers.org/)
People learn about (1) when setting up their new Internet connection, learn about (2) when setting up a new web site, but
almost no one is aware of (3), but all of this comes together in a few hundred milliseconds.   Can the _trusted_ systems that manage (1) DNS lookup (2) database zone records (3) TLD root servers be replaced with completely _trustless_ DNS protocols?  

We believe the answer is **YES**.   

Recently, the lowest level (3) has been developed into a new blockchain protocol by handshake.org, where new TLDs (e.g. `.hacker`) may be secured.    This DNS3 project is a prototype that aims to show how the more familiar layers of (1) name look up + (2) zone files of domains can managed in a _trustless_ way with 2 different State Models:
* _Ethereum State Model 1_: A DNS3 Smart contract manages Zone file hashes, where owners register domains, update domain zone files, and can sell their domain.  At the core, a zone file is update recorded with a transaction submitted to the Ethereum contract:
```
> eth.submitZone('eth.hacker', `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`)
```
* _DNS3 State Model 2_: A DNS3 Blockchain manages domain-zone file hashes in a *Sparse Merkle Tree*:
```
> dns3.submitZone('eth.hacker', `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`)
```
In both model, zone files are held in decentralized storage (IPFS), where a zone file is uniquely retrievable and verifiable by their zone file hash [this snippet on IPFS](https://cloudflare-ipfs.com/ipfs/QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89) and move from state `QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89` to state `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J` when the DNS entries for a domain changes to have a new entry like:
 ```
 www.eth.hacker.   3600    IN  A   104.154.155.234
 ```

Each model is described below.

## Model 1: Use Ethereum DNS3 Smart Contract to Read/Write DNS Info

A contract posted on Ethereum holds the latest zone hash for every domain.  The owner of `eth.hacker` updates the zone hash entry in this contract.    When users look up `www.eth.hacker`, a call is made by the `dns3` node to the `dns3.sol` Ethereum contract:

```
> eth.getZone('eth.hacker')
{"QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J"}
```
The Ethereum Smart Contract returns back a 32-byte hash that is used by the local `dns3` node to retrieve the zone file from decentralized storage (an IPFS hash).  

To prevent domain squatting (common to NameCoin, ENS, and the current Domain Name Registration system), we adapt a "Radical Markets" technique proposed by Weyl and Posner:
* domain name registrants specify a _sale price_ when they register.  When they do so, they commit to paying a fixed percentage of that sale price every 1MM blocks.  Otherwise, anyone can pay that _sale price_ and secure the rights to the domain.  A grace period of 7 days is offered to ensure that a transition can be smooth, or for the current owner to increase his sale price to override the transfer, but the override cost must be at least 10x higher.      
* Example: Alice purchases a new domain `dns3.hacker` for a price of .1 ETH but sets her _sale price_ to 10ETH.  Bob sees `dns3.hacker` and submits a `requestPurchase("dns3.hacker")` transaction for 10ETH to get it.   If Alice does nothing, after 7 days, Bob (or anyone) can finalize the purchase by calling `finalizePurchase("dns3.hacker")`.  If Alice pays 100ETH to keep the `dns3.hacker` domain, she (and only she) can call `challengePurchase("dns3.hacker")`

Implementation:
1. `DNS3.sol` - the Ethereum contract to manage the key operations: `getZone`, `submitZone`

2. Run a server locally from `github.com/miekg/dns`.  When a DNS lookup is performed, 

3. `submitzone.html`: Javascript-based UI for domain owner to submit new zone by posting `submitZone` transaction


## Model 2: DNS3 Blockchain with SMT roots

In this model, a specialized DNS3 blockchain replaces the Ethereum blockchain, with its own `DNSCoin` currency.  Anyone can run a DNS3 node and earn DNSCoin for:
1. storing zone files submitted by _registered_ owners of a domain
2. responding to _signed_ DNS Queries
Anyone can spend DNSCoin to:
1. store zone files
2. get DNS Info
DNS Zone storage is very special case of File storage (tackled by IPFS/Filecoin, Sia, Storj, among others).

DNS3 nodes participate in a consensus protocol to package `submitZone` transactions into a block.  The situation with `register` is
 with two key state variables for 2 Sparse Merkle Trees:

1. the new root hash of a _Domain_ SMT, where the keys are 256-bit hashes of domain strings and the values contain
 * the latest state of a domain (e.g. `eth.hacker`)'s zone file;
 * the current owner

2. the new root hash of a _Balance_ SMT, where the keys are 160-bit addresses (just like Ethereum) and the values are the 32-byte `uint256` balances of each address.

Here are the core data structures for the DNS3 Blockchain:
```
type Block {
  DomainRoot  common.Hash    // key: 256-bit hashes of domains (eg. Keccak256("eth.hacker")); values: Domain
  BalanceRoot  common.Hash   // key: 160-bit address; values: uint256 Balances  
  Transaction []*Transaction // collation of 5 different transaction types
}

type Transaction {
  transactionType int       // see 5 below
  domain     common.Hash
  zoneHash   common.Hash
  sig        []byte
}

type Domain {
  Owner     common.Address
  SalePrice big.Int
}

// TransactionType
const (
  registerTransaction = 1          // adds new key in DomainRoot (implemented), reduces BalanceRoot
  submitZoneTransaction = 2        // updates value for existing key (implemented)
  requestPurchaseTransaction = 3   // not implemented
  finalizePurchaseTransaction = 4  // not implemented, decreases balanceRoot
  challengePurchaseTransaction = 5 // not implemented
  cashTransaction = 6              // not implemented
)
```

In our implementation thus far, we have focussed not on the permissionless consensus protocol nor on the cryptoeconomics of mining rewards, but solely on the state transition model for the _Domain_ SMT.  The actions of the _Domain_ SMT has several operations:
 * `register` will insert a new
 * `submitZone` will insert a new key-value pair into the SMT
 * `getZone` will read the value of the key from the SMT
We have left all implementations concerning `BalanceRoot` `requestPurchase`, `finalizePurchase`, `challengePurchase` transactions for future work.

### Why Trustless DNS?

As all decentralization enthusiasts know, the vast majority of people are happy using trusted DNS system, and will be until ... they aren't happy ... and then ... they won't be able to do anything about it.   There is massive risk that _trusted_ parties like this one will use their increasingly dominant stack to:
 1. connect Public DNS systems (https://en.wikipedia.org/wiki/Google_Public_DNS) to other elements of the stack (https://en.wikipedia.org/wiki/Google_Chrome#User_tracking_concerns) Chrome 70 by default to collect user data in the name of faster Internet [read: increased revenue] or Attack:
 2. connect Google DNS, Cloudflare to mask DNS entries for better web site security / user privacy [read: collect data]
or in malicious ways simple redirect DNS queries for direct financial gain.  

Currently, the security of services (HTTP, Mail, FTP, messaging protocols,  etc.) run under domains may be attacked by:
 1. gaining access to Godaddy/Google/Cloudflare DNS with a username / password and hijacking the zone file editor
 2. gaining access to whatever is running 8.8.8.8 name server
and running malicious services that trick users into thinking that they are interacting with the owner of the domain when they in fact are not.  (If `myetherwallet.com` or `etherscan.io` DNS were hacked, these centralized services would result in attacks where ETH would be stolen, fake transactions could be posted, etc.)
