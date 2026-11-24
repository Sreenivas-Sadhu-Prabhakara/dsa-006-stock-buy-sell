# Best Time to Buy and Sell Stock

> Striver SDE Sheet 006 · Day 1 Arrays — one pass tracking the cheapest day so far and the best sale against it. Go, stdlib only, tested.

## Problem

Given daily stock prices, choose one day to buy and a later day to sell to maximize profit. If every trade loses money, make no trade and return 0.

## Approach

For each day, the best sale on that day is its price minus the cheapest price seen before it. So one pass suffices: carry the minimum price so far, measure today's price against it, and keep the best difference. No need to consider pairs — the running minimum already summarizes every earlier buy candidate.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
