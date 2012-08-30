// package main

function main() {
    for (var a in [1, 2, 3, 4]) {
        var b = [1, 2, 3, 4][a];
        log.Println(a + b);
    }
}
