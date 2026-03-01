// var , (let , const (ES6))


// Scope

var a = 5;//Global Scope
console.log(a);

{
    const b = 10; //Block Scope For const and let
    console.log(b);
}


// Variable Shadowing
var x = 20; // Global Scope

function shadowingExample() {
    var x = 10; // Local Scope (shadows the global variable)
    console.log(x); // Output: 10
}

// Illegal Shadowing
let y = 30; // Global Scope

function illegalShadowing() {
    let y = 15; // This will cause an error because 'y' is already declared in the global scope
    console.log(y);
}

// Declaration
var c = 50; // Can be redeclared and updated without initialization
var c = 50; 

let a = 100; // Cannot be redeclared, but can be updated with initialization

const d = 200; // Cannot be redeclared or updated without initialization

// Hoisting
// JavaScript mechanism where variables and function declarations are moved to the top of their containing scope during the compilation phase. 
// This means that you can use variables and functions before they are declared in the code.

// Creation Phase: 
// During this phase, the JavaScript engine allocates memory for variables and functions. 
// For variables declared with var, they are initialized with undefined. For variables declared with let and const, they are not initialized and are in a "temporal dead zone" until their declaration is encountered in the code.

// Execution Phase: 
// During this phase, the JavaScript engine executes the code line by line. 
// If a variable is accessed before its declaration, it will result in a ReferenceError for let and const, while var will return undefined.
 
var e = 10; // This will be hoisted and initialized with undefined
console.log(e , f ,g ); // Output: undefined
let f = 20; // This will be hoisted but not initialized, resulting in a ReferenceError if accessed before declaration
const g = 30; // This will be hoisted but not initialized, resulting in a ReferenceError if accessed before declaration
// Temporal Dead Zone (TDZ)
// The temporal dead zone (TDZ) is a behavior in JavaScript that occurs when you try to access a variable declared with let or const before it has been initialized. During the TDZ, the variable is in an uninitialized state and cannot be accessed, resulting in a ReferenceError if you attempt to use it.