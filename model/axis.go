package model

const (
	AxisDataMin = "dataMin"
	AxisDataMax = "dataMax"
)

type Position string

const (
	PositionTop               Position = "top"
	PositionLeft              Position = "left"
	PositionRight             Position = "right"
	PositionBottom            Position = "bottom"
	PositionInside            Position = "inside"
	PositionInsideLeft        Position = "insideLeft"
	PositionInsideRight       Position = "insideRight"
	PositionInsideTop         Position = "insideTop"
	PositionInsideBottom      Position = "insideBottom"
	PositionInsideTopLeft     Position = "insideTopLeft"
	PositionInsideBottomLeft  Position = "insideBottomLeft"
	PositionInsideTopRight    Position = "insideTopRight"
	PositionInsideBottomRight Position = "insideBottomRight"
)

type AxisType string

const (
	// Numerical axis, suitable for continuous data.
	AxisTypeValue AxisType = "value"

	// Category axis, suitable for discrete category data. Category data can be auto retrieved from series.data or dataset.source, or can be specified via xAxis.data.
	AxisTypeCategory AxisType = "category"

	// Time axis, suitable for continuous time series data. As compared to value axis, it has a better formatting for time and a different tick calculation method. For example, it decides to use month, week, day or hour for tick based on the range of span.
	AxisTypeTime AxisType = "time"

	// Log axis, suitable for log data.
	AxisTypeLog AxisType = "log"

	AxisTypeEmpty AxisType = ""
)

type Location string

const (
	LocationStart  Location = "start"
	LocationMiddle Location = "middle"
	LocationCenter Location = "center"
	LocationEnd    Location = "end"
)

// SplitArea is the option set for a split area.
type SplitArea struct {
	// Set this to true to show the splitArea.
	Show bool `json:"show"`

	// Split area style.
	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show bool `json:"show"`

	// Split line style.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// AxisLabel settings related to axis label.
// https://echarts.apache.org/en/option.html#xAxis.axisLabel
type AxisLabel struct {
	// Set this to false to prevent the axis label from appearing.
	Show bool `json:"show"`

	// Interval of AxisXY label, which is available in category axis.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label.
	// And if it is set to be 2, it means labels are shown once after two labels, and so on.
	Interval string `json:"interval,omitempty"`

	// Set this to true so the axis labels face the inside direction.
	Inside bool `json:"inside,omitempty"`

	// Rotation degree of axis label, which is especially useful when there is no enough space for category axis.
	// Rotation degree is from -90 to 90.
	Rotate float64 `json:"rotate,omitempty"`

	// The margin between the axis label and the axis line.
	Margin float64 `json:"margin,omitempty"`

	// Formatter of axis label, which supports string template and callback function.
	//
	// Example:
	//
	// Use string template; template variable is the default label of axis {value}
	// formatter: '{value} kg'
	//
	// Use callback function; function parameters are axis index
	//
	//
	//  formatter: function (value, index) {
	//    // Formatted to be month/day; display year only in the first label
	//    var date = new Date(value);
	//    var texts = [(date.getMonth() + 1), date.getDate()];
	//    if (idx === 0) {
	//        texts.unshift(date.getYear());
	//    }
	//    return texts.join('/');
	// }
	Formatter string `json:"formatter,omitempty"`

	ShowMinLabel bool `json:"showMinLabel"`
	ShowMaxLabel bool `json:"showMaxLabel"`

	// Color of axis label is set to be axisLine.lineStyle.color by default. Callback function is supported,
	// in the following format:
	//
	// (val: string) => Color
	// Parameter is the text of label, and return value is the color. See the following example:
	//
	// textStyle: {
	//    color: function (value, index) {
	//        return value >= 0 ? 'green' : 'red';
	//    }
	// }
	Color string `json:"color,omitempty"`

	// axis label font style
	FontStyle string `json:"fontStyle,omitempty"`
	// axis label font weight
	FontWeight string `json:"fontWeight,omitempty"`
	// axis label font family
	FontFamily string `json:"fontFamily,omitempty"`
	// axis label font size
	FontSize string `json:"fontSize,omitempty"`
	// Horizontal alignment of axis label
	Align string `json:"align,omitempty"`
	// Vertical alignment of axis label
	VerticalAlign string `json:"verticalAlign,omitempty"`
	// Line height of the axis label
	LineHeight string `json:"lineHeight,omitempty"`
}

type AxisTick struct {
	// Set this to false to prevent the axis tick from showing.
	Show bool `json:"show"`

	// interval of axisTick, which is available in category axis. is set to be the same as axisLabel.interval by default.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label. And if it is set to be 2, it means labels are shown once after two labels, and so on.
	// On the other hand, you can control by callback function, whose format is shown below:
	// (index:number, value: string) => boolean
	// The first parameter is index of category, and the second parameter is the name of category. The return values decides whether to display label.
	Interval string `json:"interval,omitempty"`
}

type MinorTick struct {
	// If show minor ticks.
	Show bool `json:"show"`

	// Number of interval splited by minor ticks.
	SplitNumber int `json:"splitNumber,omitempty"`

	// Length of minor ticks lines。
	Length int `json:"length,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// AxisLine controls settings related to axis line.
// https://echarts.apache.org/en/option.html#yAxis.axisLine
type AxisLine struct {
	// Set this to false to prevent the axis line from showing.
	Show bool `json:"show"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	OnZero bool `json:"onZero,omitempty"`

	// When multiple axes exists, this option can be used to specify which axis can be "onZero" to.
	OnZeroAxisIndex int `json:"onZeroAxisIndex,omitempty"`

	// Symbol of the two ends of the axis. It could be a string, representing the same symbol for two ends; or an array
	// with two string elements, representing the two ends separately. It's set to be 'none' by default, meaning no
	//arrow for either end. If it is set to be 'arrow', there shall be two arrows. If there should only one arrow
	//at the end, it should set to be ['none', 'arrow'].
	Symbol string `json:"symbol,omitempty"`

	// Size of the arrows at two ends. The first is the width perpendicular to the axis, the next is the width parallel to the axis.
	SymbolSize []float64 `json:"symbolSize,omitempty"`

	// Arrow offset of axis. If is array, the first number is the offset of the arrow at the beginning, and the second
	// number is the offset of the arrow at the end. If is number, it means the arrows have the same offset.
	SymbolOffset []float64 `json:"symbolOffset,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}
