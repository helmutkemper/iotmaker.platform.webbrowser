package factoryColor

import "image/color"

func NewTanTransparent() color.RGBA {
	return color.RGBA{R: 0xd2, G: 0xb4, B: 0x8c, A: 0x00} // rgb(210, 180, 140)
}
