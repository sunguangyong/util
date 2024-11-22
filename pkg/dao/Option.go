package dao

import "gorm.io/gorm"

type Option func(db *gorm.DB) *gorm.DB
