package internal

func NewGame(secret string) *Game {
	return &Game{secret: secret}
}

type Game struct {
	secret string
}

type GuessResult struct {
	Solved bool
}

func (g *Game) Guess(guess string) GuessResult  {
	return GuessResult{
		Solved: guess == g.secret,
	}
}
