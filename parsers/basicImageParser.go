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
}

type SimpleParser struct {
	img			image.Image
	storage		containers.SimpleDataManager
}

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

	

	return nil
}