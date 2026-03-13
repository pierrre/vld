package vld

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
