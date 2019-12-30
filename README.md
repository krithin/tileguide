# tileguide
Visualize slippy tile bounds over a mapbox map.

This supports XYZ tiles in the format used by [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames) and [Google Maps](https://developers.google.com/maps/documentation/javascript/coordinates#tile-coordinates), which is notably different from the [TMS scheme](https://gist.github.com/tmcw/4954720) used by some other services.

See it in action: https://onetwotwo.sg/tileguide

## Motivation
I often work with tiled geospatial datasets (whether that's raster satellite imagery or vector tiles containing information about a road network). Getting tile references right when working with external datasets is occasionally challenging, so I created this tool to help visualize standard [OSM](https://wiki.openstreetmap.org/wiki/Slippy_Map) / [Mapbox](https://docs.mapbox.com/vector-tiles/reference/)-compatible [slippy tile](https://en.wikipedia.org/wiki/Tiled_web_map) numbers to make it easier to compare them to the tile identifiers used in an dataset.

## Usage

You can use an existing, deployed instance of this tool at https://onetwotwo.sg/tileguide.

Alternatively, deploy it for yourself with Docker.

### Deploy with docker

1. Set up the tileguide server:

    docker build -t tileguide .
    docker run --rm -p 8080:8080 tileguide:latest

1. Move `index.html` somewhere where your HTTP server can read from it.

1. Modify `index.html`:
	1. Set `mapboxgl.accessToken` to your [Mapbox access token](https://docs.mapbox.com/help/how-mapbox-works/access-tokens/)
	1. Set `tileguide_server_url` to point to your tileguide server from step 1.

That's it!

### Manual deploy

As an alternative to the docker build you can also clone this repository into your `$GOPATH` and use standard golang tooling (`go get && go build`) to build the server.
