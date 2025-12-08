# Technical Analysis: Post-Fusaka Prysm Client Stability Incident

## Executive Summary
On December 4, 2025, following the Fusaka network upgrade, Prysm consensus client operators encountered stability issues requiring immediate reconfiguration. This report provides an independent technical analysis of the incident, tracing the root cause to a fork-choice compatibility issue with Fusaka's PeerDAS implementation and explaining the emergency mitigation.

---

## 1. Incident Timeline & Impact
- **~Dec 4, 2025**: Prysm team issues PSA regarding mainnet stability.
- **Affected Version**: Prysm v7.0.0 (initial Fusaka-compatible release).
- **Emergency Mitigation**: Operators instructed to add `--disable-last-epoch-targets` flag.
- **Impact Scope**: Prysm beacon nodes experiencing crashes or consensus failures; validator clients unaffected.

---

## 2. Root Cause Analysis

### 2.1 The Problematic Code Path
The issue was isolated to the fork-choice algorithm in  
`beacon-chain/forkchoice/doubly-linked-tree/forkchoice.go`,  
specifically the `IsViableForCheckpoint()` function.

**Key Code Segment (lines 243–249):**

```go
if !features.Get().DisableLastEpochTargets {
    // Allow any node from the checkpoint epoch - 1 to be viable.
    nodeEpoch := slots.ToEpoch(node.slot)
    if nodeEpoch+1 == cp.Epoch {
        return true, nil
    }
}
```

### 2.2 Functional Interpretation
Under normal operation (`DisableLastEpochTargets=false`), this logic allows blocks from the previous epoch (**epoch − 1**) to be considered valid candidates for checkpoint synchronization.

This “leniency” requires the node to **generate or access historical states** to validate attestations pointing to these older blocks.

### 2.3 Fusaka Incompatibility
The Fusaka upgrade introduces **PeerDAS (EIP-7594)**, which fundamentally changes how historical state data is stored and retrieved.

The previous mechanism for generating “old states” — likely reliant on locally available complete state data — became:

- ineffective  
- incompatible  
- or computationally pathological  

under the new data availability model.

### 2.4 Failure Mode
When Prysm nodes encountered attestations for **epoch − 1** targets post-Fusaka:

1. The fork-choice logic triggered the **old state generation** process.  
2. This process **failed or hung** due to PeerDAS incompatibility.  

**Result:**

- Node crash  
- Consensus failure  
- Severe performance degradation

---

## 3. The Fix: `--disable-last-epoch-targets`

### 3.1 Mechanism
The flag sets:

```
DisableLastEpochTargets = true
```

This **skips the entire conditional block** shown earlier.

Effectively:

- Removes allowance for epoch − 1 blocks  
- Eliminates the need to generate incompatible historical states  

### 3.2 Trade-offs
**Stability Restored**  
Nodes immediately stabilize.

**Theoretical Impact**  
Minor reduction in fork-choice flexibility when reorganizations involve previous-epoch targets.

**Practical Reality**  
Mainnet continues finalizing seamlessly, indicating this was a **non-critical** code path under normal conditions.

---

## 4. Broader Implications for Protocol Security

### 4.1 Client–Protocol Coupling
This incident shows how deeply client implementations couple to protocol specifications.

A change in **data availability** (PeerDAS) surfaced a **latent bug** in an otherwise unrelated fork-choice logic.

### 4.2 The Value of Client Diversity
While Prysm nodes required intervention, other consensus clients:

- Lighthouse  
- Teku  
- Nimbus  
- Lodestar  

continued operating normally.

This is a real-world demonstration of why **client diversity is essential** for Ethereum’s resilience.

### 4.3 Testing Limitations
The bug survived multiple testnet deployments.

This reveals gaps in:

- state compatibility testing  
- cross-upgrade validation  
- old-state generation regression tests  

Future test plans should include simulations of state reconstruction across protocol transitions.

---

## 5. Conclusion
The Prysm Fusaka incident was a targeted **compatibility failure** in historical state handling, mitigated quickly through a precise configuration change.

It demonstrates the Ethereum ecosystem’s operational maturity:

- issues are detected fast  
- communication is timely  
- mitigation is clear  
- network impact remains minimal  

---

## 6. References & Further Reading
- **Prysm Team PSA on X**  
  https://twitter.com/prysmaticlabs/status/...

- **Prysm v7.0.0 Release Notes**  
  https://github.com/prysmaticlabs/prysm/releases/tag/v7.0.0

- **EIP-7594: PeerDAS Specification**  
  https://eips.ethereum.org/EIPS/eip-7594

---

## 7. About This Analysis
Author: 0xc黑冰e, Independent Protocol Security Researcher
Methodology: Public data collection, code review, and logical reconstruction
Repository: Full Research Portfolio
Date: December 2025

Disclaimer: This is an independent analysis based on publicly available information. It does not represent official positions of the Ethereum Foundation or Prysmatic Labs.