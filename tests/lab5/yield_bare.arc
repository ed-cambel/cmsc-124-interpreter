spell early(grain n) {
    when (n < 0) {
        yield;
    }
    yield n;
}
reveal early(-1); /> expect: nil </
reveal early(5);  /> expect: 5 </