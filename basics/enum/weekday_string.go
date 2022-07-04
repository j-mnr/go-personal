// Code generated by "stringer -type=Weekday"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Monday-1]
	_ = x[Tuesday-2]
	_ = x[Wednesday-4]
	_ = x[Thursday-8]
	_ = x[Friday-16]
	_ = x[Saturday-32]
	_ = x[Sunday-64]
}

const (
	_Weekday_name_0 = "MondayTuesday"
	_Weekday_name_1 = "Wednesday"
	_Weekday_name_2 = "Thursday"
	_Weekday_name_3 = "Friday"
	_Weekday_name_4 = "Saturday"
	_Weekday_name_5 = "Sunday"
)

var (
	_Weekday_index_0 = [...]uint8{0, 6, 13}
)

func (i Weekday) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Weekday_name_0[_Weekday_index_0[i]:_Weekday_index_0[i+1]]
	case i == 4:
		return _Weekday_name_1
	case i == 8:
		return _Weekday_name_2
	case i == 16:
		return _Weekday_name_3
	case i == 32:
		return _Weekday_name_4
	case i == 64:
		return _Weekday_name_5
	default:
		return "Weekday(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
