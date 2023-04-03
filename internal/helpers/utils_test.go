package helpers

import (
	"reflect"
	"testing"
)

func TestGetByteArray(t *testing.T) {
	type args struct {
		any     interface{}
		keyType string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Hexadecimal",
			args: args{
				any:     "68656c6c6f",
				keyType: "hexadecimal",
			},
			want: []byte("hello"),
		},
		{
			name: "Integer",
			args: args{
				any:     "123",
				keyType: "integer",
			},
			want: intToByteArray(123),
		},
		{
			name: "Default",
			args: args{
				any:     "test",
				keyType: "",
			},
			want: []byte("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetByteArray(tt.args.any, tt.args.keyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byteArrayToType(t *testing.T) {
	type args struct {
		b     []byte
		bType DataType
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "String",
			args: args{
				b:     []byte("hello"),
				bType: stringT,
			},
			want: "hello",
		},
		{
			name: "Hex",
			args: args{
				b:     []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f},
				bType: hexT,
			},
			want: "68656C6C6F",
		},
		{
			name: "Boolean",
			args: args{
				b:     []byte("true"),
				bType: booleanT,
			},
			want: true,
		},
		{
			name: "ByteArray",
			args: args{
				b:     []byte{0x01, 0x02, 0x03},
				bType: byteArrayT,
			},
			want: []byte{0x01, 0x02, 0x03},
		},
		{
			name: "Integer",
			args: args{
				b:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
				bType: integerT,
			},
			want: uint64(123),
		},
		{
			name: "int32LittleEndian",
			args: args{
				b:     []byte{0x7b, 0x00, 0x00, 0x00},
				bType: int32LittleEndian,
			},
			want: int32(123),
		},
		{
			name: "int32BigEndian",
			args: args{
				b:     []byte{0x00, 0x00, 0x00, 0x7b},
				bType: int32BigEndian,
			},
			want: int32(123),
		},
		{
			name: "uint32LittleEndian",
			args: args{
				b:     []byte{0x7b, 0x00, 0x00, 0x00},
				bType: uint32LittleEndian,
			},
			want: uint32(123),
		},
		{
			name: "uint32BigEndian",
			args: args{
				b:     []byte{0x00, 0x00, 0x00, 0x7b},
				bType: uint32BigEndian,
			},
			want: uint32(123),
		},
		{
			name: "int64LittleEndian",
			args: args{
				b:     []byte{0x7b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				bType: int64LittleEndian,
			},
			want: int64(123),
		},
		{
			name: "int64BigEndian",
			args: args{
				b:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
				bType: int64BigEndian,
			},
			want: int64(123),
		},
		{
			name: "uint64LittleEndian",
			args: args{
				b:     []byte{0x7b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				bType: uint64LittleEndian,
			},
			want: uint64(123),
		},
		{
			name: "uint64BigEndian",
			args: args{
				b:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
				bType: uint64BigEndian,
			},
			want: uint64(123),
		},
		{
			name: "float32LittleEndian",
			args: args{
				b:     []byte{0x00, 0x20, 0xf1, 0x47},
				bType: float32LittleEndian,
			},
			want: float32(123456.0),
		},
		{
			name: "float32BigEndian",
			args: args{
				b:     []byte{0x47, 0xf1, 0x20, 0x00},
				bType: float32BigEndian,
			},
			want: float32(123456.0),
		},
		{
			name: "float64LittleEndian",
			args: args{
				b:     []byte{0x77, 0xbe, 0x9f, 0x1a, 0x2f, 0xdd, 0x5e, 0x40},
				bType: float64LittleEndian,
			},
			want: 123.456,
		},
		{
			name: "float64BigEndian",
			args: args{
				b:     []byte{0x40, 0x5e, 0xdd, 0x2f, 0x1a, 0x9f, 0xbe, 0x77},
				bType: float64BigEndian,
			},
			want: 123.456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteArrayToType(tt.args.b, tt.args.bType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("byteArrayToType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToByteArray(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Positive Integer",
			args: args{
				num: 123456,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xe2, 0x40},
		},
		{
			name: "Negative Integer",
			args: args{
				num: -123456,
			},
			want: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x1d, 0xc0},
		},
		{
			name: "Zero",
			args: args{
				num: 0,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			name: "Small Positive Integer",
			args: args{
				num: 42,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a},
		},
		{
			name: "Small Negative Integer",
			args: args{
				num: -42,
			},
			want: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xd6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToByteArray(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyByteArray(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Empty ByteArray",
			args: args{
				b: []byte{},
			},
			want: []byte{},
		},
		{
			name: "Single Byte",
			args: args{
				b: []byte{0x42},
			},
			want: []byte{0x42},
		},
		{
			name: "Multiple Bytes",
			args: args{
				b: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			},
			want: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
		},
		{
			name: "Text",
			args: args{
				b: []byte("Hello, World!"),
			},
			want: []byte("Hello, World!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyByteArray(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
