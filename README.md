# GoLang copy/allocation cost comparison

```
%  go test -bench . -benchmem  
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               15779986                73.62 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 10596181               111.4 ns/op            80 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                101122             11481 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   75417             15697 ns/op           12672 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                         101269             11987 ns/op           10880 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            74548             15479 ns/op           12672 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              6631846               182.4 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               25751415                46.14 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        16807734                70.39 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         84566             14143 ns/op           21760 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      95708             12303 ns/op           13184 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   71443             15947 ns/op           13824 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       1000000000               0.9446 ns/op          0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         1000000000               1.151 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            107515 ns/op           25436 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            110058 ns/op           25436 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 14175             85666 ns/op           25503 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   10000            103051 ns/op           38517 B/op        291 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                21392             56237 ns/op           42178 B/op        158 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  17725             66627 ns/op           44483 B/op        302 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          4724            240013 ns/op            9210 B/op        428 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            4435            234247 ns/op            9459 B/op        428 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               15267             79431 ns/op           32301 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 15098             79161 ns/op           32230 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              10000            103567 ns/op           11770 B/op       1156 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                10605            113896 ns/op           14073 B/op       1300 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison    36.767s

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