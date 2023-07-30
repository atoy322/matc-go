package matc

/*
#cgo LDFLAGS: -lmatc
#include <display.h>
#include <stdlib.h>
*/
import "C"
import (
    "unsafe"
)


func Display(mat Matrix) {
    C.matcDisplay(to_cmat(mat))
}

func Displayf(format string, mat Matrix) {
    cformat := C.CString(format)
    defer C.free(unsafe.Pointer(cformat))

    C.matcDisplayf(cformat, to_cmat(mat))
}
