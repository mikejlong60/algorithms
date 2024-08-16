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

// Macro to generate functions
#define DEFINE_OPERATION(T1, T2, R, NAME) R NAME##_##T1##_##T2(T1 a, T2 b)

// Define specific operations
DEFINE_OPERATION(int, double, double, foldr) {
    float r = a + b;
    printf("Result (int + double to float type) : %.2f\n", r);
    return r;
}

DEFINE_OPERATION(double, int, int, foldr) {
    int r = a + b;
    printf("Result (double + int to int type): %d\n", r);
    return r;
}


DEFINE_OPERATION(char, char, char*, foldr) {
    //char r = sprintf("a:%s --- b%s\n", a ,b);

    char* r = "fred";
    printf("%s\n", r);
    return r;
}


// Macro to call the correct function
#define OPERATE(T1, T2, a, b) NAME##_##T1##_##T2(a, b)

int main() {
    int x = 5;
    double y = 3.14;

    double result1 = foldr_int_double(x, y);

    double a = 2.5;
    int z = 4;
    int result3 = foldr_double_int(a, z);

    char b = 1;
    char z1 = 2;
    //int (*)[5] = &b;
    char *result4 = foldr_char_char(b, z1);
    printf("result4:%s\n", result4);

    return 0;
}
