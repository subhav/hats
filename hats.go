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
	return ((x % n) + n) % n
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
	fmt.Println(try(s2, 2))

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
	fmt.Println(try(s3, 3))

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
	fmt.Println(try(s3Alt, 3))

	// Let's generate these functions programmatically
	s3Gen := generateStrategy(3)
	fmt.Println(try(s3Gen, 3))

	// Wait, hold on. It works...?
	s4 := generateStrategy(4)
	fmt.Println(try(s4, 4))

	// But how?!
	s7 := generateStrategy(7)
	fmt.Println(try(s7, 7))
}

// Scaling my solution for 3 above:
// g_0 = v_1 + v_2 + ... + v_n
// g_i = v_0 - v_1 - ... - v_i - ... - v_n-1 + i
func generateStrategy(n int) []guess {
	g := make([]guess, n)
	g[0] = func(v []int) int {
		var res int
		for i := 1; i < n; i++ {
			res += v[i]
		}
		return mod(res, n)
	}
	for i := 1; i < n; i++ {
		i := i
		g[i] = func(v []int) int {
			res := v[0]
			for j := 1; j < n; j++ {
				if j == i {
					continue
				}
				res -= v[j]
			}
			res += i
			return mod(res, n)
		}
	}
	return g
}
