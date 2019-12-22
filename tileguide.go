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
	s := tilePattern.FindStringSubmatch(r.URL.Path)
	if len(s) != 4 {
		fmt.Fprintf(w, "Failed to match tile expression!")
		return
	}
	//fmt.Fprintf(w, "%s %s %s %s\n", s[0], s[1], s[2], s[3])
	
	var x, y, z uint32
	var tile maptile.Tile
	
	fc := geojson.NewFeatureCollection()
	if z64, err := strconv.ParseUint(s[1], 10, 32); err == nil {
		if x64, err := strconv.ParseUint(s[2], 10, 32); err == nil {
			if y64, err := strconv.ParseUint(s[3], 10, 32); err == nil {
				//fmt.Fprintf(w, "%d %d %d\n", z, x, y)
				x, y, z = uint32(x64), uint32(y64), uint32(z64)
				tile = maptile.New(x, y, maptile.Zoom(z))
				center := tile.Center()
				//fmt.Fprintf(w, "%f %f\n", center[0], center[1])
				fc.Append(geojson.NewFeature(center))
			}
		}
	}

	//rawJSON, _ := fc.MarshalJSON()
	//fmt.Fprintf(w, "%s", rawJSON)
	collections := map[string]*geojson.FeatureCollection{
		"centers": fc,
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


