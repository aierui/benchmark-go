# Benchmark-Go


```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
```


## go1.11.2 vs go1.16.5
```
➜  benchmark-go git:(master) ✗ sh benchcmp.sh go1.11.2 go1.16.5
allocate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
BenchmarkMakeSmallSlice-8       40.7          37.4          -8.11%
BenchmarkMakeSmallMap-8         410           294           -28.32%
BenchmarkMakeSmallChannel-8     189           149           -21.06%
BenchmarkMakeBigSlice-8         117111        168318        +43.73%
BenchmarkMakeBigMap-8           2246666       2775862       +23.55%
BenchmarkMakeBigChannel-8       761448        918614        +20.64%
BenchmarkNewSmallObject-8       0.32          0.32          -0.72%
BenchmarkNewBigObject-8         655206        688248        +5.04%

channel
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                          old ns/op     new ns/op     delta
BenchmarkSendBlockChannel-8        204           209           +2.50%
BenchmarkGetBlockChannel-8         209           244           +16.84%
BenchmarkSendNonBlockChannel-8     55.0          52.6          -4.44%
BenchmarkGetNonBlockChannel-8      54.8          51.5          -5.97%

copy
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
Benchmark_GOBDeepCopy-8         161725        172745        +6.81%
Benchmark_ReflectDeepCopy-8     9622          10261         +6.64%

gc
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                    old ns/op     new ns/op     delta
BenchmarkSmallMemoryGC-8     10477         13134         +25.36%
BenchmarkBigMemoryGC-8       200133569     42074140      -78.98%

goroutine
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                         old ns/op     new ns/op     delta
BenchmarkGoroutineMergeSort-8     1631871       1827338       +11.98%

iterate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                        old ns/op     new ns/op     delta
BenchmarkRangeMap-8              18917         19482         +2.99%
BenchmarkRangeSlices-8           348           407           +17.04%
BenchmarkRangeArrays-8           337           341           +1.28%
BenchmarkRangeAndUseMap-8        15243         12933         -15.15%
BenchmarkRangeAndUseSlices-8     590           328           -44.36%
BenchmarkRangeAndUseArrays-8     674           738           +9.44%
BenchmarkForAndUseSlices-8       353           441           +24.90%
BenchmarkForAndUseArrays-8       357           398           +11.43%

sync
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark             old ns/op     new ns/op     delta
BenchmarkMutex-8      360           324           -10.14%
BenchmarkAtomic-8     291           325           +11.65%
```


## go1.11.2 vs go1.17.2
```
➜  benchmark-go git:(master) ✗ sh benchcmp.sh go1.11.2 go1.17.2
allocate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
BenchmarkMakeSmallSlice-8       40.7          51.4          +26.19%
BenchmarkMakeSmallMap-8         410           352           -14.07%
BenchmarkMakeSmallChannel-8     189           149           -21.11%
BenchmarkMakeBigSlice-8         117111        96606         -17.51%
BenchmarkMakeBigMap-8           2246666       2185768       -2.71%
BenchmarkMakeBigChannel-8       761448        896100        +17.68%
BenchmarkNewSmallObject-8       0.32          0.32          -0.69%
BenchmarkNewBigObject-8         655206        748333        +14.21%

channel
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                          old ns/op     new ns/op     delta
BenchmarkSendBlockChannel-8        204           244           +19.41%
BenchmarkGetBlockChannel-8         209           198           -5.41%
BenchmarkSendNonBlockChannel-8     55.0          50.0          -9.11%
BenchmarkGetNonBlockChannel-8      54.8          48.2          -12.06%

copy
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
Benchmark_GOBDeepCopy-8         161725        98527         -39.08%
Benchmark_ReflectDeepCopy-8     9622          8588          -10.75%

gc
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                    old ns/op     new ns/op     delta
BenchmarkSmallMemoryGC-8     10477         19217         +83.42%
BenchmarkBigMemoryGC-8       200133569     43880055      -78.07%

goroutine
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                         old ns/op     new ns/op     delta
BenchmarkGoroutineMergeSort-8     1631871       1396981       -14.39%

iterate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                        old ns/op     new ns/op     delta
BenchmarkRangeMap-8              18917         11757         -37.85%
BenchmarkRangeSlices-8           348           335           -3.82%
BenchmarkRangeArrays-8           337           345           +2.49%
BenchmarkRangeAndUseMap-8        15243         16028         +5.15%
BenchmarkRangeAndUseSlices-8     590           328           -44.49%
BenchmarkRangeAndUseArrays-8     674           758           +12.48%
BenchmarkForAndUseSlices-8       353           506           +43.40%
BenchmarkForAndUseArrays-8       357           339           -5.07%

sync
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark             old ns/op     new ns/op     delta
BenchmarkMutex-8      360           267           -25.92%
BenchmarkAtomic-8     291           268           -7.73%
```

## go1.14.3 vs go1.17.2
```
➜  benchmark-go git:(master) ✗ sh benchcmp.sh go1.14.3 go1.17.2
allocate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
BenchmarkMakeSmallSlice-8       43.6          51.4          +17.80%
BenchmarkMakeSmallMap-8         316           352           +11.49%
BenchmarkMakeSmallChannel-8     143           149           +4.27%
BenchmarkMakeBigSlice-8         96684         96606         -0.08%
BenchmarkMakeBigMap-8           2170528       2185768       +0.70%
BenchmarkMakeBigChannel-8       779263        896100        +14.99%
BenchmarkNewSmallObject-8       0.32          0.32          +0.89%
BenchmarkNewBigObject-8         656293        748333        +14.02%

channel
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                          old ns/op     new ns/op     delta
BenchmarkSendBlockChannel-8        190           244           +28.21%
BenchmarkGetBlockChannel-8         174           198           +13.62%
BenchmarkSendNonBlockChannel-8     52.2          50.0          -4.23%
BenchmarkGetNonBlockChannel-8      50.6          48.2          -4.76%

copy
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
Benchmark_GOBDeepCopy-8         104377        98527         -5.60%
Benchmark_ReflectDeepCopy-8     8594          8588          -0.07%

gc
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                    old ns/op     new ns/op     delta
BenchmarkSmallMemoryGC-8     18752         19217         +2.48%
BenchmarkBigMemoryGC-8       41593999      43880055      +5.50%

goroutine
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                         old ns/op     new ns/op     delta
BenchmarkGoroutineMergeSort-8     1726637       1396981       -19.09%

iterate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                        old ns/op     new ns/op     delta
BenchmarkRangeMap-8              18142         11757         -35.19%
BenchmarkRangeSlices-8           372           335           -10.03%
BenchmarkRangeArrays-8           363           345           -4.85%
BenchmarkRangeAndUseMap-8        13291         16028         +20.59%
BenchmarkRangeAndUseSlices-8     342           328           -4.24%
BenchmarkRangeAndUseArrays-8     752           758           +0.81%
BenchmarkForAndUseSlices-8       369           506           +37.18%
BenchmarkForAndUseArrays-8       481           339           -29.54%

sync
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark             old ns/op     new ns/op     delta
BenchmarkMutex-8      286           267           -6.75%
BenchmarkAtomic-8     279           268           -3.76%
```


## go1.16.5 vs go1.17.2

```
➜  benchmark-go git:(master) ✗ sh benchcmp.sh go1.16.5 go1.17.2
allocate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
BenchmarkMakeSmallSlice-8       37.4          51.4          +37.33%
BenchmarkMakeSmallMap-8         294           352           +19.87%
BenchmarkMakeSmallChannel-8     149           149           -0.07%
BenchmarkMakeBigSlice-8         168318        96606         -42.61%
BenchmarkMakeBigMap-8           2775862       2185768       -21.26%
BenchmarkMakeBigChannel-8       918614        896100        -2.45%
BenchmarkNewSmallObject-8       0.32          0.32          +0.03%
BenchmarkNewBigObject-8         688248        748333        +8.73%

channel
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                          old ns/op     new ns/op     delta
BenchmarkSendBlockChannel-8        209           244           +16.50%
BenchmarkGetBlockChannel-8         244           198           -19.04%
BenchmarkSendNonBlockChannel-8     52.6          50.0          -4.89%
BenchmarkGetNonBlockChannel-8      51.5          48.2          -6.48%

copy
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                       old ns/op     new ns/op     delta
Benchmark_GOBDeepCopy-8         172745        98527         -42.96%
Benchmark_ReflectDeepCopy-8     10261         8588          -16.30%

gc
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                    old ns/op     new ns/op     delta
BenchmarkSmallMemoryGC-8     13134         19217         +46.31%
BenchmarkBigMemoryGC-8       42074140      43880055      +4.29%

goroutine
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                         old ns/op     new ns/op     delta
BenchmarkGoroutineMergeSort-8     1827338       1396981       -23.55%

iterate
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                        old ns/op     new ns/op     delta
BenchmarkRangeMap-8              19482         11757         -39.65%
BenchmarkRangeSlices-8           407           335           -17.82%
BenchmarkRangeArrays-8           341           345           +1.20%
BenchmarkRangeAndUseMap-8        12933         16028         +23.93%
BenchmarkRangeAndUseSlices-8     328           328           -0.24%
BenchmarkRangeAndUseArrays-8     738           758           +2.78%
BenchmarkForAndUseSlices-8       441           506           +14.81%
BenchmarkForAndUseArrays-8       398           339           -14.81%

sync
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark             old ns/op     new ns/op     delta
BenchmarkMutex-8      324           267           -17.56%
BenchmarkAtomic-8     325           268           -17.36%

```



## Notes

- Thanks [CadmusJiang/TestGo](https://github.com/CadmusJiang/TestGo)
- This is for your reference only