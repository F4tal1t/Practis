// Open-Closed Principle (OCP)
// Software entities (classes, modules, functions, etc.) should be open for extension,
// but closed for modification.

// Base class for payment processing
class PaymentProcessor {
    processPayment(amount) { // Pure virtual function
        throw new Error("processPayment must be implemented");
    }
}

// Credit card payment processor
class CreditCardPaymentProcessor extends PaymentProcessor {
    processPayment(amount) {
        console.log(`Processing credit card payment of $${amount}`);
    }
}


// PayPal payment processor
class PayPalPaymentProcessor extends PaymentProcessor {
    processPayment(amount) {
        console.log(`Processing PayPal payment of $${amount}`);
    }
}



// Function to process payment
function processPayment(processor, amount) {
    processor.processPayment(amount);
}

function main() {
    const creditCardProcessor = new CreditCardPaymentProcessor();
    const payPalProcessor = new PayPalPaymentProcessor();

    processPayment(creditCardProcessor, 100.00); // Processing credit card payment
    processPayment(payPalProcessor, 150.00);     // Processing PayPal payment
}

main();

/* 
Base Class (PaymentProcessor): 
This is an abstract base class with a pure virtual function processPayment().
It defines a common interface for all payment processors.

CreditCardPaymentProcessor: 
This class implements the payment processing logic for credit card payments.

PayPalPaymentProcessor: 
This new class extends the functionality by implementing the payment processing for PayPal payments.

Main Function: 
The processPayment function takes a pointer to a PaymentProcessor and calls the processPayment() method.
This allows you to use any processor that implements the PaymentProcessor interface without changing existing code.
*/