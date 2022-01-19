package variable

import (
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
	if len(value) != dataLen {
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
	if len(value) != dataLen {
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

func Benchmark_JustReturnSliceValue(b *testing.B) {
	target := NewValueSampleRepository()
	b.ResetTimer()
	value, _ := target.FindBySampleId(1)

	for i := 0; i < b.N; i++ {
		value = JustReturnValue(value)
	}
	if len(value) != dataLen {
		b.Error("invalid data")
	}
}

func Benchmark_JustReturnSlicePtr(b *testing.B) {
	target := NewPtrSampleRepository()
	b.ResetTimer()
	value, _ := target.FindBySampleId(1)
	for i := 0; i < b.N; i++ {
		value = JustReturnPtr(value)
	}
	if len(value) != dataLen {
		b.Error("invalid data")
	}
}
