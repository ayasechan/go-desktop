package desktop

import "github.com/go-ole/go-ole"

func hrToErr(hr uintptr) error {
	if hr > 0 {
		return ole.NewError(hr)
	}
	return nil
}
