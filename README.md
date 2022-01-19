# GoLang copy/allocation cost comparison

```
% go test -bench=. ./... -benchmem 
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/fixed
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               639632342                1.878 ns/op           0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 59168506                19.55 ns/op           32 B/op          1 allocs/op
Benchmark_GetSliceValue-16                               2236336               539.4 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                  355158              3141 ns/op            5760 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                        2010330               600.5 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                           353223              3130 ns/op            5760 B/op        145 allocs/op
Benchmark_CopySliceValue-16                             16268889                72.65 ns/op            0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               16507064                72.18 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        25215780                48.02 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                        957078              1238 ns/op            9728 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                     868753              1348 ns/op            7168 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                  308896              3523 ns/op            6912 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       1000000000               1.210 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         960750206                1.218 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           70633             16832 ns/op            8226 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             67092             18016 ns/op            8230 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                143126              8204 ns/op            8234 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                  135495              8612 ns/op            8238 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                33772             35551 ns/op           37389 B/op         11 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  33004             36994 ns/op           37389 B/op         11 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          9902            121223 ns/op             506 B/op        140 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16           10000            118417 ns/op             505 B/op        140 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               71272             16622 ns/op            8206 B/op          1 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 71608             16831 ns/op            8202 B/op          1 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              23684             49833 ns/op              72 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                22550             52696 ns/op              72 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/fixed      37.686s
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/variable
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               16452319                70.17 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 11084190               103.4 ns/op            64 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                104545             11576 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   76342             16467 ns/op           10368 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                          94794             12216 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            73737             16600 ns/op           10368 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              7319156               157.2 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               24534486                49.39 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        15677638                72.56 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         89682             13800 ns/op           16384 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      88424             13013 ns/op           10496 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   64146             16629 ns/op           11520 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       988658186                1.228 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         996603160                1.206 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            101550 ns/op           20562 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            104986 ns/op           20567 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 14911             81960 ns/op           20572 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   14511             83244 ns/op           20569 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                22752             52507 ns/op           42000 B/op        155 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  22455             52807 ns/op           42000 B/op        155 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          5404            211293 ns/op            7423 B/op        284 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            5624            205834 ns/op            7418 B/op        284 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               16311             74493 ns/op           27428 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 16353             75626 ns/op           27423 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              18770             63316 ns/op              73 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                17954             65750 ns/op              72 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/variable   38.235s

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