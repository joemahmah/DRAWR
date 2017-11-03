package generators

import (
	"image"
	"image/color"
	"image/png"
	"image/draw"
	_ "image/jpeg"
	"github.com/joemahmah/DRAWR/containers"
	"os"
	"math/rand"
	"time"
	//"fmt"
)

var RandomGen *rand.Rand
var UseNoexpGen bool

func init(){
	RandomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	UseNoexpGen = false;
}

type Generator interface{
	SaveImage(string) error
	Generate() error
	
	GetUpperBound() containers.Pixel
	GetLeftBound() containers.Pixel
}

type SimpleGenerator struct {
	Img			image.Image
	Storage		containers.SimpleDataManager
}

/////////////
//Functions//
/////////////

func convertPixelto32BitRGBA(pixel containers.Pixel) color.NRGBA{
	return color.NRGBA{pixel.R, pixel.G, pixel.B, pixel.A}
	
}

func convert32BitRGBAtoPixel(r,g,b,a uint32) containers.Pixel{
	return containers.Pixel{byte(r/257),byte(g/257),byte(b/257),byte(a/257),0}
}

func MakeSimpleGenerator(sizex, sizey int) *SimpleGenerator{
	sg := new(SimpleGenerator)
	sg.Img = image.NewRGBA(image.Rect(0, 0, sizex, sizey))
	
	return sg
}

///////////
//Methods//
///////////

func (generator *SimpleGenerator) SaveImage(path string) error {
	imageFile, err := os.Create(path)
	defer imageFile.Close()
	
	if (err != nil) {
		return err
	}
	
	err = png.Encode(imageFile, generator.Img)

	return err
}

func (generator *SimpleGenerator) SetStorage(storage containers.SimpleDataManager) {
	generator.Storage = storage
}

func (generator *SimpleGenerator) Generate() error {
	imageBounds := generator.Img.Bounds()
	boundsX, boundsY := imageBounds.Max.X, imageBounds.Max.Y

	var currentPixel containers.Pixel
	_ = currentPixel //"currentPixel declared and not used" even with loop below...	
	
	var leftPixel containers.Pixel
	var upPixel containers.Pixel
	
	_ = leftPixel
	_ = upPixel
	
	//Container for all possible pixels
	pixelPool := make([]containers.Pixel,0,0)
	
    for y := 0; y < boundsY; y++ {
			//fmt.Println(y)
			
        for x := 0; x < boundsX; x++ {
			//fmt.Println(x,y)
		
		
			//Set left pixel
			if(x == 0){
				leftPixel = generator.GetLeftBound()
			} else {
				leftPixel = convert32BitRGBAtoPixel(generator.Img.At(x-1,y).RGBA())
			}
			
			//Set up pixel
			if(y == 0){
				upPixel = generator.GetUpperBound()
			} else {
				upPixel = convert32BitRGBAtoPixel(generator.Img.At(x,y-1).RGBA())
			}
			
			//Add all possible pixels into pool
			leftData, _ := generator.Storage.GetPixelData(leftPixel)
			upData, _ := generator.Storage.GetPixelData(upPixel)
			
			rightPixels := leftData.GetPixelsRight();
			downPixels := upData.GetPixelsBelow();
			
			if(UseNoexpGen){
				pixelPool = append(pixelPool, rightPixels...)
				pixelPool = append(pixelPool, downPixels...)
			
			} else {
				for _, data := range rightPixels {
					if(containsPixel(downPixels, data)){
						pixelPool = append(pixelPool, data)
					}
				}
				
				if(len(pixelPool) == 0){
					if(len(rightPixels) != 0){
						pixelPool = append(pixelPool, rightPixels...)
					} else if(len(downPixels) != 0) {
						pixelPool = append(pixelPool, downPixels...)
					} else { //No possible pixels
						pixelPool = append(pixelPool, containers.Pixel{255,255,255,255,0}) //add white
					}
				}
			}
			
			
			//fmt.Println(len(pixelPool))
			randLoc := RandomGen.Intn(len(pixelPool))
			currentPixel = pixelPool[randLoc]
			
			
			generator.Img.(draw.Image).Set(x,y,convertPixelto32BitRGBA(currentPixel)) 
			
			pixelPool = pixelPool[:0] //Empty slice
        }
    }

	return nil
}

func withinRange(a,b containers.Pixel, diviation int) bool {
	var good bool
	
	R := int(a.R) - int(b.R)
	G := int(a.G) - int(b.G)
	B := int(a.B) - int(b.B)
	
	good = R > -diviation && R < diviation
	good = good && G > -diviation && G < diviation
	good = good && B > -diviation && B < diviation
	
	/*if(good){
		fmt.Println(R,G,B)
	}*/
	return good
}

func containsPixel(slice []containers.Pixel, pix containers.Pixel) bool{
	for _, elem := range slice {
		if /*(elem == pix) {*/ (withinRange(elem, pix, 4)){
			return true
		}
	}
	
	return false;
}

func (generator *SimpleGenerator) GetUpperBound() containers.Pixel {
	return containers.Pixel{0,0,0,0,2}
}

func (generator *SimpleGenerator) GetLeftBound() containers.Pixel {
	return containers.Pixel{0,0,0,0,1}
}