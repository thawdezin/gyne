package main

import (
    "os"
    "github.com/therecipe/qt/widgets"
)

func main() {
    // Initialize the application
    app := widgets.NewQApplication(len(os.Args), os.Args)

    // Create a main window
    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("Go Qt GUI Example")

    // Create a central widget for the main window
    centralWidget := widgets.NewQWidget(nil, 0)
    centralWidget.SetLayout(widgets.NewQVBoxLayout())

    // Create a button
    button := widgets.NewQPushButton2("Click Me", nil)
    button.ConnectClicked(func(checked bool) {
        widgets.QMessageBox_Information(nil, "Message", "Hello, Go Qt GUI!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
    })

    // Add the button to the central widget
    centralWidget.Layout().AddWidget(button)

    // Set the central widget for the main window
    window.SetCentralWidget(centralWidget)

    // Show the main window
    window.Show()

    // Enter the main event loop
    app.Exec()
}


// package main

// import (
//     "fyne.io/fyne/v2/app"
//     "fyne.io/fyne/v2/container"
//     "fyne.io/fyne/v2/widget"
// )

// func main() {
//     // Create a new Fyne application
//     myApp := app.New()

//     // Create a new window
//     myWindow := myApp.NewWindow("Fyne Example")

//     // Create a label and a button
//     myLabel := widget.NewLabel("Hello, Fyne!")
//     myButton := widget.NewButton("Click Me", func() {
//         myLabel.SetText("Button Clicked!")
//     })

//     // Create a container to hold the label and button
//     myContainer := container.NewVBox(myLabel, myButton)

//     // Set the container as the window's content
//     myWindow.SetContent(myContainer)

//     // Show the window
//     myWindow.ShowAndRun()
// }
