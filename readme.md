## Gobelt

> Warning: This project is still in development

Gobelt is a code generation tool written in Golang. It produces code based
on existing Golang code. It is mainly focused on generating boilerplate and
repetitive code such as test suites, data access objects or otherwise templated
code like observables.

## Contributing to Gobelt

`Develop` branch of Gobelt has a dependency of [**Gogen**](https://github.com/flowup/gogen)
also on branch `develop`. When issuing pull requests, please make sure you are
building on top of the `develop` branch.

## Installation

```cmd
$ go get github.com/flowup/gobelt/cmd/gobelt

$ gobelt # for usage information
```

## Features

- [x] Testify test suites
- [x] Observables
- [x] Data Access Objects (based on Gorm)
- [ ] Map, Filter and Reduce operations
- [ ] Data Access Objects (pure SQL)
- [ ] Function RPC (for microservices)
- [ ] More to come soon

### DAO generator

This part of gobelt generates data access objects based on models. Please note that generated methods ending with `T` (`ReadT` for example) are methods that return the db transaction instead of the models. These methods can be used to minimize number of redundant requests to DB.
