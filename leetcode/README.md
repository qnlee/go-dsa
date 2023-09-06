## Directory Structure
    leetcode/
        <difficulty>/
            lc#_problem_title/
                solution.go
                solution_test.go
        <random_helper_util_stuff>/
            <stuff>/
                ...

--- 
## How to Test
### Run all tests
> Not recommended, because this repo is quite large...

From project root, run:
* tests: `make test`
* bench: `make bench`

### Run tests for `leetcode/easy` 
From project root, run: 
* unit tests: `make test_lc_easy`
* bench: `make bench_lc_easy`

### Run tests for `leetcode/medium` 
From project root, run:
* unit tests: `make test_lc_medium`
* bench: `make bench_lc_medium`

---
## Problems Covered / Statuses 

### Leetcode Easy
| Leetcode # | Problem                                                                                       | Last Revisited |
|------------|-----------------------------------------------------------------------------------------------|----------------|
| 1          | [Two Sum](https://leetcode.com/problems/two-sum/)                                             | ;)             | 
| 20         | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)                         | 7/10/23        |
| 35         | [Search Insert Position](https://leetcode.com/problems/search-insert-position/)               | 7/10/23        |
| 70         | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                             | 5/22/23        |
| 104        | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)   | 8/31/23        |
| 121        | [Best Time to Buy Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | 5/2/23         |
| 125        | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)                           | 5/16/23        |
| 226        | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)                       | 8/31/23        |
| 290        | [Word Pattern](https://leetcode.com/problems/word-pattern/)                                   | 6/23/23        |
| 543        | [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)             | 9/1/23         |
| 704        | [Binary Search](https://leetcode.com/problems/binary-search/)                                 | 5/21/23        |

### Leetcode Medium
| Leetcode # | Problem                                                                                                                         | Last Revisited |
|------------|---------------------------------------------------------------------------------------------------------------------------------|----------------|
| 3          | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | 8/19/23        | 
| 11         | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/)                                           | 7/6/23         |
| 22         | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)                                                     | 7/12/23        |
| 33         | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)                                 | 6/14/23        |
| 36         | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)                                                                     | 6/14/23        |
| 73         | [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)                                                           | 6/14/23        |
| 74         | [Search 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)                                                           | 6/14/23        |
| 91         | [Decode Ways](https://leetcode.com/problems/decode-ways/)                                                                       | 6/23/23        |
| 150        | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)                             | 7/11/23        |
| 153        | [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)                     | 6/14/23        |
| 155        | [Min Stack](https://leetcode.com/problems/min-stack/)                                                                           | 7/7/23         |
| 347        | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)                                               | 6/14/23        |
| 424        | [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)               | 8/21/23        |
| 547        | [Number of Provinces](https://leetcode.com/problems/number-of-provinces/)                                                       | 7/21/23        |
| 673        | [Number of Longest Increasing Subsequences](https://leetcode.com/problems/number-of-longest-increasing-subsequence/)            | 7/21/23        |
| 739        | [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)                                                         | 7/12/23        |
| 853        | [Car Fleet](https://leetcode.com/problems/car-fleet/)                                                                           | 7/13/23        |
| 875        | [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)                                                       | 6/14/23        |
| 891        | [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/)                                         | 7/25/23        |
| 1492       | [The kth Factor of n ](https://leetcode.com/problems/the-kth-factor-of-n/)                                                      | 7/22/23        |
| 2405       | [Optimal Partition of a String](https://leetcode.com/problems/optimal-partition-of-string/)                                     | 8/17/23        |



### Leetcode Hard
| Leetcode # | Problem                                                                   | Last Revisited |
|------------|---------------------------------------------------------------------------|----------------|
| 42         | [Trapping Rainwater](https://leetcode.com/problems/trapping-rain-water/)  | 7/7/23         |

### Derivatives
| Problem                                     | Last Revisited | 
|---------------------------------------------|----------------|
| Check if input string follows given pattern | 6/23/23        |
| Particle Velocity                           | 6/23/23        |
| Transaction Tracker                         | 6/23/23        |
| Transition Words                            | 6/14/23        |
| Word Machine                                | 7/6/23         |