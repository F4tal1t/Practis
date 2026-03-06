// Call, Bind, Apply

// Q1 - What is call

var obj = {
    name: "Dibby",
}
function sayHello(age , profession) {
    return "Hello " + this.name + " you are " + age + " and you are a " + profession;
}

console.log(sayHello.call(obj, 25, "Engineer"))


// Q2 - What is Apply
console.log(sayHello.apply(obj, [25, "Engineer"]))

// Q3 - What is Bind
var sayHelloBind = sayHello.bind(obj);
console.log(sayHelloBind(25, "Engineer"))
// Provides an reusable function 


// Q1 - call with functon inside object

const age = 10;
var person = {
    name: "Dibby",
    age: 20,
    getAge: function() {
        return this.age;
    }
};
var person2 = { age:30};
person.getAge.call(person2) // 24

// Q2 - O/P
var status = "Global";
setTimeout(()=>{
    var status = "Local"; 
    const dataa ={
        status: "Data",
        getStatus: function() {
            return this.status;
        }
    }
    console.log(dataa.getStatus()) // Data
    console.log(dataa.getStatus.call(this)) // Global
},0)

// Q3 - Bind Chaining
function f() {
    console.log(this.name)
}
f = f.find({name: "Dibby"}).bind({name: "Dibby2"})
f() // Dibby2
// Bind Chaining Not possible

