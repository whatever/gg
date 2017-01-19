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

	e := NewQuadtree(
		NewQuadtreeLeaf(B, B, B, W),
		NewQuadtreeLeaf(W, W, B, W),
		NewQuadtreeWhite(),
		NewQuadtreeBlack(),
	)

	e_nw := e.NW
	e_ne := e.NE

	if e.Color != G || e.NW.Color != G || e.NE.Color != G || e.SE.Color != W || e.SW.Color != B {
		t.Fail()
	}

	if e_nw.NW.Color != B || e_nw.NE.Color != B || e_nw.SE.Color != B || e_nw.SW.Color != W {
		t.Fail()
	}

	if e_ne.NW.Color != W || e_ne.NE.Color != W || e_ne.SE.Color != B || e_ne.SW.Color != W {
		t.Fail()
	}
}
