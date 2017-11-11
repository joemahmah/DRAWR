package containers

import (
	"errors"
	"math/rand"
	"time"
)

var RandomGen *rand.Rand

const (
	PixelError = 0
	PixelEqual = 1
	PixelLess = 2
	PixelGreater = 3
)

//Package Init//
func init(){
	RandomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Pixel struct {
	R		byte
	G		byte
	B		byte
	A		byte
	
	Edge	byte //0 = no, 1 = left, 2 = top
}

type PixelTreeNode struct {
	Key			Pixel
	Count		int
	LeftNode	*PixelTreeNode
	RightNode	*PixelTreeNode
}

type PixelTree struct {
	RootNode *PixelTreeNode
}

type PixelData interface{
	AddPixelBelow(Pixel)
	AddPixelRight(Pixel)
	GetPixelsBelow() []Pixel
	GetPixelsRight() []Pixel
	GetRandomPixelBelow() Pixel
	GetRandomPixelRight() Pixel
	
	GetColor() Pixel
}

type SimplePixelData struct {
	PixelBelow	[]Pixel
	PixelRight	[]Pixel
	
	Color		Pixel
}

type DataManager interface{
	GetPixelData(Pixel) (PixelData, error)
	GetPixelDataCreateIfNotExist(Pixel) PixelData
}

type SimpleDataManager struct {
	Data		map[Pixel]*SimplePixelData
}

/////////////
//Functions//
/////////////

func MakeSimpleDataManager() *SimpleDataManager{
	sdm := new(SimpleDataManager)
	sdm.Data = make(map[Pixel]*SimplePixelData)
	
	return sdm
}

func MakePixelTreeNode(pixel Pixel, count int) *PixelTreeNode{
	newNode := new(PixelTreeNode)
	newNode.Key = pixel
	newNode.Count = count
	newNode.LeftNode = nil
	newNode.RightNode = nil
	
	return newNode
}


//Note: 1=equal, 2=lessthan, 3=greaterthan, 0=error
//Equality based on R then G then B then A then Edge
func GetPixelRelation(pixel1 *Pixel, pixel2 *Pixel) int{

	if(pixel1.R > pixel2.R){
		return PixelGreater
		
	} else if (pixel1.R < pixel2.R) {
		return PixelLess
		
	} else { //If reds are the same
		if(pixel1.G > pixel2.G){
			return PixelGreater
			
		} else if (pixel1.G < pixel2.G) {
			return PixelLess
			
		} else { //If greens are the same
			if(pixel1.B > pixel2.B){
				return PixelGreater
				
			} else if (pixel1.B < pixel2.B) {
				return PixelLess
				
			} else { //If blues are the same
					if(pixel1.A > pixel2.A){
						return PixelGreater
						
					} else if (pixel1.A < pixel2.A) {
						return PixelLess
						
					} else { //If alphas are the same
						if(pixel1.Edge > pixel2.Edge){
							return PixelGreater
							
						} else if (pixel1.Edge < pixel2.Edge) {
							return PixelLess
							
						} else { //If edge states are the same
							//These are the same
							return PixelEqual
						}
					}
			}
		}
	}
	
	//This should never happen
	//But if it does...
	return PixelError
}

///////////
//Methods//
///////////

// PixelTreeNode
func (node *PixelTreeNode) Add(pixel Pixel, count int){
	relation := GetPixelRelation(&node.Key, &pixel)
	
	switch relation {
	case PixelEqual:
		node.Count += count;
		
	case PixelGreater:
		if(node.RightNode == nil){ //If node empty
			node.RightNode = MakePixelTreeNode(pixel, count)
		} else { //If node full
			node.RightNode.Add(pixel, count)
		}
	
	case PixelLess:
		if(node.LeftNode == nil){ //If node empty
			node.LeftNode = MakePixelTreeNode(pixel, count)
		} else { //If node full
			node.LeftNode.Add(pixel, count)
		}

	default:
		//TODO: Add error handling
	}
}

//returns the node with key pixel (if not exist, nil)
func (node *PixelTreeNode) GetNode(pixel Pixel) *PixelTreeNode{
	relation := GetPixelRelation(&node.Key, &pixel)
	
	switch relation {
	case PixelEqual:
		return node
		
	case PixelGreater:
		if(node.RightNode == nil){ //If node empty
			return nil
		} else { //If node full
			return node.RightNode.GetNode(pixel)
		}
	
	case PixelLess:
		if(node.LeftNode == nil){ //If node empty
			return nil
		} else { //If node full
			return node.LeftNode.GetNode(pixel)
		}

	default:
		return nil
	}
	
	//Should never hit here
	return nil
}

//returns the node with key pixel (if not exist, nil)
func (node *PixelTreeNode) GetCount(pixel Pixel) (int, error){
	relation := GetPixelRelation(&node.Key, &pixel)
	
	switch relation {
	case PixelEqual:
		return node.Count, nil
		
	case PixelGreater:
		if(node.RightNode == nil){ //If node empty
			return 0, errors.New("Node does not exist.")
		} else { //If node full
			return node.RightNode.GetCount(pixel)
		}
	
	case PixelLess:
		if(node.LeftNode == nil){ //If node empty
			return 0, errors.New("Node does not exist.")
		} else { //If node full
			return node.LeftNode.GetCount(pixel)
		}

	default:
		return 0, errors.New("Node does not exist.")
	}
	
	//Should never hit here
	return 0, errors.New("Node does not exist.")
}

//PixelTree
func (tree *PixelTree) Add(pixel Pixel, count int){
	if(tree.RootNode != nil) {
		tree.RootNode.Add(pixel, count)
	} else { //if tree empty
		tree.RootNode = MakePixelTreeNode(pixel, count)
	}

}

func (tree *PixelTree) GetNode(pixel Pixel) *PixelTreeNode{
	if(tree.RootNode != nil) {
		return tree.RootNode.GetNode(pixel);
	} else { //if tree empty
		return nil
	}

}

func (tree *PixelTree) GetCount(pixel Pixel) (int, error){
	if(tree.RootNode != nil) {
		return tree.RootNode.GetCount(pixel);
	} else { //if tree empty
		return 0, errors.New("Tree empty.")
	}
}

//Pixel
func (p *Pixel) SetRGB(r,g,b byte){
	p.R = r
	p.G = g
	p.B = b
}

func (p *Pixel) SetRGBA(r,g,b,a byte){
	p.R = r
	p.G = g
	p.B = b
	p.A = a
}

func (p *Pixel) SetAlpha(a byte){
	p.A = a
}

func (p *Pixel) SetEdge(e byte){
	p.Edge = e
}

func (p *Pixel) getRGB() (byte,byte,byte){
	return p.R, p.G, p.B
}

func (p *Pixel) getRGBA() (byte,byte,byte,byte){
	return p.R, p.G, p.B, p.A
}

func (p *Pixel) IsEdge() bool{
	return p.Edge != 0
}

//SimplePixelData
func (p *SimplePixelData) AddPixelBelow(pixel Pixel){
	p.PixelBelow = append(p.PixelBelow, pixel)
}

func (p *SimplePixelData) AddPixelRight(pixel Pixel){
	p.PixelRight = append(p.PixelRight, pixel)
}

func (p *SimplePixelData) GetColor() Pixel {
	return p.Color
}

func (p *SimplePixelData) GetPixelsBelow() []Pixel{
	return p.PixelBelow
}

func (p *SimplePixelData) GetPixelsRight() []Pixel{
	return p.PixelRight
}

func (p *SimplePixelData) GetRandomPixelBelow() Pixel{
	return p.PixelBelow[RandomGen.Intn(len(p.PixelBelow))]
}

func (p *SimplePixelData) GetRandomPixelRight() Pixel{
	return p.PixelRight[RandomGen.Intn(len(p.PixelRight))]
}

//SimpleDataManager
func (dm *SimpleDataManager) GetPixelData(pixel Pixel) (PixelData,error) {
	pixelData, exists := dm.Data[pixel]
	
	if(!exists){
		return nil,errors.New("Pixel has not been stored")
	}
	
	return pixelData,nil
}

func (dm *SimpleDataManager) GetPixelDataCreateIfNotExist(pixel Pixel) PixelData{
	pixelData, exists := dm.Data[pixel]
	
	if(!exists){
		dm.Data[pixel] = new(SimplePixelData)
		pixelData, _ = dm.Data[pixel]
	}
	
	
	return pixelData
}