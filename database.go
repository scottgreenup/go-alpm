package alpm

// #include <alpm.h>
import "C"

import (
    "unsafe"
)

type Database struct {
    ref    *C.alpm_db_t
    handle *Handle
}

func NewDatabase(handle *Handle) (*Database, error) {
    database := C.alpm_get_localdb(handle.Ref())

    if database == nil {
        err := C.alpm_errno(handle.Ref())
        return nil, NewError(err)
    }

    return &Database{
        ref: database,
        handle: handle,
    }, nil
}

// alpm_list_t
type list struct {
    Data unsafe.Pointer
    Prev *list
    Next *list
}

// alpm_pkg_t
type Package struct {
    ref *C.alpm_pkg_t
}

func (p *Package) Name() string {
    return C.GoString(C.alpm_pkg_get_name(p.ref))
}

func (d *Database) packageList() (*list, error) {
    packageCache := C.alpm_db_get_pkgcache(d.ref)
    if packageCache == nil {
        err := C.alpm_errno(d.handle.Ref())
        return nil, NewError(err)
    }

    return (*list)(unsafe.Pointer(packageCache)), nil
}

func (d *Database) Packages() ([]Package, error) {
    packageList, err := d.packageList()
    if err != nil {
        return nil, err
    }

    packages := []Package{}

    for packageList != nil {
        var packageRef *C.alpm_pkg_t
        packageRef =  (*C.alpm_pkg_t)(packageList.Data)

        packages = append(packages, Package{
            ref: packageRef,
        })

        packageList = packageList.Next
    }

    return packages, nil
}
