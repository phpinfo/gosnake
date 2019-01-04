package snake

import (
	"time"
	"github.com/nsf/termbox-go"
	"math/rand"
)

const (
	BoxRectX      = 1
	BoxRectY      = 2
	BoxRectWidth  = 50
	BoxRectHeight = 25

	SnakeLength = 4

	TitleText = "Go Snake!"
	TitleX    = 1
	TitleY    = 0

	ScoreValue = 0
	ScoreX     = 48
	ScoreY     = 0
)

type Game struct {
	score    int
	Renderer *Renderer
	lblTitle *Label
	lblScore *Counter
	Snake    *Snake
	Box      *Rect
	Food     *Food
	isQuit   bool
	isPause  bool
	aPause   *Alert
}

func NewGame() (*Game) {
	var (
		rect  = NewRect(BoxRectX, BoxRectY, BoxRectWidth, BoxRectHeight)
		snake = initSnake(rect, SnakeLength)
		food  = initFood(rect, snake)
	)

	game := &Game{
		score:    ScoreValue,
		Renderer: NewRenderer(),
		lblTitle: NewLabel(TitleText, NewPoint(TitleX, TitleY)),
		Snake:    snake,
		Box:      rect,
		Food:     food,
		isQuit:   false,
		isPause:  false,
		aPause:   initPauseAlert(rect),
	}

	lblScore := NewCounter(&game.score, NewPoint(ScoreX, ScoreY))
	game.lblScore = lblScore

	return game
}

func (game *Game) Start() {
	game.Renderer.Init()
	defer game.Renderer.Close()

	rand.Seed(time.Now().UnixNano())

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	loop:
	for {
		select {
		case event := <-eventQueue:
			game.handleKeyEvents(event)
			game.handleControlKeyEvents(event)

		default:
			if game.isQuit {
				break loop
			}
			game.tick()
		}
	}
}

func (game *Game) tick() {
	if game.isPause {
		return
	}

	game.Snake.Move()

	if game.isCollision() {
		game.quit()
		return
	}

	if game.isFood() {
		game.Snake.Eat()
		game.score++
		game.Food = initFood(game.Box, game.Snake)
	}

	game.Render()
	time.Sleep(game.getMoveInterval())
}

func (game *Game) isCollision() bool {
	if !game.Box.Contains(game.Snake.Head()) {
		return true
	}

	if game.Snake.SelfCollides() {
		return true
	}

	return false
}

func (game *Game) isFood() bool {
	return game.Snake.Head().Equals(game.Food.Point)
}

func (game *Game) getMoveInterval() time.Duration {
	ms := 100 - game.score / 1
	return time.Duration(ms) * time.Millisecond
}

func (game *Game) quit() {
	game.isQuit = true
}

func (game *Game) pause() {
	game.isPause = !game.isPause
}

func initSnake(rect *Rect, length int) *Snake {
	var (
		point = rect.LeftTopPoint.Add(rect.Dimensions.Width / 2, rect.Dimensions.Height - 2)
		body = []*Point{point}
	)

	for dy := -1; dy > -length; dy-- {
		body = append(body, point.Add(0, dy))
	}

	return NewSnake(body, DirectionUp)
}

func initFood(rect *Rect, snake *Snake) *Food {
	for {
		var (
			x = rand.Intn(rect.Dimensions.Width)
			y = rand.Intn(rect.Dimensions.Height)
			point = NewPoint(x, y).Add(rect.Left, rect.Top)
		)

		if !snake.Contains(point) {
			return NewFood(point)
		}
	}
}

func initPauseAlert(rect *Rect) *Alert {
	text := "          PAUSED\n\n" +
		    "<Press any key to continue>"

	alert := NewAlert(text, NewPoint(0, 0))

	point := NewPoint(
		rect.Dimensions.Width / 2 - alert.getWidth() / 2 + rect.Left,
		rect.Dimensions.Height / 2 - alert.getHeight() / 2 + rect.Top,
	)

	alert.Move(point)

	return alert
}
