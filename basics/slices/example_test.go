package slices_test

import "basics/slices"

func ExampleSliceBasic() {
	slices.SliceBasic()
	// Output:
	// empty: [  ]
	// full: [|set zeroeth value| |set first value| |set second value|]
	// pick a value: |set second value|
	// capacity: 3
	// length: 3
	// Can be declared inline [0 1 2 3 4]
}

func ExampleSliceAppend() {
	slices.SliceAppend()
	// Output:
	// capacity: 0
	// length: 0
	// capacity: 4
	// length: 4
	// We had to go find more space! Which takes time and effort!
	// slice: [append a single value append multiple values]
}

func ExampleSliceCopy() {
	slices.SliceCopy()
	// Output:
	// empty srcSlice: [0 0 0 0 0 0 0 0 0 0]
	// full srcSlice: [0 1 2 3 4 5 6 7 8 9]
	// empty dstSlice: [0 0 0 0 0 0 0 0 0 0]
	// full dstSlice: [0 1 2 3 4 5 6 7 8 9]
}

func ExampleSliceIndexOutOfRangePanic() {
	slices.SliceIndexOutOfRangePanic()
	// Output:
	// slice paniced!
	//  runtime error: index out of range [-1]
}

func ExampleSliceSlices() {
	slices.SliceSlices()
	// Output:
	// sliceUpToThirdIndex: [zero one two three four five]
	// length: 6 capacity: 6
	// sliceUpToThirdIndex: [zero one two]
	// length: 3 capacity: 6
	// sliceStartAtIndexTwo: [two three four five]
	// length: 4 capacity: 4
	// sliceFromOneUpToFour: [one two three]
	// length: 3 capacity: 5
	// Efficiency to the Max for substrings
}

func ExampleSliceMatrix() {
	slices.SliceMatrix()
	// Output:
	// matrix empty: [[] [] []]
	// matrix full: [[0] [1 2] [2 3 4]]
	// matrix append first slice with value: [[0 21] [1 2] [2 3 4]]
}
