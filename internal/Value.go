package internal

type Val interface {
    Value(name string) float64
}
