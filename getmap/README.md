# GetMap for Go

Golang package to download tiles from a Web Map Service. With informative error handling.

```go
service, err := getmap.New(getmap.AddURL("http://ows.terrestris.de/osm/service"), getmap.AddFormat("image/png"))

service.AddEPSG(900913)
service.AddLayers("OSM-WMS")
...
r, width, height, err := service.GetMap(1110998.5364747, 7078816.8197878, 1119515.6176321, 7087113.4793097, getmap.WidthOption(1000))
// http://ows.terrestris.de/osm/service?SERVICE=WMS&REQUEST=GetMap&VERSION=1.1.1&FORMAT=image/png&LAYERS=OSM-WMS&STYLES=&SRS=EPSG:900913&WIDTH=1000&HEIGHT=1000&BBOX=1110998.5364747,7078816.8197878,1119515.6176321,7087113.4793097

// for example: github.com/disintegration/imaging
img, _ := imaging.Decode(r)
```