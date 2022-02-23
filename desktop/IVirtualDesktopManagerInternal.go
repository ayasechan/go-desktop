package desktop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/gonutz/w32/v2"
)

var (
	CLSID_VirtualDesktopManagerInternal = ole.NewGUID("{C5E0CDCA-7B6E-41B2-9FC4-D93975CC467B}")
	IID_IVirtualDesktopManagerInternal  = func() *ole.GUID {
		iid := "{F31574D6-B682-4CDC-BD56-1827860ABEC6}"
		if w32.RtlGetVersion().BuildNumber >= 20231 {
			iid = "{094AFE11-44F2-4BA0-976F-29A97E263EE0}"
		}
		return ole.NewGUID(iid)
	}()
)

type IVirtualDesktopManagerInternal struct {
	ole.IUnknown
}

type IVirtualDesktopManagerInternalVtbl struct {
	ole.IUnknownVtbl
	GetCount            uintptr
	MoveViewToDesktop   uintptr
	CanViewMoveDesktops uintptr
	GetCurrentDesktop   uintptr
	GetDesktops         uintptr
	GetAdjacentDesktop  uintptr
	SwitchDesktop       uintptr
	CreateDesktopW      uintptr
	RemoveDesktop       uintptr
	FindDesktop         uintptr
}

func (v *IVirtualDesktopManagerInternal) VTable() *IVirtualDesktopManagerInternalVtbl {
	return (*IVirtualDesktopManagerInternalVtbl)(unsafe.Pointer(v.RawVTable))
}
func (v *IVirtualDesktopManagerInternal) GetCount() error { return ErrNotImplement }

// Since build 10240
func (v *IVirtualDesktopManagerInternal) MoveViewToDesktop() error   { return ErrNotImplement }
func (v *IVirtualDesktopManagerInternal) CanViewMoveDesktops() error { return ErrNotImplement }
func (v *IVirtualDesktopManagerInternal) GetCurrentDesktop() (*IVirtualDesktop, error) {
	var desktop *IVirtualDesktop
	hr, _, _ := syscall.Syscall(
		v.VTable().GetCurrentDesktop,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&desktop)),
		0,
	)
	return desktop, hrToErr(hr)
}
func (v *IVirtualDesktopManagerInternal) GetDesktops() ([]*IVirtualDesktop, error) {
	var pObjectArray *IObjectArray
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDesktops,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&pObjectArray)),
		0,
	)
	err := hrToErr(hr)
	if err != nil {
		return nil, err
	}
	count, err := pObjectArray.GetCount()
	if err != nil {
		return nil, err
	}
	desktops := make([]*IVirtualDesktop, 0, count)
	for i := 0; i < int(count); i++ {
		var desktop *IVirtualDesktop
		err = pObjectArray.GetAt(uint(i), IID_IVirtualDesktop, uintptr(unsafe.Pointer(&desktop)))
		if err != nil {
			return nil, err
		}
		desktops = append(desktops, desktop)
	}
	return desktops, nil
}

func (v *IVirtualDesktopManagerInternal) GetAdjacentDesktop() error { return ErrNotImplement }
func (v *IVirtualDesktopManagerInternal) SwitchDesktop(desktop *IVirtualDesktop) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().SwitchDesktop,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(desktop)),
		0,
	)
	return hrToErr(hr)
}
func (v *IVirtualDesktopManagerInternal) CreateDesktopW() (*IVirtualDesktop, error) {
	var desktop *IVirtualDesktop
	hr, _, _ := syscall.Syscall(
		v.VTable().CreateDesktopW,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&desktop)),
		0,
	)
	return desktop, hrToErr(hr)
}

// rm 要移除的桌面
// fallback 移除移动到该桌面
func (v *IVirtualDesktopManagerInternal) RemoveDesktop(rm, fallback *IVirtualDesktop) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().RemoveDesktop,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(rm)),
		uintptr(unsafe.Pointer(fallback)),
	)
	return hrToErr(hr)
}

// Since build 10240
func (v *IVirtualDesktopManagerInternal) FindDesktop(id ole.GUID) (*IVirtualDesktop, error) {
	var desktop *IVirtualDesktop
	hr, _, _ := syscall.Syscall(
		v.VTable().FindDesktop,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&id)),
		uintptr(unsafe.Pointer(&desktop)),
	)
	return desktop, hrToErr(hr)
}
