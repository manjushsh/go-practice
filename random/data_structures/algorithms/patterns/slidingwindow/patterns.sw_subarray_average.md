# The Sliding Window Average - Like a Moving Camera! 📹

## 🎬 What Does This Do?

Imagine you have a row of colorful blocks on the floor, and you want to find the average number on groups of blocks using a special "camera" that can only see a certain number of blocks at once.

This code finds the average of groups of numbers in a list. If you have numbers `[1, 2, 3, 4, 5]` and want groups of 3, it finds the average of `[1,2,3]`, then `[2,3,4]`, then `[3,4,5]`.

## 🎨 Visual Camera Metaphor

```
┌─────────────────────────────────────────────────────────────┐
│           SLIDING WINDOW CAMERA VISUALIZATION               │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│ Blocks: [🔵1] [🟢2] [🟡3] [🔴4] [🟣5]                      │
│                                                             │
│ 📷 Camera View (width = 3 blocks):                         │
│                                                             │
│ Position 1: [🔵1] [🟢2] [🟡3] 🔴 🟣                        │
│             └─── 📷 VIEW ───┘                              │
│             Average: (1+2+3)/3 = 2.0 📊                   │
│                                                             │
│ Position 2: 🔵 [🟢2] [🟡3] [🔴4] 🟣                        │
│                └─── 📷 VIEW ───┘                           │
│                Average: (2+3+4)/3 = 3.0 📊                │
│                                                             │
│ Position 3: 🔵 🟢 [🟡3] [🔴4] [🟣5]                        │
│                   └─── 📷 VIEW ───┘                        │
│                   Average: (3+4+5)/3 = 4.0 📊             │
│                                                             │
│ 🎯 Result: [2.0, 3.0, 4.0]                                │
└─────────────────────────────────────────────────────────────┘
```

## 🧠 Smart vs Naive Approach

### 🐌 The Slow Way (Naive):
```
┌─────────────────────────────────────────────────────────────┐
│                    NAIVE APPROACH                           │
├─────────────────────────────────────────────────────────────┤
│ For each window position:                                   │
│   • Calculate sum from scratch                              │
│   • Divide by window size                                   │
│                                                             │
│ Window 1: 1+2+3 = 6, avg = 6/3 = 2.0  [🔢🔢🔢]          │
│ Window 2: 2+3+4 = 9, avg = 9/3 = 3.0  [🔢🔢🔢]          │ 
│ Window 3: 3+4+5 = 12, avg = 12/3 = 4.0 [🔢🔢🔢]         │
│                                                             │
│ Total operations: 9 additions + 3 divisions = 12 ops       │
│ Time complexity: O(n × k) where k = window size           │
└─────────────────────────────────────────────────────────────┘
```

### ⚡ The Fast Way (Sliding Window):
```
┌─────────────────────────────────────────────────────────────┐
│                  SLIDING WINDOW APPROACH                    │
├─────────────────────────────────────────────────────────────┤
│ Maintain running sum, adjust by +new -old:                 │
│                                                             │
│ Initial: 1+2+3 = 6, avg = 6/3 = 2.0    [🔢🔢🔢]         │
│ Slide 1: 6-1+4 = 9, avg = 9/3 = 3.0    [➖➕]             │
│ Slide 2: 9-2+5 = 12, avg = 12/3 = 4.0  [➖➕]             │
│                                                             │
│ Total operations: 3+2+2 = 7 additions + 3 divisions = 10   │
│ Time complexity: O(n) - much faster! 🚀                   │
│ Efficiency gain: ~40% faster for this example!             │
└─────────────────────────────────────────────────────────────┘
```

## 🎮 Step-by-Step Algorithm Walkthrough

### Input: `[1, 3, 2, 6, -1, 4, 1, 8, 2]` with window size `k = 5`

```
┌─────────────────────────────────────────────────────────────┐
│                     INITIALIZATION                          │
├─────────────────────────────────────────────────────────────┤
│ windowEnd = 0, windowStart = 0                             │
│ sum = 0                                                     │
│ results = [ _, _, _, _, _ ] (5 slots for 5 possible windows)│
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ EXPANDING PHASE: Building first window                     │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│ Step 1: windowEnd=0                                        │
│ [1] 3  2  6 -1  4  1  8  2                                │
│  ↑🔴                                                        │
│ windowEnd                                                   │
│ sum += 1 → sum = 1                                         │
│ windowEnd(0) < k-1(4)? YES → continue                     │
│                                                             │
│ Step 2: windowEnd=1                                        │
│ [1][3] 2  6 -1  4  1  8  2                                │
│     ↑🔴                                                     │
│ sum += 3 → sum = 4                                         │
│ windowEnd(1) < k-1(4)? YES → continue                     │
│                                                             │
│ Step 3: windowEnd=2                                        │
│ [1][3][2] 6 -1  4  1  8  2                                │
│        ↑🔴                                                  │
│ sum += 2 → sum = 6                                         │
│ windowEnd(2) < k-1(4)? YES → continue                     │
│                                                             │
│ Step 4: windowEnd=3                                        │
│ [1][3][2][6]-1  4  1  8  2                                │
│           ↑🔴                                               │
│ sum += 6 → sum = 12                                        │
│ windowEnd(3) < k-1(4)? YES → continue                     │
│                                                             │
│ Step 5: windowEnd=4                                        │
│ [1][3][2][6][-1] 4  1  8  2                               │
│              ↑🔴                                            │
│ sum += (-1) → sum = 11                                     │
│ windowEnd(4) >= k-1(4)? YES → First window complete!      │
│ results[0] = 11/5 = 2.2                                    │
│ sum -= arr[windowStart] → sum = 11-1 = 10                 │
│ windowStart++ → windowStart = 1                            │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ SLIDING PHASE: Efficient window movement                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│ Step 6: windowEnd=5                                        │
│  1 [3][2][6][-1][4] 1  8  2                               │
│  ↑🔷               ↑🔴                                      │
│ windowStart      windowEnd                                  │
│ sum += 4 → sum = 14                                        │
│ results[1] = 14/5 = 2.8                                    │
│ sum -= arr[1] → sum = 14-3 = 11                           │
│ windowStart++ → windowStart = 2                            │
│                                                             │
│ Step 7: windowEnd=6                                        │
│  1  3 [2][6][-1][4][1] 8  2                               │
│     ↑🔷                ↑🔴                                  │
│ sum += 1 → sum = 12                                        │
│ results[2] = 12/5 = 2.4                                    │
│ sum -= arr[2] → sum = 12-2 = 10                           │
│ windowStart++ → windowStart = 3                            │
│                                                             │
│ Step 8: windowEnd=7                                        │
│  1  3  2 [6][-1][4][1][8] 2                               │
│        ↑🔷                 ↑🔴                              │
│ sum += 8 → sum = 18                                        │
│ results[3] = 18/5 = 3.6                                    │
│ sum -= arr[3] → sum = 18-6 = 12                           │
│ windowStart++ → windowStart = 4                            │
│                                                             │
│ Step 9: windowEnd=8                                        │
│  1  3  2  6 [-1][4][1][8][2]                              │
│           ↑🔷                 ↑🔴                           │
│ sum += 2 → sum = 14                                        │
│ results[4] = 14/5 = 2.8                                    │
│                                                             │
│ 🎯 Final results: [2.2, 2.8, 2.4, 3.6, 2.8]              │
└─────────────────────────────────────────────────────────────┘
```


# Other Explanations

## 1. Reason for `results := make([]float64, len(nums)-k+1)`

The expression `len(nums) - k + 1` calculates the number of contiguous subarrays of length `k` you can form from the slice `nums`.

**Explanation:**

- `len(nums)` is the total number of elements in the array.
- For each subarray of size `k`, you need `k` consecutive elements.
- The first subarray starts at index `0`, the next at index `1`, and so on, up to index `len(nums) - k`.

So, the last possible starting index is `len(nums) - k`. That means there are `len(nums) - k + 1` possible starting positions, and thus that many subarrays of size `k`.

**Example:**

If `nums = [1, 2, 3, 4, 5]` and `k = 3`:

- Possible subarrays: `[1,2,3]`, `[2,3,4]`, `[3,4,5]`
- Number of subarrays: `5 - 3 + 1 = 3`

This ensures the `results` slice is exactly the right size to store all averages (or results) for each subarray of length `k`.