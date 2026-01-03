# Superclass (Parent)
class Animal:
    def eat(self):
        print("Animal is eating...")

    def sleep(self):
        print("Animal is sleeping...")

# Subclass (Child) - Inherits from Animal
class Dog(Animal):
    def bark(self):
        print("Dog is barking!")

# Main
if __name__ == '__main__':
    myDog = Dog()

    # Inherited methods (from Animal)
    myDog.eat()
    myDog.sleep()

    # Child class method
    myDog.bark()