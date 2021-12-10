package main

import (
	"fmt"
)

var current []int
var visited []bool

func main() {
	input := []int{2, 3, 1}

	visited = make([]bool, len(input))
	for i := range visited {
		visited[i] = false
	}

	println(getProfit(input, 3))
	// getPermutations(input)
}

func getProfit(input []int, del int) int {
	var profit int
	var curr int

	for _, i := range input {
		if (curr + i) > 4 {
			curr = curr + i - 4
			profit += profitOnBlock(i - curr)
			profit += profitOnBlock(curr)
		} else {
			profit += profitOnBlock(i)
			curr += i

			if curr == 4 {
				curr = 0
			}
		}
	}

	return profit
}

func profitOnBlock(l int) int {
	switch l {
	case 1:
		return -1
	case 2:
		return 3
	case 3:
		return 1
	case 4:
		return 6
	default:
		return 0
	}
}

func getPermutations(input []int) {
	if len(current) == len(input) {
		fmt.Printf("%+v", current)
	}

	for i := 0; i < len(input); i++ {
		if visited[i] {
			continue
		}

		if i > 0 && input[i] == input[i-1] && visited[i-1] {
			continue
		}

		visited[i] = true

		current = append(current, input[i])

		getPermutations(input)

		visited[i] = false

		current = current[:len(current)-1]
	}
}

// func readFromStdin() map[int][][]int {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var res map[int][][]int
// 	var counter int

// 	for scanner.Scan() {
// 		riverLines, err := strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			panic(err)
// 		}

// 		if riverLines == 0 {
// 			break
// 		}

// 		riversInput := make([][]int, riverLines)
// 		for i := 0; i < riverLines; i++ {
// 			scanner.Scan()
// 			riversInput = append(riversInput, stringArrayToIntArray(scanner.Text()))
// 		}

// 		res[counter] = make([][]int, riverLines)
// 		res[counter] = riversInput
// 		counter++
// 	}

// 	return res
// }

// func stringArrayToIntArray(s string) []int {
// 	var err error

// 	strs := strings.Split(s, " ")
// 	ary := make([]int, len(strs))
// 	for i := range ary {
// 		ary[i], err = strconv.Atoi(strs[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return ary
// }
