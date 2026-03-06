// Promises in Js


//Callbacks
function getData(data, callback) {
    setTimeout(() => {
        callback(data)
    }, 1000)
}

getData("Dibby", (data) => {
    console.log(data)
})


// Promises

console.log("Start")
const sub = new Promise((resolve, reject) => {
    setTimeout(()=>{
        const result = 10;
        if(result > 5) {
            resolve(result)
        } else {
            reject("Error: Result is less than 5")
        }
    },1000)
})
sub.then((result)=>{
    console.log("Result: " + result)
}).catch((error)=>{
    console.log(error)
})

console.log("End")

// Promise.resolve

Promise.resolve("Hello").then((result)=>{
    console.log(result)
})

// Promise.reject

Promise.reject("Error").catch((error)=>{
    console.log(error)
})


// Promise Chaining
function addTwo(num){
    console.log("Adding 2 to " + num)
    return Promise.resolve(num + 2)
}

addTwo(5).then((result)=>{
    console.log("Result after adding 2: " + result)
    return addTwo(result)
}).then((result)=>{
    console.log("Result after adding 2 again: " + result)
}).catch((error)=>{
    console.log(error)
})

// Promise Combination
// Promise.all
const p1 = Promise.resolve(10);
const p2 = Promise.resolve(20);
const p3 = Promise.resolve(30);
Promise.all([p1, p2, p3]).then((results)=>{
    console.log("Results from Promise.all: " + results)
})

// Promise.race()
Promise.race([p1, p2, p3]).then((result)=>{
    console.log("Result from Promise.race: " + result)
})

//Promise.any()  same as Promise.race but ignores rejected promises and only returns the first resolved promise


// Q1 : O/P

console.log("Start")
const fn = () => {
    new Promise((resolve, reject) => {
        console.log(1)
        resolve("Success")
    })
}
console.log("Middle")
fn().then((result)=>{
    console.log(result)
})
console.log("End")

// O/P : Start, Middle, 1, End, Success

