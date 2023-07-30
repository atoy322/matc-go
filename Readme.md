# Go binding for Matc
[Matc](https://github.com/atoy322/matc) is a small library treating a matrix.

## Building for Windows
### Clone and Make Matc
```bash
git clone https://github.com/atoy322/matc
cd matc
make static # "libmatc" directory will be created
```

### Add a path to env variables
Before doing `go build` the folowing steps should be done.
```bash
set C_INCLUDE_PATH=C:\path\to\libmatc\include
set LIBRARY_PATH=C:\path\to\libmatc\lib
```

### go get
```bash
go get github.com/atoy322/matc-go
```

## Example
source:
```golang
package main

import "github.com/atoy322/matc-go"

func main() {
    // allocate
    mat := matc.Init(5, 5)
    matc.Eye(5, mat)

    // show
    matc.Displayf("%2.0f", mat)

    // free
    matc.Deinit(mat)
}
```
output:
```
┌                    ┐
│  1   0   0   0   0 │
│  0   1   0   0   0 │
│  0   0   1   0   0 │
│  0   0   0   1   0 │
│  0   0   0   0   1 │
└                    ┘
```