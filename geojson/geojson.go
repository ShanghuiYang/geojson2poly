package geojson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Geojson ...
type Geojson struct {
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}

// Geometry ...
type Geometry struct {
	Coordinates [][][][]float64 `json:"coordinates"`
	Type        string          `json:"type"`
}

// Properties ...
type Properties struct {
	Name string `json:"name"`
}

// NewGeojson ...
func NewGeojson() *Geojson {
	return &Geojson{}
}

// Load ...
func (g *Geojson) Load(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, g); err != nil {
		return err
	}

	return nil
}

// ToPoly ...
func (g *Geojson) ToPoly(filePath string) error {
	s := g.Properties.Name
	if s == "" {
		s = "No Name"
	}
	s += "\n"

	for i, polygon := range g.Geometry.Coordinates {
		s += fmt.Sprintf("%v\n", i+1)
		for _, linears := range polygon {
			for _, points := range linears {
				s += fmt.Sprintf("\t%v %v\n", points[0], points[1])
			}
		}
		s += "END\n"
	}
	s += "END\n"

	if err := ioutil.WriteFile(filePath, []byte(s), 0644); err != nil {
		return err
	}

	return nil
}
