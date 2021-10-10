package render

import (
	"fmt"
	"regexp"
	"strings"

)

var funcPat = regexp.MustCompile(`\n|\t`)

const funcMarker = "__f__"

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, funcPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts is the option set for handling function type.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	return fmt.Sprintf("%s%s%s", funcMarker, funcPat.ReplaceAllString(fn, ""), funcMarker)
}



// Assets contains options for static assets.
type Assets struct {
	JSAssets  OrderedSet
	CSSAssets OrderedSet

	CustomizedJSAssets  OrderedSet
	CustomizedCSSAssets OrderedSet
}

// InitAssets inits the static assets storage.
func (opt *Assets) InitAssets() {
	opt.JSAssets.Init("echarts.min.js")
	opt.CSSAssets.Init()

	opt.CustomizedJSAssets.Init()
	opt.CustomizedCSSAssets.Init()
}

// AddCustomizedJSAssets adds the customized javascript assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedJSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedJSAssets.Add(assets[i])
	}
}

// AddCustomizedCSSAssets adds the customized css assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedCSSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedCSSAssets.Add(assets[i])
	}
}

// Validate validates the static assets configurations
func (opt *Assets) Validate(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		if !strings.HasPrefix(opt.JSAssets.Values[i], host) {
			opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
		}
	}

	for i := 0; i < len(opt.CSSAssets.Values); i++ {
		if !strings.HasPrefix(opt.CSSAssets.Values[i], host) {
			opt.CSSAssets.Values[i] = host + opt.CSSAssets.Values[i]
		}
	}
}


// OrderedSet represents an ordered set.
type OrderedSet struct {
	filter map[string]bool
	Values []string
}

// Init creates a new OrderedSet instance, and adds any given items into this set.
func (o *OrderedSet) Init(items ...string) {
	o.filter = make(map[string]bool)
	for _, item := range items {
		o.Add(item)
	}
}

// Add adds a new item into the ordered set.
func (o *OrderedSet) Add(item string) {
	if !o.filter[item] {
		o.filter[item] = true
		o.Values = append(o.Values, item)
	}
}
