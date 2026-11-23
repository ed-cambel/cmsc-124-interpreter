grain i -> 0;
grain sum -> 0;
ritual (i < 4) {
    sum -> sum + i;
    i -> i + 1;
}
reveal sum; /> expect: 6 </
reveal i;   /> expect: 4 </