var lib = require('../../lib.js');

function Address(Number, Street, Name, ZipCode) {
    this.Number = Number;
    this.Street = Street;
    this.Name = Name;
    this.ZipCode = ZipCode;
}
var main = function() {
    var adds = [];
    for (var i = 0; i < 10; i++) {
        var temp = new Address(21, "Jump st", "John Doe", 10001);
        adds = lib.append(adds, temp);
        console.log(adds);
    }
}
main();
