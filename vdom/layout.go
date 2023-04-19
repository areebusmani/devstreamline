package vdom

import (
	"devstreamline/figma"
)

// Populates the layout (using flex or grid) recursively until the
// function finds a ui component like text input or button which
// has to be processed differently.
func PopulateLayout(figmaNode *figma.Node, vdomNode *Node) {
}
