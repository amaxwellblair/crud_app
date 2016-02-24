#### RESTful app using Go stdlib and BoltDB

##### Objectives
- Explore net/http package
- Explore template package
- Quickly implement a database

##### Summary
Created a basic RESTful app using minimal dependancies. Chose [BoltDB](https://github.com/boltdb/bolt)
 as a database given its simple user interface and near zero setup time. BoltDB
 is a simple key/value store and is used in live productions today. The database
 manager is located in `store.go`

The application supports index, new, create, show, edit, update and delete actions.
These actions are taken care of in `main.go` in  the robots handler. Uses Golang
 templates for basic iteration and argument passing.

##### TODO
- [ ] Basic styling
- [ ] Abstract handlers from `main.go`
- [ ] Refactor templates and create partials
- [ ] Additional `store.go` testing
- [ ] Create test suite for handler files
