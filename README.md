# gamekit

A collection of utility packages for making games in Go.

- `sprite` package contains implemenation of sprite collection.
**Sprite** can consist of several frames and contain properties for animation.
**Anim** is another structure that tracks changing of sprite frames.
Collection is initialized from the yaml configuration file (see below)
This is pure logical package without dependency on any specific graphics library.
- `ebitenkit` package contains [ebiten](https://github.com/hajimehoshi/ebiten) 
specific collection of sprites that stores images and can draw them on an
ebiten image.

## Configuration file example

```yaml
files:
  - name: sprites/character.png  # sprite sheet image file name (relative to the executable)
    sprites:
      player-idle:     # sprite name
        width: 32      # width and height of the sprite
        height: 32
        xOffset: 0     # offset of the sprite in the image
        yOffset: 0
        xOrigin: 16    # origin of the sprite is at the
        yOrigin: 32    #   foot level
        frameCount: 2  # animation consists of 2 frames
        animFirst: 0   #   represented by a horizontal stripe 
        animLast: 1    #   of the same size frames
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
        animFirst: 0
        animLast: 7
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
