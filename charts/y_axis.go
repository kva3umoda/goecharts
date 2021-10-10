package charts

import "github.com/kva3umoda/goecharts/model"

type YAxis struct {
	axis *model.YAxis
}

func NewYAxis(tp model.AxisType) *YAxis {
	ax := &YAxis{
		axis: &model.YAxis{
			Type:        tp,
			BoundaryGap: false,
		},
	}
	return ax
}

// It specifies whether not to contain zero position of axis compulsively.
// When it is set to be true, the axis may not contain zero position, which is useful in the scatter chart for both value axes.
func (axis *YAxis) ScaleEnabled() *YAxis {
	axis.axis.Scale = true

	return axis
}

func (axis *YAxis) ScaleDisabled() *YAxis {
	axis.axis.Scale = false

	return axis
}

// Split area of axis in grid area, not shown by default.
func (axis *YAxis) SplitAreaEnabled() *YAxis {
	if axis.axis.SplitArea == nil {
		axis.axis.SplitArea = &model.SplitArea{}
	}

	axis.axis.SplitArea.Show = true

	return axis
}

func (axis *YAxis) SplitAreaDisabled() *YAxis {
	if axis.axis.SplitArea != nil {
		axis.axis.SplitArea.Show = false
	}

	return axis
}

// Split line of axis in grid area.
func (axis *YAxis) SplitLineEnabled() *YAxis {
	if axis.axis.SplitLine == nil {
		axis.axis.SplitLine = &model.SplitLine{}
	}

	axis.axis.SplitLine.Show = true

	return axis
}

func (axis *YAxis) SplitLineDisabled() *YAxis {
	if axis.axis.SplitLine != nil {
		axis.axis.SplitLine.Show = false
	}

	return axis
}

// The minimun value of axis.
func (axis *YAxis) Min(min float32) *YAxis {
	axis.axis.Min = min
	return axis
}

func (axis *YAxis) MinSetDataMin() *YAxis {
	axis.axis.Min = model.AxisDataMin
	return axis
}

// The maximum value of axis.
func (axis *YAxis) Max(max float32) *YAxis {
	axis.axis.Max = max
	return axis
}

func (axis *YAxis) MaxSetDataMax() *YAxis {
	axis.axis.Max = model.AxisDataMax
	return axis
}

// The boundary gap on both sides of a coordinate axis. The setting and behavior of category axes and non-category axes are different.
func (axis *YAxis) BoundaryGap(value bool) *YAxis {
	axis.axis.BoundaryGap = value
	return axis
}

func (axis *YAxis) Build() *model.YAxis {
	return axis.axis
}
