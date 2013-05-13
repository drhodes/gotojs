var lib = require('../../lib.js');
var square = function(x) {
    return x * x;
}
var main = function() {
    var __defer_stack = [];
    var __retvals = function() {
        __defer_stack.push(function() {
            square(4)
        })
        var temp = square(4);
        temp = square(temp);
    }();

    while (__defer_stack.length != 0) {
        __defer_stack.pop()();
    }
    return __retvals;
}

main();
