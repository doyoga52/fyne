package dialog

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func createTextDialog(title, message string, icon fyne.Resource, parent fyne.Window) Dialog {
	return createTextDialogWithCallback(title, message, icon, nil, parent)
}

func createTextDialogWithCallback(title, message string, icon fyne.Resource, callback func(bool), parent fyne.Window) Dialog {
	d := newDialog(title, message, icon, callback, parent)

	d.dismiss = &widget.Button{Text: "OK",
		OnTapped: d.Hide,
	}
	d.setButtons(newButtonList(d.dismiss))

	return d
}

// NewInformation creates a dialog over the specified window for user information.
// The title is used for the dialog window and message is the content.
// After creation you should call Show().
func NewInformation(title, message string, parent fyne.Window) Dialog {
	return createTextDialog(title, message, theme.InfoIcon(), parent)
}

// The title is used for the dialog window and message is the content.
// After creation you should call Show().
func NewInformationWithCallback(title, message string, callback func(bool), parent fyne.Window) Dialog {
	return createTextDialogWithCallback(title, message, theme.InfoIcon(), callback, parent)
}

// ShowInformation shows a dialog over the specified window for user
// information. The title is used for the dialog window and message is the content.
func ShowInformation(title, message string, parent fyne.Window) {
	NewInformation(title, message, parent).Show()
}

// ShowInformation shows a dialog over the specified window for user
// information. The title is used for the dialog window and message is the content.
func ShowInformationWithCallback(title, message string, callback func(bool), parent fyne.Window) {
	NewInformationWithCallback(title, message, callback, parent).Show()
}

// ShowError shows a dialog over the specified window for an application
// error. The title and message are extracted from the provided error.
func ShowError(err error, parent fyne.Window) {
	createTextDialog("Error", err.Error(), theme.WarningIcon(), parent).Show()
}

// ShowError shows a dialog over the specified window for an application
// error. The title and message are extracted from the provided error.
func ShowErrorWithCallback(err error, callback func(bool), parent fyne.Window) {
	createTextDialogWithCallback("Error", err.Error(), theme.WarningIcon(), callback, parent).Show()
}
