package desktop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/gonutz/w32/v2"
)

var IID_IVirtualDesktop = func() *ole.GUID {
	iid := "{FF72FFDD-BE7E-43FC-9C03-AD81681E88E4}"
	if w32.RtlGetVersion().BuildNumber >= 20231 {
		iid = "{62FDF88B-11CA-4AFB-8BD8-2296DFAE49E2}"
	}
	return ole.NewGUID(iid)
}()

type IVirtualDesktop struct {
	ole.IUnknown
}

type IVirtualDesktopVtbl struct {
	ole.IUnknownVtbl
	IsViewVisible uintptr
	GetID         uintptr
}

func (v *IVirtualDesktop) VTable() *IVirtualDesktopVtbl {
	return (*IVirtualDesktopVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IVirtualDesktop) IsViewVisible(pView *IApplicationView, pfVisible *int) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().IsViewVisible,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pView)), // TODO 难道不是双指针吗
		uintptr(unsafe.Pointer(pfVisible)),
	)

	return hrToErr(hr)
}

func (v *IVirtualDesktop) GetID() (ole.GUID, error) {
	var id ole.GUID
	hr, _, _ := syscall.Syscall(
		v.VTable().GetID,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&id)),
		0,
	)

	return id, hrToErr(hr)
}
