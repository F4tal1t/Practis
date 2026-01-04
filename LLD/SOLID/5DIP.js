// Dependency Inversion Principle (DIP)
// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.

// Interface for version control system simulated via convention
class IVersionControl {
    commit(message) { throw new Error("commit() must be implemented"); }
    push() { throw new Error("push() must be implemented"); }
    pull() { throw new Error("pull() must be implemented"); }
}

// Git version control implementation
class GitVersionControl extends IVersionControl {
    commit(message) {
        console.log("Committing changes to Git with message: " + message);
    }

    push() {
        console.log("Pushing changes to remote Git repository.");
    }

    pull() {
        console.log("Pulling changes from remote Git repository.");
    }
}

// Team class that relies on version control
class DevelopmentTeam {
    constructor(vc) {
        this.versionControl = vc;
    }

    makeCommit(message) {
        this.versionControl.commit(message);
    }

    performPush() {
        this.versionControl.push();
    }

    performPull() {
        this.versionControl.pull();
    }
}

function main() {
    const git = new GitVersionControl();
    const team = new DevelopmentTeam(git);

    team.makeCommit("Initial commit");
    team.performPush();
    team.performPull();
}

main();

/*
IVersionControl Interface: 
This defines the operations that any version control system should support, like commit, push, and pull. 
It serves as an abstraction that decouples high-level code from low-level implementations.

GitVersionControl Class: 
This class implements the IVersionControl interface,
providing specific functionality for managing version control using Git.

DevelopmentTeam Class: 
This class relies on the IVersionControl interface, meaning it can work with any version control implementation that adheres to the interface.
It does not need to know the details of how Git works internally.
*/