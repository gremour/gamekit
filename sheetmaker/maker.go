package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/gremour/gamekit/sprite"
	"gopkg.in/yaml.v2"
)

type Maker struct {
	Config Config

	sheetConfig sprite.Config
}

func (m *Maker) Run() error {
	err := m.Config.Tidy()
	if err != nil {
		return err
	}

	for _, s := range m.Config.Sheets {
		f, img, err := m.createSheet(s)
		if err != nil {
			return fmt.Errorf("failed to create sheet %v: %w", s.Name, err)
		}
		fim, err := os.Create(s.Name)
		if err != nil {
			return fmt.Errorf("failed to create sheet file: %w", err)
		}
		err = png.Encode(fim, img)
		if err != nil {
			return fmt.Errorf("failed to encode sheet file: %w", err)
		}
		err = fim.Close()
		if err != nil {
			return fmt.Errorf("failed to write to sheet file: %w", err)
		}

		m.sheetConfig.Files = append(m.sheetConfig.Files, f)
	}
	m.sheetConfig.Files = append(m.sheetConfig.Files, m.Config.Embeds...)

	payload, err := yaml.Marshal(m.sheetConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal sheet config file: %w", err)
	}

	err = ioutil.WriteFile(m.Config.Name, payload, 0644)
	if err != nil {
		return fmt.Errorf("failed to write sheet config file: %w", err)
	}

	return nil
}

// Create sprite sheet image, save it to a file and return config entry for the created file.
func (m *Maker) createSheet(s *Sheet) (*sprite.File, image.Image, error) {
	var x, y, maxh int
	sheetImg := image.NewRGBA(image.Rect(0, 0, s.Width, s.Width*2))
	fc := &sprite.File{
		Name:    s.Name,
		Sprites: make(map[string]*sprite.Sprite),
	}
	for _, f := range s.Files {
		file, err := os.Open(f.Name)
		if err != nil {
			return nil, nil, err
		}
		img, _, err := image.Decode(file)
		if err != nil {
			return nil, nil, err
		}
		min := img.Bounds().Min
		max := img.Bounds().Max
		w := max.X - min.X
		h := max.Y - min.Y
		if x+w > s.Width {
			y += maxh
			maxh = 0
			x = 0
		}
		if h > maxh {
			maxh = h
		}
		frames := w / h
		if f.FrameCount != 0 {
			frames = f.FrameCount
		}
		sw := w / frames
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				sheetImg.Set(x+j, y+i, img.At(j, i))
			}
		}
		fms := f.FrameMS
		if fms == 0 {
			fms = s.UniformDelayMS
		}
		spr := &sprite.Sprite{
			Name:       f.SpriteName,
			XOffset:    x,
			YOffset:    y,
			XOrigin:    f.XOrigin,
			YOrigin:    f.YOrigin,
			Width:      sw,
			Height:     h,
			FrameCount: frames,
			AnimLoop:   f.AnimLoop,
			AnimNext:   f.AnimNext,
			FrameMS:    fms,
		}
		fc.Sprites[spr.Name] = spr
		x += w
	}
	img := sheetImg.SubImage(image.Rect(0, 0, s.Width, y+maxh))
	fmt.Printf("Prepared sheet image %v (%vx%v)\n", s.Name, s.Width, y+maxh)

	return fc, img, nil
}
