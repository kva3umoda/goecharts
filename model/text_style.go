package model


// TextStyle is the option set for a text style component.
type TextStyle struct {
	// Font color
	Color string `json:"color,omitempty"`

	// Font style
	// Options: 'normal', 'italic', 'oblique'
	FontStyle FontStyle `json:"fontStyle,omitempty"`

	// main title font thick weight.
	// Options: 'normal', 'bold', 'bolder', 'lighter', 100 | 200 | 300 | 400...
	FontWeight FontWeight `json:"fontWeight,omitempty"`

	// Font family the main title font family.
	// Options: "sans-serif", 'serif' , 'monospace', 'Arial', 'Courier New', 'Microsoft YaHei', ...
	FontFamily string `json:"fontFamily,omitempty"`

	// Font size
	FontSize int `json:"fontSize,omitempty"`
}


type (
	FontStyle string
	FontWeight string
)


const (
	FontStyleNormal  FontStyle = "normal"
	FontStyleItalic  FontStyle = "italic"
	FontStyleOblique FontStyle = "oblique"

	FontWeightNormal  FontWeight = "normal"
	FontWeightBold    FontWeight = "bold"
	FontWeightBolder  FontWeight = "bolder"
	FontWeightLighter FontWeight = "lighter"
)