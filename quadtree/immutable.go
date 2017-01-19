package gg

type Qtype int

const (
	WHITE Qtype = 0
	BLACK       = 1
	GRAY        = 2
)

type Qregion int

const (
	NW Qregion = 0
	NE Qregion = 1
	SE Qregion = 2
	SW Qregion = 3
)

type Qtree struct {
	/*
		WW -> W
		BW -> B
		BB -> B
		GB -> *RECURSE*
		GW -> *RECURSE*
	*/
	Color Qtype

	/*
		If gray, regions look l ike:
			+----+----+
			| NW | NE |
			+----+----+
			| SW | SE +
			+----+----+
	*/
	Regions []Qtree
}

func QtreeBlack() Qtree {
	return Qtree{BLACK, nil}
}

func QtreeWhite() Qtree {
	return Qtree{WHITE, nil}
}

func NewQtree(nw, ne, se, sw Qtype) Qtree {
	if nw == ne && ne == se && se == sw && sw == nw {
		return Qtree{nw, nil}
	}

	return Qtree{
		GRAY,
		[]Qtree{
			Qtree{nw, nil},
			Qtree{ne, nil},
			Qtree{se, nil},
			Qtree{sw, nil},
		},
	}
}
