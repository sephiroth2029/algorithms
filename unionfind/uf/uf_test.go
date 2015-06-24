package uf

import "testing"

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedReflective(t *testing.T) {
	var uf UF
	uf.N = 2

	if !uf.Connected(0, 0) {
		t.Error("The reflective property isn't working.")
	}
}

// TestUnion tests that the algorithm is properly connecting two objects.
func TestUnion(t *testing.T) {
	var uf UF
	uf.N = 2

	uf.Union(0, 1)
}
