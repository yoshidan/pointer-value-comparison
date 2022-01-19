package fixed

import (
	"encoding/json"
	gojson "github.com/goccy/go-json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

func Benchmark_SerializeSliceValueWithStandardJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = json.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSlicePtrWithStandardJson(b *testing.B) {
	target := NewPtrSampleRepository()
	var str []byte
	value, _ := target.FindBySampleId(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = json.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSliceValueWithGoJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = gojson.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSlicePtrWithGoJson(b *testing.B) {
	target := NewPtrSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = gojson.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSliceValueWithMsgPack(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = msgpack.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSlicePtrWithMsgPack(b *testing.B) {
	target := NewPtrSampleRepository()
	var str []byte
	value, _ := target.FindBySampleId(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = msgpack.Marshal(value)
	}
	if len(str) < 100 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSliceValueWithStandardJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := json.Marshal(value)
	var res []Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSlicePtrWithStandardJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := json.Marshal(value)
	var res []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSliceValueWithGoJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := gojson.Marshal(value)
	var res []Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = gojson.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSlicePtrWithGoJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := gojson.Marshal(value)
	var res []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = gojson.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSliceValueWithMsgPack(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := msgpack.Marshal(value)
	var res []Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSlicePtrWithMsgPack(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := msgpack.Marshal(value)
	var res []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}
