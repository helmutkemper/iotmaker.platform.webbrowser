package canvas

// GetImageDataJsValue
//
// English:
//
//  Returns an ImageData js.Value that copies the pixel data for the specified rectangle on a canvas
//
//   Input:
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//
//   Output:
//     data: js.Value
//
//   Note:
//     * return x and y are NOT relative to the coordinate (0,0) on the image, are relative to the
//       coordinate (0,0) on the canvas;
//     * The ImageData object is not a picture, it specifies a part (rectangle) on the canvas, and
//       holds information of every pixel inside that rectangle.
//
// For every pixel in the map[x][y] there are four pieces of information, the color.RGBA values:
//   R - The color red (from 0-255)
//   G - The color green (from 0-255)
//   B - The color blue (from 0-255)
//   A - The alpha channel (from 0-255; 0 is transparent and 255 is fully visible)
//
//   Tip:
//   * After you have manipulated the color/alpha information in the map[x][y], you can copy the image
//     data back onto the canvas with the putImageData() method.
//
// pr_br: Retorna um js.Value com parte dos dados da imagem contida no retângulo especificado.
//
//   Entrada:
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão ser copiados;
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão ser copiados;
//     width: comprimento do retângulo a ser copiado;
//     height: altura do retângulo a ser copiado;
//
//   Saída:
//     data: js.Value
//
//   Nota:
//     * x e y do retorno não são relativos a coordenada (0,0) da imagem, são relativos a coordenada
//       (0,0) do canvas;
//     * Os dados da imagem não são uma figura, eles representam uma parte retangular do canvas e
//       guardam informações de cada pixel contido nessa área.
//
// Para cada pixel contido no mapa há quatro peças de informação com valores no formato de color.RGBA:
//   R - Cor vermelha (de 0-255)
//   G - Cor verde (de 0-255)
//   B - Cor azul (de 0-255)
//   A - Canal alpha (de 0-255; onde, 0 é transparente e 255 é totalmente visível)
//
//   Dica:
//     * Depois de manipular as informações de cor/alpha contidas no map[x][y], elas podem ser
//       colocadas de volta no canvas com o método putImageData().
func (el *Canvas) GetImageDataJsValue(x, y, width, height int) (data interface{}) {
	dataInterface := el.SelfContext.Call("getImageData", x, y, width, height)
	return dataInterface.Get("data")
}
