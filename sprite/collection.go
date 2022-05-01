// Package sprite implements sprite collection loading from config file.
// Sprite can contain several frames and animation properties.
package sprite

import (
	"image/color"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Collection contains information about all sprites & animation
// properties, as well as a map to lookup sprites by their names.
type Collection struct {
	Sprites map[string]*Sprite
	Config  Config
}

// Config is the structure to store unmarshalled config.
type Config struct {
	Files []*File `yaml:"files"`
}

// File represents one sprite sheet file.
type File struct {
	Name    string
	Sprites map[string]*Sprite `yaml:"sprites"`
}

// Sprite contains information about position and size of
// sprite in the sheet, it's origin point and animation properties.
// Sprite can contain several frames.
type Sprite struct {
	Name       string
	XOffset    int    `yaml:"xOffset"`
	YOffset    int    `yaml:"yOffset"`
	XOrigin    int    `yaml:"xOrigin"`
	YOrigin    int    `yaml:"yOrigin"`
	Width      int    `yaml:"width"`
	Height     int    `yaml:"height"`
	FrameCount int    `yaml:"frameCount"`
	AnimFirst  int    `yaml:"animFirst"`
	AnimLast   int    `yaml:"animLast"`
	AnimLoop   bool   `yaml:"animLoop"`
	AnimNext   string `yaml:"animNext"`
	FrameMS    int    `yaml:"frameMS"`
}

// DrawOpts contains options for drawing sprite.
// Name is mandatory. Other fields can be left empty.
type DrawOpts struct {
	Name   string
	Frame  int
	X      float64
	Y      float64
	ScaleX float64
	ScaleY float64
	// Rotation angle in radians.
	Rotation float64
	Color    color.Color
}

// NewCollectionFromFile creates sprite collection from yaml configuration file.
func NewCollectionFromFile(fileName string) (*Collection, error) {
	payload, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	conf := Config{}
	err = yaml.Unmarshal(payload, &conf)
	if err != nil {
		return nil, err
	}

	c := &Collection{
		Sprites: make(map[string]*Sprite),
	}

	for _, fi := range conf.Files {
		for name, spr := range fi.Sprites {
			if spr.FrameCount == 0 {
				spr.FrameCount = 1
			}
			c.Sprites[name] = spr
		}
	}

	return c, nil
}
