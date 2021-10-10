package series

import "github.com/kva3umoda/goecharts/model"

type MarkLines struct {
	markLines *model.MarkLines
}

func newMarkLines(series *model.Series) *MarkLines {
	ml := &MarkLines{
		markLines: &model.MarkLines{
			Silent: false,
			Symbol: model.SymbolNone,
		},
	}
	series.MarkLine = ml.markLines
	return ml
}

func (ml *MarkLines) AddVertical(name string, xvalue interface{}) *MarkLines {
	ml.markLines.Data = append(ml.markLines.Data,
		model.MarkLineData{
			Name:  name,
			XAxis: xvalue,
		},
	)

	return ml
}

func (ml *MarkLines) AddHorizontal(name string, yvalue interface{}) *MarkLines {
	ml.markLines.Data = append(ml.markLines.Data,
		model.MarkLineData{
			Name:  name,
			YAxis: yvalue,
		},
	)

	return ml
}

func (ml *MarkLines) AddSpecial(name string, tp model.MarkLineType, dimension ...string) *MarkLines {
	d := model.MarkLineData{
		Name: name,
		Type: tp,
	}
	if len(dimension) >= 1 {
		d.ValueDim = dimension[0]
	}

	ml.markLines.Data = append(ml.markLines.Data, d)

	return ml
}

func (ml *MarkLines) AddByPoints(name string,
	x1, y1 interface{}, symbol1 model.Symbol,
	x2, y2 interface{}, symbol2 model.Symbol) *MarkLines {
	ml.markLines.Data = append(ml.markLines.Data,
		[]model.MarkLineData{
			{
				Name:   name,
				Coord:  []interface{}{x1, y1},
				Symbol: symbol1,
			},
			{
				Coord:  []interface{}{x2, y2},
				Symbol: symbol2,
			},
		},
	)

	return ml
}
