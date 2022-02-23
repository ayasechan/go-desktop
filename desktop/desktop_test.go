package desktop

import (
	"testing"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/gonutz/w32/v2"
	"github.com/stretchr/testify/assert"
)

func TestIVirtualDesktopManager(t *testing.T) {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	defer ole.CoUninitialize()

	manager, err := NewIVirtualDesktopManager()
	assert.Nil(t, err)
	hwnd := w32.GetForegroundWindow()
	t.Log("GetForegroundWindow:", hwnd)

	ok, err := manager.IsWindowOnCurrentVirtualDesktop(uintptr(hwnd))
	assert.Nil(t, err)
	assert.True(t, ok)

	id, err := manager.GetWindowDesktopId(uintptr(hwnd))
	assert.Nil(t, err)
	t.Log("GetWindowDesktopId:", id.String())

	err = manager.MoveWindowToDesktop(uintptr(hwnd), id)
	assert.Nil(t, err)

}
func TestIVirtualDesktopManagerInternal(t *testing.T) {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	defer ole.CoUninitialize()

	manager, err := NewIVirtualDesktopManagerInternal()
	assert.Nil(t, err)

	desktops, err := manager.GetDesktops()
	assert.Nil(t, err)
	curDesktop, err := manager.GetCurrentDesktop()
	assert.Nil(t, err)
	curDesktopId, err := curDesktop.GetID()
	assert.Nil(t, err)
	t.Log("curDesktop.GetID:", curDesktopId.String())

	for _, desktop := range desktops {
		id, err := desktop.GetID()
		assert.Nil(t, err)
		t.Log("desktop id:", id.String())
	}

	newDesktop, err := manager.CreateDesktopW()
	assert.Nil(t, err)
	time.Sleep(time.Second)
	err = manager.SwitchDesktop(newDesktop)
	assert.Nil(t, err)
	time.Sleep(time.Second)
	err = manager.RemoveDesktop(newDesktop, curDesktop)
	assert.Nil(t, err)

}
