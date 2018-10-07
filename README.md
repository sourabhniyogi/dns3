# DNS3.0: Trustless DNS with Ethereum, Local DNS and Radical Markets

Team:
* Michael Chung (michael@wolk.com)
* Sourabh Niyogi (sourabh@wolk.com)
* Rodney Witcher (rodney@wolk.com)

This DNS3.0 "hackathon" project develops a piece of a DNS 3.0 stack, aiming for trustless decentralized DNS systems integrating with Handshake's decentralized TLD.  

DNS is managed with _Zone_ files like this one, typically by a web master updating DNS records in a website UI in Godaddy / Cloudflare / Google / ... DNS Manager, so that when users visit a name-friendly url like "http://www.eth.hacker" their browser will map them to an IP `104.154.155.233`:

```
;; eth.hacker zone entries
;; A Records
www.eth.hacker.     3600   IN  A   104.154.155.233
ganache.eth.hacker. 1      IN  A   35.224.4.165
remix.eth.hacker.   1      IN  A   52.4.38.70
```

Behind the scenes, when users browsers ask for `www.eth.hacker` there is:
 1. a user's _trusted_ DNS provider (e.g. 8.8.8.8 run by Google, or your local monopolist ISP) that figures out the name server for `eth.hacker` and asks that name server to look up the value for `www`
 2. a name server for `eth.hacker` hosted by a _trusted_ Zone file database editor like Godaddy or Google DNS
 3. a root "Top level domain" (or "TLD") server (e.g. ".com, .edu, .org"), which have historically been managed with a dozen or so organizations (see http://www.root-servers.org/)
People learn about (1) when setting up their new Internet connection, learn about (2) when setting up a new web site, but almost no one is aware of (3).  However, all of this comes together in a few hundred milliseconds.  

Can the _trusted_ systems that manage (1) DNS lookup (2) database zone records (3) TLD root servers be replaced with completely _trustless_ DNS protocols?  We believe the answer is **YES**.   

Recently, the lowest level (3) has been developed into a new blockchain protocol by handshake.org, where new TLDs (e.g. `.hacker`) can be bought and queries resolved in a trustless way, replacing 12-13 truste parties.   Before the rise of Ethereum, a special purpose blockchain named `Namecoin` built a special purpose blockchain to handle (1) + (2), but it never really took off.  This DNS3 project is a prototype that aims to show how  (1) name look up + (2) zone files of domains can managed in a _trustless_ way with a local DNS server that connects to a DNS3.sol Smart contract that keeps the Zone file hashes in a decentralized storage backend (IPFS).

# Demonstration of DNS3

### Part 1: DNS3.sol Ethereum Smart Contract

The core idea is that core DNS zone data are kept in an Ethereum Smart Contract like `DNS3.sol`:

 https://rinkeby.etherscan.io/address/0x8116a77cf44457a455ffc24001c521ddeebc9606

Domain owners for new TLDs like `.hacker` will manage their DNS entries by:
 * registering their domain with `registerDomain(string _domain)`
 https://rinkeby.etherscan.io/tx/0xa17097bc57d65c2a1a2b3510cda36bf5390e07a9841adc756ba70078a9000730
 where a domain `eth.hacker` is represented on Ethereum with a domainHash like: `0xb63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b`

 * updating the zone record for the `domainHash` with `submitZone(bytes ipfsHashByte, bytes32 domainHash)`
 https://rinkeby.etherscan.io/tx/0x4d5eb503fb3fe8006728724fe94242b142285fb1971c6d7362749a9f3d8262eb
where the zone file hash `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J` is represented in a 34-byte `ipfsHashByte`

In this model, zone files are held in decentralized storage (IPFS), where a zone file is uniquely retrievable and verifiable by their zone _file hash_ such as this one:
```
https://cloudflare-ipfs.com/ipfs/QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89
```
When someone adds a new record like:
```
dev.eth.hacker.   3600    IN  A   35.77.66.55
```
the owner of `eth.hacker` updates the zone hash entry in this contract.

### Part 2: Local DNS resolution

In DNS3.0, when devices resolve DNS entries like `www.eth.hacker` or `dev.eth.hacker`, instead of resolving to 8.8.8.8 Trusted server, their local resolver will:
 1. read the latest zone file 32-byte hash with `getZone(bytes32 domainHash)` by  hashing `eth.hacker` into the `domainHash` and retrieving `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`
 2. do a IPFS fetch to get the ZONE file
```
https://cloudflare-ipfs.com/ipfs/QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89
```

We demonstrate a proof-of-concept showing that a local DNS server (`github.com/miekg/dns`)
can do local domain resolution:

```
$ go test DNSRequest
DNS3 Request:    dev.eth.hacker
  domain:        eth.hacker
  domainHash:    0xb63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b
DNS3.sol Call:   getZone(0xb63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b)
  ipfsHash:      QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J
IPFS Lookup:     https://cloudflare-ipfs.com/ipfs/QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J ... DONE
DNS3 Result:     35.77.66.55
PASS
```
Given a DNS3 Request of `dev.eth.hacker`, the above test computes the domain `eth.hacker`, calls `getZone(bytes32 domainHash)` with the Keccak hash of `eth.hacker` and gets back the latest zone file hash `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`.  

Then it does a HTTP GET to IPFS to get the latest Zone file, and then retrieves the latest value `35.77.66.55`!

When everyone runs local DNS Servers, we get _totally trustless DNS_!  


### ... One more thing: Radical Markets

We believe this trustless DNS3.0 approach can be adapted for new TLDs managed entirely with Handshake.  To reduce domain squatting (common to NameCoin, ENS, and the current Domain Name Registration system), we propose to adapt a "Radical Markets" technique proposed by Weyl and Posner in our DNS3.sol smart contract:
* domain name registrants specify a _sale price_ when they register.  When they do so, they commit to paying a fixed percentage of that sale price every 1MM blocks.  Otherwise, anyone can pay that _sale price_ and secure the rights to the domain.  A grace period of 7 days is offered to ensure that a transition can be smooth, or for the current owner to increase his sale price to override the transfer, but the override cost must be at least 10x higher.      
* Example: Alice purchases a new domain `dns3.hacker` for a price of .1 ETH but sets her _sale price_ to 10ETH.  Bob sees `dns3.hacker` and submits a `requestPurchase("dns3.hacker")` transaction for 10ETH to get it.   If Alice does nothing, after 7 days, Bob (or anyone) can finalize the purchase by calling `finalizePurchase("dns3.hacker")`.  If Alice pays 100ETH to keep the `dns3.hacker` domain, she (and only she) can call `challengePurchase("dns3.hacker")`

### Welcome to Trustless DNS3.0 with Ethereum!!!
