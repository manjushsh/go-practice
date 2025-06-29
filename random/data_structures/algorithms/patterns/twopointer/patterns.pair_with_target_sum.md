# Two Pointer Pattern: Pair with Target Sum

## ELI5 (Explain Like I'm 5) 🧒

Imagine you have a line of kids arranged by height from shortest to tallest. You want to find two kids whose heights add up to exactly a specific number.

**The Smart Way (Two Pointer):**
- Put one finger on the shortest kid (left side)
- Put another finger on the tallest kid (right side)
- Add their heights together
- If the sum is too small, move your left finger to a taller kid
- If the sum is too big, move your right finger to a shorter kid
- If the sum is just right, you found your pair!

This is much faster than checking every possible pair of kids!

## What This Algorithm Does

The `PairWithTargetSum` function finds two numbers in a **sorted array** that add up to a target sum. It returns the **indices** (positions) of these two numbers.

## How It Works Step by Step

### Input Requirements
- **Sorted array**: The array must be sorted in ascending order
- **Target sum**: The number we want two elements to add up to

### 🎨 Enhanced Visual Algorithm Flow

```
┌─────────────────────────────────────────────────────────────┐
│               INITIALIZATION                                │
├─────────────────────────────────────────────────────────────┤
│ Input: [1, 2, 3, 4, 6]  Target: 6                         │
│        ↑🔵           ↑🔴                                    │
│    startPointer   endPointer                               │
│    (index 0)      (index 4)                               │
│                                                             │
│ 🎯 Goal: Find two numbers that sum to 6                   │
│ 💭 Strategy: Squeeze pointers toward center               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   ITERATION 1                              │
├─────────────────────────────────────────────────────────────┤
│ Array: [1] [2] [3] [4] [6]                                 │
│         ↑🔵           ↑🔴                                   │
│     start(0)      end(4)                                   │
│                                                             │
│ 🔢 Calculate: 1 + 6 = 7                                   │
│ ❌ 7 > 6 (too big!)                                        │
│ 💡 Decision: Sum too large, need smaller number           │
│ 🔄 Action: Move endPointer left ⬅️                        │
│                                                             │
│ 🧠 Logic: Since array is sorted, moving right pointer     │
│          left gives us a smaller value                     │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   ITERATION 2                              │
├─────────────────────────────────────────────────────────────┤
│ Array: [1] [2] [3] [4] [6]                                 │
│         ↑🔵       ↑🔴                                       │
│     start(0)   end(3)                                      │
│                                                             │
│ 🔢 Calculate: 1 + 4 = 5                                   │
│ ❌ 5 < 6 (too small!)                                      │
│ 💡 Decision: Sum too small, need larger number            │
│ 🔄 Action: Move startPointer right ➡️                     │
│                                                             │
│ 🧠 Logic: Since array is sorted, moving left pointer      │
│          right gives us a larger value                     │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   ITERATION 3                              │
├─────────────────────────────────────────────────────────────┤
│ Array: [1] [2] [3] [4] [6]                                 │
│            ↑🔵   ↑🔴                                        │
│        start(1) end(3)                                     │
│                                                             │
│ 🔢 Calculate: 2 + 4 = 6                                   │
│ ✅ 6 == 6 (perfect match!)                                 │
│ 🎉 Found the pair!                                         │
│ 📤 Return: [1, 3] (indices of elements 2 and 4)          │
│                                                             │
│ 🏆 Success: Elements at positions 1 and 3 sum to target   │
└─────────────────────────────────────────────────────────────┘
```

### 🧠 Decision Tree Visualization

```
┌─────────────────────────────────────────────────────────────┐
│              POINTER MOVEMENT DECISION TREE                 │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│        Calculate: sum = arr[left] + arr[right]             │
│                          ↓                                  │
│          ┌─────────────────────────────┐                   │
│          │    sum compared to target   │                   │
│          └─────────────────────────────┘                   │
│                          ↓                                  │
│    ┌─────────────┬───────────────┬─────────────┐           │
│    │   sum < target    │  sum == target   │  sum > target │ │
│    │        ↓          │       ↓          │       ↓       │ │
│    │  Move LEFT ➡️      │   🎯 FOUND!     │  Move RIGHT ⬅️│ │
│    │  pointer right    │   Return [L,R]   │  pointer left │ │
│    │  (need larger)    │                  │ (need smaller)│ │
│    │                   │                  │               │ │
│    └─────────────┬─────┴─────┬─────────────┴───────────────┘ │
│                  │           │                               │
│            Continue loop while left < right                  │
│            If loop ends without finding: return [-1, -1]     │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Code Breakdown

```go
func PairWithTargetSum(arr []int, targetSum int) []int {
    startPointer := 0              // Start at beginning
    endPointer := len(arr) - 1     // Start at end
    
    for startPointer < endPointer {
        currentSum := arr[startPointer] + arr[endPointer]
        
        if currentSum == targetSum {
            return []int{startPointer, endPointer}  // Found it!
        }
        if currentSum < targetSum {
            startPointer++  // Need bigger sum, move left pointer right
        } else {
            endPointer--    // Need smaller sum, move right pointer left
        }
    }
    return []int{-1, -1}  // Not found
}
```

## Why This Works

1. **Sorted Array Advantage**: Because the array is sorted, we know:
   - Moving the left pointer right increases the sum
   - Moving the right pointer left decreases the sum

2. **Elimination Strategy**: In each step, we eliminate one possibility:
   - If sum is too small: we'll never find the answer with current left element
   - If sum is too big: we'll never find the answer with current right element

## Time & Space Complexity

- **Time Complexity**: O(n) - We visit each element at most once
- **Space Complexity**: O(1) - We only use two pointers

Compare this to the brute force approach which would be O(n²)!

## Example Walkthrough

```go
arr := []int{1, 2, 3, 4, 6}
targetSum := 6

// Initial state
startPointer = 0, endPointer = 4
arr[0] + arr[4] = 1 + 6 = 7 > 6, so endPointer--

// Step 1
startPointer = 0, endPointer = 3
arr[0] + arr[3] = 1 + 4 = 5 < 6, so startPointer++

// Step 2
startPointer = 1, endPointer = 3
arr[1] + arr[3] = 2 + 4 = 6 == 6, return [1, 3]
```

## When to Use This Pattern

✅ **Good for:**
- Finding pairs in sorted arrays
- Two sum problems
- Finding triplets that meet criteria
- Palindrome checking

❌ **Not suitable for:**
- Unsorted arrays (sort first, or use hash map)
- When you need all pairs (this finds just one)
- Single element searches

## Common Variations

1. **Find all pairs** that sum to target
2. **Triplet sum** (three pointers)
3. **Closest pair** to target sum
4. **Pair difference** equals target

## Real-World Applications

- **Shopping**: Find two items that fit your exact budget
- **Chemistry**: Mix two solutions to get specific concentration
- **Gaming**: Combine two cards with specific total power
- **Finance**: Find two investments that total your available funds