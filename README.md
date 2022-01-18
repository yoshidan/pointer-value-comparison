# GoLang copy/allocation cost comparison

```
%  go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               16677990                71.56 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 10463700               110.9 ns/op            80 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                 92354             12338 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   75403             16108 ns/op           12672 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                          94179             12559 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            70450             16507 ns/op           12672 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              6585876               182.2 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               25266490                46.73 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        16239200                71.31 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         70276             15208 ns/op           21760 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      93792             13149 ns/op           13184 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   72093             16596 ns/op           13824 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            110473 ns/op           25448 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            111598 ns/op           25448 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 14106             84275 ns/op           25358 B/op        145 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   10000            103265 ns/op           38484 B/op        290 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          4538            239224 ns/op            9466 B/op        428 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            4887            238791 ns/op            9459 B/op        428 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                20384             61112 ns/op           42178 B/op        158 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  17089             69984 ns/op           44483 B/op        302 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              10000            109268 ns/op           11770 B/op       1156 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                 9813            122801 ns/op           14073 B/op       1300 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison    29.760s

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
    
    // Use ptrs instead of values to avoid copy element value.
    UseSlice(ptrs) 
    UseSlice(ptrs) 
    UseSlice(ptrs) 
    UseSlice(ptrs) 
```