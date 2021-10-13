package examples

import (
	"io"
	"os"
	"path"

	"github.com/kva3umoda/goecharts"
	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/render"
)

func candlestickFirst() *goecharts.Charts {
	const (
		upColor         = "#ec0000"
		upBorderColor   = "#8A0000"
		downColor       = "#00da3c"
		downBorderColor = "#008F28"
	)
	chart := goecharts.NewCharts("test one")

	chart.DataZoomYAxis(model.DataZoomTypeInside, 50, 100).
		DataZoomXAxis(model.DataZoomTypeSlider, 50, 100).
		ToolTip(model.TriggerTypeAxis, model.PointerTypeCross)

	dataset := chart.Dataset()
	dataset.AddDimension("date", model.DataTime).
		AddDimension("open", model.DataFloat).
		AddDimension("close", model.DataFloat).
		AddDimension("higher", model.DataFloat).
		AddDimension("lower", model.DataFloat)

	for _, values := range dataCandlestick {
		dataset.AddDataRow(values...)
	}

	chart2d := chart.AddChart2D("candlestick", model.AxisTypeCategory, model.AxisTypeValue)
	chart2d.YAxis().
		Scale(true).
		SplitArea(true)

	chart2d.XAxis().
		Scale(false).
		BoundaryGap(false).
		SplitLine(true)

	chart2d.AddSeriesCandlestick("APPL", "date", "open", "close", "higher", "lower").
		ItemStyleUpColor(upColor).
		ItemStyleUpBorderColor(upBorderColor).
		ItemStyleDownColor(downColor).
		ItemStyleDownBorderColor(downBorderColor)

	dataset.AddDataColumn("ma5", model.DataFloat, calculateMA(5, dataCandlestick)).
		AddDataColumn("ma10", model.DataFloat, calculateMA(10, dataCandlestick)).
		AddDataColumn("ma20", model.DataFloat, calculateMA(20, dataCandlestick)).
		AddDataColumn("ma30", model.DataFloat, calculateMA(30, dataCandlestick))

	chart2d.AddSeriesLine("MA5", "date", "ma5").
		Smooth(true).
		LineStyleOpacity(0.5)

	chart2d.AddSeriesLine("MA10", "date", "ma10").
		Smooth(true).
		LineStyleOpacity(0.5)

	chart2d.AddSeriesLine("MA20", "date", "ma20").
		Smooth(true).
		LineStyleOpacity(0.5)

	chart2d.AddSeriesLine("MA30", "date", "ma30").
		Smooth(true).
		LineStyleOpacity(0.5)

	return chart
}

type CandlestickExamples struct{}

func (CandlestickExamples) Examples(workdir string) {
	page := render.NewPage()
	page.AddCharts(
		candlestickFirst(),
	)
	f, err := os.Create(path.Join(workdir, "candlestick.html"))
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
