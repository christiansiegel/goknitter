package generator

type Matrix struct {
	Vals   []uint8
	Stride int
}

func NewMatrix(width, height int) *Matrix {
	vals := make([]uint8, width*height)
	return &Matrix{vals, width}
}

func (m *Matrix) Width() int {
	return m.Stride
}

func (m *Matrix) Height() int {
	return len(m.Vals) / m.Stride
}

func (m *Matrix) ValOffset(x, y int) int {
	if x < 0 || y < 0 || x >= m.Width() || y >= m.Height() {
		return -1
	}
	return x + m.Stride*y
}

func (m *Matrix) Set(x, y int, v uint8) {
	offset := m.ValOffset(x, y)
	if offset < 0 {
		return
	}
	m.Vals[offset] = v
}

func (m *Matrix) At(x, y int) uint8 {
	offset := m.ValOffset(x, y)
	if offset < 0 {
		return 0
	}
	return m.Vals[offset]
}
