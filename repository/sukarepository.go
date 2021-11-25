package repository

import (
	"example/micro/model"
	"fmt"
)

func FindAllSuka() ([]model.Suka, error) {
	// An albums slice to hold data from returned rows.
	var sukas []model.Suka

	rows, err := db.Query("SELECT * FROM suka")
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb model.Suka
		if err := rows.Scan(&alb.Id, &alb.Name, &alb.Surname); err != nil {
			return nil, fmt.Errorf("albumsByArtist : %v", err)
		}
		sukas = append(sukas, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %v", err)
	}
	return sukas, nil
}

func CreateSuka(suka model.Suka) (model.Suka, error) {

	insertQuery := "INSERT INTO suka (name, surname) VALUES (?, ?)"
	result, err := db.Exec(insertQuery, suka.Name, suka.Surname)

	if err != nil {
		return suka, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return suka, fmt.Errorf("addAlbum: %v", err)
	}
	suka.Id = int(id)
	return suka, nil
}
