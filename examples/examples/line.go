package examples

import (
	"io"
	"os"
	"path"

	"github.com/kva3umoda/goecharts"
	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/render"
)

func lineBasicLineChart() *goecharts.Charts {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}

	charts := goecharts.NewCharts("Basic Line Chart")
	charts.Dataset().
		AddDataColumnString("weekday", weeks).
		AddDataColumnInt("value", value)

	charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue).
		AddSeriesLine("line", "weekday", "value")

	return charts
}

func lineSmoothedLineChart() *goecharts.Charts {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}

	charts := goecharts.NewCharts("Smoothed Line Chart")
	charts.Dataset().
		AddDataColumnString("weekday", weeks).
		AddDataColumnInt("value", value)

	charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue).
		AddSeriesLine("line", "weekday", "value").
		Smooth(true)

	return charts
}

func lineBasicAreaChart() *goecharts.Charts {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}
	charts := goecharts.NewCharts("Basic Area Chart")
	charts.Dataset().
		AddDataColumnString("weekday", weeks).
		AddDataColumnInt("value", value)

	chart2d := charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue)

	chart2d.YAxis().BoundaryGap(true)

	chart2d.AddSeriesLine("line", "weekday", "value").
		Area(true)

	return charts
}

func lineStackedLineChart() *goecharts.Charts {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	valueEmail := []int32{120, 132, 101, 134, 90, 230, 210}
	valueUnionAds := []int32{220, 182, 191, 234, 290, 330, 310}
	valueVideoAds := []int32{150, 232, 201, 154, 190, 330, 410}
	valueDirect := []int32{320, 332, 301, 334, 390, 330, 320}
	valueSearchEngine := []int32{820, 932, 901, 934, 1290, 1330, 1320}

	charts := goecharts.NewCharts("Stacked Line Chart").
		ToolTip("axis", model.PointerTypeCross)

	charts.Dataset().
		AddDataColumnString("weekday", weeks).
		AddDataColumnInt("email", valueEmail).
		AddDataColumnInt("unionAds", valueUnionAds).
		AddDataColumnInt("videoAds", valueVideoAds).
		AddDataColumnInt("direct", valueDirect).
		AddDataColumnInt("searchEngine", valueSearchEngine)

	charts.ToolTip("axis", model.PointerTypeCross)

	chart2d := charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue)

	chart2d.AddSeriesLine("email", "weekday", "email").Stack(true)
	chart2d.AddSeriesLine("unionAds", "weekday", "unionAds").Stack(true)
	chart2d.AddSeriesLine("videoAds", "weekday", "videoAds").Stack(true)
	chart2d.AddSeriesLine("direct", "weekday", "direct").Stack(true)
	chart2d.AddSeriesLine("searchEngine", "weekday", "searchEngine").Stack(true)

	return charts
}

func lineStackedAreaChart() *goecharts.Charts {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	valueEmail := []int32{120, 132, 101, 134, 90, 230, 210}
	valueUnionAds := []int32{220, 182, 191, 234, 290, 330, 310}
	valueVideoAds := []int32{150, 232, 201, 154, 190, 330, 410}
	valueDirect := []int32{320, 332, 301, 334, 390, 330, 320}
	valueSearchEngine := []int32{820, 932, 901, 934, 1290, 1330, 1320}

	charts := goecharts.NewCharts("Stacked Area Chart").
		ToolTip("axis", model.PointerTypeCross)

	charts.Dataset().
		AddDataColumnString("weekday", weeks).
		AddDataColumnInt("email", valueEmail).
		AddDataColumnInt("unionAds", valueUnionAds).
		AddDataColumnInt("videoAds", valueVideoAds).
		AddDataColumnInt("direct", valueDirect).
		AddDataColumnInt("searchEngine", valueSearchEngine)

	charts.ToolTip("axis", model.PointerTypeCross)

	chart2d := charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue)

	chart2d.AddSeriesLine("email", "weekday", "email").
		Stack(true).
		Area(true).
		EmphasisScale(true).
		EmphasisFocus(model.FocusSeries).
		EmphasisBlurScope(model.BlurScopeCoordinateSystem)

	chart2d.AddSeriesLine("unionAds", "weekday", "unionAds").
		Stack(true).
		Area(true).
		EmphasisScale(true).
		EmphasisFocus(model.FocusSeries).
		EmphasisBlurScope(model.BlurScopeCoordinateSystem)
	chart2d.AddSeriesLine("videoAds", "weekday", "videoAds").
		Stack(true).
		Area(true).
		EmphasisScale(true).
		EmphasisFocus(model.FocusSeries).
		EmphasisBlurScope(model.BlurScopeCoordinateSystem)
	chart2d.AddSeriesLine("direct", "weekday", "direct").
		Stack(true).
		Area(true).
		EmphasisScale(true).
		EmphasisFocus(model.FocusSeries).
		EmphasisBlurScope(model.BlurScopeCoordinateSystem)
	chart2d.AddSeriesLine("searchEngine", "weekday", "searchEngine").
		Stack(true).
		Area(true).
		EmphasisScale(true).
		EmphasisFocus(model.FocusSeries).
		EmphasisBlurScope(model.BlurScopeCoordinateSystem).
		LabelShow(true).
		LabelPosition(model.PositionTop)

	return charts
}

func lineLineWithMarkLines() *goecharts.Charts {
	dataX := []string{"A", "B", "C", "D", "E", "F"}
	dataY := []float64{0.3, 1.4, 1.2, 1, 0.6}

	charts := goecharts.NewCharts("Line With Mark Lines").
		ToolTip("axis", model.PointerTypeCross)

	charts.Dataset().
		AddDataColumnString("x", dataX).
		AddDataColumnFloat("y", dataY)

	chart2d := charts.AddChart2D("", model.AxisTypeCategory, model.AxisTypeValue)

	chart2d.XAxis().BoundaryGap(true)
	chart2d.YAxis().Max(2)

	series := chart2d.AddSeriesLine("l", "x", "y")

	series.MarkLines().
		AddSpecial("max", model.MarkLineTypeMax).
		AddSpecial("min", model.MarkLineTypeMin).
		AddSpecial("median", model.MarkLineTypeMedian).
		AddSpecial("average", model.MarkLineTypeAverage).
		AddVertical("vertical", "B").
		AddHorizontal("horizontal", 1.8).
		AddByPoints("points", "B", 0.3, model.SymbolCircle, "F", 1.7, model.SymbolArrow)

	return charts
}

type LineExamples struct{}

func (LineExamples) Examples(workdir string) {
	page := render.NewPage()
	page.AddCharts(
		lineBasicLineChart(),
		lineSmoothedLineChart(),
		lineBasicAreaChart(),
		lineStackedLineChart(),
		lineStackedAreaChart(),
		lineLineWithMarkLines(),
	)
	f, err := os.Create(path.Join(workdir, "line.html"))
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
