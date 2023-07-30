package matc

import (
	"fmt"
	"testing"
)

func TestDisplay(t *testing.T) {
    mat := Init(4, 4)
    Eye(4, mat)
    Displayf("%2.0f", mat)

    fmt.Println(mat.Get(0, 0))

    //Dump("matrix.mat", mat, 'A')
    Load("matrix.mat", mat, 'A')

    Display(mat)

    var det float64
    Det(mat, &det)
    fmt.Println(det)

    Rref(mat)
    Display(mat)

    Deinit(mat)
}
