spell fact(grain n) {
    when (n <= 1) yield 1;
    yield n * fact(n - 1);
}
spell fib(grain n) {
    when (n <= 1) yield n;
    yield fib(n - 1) + fib(n - 2);
}
reveal fact(5); /> expect: 120 </
reveal fib(10); /> expect: 55 </