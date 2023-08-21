package main

import (
    "fmt"
)


func main() {
    iris := &IrisPlant{}

    fmt.Print("Sepal length (cm): ")
    fmt.Scanf("%f\n", &iris.SepalLength)
    
    fmt.Print("Sepal width (cm): ")
    fmt.Scanf("%f\n", &iris.SepalWidth)

    fmt.Print("Petal length (cm): ")
    fmt.Scanf("%f\n", &iris.PetalLength)

    fmt.Print("Petal width (cm): ")
    fmt.Scanf("%f\n", &iris.PetalWidth)

    iris.PredictName()

    fmt.Printf("It is %s\n", iris.Name)

    return
}
