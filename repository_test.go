package main

import (
	"encoding/json"
	gojson "github.com/goccy/go-json"
	"github.com/vmihailenco/msgpack"
	"testing"
)

func Benchmark_GetUnitValue(b *testing.B) {
	target := NewValueSampleRepository()
	value := Sample{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindByPK(1, 1)
	}
	if value.Quantity != 3 {
		print(value.SampleID)
	}
}

func Benchmark_GetUnitPtr(b *testing.B) {
	target := NewPtrSampleRepository()
	value := &Sample{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindByPK(1, 1)
	}
	if value.Quantity != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_GetSliceValue(b *testing.B) {
	target := NewValueSampleRepository()
	var value []Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
	}
	if len(value) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_GetSlicePtr(b *testing.B) {
	target := NewPtrSampleRepository()
	var value []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
	}
	if len(value) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_GetSliceValueAndCopy(b *testing.B) {
	target := NewValueSampleRepository()
	var value []Sample
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
		for _, v := range value {
			h = v.Quantity
		}
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_GetSlicePtrAndCopy(b *testing.B) {
	target := NewPtrSampleRepository()
	var value []*Sample
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
		for _, v := range value {
			h = v.Quantity
		}
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_CopySliceValue(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range value {
			h = v.Quantity
		}
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_CopySlicePtr(b *testing.B) {
	target := NewPtrSampleRepository()
	value, _ := target.FindBySampleId(1)
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range value {
			h = v.Quantity
		}
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_AvoidSliceValueCopy(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := range value {
			h = (&value[j]).Quantity
		}
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_UseSliceValueWithCopy(b *testing.B) {
	target := NewValueSampleRepository()
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ := target.FindBySampleId(1)
		v := FilterValue(value)
		h = (&v[0]).Quantity
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_UseSliceValueWithoutCopy(b *testing.B) {
	target := NewValueSampleRepository()
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ := target.FindBySampleId(1)
		ptr := make([]*Sample, len(value))
		for j := range value {
			ptr[j] = &value[j]
		}
		v := FilterPtr(ptr)
		h = v[0].Quantity
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_UseSlicePtr(b *testing.B) {
	target := NewPtrSampleRepository()
	var h int64
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ptr, _ := target.FindBySampleId(1)
		v := FilterPtr(ptr)
		h = v[0].Quantity
	}
	if h != 3 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSliceValueWithStandardJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = json.Marshal(value)
	}
	if len(str) == 0 {
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
	if len(str) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSliceValueWithGoJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, _ = gojson.MarshalNoEscape(value)
	}
	if len(str) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_SerializeSlicePtrWithGoJson(b *testing.B) {
	target := NewPtrSampleRepository()
	var str []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, _ := target.FindBySampleId(1)
		str, _ = gojson.MarshalNoEscape(value)
	}
	if len(str) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSliceValueWithStandardJson(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := json.Marshal(&value)
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
	str, _ := json.Marshal(&value)
	var res []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(str, &res)
	}
	if len(res) == 0 {
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
	if len(str) == 0 {
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
	if len(str) == 0 {
		b.Error("invalid data")
	}
}

func Benchmark_DeserializeSliceValueWithMsgPack(b *testing.B) {
	target := NewValueSampleRepository()
	value, _ := target.FindBySampleId(1)
	str, _ := msgpack.Marshal(&value)
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
	str, _ := msgpack.Marshal(&value)
	var res []*Sample
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(str, &res)
	}
	if len(res) == 0 {
		b.Error("invalid data")
	}
}
