package axis

import "github.com/kva3umoda/goecharts/model"

type XYAxis struct {
	index int
	model *model.XYAxis
}

func NewXYAxis(index int, tp model.AxisType, axis *model.XYAxis) *XYAxis {
	ax := &XYAxis{
		index: index,
		model: axis,
	}

	ax.model.BoundaryGap = false
	ax.model.Type = tp
	ax.model.Min = model.AxisDataMin
	ax.model.Max = model.AxisDataMax

	return ax
}

func (axis *XYAxis) Index() int {
	return axis.index
}

// It specifies whether not to contain zero position of axis compulsively.
// When it is set to be true, the axis may not contain zero position, which is useful in the scatter chart for both value axes.
func (axis *XYAxis) Scale(enable bool) *XYAxis {
	axis.model.Scale = enable

	return axis
}

// The minimun value of axis.
func (axis *XYAxis) Min(min interface{}) *XYAxis {
	axis.model.Min = min

	return axis
}

// The maximum value of axis.
func (axis *XYAxis) Max(max interface{}) *XYAxis {
	axis.model.Max = max
	return axis
}

// The boundary gap on both sides of a coordinate axis. The setting and behavior of category axes and non-category axes are different.
func (axis *XYAxis) BoundaryGap(enable bool) *XYAxis {
	axis.model.BoundaryGap = true
	return axis
}

// Split line of axis in grid area.
func (axis *XYAxis) SplitLine(enable bool) *XYAxis {
	if axis.model.SplitLine == nil {
		axis.model.SplitLine = &model.SplitLine{}
	}

	axis.model.SplitLine.Show = enable

	return axis
}

// Split area of axis in grid area, not shown by default.
func (axis *XYAxis) SplitArea(enable bool) *XYAxis {
	if axis.model.SplitArea == nil {
		axis.model.SplitArea = &model.SplitArea{}
	}

	axis.model.SplitArea.Show = enable

	return axis
}
