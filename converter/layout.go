package converter

import (
	"devstreamline/figma"
	"math"
)

const (
	PIXEL_TOLERANCE = 10
)

// Populates the layout (using flex or grid) recursively until the
// function finds a ui component like text input or button which
// needs to be processed differently.
func PopulateLayout(figmaNodePtr *figma.Node, domNodePtr *Node) {
	figmaNode := *figmaNodePtr
	vdomNode := *domNodePtr

	vdomNode.FigmaID = figmaNode.ID
	vdomNode.LayoutProperties.LayoutType = getLayoutType(figmaNodePtr)
}

func getLayoutType(figmaNodePtr *figma.Node) (layoutType LayoutType) {
	figmaNode := *figmaNodePtr
	isCandidateForVertical := true
	isCandidateForHorizontal := true
	var lastHorizontalPosition float64
	var lastVerticalPosition float64

	for index, childNode := range figmaNode.Children {
		if index != 0 {
			if math.Abs(childNode.AbsoluteBoundingBox.X-lastHorizontalPosition) > PIXEL_TOLERANCE {
				isCandidateForHorizontal = false
			}
			if math.Abs(childNode.AbsoluteBoundingBox.Y-lastVerticalPosition) > PIXEL_TOLERANCE {
				isCandidateForVertical = false
			}
		}
		lastHorizontalPosition = childNode.AbsoluteBoundingBox.X
		lastVerticalPosition = childNode.AbsoluteBoundingBox.Y
	}
	if isCandidateForHorizontal {
		layoutType = LayoutTypeHorizontal
	} else if isCandidateForVertical {
		layoutType = LayoutTypeVertical
	}
	// TODO: Handle case for grid layout and unknown layout
	return
}
