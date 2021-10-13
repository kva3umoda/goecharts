package model

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type Option struct {
	// Title component, including main title and subtitle.
	Title *Title `json:"title,omitempty"`

	// Legend component shows symbol, color and name of different series. You can click legends to toggle displaying series in the chart.
	Legend *Legend `json:"legend,omitempty"`

	// Drawing grid in rectangular coordinate. In a single grid, at most two X and Y axes each is allowed. Line chart, bar chart, and scatter chart (bubble chart) can be drawn in grid.
	Grids []*Grid `json:"grid,omitempty"`

	// The x axis in cartesian(rectangular) coordinate. Usually a single grid component can place at most 2 x axis, one on the bottom and another on the top. offset can be used to avoid overlap when you need to put more than two x axis.
	XAxis []*XYAxis `json:"xAxis,omitempty"`

	// The y axis in cartesian(rectangular) coordinate. Usually a single grid component can place at most 2 y axis, one on the left and another on the right. offset can be used to avoid overlap when you need to put more than two y axis.
	YAxis []*XYAxis `json:"yAxis,omitempty"`

	// Polar coordinate can be used in scatter and line chart. Every polar coordinate has an angleAxis and a radiusAxis.
	Polar *Polar `json:"polar,omitempty"`

	// Radial axis of polar coordinate.
	RadiusAxis *RadiusAxis `json:"radiusAxis,omitempty"`

	// The angle axis in Polar Coordinate.
	AngleAxis *AngleAxis `json:"angleAxis,omitempty"`

	// Coordinate for radar charts. This component is equal to the polar component in ECharts 2. Because the polar component in the echarts 3 is reconstructed to be the standard polar coordinate component, this component is renamed to be radar to avoid confusion.
	Radar *Radar `json:"radar,omitempty"`

	// dataZoom component is used for zooming a specific area, which enables user to investigate data in detail, or get an overview of the data, or get rid of outlier points.
	DataZoom []interface{} `json:"dataZoom,omitempty"`

	// visualMap is a type of component for visual encoding, which maps the data to visual channels, including:
	VisualMap []VisualMap `json:"visualMap,omitempty"`

	// Tooltip component.
	Tooltip *Tooltip `json:"tooltip,omitempty"`

	// This is the global configurations of axisPointer.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`

	// A group of utility tools, which includes export, data view, dynamic type switching, data area zooming, and reset.
	Toolbox *Toolbox `json:"toolbox,omitempty"`

	// brush is an area-selecting component, with which user can select part of data from a chart to display in detail, or do calculations with them.
	Brush *Brush `json:"brush,omitempty"`

	// Geographic coorinate system component is used to draw maps, which also supports scatter series, and line series.
	Geo *GeoComponent `json:"geo,omitempty"`

	// Parallel Coordinates is a common way of visualizing high-dimensional geometry and analyzing multivariate data.
	Parallel *ParallelComponent `json:"parallel,omitempty"`

	// Parallel Coordinates is a common way of visualizing high-dimensional geometry and analyzing multivariate data.
	ParallelAxis *ParallelAxis `json:"parallelAxis,omitempty"`

	// An axis with a single dimension. It can be used to display data in one dimension. For example:
	SingleAxis *SingleAxis `json:"singleAxis,omitempty"`

	// timeline component, which provides functions like switching and playing between multiple ECharts options.
	Timeline *Timeline `json:"timeline,omitempty"`
	// graphic component enables creating graphic elements in ECharts.
	// Those graphic type are supported: image, text, circle, sector, ring, polygon, polyline, rect, line, bezierCurve, arc, group,
	Graphic []*Graphic `json:"graphic,omitempty"`
	// Calendar coordinates.
	Calendar *Calendar `json:"calendar,omitempty"`

	Aria    *Aria     `json:"aria,omitempty"`
	Series  []*Series `json:"series,omitempty"`

	// The color list of palette. If no color is set in series, the colors would be adopted sequentially and circularly from this list as the colors of series.
	Color []string `json:"color,omitempty"`

	// Background color. Defaults to have no background.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Global font style.
	TextStyle *TextStyle `json:"textStyle,omitempty"`

	// Whether to enable animation.
	Animation bool `json:"animation,omitempty"`
	// Whether to set graphic number threshold to animation. Animation will be disabled when graphic number is larger than threshold.
	AnimationThreshold int `json:"animationThreshold,omitempty"`

	// Duration of the first animation, which supports callback function for different data to have different animation effect:
	AnimationDuration int `json:"animationDuration,omitempty"`

	// Easing method used for the first animation. Varied easing effects can be found at easing effect example.
	AnimationEasing interface{} `json:"animationEasing,omitempty"`

	// Delay before updating the first animation, which supports callback function for different data to have different animation effect.
	AnimationDelay int `json:"animationDelay,omitempty"`

	// Time for animation to complete, which supports callback function for different data to have different animation effect:
	AnimationDurationUpdate int `json:"animationDurationUpdate,omitempty"`

	// Delay before updating animation, which supports callback function for different data to have different animation effects.
	AnimationDelayUpdate int `json:"animationDelayUpdate,omitempty"`

	// Animation configurations of state switchment. Can be configured in each series individually.
	StateAnimation interface{} `json:"stateAnimation,omitempty"`

	//
	BlendMode interface{} `json:"blendMode,omitempty"`

	// When the number of element of the whole chart is larger than hoverLayerThreshold, a seperate hover layer is used to render hovered elements.
	HoverLayerThreshold int `json:"hoverLayerThreshold,omitempty"`

	// Whether to use UTC in display.
	UseUTC bool `json:"useUTC,omitempty"`

	// dataset brings convenience in data management separated with styles and enables data reuse by different series. More importantly, it enables data encoding from data to visual, which brings convenience in some scenarios.
	Dataset *Dataset  `json:"dataset,omitempty"`
}

/*
func (op *Option) JSON() map[string]interface{} {
	return op.json()
}
*/
// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (op *Option) JSONNotEscaped() template.HTML {

	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	err := enc.Encode(op)
	if err != nil {
		panic(err)
	}

	return template.HTML(buff.String())
}
