package svg

import mt "github.com/rustyoz/Mtransform"

type Ellipse struct {
	Id        string `xml:"id,attr"`
	Transform string `xml:"transform,attr"`
	Style     string `xml:"style,attr"`
	Cx        string `xml:"cx,attr"`
	Cy        string `xml:"cy,attr"`
	Rx        string `xml:"rx,attr"`
	Ry        string `xml:"ry,attr"`

	transform mt.Transform
	group     *Group
}
