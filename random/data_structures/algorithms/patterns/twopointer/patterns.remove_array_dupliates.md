# Remove Array Duplicates - Two Pointer Pattern

## 🎯 What does this function do?

The `RemoveArrayDuplicates` function removes duplicate elements from a **sorted array** and returns a new array with only unique elements, while maintaining the original order.

## 🧠 ELI5 (Explain Like I'm 5)

Imagine you have a line of toy cars that are arranged in order by color: 🚗🚗🚙🚙🚙🚕🚕

You want to keep only one car of each color. You use two fingers:
- **Left finger (nextNonDup)**: Points to where the next unique car should go
- **Right finger (currentElement)**: Moves through all cars to find new colors

You start with your left finger on the second position (index 1) because the first car is always unique.

## 🔍 How it works step by step

### Algorithm Steps:
1. **Start**: Left pointer at index 1, right pointer at index 1
2. **Compare**: Is the current car different from the previous unique car?
3. **If different**: Copy the car to the left pointer position and move left pointer forward
4. **If same**: Just move the right pointer forward (skip the duplicate)
5. **Repeat**: Until we've checked all cars

### 🎨 Enhanced Visual Example with Color Coding:

```
┌─────────────────────────────────────────────────────────────┐
│            REMOVING DUPLICATES VISUALIZATION                │
├─────────────────────────────────────────────────────────────┤
│ Original: [1, 1, 2, 3, 3, 3, 4, 4, 5]                     │
│           ↑🔷 nextNonDup                                    │
│           ↑🔴 currentElement                                │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ STEP 1: Compare positions 0 and 1                          │
│ [1] [1] [2] [3] [3] [3] [4] [4] [5]                        │
│  ↑🟢  ↑🔴                                                   │
│ prev current                                                │
│                                                             │
│ Compare: 1 == 1? ✅ YES → Skip duplicate                   │
│ Action: Move currentElement forward only                    │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ STEP 2: Compare positions 0 and 2                          │
│ [1] [1] [2] [3] [3] [3] [4] [4] [5]                        │
│  ↑🟢      ↑🔴                                               │
│ prev    current                                             │
│                                                             │
│ Compare: 1 == 2? ❌ NO → Found new unique!                 │
│ Action: Copy 2 to position 1, move nextNonDup              │
│ Result: [1] [2] [2] [3] [3] [3] [4] [4] [5]                │
│             ↑🔷                                             │
│            nextNonDup                                       │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ STEP 3: Compare positions 1 and 3                          │
│ [1] [2] [2] [3] [3] [3] [4] [4] [5]                        │
│      ↑🟢      ↑🔴                                           │
│     prev    current                                         │
│                                                             │
│ Compare: 2 == 3? ❌ NO → Found new unique!                 │
│ Action: Copy 3 to position 2, move nextNonDup              │
│ Result: [1] [2] [3] [3] [3] [3] [4] [4] [5]                │
│                 ↑🔷                                         │
│                nextNonDup                                   │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ STEP 4-5: Skip duplicates of 3                             │
│ [1] [2] [3] [3] [3] [3] [4] [4] [5]                        │
│          ↑🟢          ↑🔴                                   │
│         prev        current                                 │
│                                                             │
│ Compare: 3 == 3? ✅ YES → Skip                             │
│ Compare: 3 == 3? ✅ YES → Skip                             │
│ Keep moving currentElement...                               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ STEP 6: Found 4                                            │
│ [1] [2] [3] [3] [3] [3] [4] [4] [5]                        │
│          ↑🟢              ↑🔴                               │
│         prev            current                             │
│                                                             │
│ Compare: 3 == 4? ❌ NO → Found new unique!                 │
│ Action: Copy 4 to position 3, move nextNonDup              │
│ Result: [1] [2] [3] [4] [3] [3] [4] [4] [5]                │
│                     ↑🔷                                     │
│                   nextNonDup                               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ FINAL STEPS: Skip 4, Add 5                                 │
│ [1] [2] [3] [4] [3] [3] [4] [4] [5]                        │
│              ↑🟢                  ↑🔴                       │
│             prev                current                     │
│                                                             │
│ Skip duplicate 4, then copy 5:                             │
│ Result: [1] [2] [3] [4] [5] [3] [4] [4] [5]                │
│                         ↑                                   │
│                    nextNonDup=5                             │
│                                                             │
│ 🎯 Return: originalArray[:5] = [1, 2, 3, 4, 5]            │
└─────────────────────────────────────────────────────────────┘
```

## 💻 Code Walkthrough

```go
func RemoveArrayDuplicates(originalArray []int) []int {
    nextNonDup := 1  // Position for next unique element (skip first element)
    
    // Start from second element
    for currentElement := 1; currentElement < len(originalArray); currentElement++ {
        // Compare current with previous unique element
        if originalArray[nextNonDup-1] != originalArray[currentElement] {
            // Found a new unique element
            originalArray[nextNonDup] = originalArray[currentElement]
            nextNonDup++  // Move to next position for unique elements
        }
        // If duplicate, just continue (skip it)
    }
    
    // Return slice with only unique elements
    return originalArray[:nextNonDup]
}
```

## 📝 Detailed Examples

### Example 1: Basic case
```go
Input:  [1, 1, 2, 3, 3, 3, 4, 4, 5]
Output: [1, 2, 3, 4, 5]

Trace:
- nextNonDup=1, currentElement=1: originalArray[0]=1, originalArray[1]=1 → Same, skip
- nextNonDup=1, currentElement=2: originalArray[0]=1, originalArray[2]=2 → Different!
  * Copy: originalArray[1] = 2, nextNonDup=2
- nextNonDup=2, currentElement=3: originalArray[1]=2, originalArray[3]=3 → Different!
  * Copy: originalArray[2] = 3, nextNonDup=3
- nextNonDup=3, currentElement=4: originalArray[2]=3, originalArray[4]=3 → Same, skip
- nextNonDup=3, currentElement=5: originalArray[2]=3, originalArray[5]=3 → Same, skip
- nextNonDup=3, currentElement=6: originalArray[2]=3, originalArray[6]=4 → Different!
  * Copy: originalArray[3] = 4, nextNonDup=4
- nextNonDup=4, currentElement=7: originalArray[3]=4, originalArray[7]=4 → Same, skip
- nextNonDup=4, currentElement=8: originalArray[3]=4, originalArray[8]=5 → Different!
  * Copy: originalArray[4] = 5, nextNonDup=5

Final: return originalArray[:5] = [1, 2, 3, 4, 5]
```

### Example 2: No duplicates
```go
Input:  [1, 2, 3, 4, 5]
Output: [1, 2, 3, 4, 5]

Trace: Every comparison finds different elements, so all are kept.
```

### Example 3: All duplicates
```go
Input:  [7, 7, 7, 7, 7]
Output: [7]

Trace: Only the first element is kept, all others are duplicates.
```

### Example 4: Single element
```go
Input:  [42]
Output: [42]

Trace: Loop doesn't run (length=1), return originalArray[:1] = [42]
```

## ⚡ Key Points

1. **Sorted Array Required**: This algorithm only works on sorted arrays
2. **In-Place Modification**: The original array is modified, but we return a slice
3. **Two Pointers**: 
   - `nextNonDup`: Tracks where to place the next unique element
   - `currentElement`: Scans through the array
4. **Time Complexity**: O(n) - single pass through array
5. **Space Complexity**: O(1) - only using constant extra space

## 🤔 Why does this work?

1. **Sorted Array Property**: In a sorted array, all duplicates are adjacent
2. **Skip Strategy**: We can safely skip duplicates because we've already seen that value
3. **Overwrite Strategy**: We overwrite the array in-place, keeping track of the "clean" portion

## 🚨 Important Notes

- **Input must be sorted**: `[3, 1, 2, 1]` won't work correctly
- **Modifies original array**: The input array is changed
- **Returns slice**: We return `originalArray[:nextNonDup]` to get only the unique part

## 🎮 Try it yourself

Test these cases:
- `[1, 1, 1, 1]` → `[1]`
- `[1, 2, 2, 3, 4, 4, 4, 5]` → `[1, 2, 3, 4, 5]`
- `[]` → `[]` (empty array)
- `[5]` → `[5]` (single element)