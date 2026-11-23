grain x -> 1;
spell f() {
    grain x -> 2;
    {
        grain x -> 3;
        yield x;
    }
}
reveal f(); /> expect: 3 </
reveal x;   /> expect: 1 </