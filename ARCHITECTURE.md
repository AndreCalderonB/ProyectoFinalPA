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
