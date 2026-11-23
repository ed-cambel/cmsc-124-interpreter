/> a skipped right operand must not run its assignment </
grain x -> 0;
bane && (x -> 1);
reveal x;                       /> expect: 0 </
boon || (x -> 2);
reveal x;                       /> expect: 0 </
/> an evaluated right operand does run it </
grain y -> 0;
boon && (y -> 5);
reveal y;                       /> expect: 5 </
bane || (y -> 7);
reveal y;                       /> expect: 7 </