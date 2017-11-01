package containers

type Pixel struct {
	R	int
	G	int
	B	int
	A	int
}

type PixelData struct {
	pixelBelow	[]Pixel
	pixelRight	[]Pixel
	
	color		Pixel
}

type DataManager struct {
	data		map[Pixel]PixelData
}

func (p *PixelData) AddPixelBelow(pixel Pixel){
	p.pixelBelow = append(p.pixelBelow, pixel)
}

func (p *PixelData) AddPixelRight(pixel Pixel){
	p.pixelBelow = append(p.pixelBelow, pixel)
}