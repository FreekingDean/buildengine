package buildengine

// Sector is similar to a room in the build engine.
// For more info visit http://www.shikadi.net/moddingwiki/MAP_Format_(Build)#File_format
// under sector.
type Sector struct {
	parent *Map
	raw    rawSector
}

type rawSector struct {
	WallPtr     int16
	WallCount   int16
	CeilingZ    int32
	FloorZ      int32
	CeilingStat int16
	FloorStat   int16

	CeilingTexturePtr     int16
	CeilingAngle          int16
	CeilingShade          int8
	CeilingPallete        uint8
	CeilingTextureXOffset uint8
	CeilingTextureYOffset uint8

	FloorTexturePtr     int16
	FloorAngle          int16
	FloorShade          int8
	FloorPallete        uint8
	FloorTextureXOffset uint8
	FloorTextureYOffset uint8

	Visibility uint8
	Filler     uint8

	Lotag int16
	Hitag int16
	Extra int16
}

func (s *Sector) Walls() []*Wall {
	return s.parent.walls[s.raw.WallPtr : s.raw.WallPtr+s.raw.WallCount]
}

func (s *Sector) FloorZ() int {
	return int(s.raw.FloorZ) / 16
}

func (s *Sector) CeilingZ() int {
	return int(s.raw.CeilingZ) / 16
}

func (s *Sector) CeilingStat() int {
	return int(s.raw.CeilingStat)
}

func (s *Sector) FloorStat() int {
	return int(s.raw.FloorStat)
}
