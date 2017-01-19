package gg

import (
	"fmt"
	"testing"
)

func _(t *testing.T) {
	fmt.Println("x_x")
}

func TestBasic(t *testing.T) {
	a := Quadtree{B, nil, nil, nil, nil}
	b := Quadtree{W, nil, nil, nil, nil}

	if a.Color != B || b.Color != W {
		t.Fail()
	}

	c := NewQuadtreeLeaf(B, B, W, W)

	if c.NW.Color != B || c.NE.Color != B || c.SE.Color != W || c.SW.Color != W {
		t.Fail()
	}

	d := NewQuadtreeLeaf(W, W, B, W)

	if d.NW.Color != W || d.NE.Color != W || d.SE.Color != B || d.SW.Color != W {
		t.Fail()
	}
}
