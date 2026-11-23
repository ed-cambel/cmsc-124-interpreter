spell find(grain target) {
    grain i -> 0;
    ritual (i < 10) {
        when (i == target) {
            yield i * 10;
        }
        i -> i + 1;
    }
    yield -1;
}
reveal find(3);  /> expect: 30 </
reveal find(99); /> expect: -1 </