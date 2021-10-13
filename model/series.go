package model

type Series struct {
	Type             SeriesType       `json:"type"`
	CoordinateSystem CoordinateSystem `json:"coordinateSystem"`
	Name             string           `json:"name"`

	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Mark point in a chart.
	MarkPoint *MarkPoints `json:"markPoint,omitempty"`
	// Use a line in the chart to illustrate.
	MarkLine *MarkLines `json:"markLine,omitempty"`

	// For Line
	// The style of label line
	Smooth    bool       `json:"smooth,omitempty"`
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
	Encode    *Encode    `json:"encode,omitempty"`

	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`

	Stack string `json:"stack,omitempty"`

	// Highlight style of the graphic.
	Emphasis *Emphasis `json:"emphasis,omitempty"`

	// Text label of , to explain some data information about graphic item like value, name and so on
	Label *Label `json:"label,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int `json:"xAxisIndex,omitempty"`

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int `json:"yAxisIndex,omitempty"`
}

type SeriesType string

const (
	SeriesTypeLine          SeriesType = "line"
	SeriesTypeBar           SeriesType = "bar"
	SeriesTypePie           SeriesType = "pie"
	SeriesTypeScatter       SeriesType = "scatter"
	SeriesTypeEffectScatter SeriesType = "effectScatter"
	SeriesTypeRadar         SeriesType = "radar"
	SeriesTypeTree          SeriesType = "tree"
	SeriesTypeTreemap       SeriesType = "treemap"
	SeriesTypeSunburst      SeriesType = "sunburst"
	SeriesTypeBoxplot       SeriesType = "boxplot"
	SeriesTypeCandlestick   SeriesType = "candlestick"
	SeriesTypeHeatmap       SeriesType = "heatmap"
	SeriesTypeMap           SeriesType = "map"
	SeriesTypeParallel      SeriesType = "parallel"
	SeriesTypeLines         SeriesType = "lines"
	SeriesTypeGraph         SeriesType = "graph"
	SeriesTypeSankey        SeriesType = "sankey"
	SeriesTypeFunnel        SeriesType = "funnel"
	SeriesTypeGauge         SeriesType = "gauge"
	SeriesTypePictorialBar  SeriesType = "pictorialBar"
	SeriesTypeThemeRiver    SeriesType = "themeRiver"
	SeriesTypeCustom        SeriesType = "custom"
)

type CoordinateSystem string

const (
	CoordinateSystemCartesian2d CoordinateSystem = "cartesian2d"
	CoordinateSystemPolar       CoordinateSystem = "polar"
)

type Encode struct {
	// These properties only work in cartesian(grid) coordinate system:
	X []string `json:"x,omitempty"`
	Y []string `json:"y,omitempty"`

	// These properties only work in polar coordinate system
	Radius []string `json:"radius,omitempty"`
	Angle  []string `json:"angle,omitempty"`

	// These properties only work in geo coordinate system
	Lng []string `json:"lng,omitempty"`
	Lat []string `json:"lat,omitempty"`

	// For some type of series that are not in any coordinate system, like 'pie', 'funnel' etc.:
	Value []string `json:"value,omitempty"`
}

type Emphasis struct {
	// Whether to scale to highlight the data in emphasis state.
	Scale bool `json:"scale,omitempty"`
	// When the data is highlighted, whether to fade out of other data to focus the highlighted
	Focus Focus `json:"focus,omitempty"`
	// The range of fade out when focus is enabled. Support the following configurations
	BlurScope BlurScope `json:"blurScope,omitempty"`
}

type Focus string

const (
	// Do not fade out other data, it's by default.
	FocusNone Focus = "none"
	// Only focus (not fade out) the element of the currently highlighted data.
	FocusSelf Focus = "self"
	// Focus on all elements of the series which the currently highlighted data belongs to.
	FocusSeries Focus = "series"
)

type BlurScope string

const (
	BlurScopeCoordinateSystem BlurScope = "coordinateSystem"
	BlurScopeSeries           BlurScope = "series"
	BlurScopeGlobal           BlurScope = "global"
)

type Label struct {
	Show     bool     `json:"show"`
	Position Position `json:"position"`
}
