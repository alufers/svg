package svg

import mt "github.com/rustyoz/Mtransform"

type Line struct {
	Id        string `xml:"id,attr"`
	Transform string `xml:"transform,attr"`
	Style     string `xml:"style,attr"`
	X1        string `xml:"x1,attr"`
	X2        string `xml:"x2,attr"`
	Y1        string `xml:"y1,attr"`
	Y2        string `xml:"y2,attr"`

	transform mt.Transform
	group     *Group
}
