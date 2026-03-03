// 'this' keyword

this.name = 'Global';
function getParam(){
    console.log(this.name);
}

let user = {
    name: "Dibby",
    getName(){
        console.log(this.name);
    },
}

user.getName();

// Implicit Binding
// Implicit binding occurs when a function is called as a method of an object. In this case, 'this' refers to the object that the method is called on.

// Explicit Binding
// Explicit binding occurs when we use the call(), apply(), or bind() methods to explicitly set the value of 'this' for a function.

// Q1
const person = {
    name: "Dibby",
    greet() {
        return `Hello, my name is ${this.name}`;
    },
    farewell: () => {
        return `Goodbye from ${this.name}`;
    }
}

console.log(person.greet()); // Output: Hello, my name is Dibby 
console.log(person.farewell()); // Output: Goodbye from undefined   



// Q2 - Calculator

let calculator = {
    read(){
        this.a = +prompt('a=',0);
        this.b = +prompt('b=',0);
    },
    sum(){
        return this.a + this.b;
    },
    mul(){
        return this.a * this.b;
    }
}


// Q3 - Callback

var length = 4;
function callback(){
    console.log(this.length);
}
const obj = {
    length: 5,
    method(callback){
        callback();
    }
}
obj.method(callback); // Output: 4

