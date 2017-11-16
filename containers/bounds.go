package containers

func GetInvalidPixel() Pixel {
	return Pixel{255,255,255,255,255}
}

func GetUpperBound() Pixel {
	return Pixel{0,0,0,0,128}
}

func GetLeftBound() Pixel {
	return Pixel{0,0,0,0,1}
}

func GetBound(x, y byte) Pixel {
	//x is lower 4 bits while y is upper 4 bits
	return Pixel{0,0,0,0,(x % 16) + (y % 16) * 16}
}

func GetBoundCoord(pixel Pixel) (byte, byte) {
	x := (pixel.Edge << 4) >> 4
	y := pixel.Edge >> 4
	return x, y
}