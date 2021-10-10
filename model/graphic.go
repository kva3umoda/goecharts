package model

type Graphic struct {
	Elements []GraphicElement `json:"elements,omitempty"`
}

type GraphicElement struct {
	Type ElementType `json:"type"`
}

type ElementType string

const (
	ElementTypeGroup       ElementType = "group"
	ElementTypeImage       ElementType = "image"
	ElementTypeText        ElementType = "text"
	ElementTypeRect        ElementType = "rect"
	ElementTypeCircle      ElementType = "circle"
	ElementTypeRing        ElementType = "ring"
	ElementTypeSector      ElementType = "sector"
	ElementTypeArc         ElementType = "arc"
	ElementTypePolygon     ElementType = "polygon"
	ElementTypePolyline    ElementType = "polyline"
	ElementTypeLine        ElementType = "line"
	ElementTypeBezierCurve ElementType = "bezierCurve"
)
