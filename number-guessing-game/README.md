# Number Guessing Game

An interactive console-based number guessing game written in Go. Players must guess a random number between 1 and 100 with different difficulty levels and limited attempts.

## 🎮 Game Features

- 🎯 **Random number generation** between 1-100
- 🏆 **Three difficulty levels**:
  - Easy: 10 attempts
  - Medium: 5 attempts
  - Hard: 3 attempts
- 💡 **Smart hints** (higher/lower) after each guess
- 🔄 **Replay functionality** - play multiple rounds
- 📊 **Attempt tracking** with clear feedback
- 🎉 **Victory celebration** when you win

## Installation

1. Clone or download this repository
2. Navigate to the number-guessing-game directory
3. Run the game with Go:

```bash
go run main.go
```

## How to Play

### Starting the Game

```bash
go run main.go
```

### Game Flow

1. **Select Difficulty**: Choose from Easy (10 chances), Medium (5 chances), or Hard (3 chances)
2. **Make Guesses**: Enter numbers between 1 and 100
3. **Get Hints**: After each guess, the game tells you if the number is higher or lower
4. **Win or Lose**: Guess correctly to win, or run out of attempts to lose
5. **Play Again**: Choose to start a new game

### Example Game Session

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Attempt 1 - Enter your guess: 50
Incorrect! The number is greater than 50.

Attempt 2 - Enter your guess: 75
Incorrect! The number is less than 75.

Attempt 3 - Enter your guess: 62
🎉 Congratulations! You guessed the correct number in 3 attempts.

Would you like to play again? (y/n): n
```

## Difficulty Levels

| Level  | Attempts | Description             |
| ------ | -------- | ----------------------- |
| Easy   | 10       | Best for beginners      |
| Medium | 5        | Balanced challenge      |
| Hard   | 3        | For experienced players |

## Game Logic

- **Random Number**: Generated between 1 and 100 (inclusive)
- **Hints**: After each guess, tells you if the target is higher or lower
- **Attempt Tracking**: Shows current attempt number
- **Win Condition**: Correct guess within allowed attempts
- **Lose Condition**: Run out of attempts without guessing correctly

## Requirements

- Go 1.16 or higher
- No external dependencies

## File Structure

```
number-guessing-game/
├── main.go     # Main game logic
└── README.md   # This file
```

## Learning Objectives

This project demonstrates:

- **User input handling** with `fmt.Scan()`
- **Random number generation** with `math/rand`
- **Control structures** (loops, conditionals, switch statements)
- **Game state management**
- **Interactive console applications**

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.
