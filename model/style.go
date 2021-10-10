package model

// LineStyle is the option set for a link style component.
type LineStyle struct {
	// Line color
	Color string `json:"color,omitempty"`

	// Width of line. default 1
	Width float32 `json:"width,omitempty"`

	// Type of line，options: "solid", "dashed", "dotted". default "solid"
	Type LineType `json:"type,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`

	// Curveness of edge. The values from 0 to 1 could be set.
	// it would be larger as the the value becomes larger. default 0
	Curveness float32 `json:"curveness,omitempty"`
}

type LineType string

const (
	LineTypeSolid  LineType = "solid"
	LineTypeDashed LineType = "dashed"
	LineTypeDotted LineType = "dotted"
)

// AreaStyle is the option set for an area style component.
type AreaStyle struct {
	// Fill area color.
	Color string `json:"color,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`
}
