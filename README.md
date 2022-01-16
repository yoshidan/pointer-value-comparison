# GoLang copy/allocation cost comparison

```
% go test -bench . -benchmem
1goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz

Benchmark_GetUnitValue-16               111116719954           77.39 ns/op   0 B/op        0 allocs/op
1Benchmark_GetUnitPtr-16                1111 9546427           117.1 ns/op   80 B/op       1 allocs/op

144Benchmark_GetSliceValue-16           144144144     101185   11720 ns/op   10880 B/op    1 allocs/op
144Benchmark_GetSlicePtr-16             144144144     75032    15888 ns/op   12672 B/op    145 allocs/op

1443Benchmark_GetSliceValueAndCopy-16   144314431443  101310   11895 ns/op   10880 B/op    1 allocs/op
1443Benchmark_GetSlicePtrAndCopy-16     144314431443  73938    15536 ns/op   12672 B/op    145 allocs/op

3Benchmark_CopySliceValue-16            3333 6362424           184.1 ns/op   0 B/op        0 allocs/op
3Benchmark_CopySlicePtr-16              333324827012           47.50 ns/op   0 B/op        0 allocs/op
3Benchmark_AvoidSliceValueCopy-16       333324841756           47.00 ns/op   0 B/op        0 allocs/op

PASS
ok      github.com/yoshidan/pointer-value-comparison        12.077s
```

## Conclusion
* It is better to generate a struct by using a value instead of a pointer.
* `for i := range value { attr := (&value[i]).Attr }` is better than `for i, v := range value { attr := v.Attr }` when looping through the value slice.
