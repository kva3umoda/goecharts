package series

import "github.com/kva3umoda/goecharts/model"

var _ Series = (*CandleStick)(nil)

type CandleStick struct {
	series     *model.Series
	markLines  *MarkLines
	markPoints *MarkPoints
}

func NewCandleStick(name string) *CandleStick {
	cs := &CandleStick{
		series: &model.Series{
			Name:             name,
			Type:             model.SeriesTypeCandlestick,
			CoordinateSystem: model.CoordinateSystemCartesian2d,
			Encode:           &model.Encode{},
		},
	}

	return cs
}

func (s *CandleStick) Type() model.SeriesType {
	return s.series.Type
}

func (s *CandleStick) Build() *model.Series {
	return s.series
}

func (s *CandleStick) ItemStyleUpColor(upColor string) *CandleStick {
	if s.series.ItemStyle == nil {
		s.series.ItemStyle = &model.ItemStyle{}
	}
	s.series.ItemStyle.Color = upColor

	return s
}

func (s *CandleStick) ItemStyleUpBorderColor(upBorderColor string) *CandleStick {
	if s.series.ItemStyle == nil {
		s.series.ItemStyle = &model.ItemStyle{}
	}
	s.series.ItemStyle.BorderColor = upBorderColor

	return s
}

func (s *CandleStick) ItemStyleDownColor(downColor string) *CandleStick {
	if s.series.ItemStyle == nil {
		s.series.ItemStyle = &model.ItemStyle{}
	}

	s.series.ItemStyle.Color0 = downColor

	return s
}

func (s *CandleStick) ItemStyleDownBorderColor(downBorderColor string) *CandleStick {
	if s.series.ItemStyle == nil {
		s.series.ItemStyle = &model.ItemStyle{}
	}

	s.series.ItemStyle.BorderColor0 = downBorderColor

	return s
}

func (s *CandleStick) ItemStyleOpacity(opacity float32) *CandleStick {
	if s.series.ItemStyle == nil {
		s.series.ItemStyle = &model.ItemStyle{}
	}

	s.series.ItemStyle.Opacity = opacity

	return s
}

func (s *CandleStick) Dimension(date, close, open, high, low string) *CandleStick {
	s.series.Encode.X = []string{date}
	s.series.Encode.Y = []string{close, open, high, low}

	return s
}

func (s *CandleStick) MarkLines() *MarkLines {
	if s.markLines != nil {
		return s.markLines
	}

	s.markLines = newMarkLines(s.series)

	return s.markLines
}

func (s *CandleStick) MarkPoints() *MarkPoints {
	if s.markPoints != nil {
		return s.markPoints
	}

	s.markPoints = newMarkPoints(s.series)

	return s.markPoints
}
