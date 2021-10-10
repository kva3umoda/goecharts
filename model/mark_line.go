package model

// MarkLines represents a series of marklines.
type MarkLines struct {
	// Whether to ignore mouse events. Default value is false, for triggering and responding to mouse events.
	Silent bool `json:"silent,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	Symbol Symbol `json:"symbol,omitempty"`
	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// Data array of marking line.
	Data []interface{} `json:"data,omitempty"`
}

type MarkLineData struct {
	// Name of the marker, which will display as a label.
	Name string `json:"name,omitempty"`
	// Special label types, are used to label maximum value, minimum value and so on.
	Type MarkLineType `json:"type,omitempty"`
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

type MarkLineType string

const (
	MarkLineTypeMin     = "min"
	MarkLineTypeMax     = "max"
	MarkLineTypeAverage = "average"
	MarkLineTypeMedian  = "median"
)

type Symbol string

const (
	SymbolNone      = "none"
	SymbolCircle    = "circle"
	SymbolRect      = "rect"
	SymbolRoundRect = "roundRect"
	SymbolTriangle  = "triangle"
	SymbolDiamond   = "diamond"
	SymbolPin       = "pin"
	SymbolArrow     = "arrow"
)
