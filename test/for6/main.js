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
        var adds = append(adds, new Address(21, "Jump st", "John Doe", 10001));
    }
}
main();

