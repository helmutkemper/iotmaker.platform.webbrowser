package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"log"
	"syscall/js"
)

type TagCanvas struct {
	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class

	context js.Value
	width   int
	height  int
}

// Id
//
// English:
//
//  Specifies a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//  Especifica um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (el *TagCanvas) Id(id string) (ref *TagCanvas) {
	el.id = id
	el.selfElement.Set("id", id)
	return el
}

// CreateElement
//
// English:
//
//  In an HTML document, the Document.createElement() method creates the specified HTML element or an
//  HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//  Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//  HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (el *TagCanvas) CreateElement(tag Tag, width, height int) (ref *TagCanvas) {
	el.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if el.selfElement.IsUndefined() == true || el.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}
	el.tag = tag

	el.context = el.selfElement.Call("getContext", "2d")
	el.selfElement.Set("width", width)
	el.selfElement.Set("height", height)

	el.width = width
	el.height = height

	return el
}

// DrawImage
//
// English:
//
//  Draws a preloaded image on the canvas element.
//
//   Input:
//     image: js.Value object with a preloaded image.
//
// Português:
//
//  Desenha uma imagem pre carregada no elemento canvas.
//
//   Entrada:
//     image: objeto js.Value com uma imagem pré carregada.
func (el *TagCanvas) DrawImage(image interface{}) (ref *TagCanvas) {
	el.context.Call("drawImage", image, 0, 0, el.width, el.height)
	return el
}

// AppendById
//
// English:
//
//  Adds a node to the end of the list of children of a specified parent node. If the node already
//  exists in the document, it is removed from its current parent node before being added to the
//  new parent.
//
//   Input:
//     appendId: id of parent element.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//  documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: id do elemento pai.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (el *TagCanvas) AppendById(appendId string) (ref *TagCanvas) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return el
	}

	toAppend.Call("appendChild", el.selfElement)
	return el
}

// Append
//
// English:
//
//  Adds a node to the end of the list of children of a specified parent node. If the node already
//  exists in the document, it is removed from its current parent node before being added to the new
//  parent.
//
//   Input:
//     append: element in js.Value format.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//  documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: elemento no formato js.Value.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (el *TagCanvas) Append(append interface{}) (ref *TagCanvas) {
	switch append.(type) {
	case *TagCanvas:
		el.selfElement.Call("appendChild", append.(*TagCanvas).selfElement)
	case js.Value:
		el.selfElement.Call("appendChild", append)
	case string:
		toAppend := js.Global().Get("document").Call("getElementById", append.(string))
		if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
			log.Print(KIdToAppendNotFound, append.(string))
			return el
		}

		toAppend.Call("appendChild", el.selfElement)
	}

	return el
}

// GetCollisionData
//
// English:
//
//  Returns an array (x,y) with a boolean indicating transparency.
//
// The collision map is a quick way to preload data about the coordinates of where there are parts of
// the image.
//
//   Output:
//     data: [x][y]transparent
//
// Português:
//
//  Retorna uma array (x,y) com um booleano indicando transparência.
//
// O mapa de colisão é uma forma rápida de deixar um dado pre carregado sobre as coordenadas de onde
// há partes da imagem.
//
//   Saída:
//     data: [x][y]transparente
func (el *TagCanvas) GetCollisionData() (data [][]bool) {

	dataInterface := el.context.Call("getImageData", 0, 0, el.width, el.height)
	dataJs := dataInterface.Get("data")

	var rgbaLength = 4

	var i = 0
	var x int
	var y int

	// [x][y]bool
	data = make([][]bool, el.width)
	for x = 0; x != el.width; x += 1 {
		data[x] = make([]bool, el.height)
		for y = 0; y != el.height; y += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			data[x][y] = dataJs.Index(i+3).Int() != 0

			i += rgbaLength
		}
	}

	return
}

// GetImageData
//
// English:
//
//  Returns an array copy of the image.
//
//   Input:
//     x: x position of the image;
//     y: y position of the image;
//     width: image width;
//     height: image height.
//
//   Output:
//     data: image in matrix format.
//       [x][y][0]: red color value between 0 and 255
//       [x][y][1]: green color value between 0 and 255
//       [x][y][2]: blue color value between 0 and 255
//       [x][y][3]: alpha color value between 0 and 255
//
// Português:
//
//  Retorna uma cópia matricial da imagem.
//
//   Entrada:
//     x: Posição x da imagem;
//     y: Posição y da imagem;
//     width: comprimento da imagem;
//     height: altura da imagem.
//
//   Saída:
//     data: imagem em formato matricial.
//       [x][y][0]: valor da cor vermelha entre 0 e 255
//       [x][y][1]: valor da cor verde entre 0 e 255
//       [x][y][2]: valor da cor azul entre 0 e 255
//       [x][y][3]: valor da cor alpha entre 0 e 255
func (el *TagCanvas) GetImageData(x, y, width, height int) (data [][][]uint8) {

	dataInterface := el.context.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	var rgbaLength = 4

	var i = 0
	x = 0
	y = 0

	// [x][y][4-channel]
	data = make([][][]uint8, width)
	for x = 0; x != width; x += 1 {
		data[x] = make([][]uint8, height)
		for y = 0; y != height; y += 1 {
			data[x][y] = make([]uint8, 4)

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			data[x][y][0] = uint8(dataJs.Index(i + 0).Int())
			data[x][y][1] = uint8(dataJs.Index(i + 1).Int())
			data[x][y][2] = uint8(dataJs.Index(i + 2).Int())
			data[x][y][3] = uint8(dataJs.Index(i + 3).Int())

			i += rgbaLength
		}
	}

	return
}

// PutImageData
//
// English:
//
//  Transform an array of data into an image.
//
//   Input:
//     imgData: data array with the new image;
//       [x][y][0]: red color value between 0 and 255;
//       [x][y][1]: green color value between 0 and 255;
//       [x][y][2]: blue color value between 0 and 255;
//       [x][y][3]: alpha color value between 0 and 255.
//     width: image width;
//     height: image height.
//
// Português:
//
//  Transforma uma matrix de dados em uma imagem.
//
//   Entrada:
//     imgData: array de dados com o a nova imagem;
//       [x][y][0]: valor da cor vermelha entre 0 e 255;
//       [x][y][1]: valor da cor verde entre 0 e 255;
//       [x][y][2]: valor da cor azul entre 0 e 255;
//       [x][y][3]: valor da cor alpha entre 0 e 255.
//     width: comprimento da imagem;
//     height: altura da imagem.
func (el *TagCanvas) PutImageData(imgData [][][]uint8, width, height int) (ref *TagCanvas) {

	dataJs := el.context.Call("createImageData", width, height)

	var rgbaLength = 4

	var i = 0
	var x int
	var y int
	for x = 0; x != width; x += 1 {
		for y = 0; y != height; y += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			dataJs.Get("data").SetIndex(i+0, imgData[x][y][0])
			dataJs.Get("data").SetIndex(i+1, imgData[x][y][1])
			dataJs.Get("data").SetIndex(i+2, imgData[x][y][2])
			dataJs.Get("data").SetIndex(i+3, imgData[x][y][3])

			i += rgbaLength
		}
	}

	el.context.Call("putImageData", dataJs, 0, 0, 0, 0, width, height)
	return el
}

// Arc
//
// English:
//
// Creates an arc/curve (used to create circles, or parts of circles)
//   Input:
//     x: The horizontal coordinate of the arc's center.
//     y: The vertical coordinate of the arc's center.
//     radius: The arc's radius. Must be positive.
//     startAngle: The angle at which the arc starts in radians, measured from the positive x-axis.
//     endAngle: The angle at which the arc ends in radians, measured from the positive x-axis.
//     anticlockwise: An optional Boolean. If true, draws the arc counter-clockwise between the start
//       and end angles. The default is false (clockwise).
//
//     Example:
//     Arc(100, 75, 50, 0, 2 * Math.PI, false);
func (el *TagCanvas) Arc(x, y, radius, startAngle int, endAngle float64, anticlockwise bool) (ref *TagCanvas) {
	el.context.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
	return el
}

// ArcTo
//
// English:
//
//  Creates an arc/curve between two tangents
//   Input:
//     x1:     The x-axis coordinate of the first control point.
//     y1:     The y-axis coordinate of the first control point.
//     x2:     The x-axis coordinate of the second control point.
//     y2:     The y-axis coordinate of the second control point.
//     radius: The arc's radius. Must be non-negative.
//
//   Example:
//     factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//       // Create a starting point
//       MoveTo(20, 20).
//       // Create a horizontal line
//       LineTo(100, 20).
//       // Create an arc
//       ArcTo(150, 20, 150, 70, 50).
//       // Continue with vertical line
//       LineTo(150, 120).
//       // Draw it
//       Stroke().
//		   AppendById("stage")
func (el *TagCanvas) ArcTo(x1, y1, x2, y2, radius int) (ref *TagCanvas) {
	el.context.Call("arcTo", x1, y1, x2, y2, radius)
	return el
}

// BeginPath
//	en: Begins a path, or resets the current path
//      Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
//      Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Inicia ou reinicializa uma nova rota no desenho
//      Dica: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), e arc(), para criar uma nova rota no desenho
//      Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *TagCanvas) BeginPath() (ref *TagCanvas) {
	el.context.Call("beginPath")
	return el
}

// BezierCurveTo
//
// English:
//
//  Creates a cubic Bézier curve
//   Input:
//     cp1x: The x-axis coordinate of the first control point.
//     cp1y: The y-axis coordinate of the first control point.
//     cp2x: The x-axis coordinate of the second control point.
//     cp2y: The y-axis coordinate of the second control point.
//     x: The x-axis coordinate of the end point.
//     y: The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.bezierCurveTo(20, 100, 200, 100, 200, 20);
//     ctx.stroke();
func (el *TagCanvas) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) (ref *TagCanvas) {
	el.context.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
	return el
}

// en: Clears the specified pixels within a given rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle to clear
//     y:      The y-coordinate of the upper-left corner of the rectangle to clear
//     width:  The width of the rectangle to clear, in pixels
//     height: The height of the rectangle to clear, in pixels
//
//     The clearRect() method clears the specified pixels within a given rectangle.
//     JavaScript syntax: context.clearRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 300, 150);
//     ctx.clearRect(20, 20, 100, 50);
func (el *TagCanvas) ClearRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("clearRect", x, y, width, height)
	return el
}

// en: Clips a region of any shape and size from the original canvas
//     The clip() method clips a region of any shape and size from the original canvas.
//     Tip: Once a region is clipped, all future drawing will be limited to the clipped region (no access to other
//     regions on the canvas). You can however save the current canvas region using the save() method before using the
//     clip() method, and restore it (with the restore() method) any time in the future.
func (el *TagCanvas) Clip(x, y int) (ref *TagCanvas) {
	el.context.Call("clip", x, y)
	return el
}

// en: Creates a path from the current point back to the starting point
//     The closePath() method creates a path from the current point back to the starting point.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
//     another color/gradient.
func (el *TagCanvas) ClosePath(x, y int) (ref *TagCanvas) {
	el.context.Call("closePath", x, y)
	return el
}

// en: Creates a linear gradient (to use on canvas content)
//     x0: The x-coordinate of the start point of the gradient
//     y0: The y-coordinate of the start point of the gradient
//     x1: The x-coordinate of the end point of the gradient
//     y1: The y-coordinate of the end point of the gradient
//
//     The createLinearGradient() method creates a linear gradient object.
//     The gradient can be used to fill rectangles, circles, lines, text, etc.
//     Tip: Use this object as the value to the strokeStyle or fillStyle properties.
//     Tip: Use the addColorStop() method to specify different colors, and where to position the colors in the gradient object.
//     JavaScript syntax:	context.createLinearGradient(x0, y0, x1, y1);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var grd = ctx.createLinearGradient(0, 0, 170, 0);
//     grd.addColorStop(0, "black");
//     grd.addColorStop(1, "white");
//     ctx.fillStyle = grd;
//     ctx.fillRect(20, 20, 150, 100);
func (el *TagCanvas) CreateLinearGradient(x0, y0, x1, y1 int) (ref *TagCanvas) {
	el.context.Call("createLinearGradient", x0, y0, x1, y1)
	return el
}

// en: Repeats a specified element in the specified direction
//     image: Specifies the image, canvas, or video element of the pattern to use
//     repeatedElement
//          repeat: Default. The pattern repeats both horizontally and vertically
//          repeat-x: The pattern repeats only horizontally
//          repeat-y: The pattern repeats only vertically
//          no-repeat: The pattern will be displayed only once (no repeat)
//
//     The createPattern() method repeats the specified element in the specified direction.
//     The element can be an image, video, or another <canvas> element.
//     The repeated element can be used to draw/fill rectangles, circles, lines etc.
//     JavaScript syntax:	context.createPattern(image, "repeat|repeat-x|repeat-y|no-repeat");
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var img = document.getElementById("lamp");
//     var pat = ctx.createPattern(img, "repeat");
//     ctx.rect(0, 0, 150, 100);
//     ctx.fillStyle = pat;
//     ctx.fill();
func (el *TagCanvas) CreatePattern(image js.Value, repeatRule CanvasRepeatRule) (ref *TagCanvas) {
	el.context.Call("createPattern", image, repeatRule)
	return
}

// en: Creates a radial/circular gradient (to use on canvas content)
//     x0: The x-coordinate of the starting circle of the gradient
//     y0: The y-coordinate of the starting circle of the gradient
//     r0: The radius of the starting circle
//     x1: The x-coordinate of the ending circle of the gradient
//     y1: The y-coordinate of the ending circle of the gradient
//     r1: The radius of the ending circle
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var grd = ctx.createRadialGradient(75, 50, 5, 90, 60, 100);
//     grd.addColorStop(0, "red");
//     grd.addColorStop(1, "white");
//     // Fill with gradient
//     ctx.fillStyle = grd;
//     ctx.fillRect(10, 10, 150, 100);
func (el *TagCanvas) CreateRadialGradient(x0, y0, r0, x1, y1 int, r1 float64) (ref *TagCanvas) {
	el.context.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
	return el
}

// en: Draws a "filled" rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The fillRect() method draws a "filled" rectangle. The default color of the fill is black.
//     Tip: Use the fillStyle property to set a color, gradient, or pattern used to fill the drawing.
//     JavaScript syntax: context.fillRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillRect(20, 20, 150, 100);
func (el *TagCanvas) FillRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("fillRect", x, y, width, height)
	return el
}

// en: Draws "filled" text on the canvas
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The fillText() method draws filled text on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the fillStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.fillText(text, x, y, maxWidth);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.fillText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.fillStyle = gradient;
//     ctx.fillText("Big smile!", 10, 90);
func (el *TagCanvas) FillText(text string, x, y, maxWidth int) (ref *TagCanvas) {
	el.context.Call("fillText", text, x, y, maxWidth)
	return el
}

// en: Sets or returns the current font properties for text content
//     font-style:            Specifies the font style. Possible values:
//          normal | italic | oblique
//
//     font-variant:          Specifies the font variant. Possible values:
//          normal | small-caps
//
//     font-weight:           Specifies the font weight. Possible values:
//          normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900
//
//     font-size/line-height: Specifies the font size and the line-height, in pixels
//     font-family:           Specifies the font family
//     caption:               Use the font captioned controls (like buttons, drop-downs, etc.)
//     icon:                  Use the font used to label icons
//     menu:                  Use the font used in menus (drop-down menus and menu lists)
//     message-box:           Use the font used in dialog boxes
//     small-caption:         Use the font used for labeling small controls
//     status-bar:            Use the fonts used in window status bar
//
//     The font property sets or returns the current font properties for text content on the canvas.
//     The font property uses the same syntax as the CSS font property.
//     Default value: 10px sans-serif
//     JavaScript syntax: context.font = "italic small-caps bold 12px arial";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "30px Arial";
//     ctx.fillText("Hello World", 10, 50);
func (el *TagCanvas) Font(font Font) (ref *TagCanvas) {
	el.context.Set("font", font.String())
	return el
}

// en: Sets or returns the current alpha or transparency value of the drawing
//     number: The transparency value. Must be a number between 0.0 (fully transparent) and 1.0 (no transparancy)
//
//     Default value: 1.0
//     JavaScript syntax: context.globalAlpha = number;
//
//     The globalAlpha property sets or returns the current alpha or transparency value of the drawing.
//     The globalAlpha property value must be a number between 0.0 (fully transparent) and 1.0 (no transparancy)
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(20, 20, 75, 50);
//     // Turn transparency on
//     ctx.globalAlpha = 0.2;
//     ctx.fillStyle = "blue";
//     ctx.fillRect(50, 50, 75, 50);
//     ctx.fillStyle = "green";
//     ctx.fillRect(80, 80, 75, 50);
func (el *TagCanvas) GlobalAlpha(value float64) (ref *TagCanvas) {
	el.context.Set("globalAlpha", value)
	return el
}

// en: Sets or returns how a new image are drawn onto an existing image
//
//     The globalCompositeOperation property sets or returns how a source (new) image are drawn onto a destination
//     (existing) image.
//     source image = drawings you are about to place onto the canvas.
//     destination image = drawings that are already placed onto the canvas.
//
//     Default value: source-over
//     JavaScript syntax: context.globalCompositeOperation = "source-in";
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(20, 20, 75, 50);
//     ctx.globalCompositeOperation = "source-over";
//     ctx.fillStyle = "blue";
//     ctx.fillRect(50, 50, 75, 50);
//     ctx.fillStyle = "red";
//     ctx.fillRect(150, 20, 75, 50);
//     ctx.globalCompositeOperation = "destination-over";
//     ctx.fillStyle = "blue";
//     ctx.fillRect(180, 50, 75, 50);
func (el *TagCanvas) GlobalCompositeOperation(value CompositeOperationsRule) (ref *TagCanvas) {
	el.context.Set("globalCompositeOperation", value.String())
	return el
}

// en: Returns the height of an ImageData object
//
//     The height property returns the height of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.height;
func (el *TagCanvas) Height() (height int) {
	return el.context.Get("height").Int()
}

// en: Returns true if the specified point is in the current path, otherwise false
//     x: The x-axis coordinate of the point to check.
//     y: The y-axis coordinate of the point to check.
//     fillRule: The algorithm by which to determine if a point is inside or outside the path.
//          "nonzero": The non-zero winding rule. Default rule.
//          "evenodd": The even-odd winding rule.
//     path: A Path2D path to check against. If unspecified, the current path is used.
//
//    Example:
//    var c = document.getElementById("myCanvas");
//    var ctx = c.getContext("2d");
//    ctx.rect(20, 20, 150, 100);
//    if (ctx.isPointInPath(20, 50)) {
//      ctx.stroke();
//    };
func (el *TagCanvas) IsPointInPath(path js.Value, x, y int, fillRule FillRule) (isPointInPath bool) {
	return el.context.Call("isPointInPath", path, x, y, fillRule.String()).Bool()
}

// en: Sets or returns the style of the end caps for a line
//     PlatformBasicType: "butt|round|square"
//
//     The lineCap property sets or returns the style of the end caps for a line.
//     Note: The value "round" and "square" make the lines slightly longer.
//
//     Default value: butt
//     JavaScript syntax: context.lineCap = "butt|round|square";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.lineCap = "round";
//     ctx.moveTo(20, 20);
//     ctx.lineTo(20, 200);
//     ctx.stroke();
func (el *TagCanvas) LineCap(value CanvasCapRule) (ref *TagCanvas) {
	el.context.Set("lineCap", value.String())
	return el
}

// en: Sets or returns the type of corner created, when two lines meet
//     PlatformBasicType: "bevel|round|miter"
//
//     The lineJoin property sets or returns the type of corner created, when two lines meet.
//     Note: The "miter" value is affected by the miterLimit property.
//     Default value:	miter
//     JavaScript syntax:	context.lineJoin = "bevel|round|miter";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.lineJoin = "round";
//     ctx.moveTo(20, 20);
//     ctx.lineTo(100, 50);
//     ctx.lineTo(20, 100);
//     ctx.stroke();
func (el *TagCanvas) LineJoin(value CanvasJoinRule) (ref *TagCanvas) {
	el.context.Set("lineJoin", value.String())
	return el
}

// Stroke
// en: The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
//
// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e lineTo() usando a cor padrão, preta.
//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um gradiente
func (el *TagCanvas) Stroke() (ref *TagCanvas) {
	el.context.Call("stroke")
	return el
}

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     The moveTo() method moves the path to the specified point in the canvas, without creating a line.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *TagCanvas) MoveTo(x, y int) (ref *TagCanvas) {
	el.context.Call("moveTo", x, y)
	return el
}

// en: Adds a new point and creates a line from that point to the last specified point in the canvas
//     x: The x-coordinate of where to create the line to
//     y: The y-coordinate of where to create the line to
//     The lineTo() method adds a new point and creates a line from that point to the last specified point in the canvas
//     (this method does not draw the line).
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *TagCanvas) LineTo(x, y int) (ref *TagCanvas) {
	el.context.Call("lineTo", x, y)
	return el
}

// en: Sets or returns the current line width
//     PlatformBasicType: The current line width, in pixels
//
//     The lineWidth property sets or returns the current line width, in pixels.
//     Default value: 1
//     JavaScript syntax: context.lineWidth = number;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.strokeRect(20, 20, 80, 100);
func (el *TagCanvas) LineWidth(value int) (ref *TagCanvas) {
	el.context.Set("lineWidth", value)
	return el
}

// en: Returns an object that contains the width of the specified text
//     text: The text to be measured
//
//     The measureText() method returns an object that contains the width of the specified text, in pixels.
//     Tip: Use this method if you need to know the width of a text, before writing it on the canvas.
//     JavaScript syntax: context.measureText(text).width;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "30px Arial";
//     var txt = "Hello World"
//     ctx.fillText("width:" + ctx.measureText(txt).width, 10, 50)
//     ctx.fillText(txt, 10, 100);
func (el *TagCanvas) MeasureText(text string) (ref *TagCanvas) {
	el.context.Call("measureText", text)
	return el
}

// en: Sets or returns the maximum miter length
//     PlatformBasicType: A positive number that specifies the maximum miter length. If the current miter length exceeds the
//            miterLimit, the corner will display as lineJoin "bevel"
//
//     The miterLimit property sets or returns the maximum miter length.
//     The miter length is the distance between the inner corner and the outer corner where two lines meet.
//
//     Default value: 10
//     JavaScript syntax: context.miterLimit = number;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.lineJoin = "miter";
//     ctx.miterLimit = 5;
//     ctx.moveTo(20, 20);
//     ctx.lineTo(50, 27);
//     ctx.lineTo(20, 34);
//     ctx.stroke();
func (el *TagCanvas) MiterLimit(value int) (ref *TagCanvas) {
	el.context.Set("miterLimit", value)
	return el
}

// en: Creates a quadratic Bézier curve
//     cpx: The x-axis coordinate of the control point.
//     cpy: The y-axis coordinate of the control point.
//     x:   The x-axis coordinate of the end point.
//     y:   The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.quadraticCurveTo(20, 100, 200, 20);
//     ctx.stroke();
func (el *TagCanvas) QuadraticCurveTo(cpx, cpy, x, y int) (ref *TagCanvas) {
	el.context.Call("quadraticCurveTo", cpx, cpy, x, y)
	return el
}

// en: Creates a rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The rect() method creates a rectangle.
//     Tip: Use the stroke() or the fill() method to actually draw the rectangle on the canvas.
//     JavaScript syntax: context.rect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.rect(20, 20, 150, 100);
//     ctx.stroke();
func (el *TagCanvas) Rect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("rect", x, y, width, height)
	return el
}

// en: Returns previously saved path state and attributes
func (el *TagCanvas) Restore() (ref *TagCanvas) {
	el.selfElement.Call("restore")
	return el
}

// en: Rotates the current drawing
//     angle: The rotation angle, in radians.
//            To calculate from degrees to radians: degrees*Math.PI/180.
//            Example: to rotate 5 degrees, specify the following: 5*Math.PI/180
//
//     The rotate() method rotates the current drawing.
//     Note: The rotation will only affect drawings made AFTER the rotation is done.
//     JavaScript syntax: context.rotate(angle);
//
//     Example:
//     var c = document.getElementById("my Canvas");
//     var ctx = c.getContext("2d");
//     ctx.rotate(20 * Math.PI / 180);
//     ctx.fillRect(50, 20, 100, 50);
func (el *TagCanvas) Rotate(angle float64) (ref *TagCanvas) {
	el.context.Call("rotate", angle)
	return el
}

// en: Saves the state of the current context
func (el *TagCanvas) Save() (ref *TagCanvas) {
	el.selfElement.Call("save")
	return el
}

// en: Scales the current drawing bigger or smaller
//     scaleWidth:  Scales the width of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//     scaleHeight: Scales the height of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//
//     The scale() method scales the current drawing, bigger or smaller.
//     Note: If you scale a drawing, all future drawings will also be scaled. The positioning will also be scaled. If
//     you scale(2,2); drawings will be positioned twice as far from the left and top of the canvas as you specify.
//     JavaScript syntax: context.scale(scalewidth, scaleheight);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.strokeRect(5, 5, 25, 15);
//     ctx.scale(2, 2);
//     ctx.strokeRect(5, 5, 25, 15);
func (el *TagCanvas) Scale(scaleWidth, scaleHeight int) (ref *TagCanvas) {
	el.context.Call("scale", scaleWidth, scaleHeight)
	return
}

// en: Resets the current transform to the identity matrix. Then runs transform()
//     a: Scales the drawings horizontally
//     b: Skews the drawings horizontally
//     c: Skews the drawings vertically
//     d: Scales the drawings vertically
//     e: Moves the the drawings horizontally
//     f: Moves the the drawings vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The setTransform() method resets the current transform to the identity matrix, and then runs transform() with the
//     same arguments.
//     In other words, the setTransform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the setTransform method is called.
//     JavaScript syntax: context.setTransform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *TagCanvas) SetTransform(a, b, c, d, e, f float64) (ref *TagCanvas) {
	el.context.Call("setTransform", a, b, c, d, e, f)
	return el
}

// en: Draws a rectangle (no fill)
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The strokeRect() method draws a rectangle (no fill). The default color of the stroke is black.
//     Tip: Use the strokeStyle property to set a color, gradient, or pattern to style the stroke.
//     JavaScript syntax: context.strokeRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.strokeRect(20, 20, 150, 100);
func (el *TagCanvas) StrokeRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("strokeRect", x, y, width, height)
	return el
}

// en: Sets or returns the color, gradient, or pattern used for strokes
//     The strokeStyle property sets or returns the color, gradient, or pattern used for strokes.
//     Default value: #000000
//     JavaScript syntax: context.strokeStyle = color|gradient|pattern;
func (el *TagCanvas) StrokeStyle(value string) (ref *TagCanvas) {
	el.context.Set("strokeStyle", value)
	return el
}

// en: Draws text on the canvas (no fill)
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The strokeText() method draws text (with no fill) on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the strokeStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.strokeText(text, x, y, maxWidth);
//
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.strokeText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.strokeStyle = gradient;
//     ctx.strokeText("Big smile!", 10, 90);
func (el *TagCanvas) StrokeText(text string, x, y, maxWidth int) (ref *TagCanvas) {
	el.context.Call("strokeText", text, x, y, maxWidth)
	return el
}

// en: Sets or returns the current alignment for text content
//
//     The textAlign property sets or returns the current alignment for text content, according to the anchor point.
//     Normally, the text will START in the position specified, however, if you set textAlign="right" and place the text in position 150, it means that the text should END in position 150.
//     Tip: Use the fillText() or the strokeText() method to actually draw and position the text on the canvas.
//     Default value: start
//     JavaScript syntax: context.textAlign = "center | end | left | right | start";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     // Create a red line in position 150
//     ctx.strokeStyle = "red";
//     ctx.moveTo(150, 20);
//     ctx.lineTo(150, 170);
//     ctx.stroke();
//     ctx.font = "15px Arial";
//     // Show the different textAlign values
//     ctx.textAlign = "start";
//     ctx.fillText("textAlign = start", 150, 60);
//     ctx.textAlign = "end";
//     ctx.fillText("textAlign = end", 150, 80);
//     ctx.textAlign = "left";
//     ctx.fillText("textAlign = left", 150, 100);
//     ctx.textAlign = "center";
//     ctx.fillText("textAlign = center", 150, 120);
//     ctx.textAlign = "right";
//     ctx.fillText("textAlign = right", 150, 140);
func (el *TagCanvas) TextAlign(value FontAlignRule) (ref *TagCanvas) {
	el.context.Set("textAlign", value.String())
	return el
}

// en: Sets or returns the current text baseline used when drawing text
//     PlatformBasicType:
//          alphabetic:  Default. The text baseline is the normal alphabetic baseline
//          top:         The text baseline is the top of the em square
//          hanging:     The text baseline is the hanging baseline
//          middle:      The text baseline is the middle of the em square
//          ideographic: The text baseline is the ideographic baseline
//          bottom:      The text baseline is the bottom of the bounding box
//
//     The textBaseline property sets or returns the current text baseline used when drawing text.
//     Note: The fillText() and strokeText() methods will use the specified textBaseline value when positioning the text
//     on the canvas.
//     Default value: alphabetic
//     JavaScript syntax: context.textBaseline = "alphabetic|top|hanging|middle|ideographic|bottom";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     //Draw a red line at y=100
//     ctx.strokeStyle = "red";
//     ctx.moveTo(5, 100);
//     ctx.lineTo(395, 100);
//     ctx.stroke();
//     ctx.font = "20px Arial"
//     //Place each word at y=100 with different textBaseline values
//     ctx.textBaseline = "top";
//     ctx.fillText("Top", 5, 100);
//     ctx.textBaseline = "bottom";
//     ctx.fillText("Bottom", 50, 100);
//     ctx.textBaseline = "middle";
//     ctx.fillText("Middle", 120, 100);
//     ctx.textBaseline = "alphabetic";
//     ctx.fillText("Alphabetic", 190, 100);
//     ctx.textBaseline = "hanging";
//     ctx.fillText("Hanging", 290, 100);
func (el *TagCanvas) TextBaseline(value TextBaseLineRule) (ref *TagCanvas) {
	el.context.Set("textBaseline", value.String())
	return el
}

// en: Replaces the current transformation matrix for the drawing
//     a: Scales the drawing horizontally
//     b: Skew the the drawing horizontally
//     c: Skew the the drawing vertically
//     d: Scales the drawing vertically
//     e: Moves the the drawing horizontally
//     f: Moves the the drawing vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The transform() method replaces the current transformation matrix. It multiplies the current transformation
//     matrix with the matrix described by:
//
//     a | c | e
//    -----------
//     b | d | f
//    -----------
//     0 | 0 | 1
//
//     In other words, the transform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the transform() method is called.
//     Note: The transform() method behaves relatively to other transformations made by rotate(), scale(), translate(),
//     or transform(). Example: If you already have set your drawing to scale by two, and the transform() method scales
//     your drawings by two, your drawings will now scale by four.
//     Tip: Check out the setTransform() method, which does not behave relatively to other transformations.
//     JavaScript syntax: context.transform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *TagCanvas) Transform(a, b, c, d, e, f float64) (ref *TagCanvas) {
	el.context.Call("transform", a, b, c, d, e, f)
	return el
}

// en: Remaps the (0,0) position on the canvas
//     x: The value to add to horizontal (x) coordinates
//     y: The value to add to vertical (y) coordinates
//
//     The translate() method remaps the (0,0) position on the canvas.
//     Note: When you call a method such as fillRect() after translate(), the value is added to the x- and y-coordinate
//     values.
//     JavaScript syntax: context.translate(x, y);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillRect(10, 10, 100, 50);
//     ctx.translate(70, 70);
//     ctx.fillRect(10, 10, 100, 50);
func (el *TagCanvas) Translate(x, y int) (ref *TagCanvas) {
	el.context.Call("translate", x, y)
	return el
}

// en: Sets or returns the color, gradient, or pattern used to fill the drawing
//     The fillStyle property sets or returns the color, gradient, or pattern used to fill the drawing.
//     Default value:	#000000
//     JavaScript syntax:	context.fillStyle = color|gradient|pattern;
func (el *TagCanvas) FillStyle(value string) (ref *TagCanvas) {
	el.context.Set("fillStyle", value)
	return el
}

// en: Fills the current drawing (path)
//     The fill() method fills the current drawing (path). The default color is black.
//     Tip: Use the fillStyle property to fill with another color/gradient.
//     Note: If the path is not closed, the fill() method will add a line from the last point to the startpoint of the
//     path to close the path (like closePath()), and then fill the path.
func (el *TagCanvas) Fill() (ref *TagCanvas) {
	el.context.Call("fill")
	return el
}
