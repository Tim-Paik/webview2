// +build windows

package webview2

import (
	"github.com/gen2brain/dlgs"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func GetWebview2Runtime() error {
	willDownload, err := dlgs.Question(`Require Microsoft Webview2 Runtime`,
		`Missing Microsoft Webview2 Runtime. 
Do you want to download Microsoft Webview2 Runtime now? 
The program will exit.`, false)
	if err != nil {
		return err
	}
	if willDownload {
		cmd := exec.Command(`cmd`, `/c`, `start`, `https://go.microsoft.com/fwlink/p/?LinkId=2124703`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return cmd.Start()
	}
	return err
}

func checkRuntime(err error) {
	if err == nil {
		return
	}
	if err != registry.ErrNotExist {
		dlgs.Error(`Microsoft Webview2 Runtime`, `Webview2 Runtime Error: `+err.Error())
	} else {
		if err := GetWebview2Runtime(); err != nil {
			dlgs.Error(`Microsoft Webview2 Runtime`, `Get Webview2 Runtime Error: `+err.Error())
		}
	}
	os.Exit(1)
}

func init() {
	// Enable High Dpi Support
	windows.NewLazySystemDLL("Shcore").NewProc("SetProcessDpiAwareness").Call(1)

	var key registry.Key
	var err error = nil
	switch runtime.GOARCH {
	case "amd64":
		key, err = registry.OpenKey(registry.LOCAL_MACHINE,
			`SOFTWARE\WOW6432Node\Microsoft\EdgeUpdate\Clients\{F3017226-FE2A-4295-8BDF-00C3A9A7E4C5}`,
			registry.READ)
	case "386":
		key, err = registry.OpenKey(registry.LOCAL_MACHINE,
			`SOFTWARE\Microsoft\EdgeUpdate\Clients\{F3017226-FE2A-4295-8BDF-00C3A9A7E4C5}`,
			registry.READ)
	default:
		return
	}
	defer key.Close()
	checkRuntime(err)
	_, _, err = key.GetStringValue(`pv`)
	checkRuntime(err)
}
