package models

type PokemonMoves struct {
	MoveName string `json:"move_name"`
	Effect   string `json:"effect"`
	Accuracy uint   `json:"accuracy"`
	PP       uint   `json:"pp"`
	Power    uint   `json:"power"`
}

type PokemonAbility struct {
	AbilityName string `json:"ability_name"`
	Effect      string `json:"effect"`
}

type PokemonStats struct {
	HP    uint `json:"hp"`
	Atk   uint `json:"atk"`
	Def   uint `json:"def"`
	SpAtk uint `json:"spatk"`
	SpDef uint `json:"spdef"`
	Speed uint `json:"speed"`
}

type PokemonMoveset struct {
	Move1 PokemonMoves `json:"move1"`
	Move2 PokemonMoves `json:"move2"`
	Move3 PokemonMoves `json:"move3"`
	Move4 PokemonMoves `json:"move4"`
}

type Pokemon struct {
	ID      uint           `gorm:"primary key;autoIncrement" json:"id"`
	Name    string         `json:"name"`
	Stats   PokemonStats   `json:"stats"`
	Ability PokemonAbility `json:"ability"`
	MoveSet PokemonMoveset `json:"moveset"`
	TeamID  uint           `gorm:"index" json:"team_id"`
	Team    Team           `gorm:"foreignKey:TeamID"`
}
