package main

import "github.com/Tim-Paik/webview"
import _ "github.com/Tim-Paik/webview/manifest"

func main() {
	w := webview.New(false)
	if w == nil {
		println("Failed to load webview.")
		return
	}
	defer w.Destroy()
	w.SetTitle(`Minimal webview example`)
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate(`https://www.google.com`)
	w.Run()
}
