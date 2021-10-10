package charts

import "github.com/kva3umoda/goecharts/model"

type XAxis struct {
	axis *model.XAxis
}

func NewXAxis(tp model.AxisType) *XAxis {
	ax := &XAxis{
		axis: &model.XAxis{
			Type:        tp,
			BoundaryGap: false,
		},
	}

	return ax
}


// It specifies whether not to contain zero position of axis compulsively.
// When it is set to be true, the axis may not contain zero position, which is useful in the scatter chart for both value axes.
func (axis *XAxis) ScaleEnabled() *XAxis {
	axis.axis.Scale = true

	return axis
}

func (axis *XAxis) ScaleDisabled() *XAxis {
	axis.axis.Scale = false

	return axis
}


// The minimun value of axis.
func (axis *XAxis) Min(min float32) *XAxis {
	axis.axis.Min = min
	return axis
}

func (axis *XAxis) MinSetDataMin() *XAxis {
	axis.axis.Min = model.AxisDataMin
	return axis
}

// The maximum value of axis.
func (axis *XAxis) Max(max float32) *XAxis {
	axis.axis.Max = max
	return axis
}

func (axis *XAxis) MaxSetDataMax() *XAxis {
	axis.axis.Max = model.AxisDataMax
	return axis
}

func (axis *XAxis) Scale(scale bool) *XAxis {
	axis.axis.Scale = scale
	return axis
}

// The boundary gap on both sides of a coordinate axis. The setting and behavior of category axes and non-category axes are different.
func (axis *XAxis) BoundaryGap(value bool) *XAxis {
	axis.axis.BoundaryGap = value
	return axis
}


func (axis *XAxis) SplitLine(show bool) *XAxis {
	if axis.axis.SplitLine == nil {
		axis.axis.SplitLine = &model.SplitLine{}
	}

	axis.axis.SplitLine.Show = show

	return axis
}

func (axis *XAxis) Build() *model.XAxis {
	return axis.axis
}
