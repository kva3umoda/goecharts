package model

// VisualMap is a type of component for visual encoding, which maps the data to visual channels.
// https://echarts.apache.org/en/option.html#visualMap
type VisualMap struct {
	// Mapping type.
	// Options: "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`

	// Whether show handles, which can be dragged to adjust "selected range".
	Calculable bool `json:"calculable"`

	// Specify the min dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Min float32 `json:"min,omitempty"`

	// Specify the max dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Max float32 `json:"max,omitempty"`

	// Specify selected range, that is, the dataValue corresponding to the two handles.
	Range []float32 `json:"range,omitempty"`

	// The label text on both ends, such as ['High', 'Low'].
	Text []string `json:"text,omitempty"`

	// Define visual channels that will mapped from dataValues that are in selected range.
	InRange *VisualMapInRange `json:"inRange,omitempty"`
}


// VisualMapInRange is a visual map instance in a range.
type VisualMapInRange struct {
	// Color
	Color []string `json:"color,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`
}
