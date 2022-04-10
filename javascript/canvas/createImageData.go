package canvas

import (
	"image/color"
)

// fixme: pixelColor??

func (el *Canvas) CreateImageData(width, height interface{}, pixelColor color.RGBA) interface{} {
	imageData := el.SelfContext.Call("createImageData", width, height)
	return imageData.Get("data")
}
