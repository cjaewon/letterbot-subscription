package models

import "github.com/jinzhu/gorm"

// Webhook : webhook model
type Webhook struct {
	gorm.Model
	URL string
}
