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
| 70         | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                             | 5/22/23        |
| 121        | [Best Time to Buy Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | 5/2/23         |
| 125        | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)                           | 5/16/23        |
| 290        | [Word Pattern](https://leetcode.com/problems/word-pattern/)                                   | 6/23/23        |
| 704        | [Binary Search](https://leetcode.com/problems/binary-search/)                                 | 5/21/23        |

### Leetcode Medium
| Leetcode # | Problem                                                                                                     | Last Revisited |
|------------|-------------------------------------------------------------------------------------------------------------|----------------|
| 33         | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)             | 6/14/23        |
| 36         | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/description/)                                     | 6/14/23        |
| 73         | [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)                                       | 6/14/23        |
| 74         | [Search 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)                                       | 6/14/23        |
| 91         | [Decode Ways](https://leetcode.com/problems/decode-ways/description/)                                       | 6/23/23        |
| 153        | [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | 6/14/23        |
| 347        | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)                           | 6/14/23        |
| 875        | [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)                                   | 6/14/23        |

### Leetcode Hard
TODO

### Derivatives
| Problem                                     | Last Revisited | 
|---------------------------------------------|----------------|
| Check if input string follows given pattern | 6/23/23        |
| Particle Velocit y                          | 6/23/23        |
| Transaction Tracker                         | 6/23/23        | 
| Transition Words                            | 6/14/23        | 