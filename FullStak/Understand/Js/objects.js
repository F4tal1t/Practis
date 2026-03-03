// Objects

const user = {
    name: "Dibby",
    age: 25,
    city: "Bengaluru",
    isStudent: true,
}

user.name = "Bibby";

delete user.isStudent;


// Q1 = What is JSON.stringify() and JSON.parse()?

const use = {
    name: "Dibby",
    age:20
};

console.log(JSON.stringify(user)); // Output: {"name":"Dibby","age":20}

// Q2 = Shallow copy vs Deep copy

const userObj = {
    name: "Dibby",
    age: 20
}

// Shallow copy
const shallowCopy = {...userObj};
shallowCopy.name = "Bibby";

console.log(userObj); // Output: {"name":"Dibby","age":20}
console.log(shallowCopy); // Output: {"name":"Bibby","age":20}

// Deep copy
const deepCopy = JSON.parse(JSON.stringify(userObj));
deepCopy.name = "Bibby";