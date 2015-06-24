package uf

import "testing"

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedReflective(t *testing.T) {
	ufAlg := NewUF(2)

	if !ufAlg.Connected(0, 0) {
		t.Error("The reflective property isn't working.")
	}
}

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedAfterUnion(t *testing.T) {
	ufAlg := NewUF(10)

	ufAlg.Union(0, 1)
	if !ufAlg.Connected(0, 1) {
		t.Error("Union and Connected aren't working.")
	}
}
