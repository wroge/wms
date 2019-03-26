package getcap

import (
	"strconv"
	"strings"
)

// GetLayers returns Layers based on specific layer names
func (a Abilities) GetLayers(layers ...string) Layers {
	var result Layers
	for _, l := range layers {
		if a.GetLayer(l).Name != "" {
			result = append(result, a.GetLayer(l))
		}
	}
	return result
}

// GetLayer returns a Layer based on a specific layer name
func (a Abilities) GetLayer(layer string) Layer {
	return a.Layers.GetLayer(layer)
}

// GetLayer returns a Layer based on a specific layer name
func (ll Layers) GetLayer(layer string) Layer {
	for _, l := range ll {
		if l.Name == layer {
			return l
		}
	}
	return Layer{}
}

// GetBBox returns a BBox based on a specific EPSG code
func (a Abilities) GetBBox(epsg int) BBox {
	return a.Layers.GetBBox(epsg)
}

// GetBBox returns a BBox based on a specific EPSG code
func (ll Layers) GetBBox(epsg int) BBox {
	return ll.GetBBoxes().GetBBox(epsg)
}

// GetBBox returns a BBox based on a specific EPSG code
func (l Layer) GetBBox(epsg int) BBox {
	return l.BBoxes.GetBBox(epsg)
}

// GetBBox returns a BBox based on a specific EPSG code
func (bb BBoxes) GetBBox(epsg int) BBox {
	var result BBox
	for _, b := range bb {
		if b.GetEPSG() == epsg {
			result = b
			break
		}
	}
	return result
}

// GetBBoxes returns BBoxes merged from all Layers
func (a Abilities) GetBBoxes() BBoxes {
	return a.Layers.GetBBoxes()
}

// GetBBoxes returns BBoxes merged from Layers
func (ll Layers) GetBBoxes() BBoxes {
	bboxMap := make(map[string]BBox)
	for _, l := range ll {
		for _, b := range l.BBoxes {
			if strings.Contains(b.CRS+b.SRS, "EPSG") {
				bboxMap[b.CRS+b.SRS] = b
			}
		}
	}
	var result BBoxes
	for _, b := range bboxMap {
		result = append(result, b)
	}
	return result
}

// GetEPSG returns EPSG codes merged from BBoxes
func (bb BBoxes) GetEPSG() []int {
	var result []int
	for _, b := range bb {
		if b.GetEPSG() != 0 {
			result = append(result, b.GetEPSG())
		}
	}
	return result
}

// GetEPSG returns the EPSG code from a BBox
func (b BBox) GetEPSG() int {
	crs := b.CRS + b.SRS
	split := strings.Split(crs, ":")
	if len(split) != 2 {
		return 0
	}
	if split[0] != "EPSG" {
		return 0
	}
	epsg, err := strconv.Atoi(split[1])
	if err != nil {
		return 0
	}
	return epsg
}

// GetLayerNames returns all layer names
func (a Abilities) GetLayerNames() []string {
	return a.Layers.GetLayerNames()
}

// GetLayerNames returns all layer names
func (ll Layers) GetLayerNames() []string {
	var result []string
	for _, l := range ll {
		result = append(result, l.Name)
	}
	return result
}
