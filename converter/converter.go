package converter

import (
	"devstreamline/figma"
	"devstreamline/vdom"
)

// Converts a figma node to a virtual dom node.
//   - Figma node is the node selected by the user to be converted to frontend code.
//   - Virtual dom is an internal representation of the output which can be directly
//     serialized to the desired frontend code.
func Convert(figmaNode *figma.Node) (vdomNode *vdom.Node) {
	node := vdom.Node{}
	vdom.PopulateLayout(figmaNode, &node)
	return &node
}
