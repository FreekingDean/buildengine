package buildengine

type Wall struct {
	parent *Map
	raw    rawWall
}

type rawWall struct {
	XLeft int32
	YLeft int32

	RightWallPtr int16
	OtherWall    int16
	NextSector   int16

	Stat int16

	TexturePtr     int16
	OverTexturePtr int16
	Shade          int8
	Pallette       uint8
	XRepeatTexture uint8
	YRepeatTexture uint8
	XTextureOffset uint8
	YTextureOffset uint8

	Lotag int16
	Hitag int16
	Extra int16
}

func (w *Wall) Left() *Point {
	return &Point{
		X: int(w.raw.XLeft),
		Y: int(w.raw.YLeft),
	}
}

func (w *Wall) Right() *Wall {
	if w.raw.RightWallPtr == -1 {
		return nil
	}
	return w.parent.walls[w.raw.RightWallPtr]
}
