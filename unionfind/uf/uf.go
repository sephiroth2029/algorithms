package uf

// UF is the current Union-Find algorithm. For more details of the purpose of this
// algorithm check: http://algs4.cs.princeton.edu/15uf/
type UF struct {
	N  int
	id []int
	sz []int
}

// NewUF creates and initializes a new implementation of the algorithm.
func NewUF(N int) *UF {
	var uf UF

	uf.N = N
	uf.id = make([]int, N)
	uf.sz = make([]int, N)
	for i := 0; i < N; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}

	return &uf
}

func (uf *UF) root(i int) int {
	for i != uf.id[i] {
		i = uf.id[i]
	}

	return i
}

// Connected checks if two objects are in the same component
func (uf *UF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

// Union connects the given objects.
func (uf *UF) Union(p, q int) {
	i := uf.root(p)
	j := uf.root(q)

	if i == j {
		return
	}
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}
}
