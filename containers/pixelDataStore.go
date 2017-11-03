package containers

import (
	"errors"
	"math/rand"
	"time"
)

var RandomGen *rand.Rand


//Package Init//
func init(){
	RandomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

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
	GetPixelsBelow() []Pixel
	GetPixelsRight() []Pixel
	GetRandomPixelBelow() Pixel
	GetRandomPixelRight() Pixel
	
	GetColor() Pixel
}

type SimplePixelData struct {
	PixelBelow	[]Pixel
	PixelRight	[]Pixel
	
	Color		Pixel
}

type DataManager interface{
	GetPixelData(Pixel) (PixelData, error)
	GetPixelDataCreateIfNotExist(Pixel) PixelData
}

type SimpleDataManager struct {
	Data		map[Pixel]*SimplePixelData
}

/////////////
//Functions//
/////////////

func MakeSimpleDataManager() *SimpleDataManager{
	sdm := new(SimpleDataManager)
	sdm.Data = make(map[Pixel]*SimplePixelData)
	
	return sdm
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
	p.PixelBelow = append(p.PixelBelow, pixel)
}

func (p *SimplePixelData) AddPixelRight(pixel Pixel){
	p.PixelRight = append(p.PixelRight, pixel)
}

func (p *SimplePixelData) GetColor() Pixel {
	return p.Color
}

func (p *SimplePixelData) GetPixelsBelow() []Pixel{
	return p.PixelBelow
}

func (p *SimplePixelData) GetPixelsRight() []Pixel{
	return p.PixelRight
}

func (p *SimplePixelData) GetRandomPixelBelow() Pixel{
	return p.PixelBelow[RandomGen.Intn(len(p.PixelBelow))]
}

func (p *SimplePixelData) GetRandomPixelRight() Pixel{
	return p.PixelRight[RandomGen.Intn(len(p.PixelRight))]
}

//SimpleDataManager
func (dm *SimpleDataManager) GetPixelData(pixel Pixel) (PixelData,error) {
	pixelData, exists := dm.Data[pixel]
	
	if(!exists){
		return nil,errors.New("Pixel has not been stored")
	}
	
	return pixelData,nil
}

func (dm *SimpleDataManager) GetPixelDataCreateIfNotExist(pixel Pixel) PixelData{
	pixelData, exists := dm.Data[pixel]
	
	if(!exists){
		dm.Data[pixel] = new(SimplePixelData)
		pixelData, _ = dm.Data[pixel]
	}
	
	
	return pixelData
}