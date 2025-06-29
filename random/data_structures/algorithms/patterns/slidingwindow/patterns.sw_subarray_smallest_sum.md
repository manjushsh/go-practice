# Smallest and Largest Subarray with Given Sum 🎯

## 🍬 SmallestSubarrayWithGivenSum - ELI5

### 🎯 What it does
Find the **shortest** row of numbers that adds up to **AT LEAST** your target.

### 🍽️ Buffet Analogy 
Imagine you're at a magical buffet with plates of food, each with a calorie count. You want to eat **AT LEAST 500 calories** but using the **FEWEST plates possible**.

```
┌─────────────────────────────────────────────────────────────┐
│  🍽️ MAGICAL BUFFET CHALLENGE 🍽️                          │
├─────────────────────────────────────────────────────────────┤
│  Plates: [150] [100] [120] [200] [180] [50] calories       │
│  Goal: Get ≥ 500 calories using FEWEST plates             │
│  Solution: [200] + [180] + [120] = 500 calories (3 plates) │
│  ✅ Perfect! Only 3 plates needed!                        │
└─────────────────────────────────────────────────────────────┘
```

### 🔬 Visual Algorithm Flow

```
┌─────────────────────────────────────────────────────────────┐
│               DYNAMIC WINDOW EXPANSION                      │
├─────────────────────────────────────────────────────────────┤
│ Numbers: [2, 1, 2, 4, 3, 1]  Target: ≥ 7                 │
│                                                             │
│ Phase 1: EXPAND until sum ≥ target                        │
│ ┌─┐                                                         │
│ │2│ → sum=2 < 7 ❌                                         │
│ └─┘                                                         │
│                                                             │
│ ┌───┐                                                       │
│ │2,1│ → sum=3 < 7 ❌                                       │
│ └───┘                                                       │
│                                                             │
│ ┌─────┐                                                     │
│ │2,1,2│ → sum=5 < 7 ❌                                     │
│ └─────┘                                                     │
│                                                             │
│ ┌───────┐                                                   │
│ │2,1,2,4│ → sum=9 ≥ 7 ✅ Length=4                         │
│ └───────┘                                                   │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│              DYNAMIC WINDOW CONTRACTION                     │
├─────────────────────────────────────────────────────────────┤
│ Phase 2: CONTRACT while maintaining sum ≥ target          │
│                                                             │
│   ┌─────┐                                                   │
│ 2 │1,2,4│ → Remove 2, sum=7 ≥ 7 ✅ Length=3 (Better!)    │
│   └─────┘                                                   │
│                                                             │
│     ┌───┐                                                   │
│ 2,1 │2,4│ → Remove 1, sum=6 < 7 ❌ Can't contract more   │
│     └───┘                                                   │
│                                                             │
│ Continue sliding and find: [4,3] = 7 ≥ 7 ✅ Length=2!    │
│ 🏆 WINNER: Smallest length = 2                            │
└─────────────────────────────────────────────────────────────┘
```

---

## 🎵 MaxSubarrayWithGivenSum - ELI5

### 🎯 What it does
Find the **longest** row of numbers that adds up to **EXACTLY** your target.

### 🎵 Playlist Analogy
You're making a playlist where each song has a certain length (in minutes). You want a playlist that's **EXACTLY 15 minutes long**, but you want as **MANY songs as possible**.

```
┌─────────────────────────────────────────────────────────────┐
│  🎵 PERFECT PLAYLIST CHALLENGE 🎵                          │
├─────────────────────────────────────────────────────────────┤
│  Songs: [3] [5] [2] [4] [1] minutes each                   │
│  Goal: Get EXACTLY 15 minutes with MOST songs             │
│  Solution: [3] + [5] + [2] + [4] + [1] = 15 min (5 songs) │
│  🎶 Perfect! All songs fit!                               │
└─────────────────────────────────────────────────────────────┘
```

### 🔬 Visual Algorithm Flow

```
┌─────────────────────────────────────────────────────────────┐
│              EXACT SUM WINDOW SEARCH                       │
├─────────────────────────────────────────────────────────────┤
│ Numbers: [1, 2, 1, 2, 1]  Target: EXACTLY 4               │
│                                                             │
│ 🔍 Searching for all windows that sum to exactly 4:        │
│                                                             │
│ ┌───────┐                                                   │
│ │1,2,1,2│ → sum=6 > 4 ❌ Too big! Contract...             │
│ └───────┘                                                   │
│                                                             │
│   ┌─────┐                                                   │
│ 1 │2,1,2│ → sum=5 > 4 ❌ Still too big! Contract...       │
│   └─────┘                                                   │
│                                                             │
│     ┌───┐                                                   │
│ 1,2 │1,2│ → sum=3 < 4 ❌ Too small! Expand...             │
│     └───┘                                                   │
│                                                             │
│     ┌─────┐                                                 │
│ 1,2 │1,2,1│ → sum=4 = 4 ✅ Found match! Length=3          │
│     └─────┘                                                 │
│                                                             │
│ Continue searching... Found: [2,1,2] and [1,2,1]          │
│ 🏆 All have length 3 - Maximum length = 3                 │
└─────────────────────────────────────────────────────────────┘
```

### � Side-by-Side Comparison

```
┌─────────────────────────────────────────────────────────────┐
│          SMALLEST vs LARGEST SUBARRAY                      │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│ 🔎 SmallestSubarrayWithGivenSum    🔍 MaxSubarrayWithGivenSum │
│ ──────────────────────────────      ─────────────────────── │
│ 🎯 Goal: sum ≥ target               🎯 Goal: sum = target   │
│ 📏 Find: SHORTEST length            📏 Find: LONGEST length │
│ � Strategy: Expand then contract   💡 Strategy: Exact match│
│ ⚡ Allows overshoot                 ⚡ Must be precise     │
│                                                             │
│ Example: target=7, array=[2,1,2,4,3,1]                     │
│ ✅ Answer: length=2 [4,3]=7         ✅ Answer: length=3     │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Key Differences 🔍

| SmallestSubarrayWithGivenSum | MaxSubarrayWithGivenSum |
|------------------------------|-------------------------|
| Finds **shortest** group     | Finds **longest** group |
| Sum must be **≥ targetSum**  | Sum must be **= exactSum** |
| Allows sum to be bigger      | Sum must be exact       |

## Time Complexity ⏰
Both functions use the **sliding window** technique:
- **Time:** O(n) - Each element is visited at most twice
- **Space:** O(1) - Only using a few variables

## When to Use 🤔
- **SmallestSubarrayWithGivenSum:** When you need "at least" this much
- **MaxSubarrayWithGivenSum:** When you need "exactly" this much
