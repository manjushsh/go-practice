# Smallest and Largest Subarray with Given Sum

## ELI5

### SmallestSubarrayWithGivenSum 🍬
**What it does:** Find the shortest row of numbers that adds up to AT LEAST your target.

**Simple analogy:** 
Imagine you're at a buffet with plates of food, each with a calorie count. You want to eat AT LEAST 500 calories but in the FEWEST plates possible.

- 🍽️ You have plates: [150, 100, 120, 200, 180, 50] calories
- 🎯 Goal: Get ≥ 500 calories using the FEWEST plates
- 🔍 Solution: Plates [200, 180, 120] = 500 calories (3 plates) ✅

**The smart way:** 
1. Start picking plates from left to right
2. When you hit 500+ calories, try removing plates from the left 
3. Keep the smallest group that still gives you 500+ calories

**How it works:**
1. 🍬 Start with two pointers: `leftPointer` and `rightPointer` both at the beginning
2. 📏 **Expand**: Move `rightPointer` right, adding numbers to `windowSum` until sum ≥ targetSum
3. ✂️ **Contract**: Try moving `leftPointer` right to make the window smaller while keeping sum ≥ targetSum
4. 📝 Remember the smallest window length that worked
5. 🔄 **Continue sliding**: Keep expanding/contracting until you've checked all possibilities

**🔑 Key insight:** The window size changes dynamically! It's NOT checking fixed sizes (2, then 3, then 4...). Instead, it naturally finds the optimal size by expanding when needed and contracting when possible.

**Example:**
- Numbers: [2, 1, 2, 4, 3, 1], TargetSum: 7
- **Expand phase**: [2] → [2,1] → [2,1,2] → [2,1,2,4] = 9 ≥ 7 ✅ (length 4)
- **Contract phase**: Remove 2: [1,2,4] = 7 ≥ 7 ✅ (length 3)
- **Contract more**: Remove 1: [2,4] = 6 < 7 ❌ (can't contract further)
- **Continue sliding**: [4,3] = 7 ≥ 7 ✅ (length 2) - Better!
- **Answer**: Smallest length is 2

**💡 Why this works for any target:** If targetSum was 10, the window would automatically expand to [2,1,2,4,3] (length 5), then contract to [1,2,4,3] (length 4). No need to manually try different sizes!

---

### MaxSubarrayWithGivenSum 🎯
**What it does:** Find the longest row of numbers that adds up to EXACTLY your target.

**Simple analogy:**
You're making a playlist where each song has a certain length (in minutes). You want a playlist that's EXACTLY 15 minutes long, but you want as MANY songs as possible.

- 🎵 Songs available: [3, 5, 2, 4, 1] minutes each
- 🎯 Goal: Get EXACTLY 15 minutes with the MOST songs
- 🔍 Solution: Songs [3, 5, 2, 4, 1] = 15 minutes (5 songs) ✅

**The smart way:**
1. Start adding songs from left to right
2. If you go over 15 minutes, remove songs from the left
3. When you hit exactly 15 minutes, check if this is your longest playlist yet

**How it works:**
1. 🍬 Start with two pointers: `leftPointer` and `rightPointer` both at the beginning
2. 📏 Move `rightPointer` right, adding numbers to `windowSum`
3. ✂️ When sum > exactSum, move `leftPointer` right to reduce the sum
4. 🎯 When sum = exactSum, check if this is the longest window so far
5. 🔄 Repeat until you've checked all possibilities

**Example:**
- Numbers: [1, 2, 1, 2, 1], ExactSum: 4
- Found groups that equal 4: [1, 2, 1], [2, 1, 2], [1, 2, 1]
- All have length 3, so answer is length 3

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
