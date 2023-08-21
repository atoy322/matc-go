package main

import (
    "github.com/atoy322/matc-go"
)


type IrisPlant struct {
    SepalLength float64
    SepalWidth float64
    PetalLength float64
    PetalWidth float64
    Name string
}

func argmax_1xn(mat matc.Matrix) int {
    max := 0.0
    idx := -1
    for j:=0; j<mat.Col; j++ {
        tmp := mat.Get(0, j)
        if max < tmp {
            max = tmp
            idx = j
        }
    }
    return idx
}

func (iris *IrisPlant) PredictName() string {
    X := matc.Init(1, 4)
    A := matc.Init(4, 10) // Layer1の重み
    B := matc.Init(1, 10) // Layer1のバイアス
    C := matc.Init(10, 3) // Layer2の重み
    D := matc.Init(1, 3) // Layer2のバイアス
    TMP1 := matc.Init(1, 10)
    TMP2 := matc.Init(1, 3)

    defer matc.Deinit(X)
    defer matc.Deinit(A)
    defer matc.Deinit(B)
    defer matc.Deinit(C)
    defer matc.Deinit(D)
    defer matc.Deinit(TMP1)
    defer matc.Deinit(TMP2)

    X.Set(0, 0, iris.SepalLength)
    X.Set(0, 1, iris.SepalWidth)
    X.Set(0, 2, iris.PetalLength)
    X.Set(0, 3, iris.PetalWidth)
    matc.Load("trained_params.mat", A, 'A')
    matc.Load("trained_params.mat", B, 'B')
    matc.Load("trained_params.mat", C, 'C')
    matc.Load("trained_params.mat", D, 'D')

    matc.Displayf("%7.3f", X)

    //np.dot(np.dot(np.array([5.8, 2.8, 5.1, 2.4]), W1) + b1, W2) + b2
    matc.Dot(X, A, TMP1)
    matc.Add(TMP1, B, TMP1)
    matc.Dot(TMP1, C, TMP2)
    matc.Add(TMP2, D, TMP2)

    matc.Displayf("%7.3f", TMP2)

    idx := argmax_1xn(TMP2)
    if idx == 0 {
        iris.Name = "setosa"
        return "setosa"
    } else if idx == 1 {
        iris.Name = "versicolor"
        return "versicolor"
    } else if idx == 2 {
        iris.Name = "virginica"
        return "virginica"
    } else {
        iris.Name = "unknown"
        return "unknown"
    }
}
