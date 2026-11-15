package pages

import "github.com/at-least/ciderui/internal/signedin"

var media = func() signedin.Spec {
	thumb := []string{"border-top-left-radius", "border-top-width", "border-top-style", "background-color", "object-fit", "cursor", "box-shadow"}
	return signedin.Spec{
		Path: "/apps/<app>/distribution/ios/version/inflight",
		Probes: []signedin.Probe{
			{Name: "thumb.img", Select: `img[class*="Img-sc"]`, Props: thumb, Box: true},
			{Name: "thumb.frame", Select: `img[class*="Img-sc"]`, Up: 1, Props: thumb, Box: true},
			{Name: "thumb.item", Select: `img[class*="Img-sc"]`, Up: 3, Props: []string{"margin-left", "margin-right", "cursor"}, Box: true},
			{Name: "thumb.row", Select: `img[class*="Img-sc"]`, Up: 5, Props: []string{"display", "margin-top", "gap"}},
		},
	}
}()
