const __gojs_version = 0.1;

exports.append = function(xs, x) {
    xs.push(x);
    return xs;
};

exports.makemap = function(xs) {
    var obj = {};
    for (var i in xs) {
        obj[xs[i][0]] = xs[i][1];
    }
};