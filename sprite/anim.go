package sprite

// Anim controls sprite animation.
// Call Progress providing time in seconds passed since last frame.
// Call Current to obtain sprite name and current frame number.
type Anim struct {
	collection        *Collection
	sprite            string
	frameCount        int
	frameCurrent      int
	loop              bool
	next              string
	uniformDuration   float64
	remainingDuration float64
	reverse           bool
}

// NewAnim creates new animation.
func NewAnim(col *Collection, name string, reverse bool) *Anim {
	if col == nil {
		return nil
	}
	a := &Anim{
		collection: col,
		reverse:    reverse,
	}
	ok := a.FromSprite(name)
	if !ok {
		return nil
	}
	return a
}

// fromSprite sets up current animation based on sprite name from collection.
func (a *Anim) FromSprite(name string) bool {
	spr, ok := a.collection.Sprites[name]
	if !ok {
		return false
	}
	a.sprite = spr.Name
	a.frameCount = spr.FrameCount
	a.frameCurrent = 0
	a.next = spr.AnimNext
	a.loop = spr.AnimLoop
	a.uniformDuration = float64(spr.FrameMS) * 1e-3
	a.remainingDuration = a.uniformDuration
	return true
}

// Progress advances animation by given time in seconds.
func (a *Anim) Progress(dt float64) {
	a.remainingDuration -= dt
	for a.remainingDuration < 0 {
		a.remainingDuration += a.uniformDuration
		a.frameCurrent++
		if a.frameCurrent >= a.frameCount {
			if a.loop {
				a.frameCurrent = 0
			} else {
				if a.next != "" {
					a.FromSprite(a.next)
				}
				return
			}
		}
	}
}

// SetDirection sets animation direction.
// If reverse is true, animation will play from end to start frames.
// If reset is true, animation will start from the beginning.
func (a *Anim) SetDirection(reverse, reset bool) {
	a.reverse = reverse
	if reset {
		a.frameCurrent = 0
		a.remainingDuration = a.uniformDuration
	}
}

// Current returns current sprite name and frame.
func (a *Anim) Current() (name string, frame int) {
	f := a.frameCurrent
	if a.reverse {
		f = a.frameCount - 1 - a.frameCurrent
	}
	return a.sprite, f
}

func (a *Anim) SetLoop(loop bool) {
	a.loop = loop
}
