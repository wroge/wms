// Package cap parses a GetCapabilities-Request
package getcap

import (
	"encoding/xml"
	"io"

	"github.com/wroge/wms/content"
)

// Formats of a GetCapabilities-Request
type Formats []string

// Styles of a GetCapabilities-Request
type Styles []string

// Layers of a GetCapabilities-Request
type Layers []Layer

// BBoxes of a GetCapabilities-Request
type BBoxes []BBox

// Abilities (Capabilities) of a GetCapabilities-Request
type Abilities struct {
	XMLName  xml.Name
	Version  string  `xml:"version,attr"`
	Name     string  `xml:"Service>Name"`
	Title    string  `xml:"Service>Title"`
	Abstract string  `xml:"Service>Abstract"`
	Formats  Formats `xml:"Capability>Request>GetMap>Format"`
	Layers   Layers  `xml:"Capability>Layer>Layer"`
}

// Layer of a GetCapabilities-Request
type Layer struct {
	Name   string `xml:"Name"`
	Styles Styles `xml:"Style>Name"`
	BBoxes BBoxes `xml:"BoundingBox"`
}

// BBox of a GetCapabilities-Request
type BBox struct {
	SRS  string  `xml:"SRS,attr"`
	CRS  string  `xml:"CRS,attr"`
	MinX float64 `xml:"minx,attr"`
	MinY float64 `xml:"miny,attr"`
	MaxX float64 `xml:"maxx,attr"`
	MaxY float64 `xml:"maxy,attr"`
}

// Get Capabilities of a WMS service
func From(url, version, user, password string) (Abilities, error) {
	var c Abilities
	request := url + "?SERVICE=WMS&REQUEST=GetCapabilities"
	if version != "" {
		request += "&VERSION=" + version
	}
	reader, err := content.From(request, user, password)
	if err != nil {
		return c, err
	}
	c, err = Read(reader)
	if err != nil {
		return c, err
	}
	return c, nil
}

// Read Capabilities from a GetCapabilities-Document
func Read(data io.Reader) (Abilities, error) {
	var c Abilities
	decoder := xml.NewDecoder(data)
	err := decoder.Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
