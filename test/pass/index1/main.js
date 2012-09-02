var lib = require('../../lib.js');
var main = function() {
    var temp = [1, 2, 3, 4, 5, 6];
    console.log(temp.slice(2, 5));
    console.log(temp.slice(0, 5));
    console.log(temp.slice(5, temp.length));
    console.log(temp.slice(0, temp.length));
}
main();