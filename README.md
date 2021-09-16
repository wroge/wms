[![GitHub release](https://img.shields.io/github/release/wroge/wms.svg)](https://github.com/wroge/wms/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/wroge/wms)](https://goreportcard.com/report/github.com/wroge/wms)

# Web Map Service - Manager

A command-line-tool to simplify the use of Web Map Services.
You can download WMS-Tiles and check the Capabilities of a service. Including:

- Set specific requests in a configuration file
- Get helpful error messages
- Automatic coordinate transformation into a supported reference system [by wroge/wgs84](https://github.com/wroge/wgs84)
- Download several bounding boxes at the same time

## Install

- Linux (i386/x86_64)
- MacOS (i386/x86_64)
- Windows (i386/x86_64)

[Releases](https://github.com/wroge/wms/releases)

Alternatively, you can install ```wms``` via Homebrew, Scoop, Snapcraft or Docker. Of course, you can also create the executable file from source.

### Homebrew (MacOS)

```
brew install wroge/tap/wms
```

### Scoop (Windows)

```
scoop bucket add app https://github.com/wroge/scoop-bucket
scoop install wms
```

### Snapcraft (Linux)

```
snap install wms
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

Supported by [wroge/wgs84](https://github.com/wroge/wgs84).  Some WMS allow only a few coordinate reference systems. With ```wms map``` you can choose from a larger number of EPSG codes. Please open an issue in the [wroge/wgs84](https://github.com/wroge/wgs84) repository to support your specific system.

```
wms cap -u http://ows.terrestris.de/osm/service -e
[4326 3857 900913]

wms map -u http://ows.terrestris.de/osm/service -e 12345
Error: Invalid EPSG: 12345
 Valid EPSGs: [900913 4326 3857 32650 32659 5650 25838 31466 32603 32613 32639 
        102014 6962 27700 32627 32633 32647 32652 32632 32636 3067 25829 25834 
        32605 32607 32622 32643 32654 4258 102008 102013 2154 5669 32630 32645 
        32655 102010 32656 32657 32614 32616 32619 32621 32642 32649 32638 32648 
        6870 31467 31468 32611 32628 32634 32601 32608 32612 32640 32651 32631 32660 
        5668 25833 25836 32609 32615 4647 25832 31469 32620 32641 32644 4277 25837 
        32610 32646 54027 32624 32629 32637 32658 4462 25830 25831 32625 32626 32635 
        4978 32606 32618 3395 25835 32602 32617 32623 32653 32604]
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
wms map -u http://ows.terrestris.de/osm/service -n example -e 25832 -l OSM-WMS/default --save terrestris --dry-run
Saving service: terrestris
URL: http://ows.terrestris.de/osm/service
Version: 1.1.1
Format: image/jpeg
Layers: [OSM-WMS]
Styles: [default]
EPSG: 25832
File name: example

wms cap terrestris
...

wms map terrestris -b 565000,5930000,570000,5935000 -w 1000
```

### Automatic image size calculation

GetMap-Requests require a ```WIDTH```and a ```HEIGHT```parameter. ```wms map``` calculates these parameters based on the UTM-size (ideal for small regions). You can define ```--height```, ```--width```or ```--dpi & --scale```.

```
wms map terrestris -b 565000,5930000,570000,5935000 --width 1000
http://ows.terrestris.de/osm/service?SERVICE=WMS&REQUEST=GetMap&VERSION=1.1.1&FORMAT=image/jpeg&LAYERS=OSM-WMS&STYLES=default&SRS=EPSG:900913&WIDTH=1000&HEIGHT=1000&BBOX=1110998.5409540,7078815.1864107,1119515.6232213,7087111.5778055
Done. Your requested file is here: <output>

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
