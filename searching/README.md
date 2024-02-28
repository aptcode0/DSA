# Searching Algorithms

| Algorithm            | Time Complexity                          | Space Complexity |
|----------------------|------------------------------------------|------------------|
| Linear Search        | O(n)                                     | O(1)             |
| Binary Search        | O(log n)                                 | O(1)             |
| Interpolation Search | O(log log n) on average, O(n) worst-case | O(1)             |
| Jump Search          | O(√n)                                    | O(1)             |
| Exponential Search   | O(log n)                                 | O(1)             |
| Fibonacci Search     | O(log n)                                 | O(1)             |
| Ternary Search       | O(log3 n)                                | O(1)             |

- **Linear Search**:
    - **When to Use**:
        - Use when the list is small or unsorted.
    - **Pros**:
        - Simple and easy to implement.
        - Works on unsorted arrays.
    - **Cons**:
        - Inefficient for large or sorted arrays.
        - Performs poorly compared to other algorithms for large datasets.

- **Binary Search**:
    - **When to Use**:
        - Use when the list is sorted.
    - **Pros**:
        - Efficient for large sorted arrays.
        - Guarantees a logarithmic time complexity.
    - **Cons**:
        - Requires a sorted array as input.
        - Not suitable for dynamic datasets that frequently change.

- **Interpolation Search**:
    - **When to Use**:
        - Use when the values in the array are uniformly distributed.
    - **Pros**:
        - Faster than binary search for uniformly distributed data.
    - **Cons**:
        - Inefficient for non-uniformly distributed data.
        - Requires uniformly distributed data to achieve its optimal time complexity.

- **Jump Search**:
    - **When to Use**:
        - Use when the list is sorted.
    - **Pros**:
        - Faster than linear search for large arrays.
        - Works well with uniformly distributed data.
    - **Cons**:
        - Requires a sorted array as input.
        - Not as efficient as binary search.

- **Exponential Search**:
    - **When to Use**:
        - Use when the list is sorted and the size of the array is unknown.
    - **Pros**:
        - Efficient for large sorted arrays when the upper bound is unknown.
    - **Cons**:
        - Requires a sorted array as input.
        - Performs poorly for small arrays.

- **Fibonacci Search**:
    - **When to Use**:
        - Use when the list is sorted and the cost of comparisons is high.
    - **Pros**:
        - Efficient for searching in large sorted arrays.
        - Works well when the distribution of elements is not known.
    - **Cons**:
        - Requires a sorted array as input.
        - Not as efficient as binary search in some scenarios.

- **Ternary Search**:
    - **When to Use**:
        - Use when the list is sorted and you want to find an element that satisfies some condition.
    - **Pros**:
        - Can be more efficient than binary search for certain types of functions or data distributions.
        - Works well for finding the maximum or minimum of a unimodal function.
    - **Cons**:
        - Not suitable for all types of data distributions.
        - Requires the function to be unimodal (having a single maximum or minimum) for efficient operation.
        - May require more comparisons compared to binary search in some cases.
