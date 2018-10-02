package offheap

type Small struct {
	s int8
}

func NewSmall() *Small {
	return &Small{s: int8(1)}
}

type Large struct {
	a [10000000]byte
}

func NewLarge() *Large {
	return &Large{}
}
