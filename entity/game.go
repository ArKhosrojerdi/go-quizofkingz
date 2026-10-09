package entity

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	PlayerIDs   []uint
	WinnerId    uint
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
}
