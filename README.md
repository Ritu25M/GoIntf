# Go Project: Avoiding Cyclic Imports

This project demonstrates how to avoid cyclic imports in a Go project by using interfaces. Cyclic imports occur when two or more packages directly or indirectly import each other, causing compilation errors and making code maintenance challenging. By using interfaces, we can decouple packages and break cyclic dependencies.

## Project Structure

The project consists of two packages:

1. **Species**: Contains functionalities related to species.
2. **Tag**: Handles tag generation for different species .

## How It Works
In the package `tag`, we create a struct that holds an instance of the `species` interface to handle `tag` related methods.
