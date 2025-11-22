// ...existing code...
## 🧭 My Protocol Security Learning Journey

> “(ps: I've been preparing product manager work recently, so this was delayed) I used to build security tools at the application layer, which felt like treating symptoms; now I want to learn how to protect the system's 'immune system' as a whole.”

### Why shift to protocol security?

When building wallet risk detection tools, I focused on analyzing on-chain transaction patterns. Over time I realized many application-layer risks—like the recent Balancer incident I analyzed—stem from the underlying protocol design and interactions. This sparked deeper curiosity:

- When we interact with a DApp in the frontend, how do the execution layer and consensus layer coordinate trust via engine APIs?
- How does EIP-1559's dynamic fee mechanism strengthen network security from an economic perspective?
- Does EIP-4844's introduction of blob space for L2s, while improving scalability, introduce new cryptographic attack surfaces?

I found that answering these questions—understanding and safeguarding the most fundamental, global trust layers—is more intellectually challenging and rewarding than identifying individual risk patterns at the application layer.

### My practices and exploration

To build systematic understanding, I started the following practices:

1.  Run a node: operated a geth node on the Sepolia testnet to directly interact with the blockchain's "heartbeat".
2.  Study EIPs: read EIP-1559 and EIP-4844 in depth, analyzing the core problems they solve and the new security considerations they introduce.
3.  Cryptography fundamentals: reviewed the roles of hashes, Merkle trees, ECDSA, and BLS signatures within protocol design.

### Next steps

This is only the beginning. I hope to turn this passion for fundamentals into substantive contributions to protocol security at organizations like the Ethereum Foundation. My goal is to grow from a user and observer of tools into a maintainer and builder of core ecosystem infrastructure.

*(Last updated: November 2025)*