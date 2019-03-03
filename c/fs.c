#include <time.h>
#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include <sys/sysinfo.h>
#include <string.h>

#define MIN(a,b) (((a)<(b))?(a):(b))
#define int64_t long long int
#define NSEC_PER_SEC 1000000000L

const int64_t N = 10e7;

int64_t CPUS;

void fill(int64_t arr[N])
{
	for (int64_t i = 0; i < N; ++i) {
		int64_t rnd = (rand() % (5 - 1)) + 1;

		arr[i] = rnd;
	}
}

int64_t maxprocs()
{
	const char* s = getenv("CMAXPROCS");
	if (s != NULL && strlen(s) > 0) {
		int64_t n = atoi(s);
		if (n <= 0) {
			return 1;
		}

		return n;
	}

	return get_nprocs();
}

int64_t sumLinear(int64_t arr[N])
{
	int64_t res = 0;
	for (int64_t i = 0; i < N; ++i) {
		res += arr[i];
	}
	return res;
}

struct thread_info_fs { 
   pthread_t 	thread_id;
   int64_t*     arr; 
   int64_t 		start;
   int64_t 		end;
   int64_t* 	res;
};

static void* thread_sum_fs(void *arg)
{
	struct thread_info_fs *thread = arg;

	for (int64_t i = (*thread).start; i < (*thread).end; ++i) {
		*((*thread).res) += (*thread).arr[i];
	}

	pthread_exit(0);
}

int64_t sumParallelFalseSharing(int64_t arr[N])
{
	int64_t* results = (int64_t*) calloc(CPUS, sizeof(int64_t));

	struct thread_info_fs *threads = (struct thread_info_fs*) calloc(CPUS, sizeof(struct thread_info_fs));
	int64_t blockSize = (N + CPUS - 1) / CPUS;

	for (int64_t i = 0; i < CPUS; ++i) {
		int64_t start = i * blockSize;
		int64_t end = MIN(blockSize*(i+1), N);

		threads[i].arr = arr;
		threads[i].start = start;
		threads[i].end = end;
		threads[i].res = &results[i];

		pthread_create(&threads[i].thread_id, NULL, thread_sum_fs, &threads[i]); 
	}

    for (int64_t i = 0; i < CPUS; ++i) {
		pthread_join(threads[i].thread_id, NULL); 
	}

	int64_t res = 0;
	for (int64_t i = 0; i < CPUS; ++i) {
		res += results[i];
	}

	return res;
}


struct res_padded {
	int64_t		res;
	int64_t 	_padding[7];
};

struct thread_info_padded { 
   pthread_t 			thread_id;
   int64_t*     		arr; 
   int64_t 				start;
   int64_t 				end;
   struct res_padded* 	res;
};

static void* thread_sum_padded(void *arg)
{
	struct thread_info_padded *thread = arg;

	for (int64_t i = (*thread).start; i < (*thread).end; ++i) {
		(*((*thread).res)).res += (*thread).arr[i];
	}

	pthread_exit(0);
}


int64_t sumParallelPadded(int64_t arr[N])
{
	struct res_padded* results = (struct res_padded*) calloc(CPUS, sizeof(struct res_padded));

	struct thread_info_padded *threads = (struct thread_info_padded*) calloc(CPUS, sizeof(struct thread_info_padded));
	int64_t blockSize = (N + CPUS - 1) / CPUS;

	for (int64_t i = 0; i < CPUS; ++i) {
		int64_t start = i * blockSize;
		int64_t end = MIN(blockSize*(i+1), N);

		threads[i].arr = arr;
		threads[i].start = start;
		threads[i].end = end;
		threads[i].res = &results[i];

		pthread_create(&threads[i].thread_id, NULL, thread_sum_padded, &threads[i]); 
	}

    for (int64_t i = 0; i < CPUS; ++i) {
		pthread_join(threads[i].thread_id, NULL); 
	}

	int64_t res = 0;
	for (int64_t i = 0; i < CPUS; ++i) {
		res += results[i].res;
	}

	return res;
}

const int64_t BENCH_N = 10;

int64_t res1, res2, res3;

double benchLinear(int64_t A[N])
{
	struct timespec start, finish;
	int64_t dur = 0;

	int64_t sum;
	for (int64_t i = 0; i < BENCH_N; ++i) {
		clock_gettime(CLOCK_MONOTONIC, &start);

		sum = sumLinear(A);

		clock_gettime(CLOCK_MONOTONIC, &finish);

		int64_t nsec = (finish.tv_sec - start.tv_sec) * NSEC_PER_SEC +
	       (finish.tv_nsec - start.tv_nsec);

	    dur += nsec;
	}
	res1 = sum;

	return (double)dur/(double)BENCH_N;
}


double benchSumParallelFalseSharing(int64_t A[N])
{
	struct timespec start, finish;
	int64_t dur = 0;

	int64_t sum;
	for (int64_t i = 0; i < BENCH_N; ++i) {
		clock_gettime(CLOCK_MONOTONIC, &start);

		sum = sumParallelFalseSharing(A);

		clock_gettime(CLOCK_MONOTONIC, &finish);

		int64_t nsec = (finish.tv_sec - start.tv_sec) * NSEC_PER_SEC +
	       (finish.tv_nsec - start.tv_nsec);

	    dur += nsec;
	}
	res2 = sum;

	return (double)dur/(double)BENCH_N;
}


double benchSumParallelPadded(int64_t A[N])
{
	struct timespec start, finish;
	int64_t dur = 0;

	int64_t sum;
	for (int64_t i = 0; i < BENCH_N; ++i) {
		clock_gettime(CLOCK_MONOTONIC, &start);

		sum = sumParallelPadded(A);

		clock_gettime(CLOCK_MONOTONIC, &finish);

		int64_t nsec = (finish.tv_sec - start.tv_sec) * NSEC_PER_SEC +
	       (finish.tv_nsec - start.tv_nsec);

	    dur += nsec;
	}
	res3 = sum;

	return (double)dur/(double)BENCH_N;
}


int main(void)
{
	srand(time(NULL));
	CPUS = maxprocs();
	printf("CPUS=%lld\n", CPUS);

	int64_t* A = (int64_t*) malloc(N*sizeof(int64_t));
	fill(A);

	 double lineasOpNs = benchLinear(A);
	 printf("benchLinear ns/op = %f\n", lineasOpNs);

	double sumParallelFalseSharingOpNs = benchSumParallelFalseSharing(A);
	printf("sumParallelFalseSharing ns/op = %f\n", sumParallelFalseSharingOpNs);

	 double sumParallelPaddedOpNs = benchSumParallelPadded(A);
	 printf("sumParallelPadded ns/op = %f\n", sumParallelPaddedOpNs);

	return 0;
}