package goecharts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/kva3umoda/goecharts/charts"
	"github.com/kva3umoda/goecharts/model"
)

type Charts struct {
	dataset *Dataset
	model   *model.Option
}

func NewCharts(title string) *Charts {
	chrt := &Charts{
		model: &model.Option{
			Legend: &model.Legend{
				Show: true,
				Type: model.LegendTypePain,
			},
			Title: &model.Title{
				Text: title,
			},

			Toolbox: &model.Toolbox{
				Show:   true,
				Orient: model.OrientHorizontal,
				Feature: &model.ToolBoxFeature{
					SaveAsImage: &model.ToolBoxFeatureSaveAsImage{
						Show: true,
					},
					DataView: &model.ToolBoxFeatureDataView{
						Show: true,
					},
					DataZoom: &model.ToolBoxFeatureDataZoom{
						Show: true,
					},
					Restore: &model.ToolBoxFeatureRestore{
						Show: true,
					},
					MagicType: &model.ToolBoxFeatureMagicType{
						Show: true,
						Type: []string{"line", "bar", "stack"},
					},
					Brush: &model.ToolBoxFeatureBrush{
						Show: false,
						Type: []string{"rect", "polygon", "lineX", "lineY", "keep", "clear"},
					},
				},
			},

			Series:    make([]*model.Series, 0),
			Grids:     make([]*model.Grid, 0),
			VisualMap: make([]model.VisualMap, 0),
		},
	}
	dataset := &model.Dataset{}
	chrt.model.Dataset = dataset
	chrt.dataset = newDataset(1, dataset)

	return chrt
}

func (chart *Charts) ToolTip(trigger model.TriggerType, axisPointerType model.PointerType) *Charts {
	chart.model.Tooltip = &model.Tooltip{
		Trigger: trigger,
		AxisPointer: &model.AxisPointer{
			Type: axisPointerType,
		},
	}

	return chart
}

func (chart *Charts) DataZoomXAxis(tp model.DataZoomType, startPercent, endPercent float32, xAxisIndexes ...int) *Charts {
	var dataZoom interface{}
	switch tp {
	case model.DataZoomTypeInside:
		dataZoom = &model.DataZoomInside{
			Type:       tp,
			XAxisIndex: xAxisIndexes,
			Start:      startPercent,
			End:        endPercent,
		}
	case model.DataZoomTypeSlider:
		dataZoom = &model.DataZoomSlider{
			Type:       tp,
			XAxisIndex: xAxisIndexes,
			Start:      startPercent,
			End:        endPercent,
		}
	}
	chart.model.DataZoom = append(chart.model.DataZoom, dataZoom)

	return chart
}

func (chart *Charts) DataZoomYAxis(tp model.DataZoomType, startPercent, endPercent float32, yAxisIndexes ...int) *Charts {
	var dataZoom interface{}
	switch tp {
	case model.DataZoomTypeInside:
		dataZoom = &model.DataZoomInside{
			Type:       tp,
			YAxisIndex: yAxisIndexes,
			Start:      startPercent,
			End:        endPercent,
		}
	case model.DataZoomTypeSlider:
		dataZoom = &model.DataZoomSlider{
			Type:       tp,
			YAxisIndex: yAxisIndexes,
			Start:      startPercent,
			End:        endPercent,
		}
	}
	chart.model.DataZoom = append(chart.model.DataZoom, dataZoom)

	return chart
}

func (chart *Charts) Dataset() *Dataset {
	return chart.dataset
}

func (chart *Charts) AddChart2D(name string, xAxisType, yAxisType model.AxisType) *charts.Chart2D {
	ch := charts.NewChart2D(name, xAxisType, yAxisType, chart.model)

	return ch
}

func (chart *Charts) Build() *model.Option {

	return chart.model
}

// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (chart *Charts) JSONNotEscaped() template.HTML {
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
