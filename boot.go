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

	var testGen generators.SimpleGenerator
	_ = testGen
	
	var testPars parsers.SimpleParser
	data := containers.MakeSimpleDataManager()
	
	testPars.SetStorage(data)
	err := testPars.LoadImage("test.png")
	
	if(err != nil){
		fmt.Println("ERROR: ", err)
	}
	
	testPars.Parse()
	
	for key, value := range testPars.Storage.Data{
		fmt.Println("Key: " , key, "\nValue: ", value)
	}
	
	//Decl Run Param Vars
	//Interprate flags
	
	//open containers (if flag) OR create new container (if no flag)
	//Parse images (if images specified)
	//save to containers (if flag)
	
	//generate with generator setting (from flags)
	//save image (flag for location, else use a default)

}
