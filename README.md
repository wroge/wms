# Web Map Service - Command Line Tool

This Command Line Tool helps you to manage Web Map Services.

- Configuration-File ```$HOME/wms-config/.wms.yaml```
- Support-Messages
- Automatic Coordinate-Transformation by [go-coo](https://github.com/wroge/go-coo)
- Download multiple Bounding Boxes at once

## Install

### Download

- Linux (i386/x86_64)
- MacOS (i386/x86_64)
- Windows (i386/x86_64)

[Releases](https://github.com/wroge/wms/releases)

### Docker

You have Docker installed? Then try this commands.

```console
docker run -v "$(pwd)/output:/output" wroge/wms map -u http://ows.terrestris.de/osm/service -e 25832 -b 565000,5930000,570000,5935000 -w 1000
docker run wroge/wms cap -u http://ows.terrestris.de/osm/service
```
Docker-Images [@DockerHub](https://hub.docker.com/r/wroge/wms)

[Example-Script](https://github.com/wroge/wms/blob/master/wms) for better usage. (tested for macOS)

```console
echo '#!/bin/sh

if [[ $1 == "upgrade" ]]; then
    docker pull -a wroge/wms
    exit 0
fi

if [ -z $VERSION ]; then
    VERSION+=":latest"
else
    VERSION=":${VERSION}"
fi

docker run -v "$(pwd)/output:/output" -v "$HOME/wms-config:/wms-config" wroge/wms$VERSION $@' > ${PATH%%:*}/wms && chmod +x ${PATH%%:*}/wms
```

## Configuration

```$HOME/wms-config/.wms.yaml```

Example [Configuration](https://github.com/wroge/wms/blob/master/wms-config/.wms.yaml). (tested for macOS)

```console
mkdir -m 777 ${HOME}/wms-config && echo 'terrestris:
  epsg: 25832
  file-name: example
  format: ""
  layers:
  - TOPO-WMS
  url: http://ows.terrestris.de/osm/service
  version: ""' > ${HOME}/wms-config/.wms.yaml && chmod +x ${HOME}/wms-config/.wms.yaml
```

And Usage.

```console
wms cap terrestris

docker run -v $HOME/wms-config:/wms-config wroge/wms cap terrestris
```

Note: Before v0.0.5 Config-File: ```$HOME/.wms-cli.yaml```

## Help

GetMap.

```console
wms map --help
Download a WMS-Tile

Usage:
  wms map [flags]

Aliases:
  map, getmap

Flags:
  -b, --bbox strings       Set bbox in meters (minx,miny,maxx,maxy)
  -B, --bbox-file string   Set bbox file
  -C, --cut                Cuts image to unexpanded bbox (interesting for dynamically generated texts and symbols)
  -i, --dpi int            Set dpi of output image (scale required!)
      --dry-run            Validate your request without executing
  -e, --epsg int           Set epsg-code
  -E, --expand int         Expands bbox in %
  -n, --file-name string   Set file name
  -f, --format string      Set format
  -h, --height int         Set height of output image in px
  -l, --layers strings     Set layers
      --password string    Set password for Basic Authentication
      --save string        Save your request settings
  -s, --scale int          Set scale of output image (dpi required!)
  -u, --url string         Set url
      --user string        Set user for Basic Authentication
  -v, --version string     Set version
  -w, --width int          Set width of output image in px

Global Flags:
      --help   Help about any command
```

GetCapabilities.

```console
wms cap --help
Get the capabilities of a WMS

Usage:
  wms cap [flags]

Aliases:
  cap, getcap

Flags:
  -e, --epsg              Get available epsg-codes
  -f, --formats           Get available formats
  -l, --layers            Get available layers
      --password string   Set password
  -u, --url string        Set url
      --user string       Set user
  -v, --version string    Set version

Global Flags:
      --help   Help about any command
```

## Example

In this Example we use the OSM-WMS from [terrestris](https://ows.terrestris.de/dienste.html#openstreetmap-wms).

From terrestris provided EPSG-Codes: 900913, 4326, 3857

From wroge/wms provided EPSG-Codes: [go-coo](https://github.com/wroge/go-coo)

```console
docker run -v "$(pwd)/output:/output" wroge/wms map -u http://ows.terrestris.de/osm/service -e 12345
Error: Invalid EPSG: 12345
Valid EPSGs: [900913 4326 3857 31468 4462 32632 25832 25833 5668 31466 31467 4647 5650 3067 5669 31469 32633]
```

So we can use for example EPSG:25832 and the coordinates are automatically converted to a supported system. (here EPSG:3857)

```console
docker run -v "$(pwd)/output:/output" wroge/wms map -u http://ows.terrestris.de/osm/service -e 25832 -b 565000,5930000,570000,5935000 -w 1000 -n test
http://ows.terrestris.de/osm/service?SERVICE=WMS&REQUEST=GetMap&VERSION=1.1.1&FORMAT=image/jpeg&LAYERS=OSM-WMS&STYLES=&SRS=EPSG:3857&WIDTH=1000&HEIGHT=1000&BBOX=1110998.5364747,7078816.8197398,1119515.6176323,7087113.4792617
Done. Your requested file is here: /output
```
<img src="https://user-images.githubusercontent.com/44040384/54848416-7ffab480-4ce1-11e9-9fa9-b092a6e096ad.jpeg" width="50%">

Or with this [Shell-Script](#docker) and this [Configuration-File](#configuration).

<img src="https://github.com/wroge/wms/blob/master/demo.svg?sanitize=true">

```console
wms map terrestris -b 565000,5930000,570000,5935000 -w 1000
```

## FAQ

...For any problems/questions please open an issue.
