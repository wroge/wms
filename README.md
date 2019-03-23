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

Then you can use ```wms upgrade```to get the newest version.

## Configuration

```$HOME/wms-config/.wms.yaml```

Example.

```console
wms map -u http://ows.terrestris.de/osm/service -n example -e 25832 -l TOPO-WMS --save terrestris --dry-run
```

And Usage.

```console
wms cap terrestris
wms map terrestris -b 565000,5930000,570000,5935000 -w 1000
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

## FAQ

...For any problems/questions please open an issue.
