var lib = require('../../lib.js');

var main = function() {
    var temp = [1, 2, 3, 4];
    for (var a in temp) {
        var b = temp[a];
        log.Println(a + b);
    }
}
main();