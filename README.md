# GoLang copy/allocation cost comparison

```
% go test -bench=. ./... -benchmem 
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/fixed
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               564942596                2.047 ns/op           0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 55416172                22.37 ns/op           32 B/op          1 allocs/op
Benchmark_GetSliceValue-16                               1953220               586.6 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                  315903              3708 ns/op            5760 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                        1868181               648.2 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                           306661              3466 ns/op            5760 B/op        145 allocs/op
Benchmark_CopySliceValue-16                             16687314                72.36 ns/op            0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               15872484                72.71 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        24748663                48.41 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                        925766              1323 ns/op            9728 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                     862468              1309 ns/op            7168 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                  277981              3947 ns/op            6912 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       979711736                1.221 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         961103392                1.218 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           69895             16765 ns/op            8224 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             64363             17849 ns/op            8224 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                131920              8158 ns/op            8233 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                  136615              8355 ns/op            8231 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                33607             38399 ns/op           37389 B/op         11 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  28881             38579 ns/op           37389 B/op         11 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16         10000            126406 ns/op             506 B/op        140 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16           10000            117468 ns/op             505 B/op        140 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               72645             16332 ns/op            8206 B/op          1 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 68772             17749 ns/op            8202 B/op          1 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              22490             54935 ns/op              72 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                21380             56708 ns/op              72 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/fixed      37.812s
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/variable
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               16654639                72.13 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 11334466               103.2 ns/op            64 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                 99873             11641 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   73026             15732 ns/op           10368 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                         101836             11887 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            77714             15635 ns/op           10368 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              7322641               155.1 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               25421738                47.54 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        16226784                72.35 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         90504             13623 ns/op           16384 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      87691             12435 ns/op           10496 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   69937             16850 ns/op           11520 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       989181951                1.254 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         888444481                1.286 ns/op           0 B/op          0 allocs/op
Benchmark_SliceLoopPtrFirst-16                             20389             58118 ns/op          125568 B/op        245 allocs/op
Benchmark_SliceLoopValue-16                                 4935            221058 ns/op          827399 B/op        101 allocs/op
Benchmark_SliceLoopPtr-16                                  22644             53703 ns/op          124544 B/op        102 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            104840 ns/op           20562 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            105836 ns/op           20547 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 14526             85075 ns/op           20528 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   14367             82817 ns/op           20533 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                22717             51218 ns/op           42000 B/op        155 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  22568             51628 ns/op           42000 B/op        155 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          6084            200296 ns/op            7422 B/op        284 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            5878            196392 ns/op            7178 B/op        284 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               16336             72650 ns/op           27426 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 16140             76488 ns/op           27423 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              18501             65014 ns/op              72 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                17980             67276 ns/op              73 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/variable   42.989s
yoshida@KNkanrishanoMacBook-Pro pointer-value-comparison % go test -bench=. ./... -benchmem 
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/fixed
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               555121180                2.066 ns/op           0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 46607631                22.47 ns/op           32 B/op          1 allocs/op
Benchmark_GetSliceValue-16                               1981252               581.3 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                  298941              4003 ns/op            5760 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                        1814124               656.0 ns/op          4864 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                           303441              3799 ns/op            5760 B/op        145 allocs/op
Benchmark_CopySliceValue-16                             14161746                83.57 ns/op            0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               14185014                75.97 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        23046901                49.57 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                        783109              1421 ns/op            9728 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                     893193              1331 ns/op            7168 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                  298849              4157 ns/op            6912 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       997017520                1.277 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         918466622                1.303 ns/op           0 B/op          0 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           65490             17723 ns/op            8226 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             60888             20156 ns/op            8227 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                130324              8455 ns/op            8231 B/op          2 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                  132949              8850 ns/op            8231 B/op          2 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                28549             39034 ns/op           37389 B/op         11 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  29610             41641 ns/op           37389 B/op         11 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          9265            134571 ns/op             506 B/op        140 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            9052            133898 ns/op             505 B/op        140 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               70087             18196 ns/op            8206 B/op          1 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 67141             17833 ns/op            8202 B/op          1 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              21286             56992 ns/op              72 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                21055             57115 ns/op              72 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/fixed      36.130s
goos: darwin
goarch: amd64
pkg: github.com/yoshidan/pointer-value-comparison/variable
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_GetUnitValue-16                               16672654                70.97 ns/op            0 B/op          0 allocs/op
Benchmark_GetUnitPtr-16                                 11558304               107.5 ns/op            64 B/op          1 allocs/op
Benchmark_GetSliceValue-16                                 99309             12654 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtr-16                                   71176             16238 ns/op           10368 B/op        145 allocs/op
Benchmark_GetSliceValueAndCopy-16                          99676             12916 ns/op            8192 B/op          1 allocs/op
Benchmark_GetSlicePtrAndCopy-16                            68630             16927 ns/op           10368 B/op        145 allocs/op
Benchmark_CopySliceValue-16                              7199067               169.6 ns/op             0 B/op          0 allocs/op
Benchmark_CopySlicePtr-16                               23414990                53.08 ns/op            0 B/op          0 allocs/op
Benchmark_AvoidSliceValueCopy-16                        15077337                86.20 ns/op            0 B/op          0 allocs/op
Benchmark_UseSliceValueWithCopy-16                         76418             14900 ns/op           16384 B/op          2 allocs/op
Benchmark_UseSliceValueWithoutCopy-16                      96098             13649 ns/op           10496 B/op          3 allocs/op
Benchmark_UseSlicePtr-16                                   71172             16954 ns/op           11520 B/op        146 allocs/op
Benchmark_JustReturnSliceValue-16                       954010326                1.243 ns/op           0 B/op          0 allocs/op
Benchmark_JustReturnSlicePtr-16                         942600884                1.272 ns/op           0 B/op          0 allocs/op
Benchmark_SliceLoopPtr-16                                  20131             57531 ns/op          125568 B/op        245 allocs/op
Benchmark_SliceLoopValue-16                                 5499            221338 ns/op          827399 B/op        101 allocs/op
Benchmark_SliceLoopValueToPtr-16                           21424             52993 ns/op          124544 B/op        102 allocs/op
Benchmark_SerializeSliceValueWithStandardJson-16           10000            103152 ns/op           20547 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithStandardJson-16             10000            109230 ns/op           20537 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithGoJson-16                 13663             87612 ns/op           20548 B/op        146 allocs/op
Benchmark_SerializeSlicePtrWithGoJson-16                   13776             87330 ns/op           20548 B/op        146 allocs/op
Benchmark_SerializeSliceValueWithMsgPack-16                21171             56500 ns/op           42000 B/op        155 allocs/op
Benchmark_SerializeSlicePtrWithMsgPack-16                  21673             54046 ns/op           42000 B/op        155 allocs/op
Benchmark_DeserializeSliceValueWithStandardJson-16          5142            212526 ns/op            7423 B/op        284 allocs/op
Benchmark_DeserializeSlicePtrWithStandardJson-16            6057            210673 ns/op            7418 B/op        284 allocs/op
Benchmark_DeserializeSliceValueWithGoJson-16               16010             74477 ns/op           27424 B/op        289 allocs/op
Benchmark_DeserializeSlicePtrWithGoJson-16                 16725             72096 ns/op           27424 B/op        289 allocs/op
Benchmark_DeserializeSliceValueWithMsgPack-16              18514             67362 ns/op              73 B/op          2 allocs/op
Benchmark_DeserializeSlicePtrWithMsgPack-16                17002             67968 ns/op              73 B/op          2 allocs/op
PASS
ok      github.com/yoshidan/pointer-value-comparison/variable   43.849s

```

## Conclusion

```
Benchmark_SliceLoopPtr-16                                  20131             57531 ns/op          125568 B/op        245 allocs/op
Benchmark_SliceLoopValue-16                                 5499            221338 ns/op          827399 B/op        101 allocs/op
Benchmark_SliceLoopValueToPtr-16                           21424             52993 ns/op          124544 B/op        102 allocs/op
```

### Benchmark_SliceLoopPtr-16 
* There are concerns about GC because it is fast but has a large heap.
```go
func Benchmark_SliceLoopPtr(b *testing.B) {
	repository := NewPtrSampleRepository()
	var x []*Sample
	for i := 0; i < b.N; i++ {
		x, _ = repository.FindBySampleId(1)

		for j := 0; j < 100; j++ {
			x = FilterPtr(x)
		}
	}
}
```

### Benchmark_SliceLoopValue-16
* Slow due to large numbers of copies.
```go
func Benchmark_SliceLoopValue(b *testing.B) {
	repository := NewValueSampleRepository()
	var x []Sample
	for i := 0; i < b.N; i++ {
		x, _ = repository.FindBySampleId(1)

		for j := 0; j < 100; j++ {
			x = FilterValue(x)
		}
	}
}
```

### Benchmark_SliceLoopValueToPtr-16
* Faster with fewer copies and less memory allocation.
```go
func Benchmark_SliceLoopValueToPtr(b *testing.B) {
	repository := NewValueSampleRepository()
	var values []Sample
	for i := 0; i < b.N; i++ {
		values, _ = repository.FindBySampleId(1)

		// change value -> ptr
		ptrs := make([]*Sample, len(values))
		for j := range values {
			ptrs[j] = &values[j]
		}

		for j := 0; j < 100; j++ {
			ptrs = FilterPtr(ptrs)
		}
	}
}
```
