spell makeAdder(grain base) {
    spell add(grain x) {
        yield base + x;
    }
    yield add;
}
seal addTwo -> makeAdder(2);
seal addTen -> makeAdder(10);
reveal addTwo(3);  /> expect: 5 </
reveal addTen(3);  /> expect: 13 </
reveal addTwo(3);  /> expect: 5 </