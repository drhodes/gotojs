// this seems to work ok.
// Is it cheaper than {val:x, ptr:true} with
// call methods attached? Dunno.

function newptr(){
    var inner = function() {
        var val = null;
        var f = function(v) {
            if (arguments.length > 0) {
                val = v;
            }
            return val;
        };
        f.prototype.ptr = true;
        return f;
    };
    return inner();
}


var temp = newptr();
temp(4);

var temp2 = newptr();
temp2(5);

console.log(temp());
console.log(temp2());
console.log(temp());
console.log(temp.prototype.ptr);
