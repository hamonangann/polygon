# polygon

Go package to construct simple polygonal shapes. 

## Example

```go
p, _ := polygon.NewPolygon(
    polygon.Point{0, 0},
    polygon.Point{0, 3},
    polygon.Point{4, 0},
)
fmt.Println(p.Area()) // 6.0
```

```go
_, err := polygon.NewPolygon(
    polygon.Point{0, 0},
    polygon.Point{3, 3},
)
fmt.Println(err) // ErrInsufficientPoints
```

## Installation

```shell
go get github.com/hamonangann/polygon
```

## Contribution

Fork [this repository](https://github.com/hamonangann/polygon) and open a [pull request](https://github.com/hamonangann/polygon/pulls) to improve this project. Your contribution matters a lot!

If you have any inquiries, please raise an [issue](https://github.com/hamonangann/polygon/issues).