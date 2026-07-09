function findOdd(A) {
    let result = 0;
​
    for (const num of A) {
        result ^= num;
    }
​
    return result;
}