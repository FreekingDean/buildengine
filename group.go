package buildengine

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type FileType int

const (
	UnknownFile FileType = -1
	MapFile     FileType = iota
)

type FileGroup struct {
	signature string
	files     []*File
}

type File struct {
	name     string
	data     []byte
	datasize uint32
}

func DecodeFileGroup(r io.Reader) (*FileGroup, error) {
	signature := make([]byte, 12)
	_, err := r.Read(signature)
	if err != nil {
		return nil, fmt.Errorf("Could not read signature: %e", err)
	}
	var totalFiles uint32
	err = binary.Read(r, endian, &totalFiles)
	if err != nil {
		return nil, fmt.Errorf("Could not read file count: %e", err)
	}

	files := make([]*File, totalFiles)
	for i := range files {
		filename := make([]byte, 12)
		_, err := r.Read(filename)
		if err != nil {
			return nil, fmt.Errorf("Could not read filename: %s", err.Error())
		}

		var datasize uint32
		err = binary.Read(r, endian, &datasize)
		if err != nil {
			return nil, fmt.Errorf("Could not read data size for file %s: %s", filename, err.Error())
		}

		fileString := strings.TrimRightFunc(string(filename), func(r rune) bool {
			return r == rune(0)
		})
		files[i] = &File{
			name:     fileString,
			datasize: datasize,
		}
	}
	for _, file := range files {
		data := make([]byte, file.datasize)
		_, err := r.Read(data)
		if err != nil {
			return nil, fmt.Errorf("Could not read data for file %s: %s", file.name, err.Error())
		}
		file.data = data
	}
	return &FileGroup{
		signature: string(signature),
		files:     files,
	}, nil
}

func (g *FileGroup) Maps() map[string]*Map {
	maps := make(map[string]*Map)
	for _, file := range g.files {
		if file.FileType() == MapFile {
			buf := bytes.NewBuffer(file.data)
			decodedMap, err := DecodeMap(buf)
			if err == nil {
				maps[file.name] = decodedMap
			}
		}
	}
	return maps
}

func (f *File) FileType() FileType {
	if strings.HasSuffix(f.name, "MAP") {
		return MapFile
	}

	return UnknownFile
}
