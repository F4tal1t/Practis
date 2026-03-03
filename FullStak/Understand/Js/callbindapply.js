// Call, Bind, Apply

// Q1 - What is call

var obj = {
    name: "Dibby",
}
function sayHello(age){
    return "Hello " + this.name + " you are " + age;
}

console.log(sayHello.call(obj, 25))


// Q2 - What is Apply