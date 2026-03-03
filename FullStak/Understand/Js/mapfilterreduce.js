// Array Methods : Map , Filter , Reduce
const nums = [1, 2, 3, 4, 5];

const multiplyfour = nums.map((num,i,arr)=>
{   
    return num * 4 + i 
})

console.log(multiplyfour); // Output: [4, 8, 12, 16, 20]

const moreThanTwo = nums.filter((num,i,arr)=>{
    return num > 2
})

console.log(moreThanTwo);

const sum = nums.reduce((acc,sum,i,arr)=>{
    return acc + sum
},0);

console.log(sum);

// Map vs forEach
const arr = [1, 2, 3, 4, 5];

const mapResult = arr.map((num) => num * 2);
console.log(mapResult); // Output: [2, 4, 6, 8, 10]

const forEachResult = [];
arr.forEach((num) => forEachResult.push(num * 2));

// forEach does not return a new array, it modifies the existing one

// Q1 - Return only name of students is Capital

const students = [
    {name: "Alice", age: 20},
    {name: "Bob", age: 22},
    {name: "Charlie", age: 19},
    {name: "David", age: 21}
];

const names = students.map((stu)=> stu.name.toUpperCase());
console.log(names); // Output: ["ALICE", "BOB", "CHARLIE", "DAVID"]
// Q2 - Return only students whose age is above 20
const ageAbove20 = students.filter((stu)=> stu.age > 20);

// Q3 - Return the sum of ages of all studnets
const sumOfAges = students.reduce((acc,stu)=> acc + stu.age,0);