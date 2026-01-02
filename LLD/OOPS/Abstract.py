

from abc import ABC, abstractmethod

# Abstract class
class Vehicle(ABC):
    @abstractmethod
    def accelerate(self):
        pass

    @abstractmethod
    def brake(self):
        pass

    def startEngine(self):
        print("Engine started!")

# Concrete class
class Car(Vehicle):
    def accelerate(self):
        print("Car: Pressing gas pedal...")
        # Hidden complex logic: fuel injection, gear shifting, etc.

    def brake(self):
        print("Car: Applying brakes...")
        # Hidden logic: hydraulic pressure, brake pads, etc.

# Main execution
if __name__ == "__main__":
    myCar = Car()
    myCar.startEngine()
    myCar.accelerate()
    myCar.brake()