package main

var fullland = []byte{
	0b01111110,
	0b01111110,
	0b00111100,
	0b00111100,
	0b00011000,
	0b01111000,
	0b01111000,
	0b01111000,
}

type Land struct {
	Width      int
	Tick       int
	TickMax    int
	LeftWidth  int
	RightWidth int
}

const nubwidth = 20

func NewLand(boardwidth, boardheight int) Land {
	l := Land{}
	l.Width = 40
	l.Tick = 0
	l.TickMax = 1000
	return l
}
func (land *Land) render(boardwidth, boardheight int) {
	if land.Tick < land.TickMax {
		land.Tick++
	} else {
		land.Tick = 0
	}
	drawwidth := uint(land.Width)
	*DRAW_COLORS = 0x3
	if land.Tick >= 100 && land.Tick < 460 {

		rect(int(drawwidth), land.Tick-300, uint(nubwidth), uint(land.Tick-100))
		rect(boardwidth-int(drawwidth)-nubwidth, land.Tick-300, uint(nubwidth), uint(land.Tick-100))
		if land.Tick > 230 && land.Tick < 430 {
			land.LeftWidth = land.Width + nubwidth
			land.RightWidth = land.Width + nubwidth
		} else {
			land.LeftWidth = land.Width
			land.RightWidth = land.Width
		}
		//rect(boardwidth-int(drawwidth), 0, drawwidth, 160)
	} else {
		land.LeftWidth = land.Width
		land.RightWidth = land.Width
	}

	rect(0, 0, drawwidth, 160)
	rect(boardwidth-int(drawwidth), 0, drawwidth, 160)

}
func (land *Land) landViolationLeft(leftPos int) bool {
	if leftPos <= land.LeftWidth {
		return true
	}
	return false
}
func (land *Land) landViolationRight(rightPos int) bool {
	if rightPos >= 160-land.RightWidth-1 {
		return true
	}
	return false
}

func (land *Land) landViolationUp(topPos int) bool {
	return false
}
