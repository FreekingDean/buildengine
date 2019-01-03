package buildengine

import (
	"encoding/binary"
	"io"
)

type rawMap struct {
	Version     int32
	Xpos        int32
	YPoz        int32
	ZPos        int32
	Angle       int16
	StartSector int16
}

// Map is an object containing all information about a build map
type Map struct {
	raw     rawMap
	sectors []*Sector
	walls   []*Wall
	sprites []*Sprite
}

func DecodeMap(r io.Reader) (*Map, error) {
	newMap := &Map{}
	err := binary.Read(r, endian, &newMap.raw)
	if err != nil {
		return nil, err
	}

	var numsec uint16
	err = binary.Read(r, endian, &numsec)
	if err != nil {
		return nil, err
	}
	rawSectors := make([]rawSector, numsec)
	err = binary.Read(r, endian, rawSectors)
	for _, rawSector := range rawSectors {
		newMap.sectors = append(newMap.sectors, &Sector{
			parent: newMap,
			raw:    rawSector,
		})
	}

	var numwalls uint16
	err = binary.Read(r, endian, &numwalls)
	if err != nil {
		return nil, err
	}
	rawWalls := make([]rawWall, numwalls)
	err = binary.Read(r, endian, rawWalls)
	if err != nil {
		return nil, err
	}
	for _, rawWall := range rawWalls {
		newMap.walls = append(newMap.walls, &Wall{
			parent: newMap,
			raw:    rawWall,
		})
	}

	var numsprites uint16
	err = binary.Read(r, endian, &numsprites)
	if err != nil {
		return nil, err
	}
	rawSprites := make([]rawSprite, numsprites)
	err = binary.Read(r, endian, rawSprites)
	if err != nil {
		return nil, err
	}
	for _, rawSprite := range rawSprites {
		newMap.sprites = append(newMap.sprites, &Sprite{
			parent: newMap,
			raw:    rawSprite,
		})
	}

	return newMap, nil
}

func (m *Map) Sectors() []*Sector {
	return m.sectors
}

func (m *Map) StartSector() *Sector {
	return m.sectors[m.raw.StartSector]
}
