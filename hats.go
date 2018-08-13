// Verifier for the n-hats puzzle.
//
// The size-7 variant of the puzzle can be stated as:
//
// -   You and 6 others are wearing one of 7 colored hats.
// -   You can see everyone else's hat color, but not your own.
// -   Everyone guesses what their hat color is.
// -   Everyone wins if one person guesses correctly.
// -   You all can decide on a strategy before-hand.
// -   Find a strategy that guarantees a win.
// 
// This program enumerates all possible outcomes and checks a
// strategy against them.
//
// We encode a hat color as an int in 0..n-1.
// We define a "guess" as a function that takes a vector of everyone
// else's hat color and returns a participant's guessed color.
// We define a "strategy" as a list of guesses, one for each
// participant.

package main

import "fmt"

// const n = 3

type guess func(actual []int) int

// Clear out a position to make sure the strategy actually works
func prune(v []int, i int, n int) []int {
	u := make([]int, n)
	copy(u, v)
	u[i] = -1
	return u
}

func checkRow(actual []int, strategy []guess, n int) bool {
	for i := 0; i < n; i++ {
		if actual[i] == strategy[i](prune(actual, i, n)) {
			return true
		}
	}
	return false
}

func printRow(actual []int, strategy []guess, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(actual[i], " ")
	}
	fmt.Print("| ")
	for i := 0; i < n; i++ {
		fmt.Print(strategy[i](actual), " ")
	}
	fmt.Println()
}

func try(strategy []guess, n int) bool {
	cur := make([]int, n)
	// Iterator which returns whether more values exist
	next := func() bool {
		for i := n-1; i >= 0; i-- {
			cur[i] = (cur[i] + 1) % n
			if cur[i] > 0 {
				return true
			}
		}
		return false
	}

	// do-while, since we already have the first one
	for i := true; i; i = next() {
		if !checkRow(cur, strategy, n) {
			printRow(cur, strategy, n)
			return false
		}
	}
	return true
}

// Positive Modulus
func mod(x int, n int) int {
	return (x + n) % n
}

func main() {
	s2 := []guess{
		func(v []int) int {
			return v[1]
		},
		func(v []int) int {
			return (v[0] + 1) % 2
		},
	}

	// Strategy which I drew out on paper for 3
	s3 := []guess{
		func(v []int) int {
			//       g_A = C - B + 1
			return mod(v[2] - v[1] + 1, 3)
		},
		func(v []int) int {
			//       g_B = C - A + 2
			return mod(v[2] - v[0] + 2, 3)
		},
		func(v []int) int {
			//       g_C = A + B
			return mod(v[0] + v[1], 3)
		},
	}
	// Alternate version, trying to move symbols around
	// and see what works.
	s3Alt := []guess{
		func(v []int) int {
			//      g_A = B + C
			return mod(v[1] + v[2], 3)
		},
		func(v []int) int {
			//      g_B = A - C + 1
			return mod(v[0] - v[2] + 1, 3)
		},
		func(v []int) int {
			//      g_C = A - B + 2
			return mod(v[0] - v[1] + 2, 3)
		},
	}

	// Let's generate these functions programmatically
	s3Gen := make([]guess, 3)
	for i := 0; i < 3; i++ {
		s3Gen[i] = func(v []int) int {
			var res int
			for j := 0; j < 3; j++ {
				if i == j {
					continue
				}
				if j < i {
					res -= j
				} else if j > i {
					res += j
				}
			}
			res += i
			return mod(res, 3)
		}
	}

	s4 := make([]guess, 4)
	for i := 0; i < 4; i++ {
		i := i
		s4[i] = func(v []int) int {
			return v[i]
		}
	}

	fmt.Println(try(s2, 2))
	fmt.Println(try(s3, 3))
	fmt.Println(try(s3Alt, 3))
	fmt.Println(try(s3Gen, 3))
	fmt.Println(try(s4, 4))

}
