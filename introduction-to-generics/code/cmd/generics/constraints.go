package generics

type Hello interface {
	Say() string
}

type SomeNumbers interface {
	int | int64 | float32 | uint
}

type SomeUnderlayingNumbers interface {
	~int | ~int64 | ~float32 | ~uint
}

type MyInt int
