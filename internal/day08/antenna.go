package day08

type Location []int

func (l Location) x() int {
	return l[0]
}

func (l Location) y() int {
	return l[1]
}

type Antenna struct {
	character string
	location  Location
}

func (a Antenna) equals(b Antenna) bool {
	if a.location.x() != b.location.x() ||
		a.location.y() != b.location.y() ||
		a.character != b.character {
		return false
	}

	return true
}
