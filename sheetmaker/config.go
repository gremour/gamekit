package main

import (
	"fmt"
	"strings"

	"github.com/gremour/gamekit/sprite"
)

type Config struct {
	Name   string         `yaml:"name"`
	Sheets []*Sheet       `yaml:"sheets"`
	Embeds []*sprite.File `yaml:"embeds"`
}

type Sheet struct {
	Name           string  `yaml:"name"`
	Width          int     `yaml:"width"`
	UniformDelayMS int     `yaml:"uniformDelayMS"`
	Files          []*File `yaml:"files"`
}

type File struct {
	Name       string `yaml:"name"`
	SpriteName string `yaml:"spriteName"`
	XOrigin    int    `yaml:"xOrigin"`
	YOrigin    int    `yaml:"yOrigin"`
	FrameCount int    `yaml:"frameCount"`
	FrameMS    int    `yaml:"frameMS"`
	AnimLoop   bool   `yaml:"animLoop"`
	AnimNext   string `yaml:"animNext"`
}

// Tidy sets default values for the config and validates required fields.
func (c *Config) Tidy() error {
	if c.Name == "" {
		c.Name = "sprites.yaml"
	}
	for i, s := range c.Sheets {
		if s.Name == "" {
			return fmt.Errorf("Sheet number %v image file name is required", i)
		}
		if s.Width <= 0 {
			s.Width = 1024
		}
		if s.UniformDelayMS <= 0 {
			s.UniformDelayMS = 100
		}
		for i, f := range s.Files {
			if s.Name == "" {
				return fmt.Errorf("Sheet %v file number %v file name is required", s.Name, i)
			}
			if f.SpriteName == "" {
				parts := strings.Split(f.Name, `/`)
				parts = strings.Split(parts[len(parts)-1], `\`)
				f.SpriteName = strings.TrimSuffix(parts[len(parts)-1], ".png")
			}
			if f.FrameMS == 0 {
				f.FrameMS = s.UniformDelayMS
			}
		}
	}
	return nil
}
