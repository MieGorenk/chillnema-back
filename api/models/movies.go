package models

import (
	"github.com/jinzhu/gorm"
)

// Movie model
type Movie struct {
	ID       int    `gorm:"AUTO_INCREMENT;not_null"`
	Name     string `gorm:"type:varchar(75)"             json:name`
	Runtime  string `gorm:"type:varchar(150)"         json:runtime`
	Released string `gorm:"type:varchar(75)"         json:released`
	Rated    string `gorm:"type:varchar(75)"            json:rated`
	Plot     string `gorm:"type:varchar(255)"            json:plot`
	Source   string `gorm:"type:varchar(300):not_null" json:source`
}

// SaveMovie saving movie obj to database
func (movie *Movie) SaveMovie(db *gorm.DB) (*Movie, error) {
	var err error

	// Debug the whole insert operation
	err = db.Debug().Create(&movie).Error
	if err != nil {
		return &Movie{}, err
	}

	return movie, nil
}

//GetAllMovies getting all movies
func GetAllMovies(db *gorm.DB) (*[]Movie, error) {
	movies := []Movie{}
	if err := db.Debug().Table("main.movies").Find(&movies).Error; err != nil {
		return &[]Movie{}, err
	}
	return &movies, nil
}

//GetMovieByID getting movie by id
func GetMovieByID(id int, db *gorm.DB) (*Movie, error) {
	movie := &Movie{}
	if err := db.Debug().Table("main.movies").Where("id = ?", id).First(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}
