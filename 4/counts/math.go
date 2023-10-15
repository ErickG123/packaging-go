package counts

type math struct {
	a int
	b int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

// Consigo acessar essa função de outro arquivo por
// causa da letra maiuscula
func (m math) Add() int {
	return m.a + m.b
}

// Não consigo acessar essa função de outro arquivo
// por causa da letra minuscula
func (m math) sub() int {
	return m.a - m.b
}
