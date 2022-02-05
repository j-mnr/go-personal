package main

import "fmt"

func main() {
	SliceBasic()
	SliceAppend()
	SliceCopy()
	SliceSlices()
	SliceMatrix()
}

func SliceBasic() {
	// Don't add a number between the
	// `[]` brackets and we `make` slices if we
	// want to have a capacity and length
	slice := make([]string, 3)
	fmt.Println("empty:", slice)

	slice[0] = "|set zeroeth value|"
	slice[1] = "|set first value|"
	slice[2] = "|set second value|"
	fmt.Println("full:", slice)
	fmt.Println("pick a value:", slice[2])
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))
}

func SliceAppend() {
	// Why wouldn't I do this always?
	var slice []string
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))

	slice = append(slice, "append a single value")
	slice = append(slice, "append", "multiple", "values")
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))
	fmt.Println("slice:", slice)
}

func SliceCopy() {
	// src is short for source
	srcSlice := make([]int, 10)
	fmt.Println("empty srcSlice:", srcSlice)
	for i := 0; i < 10; i++ {
		srcSlice[i] = i
	}
	fmt.Println("full srcSlice:", srcSlice)

	// dst is short for destination
	dstSlice := make([]int, len(srcSlice))
	fmt.Println("empty dstSlice:", dstSlice)
	copy(dstSlice, srcSlice)
	fmt.Println("full dstSlice:", dstSlice)
}

func SliceSlices() {
	var slice = []string{"zero", "one", "two", "three", "four", "five"}
	fmt.Printf("sliceUpToThirdIndex: %v\nlength: %d capacity: %d\n",
		slice,
		len(slice),
		cap(slice))

	sliceUpToThirdIndex := slice[:3]
	fmt.Printf("sliceUpToThirdIndex: %v\nlength: %d capacity: %d\n",
		sliceUpToThirdIndex,
		len(sliceUpToThirdIndex),
		cap(sliceUpToThirdIndex))

	sliceStartAtIndexTwo := slice[2:]
	fmt.Printf("sliceStartAtIndexTwo: %v\nlength: %d capacity: %d\n",
		sliceStartAtIndexTwo,
		len(sliceStartAtIndexTwo),
		cap(sliceStartAtIndexTwo))

	sliceFromOneUpToFour := slice[1:4]
	fmt.Printf("sliceFromOneUpToFour: %v\nlength: %d capacity: %d\n",
		sliceFromOneUpToFour,
		len(sliceFromOneUpToFour),
		cap(sliceFromOneUpToFour))
}

func SliceMatrix() {
	// We will allocate three slices in a slice
	matrix := make([][]int, 3)
	fmt.Println("matrix empty:", matrix)
	for i := 0; i < 3; i++ {
		innerSliceLen := i + 1
		matrix[i] = make([]int, innerSliceLen)
		for j := 0; j < innerSliceLen; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix full:", matrix)

	// and we can treat each slice like we would any other slice.
	matrix[0] = append(matrix[0], 21)
	fmt.Println("matrix append first slice with value:", matrix)
}
