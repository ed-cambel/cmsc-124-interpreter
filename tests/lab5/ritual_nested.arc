grain row -> 0;
grain count -> 0;
ritual (row < 3) {
    grain col -> 0;
    ritual (col < 4) {
        count -> count + 1;
        col -> col + 1;
    }
    row -> row + 1;
}
reveal count;       /> expect: 12 </