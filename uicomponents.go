package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func setUpThemeListBox(settings *gtk.Settings, currentTheme string) *gtk.ListBox {
	listBox, _ := gtk.ListBoxNew()
	var rowToSelect *gtk.ListBoxRow

	for _, name := range themeNames {
		row, _ := gtk.ListBoxRowNew()

		eventBox, _ := gtk.EventBoxNew()
		box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
		eventBox.Add(box)

		lbl, _ := gtk.LabelNew(name)
		n := name
		eventBox.Connect("button-press-event", func() {
			settings.SetProperty("gtk-theme-name", n)
			fmt.Println(n)
		})
		if n == currentTheme {
			rowToSelect = row
		}

		box.PackStart(lbl, false, false, 0)

		row.Add(eventBox)
		fmt.Println(name)
		listBox.Add(row)
	}
	listBox.SelectRow(rowToSelect)

	return listBox
}

func setUpWidgetsPreview() *gtk.Frame {
	frame, _ := gtk.FrameNew("Preview")

	grid, _ := gtk.GridNew()
	grid.SetRowSpacing(6)
	grid.SetColumnSpacing(12)
	for _, prop := range margins {
		grid.SetProperty(prop, 6)
	}
	frame.Add(grid)

	checkButton, _ := gtk.CheckButtonNew()
	checkButton.SetLabel("Check Button")
	grid.Attach(checkButton, 0, 0, 1, 1)

	radioButton, _ := gtk.RadioButtonNew(nil)
	radioButton.SetLabel("Radio Button")
	grid.Attach(radioButton, 0, 1, 1, 1)

	spinButton, _ := gtk.SpinButtonNewWithRange(0, 1000, 10)
	grid.Attach(spinButton, 0, 2, 1, 1)

	button, _ := gtk.ButtonNewFromIconName("search", gtk.ICON_SIZE_BUTTON)
	button.SetLabel("Button")
	grid.Attach(button, 1, 2, 1, 1)

	scale, _ := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 0, 100, 1)
	scale.SetSizeRequest(200, 0)
	scale.SetDrawValue(true)
	scale.SetValue(60)
	grid.Attach(scale, 2, 0, 1, 1)

	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	grid.Attach(separator, 2, 1, 1, 1)

	combo, _ := gtk.ComboBoxTextNew()
	combo.Append("entry #1", "entry #1")
	combo.Append("entry #2", "entry #2")
	grid.Attach(combo, 2, 2, 1, 1)

	return frame
}
