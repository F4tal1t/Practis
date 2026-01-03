class Calculator {
    // Method for adding two integers
    add(a, b) {
        return a + b;
    }

    // Overloaded method for adding three integers
    add(a, b, c) {
        if (arguments.length === 3) {
            return a + b + c;
        }
        return this.add(a, b);
    }

    // Overloaded method for adding two doubles
    add(a, b) {
        if (typeof a === 'number' && typeof b === 'number') {
            return a + b;
        }
        return this.add(a, b);
    }
}

const myCalculator = new Calculator();

// Example usage of the overloaded methods
const sum1 = myCalculator.add(5, 3);
const sum2 = myCalculator.add(4, 6, 2);
const sum3 = myCalculator.add(3.5, 2.7);

// Output the results
console.log('Sum of 5 and 3:', sum1);
console.log('Sum of 4, 6, and 2:', sum2);
console.log('Sum of 3.5 and 2.7:', sum3);