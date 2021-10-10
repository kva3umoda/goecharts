package model

// Polar Bar options
type Polar struct {
	// Component ID, not specified by default. If specified, it can be used to refer the component in option or API.
	ID string `json:"id,omitempty"`
	// zlevel value of all graphical elements in .
	ZLevel int `json:"zlevel,omitempty"`

	// z value of all graphical elements in , which controls order of drawing graphical components. Components with smaller z values may be overwritten by those with larger z values.
	Z int `json:"z,omitempty"`

	// Center position of Polar coordinate, the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width, and the second item to the height.
	Center [2]string `json:"center,omitempty"`

	// Radius of Polar coordinate. Value can be:
	Radius [2]string `json:"radius,omitempty"`

	// tooltip settings in the coordinate system component.
	Tooltip Tooltip `json:"tooltip,omitempty"`
}

