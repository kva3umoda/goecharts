package model

type Dataset struct {
	// Source data. Generally speaking, a source data describe a table, where these format below can be applied:
	Source [][]interface{} `json:"source,omitempty"`
	// dimensions can be used to define dimension info for series.data or dataset.source.
	Dimensions []Dimension `json:"dimensions,omitempty"`

	// Whether the first row/column of dataset.source represents dimension names. Optional values:
	//null/undefine: means auto detect whether the first row/column is dimension names or data.
	//true: the first row/column is dimension names.
	//false: data start from the first row/column.
	SourceHeader bool `json:"sourceHeader,omitempty"`

	Transform []Transform `json:"transform,omitempty"`

	// Specify the input dataset for dataset.transform. If dataset.transform specified but both fromDatasetIndex and fromDatasetId are not specified, fromDatasetIndex: 0 will be used by default.
	fromDatasetIndex float64 `json:"fromDatasetIndex,omitempty"`
	// Specify the input dataset for dataset.transform.
	fromDatasetId string `json:"fromDatasetId,omitempty"`
}

type Transform struct {
	Type   TransformType `json:"type,omitempty"`
	Config interface{}   `json:"config,omitempty"`
	Print  bool          `json:"print,omitempty"`
}

type TransformType string

const (
	TransformTypeFilter TransformType = "filter"
	TransformTypeSort   TransformType = "sort"
	TransformTypeXXX    TransformType = "xxx:xxx"
)

type Dimension struct {
	Name string   `json:"name"`
	Type DataType `json:"type"`
}

type DataType string

const (
	// number
	DataNumber DataType = "number"

	// float, that is, Float64Array
	DataFloat DataType = "float"

	// int, that is, Int32Array
	DataInt DataType = "int"

	// ordinal, discrete value, which represents string generally.
	DataOrdinal DataType = "ordinal"

	// time, time value, see data to check the format of time value.
	DataTime DataType = "time"
)
