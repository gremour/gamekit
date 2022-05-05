# gamekit

A collection of utility packages for making games in the awesome Go programming language.

I've created this library to help develop my pet game using [ebiten](https://github.com/hajimehoshi/ebiten)
graphics library. And I'm sharing this in hope it can be useful.

- `sprite` package contains implementation of sprite collection.
**Sprite** can consist of several frames and contain properties for animation.
**Anim** is another structure that tracks changing of sprite frames.
Collection is initialized from the yaml configuration file (see below).
This is pure logical package without dependency on any specific graphics library;
- `ebitenkit` package contains [ebiten](https://github.com/hajimehoshi/ebiten) 
specific collection of sprites that stores images and can draw them on an
ebiten image;
- `geo` package contains structures and functions for `float64` geometry calculations;
- `sheetmaker` package implements a tool that produces sprite sheet (both `png` file
and `yaml` config file) from a number of `png` files; configuration file is used to
describe files and desired sheet properties. Run `go run ./sheetmaker` to get help.

## Configuration file example

This file can be created automatically (along with sprite sheets) by `sheetmaker` by combining multiple
`png` files containing sprites, which are the horizontal stripes of sprite frames,
as exported from the sprite editor, such as **Aseprite**.

Run `go run ./sheetmaker` for the usage help. Run `go install ./sheetmaker` to
install sheetmaker binary to your `$GOBIN` path.

```yaml
files:
  - name: sprites/character.png  # sprite sheet image file name (relative to the executable)
    sprites:
      player-idle:     # sprite name
        width: 32      # width and height of the sprite
        height: 32
        xOffset: 0     # offset of the sprite in the image
        yOffset: 0
        xOrigin: 16    # origin of the sprite is at the foot level
        yOrigin: 32
        frameCount: 2  # animation consists of 2 frames represented by a horizontal stripe 
        animLoop: true # animation loops
        frameMS: 500   # uniform number of milliseconds for each frame
      player-walk:
        width: 32
        height: 32
        xOffset: 0
        yOffset: 32    # note the offset: this sprite frames are on the second row in the image
        xOrigin: 16
        yOrigin: 32
        frameCount: 8
        animNext: player-idle # after animation ends, switch to the different animation
        frameMS: 100
  - name: sprites/terrain.png
    sprites:
      grass:
        width: 32
        height: 32
        xoffset: 0
        yoffset: 0
        frameCount: 8 # grass variations instead of animation frames
```

# Contribution

Ideas, contributions and criticisms are welcome.

Please open an issue for an idea/criticism or a pull request for a contribution.
