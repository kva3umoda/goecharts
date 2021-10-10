package model

type Brush struct {
	// Use the buttons in toolbox.
	Toolbox       []ToolboxBrush `json:"toolbox,omitempty"`

	// Links interaction between selected items in different series.
	BrushLink     interface{}    `json:"brushLink,omitempty"`

	// Default type of brush.
	BrushType     BrushType      `json:"brushType,omitempty"`

	// Default brush mode, whose value can be:
	BrushMode     BrushMode      `json:"brushMode,omitempty"`

	// Determines whether a selected box can be changed in shape or translated.
	Transformable bool           `json:"transformable,omitempty"`
}

type ToolboxBrush string

const (
	ToolboxBrushRect    ToolboxBrush = "rect"
	ToolboxBrushPolygon ToolboxBrush = "polygon"
	ToolboxBrushLineX   ToolboxBrush = "lineX"
	ToolboxBrushLineY   ToolboxBrush = "lineY"
	ToolboxBrushKeep    ToolboxBrush = "keep"
	ToolboxBrushClear   ToolboxBrush = "clear"
)

type BrushType string

const (
	BrushTypeRect    BrushType = "rect"
	BrushTypePolygon BrushType = "polygon"
	BrushTypeLineX   BrushType = "lineX"
	BrushTypeLineY   BrushType = "lineY"
)

type BrushMode string

const (
	BrushModeSingle   BrushMode = "single"
	BrushModeMultiple BrushMode = "multiple"
)
