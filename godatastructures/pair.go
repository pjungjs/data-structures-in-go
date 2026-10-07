package godatastructures

import (
	"golang.org/x/exp/constraints"
)

type Pair[T constraints.Ordered] struct {
	Key   T
	Value any
}
