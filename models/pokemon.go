package models

type Pokemon struct {
	ID          uint   `gorm:"primary key;autoIncrement" json:"id"`
	TeamId 		int `json:"team_id"` <-- forein key
	moves []moves 
	Name        string `json:"name"`
	stat []stat 
	ev []stat
	AbilityName string `json:"ability_name"`
	Team        Team   `gorm:"foreignKey:TeamID"`
	TeamID      uint   `gorm:"index" json:"team_id"`
}

type Moves struct{
	pp 
	power
	acc
	pokemonId
}

type stat struct{
	HP          uint   `json:"hp"`
	Atk         uint   `json:"atk"`
	Def         uint   `json:"def"`
	SpAtk       uint   `json:"spatk"`
	SpDef       uint   `json:"spdef"`
	Speed       uint   `json:"speed"`
	pokemonId
}
