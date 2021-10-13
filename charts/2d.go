package charts

import (
	"github.com/kva3umoda/goecharts/axis"
	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/series"
)

// The x axis in cartesian(rectangular) coordinate
type Chart2D struct {
	model  *model.Option
	name   string
	index  int
	xAxis  *axis.XYAxis
	yAxis  *axis.XYAxis
	series []series.Series
}

func NewChart2D(name string, xAxisType, yAxisType model.AxisType, option *model.Option) *Chart2D {
	chart := &Chart2D{
		model: option,
		name:  name,
	}
	// create grid
	grid := &model.Grid{}
	chart.model.Grids = append(chart.model.Grids, grid)
	chart.index = len(chart.model.Grids) - 1

	// create xAxis
	x := &model.XYAxis{
		GridIndex: chart.index,
	}
	chart.model.XAxis = append(chart.model.XAxis, x)
	chart.xAxis = axis.NewXYAxis(len(chart.model.XAxis)-1, xAxisType, x)

	// create yAxis
	y := &model.XYAxis{
		GridIndex: chart.index,
	}
	chart.model.YAxis = append(chart.model.YAxis, y)
	chart.yAxis = axis.NewXYAxis(len(chart.model.YAxis)-1, yAxisType, y)

	return chart
}

func (chart *Chart2D) Index() int {
	return chart.index
}

func (chart *Chart2D) XAxis() *axis.XYAxis {
	return chart.xAxis
}

func (chart *Chart2D) YAxis() *axis.XYAxis {
	return chart.yAxis
}

func (chart *Chart2D) AddSeriesCandlestick(name, dimDate, dimOpen, dimClose, dimHigh, dimLow string) *series.CandleStick {
	chart.model.Legend.Data = append(chart.model.Legend.Data, name)
	s := &model.Series{
		Name:             name,
		CoordinateSystem: model.CoordinateSystemCartesian2d,
		XAxisIndex:       chart.xAxis.Index(),
		YAxisIndex:       chart.yAxis.Index(),
	}
	chart.model.Series = append(chart.model.Series, s)

	series := series.NewCandleStick(len(chart.model.Series)-1, dimDate, dimOpen, dimClose, dimHigh, dimLow, s)
	chart.series = append(chart.series, series)

	return series
}

func (chart *Chart2D) AddSeriesLine(name, dimX, dimY string) *series.Line {
	chart.model.Legend.Data = append(chart.model.Legend.Data, name)
	s := &model.Series{
		Name:             name,
		CoordinateSystem: model.CoordinateSystemCartesian2d,
		XAxisIndex:       chart.xAxis.Index(),
		YAxisIndex:       chart.yAxis.Index(),
	}
	chart.model.Series = append(chart.model.Series, s)
	series := series.NewLine(len(chart.model.Series)-1, dimX, dimY, s)
	chart.series = append(chart.series, series)

	return series
}
