package render

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


// Initialization contains options for the canvas.
type Initialization struct {
	// HTML title
	PageTitle string `default:"Awesome goecharts"`

	// Width of canvas
	Width string `default:"1200px"`

	// Height of canvas
	Height string `default:"600px"`

	// BackgroundColor of canvas
	BackgroundColor string

	// ChartRender unique ID
	ChartID string

	// Assets host
	AssetsHost string `default:"https://cdn.jsdelivr.net/npm/echarts@5.2.1/dist/"`

	// Theme of chart
	Theme string `default:"white"`
}

// Validate validates the initialization configurations.
func (opt *Initialization) Validate() {
	setDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = generateUniqueID()
	}
}

// set default values for the struct field.
// origin from: https://github.com/mcuadros/go-defaults
func setDefaultValue(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	t := elem.Type()

	for i := 0; i < t.NumField(); i++ {
		// handle `default` tag only
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(elem.Field(i), defaultVal)
		}
	}
}

// setField handles String/Bool types only.
func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}

const (
	chartIDSize = 12
)

// generate the unique ID for each chart.
func generateUniqueID() string {
	var b [chartIDSize]byte
	for i := range b {
		b[i] = randByte()
	}
	return string(b[:])
}

func randByte() byte {
	c := 65 // A
	if rand.Intn(10) > 5 {
		c = 97 // a
	}
	return byte(c + rand.Intn(26))
}
