boon && boon;       /> expect: boon </
boon && bane;       /> expect: bane </
bane && boon;       /> expect: bane </
bane && bane;       /> expect: bane </
1 && 2;             /> expect: 2 </
"a" && "b";         /> expect: b </
0 && 2;             /> expect: 0 </
nil && 2;           /> expect: nil </
boon && nil;        /> expect: nil </
1 && 2 && 3;        /> expect: 3 </
1 && 0 && 3;        /> expect: 0 </
boon || boon;       /> expect: boon </
boon || bane;       /> expect: boon </
bane || boon;       /> expect: boon </
bane && bane;       /> expect: bane </
1 || 2;             /> expect: 1 </
"a" || "b";         /> expect: a </
0 || 2;             /> expect: 2 </
nil || "hi";        /> expect: hi </
bane || nil;        /> expect: nil </
nil || 0;           /> expect: 0 </
nil || bane || 3;   /> expect: 3 </