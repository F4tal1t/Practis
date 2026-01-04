// Interface Segregation Principle (ISP)
// Clients should not be forced to depend on interfaces they do not use.

// Interface simulation is done by convention in JavaScript

// Class for vegetarian menu
class VegetarianMenu {
    getVegetarianItems() {
        return ["Vegetable Curry", "Paneer Tikka", "Salad"];
    }
}

// Class for non-vegetarian menu
class NonVegetarianMenu {
    getNonVegetarianItems() {
        return ["Chicken Curry", "Fish Fry", "Mutton Biryani"];
    }
}

// Class for drinks menu
class DrinkMenu {
    getDrinkItems() {
        return ["Water", "Soda", "Juice"];
    }
}

// Function to display menu items for a vegetarian customer
function displayVegetarianMenu(menu) {
    console.log("Vegetarian Menu:");
    for (const item of menu.getVegetarianItems()) {
        console.log("- " + item);
    }
}

// Function to display menu items for a non-vegetarian customer
function displayNonVegetarianMenu(menu) {
    console.log("Non-Vegetarian Menu:");
    for (const item of menu.getNonVegetarianItems()) {
        console.log("- " + item);
    }
}

function main() {
    const vegMenu = new VegetarianMenu();
    const nonVegMenu = new NonVegetarianMenu();
    const drinkMenu = new DrinkMenu();

    displayVegetarianMenu(vegMenu);
    displayNonVegetarianMenu(nonVegMenu);
}

main();

/*
IVegetarianMenu Interface:
This interface defines a method to get vegetarian items. It ensures that only classes implementing vegetarian menus will need to provide this functionality.

INonVegetarianMenu Interface: 
Similar to the vegetarian interface, this one defines a method for getting non-vegetarian items.

IDrinkMenu Interface: 
This interface defines a method for getting drink items, keeping it separate from food items.

VegetarianMenu Class: 
Implements the IVegetarianMenu interface and provides a list of vegetarian items.

NonVegetarianMenu Class: 
Implements the INonVegetarianMenu interface and provides a list of non-vegetarian items.

DrinkMenu Class: 
Implements the IDrinkMenu interface and provides a list of drink items.
*/