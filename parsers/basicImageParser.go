package parsers

import (
	"image"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"github.com/joemahmah/DRAWR/containers"
)

type Parser interface{
	LoadImage(string) error
	SetStorage(containers.DataManager)
	Parse() error
	
	GetUpperBound() containers.Pixel
	GetLeftBound() containers.Pixel
}

type SimpleParser struct {
	img			image.Image
	storage		containers.SimpleDataManager
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


func (parser *SimpleParser) LoadImage(path string) error {
	//Open the file specified by path
	file, err := os.Open(path)
	
	//If error while reading, return error
	if(err != nil){
		return err
	}
	
	//Load the image
	parser.img, _, err = image.Decode(file)
	
	//If error while loading, return error
	if(err != nil){
		return err
	}
	
	return nil
}

func (parser *SimpleParser)	SetStorage(container containers.SimpleDataManager) {
	parser.storage = container
}

func (parser *SimpleParser) Parse() error{

	imageBounds := parser.img.Bounds()
	boundsX, boundsY := imageBounds.Max.X, imageBounds.Max.Y

	var currentPixel containers.Pixel
	_ = currentPixel //"currentPixel declared and not used" even with loop below...
	
    for y := 0; y < boundsY; y++ {
        for x := 0; x < boundsX; x++ {
			currentPixel = convert32BitRGBAtoPixel(parser.img.At(x,y).RGBA()) //Current pixel at (x,y)
			
			if(x == 0){
				//get left bound from DM
				//add this pixel
			} else {
				//get left x-1,y from DM
				//add this pixel
			}
			
			if(y == 0){
				
			} else {
				
			}
        }
    }
	
	return nil
}

func (parser *SimpleParser) GetUpperBound() containers.Pixel {
	return containers.Pixel{0,0,0,0,2}
}

func (parser *SimpleParser) GetLeftBound() containers.Pixel {
	return containers.Pixel{0,0,0,0,1}
}
