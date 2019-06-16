# Parking Lot

This project is a implementation of an Automated Parking Lot.
Program accepts File input or can be accessed via Command Line Interface(CLI)

Language : GoLang

Version : go1.12.5


### Thoughts on GoLang
This is my first experience with Golang and have tried to learn as much I can and incorporate all knowledge I could gather over night for this project. 

Go is a such a naive language still packed with features that are so vital...
I can go whole day on this, to save the time lets keep this discussion for some other day.

Thanks to this article for explaining DI so beautifully:

https://blog.gojekengineering.com/the-many-flavours-of-dependency-injection-in-go-25aa070d79a0
  


## Architecture

Project is made up of following main components : 

## 1. Parking : 
 
   This is the main component. Parking Lot is responsible for managing all the parking related operations such as Create_Parking_Lot, Park Vehicle, Remove Vehicle etc.

     Properties:
      1. parking capacity
      2. List of Spots

     Functions for :  
      1. Create Parking
      2. Park Vehicle
      3. Remove Vehicle
      4. Check Parking Status
      5. Get all Cars by Color
      6. Get all Slots by Registration Number
      7. Get all Slots by Color

## 2. Slot : 

  Slot is the place in the Parking Lot where a Vehicle can be parked. Slot have a unique slot number. There can only be one Vehicle at a spot at a time.

    Properties:
      1. Slot Number
      2. Status
      3. Vehicle

    Functions: 
      1. Add Vehicle
      2. Remove Vehicle
    
## 3. Vehicle :

  Vehicle can be parked at the Spot. Vehicle have a Registration Number and Color

    Properties:
      1. Registration Number
      2. Color
   
## 4. CLI :

  Command Line Interface(CLI) is a way of interacting with the application.
  User can executes commands one at a time

  There are few components core to CLI

    1. ICommand - Interface to the Command Class. Every Command implements this interface
   
    2. CmdRegistery - All the Commands are defined here. It contains mapping of command with its handler

    3. CommandHandler - Contains a map of all the registered commands. Its an entry point for all commands

## 5. File Input:

  Gives a mechanism to read input from the file. Each command in file is intern goes through CLI interface.


  Here is a rough architecture diagram of all the components and how they interact.

![Architecture](/architecture.png)


## Build


Navigate to 'parking_lot' directory
``` 
parking_lot$ ./bin/setup
```

## Run

Navigate to 'parking_lot' directory

### Launch CLI
``` 
parking_lot$ ./bin/parking_lot
```

### File Input

``` 
parking_lot$ ./bin/parking_lot <file_name>
```

## Test


``` 
parking_lot$ ./bin/run_functional_tests
```


## Features : 

- [x] Configurable Commands, easy to create new commands
- [x] CLI Support
- [x] File Support
- [x] Help for Commands
- [x] Highly Modular Design
- [ ] Internationalization

