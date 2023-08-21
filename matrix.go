package matc

/*
#cgo LDFLAGS: -lmatc
#include <matrix.h>

double get(double **dpp, int i, int j) {
    return dpp[i][j];
}
void set(double **dpp, int i, int j, double x) {
    dpp[i][j] = x;
}
*/
import "C"

type Matrix struct {
    Row int
    Col int
    array **C.double
}

func to_cmat(mat Matrix) (cmat C.matrix_t) {
    cmat.row = C.int(mat.Row)
    cmat.col = C.int(mat.Col)
    cmat.array = mat.array
    return cmat
}

func (mat Matrix)Get(i, j int) float64 {
    return float64(C.get(mat.array, C.int(i), C.int(j)))
}

func (mat Matrix)Set(i, j int, x float64) {
    C.set(mat.array, C.int(i), C.int(j), C.double(x))
}

func Init(row, col int) Matrix {
    cmat := C.matcInit(C.int(row), C.int(col))
    mat := Matrix{Row:row, Col:col, array:cmat.array}
    return mat
}

func Deinit(mat Matrix) int {
    return int(C.matcDeinit(to_cmat(mat)))
}

func Copy(src, dest Matrix) int {
    return int(C.matcCopy(to_cmat(src), to_cmat(dest)))
}

func Add(a, b, dest Matrix) int {
    return int(C.matcAdd(to_cmat(a), to_cmat(b), to_cmat(dest)))
}

func Dot(a, b, dest Matrix) int {
    return int(C.matcDot(to_cmat(a), to_cmat(b), to_cmat(dest)))
}

func Eye(n int, mat Matrix) {
    C.matcEye(C.int(n), to_cmat(mat))
}

func Pdot(i int, c float64, mat Matrix) {
    C.matcPdot(C.int(i), C.double(c), to_cmat(mat))
}

func Qdot(i, j int, mat Matrix) {
    C.matcQdot(C.int(i), C.int(j), to_cmat(mat))
}

func Rdot(i, j int, c float64, mat Matrix) {
    C.matcRdot(C.int(i), C.int(j), C.double(c), to_cmat(mat))
}

func Rref(mat Matrix) int {
    return int(C.matcRref(to_cmat(mat)))
}

func Coladd(A, B, dest Matrix) int {
    return int(C.matcColadd(to_cmat(A), to_cmat(B), to_cmat(dest)))
}

func Inv(A, dest Matrix) int {
    return int(C.matcInv(to_cmat(A), to_cmat(dest)))
}

func Det3x3(m33 Matrix, det *float64) int {
    return int(C.matcDet3x3(to_cmat(m33), (*C.double)(det)))
}

func Det(mat Matrix, det *float64) int {
    return int(C.matcDet(to_cmat(mat), (*C.double)(det)))
}
