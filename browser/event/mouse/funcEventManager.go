package mouse

import "syscall/js"

// EventManager
//
// English:
//
// Capture event information and format to Golang
//
//   Output:
//     data: list with all the information provided by the browser.
//
// Português:
//
// Captura as informações do evento e formata para o Golang
//
//   Saída:
//     data: lista com todas as informações fornecidas pelo navegador.
func EventManager(this js.Value, args []js.Value) (data Data) {
	var event = Event{}
	event.Object = args[0]

	data.ClientX = event.GetClientX()
	data.ClientY = event.GetClientY()
	data.MovementX = event.GetMovementX()
	data.MovementY = event.GetMovementY()
	data.OffsetX = event.GetOffsetX()
	data.OffsetY = event.GetOffsetY()
	data.PageX = event.GetPageX()
	data.PageY = event.GetPageY()
	data.ScreenX = event.GetScreenX()
	data.ScreenY = event.GetScreenY()
	data.X = event.GetX()
	data.Y = event.GetY()
	data.RelatedTarget = event.GetRelatedTarget()
	data.Region = event.GetRegion()
	data.Button = event.GetButton()
	data.AltKey = event.GetAltKey()
	data.ShiftKey = event.GetShiftKey()
	data.MetaKey = event.GetMetaKey()
	data.CtrlKey = event.GetCtrlKey()
	data.This = this

	return
}
