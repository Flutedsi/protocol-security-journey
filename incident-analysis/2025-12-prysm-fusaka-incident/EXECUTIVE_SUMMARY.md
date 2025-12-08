# Executive Summary: Prysm Fusaka Incident Analysis

## Quick Facts
- **Date**: December 2025 (post-Fusaka upgrade)
- **Affected Software**: Prysm consensus client v7.0.0+
- **Fix**: `--disable-last-epoch-targets` beacon node flag
- **Impact**: Prysm node stability; no network finalization issues

## The Problem in One Sentence
Prysm's fork-choice algorithm contained legacy logic that tried to generate pre-Fusaka historical states after the PeerDAS upgrade, causing nodes to crash.

## Technical Root Cause
1. **Where**: `forkchoice.go` → `IsViableForCheckpoint()` function
2. **What**: Special allowance for "previous epoch" blocks required old state generation
3. **Why Broken**: Fusaka's PeerDAS changed how historical data is accessed, making old state generation fail

## The Fix Explained
The `--disable-last-epoch-targets` flag turns off the problematic "previous epoch" allowance, eliminating the need to generate incompatible historical states.

## Key Takeaways
1. **Client Diversity Worked**: Only Prysm affected; network continued finalizing
2. **Protocol Changes Have Ripple Effects**: A data layer change (PeerDAS) broke consensus layer logic
3. **Ecosystem is Resilient**: Issue was quickly identified, communicated, and mitigated

## For Node Operators
- If running Prysm: Ensure you're using the flag
- Monitor client communications during upgrades
- Consider multi-client setups for critical infrastructure

## For Researchers
- Study client upgrade compatibility matrices
- Develop tests for historical state validation across hard forks
- Contribute to client implementation reviews

---
*Full technical analysis available in [README.md](./README.md)*  
*Authored by independent researcher 0xc黑冰e*