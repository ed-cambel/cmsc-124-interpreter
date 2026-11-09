boon || bane && bane;           /> expect: boon </
bane && bane || boon;           /> expect: boon </
(boon || bane) && bane;         /> expect: bane </
bane && (bane || boon);         /> expect: bane </
1 == 1 && 2 == 2;               /> expect: boon </
1 == 2 || 3 == 3;               /> expect: boon </
1 != 1 && 2 != 2;               /> expect: bane </