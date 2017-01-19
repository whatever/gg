package gg

type QuadtreeType int

const (
	W QuadtreeType = 0
	B QuadtreeType = 1
	G QuadtreeType = 2
)

/*
Quadtree interface
*/

type Quadtree struct {
	Color QuadtreeType
	NW    *Quadtree
	NE    *Quadtree
	SE    *Quadtree
	SW    *Quadtree
}

func NewQuadtreeBlack() *Quadtree {
	return &Quadtree{B, nil, nil, nil, nil}
}

func NewQuadtreeWhite() *Quadtree {
	return &Quadtree{W, nil, nil, nil, nil}
}

func NewQuadtreeWithColor(color QuadtreeType) *Quadtree {
	return &Quadtree{color, nil, nil, nil, nil}
}

func NewQuadtreeLeaf(nw, ne, se, sw QuadtreeType) *Quadtree {
	return &Quadtree{
		G,
		NewQuadtreeWithColor(nw),
		NewQuadtreeWithColor(ne),
		NewQuadtreeWithColor(se),
		NewQuadtreeWithColor(sw),
	}
}
