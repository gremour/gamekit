package main

import (
	"image/color"
	"log"
	"time"

	"github.com/gremour/gamekit/ebitenkit"
	"github.com/gremour/gamekit/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	c          *ebitenkit.Collection
	anims      []*sprite.Anim
	lastUpdate time.Time
}

func (g *game) Update() error {
	now := time.Now()
	if g.lastUpdate == (time.Time{}) {
		g.lastUpdate = now
	}
	dt := now.Sub(g.lastUpdate).Seconds()
	g.lastUpdate = now

	for _, a := range g.anims {
		a.Progress(dt)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{220, 220, 220, 255})

	// Draw second frame of the idle sprite
	g.c.Draw(screen, &sprite.DrawOpts{
		Name:  "idle",
		Frame: 1,
		X:     80,
		Y:     100,
	})

	// Draw animations
	for i, a := range g.anims {
		spr, frame := a.Current()
		g.c.Draw(screen, &sprite.DrawOpts{
			Name:  spr,
			Frame: frame,
			X:     120 + float64(i*40),
			Y:     100,
		})
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	h := 200
	w := h * outsideWidth / outsideHeight
	return w, h
}

func main() {
	c, err := sprite.NewCollectionFromFile("sprites/sprites.yaml")
	if err != nil {
		panic(err)
	}

	ec, err := ebitenkit.NewCollection(c, func(msg string) {
		log.Print(msg)
	})
	if err != nil {
		panic(err)
	}

	// Create 2 animations:
	// 1: jump with transition to idle
	anim1 := sprite.NewAnim(c, "jump", false)
	// 2: jump looped, reversed
	anim2 := sprite.NewAnim(c, "jump", true)
	anim2.SetLoop(true)
	g := &game{
		c:     ec,
		anims: []*sprite.Anim{anim1, anim2},
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
