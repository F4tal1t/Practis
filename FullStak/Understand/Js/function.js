// Functions in js
// Q1 : Function Declaration
function greet(x) {
    console.log("Hello, World!", x);
}

// Q2 : Function Expression
const greetExpression = function(x) {
    console.log("Hello, World!", x);
}

// Q3 : First-Class Functions
function firstClassExample() {
    return function() {
        console.log("I am a first-class function!");
    }}
 
// Q4 : IIFE (Immediately Invoked Function Expression)
(function() {
    console.log("This is an IIFE!");
})();


//Q5 : Function Hoisting
hoistedFunction(); // This will work because of hoisting
function hoistedFunction() {
    console.log("This function is hoisted!");
}

// Q6 : Function Hoisting O/P question

var x = 10;
var fun = function () {
    console.log(x);
    var x = 20;
}// Output: undefined
fun();

// Q7 : Parameters vs Arguments
function square(num) { // num is a parameter
    return num * num;
}  
console.log(square(5)); // 5 is an argument


// Q8 : Rest Parameters
function sum(...numbers) {
    return numbers.reduce((acc, curr) => acc + curr, 0);
}
console.log(sum(1, 2, 3, 4)); // Output: 10

// Q9 : Callback Functions
function fetchData(callback) {
    setTimeout(() => {
        const data = "Data fetched!";
        callback(data);
    }, 1000);} 

// Q10 : Arrow Functions
const arrowFunction = (x) => console.log("Hello, World!", x);
