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
	"errors"
)

var (
	RandomGen *rand.Rand
)

const (
	SimpleGeneratorModeStandard int = 0
	SimpleGeneratorModeExclusive int = 1
	SimpleGeneratorModeMultiplicative int = 2
	SimpleGeneratorModeSuperMultiplicative int = 3
	SimpleGeneratorModeInverseMultiplicative int = 4
	SimpleGeneratorModeInverseSuperMultiplicative int = 5
)

func init(){
	RandomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}


type Generator interface{
	SaveImage(string) error
	Generate(int) error
	
	GetUpperBound() containers.Pixel
	GetLeftBound() containers.Pixel
}

type SimpleGenerator struct {
	Img			image.Image
	Storage		containers.SimpleDataManager
}



/////////////
/////////////
//Functions//
/////////////
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
///////////
//Methods//
///////////
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

func (generator *SimpleGenerator) Generate(generationMode int) error {
	imageBounds := generator.Img.Bounds()
	boundsX, boundsY := imageBounds.Max.X, imageBounds.Max.Y

	var currentPixel containers.Pixel
	_ = currentPixel //"currentPixel declared and not used" even with loop below...	
	
	var leftPixel 			containers.Pixel
	var leftPixelPrevious 	containers.Pixel
	var upPixel				containers.Pixel
	var upPixelPrevious		containers.Pixel
	
	_ = leftPixel
	leftPixelPrevious = containers.Pixel{0,0,0,0,255} //Not valid edge so it will always be different from leftPixel
	_ = upPixel
	upPixelPrevious = containers.Pixel{0,0,0,0,255} //Not valid edge so it will always be different from upPixel
	
	//Container for all possible pixels
	var pixelPool 			containers.PixelTree
	
	//temp var
	var tempCount int
	
    for y := 0; y < boundsY; y++ {
			
        for x := 0; x < boundsX; x++ {
		
			//Set previous pixels
			leftPixelPrevious = leftPixel
			upPixelPrevious = upPixel
		
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
			
			
			//fmt.Println("Left: ", leftPixel, leftData,"\nUp: ", upPixel, upData)
			//fmt.Println("Err: ", err)
			
			rightPixels := leftData.GetPixelsRight();
			downPixels := upData.GetPixelsBelow();
			
			//If the previous pixels are not the same, recalculate.
			//If they are the same, just reuse the data to save time.
			if(leftPixel != leftPixelPrevious || upPixel != upPixelPrevious){
				switch generationMode{
				case SimpleGeneratorModeStandard: //Don't force any type to be more common
					pixelPool.AddTree(rightPixels)
					pixelPool.AddTree(downPixels)
				
				case SimpleGeneratorModeExclusive: //Generate images with pixels exclusivly in both up and left (if unable to do so, just take the left/up)
					for _, node := range rightPixels.GetNodeSlice() {
						if(downPixels.Contains(node.Key)){
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					if(pixelPool.IsEmpty()){
						if(!rightPixels.IsEmpty()){
							pixelPool.AddTree(rightPixels)
						} else if(!downPixels.IsEmpty()) {
							pixelPool.AddTree(downPixels)
						} else { //No possible pixels
							pixelPool.Add(containers.Pixel{255,255,255,255,0},1) //add white
						}
					}
					
				case SimpleGeneratorModeMultiplicative: //Generate images with pixels both up and left being much more common (~100x multiplier)
					for _, node := range rightPixels.GetNodeSlice() {
						if(downPixels.Contains(node.Key)){
							pixelPool.Add(node.Key, node.Count * 50)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					for _, node := range downPixels.GetNodeSlice() {
						if(rightPixels.Contains(node.Key)){
							pixelPool.Add(node.Key, node.Count * 50)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					
				case SimpleGeneratorModeSuperMultiplicative: //Generate images with pixels both up and left being much, much more common (~2000x multiplier)
					for _, node := range rightPixels.GetNodeSlice() {
						if(downPixels.Contains(node.Key)){
							pixelPool.Add(node.Key, node.Count * 1000)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					for _, node := range downPixels.GetNodeSlice() {
						if(rightPixels.Contains(node.Key)){
							pixelPool.Add(node.Key, node.Count * 1000)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
				case SimpleGeneratorModeInverseMultiplicative: //Generate images with pixels both up and left being much less common (~1/5 multiplier)
					for _, node := range rightPixels.GetNodeSlice() {
						if(downPixels.Contains(node.Key)){
							tempCount = node.Count / 10
							if(tempCount <= 0){
								tempCount = 1
							}
							pixelPool.Add(node.Key, tempCount)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					for _, node := range downPixels.GetNodeSlice() {
						if(rightPixels.Contains(node.Key)){
							tempCount = node.Count / 10
							if(tempCount <= 0){
								tempCount = 1
							}
							pixelPool.Add(node.Key, tempCount)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
				case SimpleGeneratorModeInverseSuperMultiplicative: //Generate images with pixels both up and left being much, much less common (~1/250 multiplier)
					for _, node := range rightPixels.GetNodeSlice() {
						if(downPixels.Contains(node.Key)){
							tempCount = node.Count / 500
							if(tempCount <= 0){
								tempCount = 1
							}
							pixelPool.Add(node.Key, tempCount)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					for _, node := range downPixels.GetNodeSlice() {
						if(rightPixels.Contains(node.Key)){
							tempCount = node.Count / 500
							if(tempCount <= 0){
								tempCount = 1
							}
							pixelPool.Add(node.Key, tempCount)
						} else {
							pixelPool.Add(node.Key, node.Count)
						}
					}
					
					
				default:
					return errors.New("Unknown generation mode!")
					
				}
			}
			
			
			currentPixel = pixelPool.GetRandomPixel()
			
			generator.Img.(draw.Image).Set(x,y,convertPixelto32BitRGBA(currentPixel)) 
			
			pixelPool.RootNode = nil//Empty slice
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