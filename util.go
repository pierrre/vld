package vld

import (
	"reflect"
	"runtime"
	"strings"
)

// KeyValue holds a key and a value.
type KeyValue[K, V any] struct {
	Key   K
	Value V
}

// GetKey returns the key.
func (kv KeyValue[K, V]) GetKey() K {
	return kv.Key
}

// GetValue returns the value.
func (kv KeyValue[K, V]) GetValue() V {
	return kv.Value
}

func getFuncName(f any) string {
	v := reflect.ValueOf(f)
	var s string
	if v.Kind() == reflect.Func {
		s = runtime.FuncForPC(uintptr(v.UnsafePointer())).Name()
	}
	return s
}

func getMultiValidatorString[T any](name string, vrs ...Validator[T]) string {
	sb := new(strings.Builder)
	sb.WriteString(name)
	sb.WriteString("(")
	if len(vrs) > 0 {
		sb.WriteString("\n")
		for _, vr := range vrs {
			writeStringIndent(sb, vr.String())
			sb.WriteString(",\n")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

func writeStringIndent(sb *strings.Builder, s string) {
	for line := range strings.Lines(s) {
		sb.WriteString("\t")
		sb.WriteString(line)
	}
}
