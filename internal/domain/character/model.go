package character

import (
	"time"
)

type Character struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Info básica
	Name             string `gorm:"not null" json:"name"`
	Race             string `gorm:"not null" json:"race"`  // slug de 5e-bits o custom
	Class            string `gorm:"not null" json:"class"` // slug de 5e-bits o custom
	Background       string `gorm:"not null" json:"background"`
	Alignment        string `json:"alignment"`
	ExperiencePoints int    `gorm:"default:0" json:"experience_points"`
	Level            int    `gorm:"not null;default:1" json:"level"`

	// Ability scores
	Strength     int `gorm:"not null;default:10" json:"strength"`
	Dexterity    int `gorm:"not null;default:10" json:"dexterity"`
	Constitution int `gorm:"not null;default:10" json:"constitution"`
	Intelligence int `gorm:"not null;default:10" json:"intelligence"`
	Wisdom       int `gorm:"not null;default:10" json:"wisdom"`
	Charisma     int `gorm:"not null;default:10" json:"charisma"`

	// Combat
	MaxHitPoints       int    `gorm:"not null;default:0" json:"max_hit_points"`
	CurrentHitPoints   int    `gorm:"not null;default:0" json:"current_hit_points"`
	TemporaryHitPoints int    `gorm:"default:0" json:"temporary_hit_points"`
	ArmorClass         int    `gorm:"not null;default:10" json:"armor_class"`
	Speed              int    `gorm:"not null;default:30" json:"speed"`
	HitDice            string `json:"hit_dice"`

	// Flags
	IsCustomRace  bool `gorm:"default:false" json:"is_custom_race"`
	IsCustomClass bool `gorm:"default:false" json:"is_custom_class"`

	// Backstory
	PersonalityTraits string `gorm:"type:text" json:"personality_traits"`
	Ideals            string `gorm:"type:text" json:"ideals"`
	Bonds             string `gorm:"type:text" json:"bonds"`
	Flaws             string `gorm:"type:text" json:"flaws"`
	Backstory         string `gorm:"type:text" json:"backstory"`

	// Skills (JSON con indices de skills con proficiencia)
	SkillProficiencies string `gorm:"type:text;default:'[]'" json:"skill_proficiencies"`
}
