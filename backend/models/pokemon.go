package models

type Pokemon struct {
	ID          uint   `gorm:"primary key;autoIncrement" json:"id"`
	Name        string `json:"name"`
	HP          uint   `json:"hp"`
	Atk         uint   `json:"atk"`
	Def         uint   `json:"def"`
	SpAtk       uint   `json:"spatk"`
	SpDef       uint   `json:"spdef"`
	Speed       uint   `json:"speed"`
	AbilityName string `json:"ability_name"`
	Move1ID     uint   `json:"move1_id"`
	Move2ID     uint   `json:"move2_id"`
	Move3ID     uint   `json:"move3_id"`
	Move4ID     uint   `json:"move4_id"`
	Team        Team   `gorm:"foreignKey:TeamID"`
	TeamID      uint   `gorm:"index" json:"team_id"`
}
