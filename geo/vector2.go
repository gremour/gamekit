package geo

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector2) Mul(o Vector2) Vector2 {
	return Vector2{
		X: v.X * o.X,
		Y: v.Y * o.X,
	}
}

func (v Vector2) Scale(factor float64) Vector2 {
	return Vector2{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

// Angle between vectors in radians [0 .. 2*pi)
func (v Vector2) Angle(o Vector2) float64 {
	a := math.Atan2(v.Y, v.X) - math.Atan2(o.Y, o.X)
	if a < 0 {
		a += math.Pi * 2
	} else if a >= math.Pi*2 {
		a -= math.Pi * 2
	}
	return a
}

func (v Vector2) Rotated(angle float64) Vector2 {
	sina := math.Sin(angle)
	cosa := math.Cos(angle)
	return Vector2{
		X: v.X*cosa - v.Y*sina,
		Y: v.X*sina + v.Y*cosa,
	}
}
