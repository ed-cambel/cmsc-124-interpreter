spell double(grain n) {
    yield n * 2;
}
spell addOne(grain n) {
    yield n + 1;
}
reveal double(addOne(double(3))); /> expect: 14 </