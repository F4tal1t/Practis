// Currying
// Example f(a,b) => f(a)(b)

function f(a){
    return function(b){
        return '${a} ${b}';
    }
}
console.log(f("Hello")("World")); // Output: "Hello World"

function sum(a){
    return function(b){
        return function(c){
            return a + b + c;
        }
    }    
}
console.log(sum(1)(2)(3)); // Output: 6