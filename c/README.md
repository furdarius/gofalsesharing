# False Sharing with C

Example of False Sharing using C.

* Build `make`
* Run `./fs`

You can limit the number of pthreads used with `CMAXPROCS`
```
CMAXPROCS=2 ./fs
```

## Results
Compiled with `-O3`. Measured as `ns/op`.

```bash
CPUS		sumLinear	sumParallelFalseSharing	sumParallelPadded
1		88156175	100437927		103768725
2		91034534	57754896		51013625
4		90501434	44793273		37108097
6		90821035	31686255		26418707
8		90870923	28737622		20511014
12		98200929	23287962		17670726
16		99128101	20745115		14921762
24		93731639	19948281		16111866
32		88920219	18738384		15477581
56		105746241	17935466		15875341
```

<p align="center"><img src="https://docs.google.com/spreadsheets/d/e/2PACX-1vQH4D2eONdwx-z3joRZyTQjcI_mtvMQ0OJds81MY27k4J5HAFjv257Zgf1EfoyQT6qd0HraIbRP-hF0/pubchart?oid=637363575&format=image"></p>
