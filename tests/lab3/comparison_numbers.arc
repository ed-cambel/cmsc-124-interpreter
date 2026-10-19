2 < 5;           /> expect: boon </
5 < 2;           /> expect: bane </
2 < 2;           /> expect: bane </
2 <= 5;          /> expect: boon </
5 <= 2;          /> expect: bane </
2 <= 2;          /> expect: boon </
5 > 2;           /> expect: boon </
2 > 5;           /> expect: bane </
2 > 2;           /> expect: bane </
5 >= 2;          /> expect: boon </
2 >= 5;          /> expect: bane </
2 >= 2;          /> expect: boon </
-5 < -2;         /> expect: boon </
-2 < -5;         /> expect: bane </
-3 <= -3;        /> expect: boon </
-1 < 1;          /> expect: boon </
0 <= 0;          /> expect: boon </
0 > -1;          /> expect: boon </
0.0 <= -0.0;     /> expect: boon </
0.0 < -0.0;      /> expect: bane </
2.5 < 5.5;       /> expect: boon </
2.5 < 2.5;       /> expect: bane </
2.5 <= 2.5;      /> expect: boon </
-2.5 < -1.5;     /> expect: boon </
0.1 < 0.2;       /> expect: boon </
0.5 > 0;         /> expect: boon </
(1 + 1) < 3;     /> expect: boon </
2 * 3 > 5;       /> expect: boon </
(2 + 3) <= 5;    /> expect: boon </