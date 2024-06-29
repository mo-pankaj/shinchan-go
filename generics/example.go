package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

type Points []int32

// this will return []int32 not the type we want to have
func scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, e := range s {
		r[i] = e * c
	}
	return r
}

// to have the desired type, we have to create type S ~[]E
func updatedScale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, e := range s {
		r[i] = e * c
	}
	return r
}

func (p Points) String() string {
	builder := strings.Builder{}
	for _, v := range p {
		builder.Write([]byte(strconv.Itoa(int(v))))
		builder.Write([]byte("\n"))
	}
	return builder.String()
}

func main() {
	fmt.Println("Hello, 世界")
	arr := []int32{1, 2, 3, 4}
	fmt.Println("new value by scaling by 2 ", scale(arr, 2))

	points := Points([]int32{1, 2, 3, 4})
	//np := scale(points, 2)
	//fmt.Println(np.String()) // throwing error

	np := updatedScale(points, 2)
	fmt.Println(np.String())

}
