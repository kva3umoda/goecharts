package series

import (
	"github.com/kva3umoda/goecharts/model"
)

var _ Series = (*Line)(nil)

type Line struct {
	*BaseSeries
	markLines  *MarkLines
	markPoints *MarkPoints
}

func NewLine(index int, dimX string, dimY string, series *model.Series) *Line {
	series.Type = model.SeriesTypeLine
	series.Encode = &model.Encode{
		X: []string{dimX},
		Y: []string{dimY},
	}

	cs := &Line{
		BaseSeries: newBaseSeries(index, series),
	}

	return cs
}

func (s *Line) Smooth(enable bool) *Line {
	s.model.Smooth = enable

	return s
}

func (s *Line) LineStyleColor(color string) *Line {
	if s.model.LineStyle == nil {
		s.model.LineStyle = &model.LineStyle{}
	}

	s.model.LineStyle.Color = color

	return s
}

func (s *Line) LineStyleWidth(width float32) *Line {
	if s.model.LineStyle == nil {
		s.model.LineStyle = &model.LineStyle{}
	}
	s.model.LineStyle.Width = width

	return s
}

func (s *Line) LineStyleType(tp model.LineType) *Line {
	if s.model.LineStyle == nil {
		s.model.LineStyle = &model.LineStyle{}
	}
	s.model.LineStyle.Type = tp

	return s
}

func (s *Line) LineStyleOpacity(opacity float32) *Line {
	if s.model.LineStyle == nil {
		s.model.LineStyle = &model.LineStyle{}
	}
	s.model.LineStyle.Opacity = opacity

	return s
}

// Area properties
func (s *Line) Area(enable bool) *Line {
	if enable {
		if s.model.AreaStyle == nil {
			s.model.AreaStyle = &model.AreaStyle{}
		}

	} else {
		s.model.AreaStyle = nil
	}

	return s
}

// Stack properties
func (s *Line) Stack(enable bool) *Line {
	if enable {
		s.model.Stack = "Total"
	} else {
		s.model.Stack = ""
	}

	return s
}

// Emphasis properties
func (s *Line) EmphasisEnable(enable bool) *Line {
	if enable {
		if s.model.Emphasis == nil {
			s.model.Emphasis = &model.Emphasis{}
		}

	} else {
		s.model.Emphasis = nil
	}

	return s
}

func (s *Line) EmphasisScale(scale bool) *Line {
	if s.model.Emphasis == nil {
		s.model.Emphasis = &model.Emphasis{}
	}
	s.model.Emphasis.Scale = scale
	return s
}


func (s *Line) EmphasisFocus(focus model.Focus) *Line {
	if s.model.Emphasis == nil {
		s.model.Emphasis = &model.Emphasis{}
	}
	s.model.Emphasis.Focus = focus

	return s
}


func (s *Line) EmphasisBlurScope(blurScope model.BlurScope) *Line {
	if s.model.Emphasis == nil {
		s.model.Emphasis = &model.Emphasis{}
	}

	s.model.Emphasis.BlurScope = blurScope
	return s
}

// Label properties
func (s *Line) LabelShow(show bool) *Line {
	if s.model.Label == nil {
		s.model.Label = &model.Label{}
	}
	s.model.Label.Show = show

	return s
}

func (s *Line) LabelPosition(position model.Position) *Line {
	if s.model.Label == nil {
		s.model.Label = &model.Label{}
	}
	s.model.Label.Position = position

	return s
}

func (s *Line) MarkLines() *MarkLines {
	if s.markLines != nil {
		return s.markLines
	}

	s.markLines = newMarkLines(s.model)

	return s.markLines
}

func (s *Line) MarkPoints() *MarkPoints {
	if s.markPoints != nil {
		return s.markPoints
	}

	s.markPoints = newMarkPoints(s.model)

	return s.markPoints
}
