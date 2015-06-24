package uf

// UF is the current Union-Find algorithm. For more details of the purpose of this
// algorithm check: http://algs4.cs.princeton.edu/15uf/
type UF struct {
	N int64
}

// Connected checks if two objects are in the same component
func (*UF) Connected(p, q int64) bool {
	if p == q {
		return true
	}
	return false
}

// Union connects the given objects.
func (*UF) Union(p, q int64) {

}
