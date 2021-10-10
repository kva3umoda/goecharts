package model

type AngleAxis struct {
	// Component ID, not specified by default. If specified, it can be used to refer the component in option or API.
	ID string `json:"id,omitempty"`

	// Index of radial axis in polor coordinate. It's the first axis by default.
	PolarIndex int `json:"polarIndex,omitempty"`

	// Starting angle of axis. 90 degrees by default, standing for top position of center. 0 degree stands for right position of center.
	StartAngle float64 `json:"startAngle,omitempty"`

	// Whether the positive position of axis is clockwise. True for clockwise by default.
	Clockwise bool `json:"clockwise,omitempty"`

	// Type of axis.
	Type AxisType `json:"type,omitempty"`

	// The boundary gap on both sides of a coordinate axis. The setting and behavior of category axes and non-category axes are different.
	BoundaryGap bool `json:"boundaryGap,omitempty"`

	// The minimun value of axis.
	Min float64 `json:"min,omitempty"`

	// The maximum value of axis.
	Max float64 `json:"max,omitempty"`

	//  It is available only in numerical axis, i.e., type: 'value'.
	Scale bool `json:"scale,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a recommendation, and the true segments may be adjusted based on readability.
	SplitNumber int `json:"splitNumber,omitempty"`

	// Minimum gap between split lines.
	MinInterval float64 `json:"minInterval,omitempty"`

	// Maximum gap between split lines.
	MaxInterval float64 `json:"maxInterval,omitempty"`

	// Compulsively set segmentation interval for axis.
	Interval float64 `json:"interval,omitempty"`

	// Base of logarithm, which is valid only for numeric axes with type: 'log'.
	LogBase float64 `json:"logBase,omitempty"`

	// Set this to true, to prevent interaction with the axis.
	Silent bool `json:"silent,omitempty"`

	// Set this to true to enable triggering events.
	TriggerEvent bool `json:"triggerEvent,omitempty"`

	// Settings related to axis line.
	AxisLine *AxisLine `json:"axisLine,omitempty"`

	// Settings related to axis tick.
	AxisTick *AxisTick `json:"axisTick,omitempty"`

	// Settings related minor ticks.
	MinorTick *MinorTick `json:"minorTick,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Minor split lines of axis in the grid area。It will align to the minorTick
	MinorSplitLine *SplitLine `json:"MinorSplitLine,omitempty"`

	// Split area of X axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// axisPointer settings on the axis.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}
