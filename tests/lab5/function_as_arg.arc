spell apply(spell f, grain x) {
    yield f(x);
}
spell double(grain n) {
    yield n * 2;
}
reveal apply(double, 4); /> expect: 8 </
