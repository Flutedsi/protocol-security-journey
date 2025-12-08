# Reproduction Guide: Verifying the Code Evidence

## Prerequisites
- Git installed
- Access to https://github.com/prysmaticlabs/prysm

---

## Step-by-Step Verification

### 1. Clone and Checkout the Relevant Version
```bash
git clone https://github.com/prysmaticlabs/prysm.git
cd prysm
git fetch --tags
git checkout tags/v7.0.0   # Commit 7794a77ae6
```

---

### 2. Verify the Flag Definition
```bash
grep -n -B2 -A2 '"disable-last-epoch-targets"' config/features/flags.go
grep -n -B5 -A5 "DisableLastEpochTargets" config/features/config.go
```

---

### 3. Locate the Core Logic
```bash
grep -n -B15 -A10 "DisableLastEpochTargets" \
    beacon-chain/forkchoice/doubly-linked-tree/forkchoice.go
```

---

### 4. Expected Outputs
The search commands should return the same code snippets documented in **CORE_EVIDENCE.md**.

---

## Independent Search Commands

```bash
# Find all references to the feature flag
grep -r "DisableLastEpochTargets" . --include="*.go"

# Explore attestation/state logic involving old epochs
grep -r "old.*attestation\|epoch.*target" . \
    --include="*.go" -i | head -20
```