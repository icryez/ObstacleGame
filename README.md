# Multiplayer Obstacle Game

This repository contains a simple multiplayer obstacle game implemented in Go. Players can control a character to navigate through a dynamically generated map while avoiding obstacles and jumping over gaps.

(Under Development)

![Screencast from 2024-05-05 21-11-59](https://github.com/icryez/ObstacleGame/assets/35337801/517f2e45-d504-482e-9a66-a8fee649e81b)


## Installation

To install and run the game, follow these steps:

1. Clone this repository:

    ```bash
    git clone https://github.com/icryez/ObstacleGame.git
    ```

2. Navigate to the project directory:

    ```bash
    cd ObstacleGame
    ```

3. Build and run the game:

    ```bash
    go run .
    ```

## Features
- Multiplayer support: The game architecture is designed to support multiple players simultaneously. Game can connect to a server using tcp and reads/writes data. Server Repo - https://github.com/icryez/MultiGameServer
- Dynamic map generation: The game generates a new map layout every time it is run. --To Be Implemented
- Player control: Use the 'A' and 'D' keys to move left and right, respectively. Press 'space' to jump.
- Gravity simulation: The player character is subject to gravity, causing it to fall when not supported by solid ground.
- Obstacle avoidance: Navigate around obstacles and jump over gaps to progress through the map.

## Usage

After running the game, use the following controls:

- **A**: Move left.
- **D**: Move right.
- **Space**: Jump.

## Code Structure

The codebase is organized into several packages:

- **terminal**: Provides functions for interacting with the terminal, including clearing the screen and moving the cursor.
- **cursor**: Manages cursor visibility.
- **gametick**: Contains the game loop and functions related to game mechanics such as map generation, player movement, and gravity simulation.
- **player**: Handles player-related functionality, including player position and movement.
- **keyboard**: Deals with keyboard input handling.
- **mapmodule**: Responsible for generating the game map.
- **structs**: Defines various data structures used throughout the game.
- **connection**: Connects to a server using tcp and reads/writes data.
