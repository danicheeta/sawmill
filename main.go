package main

import "fmt"

var current []int
var visited []bool

func main() {
	input := []int{1, 2}

	visited = make([]bool, len(input))
	for i := range visited {
		visited[i] = false
	}

	getPermutations(input)
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
