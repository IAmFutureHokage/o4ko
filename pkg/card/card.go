package card

type Suit int

const (
	Clubs    Suit = iota // трефы
	Diamonds             // бубны
	Hearts               // червы
	Spades               // пики
)

type Card struct {
	Suit  Suit
	Value byte
}

func NewCard(suit Suit, value byte) Card {
	return Card{Suit: suit, Value: value}
}

func (c Card) ToString() string {
	return cardValueToString(c.Value) + " of " + suitToString(c.Suit)
}

func cardValueToString(value byte) string {
	switch value {
	case 'T':
		return "10"
	case 'J':
		return "Jack"
	case 'Q':
		return "Queen"
	case 'K':
		return "King"
	case 'A':
		return "Ace"
	default:
		return string(value)
	}
}

func suitToString(suit Suit) string {
	switch suit {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return "Unknown"
	}
}
