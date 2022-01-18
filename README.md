# GoLang copy/allocation cost comparison

```
% go test -bench . -benchmem
1goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               111116737828            71.81 ns/op            0 B/op          0 allocs/op
1Benchmark_GetUnitPtr-16                                111110434010           114.6 ns/op            80 B/op          1 allocs/op
144Benchmark_GetSliceValue-16                                   144144144   99127            11749 ns/op           10880 B/op          1 allocs/op
144Benchmark_GetSlicePtr-16                                     144144144   75532            15989 ns/op           12672 B/op        145 allocs/op
1443Benchmark_GetSliceValueAndCopy-16                           144314431443   95520         12039 ns/op           10880 B/op          1 allocs/op
1443Benchmark_GetSlicePtrAndCopy-16                             144314431443   74722         15996 ns/op           12672 B/op        145 allocs/op
3Benchmark_CopySliceValue-16                            3333 6344016           185.4 ns/op             0 B/op          0 allocs/op
3Benchmark_CopySlicePtr-16                              333315454370            72.12 ns/op            0 B/op          0 allocs/op
3Benchmark_AvoidSliceValueCopy-16                       333316650676            72.14 ns/op            0 B/op          0 allocs/op
3Benchmark_UseSliceValueWithCopy-16                     333   78662          14425 ns/op           21760 B/op          2 allocs/op
3Benchmark_UseSliceValueWithoutCopy-16                  333   87982          13801 ns/op           13184 B/op          3 allocs/op
3Benchmark_UseSlicePtr-16                               333   65694          16380 ns/op           13824 B/op        146 allocs/op
16536Benchmark_SerializeSliceValueWithStandardJson-16           1653616541   10000          111768 ns/op           25440 B/op        146 allocs/op
16536Benchmark_SerializeSlicePtrWithStandardJson-16             1653616537   10000          110590 ns/op           25432 B/op        146 allocs/op
16537Benchmark_SerializeSliceValueWithGoJson-16                 165501653616538   13738      88043 ns/op           25364 B/op        145 allocs/op
16552Benchmark_SerializeSlicePtrWithGoJson-16                   1653716536   10000          106949 ns/op           38423 B/op        290 allocs/op
144Benchmark_DeserializeSliceValueWithStandardJson-16           144144    4666      244279 ns/op            9466 B/op        428 allocs/op
144Benchmark_DeserializeSlicePtrWithStandardJson-16             144144    4318      242109 ns/op            9219 B/op        428 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison    23.653s

```

## Conclusion
* It is better to generate a struct by using a value instead of a pointer.
* `for i := range value { attr := (&value[i]).Attr }` is better than `for i, v := range value { attr := v.Attr }` when looping through the value slice.
