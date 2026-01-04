// Single Responsibility Principle (SRP)
// Each class should have only one reason to change,
// meaning it should have only one job or responsibility.


// Class for baking bread
class BreadBaker {
    bakeBread() {
        console.log("Baking high-quality bread...");
    }
}

// Class for managing inventory
class InventoryManager {
    manageInventory() {
        console.log("Managing inventory...");
    }
}

// Class for ordering supplies
class SupplyOrder {
    orderSupplies() {
        console.log("Ordering supplies...");
    }
}

// Class for serving customers
class CustomerService {
    serveCustomer() {
        console.log("Serving customers...");
    }
}

// Class for cleaning the bakery
class BakeryCleaner {
    cleanBakery() {
        console.log("Cleaning the bakery...");
    }
}

function main() {
    const baker = new BreadBaker();
    const inventoryManager = new InventoryManager();
    const supplyOrder = new SupplyOrder();
    const customerService = new CustomerService();
    const cleaner = new BakeryCleaner();

    // Each class focuses on its specific responsibility
    baker.bakeBread();
    inventoryManager.manageInventory();
    supplyOrder.orderSupplies();
    customerService.serveCustomer();
    cleaner.cleanBakery();
}

main();

/*
Explanation:

BreadBaker Class: 
Responsible solely for baking bread.
This class focuses on ensuring the quality and standards of the bread without being burdened by other tasks.

InventoryManager Class: 
Handles inventory management, ensuring that the bakery has the right ingredients and supplies available.

SupplyOrder Class: 
Manages ordering supplies, ensuring that the bakery is stocked with necessary items.

CustomerService Class: 
Takes care of serving customers, providing a focused approach to customer interactions.

BakeryCleaner Class:
Responsible for cleaning the bakery, ensuring a hygienic environment.
*/