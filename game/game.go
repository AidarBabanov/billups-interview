package game

import (
	"math/rand"
	"time"
)

var (
	choices      = []string{"rock", "paper", "scissors", "lizard", "spock"}
	descriptions = map[int]string{
		bmask(1) | bmask(1): "Two solid Dwaynes.",     // two rocks played.
		bmask(1) | bmask(2): "Paper covers rock.",     // rock and paper played.
		bmask(1) | bmask(3): "Rock crushes scissors.", // ...
		bmask(1) | bmask(4): "Rock crushes lizard.",
		bmask(1) | bmask(5): "Spock vaporizes rock.",
		bmask(2) | bmask(2): "What did one piece of toilet paper say to the other? 'I feel really wiped.'",
		bmask(2) | bmask(3): "Scissors cut paper.",
		bmask(2) | bmask(4): "Lizard eats paper.",
		bmask(2) | bmask(5): "Paper disproves Spock.",
		bmask(3) | bmask(3): "Scissors look like a good position.",
		bmask(3) | bmask(4): "Scissors decapitates lizard.",
		bmask(3) | bmask(5): "Spock smashes scissors.",
		bmask(4) | bmask(4): "So, you want to start a lizard colony. Right?",
		bmask(4) | bmask(5): "Lizard poisons Spock.",
		bmask(5) | bmask(5): "Have you ever seen two Spocks?",
	}
	wins = map[int]struct{}{
		13: {}, // 1 wins 3
		14: {}, // 1 wins 4
		21: {}, // 2 wins 1
		25: {}, // ...
		32: {},
		34: {},
		42: {},
		45: {},
		53: {},
		51: {},
	}
)

// bmask converts id of choice to bitmask
func bmask(id int) int {
	return 1 << (id - 1)
}

type GService struct {
}

func NewGame() Game {
	rand.Seed(time.Now().UnixNano())
	return &GService{}
}

func (s *GService) GetChoices() []string {
	return choices
}

func (s *GService) GetChoice() (int, string) {
	id := rand.Intn(5) + 1
	return id, choices[id-1]
}

// play simulates game for the first player. If he wins return result and description of the result.
func play(pl1, pl2 int) (pl1choice, pl2choice int, res string, description string) {
	if pl1 == pl2 {
		res = "tie"
	} else if _, win := wins[pl1*10+pl2]; win {
		res = "win"
	} else {
		res = "lose"
	}
	return pl1, pl2, res, descriptions[bmask(pl1)|bmask(pl2)]
}

// Play simulates game for player with computer. If he wins return result and description of the result.
func (s *GService) Play(choiceID int) (plChoiceID, comChoiceID int, res string, description string) {
	comChoiceID = rand.Intn(5) + 1
	return play(choiceID, comChoiceID)
}
