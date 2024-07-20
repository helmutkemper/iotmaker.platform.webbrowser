// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="50" cy="50" r="50"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(50),
	)

	stage.Append(s1)

	<-done
}
