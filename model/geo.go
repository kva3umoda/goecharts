package model


// GeoComponent is the option set for geo component.
// https://echarts.apache.org/en/option.html#geo
type GeoComponent struct {
	// Map charts.
	Map string `json:"map,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Set this to true, to prevent interaction with the axis.
	Silent bool `json:"silent,omitempty"`
}

// ItemStyle represents a style of an item.
type ItemStyle struct {
	// Color of chart
	// Kline Up candle color
	Color string `json:"color,omitempty"`

	// Kline Down candle color
	Color0 string `json:"color0,omitempty"`

	// BorderColor is the hart border color
	// Kline  Up candle border color
	BorderColor string `json:"borderColor,omitempty"`

	// Kline Down candle border color
	BorderColor0 string `json:"borderColor0,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`
}
