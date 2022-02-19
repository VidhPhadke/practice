#include <stdio.h>

int factorial(int) ;
int gcd(int,int) ;

int main(){
    printf("This are going great %d", gcd(4,2)) ;    
}

int factorial(int n) {
    if(n <= 1) {
        return 1 ;
    }
    return n * factorial(n-1) ;
}

int gcd(int m, int n)
  {
    if (n == 0) return m;
    return gcd(n, m % n);
  }