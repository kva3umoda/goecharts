package series

import (
	"github.com/kva3umoda/goecharts/model"
)

var _ Series = (*Line)(nil)

type Line struct {
	series     *model.Series
	markLines  *MarkLines
	markPoints *MarkPoints
}

func NewLine(name string, coordinateSystem model.CoordinateSystem) *Line {
	cs := &Line{
		series: &model.Series{
			Name:             name,
			Type:             model.SeriesTypeLine,
			CoordinateSystem: coordinateSystem, /*polar*/
			Encode:           &model.Encode{},
			AreaStyle:        nil},
	}

	return cs
}

func (s *Line) Type() model.SeriesType {
	return s.series.Type
}

func (s *Line) Build() *model.Series {
	return s.series
}

func (s *Line) SmoothEnabled() *Line {
	s.series.Smooth = true

	return s
}

func (s *Line) SmoothDisabled() *Line {
	s.series.Smooth = false

	return s
}

func (s *Line) LineStyleColor(color string) *Line {
	if s.series.LineStyle == nil {
		s.series.LineStyle = &model.LineStyle{}
	}

	s.series.LineStyle.Color = color

	return s
}

func (s *Line) LineStyleWidth(width float32) *Line {
	if s.series.LineStyle == nil {
		s.series.LineStyle = &model.LineStyle{}
	}
	s.series.LineStyle.Width = width

	return s
}

func (s *Line) LineStyleType(tp model.LineType) *Line {
	if s.series.LineStyle == nil {
		s.series.LineStyle = &model.LineStyle{}
	}
	s.series.LineStyle.Type = tp

	return s
}

func (s *Line) LineStyleOpacity(opacity float32) *Line {
	if s.series.LineStyle == nil {
		s.series.LineStyle = &model.LineStyle{}
	}
	s.series.LineStyle.Opacity = opacity

	return s
}

func (s *Line) Dimension(x, y string) *Line {
	s.series.Encode.X = []string{x}
	s.series.Encode.Y = []string{y}

	return s
}

// Area properties
func (s *Line) AreaEnabled() *Line {
	if s.series.AreaStyle == nil {
		s.series.AreaStyle = &model.AreaStyle{}
	}

	return s
}

func (s *Line) AreaDisabled() *Line {
	s.series.AreaStyle = nil
	return s
}

// Stack properties
func (s *Line) StackEnabled() *Line {
	s.series.Stack = "Total"

	return s
}

func (s *Line) StackDisabled() *Line {
	s.series.Stack = ""

	return s
}

// Emphasis properties
func (s *Line) EmphasisEnabled(scale bool, focus model.Focus, blurScope model.BlurScope) *Line {
	if s.series.Emphasis == nil {
		s.series.Emphasis = &model.Emphasis{}
	}
	s.series.Emphasis.Scale = scale
	s.series.Emphasis.Focus = focus
	s.series.Emphasis.BlurScope = blurScope
	return s
}

func (s *Line) EmphasisDisabled() *Line {
	s.series.Emphasis = nil
	return s
}

// Label properties
func (s *Line) LabelEnabled(position model.Position) *Line {
	if s.series.Label == nil {
		s.series.Label = &model.Label{}
	}
	s.series.Label.Show = true
	s.series.Label.Position = position

	return s
}

func (s *Line) LabelDisabled() *Line {
	if s.series.Label != nil {
		s.series.Label.Show = false
	}

	return s
}

func (s *Line) MarkLines() *MarkLines {
	if s.markLines != nil {
		return s.markLines
	}

	s.markLines = newMarkLines(s.series)

	return s.markLines
}

func (s *Line) MarkPoints() *MarkPoints {
	if s.markPoints != nil {
		return s.markPoints
	}

	s.markPoints = newMarkPoints(s.series)

	return s.markPoints
}
