package series

import "github.com/kva3umoda/goecharts/model"

type Series interface {
	Index() int
	Type() model.SeriesType
	Build() *model.Series
}

type BaseSeries struct {
	index int
	model *model.Series
}

func newBaseSeries(index int, series *model.Series) *BaseSeries {
	return &BaseSeries{
		index: index,
		model: series,
	}
}

func (base *BaseSeries) Index() int {
	return base.index
}

func (base *BaseSeries) Type() model.SeriesType {
	return base.model.Type
}

func (base *BaseSeries) Build() *model.Series {
	return base.model
}
