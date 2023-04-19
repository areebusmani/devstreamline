package converter

import (
	"devstreamline/figma"
)

func Convert(figmaNode *figma.Node) *Node {
	node := Node{}
	PopulateLayout(figmaNode, &node)
	return &node
}
