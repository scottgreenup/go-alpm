package alpm

// #include <alpm.h>
import "C"

type Error struct {
    alpmError C.alpm_errno_t
    message   string
}

func NewError(err C.alpm_errno_t) error {
    return Error{
        alpmError: err,
        message: C.GoString(C.alpm_strerror(err)),
    }
}

func (e Error) Error() string {
    return e.message
}
