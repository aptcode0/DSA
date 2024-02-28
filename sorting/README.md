| Sorting Algorithm | Time Complexity (Best, Average, Worst) | Space Complexity | Stability          |
|-------------------|----------------------------------------|------------------|--------------------|
| Bubble Sort       | O(n), O(n^2), O(n^2)                   | O(1)             | Stable             |
| Selection Sort    | O(n^2), O(n^2), O(n^2)                 | O(1)             | Unstable           |
| Insertion Sort    | O(n), O(n^2), O(n^2)                   | O(1)             | Stable             |
| Merge Sort        | O(n*log(n)), O(n*log(n)), O(n*log(n))  | O(n)             | Stable             |
| Quick Sort        | O(n*log(n)), O(n*log(n)), O(n^2)       | O(log(n)), O(n)  | Generally Unstable |
| Heap Sort         | O(n*log(n)), O(n*log(n)), O(n*log(n))  | O(1)             | Unstable           |
| Counting Sort     | O(n + k)                               | O(n + k)         | Stable             |
| Radix Sort        | O(d * (n + k))                         | O(n + k)         | Stable             |
| Bucket Sort       | O(n + k), O(n + k), O(n^2)             | O(n + k)         | Generally Unstable |

d: It denotes the number of digits in the largest number present in the array. e.g. num: 1456, d=4
k: It denotes the range of values for the elements being sorted. e.g. array elements 0 to 99, k = 100