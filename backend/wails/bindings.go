package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Bindings provides helper methods exposed to the frontend.
type Bindings struct {
	ctx context.Context
}

// NewBindings creates a new Bindings instance.
func NewBindings() *Bindings {
	return &Bindings{}
}

// startup is called by Wails when the app starts.
func (b *Bindings) startup(ctx context.Context) {
	b.ctx = ctx
}

// OpenURL opens the provided URL in the default browser.
func (b *Bindings) OpenURL(url string) {
	runtime.BrowserOpenURL(b.ctx, url)
}

// SelectFolder displays a folder selection dialog and returns the chosen path.
func (b *Bindings) SelectFolder() (string, error) {
	return runtime.OpenDirectoryDialog(b.ctx, runtime.OpenDialogOptions{})
}
