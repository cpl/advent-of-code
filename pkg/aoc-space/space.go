package aoc_space

import "math"

type Space struct {
	Size    Vec
	SizeNeg Vec
}

func (sp Space) AbsSize() Vec {
	sizex := sp.Size.X + sp.SizeNeg.X
	sizey := sp.Size.Y + sp.SizeNeg.Y
	sizez := sp.Size.Z + sp.SizeNeg.Z

	return Vec{
		X: sizex,
		Y: sizey,
		Z: sizez,
	}
}

type Vec struct {
	X, Y, Z int64
}

func (v Vec) Add(o Vec) Vec {
	return Vec{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vec) Adds(others ...Vec) []Vec {
	out := make([]Vec, len(others))
	for idx, other := range others {
		out[idx] = v.Add(other)
	}

	return out
}

func (v Vec) AddInt(x, y, z int) Vec {
	return Vec{
		X: v.X + int64(x),
		Y: v.Y + int64(y),
		Z: v.Z + int64(z),
	}
}

func (v Vec) Distance(other Vec) float64 {
	dx := other.X - v.X
	dy := other.Y - v.Y
	dz := other.Z - v.Z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func (v Vec) DistanceDelta(other Vec) int64 {
	dx := other.X - v.X
	dy := other.Y - v.Y
	dz := other.Z - v.Z

	return dx*dx + dy*dy + dz*dz
}

func Vec2(x, y int) Vec {
	return Vec{X: int64(x), Y: int64(y), Z: 0}
}

func Vec3(x, y, z int) Vec {
	return Vec{X: int64(x), Y: int64(y), Z: int64(z)}
}

type Point[T any] struct {
	Value    T
	Position Vec
}
