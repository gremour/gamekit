package main

import (
	"image/color"
	"log"

	"github.com/gremour/gamekit/ebitenkit"
	"github.com/gremour/gamekit/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	c *ebitenkit.Collection
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{220, 220, 220, 255})
	g.c.Draw(screen, &sprite.DrawOpts{
		Name:  "idle",
		Frame: 0,
		X:     120,
		Y:     100,
	})
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

	g := &game{
		c: ec,
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
