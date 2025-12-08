# Core Code Evidence: Prysm Fusaka Incident Analysis

## Evidence Collected: December 2025
**Prysm Version:** v7.0.0 (`7794a77ae6`)  
**Analysis Focus:** `--disable-last-epoch-targets` flag and related fork-choice logic

---

## 1. Flag Definition & Configuration

### **File:** `config/features/flags.go`

```go
// disableLastEpochTargets is a flag to disable processing of attestations for old blocks.
disableLastEpochTargets = &cli.BoolFlag{
    Name:  "disable-last-epoch-targets",
    Usage: "Disables processing of states for attestations to old blocks",
}
```

### **File:** `config/features/config.go`

```go
// DisableLastEpochTargets disables processing of states for attestations to old blocks.
DisableLastEpochTargets bool

if ctx.IsSet(disableLastEpochTargets.Name) {
    logEnabled(disableLastEpochTargets)
    cfg.DisableLastEpochTargets = true
}
```

**Interpretation:**  
This flag directly controls whether Prysm should attempt to generate / process historical states required by attestations that target **epoch − 1** blocks.

---

## 2. Critical Logic in Fork-Choice Algorithm

### **File:** `beacon-chain/forkchoice/doubly-linked-tree/forkchoice.go`  
### **Function:** `IsViableForCheckpoint` (lines 243–249)

```go
if !features.Get().DisableLastEpochTargets {
    // Allow any node from the checkpoint epoch - 1 to be viable.
    nodeEpoch := slots.ToEpoch(node.slot)
    if nodeEpoch+1 == cp.Epoch {
        return true, nil
    }
}
```

### Context & Functional Meaning
This block determines whether a block may be considered **viable for checkpoint synchronization**.

When the flag is **disabled (false)**:

- Prysm permits checkpoint consideration of blocks from **epoch − 1**
- This triggers **old-state generation**, which became incompatible after the Fusaka PeerDAS transition

This confirms why an emergency toggle of the flag stabilized the client.

---

## 3. Functional Documentation Evidence

### Key Internal Documentation Phrase:
> “Disables processing of states for attestations to old blocks.”

This internal description directly matches:

- The scenario observed post-Fusaka  
- The hypothesis that PeerDAS made historical state access paths non-functional  
- The behavior seen when nodes attempted to validate epoch − 1 attestations  

This supports PeerDAS incompatibility as the root cause.

---