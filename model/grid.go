package model

// Drawing grid in rectangular coordinate. In a single grid, at most two X and Y axes each is allowed.
// Line chart, bar chart, and scatter chart (bubble chart) can be drawn in grid.
// https://echarts.apache.org/en/option.html#grid
type Grid struct {
	// Component ID, not specified by default.
	ID string `json:"id,omitempty"`

	// Whether to show the grid in rectangular coordinate.
	Show bool `json:"show"`

	// Distance between grid component and the left side of the container.
	//
	//left value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	//
	//If the left value is set to be 'left', 'center', or 'right', then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the top side of the container.
	//
	//top value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	//
	//If the left value is set to be 'top', 'middle', or 'bottom', then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the right side of the container.
	//
	//right value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the bottom side of the container.
	//
	//bottom value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'
	Bottom string `json:"bottom,omitempty"`

	// Width of grid component. Adaptive by default.
	Width string `json:"width,omitempty"`

	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Whether the grid region contains axis tick label of axis.
	//
	//When containLabel is false:
	//grid.left grid.right grid.top grid.bottom grid.width grid.height decide the location and size of the rectangle that is made of by xAxis and yAxis.
	//Setting to false will help when multiple grids need to be aligned at their axes.
	//When containLabel is true:
	//grid.left grid.right grid.top grid.bottom grid.width grid.height decide the location and size of the rectangle that contains the axes and the labels of the axes.
	//Setting to true will help when the length of axis labels is dynamic and is hard to approximate. This will avoid labels from overflowing the container or overlapping other components.
	ContainLabel bool `json:"containLabel,omitempty"`

	// Background color of grid, which is transparent by default.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Border color of grid. Support the same color format as backgroundColor.
	BorderColor string `json:"borderColor,omitempty"`

	// Border width of grid.
	BorderWidth int `json:"borderWidth,omitempty"`

	// Size of shadow blur. This attribute should be used along with shadowColor,shadowOffsetX, shadowOffsetY to set shadow to component.
	ShadowBlur string `json:"shadowBlur,omitempty"`

	// Shadow color. Support same format as color.
	ShadowColor string `json:"shadowColor,omitempty"`

	// Offset distance on the horizontal direction of shadow.
	ShadowOffsetX int `json:"shadowOffsetX,omitempty"`

	// Offset distance on the horizontal direction of shadow.
	ShadowOffsetY int `json:"shadowOffsetY,omitempty"`
}
