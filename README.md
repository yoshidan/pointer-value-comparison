# GoLang copy/allocation cost comparison

```
 % go test -bench=. ./... -benchmem 
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/fixed
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               624910406                1.901 ns/op           0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 58821982                18.89 ns/op           32 B/op          1 allocs/op
Benchmark_GetSliceValue-16                               2221482               546.2 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                  355455              3427 ns/op            5760 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                        1901372               606.9 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                           331419              3218 ns/op            5760 B/op        145 allocs/op
Benchmark_CopySliceValue-16                             16506649                72.41 ns/op            0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               16012574                77.27 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        24596698                49.85 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                        937972              1264 ns/op            9728 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                     857865              1361 ns/op            7168 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                  317434              3509 ns/op            6912 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       976762335                1.232 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         1000000000               1.210 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           76052             16217 ns/op            8225 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             67797             17268 ns/op            8227 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                141600              8090 ns/op            8235 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                  135304              8630 ns/op            8234 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                33186             35271 ns/op           37464 B/op         13 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  26598             44627 ns/op           39768 B/op        157 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16         10000            123829 ns/op             506 B/op        140 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            9697            120271 ns/op             505 B/op        140 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               71296             16717 ns/op            8206 B/op          1 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 72006             16938 ns/op            8202 B/op          1 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              18574             63535 ns/op            4856 B/op        580 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                16203             75184 ns/op            7160 B/op        724 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/fixed      37.054s
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/variable
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               16045844                73.37 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 10543267               101.8 ns/op            64 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                103053             11154 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   81225             15469 ns/op           10368 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                          99127             12804 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            73200             16085 ns/op           10368 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              7448168               156.7 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               24487377                48.71 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        15913574                72.56 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         88802             13554 ns/op           16384 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      98178             12387 ns/op           10496 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   74696             15952 ns/op           11520 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       982829187                1.213 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         997594731                1.209 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            102022 ns/op           20557 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16              9877            104530 ns/op           20597 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 14637             83740 ns/op           20578 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   13580             87956 ns/op           20573 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                21590             55137 ns/op           42113 B/op        158 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  18804             63133 ns/op           44419 B/op        302 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          5499            210676 ns/op            7166 B/op        284 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            5820            210101 ns/op            7418 B/op        284 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               15813             75334 ns/op           27425 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 15631             76230 ns/op           27424 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              12554             97425 ns/op            9465 B/op        868 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                10000            104331 ns/op           11769 B/op       1012 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/variable   38.095s

```
## Conclusion
* It is better to generate a struct by using a value instead of a pointer.
* If the element of the slice is a value and you need to loop many times, it is better to make a slice with the element as a pointer.
* `for i := range value { attr := (&value[i]).Attr }` is better than `for i, v := range value { attr := v.Attr }` when looping through the value slice.

```go
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