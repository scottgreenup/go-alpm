package alpm

// #cgo LDFLAGS: -lalpm
// #include <alpm.h>
// #include <stdlib.h>
import "C"
import (
    "unsafe"
)

type Handle struct {
    ref *C.alpm_handle_t
}

func NewHandle(root, dbpath string) (*Handle, error) {
    cRoot := C.CString("/")
    defer C.free(unsafe.Pointer(cRoot))
    cDBPath := C.CString("/var/lib/pacman")
    defer C.free(unsafe.Pointer(cDBPath))

    var err C.alpm_errno_t;

    handle := C.alpm_initialize(cRoot, cDBPath, &err)
    if err != 0 {
        return nil, NewError(err)
    }

    return &Handle{
        ref: handle,
    }, nil
}

func (h* Handle) Close() error {
    C.alpm_release(h.ref)
    return nil
}

func (h* Handle) Ref() *C.alpm_handle_t {
    return h.ref
}
