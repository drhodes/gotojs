// package main
var ifsquare = function(x, y) {
    if (x < 0) {
        return x * y;
    }
    return y * x;
}
var main = function() {
    ifsquare(3, 4);
}
main();
