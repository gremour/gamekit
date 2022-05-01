// Package ebitenkit contains utilities specific to
// github.com/hajimehoshi/ebiten/v2 library
package ebitenkit

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gremour/gamekit/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Collection is a wrapper around sprite.Collection,
// keeping track of ebiten images for the corresponding sprites.
type Collection struct {
	*sprite.Collection
	frames          map[string][]frame
	logFunc         func(msg string)
	reportedSprites map[string]struct{}
}

// Intermediate structure to store frames and origins.
type frame struct {
	image   *ebiten.Image
	xOrigin float64
	yOrigin float64
}

// NewCollection creates new collection, initializing
// frame images.
func NewCollection(col *sprite.Collection, logFunc func(msg string)) (*Collection, error) {
	c := &Collection{
		Collection: col,
		frames:     make(map[string][]frame, len(col.Sprites)),
		logFunc:    logFunc,
	}
	for _, fi := range col.Config.Files {
		img, _, err := ebitenutil.NewImageFromFile(fi.Name)
		if err != nil {
			return nil, err
		}
		for name, spr := range fi.Sprites {
			frames := make([]frame, 0, spr.FrameCount)
			for i := 0; i < spr.FrameCount; i++ {
				frm := frame{
					image: img.SubImage(image.Rect(
						spr.XOffset+spr.Width*i,
						spr.YOffset,
						spr.XOffset+spr.Width*(i+1),
						spr.YOffset+spr.Height)).(*ebiten.Image),
					xOrigin: float64(spr.XOrigin),
					yOrigin: float64(spr.XOrigin),
				}
				frames = append(frames, frm)
			}
			c.frames[name] = frames
		}
	}

	return c, nil
}

// FrameImage returns frame image for the sprite name and frame number.
// May return nil if such sprite is not in collection or frame number
// is invalid.
// Also retuerns sprite origins.
func (c *Collection) FrameImage(name string, frame int) (*ebiten.Image, float64, float64) {
	fs := c.frames[name]
	if frame < 0 || frame >= len(fs) {
		return nil, 0, 0
	}
	f := fs[frame]
	return f.image, f.xOrigin, f.yOrigin
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

// Draw draws the sprite frame on dest image with given offset,
// scale and rotation.
func (c *Collection) Draw(dest *ebiten.Image, do *DrawOpts) {
	im, ox, oy := c.FrameImage(do.Name, do.Frame)
	if im == nil {
		if _, ok := c.reportedSprites[do.Name]; !ok {
			if c.logFunc != nil {
				c.logFunc(fmt.Sprintf("Sprite %v frame %v not found in collection", do.Name, do.Frame))
			}
			c.reportedSprites[do.Name] = struct{}{}
		}
	}
	var eo ebiten.DrawImageOptions
	if do.Rotation != 0 {
		eo.GeoM.Rotate(do.Rotation)
	}
	if do.ScaleX != 0 && do.ScaleY != 0 {
		eo.GeoM.Scale(do.ScaleX, do.ScaleY)
	}
	eo.GeoM.Translate(do.X-ox, do.Y-oy)
	if do.Color != nil {
		eo.ColorM.Apply(do.Color)
	}
	dest.DrawImage(im, &eo)
}
