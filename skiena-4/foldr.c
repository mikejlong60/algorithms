// The efficiency of this algorithm is O(N) but it reverses the list.  Use FoldLeft instead if you don't want this.
//func FoldRight[T1, T2 any](as []T1, z T2, f func(T1, T2) T2) T2 {
//	if len(as) > 1 { //Slice has a head and a tail.
//		h, t := as[0], as[1:len(as)]
//		return f(h, FoldRight(t, z, f))
//	} else if len(as) == 1 { //Slice has a head and an empty tail.
//		h := as[0]
//		return f(h, FoldRight(Zero[T1](), z, f))
//	}
//	return z
//}

#include <stdio.h>

#define T2 GENERIC_FOLDR(T1, T2) \
 ##T2 foldr_##T1##_##T2(T1 xs, T2 z) { \
    printf(#T1 ": %d, " #T2 ": %.2f\n", xs, z); \
    return z; \
}

// Generate specific functions
double GENERIC_FOLDR(int, double)
int GENERIC_FOLDR(double, int)

int main() {
    foldr_int_double(5, 3.14);
    foldr_double_int(2.71, 10);

    return 0;
}