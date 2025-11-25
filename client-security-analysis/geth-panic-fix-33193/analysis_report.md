# Geth Security Analysis: randomDuration Boundary Condition Panic Fix (#33193)

## PR Information

- **Commit Hash**: e0d81d1e993ad6dc3e618cd06e56b7be916efd8e

- **Title**: eth: fix panic in randomDuration when min equals max

- **Affected Module**: eth/dropper.go (Peer Manager)

- **Fix Type**: Boundary Condition Handling

## Problem Description

### Background

The `randomDuration` function is used to generate random time intervals in the peer manager. The original implementation caused a panic when the `min` and `max` parameters were equal.

### Vulnerability Mechanism

```go

// Problematic code before fix return time.Duration(mrand.Int63n(int64(max-min)) + int64(min))