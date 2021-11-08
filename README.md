# Fibonacci Spiral

## Requirements

- Docker (tested in 20.10.6)

## Run locally

sh local-start.sh

Then enter to http://localhost:8080

## Run test

sh run-test.sh 

Notes: Only backend test was implemented. TODO: Implement frontend test.

## General Architecture

![image](https://user-images.githubusercontent.com/11558202/140818929-a456da83-7e52-4828-b8a9-69b9d7406cbc.png)



## Backend

Technologies: Golang.

The backend is using hexagonal architecture [More info](https://alistair.cockburn.us/hexagonal-architecture/).

#### Scaffolding
- cmd 
  - httpserver (entrypoints to the apps) 
- internal
  - core (contain all the business logic)
    - domain (contain all the business domains)
    - port (contain how the core module will interacts with the external actors(handlers and repositories)
    - service (implement the business logic)
  - handler (external actors to know how to interact with the core)
  - repository (external actors that implements the functions that core understand
- pkg (code to open to use by external modules)

#### Sequence Diagram
![image](https://user-images.githubusercontent.com/11558202/140818892-dae938d8-68a7-4084-946a-1efa6a42bb33.png)

## Frontend

Technologies: NodeJS & React

## TODO
- Add test to frontend
- Add observability
- Add authentication
- Improve frontend styles

