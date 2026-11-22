grain a -> 0;
grain b -> 0;
a -> b -> 3;
reveal a; /> expect: 3 </
reveal b; /> expect: 3 </
grain c -> (a -> 5) + 1;
reveal c; /> expect: 6 </
reveal a; /> expect: 5 </