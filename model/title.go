package model

// Title is the option set for a title component.
// https://echarts.apache.org/en/option.html#title
type Title struct {
	// The main title text, supporting \n for newlines.
	Text string `json:"text,omitempty"`
}
