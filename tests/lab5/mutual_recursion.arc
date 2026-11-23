spell isEven(grain n) {
    when (n == 0) yield boon;
    yield isOdd(n - 1);
}
spell isOdd(grain n) {
    when (n == 0) yield bane;
    yield isEven(n - 1);
}
reveal isEven(4); /> expect: boon </
reveal isOdd(4);  /> expect: bane </