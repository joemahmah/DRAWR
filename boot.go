package main

import (
	"fmt"
	"flag"
	"log"
	//"io/ioutil"
	//"os"
	"github.com/joemahmah/DRAWR/containers"
	"github.com/joemahmah/DRAWR/parsers"
	"github.com/joemahmah/DRAWR/generators"
)

func main() {

	var interactiveMode	bool
	var sizeX 			int
	var sizeY 			int
	var inputPath 		string
	var outputPath		string
	var mode			int
	
	flag.BoolVar(&interactiveMode, "im", false, "Enable interactive mode.")
	flag.BoolVar(&interactiveMode, "interactive", false, "Enable interactive mode.")
	flag.IntVar(&sizeX, "x", 300, "Set image width (default 300).")
	flag.IntVar(&sizeY, "y", 300, "Set image width (default 300).")
	flag.StringVar(&inputPath, "i", "test.png", "Set input image (default test.png).")
	flag.StringVar(&outputPath, "o", "imgOut.png", "Set output image (default imgOut.png).")
	flag.IntVar(&mode, "m", 0, "Sets the mode for the generator used (default to standard mode of 0). See README for list of modes with each generator.")
	flag.IntVar(&mode, "mode", 0, "Sets the mode for the generator used (default to standard mode of 0). See README for list of modes with each generator.")
	
	flag.Parse()
	
	if(interactiveMode){
		fmt.Println("Interactive mode not yet implemented.")
	} else {
		testGen := generators.MakeSimpleGenerator(sizeX,sizeY) //testGen is pointer
		var testPars parsers.SimpleParser
		data := containers.MakeSimpleDataManager() //data is a pointer
		
		
		testPars.SetStorage(*data)
		err := testPars.LoadImage(inputPath)
		
		if(err != nil){
			log.Fatal(err)
		}
		
		testPars.Parse()
		
		testGen.SetStorage(testPars.Storage)
		
		testGen.Generate(mode)
		testGen.SaveImage(outputPath)
	}
}