package player

import (
	"twentyoneo4ko/pkg/card"
)

type StackInterface[T any] interface {
	Length() int
	Push(value T)
	Pop() T
}

type Player struct {
	hand  StackInterface[card.Card]
	score int
	aces  int
}

func NewPlayer(deck StackInterface[card.Card]) *Player {
	return &Player{
		hand:  deck,
		score: 0,
	}
}

func (p *Player) AddCard(card card.Card) {
	p.hand.Push(card)

	switch card.Value {
	case 'A':
		// Добавляем 11 очков за туз, если это не приводит к перебору, иначе добавляем 1
		if p.score+11 > 21 {
			p.score += 1
		} else {
			p.score += 11
			p.aces += 1
		}
	case 'K', 'Q', 'J', 'T':
		p.score += 10
	default:
		p.score += int(card.Value - '0')
	}

	// Корректировка счета, если сумма очков с тузами, считающимися за 11, превышает 21
	for p.score > 21 && p.aces > 0 {
		p.score -= 10
		p.aces--
	}
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) HandSize() int {
	return p.hand.Length()
}

func (p *Player) DrawCardFromHand() card.Card {
	if p.hand.Length() == 0 {
		panic("no cards left in hand to draw")
	}
	return p.hand.Pop()
}

func (p *Player) ResetScore() {
	p.score = 0
}
