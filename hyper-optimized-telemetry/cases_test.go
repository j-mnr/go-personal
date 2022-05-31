package hotelemetry_test

var toBufferTT = map[string]struct {
	value int64
	want  [9]byte
}{
	"upper bound of int64": {
		0x7fffffffffffffff, // or (1 << 63) - 1 or 9223372036854775807
		[9]byte{0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
	},
	"lower bound of int64": {
		0x100000000, // or (1 << 32) or 4294967296
		[9]byte{0xf8, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0},
	},
	"upper bound of uint32": {
		0xffffffff, // or (1 << 32) - 1 or 4294967295
		[9]byte{0x4, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0},
	},
	"lower bound of uint32": {
		0x80000000, // or (1 << 31) or 2147483648
		[9]byte{0x4, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0},
	},
	"upper bound of int32": {
		0x7fffffff, // or (1 << 31) - 1 or 2147483647
		[9]byte{0xfc, 0xff, 0xff, 0xff, 0x7f, 0x0, 0x0, 0x0, 0x0},
	},
	"lower bound of int32": {
		0x10000, // or (1 << 16) or 65536
		[9]byte{0xfc, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"upper bound of uint16": {
		0xffff, // or (1 << 16) - 1 or 65535
		[9]byte{0x2, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"lower bound of uint16": {
		0x8000, // or (1 << 15) or 32768
		[9]byte{0x2, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"upper bound of int16": {
		0x7fff, // or (1 << 15) - 1 or 32767
		[9]byte{0x2, 0xff, 0x7f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"zero": {
		0,
		[9]byte{0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"negative upper bound of int16": {
		-1,
		[9]byte{0xfe, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"negative lower bound of int16": {
		-0x8000, // or (-1 << 15) or -32768
		[9]byte{0xfe, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"negative upper bound of int32": {
		-0x8001, // or (-1 << 15) - 1 or -32769
		[9]byte{0xfc, 0xff, 0x7f, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0},
	},
	"negative lower bound of int32": {
		-0x80000000, // or (-1 << 31) or -2147483648
		[9]byte{0xfc, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0},
	},
	"negative upper bound of int64": {
		-0x80000001, // or (-1 << 31) - 1 or -2147483649
		[9]byte{0xf8, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xff},
	},
	"negative lower bound of int64": {
		-0x8000000000000000, // or (-1 << 63) or -9223372036854775808
		[9]byte{0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
	},
}

var fromBufferTT = map[string]struct {
	buf  [9]byte
	want int64
}{
	"invalid prefix byte": {
		[9]byte{22, 0xff, 0xff, 0xff, 0x7f, 0, 0, 0, 0},
		0,
	},
	"upper bound of int64": {
		[9]byte{0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
		0x7fffffffffffffff, // or (1 << 63) - 1 or 9223372036854775807
	},
	"lower bound of int64": {
		[9]byte{0xf8, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0},
		0x100000000, // or (1 << 32) or 4294967296
	},
	"upper bound of uint32": {
		[9]byte{0x4, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0},
		0xffffffff, // or (1 << 32) - 1 or 4294967295
	},
	"lower bound of uint32": {
		[9]byte{0x4, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0},
		0x80000000, // or (1 << 31) or 2147483648
	},
	"upper bound of int32": {
		[9]byte{0xfc, 0xff, 0xff, 0xff, 0x7f, 0x0, 0x0, 0x0, 0x0},
		0x7fffffff, // or (1 << 31) - 1 or 2147483647
	},
	"lower bound of int32": {
		[9]byte{0xfc, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0},
		0x10000, // or (1 << 16) or 65536
	},
	"upper bound of uint16": {
		[9]byte{0x2, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		0xffff, // or (1 << 16) - 1 or 65535
	},
	"lower bound of uint16": {
		[9]byte{0x2, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		0x8000, // or (1 << 15) or 32768
	},
	"upper bound of int16": {
		[9]byte{0xfe, 0xff, 0x7f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		0x7fff, // or (1 << 15) - 1 or 32767
	},
	"zero": {
		[9]byte{0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		0,
	},
	"negative upper bound of int16": {
		[9]byte{0xfe, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		-1,
	},
	"negative lower bound of int16": {
		[9]byte{0xfe, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		-0x8000, // or (-1 << 15) or -32768
	},
	"negative upper bound of int32": {
		[9]byte{0xfc, 0xff, 0x7f, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0},
		-0x8001, // or (-1 << 15) - 1 or -32769
	},
	"negative lower bound of int32": {
		[9]byte{0xfc, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0},
		-0x80000000, // or (-1 << 31) or -2147483648
	},
	"negative upper bound of int64": {
		[9]byte{0xf8, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xff},
		-0x80000001, // or (-1 << 31) - 1 or -2147483649
	},
	"negative lower bound of int64": {
		[9]byte{0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
		-0x8000000000000000, // or (-1 << 63) or -9223372036854775808
	},
}
