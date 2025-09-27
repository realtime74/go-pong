# first steps

- [x] create a project repository
- [x] set up version control (e.g., Git)

- [x] write a README.md

# hello world

- [x] write a simple "Hello, World!"

# use tcell library

- [x] install tcell library
- [x] create a basic terminal application using tcell
- [x] write a hello world as TUI
- [x] move "hello world" to the middle of the screen

# title bar

- [x] create a title bar at the top of the screen
    - [x] let titlebar consume the full width of the screen
    - [x] choose a smarter color
- [x] display "go-pong" in the title bar

# racket

- [x] draw the player's racket
      on the right side of the screen
- [x] draw the other player's racket
      on the left side of the screen

# interaction

- [x] move the player's racket up and down
      using the j/k keys

# ball

- [x] draw the ball in the middle of the screen
    - [x] create a ball struct
    - [x] create a function to draw the ball
- [x] let the ball move to the right
    - [x] create a goroutine to move the ball
    - [x] update ball position in the loop
- [x] bounce off the walls
    - [x] for the moment,
          let the ball just bounce of the walls

- [x] create game controller
    - [x] move ball bounce logic to
          game controller 
- [x] bounce off the racket
    - [x] detect collision with the racket
    - [x] change ball direction on collision

- [x] detect scoring
    - [x] detect when the ball goes past a racket
    - [x] flash a red bottom line on the screen

- [x] refactor
    - [x] move game logic to a separate module

- [x] measure score
    - [x] create a footer bar
    - [x] keep track of the score for both players
    - [x] display the score in the footer bar

# ball vector

- [x] game makes much more fun if ball can fly
      in different angles
    - [x] create a vector struct
    - [x] use vector for ball movement
- [x] fix bounce in wrong direction

# add computer player

- [x] moves left cursor autonomically
    -   only allowed to watch and move
        every 10 ticks
    - sees only current ball position

------



