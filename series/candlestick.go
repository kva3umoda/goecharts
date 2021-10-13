package series

import (
	"github.com/kva3umoda/goecharts/model"
)

var _ Series = (*CandleStick)(nil)

type CandleStick struct {
	*BaseSeries

	markLines  *MarkLines
	markPoints *MarkPoints
}

func NewCandleStick(index int, dimDate, dimOpen, dimClose, dimHigh, dimLow string, series *model.Series) *CandleStick {

	series.Type = model.SeriesTypeCandlestick
	series.Encode = &model.Encode{
		X: []string{dimDate},
		Y: []string{dimClose, dimOpen, dimHigh, dimLow},
	}

	cs := &CandleStick{
		BaseSeries: newBaseSeries(index, series),
	}

	return cs
}

func (s *CandleStick) Index() int {
	return s.index
}

func (s *CandleStick) Type() model.SeriesType {
	return s.model.Type
}

func (s *CandleStick) Build() *model.Series {
	return s.model
}

func (s *CandleStick) ItemStyleUpColor(upColor string) *CandleStick {
	if s.model.ItemStyle == nil {
		s.model.ItemStyle = &model.ItemStyle{}
	}
	s.model.ItemStyle.Color = upColor

	return s
}

func (s *CandleStick) ItemStyleUpBorderColor(upBorderColor string) *CandleStick {
	if s.model.ItemStyle == nil {
		s.model.ItemStyle = &model.ItemStyle{}
	}
	s.model.ItemStyle.BorderColor = upBorderColor

	return s
}

func (s *CandleStick) ItemStyleDownColor(downColor string) *CandleStick {
	if s.model.ItemStyle == nil {
		s.model.ItemStyle = &model.ItemStyle{}
	}

	s.model.ItemStyle.Color0 = downColor

	return s
}

func (s *CandleStick) ItemStyleDownBorderColor(downBorderColor string) *CandleStick {
	if s.model.ItemStyle == nil {
		s.model.ItemStyle = &model.ItemStyle{}
	}

	s.model.ItemStyle.BorderColor0 = downBorderColor

	return s
}

func (s *CandleStick) ItemStyleOpacity(opacity float32) *CandleStick {
	if s.model.ItemStyle == nil {
		s.model.ItemStyle = &model.ItemStyle{}
	}

	s.model.ItemStyle.Opacity = opacity

	return s
}

func (s *CandleStick) DatasetDimension(date, close, open, high, low string) *CandleStick {
	s.model.Encode.X = []string{date}
	s.model.Encode.Y = []string{close, open, high, low}

	return s
}

func (s *CandleStick) MarkLines() *MarkLines {
	if s.markLines != nil {
		return s.markLines
	}

	s.markLines = newMarkLines(s.model)

	return s.markLines
}

func (s *CandleStick) MarkPoints() *MarkPoints {
	if s.markPoints != nil {
		return s.markPoints
	}

	s.markPoints = newMarkPoints(s.model)

	return s.markPoints
}
