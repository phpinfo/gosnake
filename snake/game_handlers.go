package snake

import "github.com/nsf/termbox-go"

func (game *Game) handleKeyEvents(event termbox.Event) {
	if event.Type != termbox.EventKey {
		return
	}

	switch event.Key {
	case termbox.KeyEsc:
		game.quit()
		break
	}

	switch event.Ch {
	case 'q':
		game.quit()
		break
	case 'p':
		game.pause()
		break
	}
}

func (game *Game) handleControlKeyEvents(event termbox.Event) {
	if game.isPause {
		return
	}

	if event.Type != termbox.EventKey {
		return
	}

	switch event.Key {
	case termbox.KeyArrowLeft:
		game.Snake.SetDirection(DirectionLeft)
		break
	case termbox.KeyArrowRight:
		game.Snake.SetDirection(DirectionRight)
		break
	case termbox.KeyArrowUp:
		game.Snake.SetDirection(DirectionUp)
		break
	case termbox.KeyArrowDown:
		game.Snake.SetDirection(DirectionDown)
		break
	}
}
