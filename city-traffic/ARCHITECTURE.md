# City Traffic

Traffic simulator in go.

## Architecture
**Ebiten library**

According to their [documentation](https://github.com/hajimehoshi/ebiten/): Ebiten is an open source game library for the Go programming language. Ebiten's simple API allows you to quickly and easily develop 2D games that can be deployed across multiple platforms.

**Ebiten Game Design**

Ebiten library's `ebiten.Game` interface makes game development in go very easy, there are three necessary methods that require implementation:
* **Update**: This method is where you put the game logic, which will be updated each `Tick` (1/60th of a second).
* **Draw**: Method to render the images in every frame. 
* **Layout**: Method that defines the overall game layout.


### Project Structure
The structure is made of two folders and the main.go file: 

* city-traffic 
    * imgs/ - images used for the graphic interfaces
    * scripts/ - scripts used for the simulator
    * main.go


The `scripts/` is where the game logic resides. The main entities needed for the game are:
* Game [game.go](https://github.com/AndreCalderonB/ProyectoFinalPA/blob/master/city-traffic/scripts/game.go)
* Semaphore [semaphore.go](https://github.com/AndreCalderonB/ProyectoFinalPA/blob/master/city-traffic/scripts/semaphore.go)
* Cars [cherry.go](https://github.com/AndreCalderonB/ProyectoFinalPA/blob/master/city-traffic/scripts/hud.go)
* A Hud [hud.go](https://github.com/AndreCalderonB/ProyectoFinalPA/blob/master/city-traffic/scripts/hud.go)

`game.go` will have all `ebiten.Game` interface methods implemented, in addition to the following functions and methods 
* **NewGame**: function that instanciates a new game
* Logic behind cherry eats

As we know, we need to implement the `Game` interface. This is done in `main.go` for simplicity reasons, but inside `game.go` we will code the functionality of the structure.

`semaphore.go` is the entitie that controls traffic in each lane of cars. A semaphore can do the following:
* **makeCar**. This function initializes cars at its lane and then spawns cars after the others are gone.
* **toggleLight**. A semaphore should be able to toggle the color of its own light. Each semaphore has a timer to change color asynchronously with the ohter semaphores
* **queue**. After a car is initialized on the lane it is added to a queue of waiting cars at the semaphore.
* **dequeue**. When a car position that belongs at the lane has passed the semaphore, the car is dequeued form the waiting line.

In the image below you can see all entities and methods in detail

![uml](uml.png)


### Concurrency

Semaphores are being created concurrently, the game entitie has a reference to each one so they can toggle lights on a synchronized way. Then cars are spawned also concurrently on each lane.
