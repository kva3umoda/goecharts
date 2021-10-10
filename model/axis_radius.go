package model

type RadiusAxis struct {
	// Component ID, not specified by default. If specified, it can be used to refer the component in option or API.
	ID string `json:"id,omitempty"`

	// Index of radial axis in polor coordinate. It's the first axis by default.
	PolarIndex int `json:"polarIndex,omitempty"`

	// Type of axis.
	Type AxisType `json:"type,omitempty"`

	// Name of axis.
	Name string `json:"name,omitempty"`

	// Location of axis name.
	NameLocation Location `json:"nameLocation,omitempty"`

	// Text style of axis name.
	NameTextStyle TextStyle `json:"nameTextStyle,omitempty"`

	// Gap between axis name and axis line.
	NameGap float64 `json:"nameGap,omitempty"`

	// Rotation of axis name.
	NameRotate float64 `json:"nameRotate,omitempty"`

	// Set this to true to invert the axis. This is a new option available from Echarts 3 and newer.
	Inverse bool `json:"inverse,omitempty"`

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

	// axisPointer settings on the axis.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}
