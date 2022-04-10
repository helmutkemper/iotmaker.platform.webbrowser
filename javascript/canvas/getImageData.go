package canvas

import (
	"image/color"
)

// todo: rewrite documentation

// GetImageData
// en: Returns an ImageData map[x][y]color.RGBA that copies the pixel data for the
// specified rectangle on a canvas
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     return: map[x(int)][y(int)]color.RGBA
//             Note: return x and y are NOT relative to the coordinate (0,0) on the
//             image, are relative to the coordinate (0,0) on the canvas
//
//     Note: The ImageData object is not a picture, it specifies a part (rectangle)
//     on the canvas, and holds information of every pixel inside that rectangle.
//
//     For every pixel in the map[x][y] there are four pieces of information, the
//     color.RGBA values:
//     R - The color red (from 0-255)
//     G - The color green (from 0-255)
//     B - The color blue (from 0-255)
//     A - The alpha channel (from 0-255; 0 is transparent and 255 is fully visible)
//
//     Tip: After you have manipulated the color/alpha information in the map[x][y],
//     you can copy the image data back onto the canvas with the putImageData()
//     method.
//
// pr_br: Retorna um mapa map[x][y]color.RGBA com parte dos dados da imagem contida
// no retângulo especificado.
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     width: comprimento do retângulo a ser copiado
//     height: altura do retângulo a ser copiado
//     return: map[x(int)][y(int)]color.RGBA
//             Nota: x e y do retorno não são relativos a coordenada (0,0) da imagem,
//             são relativos a coordenada (0,0) do canvas
//
//     Nota: Os dados da imagem não são uma figura, eles representam uma parte
//     retangular do canvas e guardam informações de cada pixel contido nessa área
//
//     Para cada pixel contido no mapa há quatro peças de informação com valores no
//     formato de color.RGBA:
//     R - Cor vermelha (de 0-255)
//     G - Cor verde (de 0-255)
//     B - Cor azul (de 0-255)
//     A - Canal alpha (de 0-255; onde, 0 é transparente e 255 é totalmente visível)
//
//     Dica: Depois de manipular as informações de cor/alpha contidas no map[x][y],
//     elas podem ser colocadas de volta no canvas com o método putImageData().
func (el *Canvas) GetImageData(x, y, width, height int) map[int]map[int]color.RGBA {
	dataInterface := el.SelfContext.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	ret := make(map[int]map[int]color.RGBA)

	var rgbaLength int = 4

	var tmpR, tmpG, tmpB, tmpA uint8
	var i int = 0
	var xp int
	var yp int
	for yp = 0; yp != height; yp += 1 {
		for xp = 0; xp != width; xp += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			if dataJs.Index(i+0).IsUndefined() == true {
				tmpR = 0
			} else {
				tmpR = uint8(dataJs.Index(i + 0).Int())
			}

			if dataJs.Index(i+1).IsUndefined() == true {
				tmpG = 0
			} else {
				tmpG = uint8(dataJs.Index(i + 1).Int())
			}

			if dataJs.Index(i+2).IsUndefined() == true {
				tmpB = 0
			} else {
				tmpB = uint8(dataJs.Index(i + 2).Int())
			}

			if dataJs.Index(i+3).IsUndefined() == true {
				tmpA = 0
			} else {
				tmpA = uint8(dataJs.Index(i + 3).Int())
			}

			i += rgbaLength

			if len(ret[x+xp]) == 0 {
				ret[x+xp] = make(map[int]color.RGBA)
			}

			ret[x+xp][y+yp] = color.RGBA{
				R: tmpR,
				G: tmpG,
				B: tmpB,
				A: tmpA,
			}
		}
	}

	return ret
}
