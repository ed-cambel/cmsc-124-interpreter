/> right operand must be skipped; each would raise "Division by zero." if evaluated </
reveal bane && (1 / 0);      /> expect: bane </
reveal nil && (1 / 0);       /> expect: nil </
reveal 0 && (1 / 0);         /> expect: 0 </
reveal boon || (1 / 0);       /> expect: boon </
reveal "hi" || (1 / 0);       /> expect: hi </
reveal 5 || (1 / 0);          /> expect: 5 </
/> chains stop at the first deciding operand </
reveal 1 && nil && (1 / 0); /> expect: nil </
reveal nil || bane || "x";    /> expect: x </