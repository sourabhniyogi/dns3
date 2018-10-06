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

Recently, the lowest level (3) has been developed into a new blockchain protocol by handshake.org, where new TLDs (e.g. `.hacker`) may be secured.   The blockchain (with its own currency, POW consensus later) maintains a provable authenticated data structure where the keys to the authenticated data structure access.  Transactions of { `BID`, `REVOKE`, ... }  are managed in a _trustless_

This is a prototype that aims to show how the more familiar layers of (1) name look up + (2) zone files of domains can managed in a trustless way:
1. Zone files are held in decentralized storage (IPFS), where a zone file is represented by file hashes [this snippet on IPFS](https://cloudflare-ipfs.com/ipfs/QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89) and move from state `QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89` to state `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`  When the DNS entries for a domain changes to have a new entry like:
 ```
 www.eth.hacker.   3600    IN  A   104.154.155.234
 ```
2. A Sparse Merkle Tree (SMT) data structure holds at the leaf: (a) the latest state of a domain (e.g. `eth.hacker`)'s zone file; (b) the current owner.   The SMT is a abstract data structure with 2^256 leaves indexed by 256-bit leaves; the keys are hashes of domain names.  
  * Owners can bid/win/transfer their domain (e.g. `eth.hacker`).  Following Handshake, this can be done with second-price auctions
  * Owners can update their domain's zone file (e.g. containing `www.eth.hacker`) by submitting a new transaction containing a valid zone hash:
```
> dns3.submitZone('eth.hacker', `QmWMhdVYpGYeS33BNNAXSNwbiAVCkaaiFgeSAAiywwZP3J`)
```

### Why Trustless DNS?

As all decentralization enthusiasts know, the vast majority of people are happy using trusted DNS system, and will be until ... they aren't happy ... and then ... they won't be able to do anything about.   There is massive risk that _trusted_ parties will use their increasingly dominant stack to:
 1. connect "8.8.8.8" to Chrome 70 by default to collect user data in the name of faster Internet [read: increased revenue] or Attack:
 2. buy Cloudflare to mask DNS entries for better web site security / user privacy [read: collect data]

Currently, the security of services (HTTP, Mail, FTP, messaging protocols,  etc.) run under domains may be attacked by:
 1. gaining access to Godaddy/Google/Cloudflare DNS with a username / password and hijacking the zone file editor
 2. gaining access to whatever is running 8.8.8.8 name server
and running malicious services that trick users into thinking that they are interacting with the owner of the domain when they in fact are not.  (If `myetherwallet.com` or `etherscan.io` were hacked, the ETH community would basically be destroyed.)
