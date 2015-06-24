package uf

// UF is the current Union-Find algorithm. For more details of the purpose of this
// algorithm check: http://algs4.cs.princeton.edu/15uf/
type UF struct {
	N  int
	id []int
}

// NewUF creates and initializes a new implementation of the algorithm.
func NewUF(N int) *UF {
	var uf UF

	uf.N = N
	uf.id = make([]int, N)
	for i := 0; i < N; i++ {
		uf.id[i] = i
	}

	return &uf
}

// Connected checks if two objects are in the same component
func (uf *UF) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

// Union connects the given objects.
func (uf *UF) Union(p, q int) {
	pid := uf.id[p]
	qid := uf.id[q]

	for i := range uf.id {
		if uf.id[i] == pid {
			uf.id[i] = qid
		}
	}
}
