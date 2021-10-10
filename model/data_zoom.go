package model

// DataZoom is the option set for a zoom component.
// dataZoom component is used for zooming a specific area, which enables user to
// investigate data in detail, or get an overview of the data, or get rid of outlier points.
// https://echarts.apache.org/en/option.html#dataZoom
type DataZoomInside struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type DataZoomType `json:"type" default:"inside"`

	// Whether to show the component. If is set to be false, it will not show, but its data filtering function still works.
	//Disabled bool `json:"disabled"`
	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	Start float32 `json:"start,omitempty"`

	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`

	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	Throttle float32 `json:"throttle,omitempty"`

	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}


type DataZoomSlider struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type DataZoomType `json:"type" default:"slider"`

	// Whether to show the component. If is set to be false, it will not show, but its data filtering function still works.
	//Show bool `json:"show"`
	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	Start float32 `json:"start,omitempty"`

	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`

	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	Throttle float32 `json:"throttle,omitempty"`

	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}


type DataZoomType string

const (
	// Data zoom functionalities is embeded inside coordinate systems, enable user to zoom or roam coordinate system by mouse dragging, mouse move or finger touch (in touch screen).
	DataZoomTypeInside DataZoomType = "inside"
	// A special slider bar is provided, on which coordinate systems can be zoomed or roamed by mouse dragging or finger touch (in touch screen).
	DataZoomTypeSlider DataZoomType = "slider"
)
