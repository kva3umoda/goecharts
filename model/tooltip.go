package model

// Tooltip is the option set for a tooltip component.
// https://echarts.apache.org/en/option.html#tooltip
// tooltip can be configured on different places:
// Configured on global: tooltip
// Configured in a coordinate system: grid.tooltip, polar.tooltip, single.tooltip
// Configured in a series: series.tooltip
// Configured in each item of series.data: series.data.tooltip
type Tooltip struct {
	// Whether to show the tooltip component, including tooltip floating layer and axisPointer.
	//Show bool `json:"show"`

	// Type of triggering.
	// Options:
	// * 'item': Triggered by data item, which is mainly used for charts that
	//    don't have a category axis like scatter charts or pie charts.
	// * 'axis': Triggered by axes, which is mainly used for charts that have category axes,
	//    like bar charts or line charts.
	// * 'none': Trigger nothing.
	Trigger TriggerType `json:"trigger,omitempty"`

	// Configuration item for axisPointer
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`

	// Conditions to trigger tooltip. Options:
	// * 'mousemove': Trigger when mouse moves.
	// * 'click': Trigger when mouse clicks.
	// * 'mousemove|click': Trigger when mouse clicks and moves.
	// * 'none': Do not triggered by 'mousemove' and 'click'. Tooltip can be triggered and hidden
	//    manually by calling action.tooltip.showTip and action.tooltip.hideTip.
	//    It can also be triggered by axisPointer.handle in this case.
	TriggerOn string `json:"triggerOn,omitempty"`

	// The content formatter of tooltip's floating layer which supports string template and callback function.
	//
	// 1. String template
	// The template variables are {a}, {b}, {c}, {d} and {e}, which stands for series name,
	// data name and data value and ect. When trigger is set to be 'axis', there may be data from multiple series.
	// In this time, series index can be refereed as {a0}, {a1}, or {a2}.
	// {a}, {b}, {c}, {d} have different meanings for different series types:
	//
	// * Line (area) charts, bar (column) charts, K charts: {a} for series name,
	//   {b} for category name, {c} for data value, {d} for none;
	// * Scatter (bubble) charts: {a} for series name, {b} for data name, {c} for data value, {d} for none;
	// * Map: {a} for series name, {b} for area name, {c} for merging data, {d} for none;
	// * Pie charts, gauge charts, funnel charts: {a} for series name, {b} for data item name,
	//   {c} for data value, {d} for percentage.
	//
	// 2. Callback function
	// The format of callback function:
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// The first parameter params is the data that the formatter needs. Its format is shown as follows:
	// {
	//    componentType: 'series',
	//    // Series type
	//    seriesType: string,
	//    // Series index in option.series
	//    seriesIndex: number,
	//    // Series name
	//    seriesName: string,
	//    // Data name, or category name
	//    name: string,
	//    // Data index in input data array
	//    dataIndex: number,
	//    // Original data as input
	//    data: Object,
	//    // Value of data. In most series it is the same as data.
	//    // But in some series it is some part of the data (e.g., in map, radar)
	//    value: number|Array|Object,
	//    // encoding info of coordinate system
	//    // Key: coord, like ('x' 'y' 'radius' 'angle')
	//    // value: Must be an array, not null/undefined. Contain dimension indices, like:
	//    // {
	//    //     x: [2] // values on dimension index 2 are mapped to x axis.
	//    //     y: [0] // values on dimension index 0 are mapped to y axis.
	//    // }
	//    encode: Object,
	//    // dimension names list
	//    dimensionNames: Array<String>,
	//    // data dimension index, for example 0 or 1 or 2 ...
	//    // Only work in `radar` series.
	//    dimensionIndex: number,
	//    // Color of data
	//    color: string,
	//
	//    // the percentage of pie chart
	//    percent: number,
	// }
	Formatter string `json:"formatter,omitempty"`
}

type TriggerType string

const (
	TriggerTypeItem TriggerType = "item"
	TriggerTypeAxis TriggerType = "axis"
	TriggerTypeNone TriggerType = "none"
)

// AxisPointer is a tool for displaying reference line and axis value under mouse pointer.
// https://echarts.apache.org/en/option.html#axisPointer
type AxisPointer struct {
	Show bool `json:"show,omitempty"`

	// Indicator type.
	Type PointerType `json:"type,omitempty"`

	// 	Whether snap to point automatically. The default value is auto determined.
	// This feature usually makes sense in value axis and time axis, where tiny points can be seeked automatically.
	Snap bool `json:"snap,omitempty"`

	Link []interface{} `json:"link,omitempty"`
}

type PointerType string

const (
	// line indicator.
	PointerTypeLine PointerType = "line"
	// shadow crosshair indicator.
	PointerTypeShadow PointerType = "shadow"
	// no indicator displayed.
	PointerTypeNone PointerType = "none"
	// crosshair indicator, which is actually the shortcut of enable two axisPointers of two orthometric axes.
	PointerTypeCross PointerType = "cross"
)
