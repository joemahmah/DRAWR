package containers

import (
	"errors"
)

type Pixel struct {
	R		byte
	G		byte
	B		byte
	A		byte
	
	Edge	byte //0 = no, 1 = left, 2 = top
}

type PixelData interface{
	AddPixelBelow(Pixel)
	AddPixelRight(Pixel)
	
	GetColor() Pixel
}

type SimplePixelData struct {
	pixelBelow	[]Pixel
	pixelRight	[]Pixel
	
	color		Pixel
}

type DataManager interface{
	GetPixelData(Pixel) (PixelData, error)
	SetPixelData(Pixel,PixelData) error
}

type SimpleDataManager struct {
	data		map[Pixel]SimplePixelData
}


///////////
//Methods//
///////////

//Pixel
func (p *Pixel) SetRGB(r,g,b byte){
	p.R = r
	p.G = g
	p.B = b
}

func (p *Pixel) SetRGBA(r,g,b,a byte){
	p.R = r
	p.G = g
	p.B = b
	p.A = a
}

func (p *Pixel) SetAlpha(a byte){
	p.A = a
}

func (p *Pixel) SetEdge(e byte){
	p.Edge = e
}

func (p *Pixel) getRGB() (byte,byte,byte){
	return p.R, p.G, p.B
}

func (p *Pixel) getRGBA() (byte,byte,byte,byte){
	return p.R, p.G, p.B, p.A
}

func (p *Pixel) IsEdge() bool{
	return p.Edge != 0
}

//SimplePixelData
func (p *SimplePixelData) AddPixelBelow(pixel Pixel){
	p.pixelBelow = append(p.pixelBelow, pixel)
}

func (p *SimplePixelData) AddPixelRight(pixel Pixel){
	p.pixelRight = append(p.pixelRight, pixel)
}

func (p *SimplePixelData) GetColor() Pixel {
	return p.color
}


//SimpleDataManager
func (dm *SimpleDataManager) GetPixelData(pixel Pixel) (PixelData,error) {
	pixelData, exists := dm.data[pixel]
	
	if(!exists){
		return nil,errors.New("Pixel has not been stored")
	}
	
	return &pixelData,nil
}

func (dm *SimpleDataManager) SetPixelData(pixel Pixel, pixelData PixelData) error {
	return nil
}