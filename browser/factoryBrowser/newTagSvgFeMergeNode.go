package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgFeMergeNode
//
// English:
//
// The feMergeNode takes the result of another filter to be processed by its parent <feMerge>.
//
// Português:
//
// O feMergeNode recebe o resultado de outro filtro para ser processado por seu pai <feMerge>.
func NewTagSvgFeMergeNode() (ref *html.TagSvgFeMergeNode) {
	ref = &html.TagSvgFeMergeNode{}
	ref.Init()

	return ref
}
