# GoLang copy/allocation cost comparison

```
%  go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               15903512                74.98 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                  9818766               115.0 ns/op            80 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                101590             11684 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   75370             16131 ns/op           12672 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                          95856             12119 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            71475             16155 ns/op           12672 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              6256909               184.6 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               24947664                47.58 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        16808020                71.06 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         83845             14609 ns/op           21760 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      88630             13133 ns/op           13184 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   73376             16329 ns/op           13824 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       1000000000               0.9542 ns/op          0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         1000000000               1.180 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            107812 ns/op           25428 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            110697 ns/op           25436 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 13528             85189 ns/op           25358 B/op        145 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                    9859            103436 ns/op           38464 B/op        290 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                21544             56125 ns/op           42178 B/op        158 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  18205             66775 ns/op           44483 B/op        302 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          4689            242084 ns/op            9306 B/op        428 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            5221            240597 ns/op            9283 B/op        428 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               15034             79477 ns/op           31848 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 14714             82108 ns/op           32295 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              10000            106557 ns/op           11770 B/op       1156 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                 9764            117277 ns/op           14073 B/op       1300 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison    36.140s

```

## Conclusion
* It is better to generate a struct by using a value instead of a pointer.
* If the element of the slice is a value and you need to loop many times, it is better to make a slice with the element as a pointer.
* `for i := range value { attr := (&value[i]).Attr }` is better than `for i, v := range value { attr := v.Attr }` when looping through the value slice.

``` 
    // Get value slice to avoid heap allocation of slice elements.
    // In some cases, memory allocation can be avoided even if FindAll returns pointer slice by `Escape Analysis`.
    var values []Sample
    values = repository.FindAll() 
    
    // Only 1 allocation
    ptrs := make([]*Sample, len(values)
    
    // Make pointer slice without copy element value.
    for i := range values {
        ptrs[i] = &values[i]
    }
    
    // Use ptrs instead of values to avoid copy element value in loop.
    ForEachSlice(ptrs) 
    ForEachSlice(ptrs) 
    ForEachSlice(ptrs) 
    ForEachSlice(ptrs) 
```