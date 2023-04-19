package converter

type LayoutType uint8

type NodeType uint8

const (
	// Container type. Would translate into div in the markup.
	TypeContainer NodeType = 0

	// UI Component type. Would render an external UI component.
	TypeUIComponent NodeType = 1
)

type LayoutProperties struct {
	LayoutType LayoutType
	Margin     Box
	Padding    Box
}

type Box struct {
	Top    int32
	Right  int32
	Bottom int32
	Left   int32
}

const (
	// LayoutTypeVertical is a layout with all child elements arranged vertically.
	LayoutTypeVertical LayoutType = 0

	// LayoutTypeVertical is a layout with all child elements arranged vertically.
	LayoutTypeHorizontal LayoutType = 1

	// LayoutTypeVertical is a layout with child elements arranged in a grid.
	LayoutTypeGrid LayoutType = 2
)

type UIComponentType uint8

const (
	// Input component.
	UIComponentTypeBreadcrumbs UIComponentType = 0

	// Search input component.
	UIComponentTypeSearch UIComponentType = 1

	// Tabs component.
	UIComponentTypeTabs UIComponentType = 2

	// Text Area component.
	UIComponentTypeTextArea UIComponentType = 3

	// Text Input component.
	UIComponentTypeTextInput UIComponentType = 4

	// Button component.
	UIComponentTypeButton UIComponentType = 5
)
