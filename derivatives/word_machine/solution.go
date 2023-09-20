package word_machine

import (
	"strconv"
	"strings"
)

/*
A word machine is a system that performs a sequence of simple operations on a stack of integers. Initially the stack is
empty. The sequence of operations is given as a string. Operations are separated by single spaces.

The following operations may be specified:
	an integer X (from 0 to 220 - 1): the machine pushes X onto the stack;
	"POP": the machine removes the topmost number from the stack;
	"DUP": the machine pushes a duplicate of the topmost number onto the stack;
	"+": the machine pops the two topmost elements from the stack, adds them together and pushes the sum onto the stack;
	"-": the machine pops the two topmost elements from the stack, subtracts the second one from the first (topmost)
		 one and pushes the difference onto the stack.

- After processing all the operations, the machine returns the topmost value from the stack.
- The machine processes 20-bit unsigned integers (numbers from 0 to 2^20 - 1).
- An overflow in addition or underflow in subtraction causes an error.
- The machine also reports an error when it tries to perform an operation that expects more numbers on the stack than the
  stack actually contains.
- Also, if, after performing all the operations, the stack is empty, the machine reports an error.

For example, given a string "13 DUP 4 POP 5 DUP + DUP + -", the machine performs the following operations:
	Operation | comment    										    | stack - - - - - - - - -------- [empty]
	"13"      | push 13      										| 13
	"DUP"     | duplicate 13 										| 13 13
	"4"       | push 4												| 13 13 4
	"POP"     | pop 4 												| 13 13
	"5"       | push 5 												| 13 13 5
	"DUP"     | duplicate 5 										| 13 13 5 5
	"+"       | pop 5 and 5 then adding sum of the two to stack 	| 13 13 10
	"DUP"     | duplicate 10 										| 13 13 10 10
	“+”       | pop 10 and 10 adding sum of the two to stack 		| 13 13 20
	“-“ 	  | pop 13 and 20, add difference (20-13) to stack		| 13 7
Finally, the machine will return 7.

Given a string "5 6 + -":
	the machine reports an error, since, after the addition, there is only one number on the stack,
	but the subtraction operation expects two.

Given a string "3 DUP 5 - -":
	the machine reports an error, since the second subtraction yields a negative result.

Write a function: def solution (S) that, given a string S containing a sequence of operations for the word machine,
returns the result the machine would return after processing the Operations.
The function should return -1 if the machine would report an error.

For example, given string S = "13 DUP 4 POP 5 DUP + DUP + - ":
	the function should return 7, as explained in the example above.

Given string S = "5 6 + -" or S = "3 DUP 5 - -":
	the function should return -1.

Assume that:
• the length of S is within the range (0..2,000);
• S is a valid sequence of word machine operations.
*/

type stack struct {
	elems []int
}

func (s *stack) topMost() int {
	if len(s.elems) == 0 {
		return -1
	}
	return s.elems[len(s.elems)-1]
}

func (s *stack) size() int {
	return len(s.elems)
}

func (s *stack) push(x int) {
	s.elems = append(s.elems, x)
}

func (s *stack) pop() int {
	if len(s.elems) == 0 {
		return -1
	}
	res := s.elems[len(s.elems)-1]
	s.elems = s.elems[:(len(s.elems) - 1)]
	return res
}

func solution(strOps string) int {
	ops := strings.Split(strOps, " ")
	if len(ops) == 0 {
		return -1
	}

	stack := &stack{
		elems: []int{},
	}

	for _, op := range ops {
		switch op {
		case "POP":
			topVal := stack.pop()
			if topVal == -1 {
				return -1
			}
		case "DUP":
			topVal := stack.topMost()
			if topVal == -1 {
				return -1
			}
			stack.push(topVal)
		case "+":
			if stack.size() < 2 {
				return -1
			}
			val1 := stack.pop()
			val2 := stack.pop()
			sum := val1 + val2
			if sum > 2^20-1 {
				return -1
			}
			stack.push(val1 + val2)
		case "-":
			if stack.size() < 2 {
				return -1
			}
			val1 := stack.pop()
			val2 := stack.pop()
			diff := val1 - val2
			if diff < 0 {
				return -1
			}
			stack.push(diff)
		default:
			val, err := strconv.Atoi(op)
			if err != nil || val < 0 || val > 2^20-1 {
				return -1
			}
			stack.push(val)
		}
	}

	return stack.topMost()
}
