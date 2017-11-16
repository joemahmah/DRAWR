package containers

import (
	"errors"
	"math/rand"
	"time"
	"io"
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
	RootNode 	*PixelTreeNode
	NodeSlice	[]PixelTreeNode
	IsFlat		bool
	TotalCount	int64 //Summation of node.count for all nodes in the nodeslice
}

type PixelData interface{
	AddPixelBelow(Pixel)
	AddPixelRight(Pixel)
	GetPixelsBelow() 			*PixelTree
	GetPixelsRight() 			*PixelTree
	GetPixelsAt(int,int) 		*PixelTree
	GetRandomPixelBelow() 		Pixel
	GetRandomPixelRight() 		Pixel
	GetRandomPixelAt(int,int)	Pixel
	
	GetColor() Pixel
}

type SimplePixelData struct {
	PixelBelow	PixelTree
	PixelRight	PixelTree
	
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

func (node *PixelTreeNode) GetNodeSlice() []PixelTreeNode{
	nodeSlice := make([]PixelTreeNode,0)
	
	if(node.LeftNode != nil){
		nodeSlice = append(nodeSlice, node.LeftNode.GetNodeSlice()...)
	}
	
	nodeSlice = append(nodeSlice, *node)
	
	if(node.RightNode != nil){
		nodeSlice = append(nodeSlice, node.RightNode.GetNodeSlice()...)
	}
	
	return nodeSlice
}

//returns if a node with key pixel is in tree
func (node *PixelTreeNode) Contains(pixel Pixel) bool{
	relation := GetPixelRelation(&node.Key, &pixel)
	
	switch relation {
	case PixelEqual:
		return true
		
	case PixelGreater:
		if(node.RightNode == nil){ //If node empty
			return false
		} else { //If node full
			return node.RightNode.Contains(pixel)
		}
	
	case PixelLess:
		if(node.LeftNode == nil){ //If node empty
			return false
		} else { //If node full
			return node.LeftNode.Contains(pixel)
		}

	default:
		return false
	}
	
	//Should never hit here
	return false
}

func (node *PixelTreeNode) Print(writer io.Writer) {
	node.LeftNode.Print(writer)
	writer.Write([]byte(" "))
	writer.Write([]byte(string(node.Key.R)))
	writer.Write([]byte(string(node.Key.G)))
	writer.Write([]byte(string(node.Key.B)))
	writer.Write([]byte(string(node.Key.A)))
	writer.Write([]byte(string(node.Key.Edge)))
	writer.Write([]byte(" "))
	node.RightNode.Print(writer)
}

//PixelTree
func (tree *PixelTree) Add(pixel Pixel, count int){
	if(tree.RootNode != nil) {
		tree.RootNode.Add(pixel, count)
	} else { //if tree empty
		tree.RootNode = MakePixelTreeNode(pixel, count)
	}

	tree.IsFlat = false
}

func (tree *PixelTree) AddTree(nodes *PixelTree){
	for _, node := range nodes.GetNodeSlice() {
		tree.Add(node.Key, node.Count)
	}
}

func (tree *PixelTree) GetNode(pixel Pixel) *PixelTreeNode{
	if(tree.RootNode != nil) {
		return tree.RootNode.GetNode(pixel);
	} else { //if tree empty
		return nil
	}
}

func (tree *PixelTree) Contains(pixel Pixel) bool{
	if(tree.RootNode != nil) {
		return tree.RootNode.Contains(pixel);
	} else { //if tree empty
		return false
	}
}

func (tree *PixelTree) GetCount(pixel Pixel) (int, error){
	if(tree.RootNode != nil) {
		return tree.RootNode.GetCount(pixel);
	} else { //if tree empty
		return 0, errors.New("Tree empty.")
	}
}

func (tree *PixelTree) Flatten() {
	if(tree.RootNode != nil){
		tree.NodeSlice = tree.RootNode.GetNodeSlice()
	} else {
		tree.NodeSlice = make([]PixelTreeNode, 0)
	}
	tree.TotalCount = 0;
	
	for _, node := range tree.NodeSlice {
		tree.TotalCount += int64(node.Count)
	}
	
	tree.IsFlat = true
}

func (tree *PixelTree) GetNodeSlice() []PixelTreeNode {
	if(!tree.IsFlat){
		tree.Flatten()
	}
	
	return tree.NodeSlice
}

func (tree *PixelTree) IsEmpty() bool {
	return tree.RootNode == nil
}

func (tree *PixelTree) GetRandomPixel() Pixel {
	if(!tree.IsFlat){
		tree.Flatten()
	}
	
	targetCount := RandomGen.Intn(int(tree.TotalCount)) //Get a random value
	currentCount := 0;
	
	for _, node := range tree.NodeSlice {
		if(targetCount >= currentCount && targetCount <= currentCount + node.Count){
			return node.Key
		} else {
			currentCount += node.Count
		}
	}
	
	return tree.RootNode.Key
}

func (tree *PixelTree) Print(writer io.Writer) {
	writer.Write([]byte("{"))
	tree.RootNode.Print(writer)
	writer.Write([]byte("}"))
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
	p.PixelBelow.Add(pixel, 1)
}

func (p *SimplePixelData) AddPixelRight(pixel Pixel){
	p.PixelRight.Add(pixel, 1)
}

func (p *SimplePixelData) GetColor() Pixel {
	return p.Color
}

func (p *SimplePixelData) GetPixelsBelow() *PixelTree{
	return &p.PixelBelow
}

func (p *SimplePixelData) GetPixelsRight() *PixelTree{
	return &p.PixelRight
}

//Since SimplePixelData only has above and below, this
//method is only present to satisfy the interface
//If the y value is >0, returns below else right
func (p *SimplePixelData) GetPixelsAt(x, y int) *PixelTree{
	if(y > 0){
		return &p.PixelBelow
	}
	
	return &p.PixelRight
}

func (p *SimplePixelData) GetRandomPixelBelow() Pixel{
	return p.PixelBelow.GetRandomPixel()
}

func (p *SimplePixelData) GetRandomPixelRight() Pixel{
	return p.PixelRight.GetRandomPixel()
}

func (p *SimplePixelData) GetRandomPixelAt(x, y int) Pixel{
	if(y > 0){
		return p.PixelBelow.GetRandomPixel()
	}
	
	return p.PixelRight.GetRandomPixel()
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