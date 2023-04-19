package converter

type Node struct {
	// ID of the figma node corresponding to this element.
	FigmaID string

	// Type of the node.
	Type NodeType

	// Properties of the layout like layout direction, alignment and spacings (margin and padding).
	LayoutProperties LayoutProperties

	// Child nodes of this element. Will be nil if this container
	// is a UI component.
	Children []Node

	// UI Component type.
	UIComponentType UIComponentType
}
