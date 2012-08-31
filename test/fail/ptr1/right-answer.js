// package main

function ptr_square(n_ptr) {
    var val = n_ptr();
	n_ptr(val * val);
};

function main() {
    var n_ptr = function() {
        var val = null;
        var f = function(b) {
            if (arguments.length > 0) {
                val = b;
            }
            return val;
        };
        return f;
    }();

    n_ptr(4);
    console.log(n_ptr());
	ptr_square(n_ptr);
    console.log(n_ptr());
}
main();
