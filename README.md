# Next Greater Element

> Striver SDE Sheet 080 · Day 13 Stack & Queue — a monotonic stack of unresolved indexes; each element resolves everything smaller beneath it. Go, stdlib only, tested.

## Problem

For each array element, find the first strictly greater element to its right (-1 if none). The O(n²) scan-per-element is the baseline; the monotonic stack makes it one pass.

## Approach

Walk left to right keeping a stack of indexes whose answer is still unknown — by construction their values sit in decreasing order. Each new element *is the answer* for every stacked index with a smaller value: pop and record them all, then push the newcomer to wait its turn. Every index is pushed once and popped at most once, so the nested-looking loop is linear. Whatever survives to the end has no greater element: -1. This stack discipline is the template for next-smaller, stock span, and largest-rectangle problems on the following day.

**Complexity:** O(n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
