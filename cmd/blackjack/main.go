package main

import (
	"twentyoneo4ko/internal/domain/dealer"
	"twentyoneo4ko/internal/domain/player"
	"twentyoneo4ko/internal/domain/usecase"
	"twentyoneo4ko/internal/managers"
	"twentyoneo4ko/pkg/card"
	deck "twentyoneo4ko/pkg/card_deck"
	ls "twentyoneo4ko/pkg/linked_list"
	"twentyoneo4ko/pkg/stack"
)

func main() {

	linkedList := ls.New[card.Card]()
	cardStack := stack.NewStack(linkedList)

	playerHand := stack.NewStack(ls.New[card.Card]())
	dealerHand := stack.NewStack(ls.New[card.Card]())

	hand := stack.NewStack(ls.New[card.Card]())
	cardDeck := deck.NewDeck(cardStack)
	player := player.NewPlayer(playerHand)
	dealer := dealer.NewDealer(dealerHand)

	bjManager := managers.NewBlackjackManager(cardDeck, hand, player, dealer)

	game := usecase.NewBlackjackGame(bjManager)

	game.Run()
}
