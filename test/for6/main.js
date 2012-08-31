// package main

function Address(Number, Street, Name, ZipCode) {
    this.Number = Number;
    this.Street = Street;
    this.Name = Name;
    this.ZipCode = ZipCode;
}

function main() {
    var adds = [];
    for (var i = 0; i < 10; i++) {
        var temp = new Address(21, "Jump st", "John Doe", 10001);
        var adds = append(adds, temp);
        console.log(adds);
    }
}
main();

