package database

// These are the models for the C2 Team Server database.
// They should be implemented using GORM with PostgreSQL.

import (
	"time"

	"gorm.io/gorm"
)

// Models

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Team model, represents a "group" of users with access to a set of resources, campaigns, etc.
type Team struct {
	BaseModel
	Name        string       `gorm:"not null" json:"name"`
	Description string       `json:"description"`
	Members     []TeamMember `gorm:"foreignKey:TeamID" json:"-"`
	Campaigns   []Campaign   `gorm:"foreignKey:TeamID" json:"-"`
}

// TeamMember model, represents a user in a team.
type TeamMember struct {
	BaseModel
	TeamID uint `gorm:"not null" json:"team_id"`
	UserID uint `gorm:"not null" json:"user_id"`
	Team   Team `gorm:"foreignKey:TeamID" json:"-"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
}

// User model, represents a user (regardless of role) in the C2 Team Server.
type User struct {
	BaseModel
	Name         string       `gorm:"not null" json:"name"`
	LoginString  string       `gorm:"unique;not null" json:"login_string"` // This is the "username" field for the user at login.
	PasswordHash string       `gorm:"not null" json:"password_hash"`
	TeamMembers  []TeamMember `gorm:"foreignKey:UserID" json:"-"`
}

// Campaign model, represents a campaign in the C2 Team Server.
// A campaign is a collection of targets, tasks, and other resources.
// It is used to organize and manage a set of targets and tasks for a specific purpose.
// For example, a campaign could be used to target a specific set of IP addresses, URLs, or other resources.
// A campaign is owned by a team, and can be accessed by all members of the team.
// A campaign can have multiple tasks, which are the individual tasks that will be executed on the targets.
// A task is a collection of commands that will be executed on a target.
// A command is a single command that will be executed on a target.
type Campaign struct {
	BaseModel
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	TeamID      uint           `gorm:"not null" json:"team_id"`
	Status      string         `gorm:"not null" json:"status"`
	Target      string         `json:"target"`
	TargetType  string         `json:"target_type"`
	Team        Team           `gorm:"foreignKey:TeamID" json:"-"`
	Goals       []CampaignGoal `gorm:"foreignKey:CampaignID" json:"-"`
	Tasks       []Task         `gorm:"foreignKey:CampaignID" json:"-"`
}

// Campaign Goal model, represents a goal for a campaign.
// A goal is a specific target that the campaign is trying to achieve.
// For example, a goal could be to "exfiltrate" data from a specific set of IP addresses.
type CampaignGoal struct {
	BaseModel
	CampaignID      uint     `gorm:"not null" json:"campaign_id"`
	GoalTitle       string   `gorm:"not null" json:"goal_title"`
	GoalDescription string   `json:"goal_description"`
	Status          string   `gorm:"not null" json:"status"`
	TaskID          uint     `json:"task_id"`
	Campaign        Campaign `gorm:"foreignKey:CampaignID" json:"-"`
	Task            Task     `gorm:"foreignKey:TaskID" json:"-"`
}

// Task model, represents a task in a campaign.
// A task is a collection of commands that will be executed on a target.
// A task is owned by a campaign, and can be accessed by all members of the team that owns the campaign.
type Task struct {
	BaseModel
	CampaignID uint      `gorm:"not null" json:"campaign_id"`
	Status     string    `gorm:"not null" json:"status"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID" json:"-"`
	Commands   []Command `gorm:"foreignKey:TaskID" json:"-"`
}

// Command model, represents a command in a task.
// A command is a single command that will be executed on a target.
// A command is owned by a task, and can be accessed by all members of the team that owns the task.
type Command struct {
	BaseModel
	TaskID  uint   `gorm:"not null" json:"task_id"`
	Command string `gorm:"not null" json:"command"`
	Task    Task   `gorm:"foreignKey:TaskID" json:"-"`
}
