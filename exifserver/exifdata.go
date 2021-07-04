package exifserver

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type ExifData struct {
	Data map[string]string
}

func NewExifData() *ExifData {
	return &ExifData{
		Data: make(map[string]string),
	}
}

func (e *ExifData) Walk(name exif.FieldName, tag *tiff.Tag) error {
	e.Data[string(name)] = tag.String()
	return nil
}
