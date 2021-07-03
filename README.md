# webview
This is a fork of [go-webview2](https://github.com/jchv/go-webview2)

**No EdgeHTML fallback**

### Added some functions I need:
1. High DPI support
2. Remove the blank icon in the title bar (because I don't know how to modify the icon, if you know, please open an Issue)
3. Prompt to download WebView2 runtime (based on [dlgs](https://github.com/gen2brain/dlgs))
4. Implemented disable debugging in non-debug mode:  prevent default context menus, disable DevTools, disable built in error page, disable status bar and disable zoom control.

### Tips
If you want to use a more modern dialog box, please import the following package to use comctl6 (thanks to andlabs' [winmanifest](https://github.com/andlabs/ui/tree/master/winmanifest))
```go
package main
import _ "github.com/Tim-Paik/webview2/manifest"
```

### Problem:
The ARM64 architecture is currently not supported, because I don't have an ARM64 device, and the Microsoft documentation does not mention the location of the Webview2 registry under the ARM64 architecture.

### TODO
 - [x] High DPI support
 - [x] Clean title bar
 - [x] Prompt to download WebView2 runtime
 - [x] Non-debug mode
 - [ ] Bindings

> ### go-webview2
> This is a proof of concept for embedding Webview2 into Go without CGo. It is based on [webview/webview](https://github.com/webview/webview) and provides a compatible API sans some unimplemented functionality (notably, bindings are not implemented.)
>
> #### Notice
> Because this version doesn't currently have an EdgeHTML fallback, it will not work unless you have a Webview2 runtime installed. In addition, it requires the Webview2Loader DLL in order to function. Adding an EdgeHTML fallback should be technically possible but will likely require much worse hacks since the API is not strictly COM to my knowledge.
>
> #### Demo
> For now, you'll need to install the Webview2 runtime, as it does not ship with Windows.
>
> [WebView2 runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)
>
> After that, you should be able to run go-webview2 directly:
>
> ```
> go run github.com/jchv/go-webview2/cmd/demo
> ```
>
> This will use go-winloader to load an embedded copy of WebView2Loader.dll.
>
> If this does not work, please try running from a directory that has an appropriate copy of `WebView2Loader.dll` for your GOARCH. If _that_ worked, *please* file a bug so we can figure out what's wrong with go-winloader :)
