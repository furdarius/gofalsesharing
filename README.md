# False Sharing with Go

False sharing is a common problem in shared memory parallel processing. It occurs when two or more cores hold a copy of the same memory cache line.

If one core writes, the cache line holding the memory line is invalidated on other cores. Even though another core may not be using that data (reading or writing), it may be using another element of data on the same cache line. The second core will need to reload the line before it can access its own data again.

The cache hardware ensures data coherency, but at a potentially high performance cost if false sharing is frequent. A good technique to identify false sharing problems is to catch unexpected sharp increases in last-level cache misses using hardware counters or other performance tools.

<p align="center"><img src="https://github.com/furdarius/gofalsesharing/blob/master/schema.svg"></p>

## Benchmark
Summarize elements of array (size `10^8`)

```
go test -run=XXX -bench=. -cpu=1,2,4,6,8,12,16,24,32,56 -benchtime=10s
```
#### Results
```
BenchmarkSum/Linear            	                     100	 118133518 ns/op
BenchmarkSum/Linear-2          	                     100	 123964604 ns/op
BenchmarkSum/Linear-4          	                     100	 112477528 ns/op
BenchmarkSum/Linear-6          	                     100	 123335032 ns/op
BenchmarkSum/Linear-8          	                     100	 123343898 ns/op
BenchmarkSum/Linear-12          	             100	 110501346 ns/op
BenchmarkSum/Linear-16          	             100	 120919665 ns/op
BenchmarkSum/Linear-24          	             100	 120565232 ns/op
BenchmarkSum/Linear-32          	             100	 116581446 ns/op
BenchmarkSum/Linear-56          	             100	 108527032 ns/op
BenchmarkSum/ParallelFalseSharing            	     100	 231289258 ns/op
BenchmarkSum/ParallelFalseSharing-2          	     100	 117786360 ns/op
BenchmarkSum/ParallelFalseSharing-4          	     200	  64357195 ns/op
BenchmarkSum/ParallelFalseSharing-6          	     300	  47391438 ns/op
BenchmarkSum/ParallelFalseSharing-8          	     500	  37229853 ns/op
BenchmarkSum/ParallelFalseSharing-12         	     500	  27098008 ns/op
BenchmarkSum/ParallelFalseSharing-16         	    1000	  22183358 ns/op
BenchmarkSum/ParallelFalseSharing-24         	    1000	  18418561 ns/op
BenchmarkSum/ParallelFalseSharing-32         	    1000	  16435079 ns/op
BenchmarkSum/ParallelFalseSharing-56         	    1000	  14559299 ns/op
BenchmarkSum/ParallelWithPadding             	     100	 229699936 ns/op
BenchmarkSum/ParallelWithPadding-2           	     100	 118146717 ns/op
BenchmarkSum/ParallelWithPadding-4           	     200	  59917481 ns/op
BenchmarkSum/ParallelWithPadding-6           	     300	  42033348 ns/op
BenchmarkSum/ParallelWithPadding-8           	     500	  30706079 ns/op
BenchmarkSum/ParallelWithPadding-12          	    1000	  21592191 ns/op
BenchmarkSum/ParallelWithPadding-16          	    1000	  17484888 ns/op
BenchmarkSum/ParallelWithPadding-24          	    1000	  13178152 ns/op
BenchmarkSum/ParallelWithPadding-32          	    2000	   9742292 ns/op
BenchmarkSum/ParallelWithPadding-56          	    2000	   9075207 ns/op
BenchmarkSum/ParallelLocalVariable            	     300	  47909122 ns/op
BenchmarkSum/ParallelLocalVariable-2          	     500	  24776753 ns/op
BenchmarkSum/ParallelLocalVariable-4          	    1000	  12702117 ns/op
BenchmarkSum/ParallelLocalVariable-6          	    2000	   8882727 ns/op
BenchmarkSum/ParallelLocalVariable-8          	    2000	   6479442 ns/op
BenchmarkSum/ParallelLocalVariable-12         	    3000	   4589380 ns/op
BenchmarkSum/ParallelLocalVariable-16         	    5000	   3779905 ns/op
BenchmarkSum/ParallelLocalVariable-24         	    5000	   3476589 ns/op
BenchmarkSum/ParallelLocalVariable-32         	    5000	   3301371 ns/op
BenchmarkSum/ParallelLocalVariable-56         	    5000	   2962546 ns/op
```

<p align="center"><img src="https://docs.google.com/spreadsheets/d/e/2PACX-1vQH4D2eONdwx-z3joRZyTQjcI_mtvMQ0OJds81MY27k4J5HAFjv257Zgf1EfoyQT6qd0HraIbRP-hF0/pubchart?oid=1182271586&format=image"></p>

#### Go version and CPU
```
$ go version
go version go1.12 linux/amd64

$ lscpu 
Architecture:          x86_64
CPU op-mode(s):        32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                56
On-line CPU(s) list:   0-55
Thread(s) per core:    2
Core(s) per socket:    14
Socket(s):             2
NUMA node(s):          2
Vendor ID:             GenuineIntel
CPU family:            6
Model:                 63
Model name:            Intel(R) Xeon(R) CPU E5-2697 v3 @ 2.60GHz
Stepping:              2
CPU MHz:               2600.000
CPU max MHz:           2600.0000
CPU min MHz:           1200.0000
BogoMIPS:              5193.50
Virtualization:        VT-x
L1d cache:             32K
L1i cache:             32K
L2 cache:              256K
L3 cache:              35840K
NUMA node0 CPU(s):     0-13,28-41
NUMA node1 CPU(s):     14-27,42-55
Flags:                 fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid dca sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm epb invpcid_single kaiser tpr_shadow vnmi flexpriority ept vpid fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid cqm xsaveopt cqm_llc cqm_occup_llc dtherm ida arat pln pts
```

## False Sharing detection
Using linux perf `perf c2c`

Setup max sample rate
```bash 
# echo 100000 > /proc/sys/kernel/perf_event_max_sample_rate
```
Record `BenchmarkSum/ParallelFalseSharing`
```bash
# perf c2c record -F 60000 -a --all-user go test -run=XXX -bench=BenchmarkSum/ParallelFalseSharing -cpu=4 -benchtime=5s
# perf c2c report -NN --stdio

=================================================
            Trace Event Information              
=================================================
  Total records                     :     591181
  Locked Load/Store Operations      :      56871
  Load Operations                   :     260941
  Loads - uncacheable               :          0
  Loads - IO                        :          0
  Loads - Miss                      :        452
  Loads - no mapping                :        718
  Load Fill Buffer Hit              :      65782
  Load L1D hit                      :     188522
  Load L2D hit                      :       1363
  Load LLC hit                      :       2167
  Load Local HITM                   :         43
  Load Remote HITM                  :          0
  Load Remote HIT                   :          0
  Load Local DRAM                   :       1937
  Load Remote DRAM                  :          0
  Load MESI State Exclusive         :       1937
  Load MESI State Shared            :          0
  Load LLC Misses                   :       1937
  LLC Misses to Local DRAM          :      100.0%
  LLC Misses to Remote DRAM         :        0.0%
  LLC Misses to Remote cache (HIT)  :        0.0%
  LLC Misses to Remote cache (HITM) :        0.0%
  Store Operations                  :     330240
  Store - uncacheable               :          0
  Store - no mapping                :        187
  Store L1D Hit                     :     315973
  Store L1D Miss                    :      14080
  No Page Map Rejects               :        688
  Unable to parse data source       :          0

=================================================
    Global Shared Cache Line Event Information   
=================================================
  Total Shared Cache Lines          :         35
  Load HITs on shared lines         :      34424
  Fill Buffer Hits on shared lines  :      10786
  L1D hits on shared lines          :      23583
  L2D hits on shared lines          :          2
  LLC hits on shared lines          :         49
  Locked Access on shared lines     :         63
  Store HITs on shared lines        :      38356
  Store L1D hits on shared lines    :      35934
  Total Merged records              :      38399

=================================================
```

## Reference
* https://software.intel.com/en-us/articles/avoiding-and-identifying-false-sharing-among-threads
* https://scc.ustc.edu.cn/zlsc/sugon/intel/compiler_c/main_cls/cref_cls/common/cilk_false_sharing.htm
* https://parallelcomputing2017.wordpress.com/2017/03/17/understanding-false-sharing/
* https://lrita.github.io/images/posts/debug/moc.apr.2017.c2c.pdf
* https://joemario.github.io/blog/2016/09/01/c2c-blog/