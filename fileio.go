package matc

/*
#cgo CFLAGS: -I./matc/libmatc/include
#cgo LDFLAGS: -L./matc/libmatc/lib -lmatc
#include <fileio.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Dump(filename string, mat Matrix, name rune) int {
    cfilename := C.CString(filename)
    mode := C.CString("wb")
    defer C.free(unsafe.Pointer(cfilename))
    defer C.free(unsafe.Pointer(mode))

    cfp := C.fopen(cfilename, mode)
    if cfp == nil {
        fmt.Println("cfp == nil")
        return -1
    }
    defer C.fclose(cfp)

    return int(C.matcDump(cfp, to_cmat(mat), C.char(name)))
}

func Load(filename string, mat Matrix, name rune) int {
    cfilename := C.CString(filename)
    mode := C.CString("rb")
    defer C.free(unsafe.Pointer(cfilename))
    defer C.free(unsafe.Pointer(mode))

    cfp := C.fopen(cfilename, mode)
    if cfp == nil {
        fmt.Println("cfp == nil")
        return -1
    }
    defer C.fclose(cfp)

    return int(C.matcLoad(cfp, to_cmat(mat), C.char(name)))
}
