// Closures in JavaScript
// A closure is a function that has access to its own scope, the outer function's scope, and the global scope.
// Closures are created whenever a function is created, at function creation time. 
//  They allow functions to access variables from an enclosing scope or environment 
//  even after it leaves the scope in which it was declared.

// Lexical Scoping
var outerVariable = 'I am from the outer scope';
function outerFunction() {
    var innerVariable = 'I am from the inner scope';
    function innerFunction() {
        console.log(outerVariable); // Accessing outer variable
        console.log(innerVariable); // Accessing inner variable
    }
    return innerFunction; // Returning the inner function
}


//Closure Scope Chain
var e = 10; // Global Scope
function sum(a){
    return function(b){
        // outer function scope
        return function(c){
            // inner function scope
            return a + b + c + e; // Accessing variables from all scopes
        }
    }
}
console.log(sum(1)(2)(3)); // Output: 16 (1 + 2 + 3 + 10)

// Q1 Wat will be logged in the console
let count = 0;
(function printCount() {
    if (count === 0) {
        let count = 1;
        console.log(count); // Output: 1
    }
    console.log(count);
})();

// Q2: Function 
function createBase(base) {
    return function(num) {
        console.log( base + num);
    }
}
var addSix = createBase(6);
addSix(10); // Output: 16
addSix(21); // Output: 27

// Q3 : Time Optimization with Closures 
function find(){
    let a = [];
    for (let i = 0; i < 100000; i++) {
        a[i] = i * i;
    }
    return function(index){
        console.log(a[index]);
    } 
}
const closure = find();
console.time("6")
closure(6)
console.timeEnd("6")
console.time("12")
closure(12)
console.timeEnd("12")

// Q4 : setTimeout O/P
function a(){
    for(var i = 0; i < 3; i++){
        function closure(i){
            setTimeout(function(){
                console.log(i);
            }, 1000);
        }
        closure(i);
    }
}// using closures we ignore var limitations
a(); // 3, 3, 3 var == function scope and setTimeout is asynchronous, so by the time the callbacks are executed, the loop has completed and i has the value 3.
// make it work with let

// Q5 : Counter with Closures
function counter(){
    let count = 0
    function add(inc){
        count += inc;
    }
    function retrieve(){
        console.log(count);
    }
    return {
        add,
        retrieve
    };
}
const c = counter();
c.add(5);
c.add(10);
c.retrieve(); // Output: 5

// Q6 : What is Module Pattern in JavaScript?

var Module = (function (){
    function privateMethod() {
        console.log("Priv");
    }
    return {
        publicMethod: function() {
            console.log("Pub");
        }
    };

})();
Module.publicMethod();
// Module.privateMethod();  Error: privateMethod is not defined

// Q7 : Implement Caching/Memoization with Closures
function myMemoize(fn) {
    const results = {};
    return function(...args){
        var argsCached = JSON.stringify(args);
        if (!results[argsCached]){
            results[argsCached] = fn.call(this, ...args);
            }
        return results[argsCached];
        }
    }
const product = (a, b) =>{
    for (let i = 0; i < 1000000000; i++) {} // Simulating a time-consuming operation
    return a * b;
} ;
const memoizedProduct = myMemoize(product);
console.time("First Call");
memoizedProduct(5, 6);
console.timeEnd("First Call");
console.time("Second Call");
memoizedProduct(5, 6);
console.timeEnd("Second Call");