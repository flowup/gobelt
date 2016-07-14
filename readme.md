## Gobelt

> Warning: This project is still in development

Gobelt is a code generator library written in Golang. It is based on models and can currently generate basic data access objects (DAOs) with tests for said models.
Note: Currently you need to have branch "development" set up in [**Gogen libary**](https://github.com/flowup/gogen)


## Usage
This library is used by running the main.go located in /cmd/gobelt and selecting one of available choices (currently dao or suite) and passing 

```cmd

$go run main.go dao $HOME/model.go

// OR

$go run main.go suite $HOME/model.go