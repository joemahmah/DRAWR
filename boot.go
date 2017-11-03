package main

import (
	"fmt"
	//"flag"
	//"io/ioutil"
	"github.com/joemahmah/DRAWR/containers"
	"github.com/joemahmah/DRAWR/parsers"
	"github.com/joemahmah/DRAWR/generators"
)

func main() {

	testGen := generators.MakeSimpleGenerator(300,300) //testGen is pointer
	var testPars parsers.SimpleParser
	data := containers.MakeSimpleDataManager() //data is a pointer
	
	
	testPars.SetStorage(*data)
	err := testPars.LoadImage("test.png")
	
	if(err != nil){
		fmt.Println("ERROR: ", err)
	}
	
	testPars.Parse()
	
	testGen.SetStorage(testPars.Storage)
	
	generators.UseNoexpGen = false;
	
	testGen.Generate()
	testGen.SaveImage("imgOut.png")
	
	
	//Decl Run Param Vars
	//Interprate flags
	
	//open containers (if flag) OR create new container (if no flag)
	//Parse images (if images specified)
	//save to containers (if flag)
	
	//generate with generator setting (from flags)
	//save image (flag for location, else use a default)

}
