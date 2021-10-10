package model


// ParallelComponent is the option set for parallel component.
type ParallelComponent struct {
	// Distance between parallel component and the left side of the container.
	// Left value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%';
	// and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between parallel component and the top side of the container.
	// Top value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	// and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between parallel component and the right side of the container.
	// Right value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between parallel component and the bottom side of the container.
	// Bottom value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// ParallelAxis is the option set for a parallel axis.
type ParallelAxis struct {
	// Dimension index of coordinate axis.
	Dim int `json:"dim,omitempty"`

	// Name of axis.
	Name string `json:"name,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	// In category axis, it can also be set as the ordinal number.
	Max interface{} `json:"max,omitempty"`

	// Compulsively set segmentation interval for axis.
	Inverse bool `json:"inverse,omitempty"`

	// Location of axis name. Options: "start", "middle", "center", "end"
	// default "end"
	NameLocation string `json:"nameLocation,omitempty"`

	// Type of axis.
	// Options：
	// "value"：Numerical axis, suitable for continuous data.
	// "category" Category axis, suitable for discrete category data. Category data can be auto retrieved from series.
	// "log"  Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Category data，works on (type: "category").
	Data interface{} `json:"data,omitempty"`
}
