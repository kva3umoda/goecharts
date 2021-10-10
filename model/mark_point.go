package model

// MarkPoints represents a series of markpoints.
type MarkPoints struct {

	// Whether to ignore mouse events. Default value is false, for triggering and responding to mouse events.
	Silent bool `json:"silent,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol Symbol `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`

	// Mark point text options.
	Label *Label `json:"label,omitempty"`

	Data []interface{} `json:"data,omitempty"`
}

type MarkPointData struct {
	// Mark point name.
	Name string `json:"name,omitempty"`

	// Special label types, are used to label maximum value, minimum value and so on.
	Type MarkPointType

	// Works only when type is assigned. It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x, or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
	// Coordinates of the starting point or ending point
	Coord []interface{} `json:"coord,omitempty"`
	// Markline at x at given value, which only works for single data item
	XAxis interface{} `json:"xAxis,omitempty"`
	// Markline at y at given value, which only works for single data item
	YAxis interface{} `json:"yAxis,omitempty"`
	// Symbol of starting point.
	Symbol Symbol `json:"symbol,omitempty"`
}

type MarkPointType string

const (
	MarkPointTypeMin     = "min"
	MarkPointTypeMax     = "max"
	MarkPointTypeAverage = "average"
)
