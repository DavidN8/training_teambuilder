package models

type Team struct {
	ID       uint      `gorm:"primary key;autoIncrement" json:"id"`
	Pokemons []Pokemon `json:"pokemons"`
	UserID   uint      `gorm:"index" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID"`
}
