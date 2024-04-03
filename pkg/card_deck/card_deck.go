package carddeck

import (
	"math/rand"
	"twentyoneo4ko/pkg/card"
)

type StackInterface[T any] interface {
	Length() int
	Push(value T)
	Pop() T
}

type Deck struct {
	cards StackInterface[card.Card]
}

func NewDeck(cards StackInterface[card.Card]) *Deck {
	return &Deck{
		cards: cards,
	}
}

func (d *Deck) GeneratePokerWithoutJockers() {
	if d.cards.Length() != 0 {
		panic("deck already generate")
	}
	for suit := card.Clubs; suit <= card.Spades; suit++ {
		for _, value := range []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'} {
			d.cards.Push(card.NewCard(suit, value))
		}
	}
}

func (d *Deck) Shuffle(hand StackInterface[card.Card], shuffleIntensity int) {
	for i := 0; i < shuffleIntensity; i++ {
		// Перемещаем случайное количество карт из колоды в Hand
		for j := 0; j < rand.Intn(d.cards.Length())+1; j++ {
			if d.cards.Length() > 0 {
				hand.Push(d.cards.Pop())
			}
		}

		// Запоминаем случайную карту, оставляя ее вне Hand и колоды
		var memorizedCard *card.Card
		if hand.Length() > 0 {
			tempCard := hand.Pop()
			memorizedCard = &tempCard
		}

		// Возвращаем карты из Hand обратно в колоду
		for hand.Length() > 0 {
			d.cards.Push(hand.Pop())
		}

		// Кладем "запомненную" карту обратно в колоду
		if memorizedCard != nil {
			d.cards.Push(*memorizedCard)
		}
	}
}

func (d *Deck) DrawCard() card.Card {
	if d.cards.Length() == 0 {
		panic("cannot draw from an empty deck")
	}
	return d.cards.Pop()
}

func (d *Deck) ReturnCard(card card.Card) {
	d.cards.Push(card)
}
