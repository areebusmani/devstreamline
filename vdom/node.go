package vdom

type NodeType uint8

type Node struct {
	// ID of the figma node corresponding to this element.
	FigmaID string

	// Properties of the container.
	Properties Properties

	// Child nodes of this element. Will be nil if this container
	// is a UI component.
	Children []Node

	// UI Component type.
	UIComponentType UIComponentType
}

type LayoutType uint8

type Properties struct {
	LayoutType LayoutType
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
