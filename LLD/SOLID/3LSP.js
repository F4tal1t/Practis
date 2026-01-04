// Liskov Substitution Principle (LSP)
// Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.

class Recangle {
    constructor(width, height) {
        this.width = width;
        this.height = height;
    }
    getArea() {
        return this.width * this.height;
    }
    getWidth() {
        return this.width;
    }
    getHeight() {
        return this.height;
    }
    setWidth(width) {
        this.width = width;
    }
    setHeight(height) {
        this.height = height;
    }
}

// Derived class for square
class Square extends Recangle {
    constructor(side) {
        super(side, side);
    }
    setWidth(w){
        this.width = this.height = w; // Ensure both width and height remain the same
    }
}

/*
Rectangle Class:
This is the base class that has properties for width and height. It has methods for calculating the area and for setting width and height.

Square Class: 
This class inherits from Rectangle but overrides the setWidth and setHeight methods to ensure that changing one dimension affects the other, maintaining the property that all sides are equal.
*/