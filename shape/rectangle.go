package shape

type Rectangle struct {
	Width  float32
	Length float32
}

//Area()=>nama func
func (rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Kell() float32 {
	return 2 * (rectangle.Width + rectangle.Length)
}
