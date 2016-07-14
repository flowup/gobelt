## Gobelt

> Warning: This project is still in development

Gobelt is a code generator library written in Golang. It is based on models and can currently generate basic data access objects (DAOs) with tests for said models.


## Development guideline

Note: Currently you need to have branch "development" set up in [**Gogen libary**](https://github.com/flowup/gogen) in order for Gobelt to work.


## Usage
This library is used by installing it from source, running it with gobelt command and selecting one of available choices (currently dao or suite) and passing paths to models as arguments.

```cmd

$go get github.com/flowup/gobelt
$go build

$gobelt dao $HOME/model.go

// OR

$gobelt suite $HOME/model.go

```