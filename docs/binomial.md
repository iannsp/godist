#Binomial Distribution Notes

## Int overflow
The use of int is to enforce the need to handle the overflow situation.

## Optimizations

1. Multiplication Formula X Factorial Formula  
Using the factorial formula the limit is 20! with 21! overflows while the binomial  
coefficient often fits in a integer.  
Change to mutiplicative formula increase this limit.  
 
2. Symmetry property of combinations  
This property states that the number of ways to choose k items from a set of   
n distinct items is the same as the number of ways to choose n-k items from the   
same set of n distinct items.   

$C(n, k) = n! / (k! * (n - k)!)$

$C(n, n-k) = n! / ((n - k)! * (n - (n - k))!)$

$C(n, n-k) = n! / ((n - k)! * k!)$
