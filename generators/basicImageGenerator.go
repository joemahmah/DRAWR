package generators

import (
	"image"
	"image/color"
	_ "image/png"
	_ "image/jpeg"
	"github.com/joemahmah/DRAWR/containers"
)

type Parser interface{
	SaveImage(string) error
	SetStorage(containers.DataManager)
	Generate() error
}

type SimpleGenerator struct {
	Img			image.Image
	Storage		containers.SimpleDataManager
}

/////////////
//Functions//
/////////////

func convertPixelto32BitRGBA(pixel containers.Pixel) color.RGBA{
	return color.RGBA{pixel.R, pixel.G, pixel.B, pixel.A}
	
}

///////////
//Methods//
///////////

func (generator *SimpleGenerator) SaveImage(path string) error {
	return nil
}

func (generator *SimpleGenerator) SetStorage(containers.DataManager) {
	
}

func (generator *SimpleGenerator) Generate() error {
	return nil
}