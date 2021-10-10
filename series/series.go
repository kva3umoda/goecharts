package series

import "github.com/kva3umoda/goecharts/model"

type Series interface {
	Type() model.SeriesType
	Build() *model.Series
}


