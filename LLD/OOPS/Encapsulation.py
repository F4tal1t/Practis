'''
Encapsulation is the process of combining data and methods for working with the data into a single unit called a class. 
It makes it possible to hide a class's implementation details from outside users who engage with the class via its public interface.
'''

class Employee:
    def __init__(self):
        self.__id = None
        self.__name = None

    def set_id(self, id):
        self.__id = id

    def set_name(self, name):
        self.__name = name

    def get_id(self):
        return self.__id

    def get_name(self):
        return self.__name


emp = Employee()

# Using setters
emp.set_id(101)
emp.set_name("Geek")

# Using getters
print(f'Employee ID: {emp.get_id()}')
print(f'Employee Name: {emp.get_name()}')