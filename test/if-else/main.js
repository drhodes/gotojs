// package main

function ifsquare(x, y) {
    if (x < 0) {
        return x * y;
    } else {
        return y * x;
    }
    return 0;
}

function main() {
    ifsquare(3, 4);
}
main();