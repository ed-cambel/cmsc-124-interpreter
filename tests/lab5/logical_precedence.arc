reveal boon || bane && bane;   /> expect: boon </
reveal bane && bane || boon;   /> expect: boon </
reveal (bane || boon) && bane; /> expect: bane </
reveal 1 == 1 && 2 == 3;       /> expect: bane </
reveal 1 < 2 && 3 > 2;         /> expect: boon </
reveal 1 && 2 && 3;           /> expect: 3 </