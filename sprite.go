package buildengine

type Sprite struct {
	parent *Map
	raw    rawSprite
}

type rawSprite struct {
	X int32
	Y int32
	Z int32

	Stat int16

	TexturePtr int16
	Shade      int8
	Pallette   uint8
	ClipDist   uint8
	Filler     uint8

	XRepeatTexture uint8
	YRepeatTexture uint8
	XTextureOffset uint8
	YTextureOffset uint8

	SectorNum int16
	Status    int16

	Angle int16
	Owner int16
	XVel  int16
	YVel  int16
	ZVel  int16

	Lotag int16
	Hitag int16
	Extra int16
}
