# Design Patterns in Go

This project demonstrates four fundamental design patterns implemented in Go:

## 1. Singleton Pattern (`singleton.go`)

**Purpose**: Ensures a class has only one instance and provides global access to it.

**Implementation**:
- Uses `sync.Once` to ensure thread-safe initialization
- Database connection example that creates only one instance
- Demonstrates memory address comparison to prove single instance

**Key Features**:
- Thread-safe singleton creation
- Lazy initialization
- Global access point

**When to Use**:
- Database connections
- Configuration objects
- Logging services
- Cache instances

## 2. Factory Pattern (`factory.go`)

**Purpose**: Creates objects without specifying their concrete classes.

**Implementation**:
- `Notification` interface with multiple implementations
- Factory function `CreateNotification()` that returns different types
- Support for Email, SMS, Push, and Slack notifications

**Key Features**:
- Abstraction of object creation
- Easy to extend with new notification types
- Centralized creation logic

**When to Use**:
- When you need to create objects based on runtime conditions
- To hide complex object creation logic
- When you want to decouple object creation from usage

## 3. Observer Pattern (`observer.go`)

**Purpose**: Defines a one-to-many dependency between objects so that when one object changes state, all dependents are notified.

**Implementation**:
- `Observer` interface for subscribers
- `Subject` interface for publishers
- News agency example with multiple subscriber types
- Stock market example for real-time updates

**Key Features**:
- Loose coupling between subject and observers
- Dynamic subscription/unsubscription
- Broadcast notifications to multiple observers

**When to Use**:
- Event handling systems
- Model-View architectures
- Real-time notifications
- State change propagation

## 4. Decorator Pattern (`decorator.go`)

**Purpose**: Adds new functionality to objects dynamically without altering their structure.

**Implementation**:
- `Coffee` interface as base component
- Multiple decorators: Milk, Sugar, Whip, Vanilla, Chocolate
- Builder pattern integration for easy composition
- Coffee shop example

**Key Features**:
- Dynamic behavior addition
- Flexible composition
- Maintains interface compatibility
- Stack-able decorators

**When to Use**:
- Adding features to objects at runtime
- Avoiding class explosion
- Flexible configuration systems
- Middleware implementations

## Running the Code

```bash
# Initialize Go module
go mod init backarch

# Run the demonstration
go run .
```

## Project Structure

```
BackArch/
├── main.go         # Main demonstration file
├── singleton.go    # Singleton pattern implementation
├── factory.go      # Factory pattern implementation
├── observer.go     # Observer pattern implementation
├── decorator.go    # Decorator pattern implementation
├── go.mod          # Go module file
└── README.md       # This file
```

## Pattern Benefits

### Singleton
- **Memory efficiency**: Only one instance exists
- **Global access**: Available throughout the application
- **Thread safety**: Uses sync.Once for safe initialization

### Factory
- **Flexibility**: Easy to add new product types
- **Decoupling**: Client code doesn't know concrete classes
- **Centralization**: Creation logic in one place

### Observer
- **Loose coupling**: Subjects and observers are independent
- **Dynamic relationships**: Add/remove observers at runtime
- **Open/Closed principle**: Easy to add new observer types

### Decorator
- **Flexibility**: Combine behaviors dynamically
- **Single responsibility**: Each decorator has one purpose
- **Open/Closed principle**: Add new decorators without modifying existing code

## Example Output

The program demonstrates each pattern with realistic examples:

1. **Singleton**: Shows same database instance being reused
2. **Factory**: Creates different notification types
3. **Observer**: News agency broadcasting to subscribers
4. **Decorator**: Building complex coffee orders step by step

Each pattern includes emojis and clear output to help visualize the concepts!
