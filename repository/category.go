package repository

import (
	"example/micro/model"
	"fmt"
)

func FindAllCategory() ([]model.Category, error) {
	// An albums slice to hold data from returned rows.
	category_list := make([]model.Category, 0)

	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		return nil, fmt.Errorf("FindAllCategory %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			return nil, fmt.Errorf("FindAllCategory %v", err)
		}
		category_list = append(category_list, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAllCategory %v", err)
	}
	return category_list, nil
}

func CreateCategory(category *model.Category) (*model.Category, error) {

	insertQuery := "INSERT INTO category (name) VALUES (?)"
	result, err := db.Exec(insertQuery, category.Name)

	if err != nil {
		return category, fmt.Errorf("on CreateCategory: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return category, fmt.Errorf("on CreateCategory: %v", err)
	}
	category.Id = int(id)
	return category, nil
}
