var lib = require('../../lib.js');
var main = function() {
    var temp = [];
    temp = lib.append(temp, "asdf");
    console.log(temp[0]);
}
main();