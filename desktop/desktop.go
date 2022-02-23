package desktop

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"golang.org/x/xerrors"
)

//go:generate go run github.com/ayasechan/genole -input def.yaml -output def.go -pkg desktop

var ErrNotImplement = xerrors.New("NotImplement")
var CLSID_ImmersiveShell = ole.NewGUID("{C2F03A33-21F5-47FA-B4BB-156362A2F239}")
var CLSID_VirtualDesktopPinnedApps = ole.NewGUID("{B5A399E7-1C87-46B8-88E9-FC5747B171BD}")

func QueryService(clsid, iid *ole.GUID, ppvObject uintptr) error {
	unk, err := ole.CreateInstance(CLSID_ImmersiveShell, IID_IServiceProvider)
	if err != nil {
		return xerrors.Errorf("%v", err)
	}
	defer unk.Release()

	provider := (*IServiceProvider)(unsafe.Pointer(unk))
	err = provider.QueryService(clsid, iid, ppvObject)
	return err
}

// call Release to recycle object
func NewIVirtualDesktopManagerInternal() (*IVirtualDesktopManagerInternal, error) {
	var obj *IVirtualDesktopManagerInternal
	err := QueryService(
		CLSID_VirtualDesktopManagerInternal,
		IID_IVirtualDesktopManagerInternal,
		uintptr(unsafe.Pointer(&obj)),
	)
	return obj, err
}

// call Release to recycle object
func NewIVirtualDesktopManager() (*IVirtualDesktopManager, error) {
	var obj *IVirtualDesktopManager
	err := QueryService(
		IID_IVirtualDesktopManager,
		IID_IVirtualDesktopManager,
		uintptr(unsafe.Pointer(&obj)),
	)
	return obj, err
}

// call Release to recycle object
func NewIApplicationViewCollection() (*IApplicationViewCollection, error) {
	var obj *IApplicationViewCollection
	err := QueryService(
		IID_IApplicationViewCollection,
		IID_IApplicationViewCollection,
		uintptr(unsafe.Pointer(&obj)),
	)
	return obj, err
}

// call Release to recycle object
func NewIVirtualDesktopPinnedApps() (*IVirtualDesktopPinnedApps, error) {
	var obj *IVirtualDesktopPinnedApps
	err := QueryService(
		CLSID_VirtualDesktopPinnedApps,
		IID_IVirtualDesktopPinnedApps,
		uintptr(unsafe.Pointer(&obj)),
	)
	return obj, err
}
