package testing

import (
	"testing"

	"github.com/kva3umoda/goecharts/charts"
	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/render"
	"github.com/stretchr/testify/assert"
)

func Test_Line_BasicLineChart(t *testing.T) {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}
	chart := charts.NewCartesian2D("Basic Line Chart")
	chart.DatasetColumnString("weekday", weeks)
	chart.DatesetColumnInt("value", value)

	chart.YAxis(model.AxisTypeValue)
	chart.XAxis(model.AxisTypeCategory)

	chart.AddSeriesLine("line").Dimension("weekday", "value")

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("Line_BasicLineChart.html")
	assert.NoError(t, err)
}

func Test_Line_SmoothedLineChart(t *testing.T) {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}
	chart := charts.NewCartesian2D("Smoothed Line Chart")
	chart.DatasetColumnString("weekday", weeks)
	chart.DatesetColumnInt("value", value)

	chart.YAxis(model.AxisTypeValue)
	chart.XAxis(model.AxisTypeCategory)

	chart.AddSeriesLine("line").Dimension("weekday", "value").
		SmoothEnabled()

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("Line_SmoothedLineChart.html")
	assert.NoError(t, err)
}

func Test_Line_BasicAreaChart(t *testing.T) {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	value := []int32{150, 230, 224, 218, 135, 147, 260}
	chart := charts.NewCartesian2D("Basic Area Chart")
	chart.DatasetColumnString("weekday", weeks)
	chart.DatesetColumnInt("value", value)

	chart.YAxis(model.AxisTypeValue).BoundaryGap(true)
	chart.XAxis(model.AxisTypeCategory)

	chart.AddSeriesLine("line").Dimension("weekday", "value").AreaEnabled()

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("line03.html")
	assert.NoError(t, err)
}

func Test_Line_StackedLineChart(t *testing.T) {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	valueEmail := []int32{120, 132, 101, 134, 90, 230, 210}
	valueUnionAds := []int32{220, 182, 191, 234, 290, 330, 310}
	valueVideoAds := []int32{150, 232, 201, 154, 190, 330, 410}
	valueDirect := []int32{320, 332, 301, 334, 390, 330, 320}
	valueSearchEngine := []int32{820, 932, 901, 934, 1290, 1330, 1320}

	chart := charts.NewCartesian2D("Stacked Line Chart")
	chart.DatasetColumnString("weekday", weeks)
	chart.DatesetColumnInt("email", valueEmail)
	chart.DatesetColumnInt("unionAds", valueUnionAds)
	chart.DatesetColumnInt("videoAds", valueVideoAds)
	chart.DatesetColumnInt("direct", valueDirect)
	chart.DatesetColumnInt("searchEngine", valueSearchEngine)
	chart.ToolTip("axis", model.PointerTypeCross)

	chart.YAxis(model.AxisTypeValue)
	chart.XAxis(model.AxisTypeCategory).BoundaryGap(false)

	chart.AddSeriesLine("email").Dimension("weekday", "email").StackEnabled()
	chart.AddSeriesLine("unionAds").Dimension("weekday", "unionAds").StackEnabled()
	chart.AddSeriesLine("videoAds").Dimension("weekday", "videoAds").StackEnabled()
	chart.AddSeriesLine("direct").Dimension("weekday", "direct").StackEnabled()
	chart.AddSeriesLine("searchEngine").Dimension("weekday", "searchEngine").StackEnabled()

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("Line_StackedLineChart.html")
	assert.NoError(t, err)
}

func Test_Line_StackedAreaChart(t *testing.T) {
	weeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	valueEmail := []int32{120, 132, 101, 134, 90, 230, 210}
	valueUnionAds := []int32{220, 182, 191, 234, 290, 330, 310}
	valueVideoAds := []int32{150, 232, 201, 154, 190, 330, 410}
	valueDirect := []int32{320, 332, 301, 334, 390, 330, 320}
	valueSearchEngine := []int32{820, 932, 901, 934, 1290, 1330, 1320}

	chart := charts.NewCartesian2D("Stacked Area Chart")
	chart.DatasetColumnString("weekday", weeks)
	chart.DatesetColumnInt("email", valueEmail)
	chart.DatesetColumnInt("unionAds", valueUnionAds)
	chart.DatesetColumnInt("videoAds", valueVideoAds)
	chart.DatesetColumnInt("direct", valueDirect)
	chart.DatesetColumnInt("searchEngine", valueSearchEngine)
	chart.ToolTip("axis", model.PointerTypeCross)

	chart.YAxis(model.AxisTypeValue)
	chart.XAxis(model.AxisTypeCategory).BoundaryGap(false)

	chart.AddSeriesLine("email").Dimension("weekday", "email").StackEnabled().AreaEnabled().
		EmphasisEnabled(true, model.FocusSeries, model.BlurScopeCoordinateSystem)
	chart.AddSeriesLine("unionAds").Dimension("weekday", "unionAds").StackEnabled().AreaEnabled().
		EmphasisEnabled(true, model.FocusSeries, model.BlurScopeCoordinateSystem)
	chart.AddSeriesLine("videoAds").Dimension("weekday", "videoAds").StackEnabled().AreaEnabled().
		EmphasisEnabled(true, model.FocusSeries, model.BlurScopeCoordinateSystem)
	chart.AddSeriesLine("direct").Dimension("weekday", "direct").StackEnabled().AreaEnabled().
		EmphasisEnabled(true, model.FocusSeries, model.BlurScopeCoordinateSystem)
	chart.AddSeriesLine("searchEngine").Dimension("weekday", "searchEngine").StackEnabled().AreaEnabled().
		EmphasisEnabled(true, model.FocusSeries, model.BlurScopeCoordinateSystem).
		LabelEnabled(model.PositionTop)

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("Line_StackedAreaChart.html")
	assert.NoError(t, err)
}

func Test_Line_LineWithMarkLines(t *testing.T) {
	dataX := []string{"A", "B", "C", "D", "E", "F"}
	dataY := []float64{0.3, 1.4, 1.2, 1, 0.6}

	chart := charts.NewCartesian2D("Line With Mark Lines")

	chart.DatasetColumnString("x", dataX).
		DatasetColumnFloat("y", dataY)

	chart.ToolTip("axis", model.PointerTypeCross)
	chart.XAxis(model.AxisTypeCategory).BoundaryGap(true)
	chart.YAxis(model.AxisTypeValue).Max(2)

	series := chart.AddSeriesLine("l").
		Dimension("x", "y")

	series.MarkLines().
		AddSpecial("max", model.MarkLineTypeMax).
		AddSpecial("min", model.MarkLineTypeMin).
		AddSpecial("median", model.MarkLineTypeMedian).
		AddSpecial("average", model.MarkLineTypeAverage).
		AddVertical("vertical", "B").
		AddHorizontal("horizontal", 1.8).
		AddByPoints("points", "B", 0.3, model.SymbolCircle, "F", 1.7, model.SymbolArrow)

	page := render.NewPage()
	page.AddCharts(chart)
	err := page.RenderHTML("Line_LineWithMarkLines.html")
	assert.NoError(t, err)
}
