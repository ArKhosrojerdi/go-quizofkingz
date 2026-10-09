package entity

type Question struct {
	ID         uint
	Text       string
	Answer     string
	Category   Category
	Difficulty string
	AnswerList []string
}

type PossibleAnswer struct {
	ID      uint
	Content string
}

type Answer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}
