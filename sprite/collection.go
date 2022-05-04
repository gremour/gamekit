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
	Files []*File `yaml:"files,omitempty"`
}

// File represents one sprite sheet file.
type File struct {
	Name    string
	Sprites map[string]*Sprite `yaml:"sprites,omitempty"`
}

// Sprite contains information about position and size of
// sprite in the sheet, it's origin point and animation properties.
// Sprite can contain several frames.
type Sprite struct {
	Name       string `yaml:"-"`
	XOffset    int    `yaml:"xOffset,omitempty"`
	YOffset    int    `yaml:"yOffset,omitempty"`
	XOrigin    int    `yaml:"xOrigin,omitempty"`
	YOrigin    int    `yaml:"yOrigin,omitempty"`
	Width      int    `yaml:"width,omitempty"`
	Height     int    `yaml:"height,omitempty"`
	FrameCount int    `yaml:"frameCount,omitempty"`
	AnimLoop   bool   `yaml:"animLoop,omitempty"`
	AnimNext   string `yaml:"animNext,omitempty"`
	FrameMS    int    `yaml:"frameMS,omitempty"`
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

	c := &Collection{
		Sprites: make(map[string]*Sprite),
	}
	err = yaml.Unmarshal(payload, &c.Config)
	if err != nil {
		return nil, err
	}

	for _, fi := range c.Config.Files {
		for name, spr := range fi.Sprites {
			spr.Name = name
			if spr.FrameCount == 0 {
				spr.FrameCount = 1
			}
			c.Sprites[name] = spr
		}
	}

	return c, nil
}
