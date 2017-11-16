package parsers

import (
	"image"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"github.com/joemahmah/DRAWR/containers"
	//"fmt"
)

type Parser interface{
	LoadImage(string) error
	Parse() error
}

type SimpleParser struct {
	Img			image.Image
	Storage		containers.SimpleDataManager
}

type SimpleGridParser struct {
	Img			image.Image
	Storage		containers.SimpleGridDataManager
}

/////////////
//Functions//
/////////////

func convert32BitRGBAtoPixel(r,g,b,a uint32) containers.Pixel{
	return containers.Pixel{byte(r/257),byte(g/257),byte(b/257),byte(a/257),0}
}


///////////
//Methods//
///////////

//SimpleParser
func (parser *SimpleParser) LoadImage(path string) error {
	//Open the file specified by path
	file, err := os.Open(path)
	defer file.Close()
	
	//If error while reading, return error
	if(err != nil){
		return err
	}
	
	//Load the image
	parser.Img, _, err = image.Decode(file)
	
	//If error while loading, return error
	if(err != nil){
		return err
	}
	
	return nil
}

func (parser *SimpleParser)	SetStorage(container containers.SimpleDataManager) {
	parser.Storage = container
}

func (parser *SimpleParser) Parse() error{

	imageBounds := parser.Img.Bounds()
	boundsX, boundsY := imageBounds.Max.X, imageBounds.Max.Y

	var currentPixel containers.Pixel
	_ = currentPixel //"currentPixel declared and not used" even with loop below...	
	
	var leftPixel containers.Pixel
	var upPixel containers.Pixel
	
	_ = leftPixel
	_ = upPixel
	
    for y := 0; y < boundsY; y++ {
        for x := 0; x < boundsX; x++ {
			currentPixel = convert32BitRGBAtoPixel(parser.Img.At(x,y).RGBA()) //Current pixel at (x,y)
			
			//Set left pixel
			if(x == 0){
				leftPixel = containers.GetLeftBound()
			} else {
				leftPixel = convert32BitRGBAtoPixel(parser.Img.At(x-1,y).RGBA())
			}
			
			//Set up pixel
			if(y == 0){
				upPixel = containers.GetUpperBound()
			} else {
				upPixel = convert32BitRGBAtoPixel(parser.Img.At(x,y-1).RGBA())
			}
			
			//Update data for left pixel
			parser.Storage.GetPixelDataCreateIfNotExist(leftPixel).AddPixelRight(currentPixel)
			//Update data for up pixel
			parser.Storage.GetPixelDataCreateIfNotExist(upPixel).AddPixelBelow(currentPixel)
        }
    }
	
	return nil
}

//SimpleGridParser
func (parser *SimpleGridParser) LoadImage(path string) error {
	//Open the file specified by path
	file, err := os.Open(path)
	defer file.Close()
	
	//If error while reading, return error
	if(err != nil){
		return err
	}
	
	//Load the image
	parser.Img, _, err = image.Decode(file)
	
	//If error while loading, return error
	if(err != nil){
		return err
	}
	
	return nil
}

func (parser *SimpleGridParser)	SetStorage(container containers.SimpleGridDataManager) {
	parser.Storage = container
}

func (parser *SimpleGridParser) Parse() error{

	imageBounds := parser.Img.Bounds()
	boundsX, boundsY := imageBounds.Max.X, imageBounds.Max.Y

	var currentPixel containers.Pixel
	_ = currentPixel //"currentPixel declared and not used" even with loop below...	
	
	var leftPixel containers.Pixel
	var upPixel containers.Pixel
	var diagPixel containers.Pixel
	
	_ = leftPixel
	_ = upPixel
	_ = diagPixel
	
    for y := 0; y < boundsY; y++ {
        for x := 0; x < boundsX; x++ {
			currentPixel = convert32BitRGBAtoPixel(parser.Img.At(x,y).RGBA()) //Current pixel at (x,y)
			
			//Set left pixel
			if(x == 0){
				leftPixel = containers.GetLeftBound()
			} else {
				leftPixel = convert32BitRGBAtoPixel(parser.Img.At(x-1,y).RGBA())
			}
			
			//Set up pixel
			if(y == 0){
				upPixel = containers.GetUpperBound()
			} else {
				upPixel = convert32BitRGBAtoPixel(parser.Img.At(x,y-1).RGBA())
			}
			
			//Set diag pixel
			if(x == 0 && y == 0){
				diagPixel = containers.GetBound(1,1)
			} else {
				diagPixel = convert32BitRGBAtoPixel(parser.Img.At(x-1,y-1).RGBA())
			}
			
			//Update data for left pixel
			parser.Storage.GetPixelDataCreateIfNotExist(leftPixel).AddPixelRight(currentPixel) //AddPixelAt(1,0,currentPixel)
			//Update data for up pixel
			parser.Storage.GetPixelDataCreateIfNotExist(upPixel).AddPixelBelow(currentPixel) //AddPixelAt(0,1,currentPixel)
			//Update data for diag pixel
			parser.Storage.GetPixelDataCreateIfNotExist(diagPixel).AddPixelAt(1, 1, currentPixel)
        }
    }
	
	return nil
}