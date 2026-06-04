<div align="center">

# 🧠 DSA in Go — Daily Brain Warm-Up

### _While others are taking coffee, I'm taking DSA._

**A daily ritual of logic.** One problem a day, written in idiomatic Go — a mental upliftment of the brain toward the real-life sequences and logic that power every early start.

[![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![LeetCode](https://img.shields.io/badge/LeetCode-grind-FFA116?style=flat-square&logo=leetcode&logoColor=white)](https://leetcode.com)
[![HackerRank](https://img.shields.io/badge/HackerRank-grind-2EC866?style=flat-square&logo=hackerrank&logoColor=white)](https://hackerrank.com)
[![Daily](https://img.shields.io/badge/cadence-daily-success?style=flat-square)](#-the-daily-cadence)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](#-license)

</div>

---

## ☕ The Philosophy

Most people start their day by reaching for a cup of coffee — a chemical jolt to drag the mind awake.

This repository is a different stimulant.

Every morning, before the noise of the day sets in, I solve one Data Structures & Algorithms problem in **Go**. It's not about grinding for an interview. It's about **warming the brain** — training it to see patterns, decompose problems, and reason about the logical sequences that show up everywhere in real life, not just in code.

> **Coffee wakes the body. DSA wakes the mind.**

This is a logbook of that habit — a public, version-controlled record of showing up every single day.

---

## 🎯 Intent & Goals

| Goal | Why it matters |
| --- | --- |
| **Daily consistency** | Compounding beats intensity. One problem a day, every day, builds a sharper mind than a weekend cram. |
| **Logical fluency** | DSA is structured thinking. The skill transfers to system design, debugging, and decision-making. |
| **Idiomatic Go** | Practice the language as much as the algorithm — clean, readable, production-shaped Go. |
| **A visible streak** | A git history is an honest accountability partner. Green squares don't lie. |
| **Pattern recognition** | Two pointers, sliding windows, recursion, DP — recognize the shape of a problem before solving it. |

---

## 📚 Problem Sources

Problems are pulled from a mix of platforms, so the brain never settles into one style:

- 🟧 **[LeetCode](https://leetcode.com)** — the bread and butter
- 🟩 **[HackerRank](https://www.hackerrank.com)** — challenges and structured tracks
- 🧩 **Other sources** — Codeforces, Project Euler, classic CS texts, and the occasional real-world puzzle

---

## 🗂️ Repository Structure

Each solution lives in its own self-contained Go file, named by the **day** it was solved and the **problem** it tackles:

```
dsa/
├── go.mod                      # module: github.com/mananuf/dsa
├── README.md                   # you are here
├── d1_add_two_numbers.go       # Day 1 — LeetCode #2: Add Two Numbers
└── d2_len_of_longest_substring.go  # Day 2 — LeetCode #3: Longest Substring Without Repeating Characters
```

### Naming convention

```
d<day>_<problem_name_snake_case>.go
```

| Part | Meaning | Example |
| --- | --- | --- |
| `d<day>` | The day number in the streak | `d1`, `d2`, `d42` |
| `<problem_name>` | The problem, in snake_case | `add_two_numbers` |

Every file carries the **problem statement, constraints, and worked examples as a header comment**, so each solution reads as a standalone document — no need to look anything up to understand what was solved and why.

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed and on your `PATH`

```bash
go version   # expect go1.21 or newer
```

### Clone the repo

```bash
git clone https://github.com/mananuf/dsa.git
cd dsa
```

### Run a single day's solution

Each file is its own `package main` with a `main()` driver, so you can run any day directly:

```bash
go run d1_add_two_numbers.go
```

Expected output (Day 1 — `342 + 465 = 807`, stored reverse → `[7,0,8]`):

```
7 ...
0 ...
8 <nil>
```

### Build a single file

```bash
go build -o /tmp/day1 d1_add_two_numbers.go && /tmp/day1
```

---

## 🛠️ Workflow & Conventions

A consistent shape makes the whole logbook easy to read and easy to grow.

1. **One file per problem.** Self-contained, runnable, and documented in place.
2. **Problem statement in a header comment.** Constraints and examples included verbatim from the source.
3. **Idiomatic Go.** Favor clarity over cleverness; let the standard library do the heavy lifting.
4. **Time & space complexity noted** where it sharpens the analysis.
5. **Alternative approaches kept** (often as commented-out blocks) to show the road not taken — the brute force before the elegant pass.

### Example: the anatomy of a solution

From `d1_add_two_numbers.go` — clean, linear, single-pass:

```go
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10
        tail.Next = &ListNode{Val: sum % 10}
        tail = tail.Next
    }
    return dummy.Next
}
```

> **Complexity:** `O(max(m, n))` time, `O(max(m, n))` space — a single pass with a dummy-head sentinel to keep the list-building branch-free.

---

## 📈 Progress Log

A running tally of the streak. Updated as the days roll on.

| Day | Problem | Source | Difficulty | Pattern | Solution |
| --- | --- | --- | --- | --- | --- |
| 1 | Add Two Numbers (#2) | LeetCode | 🟡 Medium | Linked List · Math · Carry | [`d1_add_two_numbers.go`](./d1_add_two_numbers.go) |
| 2 | Longest Substring Without Repeating Characters (#3) | LeetCode | 🟡 Medium | Strings · Sliding Window · Hash Map | [`d2_len_of_longest_substring.go`](./d2_len_of_longest_substring.go) |

> _More squares to come. The streak is the point._

---

## 🧭 Topics Covered (and on the horizon)

As the logbook grows, it will span the full DSA landscape:

- **Arrays & Strings** — two pointers, sliding window, prefix sums ✅ _(started Day 2)_
- **Linked Lists** — traversal, reversal, cycle detection ✅ _(started Day 1)_
- **Stacks & Queues** — monotonic stacks, BFS scaffolding
- **Hash Maps & Sets** — frequency counting, lookups, dedup
- **Trees & Graphs** — DFS, BFS, traversals, shortest paths
- **Recursion & Backtracking** — permutations, combinations, subsets
- **Sorting & Searching** — binary search, quickselect
- **Dynamic Programming** — memoization, tabulation, state machines
- **Greedy & Math** — interval scheduling, number theory

---

## 🤝 Contributing

This is primarily a personal training journal — but suggestions, cleaner approaches, and "have you considered…" notes are always welcome.

1. Fork the repo
2. Add or improve a solution following the [naming convention](#naming-convention)
3. Keep the header-comment style (problem statement + examples)
4. Open a pull request describing the approach and complexity

---

## 📜 License

Released under the **MIT License** — use it, learn from it, build your own streak on it.

---

<div align="center">

### The brain is a muscle. This is the morning workout.

**One problem. Every day. No days off.** 🧠☕

</div>
