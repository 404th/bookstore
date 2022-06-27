package postgres

import (
	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type bookCategoryPg struct {
	db *sqlx.DB
}

func (bc *bookCategoryPg) CreateBookCategory(details models.BookCategory) (string, error) {
	var res string

	query := `INSERT INTO book_category (
		id,
		category_name,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING id;`

	row := bc.db.QueryRow(query, details.ID, details.CategoryName, details.CreatedAt, details.UpdatedAt)

	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

func (bc *bookCategoryPg) GetBookCategory(id string) (models.BookCategory, error) {
	var res models.BookCategory
	query := `
				SELECT
					id,
					category_name,
					created_at,
					updated_at
				FROM
					book_category
				WHERE id=$1;
			`

	row := bc.db.QueryRow(query, id)
	if err := row.Scan(
		&res.ID,
		&res.CategoryName,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}

func (bc *bookCategoryPg) GetAllBookCategories(queryParam models.ApplicationQueryParamModel) ([]models.BookCategory, error) {
	var res []models.BookCategory = []models.BookCategory{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		category_name,
		created_at,
		updated_at
	FROM
		book_category`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (category_name ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM book_category" + filter
	row, err := bc.db.NamedQuery(cQ, params)
	if err != nil {
		return res, err
	}
	defer row.Close()

	q := query + filter + offset + limit
	rows, err := bc.db.NamedQuery(q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var aut models.BookCategory
		err = rows.Scan(
			&aut.ID,
			&aut.CategoryName,
			&aut.CreatedAt,
			&aut.UpdatedAt,
		)

		if err != nil {
			return res, err
		}
		res = append(res, aut)
	}

	return res, nil
}

func (bc *bookCategoryPg) UpdateBookCategory(details *models.UpdateBookCategory, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id

	query := `UPDATE book_category SET `

	if len(details.CategoryName) > 0 {
		params["category_name"] = details.CategoryName
		query += `category_name = :category_name,`
	}

	query += `updated_at = now() WHERE id =:id`

	result, err := bc.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (bc *bookCategoryPg) DeleteBookCategory(id string) (int64, error) {
	query := `DELETE FROM book_category WHERE id = $1`

	result, err := bc.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
