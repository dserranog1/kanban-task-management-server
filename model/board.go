package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardModel struct {
	db *gorm.DB
}

// create an gorm model for the board
type Board struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
}

// constructor
func NewBoardModel(db *gorm.DB) *BoardModel {
	return &BoardModel{db}
}

// get all boards from the database and return them
func (boardModel BoardModel) GetAllBoards() []Board {
	var boards []Board
	boardModel.db.Find(&boards)
	return boards
}
