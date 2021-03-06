package uf

import "testing"

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedAfterUnion(t *testing.T) {
	ufAlg := NewUF(10)

	ufAlg.Union(0, 1)
	if !ufAlg.Connected(0, 1) {
		t.Error("Union and Connected aren't working.")
	}
}

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedSymmetric(t *testing.T) {
	ufAlg := NewUF(10)

	ufAlg.Union(0, 5)
	if !ufAlg.Connected(0, 5) {
		t.Error("The symmetric property isn't working.")
	}

	if !ufAlg.Connected(5, 0) {
		t.Error("The symmetric property isn't working.")
	}
}

// TestConnected tests that the algorithm is properly identifying the reflective
// property.
func TestConnectedTransitive(t *testing.T) {
	ufAlg := NewUF(10)

	ufAlg.Union(0, 5)
	ufAlg.Union(5, 3)
	if !ufAlg.Connected(0, 3) {
		t.Error("The transitive property isn't working.")
	}
}
