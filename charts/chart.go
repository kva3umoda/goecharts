package charts

import (
	"html/template"
	"math/rand"
	"strconv"
	"time"

	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/series"
)

type Chart interface {
	ChartID() string
	Build() *model.Option
	JSONNotEscaped() template.HTML
}

type BaseChart struct {
	id      string
	title   *model.Title
	legend  *model.Legend
	dataset *model.Dataset

	series []series.Series
}

func NewBaseChart(title string) *BaseChart {
	ch := &BaseChart{
		id:     strconv.Itoa(int(rand.Int63n(100000000000))),
		series: make([]series.Series, 0, 1),
		legend: &model.Legend{
			Type: model.LegendTypePain,
			Show: true,
		},

		dataset: &model.Dataset{
			SourceHeader: false,
			Source:       make([][]interface{}, 0),
		},
	}

	ch.title = &model.Title{
		Text: title,
	}
	return ch
}

func (chart *BaseChart) ChartID() string {
	return chart.id
}

func (chart *BaseChart) Dimension(name string, tp model.DataType) *BaseChart {
	chart.dataset.Dimensions = append(chart.dataset.Dimensions,
		model.Dimension{Name: name, Type: tp},
	)

	return chart
}

func (chart *BaseChart) DatasetRow(values ...interface{}) *BaseChart {
	row := make([]interface{}, 0, len(chart.dataset.Dimensions))
	for _, v := range values {
		row = append(row, v)
	}
	chart.dataset.Source = append(chart.dataset.Source, row)

	return chart
}

func (chart *BaseChart) DatasetColumnFloat(dimension string, values []float64) *BaseChart {
	chart.Dimension(dimension, model.DataFloat)
	for i, v := range values {
		if len(chart.dataset.Source) < len(values) {
			chart.dataset.Source = append(chart.dataset.Source,
				make([]interface{},
					len(chart.dataset.Dimensions)-1,
					len(chart.dataset.Dimensions)))
		}
		chart.dataset.Source[i] = append(chart.dataset.Source[i], v)
	}

	return chart
}

func (chart *BaseChart) DatasetColumnString(dimension string, values []string) *BaseChart {
	chart.Dimension(dimension, model.DataOrdinal)
	for i, v := range values {
		if len(chart.dataset.Source) < len(values) {
			chart.dataset.Source = append(chart.dataset.Source,
				make([]interface{},
					len(chart.dataset.Dimensions)-1,
					len(chart.dataset.Dimensions)))
		}
		chart.dataset.Source[i] = append(chart.dataset.Source[i], v)
	}

	return chart
}

func (chart *BaseChart) DatesetColumnInt(dimension string, values []int32) *BaseChart {
	chart.Dimension(dimension, model.DataInt)
	for i, v := range values {
		if len(chart.dataset.Source) < len(values) {
			chart.dataset.Source = append(chart.dataset.Source,
				make([]interface{},
					len(chart.dataset.Dimensions)-1,
					len(chart.dataset.Dimensions)))
		}
		chart.dataset.Source[i] = append(chart.dataset.Source[i], v)
	}

	return chart
}

func (chart *BaseChart) DatasetColumnUnixTime(dimension string, values []int64) *BaseChart {
	chart.Dimension(dimension, model.DataTime)
	for i, v := range values {
		if len(chart.dataset.Source) < len(values) {
			chart.dataset.Source = append(chart.dataset.Source,
				make([]interface{},
					len(chart.dataset.Dimensions)-1,
					len(chart.dataset.Dimensions)))
		}
		chart.dataset.Source[i] = append(chart.dataset.Source[i], v)
	}

	return chart
}

func (chart *BaseChart) DatasetColumnTime(dimension string, values []time.Time) *BaseChart {
	chart.Dimension(dimension, model.DataTime)
	for i, v := range values {
		if len(chart.dataset.Source) < len(values) {
			chart.dataset.Source = append(chart.dataset.Source,
				make([]interface{},
					len(chart.dataset.Dimensions)-1,
					len(chart.dataset.Dimensions)))
		}
		chart.dataset.Source[i] = append(chart.dataset.Source[i], v.Unix())
	}

	return chart
}
