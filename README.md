[![Go Report Card](https://goreportcard.com/badge/github.com/wroge/wms)](https://goreportcard.com/report/github.com/wroge/wms)
[![GolangCI](https://golangci.com/badges/github.com/wroge/wms.svg)](https://golangci.com/r/github.com/wroge/wms)
[![GitHub release](https://img.shields.io/github/release/wroge/wms.svg)](https://github.com/wroge/wms/releases/latest)


# Web Map Service - Manager

A command-line-tool to simplify the use of Web Map Services.
You can download WMS-Tiles and check the Capabilities of a service. Including:

- Set specific requests in a configuration file
- Get helpful error messages
- Automatic Coordinate Transformation into a supported reference system
- Download several bounding boxes at the same time

## Install

- Linux (i386/x86_64)
- MacOS (i386/x86_64)
- Windows (i386/x86_64)

[Releases](https://github.com/wroge/wms/releases)

Alternatively, you can install ```wms``` via Homebrew, Scoop or Docker. Of course, you can also create the executable file from source.

### Homebrew (MacOS)

```
brew install wroge/tap/wms

// get latest version
brew upgrade wroge/tap/wms
```

### Scoop (Windows)

```
scoop bucket add app https://github.com/wroge/scoop-bucket
scoop install wms
```

### Docker

Docker-Images: [@DockerHub](https://hub.docker.com/r/wroge/wms/tags)

```
docker pull wroge/wms:latest
docker run -v "$(pwd)/output:/output" -v "$HOME/wms-config:/wms-config" wroge/wms
```

### From Source

This Go-Project is using ```go mod```. Please clone this repository outside of ```GOPATH```.

```
git clone https://github.com/wroge/wms.git
cd wms
go build -o wms ./cli
```

## Features

### Helpful error-messages

You can look up the capabilities ```wms cap``` or just try it out.

```
wms cap -u http://ows.terrestris.de/osm/service -l
[OSM-WMS OSM-Overlay-WMS TOPO-WMS TOPO-OSM-WMS SRTM30-Hillshade SRTM30-Colored SRTM30-Colored-Hillshade SRTM30-Contour]

wms map -u http://ows.terrestris.de/osm/service -l abc
Error: Invalid Layer: abc
Valid Layers: [OSM-WMS OSM-Overlay-WMS TOPO-WMS TOPO-OSM-WMS SRTM30-Hillshade SRTM30-Colored SRTM30-Colored-Hillshade SRTM30-Contour]
```

### Automatic Coordinate Transformation

Supported by [go-coo](https://github.com/wroge/go-coo).  Some WMS allow only a few coordinate reference systems. With ```wms map``` you can choose from a larger number of EPSG codes. Please open an issue in the [go-coo](https://github.com/wroge/go-coo) repository to support your specific system.

```
wms cap -u http://ows.terrestris.de/osm/service -e
[4326 3857 900913]

wms map -u http://ows.terrestris.de/osm/service -e 12345
Error: Invalid EPSG: 12345
Valid EPSGs: [2154 4326 5668 31467 32632 32633 5650 4277 3857 25832 25833 31468 3067 4462 6870 900913 4647 27700 5669 31466 31469 6962]
```

### Download several bounding boxes

You can create a text file that contains all the required bounding boxes and download them concurrently. (Using Goroutines)

```
cat $HOME/bbox-wgs84.txt
9,52,9.2,52.2
9.2,52,9.4,52.2
9.4,52,9.6,52.2

wms map -u http://ows.terrestris.de/osm/service -B $HOME/bbox-wgs84.txt -w 1000 -e 4326
```

### Configuration

You can change the configuration using ```--save``` or edit the configuration file ```$HOME/wms-config/.wms.yaml``` directly. 

```
wms map -u http://ows.terrestris.de/osm/service -n example -e 25832 -l OSM-WMS --save terrestris --dry-run
Saving service: terrestris
URL: http://ows.terrestris.de/osm/service
Version: 1.1.1
Format: image/jpeg
Layers: [OSM-WMS]
Styles: []
EPSG: 25832
File name: example

wms cap terrestris
...

wms map terrestris -b 565000,5930000,570000,5935000 -w 1000
```

### Automatic image size calculation

GetMap-Requests require a ```WIDTH```and a ```HEIGHT```parameter. ```wms map``` calculates these parameters based on the UTM-size. You can define ```--height```, ```--width```or ```--dpi & --scale```.

```
wms map terrestris -b 565000,5930000,570000,5935000 --width 1000
wms map terrestris -b 565000,5930000,570000,5935000 --height 1000
wms map terrestris -b 565000,5930000,570000,5935000 --dpi 100 --scale 10000
```

### More

- Expand & Cut bounding boxes (interesting for dynamically generated texts and symbols)
- Basic Authentication with ```--user```&```--password```
- ```wms map --help```
- ```wms cap --help```

## FAQ

...For any problems/questions please open an issue.