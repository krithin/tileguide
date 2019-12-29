package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	//"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/mvt"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
)

var /* const */ tilePattern = regexp.MustCompile("/([0-9]+)/([0-9]+)/([0-9]+).mvt")

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.mapbox-vector-tile")

	s := tilePattern.FindStringSubmatch(r.URL.Path)
	if len(s) != 4 {
		fmt.Fprintf(w, "Failed to match tile expression!")
		return
	}
	
	var x, y, z uint32
	var tile maptile.Tile
	
	centers := geojson.NewFeatureCollection()
	borders := geojson.NewFeatureCollection()
	if z64, err := strconv.ParseUint(s[1], 10, 32); err == nil {
		if x64, err := strconv.ParseUint(s[2], 10, 32); err == nil {
			if y64, err := strconv.ParseUint(s[3], 10, 32); err == nil {
				x, y, z = uint32(x64), uint32(y64), uint32(z64)
				tile = maptile.New(x, y, maptile.Zoom(z))
				center := geojson.NewFeature(tile.Center())
				center.Properties = geojson.Properties{
					"x": x, "y": y, "z": z,
				}
				centers.Append(center)
				border := geojson.NewFeature(tile.Bound().ToPolygon())
				borders.Append(border)
			}
		}
	}

	collections := map[string]*geojson.FeatureCollection{
		"centers": centers,
		"borders": borders,
	}
	layers := mvt.NewLayers(collections)
	layers.ProjectToTile(tile)
	data, err := mvt.Marshal(layers)
	if err == nil {
		fmt.Fprintf(w, "%s", data)
	}
}

func main() {
	http.HandleFunc("/tileguide/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


