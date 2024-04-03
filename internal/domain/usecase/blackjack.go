package usecase

type BlackjackManager interface {
	InitGame()
	ShuffleCards(shuffleIntensity int)
	DealCardsToDealer()
	DealCardsToPlayer()
	PlayerHits()
	DealerHits()
	ShowResults()
	CollectCards()
	AskToPlayAgain() bool
}

type BlackjackGame struct {
	bjmanager BlackjackManager
}

func NewBlackjackGame(manager BlackjackManager) *BlackjackGame {
	return &BlackjackGame{bjmanager: manager}
}

func (gm *BlackjackGame) Run() {
	gm.bjmanager.InitGame()

	for {
		gm.bjmanager.ShuffleCards(100000)
		gm.bjmanager.DealCardsToDealer()
		gm.bjmanager.DealCardsToPlayer()
		gm.bjmanager.PlayerHits()
		gm.bjmanager.DealerHits()
		gm.bjmanager.ShowResults()
		gm.bjmanager.CollectCards()

		if !gm.bjmanager.AskToPlayAgain() {
			return
		}
	}
}
