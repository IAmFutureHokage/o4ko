package managers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"twentyoneo4ko/pkg/card"
	carddeck "twentyoneo4ko/pkg/card_deck"
)

type CardDeck interface {
	GeneratePokerWithoutJockers()
	Shuffle(shuffleHand carddeck.StackInterface[card.Card], shuffleIntensity int)
	DrawCard() card.Card
	ReturnCard(card.Card)
}

type ShuffleHand[T card.Card] interface {
	Push(card.Card)
	Pop() card.Card
	Length() int
}

type Player interface {
	AddCard(card.Card)
	Score() int
	HandSize() int
	DrawCardFromHand() card.Card
	ResetScore()
}

type Dealer interface {
	AddCard(card.Card)
	Score() int
	HandSize() int
	DrawCardFromHand() card.Card
	ResetScore()
	ShouldTakeCard() bool
	OpenClosedCard() card.Card
}

type BlackjackManager struct {
	CardDeck    CardDeck
	ShuffleHand ShuffleHand[card.Card]
	Player      Player
	Dealer      Dealer
}

func NewBlackjackManager(cardDeck CardDeck, shuffleHand ShuffleHand[card.Card], player Player, dealer Dealer) *BlackjackManager {
	return &BlackjackManager{
		CardDeck:    cardDeck,
		ShuffleHand: shuffleHand,
		Player:      player,
		Dealer:      dealer,
	}
}

func (bm *BlackjackManager) InitGame() {
	bm.CardDeck.GeneratePokerWithoutJockers()
}

func (bm *BlackjackManager) ShuffleCards(shuffleIntensity int) {
	fmt.Println("\nКарты перемешиваются")
	bm.CardDeck.Shuffle(bm.ShuffleHand, shuffleIntensity)
}

func (bm *BlackjackManager) DealCardsToDealer() {
	fmt.Println("\nДилер берет себе две карты")
	bm.Dealer.AddCard(bm.CardDeck.DrawCard()) // Первая карта закрыта
	card := bm.CardDeck.DrawCard()
	fmt.Println("\nДилер показывает карту:")
	fmt.Println(card.ToString())
	bm.Dealer.AddCard(card)
}

func (bm *BlackjackManager) DealCardsToPlayer() {
	fmt.Println("\nДилер раздает Вам две карты")
	fmt.Println("\nВы получили карты:")
	for i := 0; i < 2; i++ {
		card := bm.CardDeck.DrawCard()
		fmt.Println(card.ToString())
		bm.Player.AddCard(card)
	}
}

func (bm *BlackjackManager) PlayerHits() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nВыберите действие (hit или stand):")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		switch action {
		case "hit":
			card := bm.CardDeck.DrawCard()
			fmt.Println("\nВы получили карту:")
			fmt.Println(card.ToString())
			bm.Player.AddCard(card)
			if bm.Player.Score() > 21 {
				fmt.Println("\nПеребор!")
				return
			}
		case "stand":
			return
		default:
			fmt.Println("\nНекорректное действие. Пожалуйста, выберите hit или stand.")
		}
	}
}

func (bm *BlackjackManager) DealerHits() {
	fmt.Println("\nДилер открыл карту:")
	fmt.Println(bm.Dealer.OpenClosedCard().ToString())

	for bm.Dealer.ShouldTakeCard() {
		card := bm.CardDeck.DrawCard()
		fmt.Println("\nДилер берет карту:")
		fmt.Println(card.ToString())
		bm.Dealer.AddCard(card)
	}
}

func (bm *BlackjackManager) ShowResults() {
	playerScore := bm.Player.Score()
	dealerScore := bm.Dealer.Score()

	fmt.Printf("\nСчет игрока: %d\n", playerScore)
	fmt.Printf("Счет дилера: %d\n", dealerScore)

	switch {
	case playerScore > 21:
		fmt.Println("Игрок проиграл, так как его счет больше 21.")
	case dealerScore > 21:
		fmt.Println("Игрок выиграл, так как счет дилера больше 21.")
	case playerScore > dealerScore:
		fmt.Println("Игрок выиграл, так как его счет больше счета дилера.")
	case playerScore < dealerScore:
		fmt.Println("Игрок проиграл, так как его счет меньше счета дилера.")
	default:
		fmt.Println("Ничья, так как счет игрока равен счету дилера.")
	}
}

func (bm *BlackjackManager) CollectCards() {
	fmt.Println("\nИгрок и дилер сдают карты:")

	fmt.Println("Игрок сдал карты")
	for bm.Player.HandSize() > 0 {
		card := bm.Player.DrawCardFromHand()
		bm.CardDeck.ReturnCard(card)
	}
	bm.Player.ResetScore()

	fmt.Println("Дилер сдал карты")
	for bm.Dealer.HandSize() > 0 {
		card := bm.Dealer.DrawCardFromHand()
		bm.CardDeck.ReturnCard(card)
	}
	bm.Dealer.ResetScore()
}

func (bm *BlackjackManager) AskToPlayAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nХотите сыграть еще раз? (Введите 'да' для продолжения или 'нет' для выхода):")
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Произошла ошибка при чтении ответа. Попробуйте снова.")
		return bm.AskToPlayAgain()
	}
	answer = strings.TrimSpace(answer)
	if strings.ToLower(answer) == "да" {
		return true
	} else {
		fmt.Println("Спасибо за игру! До свидания!")
		return false
	}
}
