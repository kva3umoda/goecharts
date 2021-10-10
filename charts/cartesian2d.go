package charts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/series"
)

type Cartesian2D struct {
	*BaseChart

	dataZoom []interface{}
	xAxis    *XAxis
	yAxis    *YAxis
	tooltip  *model.Tooltip
}

func NewCartesian2D(title string) *Cartesian2D {
	ch := &Cartesian2D{
		BaseChart: NewBaseChart(title),
	}

	return ch
}

func (chart *Cartesian2D) XAxis(tp model.AxisType) *XAxis {
	chart.xAxis = NewXAxis(tp)
	return chart.xAxis
}

func (chart *Cartesian2D) YAxis(tp model.AxisType) *YAxis {
	chart.yAxis = NewYAxis(tp)
	return chart.yAxis
}

func (chart *Cartesian2D) AddSeriesCandlestick(name string) *series.CandleStick {
	chart.legend.Data = append(chart.legend.Data, name)
	series := series.NewCandleStick(name)
	chart.series = append(chart.series, series)

	return series
}

func (chart *Cartesian2D) AddSeriesLine(name string) *series.Line {
	chart.legend.Data = append(chart.legend.Data, name)
	series := series.NewLine(name, model.CoordinateSystemCartesian2d)
	chart.series = append(chart.series, series)

	return series
}

func (chart *Cartesian2D) DataZoomXAxis(tp model.DataZoomType, startPercent, endPercent float32) *Cartesian2D {
	var dataZoom interface{}
	switch tp {
	case model.DataZoomTypeInside:
		dataZoom = &model.DataZoomInside{
			Type:       tp,
			XAxisIndex: []int{0},
			Start:      startPercent,
			End:        endPercent,
		}
	case model.DataZoomTypeSlider:
		dataZoom = &model.DataZoomSlider{
			Type:       tp,
			XAxisIndex: []int{0},
			Start:      startPercent,
			End:        endPercent,
		}
	}
	chart.dataZoom = append(chart.dataZoom, dataZoom)

	return chart
}

func (chart *Cartesian2D) DataZoomYAxis(tp model.DataZoomType, startPercent, endPercent float32) *Cartesian2D {
	var dataZoom interface{}
	switch tp {
	case model.DataZoomTypeInside:
		dataZoom = &model.DataZoomInside{
			Type:  tp,
			Start: startPercent,
			End:   endPercent,
		}
	case model.DataZoomTypeSlider:
		dataZoom = &model.DataZoomSlider{
			Type:  tp,
			Start: startPercent,
			End:   endPercent,
		}
	}
	chart.dataZoom = append(chart.dataZoom, dataZoom)

	return chart
}

func (chart *Cartesian2D) ToolTip(trigger model.TriggerType, axisPointerType model.PointerType) *Cartesian2D {
	chart.tooltip = &model.Tooltip{
		Trigger: trigger,
		AxisPointer: &model.AxisPointer{
			Type: axisPointerType,
		},
	}

	return chart
}

func (chart *Cartesian2D) Build() *model.Option {
	opt := &model.Option{}
	opt.Title = chart.title
	if chart.xAxis != nil {
		opt.XAxis = chart.xAxis.Build()
	}

	if chart.yAxis != nil {
		opt.YAxis = chart.yAxis.Build()
	}

	opt.DataZoom = chart.dataZoom
	opt.Legend = chart.legend
	opt.Tooltip = chart.tooltip
	opt.Dataset = chart.dataset

	for _, series := range chart.series {
		opt.Series = append(opt.Series, series.Build())
	}

	return opt
}

// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (chart *Cartesian2D) JSONNotEscaped() template.HTML {
	opt := chart.Build()
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "   ")

	if err := enc.Encode(opt); err != nil {
		panic(err)
	}

	return template.HTML(buff.String())
}
