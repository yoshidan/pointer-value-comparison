package main

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
	print(value.SampleID)
}

func Benchmark_GetUnitPtr(b *testing.B) {
	target := NewPtrSampleRepository()
	value := &Sample{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, _ = target.FindByPK(1, 1)
	}
	print(value.SampleID)
}

func Benchmark_GetSliceValue(b *testing.B) {
	target := NewValueSampleRepository()
	var value []Sample
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
	}
	print(len(value))
}

func Benchmark_GetSlicePtr(b *testing.B) {
	target := NewPtrSampleRepository()
	var value []*Sample
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
	}
	print(len(value))
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
	print(len(value), h)
}

func Benchmark_GetSlicePtrAndCopy(b *testing.B) {
	b.ResetTimer()
	target := NewPtrSampleRepository()
	var value []*Sample
	var h int64
	for i := 0; i < b.N; i++ {
		value, _ = target.FindBySampleId(1)
		for _, v := range value {
			h = v.Quantity
		}
	}
	print(len(value), h)
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
	print(h)
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
	print(h)
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
	print(h)
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
	print(h)
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
	print(h)
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
	print(h)
}
