package getcap

import (
	"fmt"
	"strings"
)

func (a Abilities) String() string {
	var result string
	if a.Version != "" {
		result += "Version: " + a.Version + "\n"
	}
	if a.Name != "" {
		result += "Name: " + a.Name + "\n"
	}
	if a.Title != "" {
		result += "Title: " + a.Title + "\n"
	}
	if a.Abstract != "" {
		result += "Abstract: " + a.Abstract + "\n"
	}
	if a.Formats != nil {
		result += "Formats: " + a.Formats.String() + "\n"
	}
	if a.Layers != nil {
		result += "\nLayers:\n"
		result += a.Layers.String()
	}
	if a.BBoxes != nil {
		result += "\nBBoxes:\n"
		result += a.BBoxes.String()
	}
	return result
}

func (ff Formats) String() string {
	return strings.Join(ff, ", ")
}

func (ll Layers) String() string {
	var result string
	for i, l := range ll {
		result += l.String() + "\n"
		if i < len(ll)-1 {
			result += "\n"
		}
	}
	return result
}

func (l Layer) String() string {
	var result string
	if l.Name != "" {
		result += "Name: " + l.Name + "\n"
	}
	if l.Styles != nil {
		result += "Styles: " + l.Styles.String() + "\n"
	}
	if l.BBoxes != nil {
		result += l.BBoxes.String()
	}
	return result
}

func (ss Styles) String() string {
	return strings.Join(ss, ", ")
}

func (bb BBoxes) String() string {
	var result string
	for i, b := range bb {
		if b.GetEPSG() != 0 {
			result += fmt.Sprintf("EPSG:%v %v", b.GetEPSG(), b)
			if i < len(bb)-1 {
				result += "\n"
			}
		}
	}
	return result
}

func (b BBox) String() string {
	return fmt.Sprintf("minx: %f miny: %f maxx: %f maxy %f", b.MinX, b.MinY, b.MaxX, b.MaxY)
}
