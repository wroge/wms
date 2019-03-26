[![](https://godoc.org/github.com/wroge/wms/getcap?status.svg)](https://godoc.org/github.com/wroge/wms/getcap)

# GetCapabilities for Go

Golang package for reading ```GetCapabilities``` documents from Web Map Services. With extensive use of the stringer interface.

```go
cap, _ := getcap.From("http://ows.terrestris.de/osm/service", "1.1.1", "user", "password")

layers := cap.GetLayerNames()
// [OSM-WMS OSM-Overlay-WMS TOPO-WMS TOPO-OSM-WMS SRTM30-Hillshade SRTM30-Colored SRTM30-Colored-Hillshade SRTM30-Contour]

formats := cap.Formats
// image/jpeg, image/png

epsg := cap.GetBBoxes().GetEPSG()
// [900913 4326 3857]

ll := cap.GetLayers("OSM-WMS", "TOPO-WMS")
// Name: OSM-WMS
// Styles: default
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

// Name: TOPO-WMS
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

fmt.Println(cap)
// Version: 1.1.1
// Name: OGC:WMS
// Title: OpenStreetMap WMS
// Abstract: OpenStreetMap WMS, bereitgestellt durch terrestris GmbH und Co. KG. Beschleunigt mit MapProxy (http://mapproxy.org/)
// Formats: image/jpeg, image/png

// Layers:
// Name: OSM-WMS
// Styles: default
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

// Name: OSM-Overlay-WMS
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

// Name: TOPO-WMS
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

// Name: TOPO-OSM-WMS
// EPSG:900913 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500
// EPSG:4326 minx: -180.000000 miny: -88.000000 maxx: 180.000000 maxy 88.000000
// EPSG:3857 minx: -20037508.342800 miny: -25819498.513500 maxx: 20037508.342800 maxy 25819498.513500

// Name: SRTM30-Hillshade
// EPSG:900913 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820
// EPSG:4326 minx: -180.000000 miny: -56.000000 maxx: 180.000000 maxy 60.000000
// EPSG:3857 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820

// Name: SRTM30-Colored
// EPSG:900913 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820
// EPSG:4326 minx: -180.000000 miny: -56.000000 maxx: 180.000000 maxy 60.000000
// EPSG:3857 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820

// Name: SRTM30-Colored-Hillshade
// EPSG:900913 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820
// EPSG:4326 minx: -180.000000 miny: -56.000000 maxx: 180.000000 maxy 60.000000
// EPSG:3857 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820

// Name: SRTM30-Contour
// EPSG:900913 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820
// EPSG:4326 minx: -180.000000 miny: -56.000000 maxx: 180.000000 maxy 60.000000
// EPSG:3857 minx: -20037508.342800 miny: -7558415.656080 maxx: 20037508.342800 maxy 8399737.889820
```