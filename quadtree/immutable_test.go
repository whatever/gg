package gg

import (
	"fmt"
	"testing"
)

func _(t *testing.T) {
	fmt.Println("x_x")
}

func TestSimpleQuadtree(t *testing.T) {
	a := Qtree{BLACK, nil}
	b := Qtree{WHITE, nil}
	c := Qtree{GRAY, nil}

	if a.Color != BLACK || b.Color != WHITE || c.Color != GRAY {
		t.Fail()
	}

	d := Qtree{
		GRAY,
		[]Qtree{
			Qtree{BLACK, nil},
			Qtree{WHITE, nil},
			Qtree{BLACK, nil},
			Qtree{WHITE, nil},
		},
	}

	switch {
	case d.Regions[NW].Color != BLACK:
		t.Fail()
	case d.Regions[NE].Color != WHITE:
		t.Fail()
	case d.Regions[SE].Color != BLACK:
		t.Fail()
	case d.Regions[SW].Color != WHITE:
		t.Fail()
	}
}

func TestQtreeHelpers(t *testing.T) {
	a := QtreeBlack()
	b := QtreeWhite()

	if a.Color != BLACK || b.Color != WHITE {
		t.Fail()
	}

	c := NewQtree(BLACK, WHITE, BLACK, WHITE)

	if c.Regions[NW].Color != BLACK || c.Regions[NE].Color != WHITE || c.Regions[SE].Color != BLACK || c.Regions[SW].Color != WHITE {
		fmt.Println(c.Regions)
		t.Fail()
	}

	d := NewQtree(BLACK, BLACK, BLACK, BLACK)

	if d.Color != BLACK {
		t.Fail()
	}

	e := NewQtree(WHITE, WHITE, WHITE, WHITE)

	if e.Color != WHITE {
		t.Fail()
	}
}
