package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeGaussianBlur
//
// English:
//
// The <feGaussianBlur> SVG filter primitive blurs the input image by the amount specified in stdDeviation, which
// defines the bell-curve.
//
// Português:
//
// A primitiva de filtro SVG <feGaussianBlur> desfoca a imagem de entrada pela quantidade especificada em stdDeviation,
// que define a curva de sino.
func NewTagSvgFeGaussianBlur(id string) (ref *html.TagSvgFeGaussianBlur) {
	ref = &html.TagSvgFeGaussianBlur{}
	ref.Init(id)

	return ref
}
