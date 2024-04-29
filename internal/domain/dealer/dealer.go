package dealer

import (
	"fmt"
	"twentyoneo4ko/internal/domain/player"
	"twentyoneo4ko/pkg/card"
)

type StackInterface[T any] interface {
	Length() int
	Push(value T)
	Pop() T
}

type Dealer struct {
	player.Player
	closedCard *card.Card
}

func NewDealer(deck StackInterface[card.Card]) *Dealer {
	return &Dealer{
		Player: *player.NewPlayer(deck),
	}
}

func (d *Dealer) AddCard(c card.Card) {
	if d.closedCard == nil {
		d.closedCard = &c
	}
	d.Player.AddCard(c)
}

func (d *Dealer) ResetScore() {
	d.Player.ResetScore()
	d.closedCard = nil
}

func (d *Dealer) OpenClosedCard() card.Card {
	if d.closedCard == nil {
		panic("closed card not found")
	}
	card := *d.closedCard
	d.closedCard = nil
	return card
}

func (d *Dealer) ShouldTakeCard() bool {
	fmt.Println(d.Score())
	response := d.Score() < 17
	return response
}
