# False Sharing with C

Example of False Sharing using C.

* Build `make`
* Run `./fs`

You can limit the number of pthreads used with `CMAXPROCS`
```
CMAXPROCS=2 ./fs
```

## Results
```bash
$ CMAXPROCS=2 ./fs
CPUS=2
benchLinear ns/op = 269212630.800000
sumParallelFalseSharing ns/op = 398702100.100000
sumParallelPadded ns/op = 173627799.100000
```