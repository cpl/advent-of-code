package aoc_space

type Grid2[T any] struct {
	space *Space
	data  []T
}

func NewGrid[T any](space *Space) *Grid2[T] {
	return &Grid2[T]{
		space: space,
		data:  make([]T, space.Size.X*space.Size.Y),
	}
}

func (grid *Grid2[T]) Space() Space {
	return *grid.space
}

func (grid *Grid2[T]) Contains(vec Vec) bool {
	space := grid.space

	x, y := vec.X, vec.Y
	sxmin, sxmax := -space.SizeNeg.X, space.Size.X
	symin, symax := -space.SizeNeg.Y, space.Size.Y

	if x < sxmin || x >= sxmax || y < symin || y >= symax {
		return false
	}

	return true
}

func (grid *Grid2[T]) absVec(vec Vec) Vec {
	sizeNeg := grid.space.SizeNeg

	return Vec{
		X: vec.X + sizeNeg.X,
		Y: vec.Y + sizeNeg.Y,
	}
}

func (grid *Grid2[T]) VecToIndex(vec Vec) int {
	if !grid.Contains(vec) {
		return -1
	}

	vec = grid.absVec(vec)
	size := grid.space.AbsSize()

	return int(vec.X + size.X*vec.Y)
}

func (grid *Grid2[T]) Set(value T, position Vec) bool {
	idx := grid.VecToIndex(position)
	if idx < 0 {
		return false
	}

	grid.data[idx] = value
	return true
}

func (grid *Grid2[T]) getFast(x, y, sizeX int64) T {
	return grid.data[x+sizeX*y]
}

func (grid *Grid2[T]) Get(position Vec) (T, bool) {
	idx := grid.VecToIndex(position)
	if idx < 0 {
		var zero T
		return zero, false
	}

	return grid.data[idx], true
}

var Grid2Neighbours = []Vec{
	Vec2(1, 0),
	Vec2(1, -1),
	Vec2(0, -1),
	Vec2(-1, -1),
	Vec2(-1, 0),
	Vec2(-1, 1),
	Vec2(0, 1),
	Vec2(1, 1),
}

func (grid *Grid2[T]) GetNeighbours(position Vec) []Point[T] {
	out := make([]Point[T], 0, 8)

	neighbourPositions := position.Adds(Grid2Neighbours...)
	for _, pos := range neighbourPositions {
		value, found := grid.Get(pos)
		if !found {
			continue
		}

		out = append(out, Point[T]{
			Value:    value,
			Position: pos,
		})
	}

	return out
}

func (grid *Grid2[T]) Iterate(f func(point Point[T])) {
	space := grid.space
	size := space.AbsSize()

	for y := -space.SizeNeg.Y; y < space.Size.Y; y++ {
		for x := -space.SizeNeg.X; x < space.Size.X; x++ {
			value := grid.getFast(x, y, size.X)

			f(Point[T]{
				Value:    value,
				Position: Vec{X: x, Y: y},
			})
		}
	}
}
