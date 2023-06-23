package lc91_decode_ways

/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:
	'A' -> "1"
	'B' -> "2"
	...
	'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the
mapping above (there may be multiple ways). For example, "11106" can be mapped into:
	"AAJF" with the grouping (1 1 10 6)
	"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.
The test cases are generated so that the answer fits in a 32-bit integer.

Example 1:
	Input: s = "12"
	Output: 2
	Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:
	Input: s = "226"
	Output: 3
	Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:
	Input: s = "06"
	Output: 0
	Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

Constraints:
	1 <= s.length <= 100
	s contains only digits and may contain leading zero(s).
*/

func numDecodingsTopDown(s string) int {
	d := dp{memo: map[string]int{}}
	return d.numWays(s)
}

type dp struct {
	// Why using a struct (dp) with a map field (memo) is preferred to just using a type map[string]int:
	// > Encapsulation:
	//		By using a struct, you encapsulate the memoization map within the dp type.
	//		This allows you to keep the map private/controlled within struct, preventing external modifications to it.
	// > Stateful behavior:
	//		With a struct, you can maintain the state of the map across multiple function calls.
	//		The struct can store the map as a field, & any modifications to the map will persist throughout the life
	//		of the dp object.
	//		This allows you to reuse the map for different inputs and efficiently retrieve memoized results.
	// > Extensibility:
	//		Using a struct provides flexibility to add more fields or methods to the dp type in the future if needed,
	//		like including additional data or behavior, which would be challenging with just a map type.
	memo map[string]int
}

func (dp *dp) numWays(s string) int {
	// If previously calculated, return from memo
	if v, ok := dp.memo[s]; ok {
		return v
	}
	// Base case: if s == "" || s is some digit in range [1,9] --> only 1 decoding
	if len(s) <= 1 && s != "0" {
		return 1
	}
	// Base case: if s starts with '0' --> no valid decoding
	if s[0] == '0' {
		return 0
	}

	// Evaluate # ways to decode for substring excluding the first digit (eg. s="123" -> dp("1") + dp("23"))
	ways := dp.numWays(s[1:])
	// If first 2 digits form num in valid range [1,26], evaluate for substring excluding first two digits
	// (eg. s="123" -> dp("12) + dp("3")) and add to total # ways
	if s[0] == '1' || (s[0] == '2' && s[1] < '7') {
		ways += dp.numWays(s[2:])
	}

	dp.memo[s] = ways
	return ways
}

func numDecodingsBottomUp(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}

	var prev, curr int // Variables to store the previous two results
	curr = 1           // Initialize curr to 1 since there's at least one valid decoding for non-empty string
	for i := 1; i <= len(s); i++ {
		ways := 0
		// Check if the current character can be decoded as a valid single digit
		if s[i-1] != '0' {
			ways = curr
		}
		// Check if the current and previous characters can be decoded as a valid two-digit number
		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			ways += prev
		}
		// Update prev and curr for the next iteration
		prev = curr
		curr = ways
	}

	return curr
}
