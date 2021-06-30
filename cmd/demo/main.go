package main

import "github.com/Tim-Paik/webview"

func main() {
	debug := true
	w := webview.New(debug)
	if w == nil {
		println("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://bing.com")
	w.Run()
}
