package goecharts

import (
	"time"

	"github.com/kva3umoda/goecharts/model"
)

type Dataset struct {
	index int
	model *model.Dataset
}

func newDataset(index int, dataset *model.Dataset) *Dataset {
	ds := &Dataset{
		index: index,
		model: dataset,
	}

	return ds
}

func (ds *Dataset) Index() int {
	return ds.index
}

func (ds *Dataset) AddDimension(name string, tp model.DataType) *Dataset {
	ds.model.Dimensions = append(ds.model.Dimensions,
		model.Dimension{Name: name, Type: tp},
	)

	return ds
}

func (ds *Dataset) AddDataRow(values ...interface{}) *Dataset {
	row := make([]interface{}, 0, len(ds.model.Dimensions))
	for _, v := range values {
		row = append(row, v)
	}
	ds.model.Source = append(ds.model.Source, row)

	return ds
}

func (ds *Dataset) AddDataColumn(dimension string, dt model.DataType, values []interface{}) *Dataset {
	ds.AddDimension(dimension,dt)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v)
	}

	return ds
}

func (ds *Dataset) AddDataColumnFloat(dimension string, values []float64) *Dataset {
	ds.AddDimension(dimension, model.DataFloat)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v)
	}

	return ds
}

func (ds *Dataset) AddDataColumnString(dimension string, values []string) *Dataset {
	ds.AddDimension(dimension, model.DataOrdinal)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v)
	}

	return ds
}

func (ds *Dataset) AddDataColumnInt(dimension string, values []int32) *Dataset {
	ds.AddDimension(dimension, model.DataInt)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v)
	}

	return ds
}

func (ds *Dataset) AddDataColumnUnixTime(dimension string, values []int64) *Dataset {
	ds.AddDimension(dimension, model.DataTime)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v)
	}

	return ds
}

func (ds *Dataset) AddDataColumnTime(dimension string, values []time.Time) *Dataset {
	ds.AddDimension(dimension, model.DataTime)
	for i, v := range values {
		if len(ds.model.Source) < len(values) {
			ds.model.Source = append(ds.model.Source,
				make([]interface{},
					len(ds.model.Dimensions)-1,
					len(ds.model.Dimensions)))
		}
		ds.model.Source[i] = append(ds.model.Source[i], v.Unix())
	}

	return ds
}
