var lib = require('../../lib.js');

function Point(x, y) {
    this.x = x;
    this.y = y;
}
Point.prototype.Add = function(other) {
    this.x += other.x;
    this.y += other.y;
    return this;
}
Point.prototype.Equal = function(other) {
    return this.x == other.x && this.y == other.y;
}
var main = function() {
    var p1 = new Point(1, 0);
    var p2 = new Point(0, 1);
    var p3 = new Point(1, 1);
    console.log(p1.Add(p2).Equal(p3));
}
main();