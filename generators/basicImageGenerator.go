package generators

import (
	"image"
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