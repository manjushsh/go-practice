# Sliding Window - Subarray Sum 🪟

## 🎯 What Does This Do?

The **Sliding Window Subarray Sum** algorithm finds the sum of all possible consecutive groups of numbers in an array. It's like having a window that you slide across a row of numbers, calculating the sum inside the window at each position.

## 🏠 Real-World Analogy

Imagine you're looking at a row of houses through a window that shows exactly 3 houses at a time:

```
Houses: [🏠1] [🏠2] [🏠3] [🏠4] [🏠5]

Position 1: [🏠1] [🏠2] [🏠3] ⬜ ⬜  → Sum houses 1+2+3
Position 2: ⬜ [🏠2] [🏠3] [🏠4] ⬜  → Sum houses 2+3+4  
Position 3: ⬜ ⬜ [🏠3] [🏠4] [🏠5]  → Sum houses 3+4+5
```

Instead of counting all houses each time, you just remove the house that left the window and add the new house that entered!

## 📊 Step-by-Step Visual Example

**Input:** Array `[1, 2, 3, 4, 5]` with window size `3`

### 🖼️ Interactive Window Visualization

```
┌─────────────────────────────────────────────────────┐
│                 SLIDING WINDOW                      │
├─────────────────────────────────────────────────────┤
│ Array: [ 1 ] [ 2 ] [ 3 ] [ 4 ] [ 5 ]              │
│        ┌─────────────────┐                          │
│        │  WINDOW (k=3)   │                          │
│        └─────────────────┘                          │
│ Step 1: Window = [1, 2, 3]                         │
│         Sum = 1 + 2 + 3 = 6 ✅                     │
│         windowStart=0, windowEnd=2                  │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│ Array: [ 1 ] [ 2 ] [ 3 ] [ 4 ] [ 5 ]              │
│             ┌─────────────────┐                      │
│             │  WINDOW (k=3)   │                      │
│             └─────────────────┘                      │
│ Step 2: Remove 1 ⬅️, Add 4 ➡️                      │
│         Sum = 6 - 1 + 4 = 9 ✅                     │
│         windowStart=1, windowEnd=3                  │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│ Array: [ 1 ] [ 2 ] [ 3 ] [ 4 ] [ 5 ]              │
│                  ┌─────────────────┐                  │
│                  │  WINDOW (k=3)   │                  │
│                  └─────────────────┘                  │
│ Step 3: Remove 2 ⬅️, Add 5 ➡️                      │
│         Sum = 9 - 2 + 5 = 12 ✅                    │
│         windowStart=2, windowEnd=4                  │
└─────────────────────────────────────────────────────┘

📋 Final Result: [6, 9, 12]
```

### 🏎️ Performance Comparison Visualization

```
┌─────────────────────────────────────────────────────┐
│                NAIVE APPROACH                       │
├─────────────────────────────────────────────────────┤
│ Window 1: 1+2+3 = 6     [🔢🔢🔢]                   │
│ Window 2: 2+3+4 = 9     [🔢🔢🔢]                   │  
│ Window 3: 3+4+5 = 12    [🔢🔢🔢]                   │
│ Total Operations: 9     [💻💻💻💻💻💻💻💻💻]      │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│              SLIDING WINDOW                         │
├─────────────────────────────────────────────────────┤
│ Initial: 1+2+3 = 6      [🔢🔢🔢]                   │
│ Slide 1: 6-1+4 = 9      [➖➕]                      │
│ Slide 2: 9-2+5 = 12     [➖➕]                      │
│ Total Operations: 7     [💻💻💻💻💻💻💻]           │
│ Efficiency: 22% faster! 🚀                         │
└─────────────────────────────────────────────────────┘
```

### The Slow Way (For Comparison):
```
Step 1: 1+2+3 = 6     (3 additions)
Step 2: 2+3+4 = 9     (3 additions)  
Step 3: 3+4+5 = 12    (3 additions)
Total: 9 operations
```

### The Fast Way (Sliding Window):
```
Step 1: 1+2+3 = 6       (3 additions)
Step 2: 6-1+4 = 9       (1 subtraction + 1 addition)
Step 3: 9-2+5 = 12      (1 subtraction + 1 addition)
Total: 7 operations (faster!)
```

## 🔧 How the Code Works

### 1. Safety Check
```go
if windowSize < 0 || windowSize > len(originalElements) {
    return make([]int, 0)  // Return empty if invalid input
}
```
**Translation:** "If the window size doesn't make sense, return nothing"

### 2. Initialize Variables
```go
windowStart := 0                    // Left edge of our window
sumOfWindowElements := 0            // Current sum inside window
results := make([]int, (len(originalElements)-windowSize)+1)
```
**Translation:** 
- `windowStart`: "Where does our window begin?"
- `sumOfWindowElements`: "What's the total of numbers in our window?"
- `results`: "Array to store all our sums"

### 3. The Main Loop
```go
for windowEnd := 0; windowEnd < len(originalElements); windowEnd++ {
    // Add new element to sum
    sumOfWindowElements += originalElements[windowEnd]
    
    // If window is full size
    if windowEnd >= windowSize-1 {
        // Save current sum
        results[windowStart] = sumOfWindowElements
        
        // Slide window: remove left element
        sumOfWindowElements -= originalElements[windowStart]
        windowStart++
    }
}
```

**Translation:**
1. **Expand window:** Add the next number to our sum
2. **When window is full:** Save the sum and slide the window
3. **Slide window:** Remove the leftmost number and move window right

## 🚀 Why It's Efficient

| Method | Time Complexity | Explanation |
|--------|----------------|-------------|
| **Naive** | O(n × k) | For each position, add k numbers |
| **Sliding Window** | O(n) | Add each number once, subtract each number once |

**Example with 1000 numbers, window size 100:**
- Naive: 1000 × 100 = 100,000 operations
- Sliding Window: 1000 + 1000 = 2,000 operations (50x faster!)

## 🎮 Interactive Algorithm Trace

Let's trace through `[2, 1, 3, 4]` with window size `2`:

```
┌─────────────────────────────────────────────────────────────┐
│ INITIALIZATION                                              │
├─────────────────────────────────────────────────────────────┤
│ Array: [2] [1] [3] [4]                                     │
│ windowStart = 0 🔷                                          │
│ sum = 0                                                     │
│ results = [ _ ] [ _ ] [ _ ]  (3 positions available)       │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ ITERATION 1: windowEnd = 0                                 │
├─────────────────────────────────────────────────────────────┤
│ Array: [2] [1] [3] [4]                                     │
│         ↑🔴 windowEnd                                       │
│         ↑🔷 windowStart                                     │
│                                                             │
│ ➕ sum += originalElements[0] = 0 + 2 = 2                 │
│ ❓ windowEnd(0) >= windowSize-1(1)? → NO                  │
│ 📝 Continue to next iteration                              │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ ITERATION 2: windowEnd = 1                                 │
├─────────────────────────────────────────────────────────────┤
│ Array: [2] [1] [3] [4]                                     │
│         ↑🔷    ↑🔴 windowEnd                               │
│         windowStart                                         │
│                                                             │
│ ➕ sum += originalElements[1] = 2 + 1 = 3                 │
│ ✅ windowEnd(1) >= windowSize-1(1)? → YES                 │
│ 💾 results[0] = 3                                          │
│ ➖ sum -= originalElements[0] = 3 - 2 = 1                 │
│ ➡️ windowStart++ = 1                                       │
│ results = [3] [ _ ] [ _ ]                                   │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ ITERATION 3: windowEnd = 2                                 │
├─────────────────────────────────────────────────────────────┤
│ Array: [2] [1] [3] [4]                                     │
│              ↑🔷    ↑🔴 windowEnd                          │
│              windowStart                                    │
│                                                             │
│ ➕ sum += originalElements[2] = 1 + 3 = 4                 │
│ ✅ windowEnd(2) >= windowSize-1(1)? → YES                 │
│ 💾 results[1] = 4                                          │
│ ➖ sum -= originalElements[1] = 4 - 1 = 3                 │
│ ➡️ windowStart++ = 2                                       │
│ results = [3] [4] [ _ ]                                     │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ ITERATION 4: windowEnd = 3                                 │
├─────────────────────────────────────────────────────────────┤
│ Array: [2] [1] [3] [4]                                     │
│                   ↑🔷    ↑🔴 windowEnd                     │
│                   windowStart                               │
│                                                             │
│ ➕ sum += originalElements[3] = 3 + 4 = 7                 │
│ ✅ windowEnd(3) >= windowSize-1(1)? → YES                 │
│ 💾 results[2] = 7                                          │
│ ➖ sum -= originalElements[2] = 7 - 3 = 4                 │
│ ➡️ windowStart++ = 3                                       │
│ results = [3] [4] [7]                                       │
└─────────────────────────────────────────────────────────────┘

🎯 Final result: [3, 4, 7] ✨
```

## 🔑 Key Insights

1. **Window grows** until it reaches the desired size
2. **Window slides** by removing the left element and adding the right element
3. **Only 2 operations per slide** (1 subtraction + 1 addition) instead of recalculating everything
4. **Perfect for problems** involving consecutive elements in arrays

## 🎯 Common Use Cases

- Finding maximum sum of k consecutive elements
- Average of all subarrays of size k  
- Detecting patterns in time-series data
- Image processing (sliding over pixel windows)
- Network packet analysis

---

*This algorithm demonstrates the power of the sliding window pattern - turning an O(n×k) problem into an O(n) solution! 🚀*