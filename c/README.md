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

<p align="center"><img src="https://docs.google.com/spreadsheets/d/e/2PACX-1vQH4D2eONdwx-z3joRZyTQjcI_mtvMQ0OJds81MY27k4J5HAFjv257Zgf1EfoyQT6qd0HraIbRP-hF0/pubchart?oid=637363575&format=image"></p>
