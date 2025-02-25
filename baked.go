package texukk

//go:generate texukk --source sprites

type coordTypes interface {
	float32 | float64 | uint8 | uint16 | uint32 | uint64
}

type AtlasCoords[T coordTypes] struct {
	X T
	Y T
}

type AtlasTexture[T coordTypes] struct {
	Min AtlasCoords[T]
	Max AtlasCoords[T]
}

func (a *AtlasTexture[T]) Width() T {
	return a.Max.X - a.Min.X
}

func (a *AtlasTexture[T]) Height() T {
	return a.Max.Y - a.Min.Y
}

func TexCoords[T coordTypes](x1, y1, x2, y2 T) AtlasTexture[T] {
	return AtlasTexture[T]{
		Min: AtlasCoords[T]{X: x1, Y: y1},
		Max: AtlasCoords[T]{X: x2, Y: y2},
	}
}
