package trees

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestTreesAreTheSame(t *testing.T) {
	t1, t2 := tree.New(5), tree.New(5)

	if !Same(t1, t2) {
		t.Error("Trees should be the same")
	}
}

func TestTreesAreNotTheSame1(t *testing.T) {
	t1, t2 := tree.New(5), tree.New(6)

	if Same(t1, t2) {
		t.Error("Trees should not be the same")
	}
}

func TestTreesAreNotTheSame2(t *testing.T) {
	t1, t2 := tree.New(6), tree.New(5)

	if Same(t1, t2) {
		t.Error("Trees should not be the same")
	}
}
