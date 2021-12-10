package perm

type Perm struct {
	input []int
	current []int
	visited []bool
	results [][]int
}

func (p *Perm) calcPermutations() {
	if len(p.current) == len(p.input) {
		buffer := make([]int, len(p.current))
		copy(buffer, p.current)
		p.results = append(p.results, buffer)
	}

	for i := 0; i < len(p.input); i++ {
		if p.visited[i] {
			continue
		}

		if i > 0 && p.input[i] == p.input[i-1] && p.visited[i-1] {
			continue
		}

		p.visited[i] = true

		p.current = append(p.current, p.input[i])

		p.calcPermutations()

		p.visited[i] = false

		p.current = p.current[:len(p.current)-1]
	}
}

func (p *Perm) GetPermutations() [][]int {
	p.calcPermutations()
	return p.results
}

func New(input []int) *Perm {
	visited := make([]bool, len(input))

	for i := range visited {
		visited[i] = false
	}

	return &Perm{
		input: input,
		visited: visited,
	}
}
