package series

import "github.com/kva3umoda/goecharts/model"

type MarkPoints struct {
	markPoints *model.MarkPoints
}

func newMarkPoints(series *model.Series) *MarkPoints {
	ml := &MarkPoints{
		markPoints: &model.MarkPoints{
			Silent: false,
			Symbol: model.SymbolNone,
		},
	}
	series.MarkPoint = ml.markPoints
	return ml
}

func (ml *MarkPoints) AddSpecial(name string, symbol model.Symbol, tp model.MarkLineType, dimension string) *MarkPoints {
	ml.markPoints.Data = append(
		ml.markPoints.Data,
		model.MarkLineData{
			Name:     name,
			Type:     tp,
			ValueDim: dimension,
			Symbol:   symbol,
		},
	)

	return ml
}

func (ml *MarkPoints) AddByPoints(name string, symbol model.Symbol, x, y interface{}) *MarkPoints {
	ml.markPoints.Data = append(ml.markPoints.Data,
		model.MarkPointData{
			Name:   name,
			Coord:  []interface{}{x, y},
			Symbol: symbol,
		},
	)

	return ml
}
