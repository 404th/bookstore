package postgres

import (
	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type authorPg struct {
	db *sqlx.DB
}

func (ar *authorPg) CreateAuthor(details models.Author) (string, error) {
	var res string

	query := `INSERT INTO author (
		id,
		firstname,
		secondname,
		email,
		age,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	) RETURNING id;`

	row := ar.db.QueryRow(query, details.ID, details.Firstname, details.Secondname, details.Email, details.Age, details.CreatedAt, details.UpdatedAt)

	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

func (ar *authorPg) GetAuthor(id string) (models.Author, error) {
	var res models.Author
	query := `
				SELECT
					id,
					firstname,
					secondname,
					email,
					age,
					created_at,
					updated_at
				FROM
					author
				WHERE id=$1;
			`

	row := ar.db.QueryRow(query, id)
	if err := row.Scan(
		&res.ID,
		&res.Firstname,
		&res.Secondname,
		&res.Email,
		&res.Age,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}

func (ar *authorPg) GetAllAuthors(queryParam models.ApplicationQueryParamModel) ([]models.Author, error) {
	var res []models.Author = []models.Author{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		firstname,
		secondname,
		email,
		age,
		created_at,
		updated_at
	FROM
		author`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (firstname ILIKE '%' || :search || '%') OR (secondname ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM author" + filter
	row, err := ar.db.NamedQuery(cQ, params)
	if err != nil {
		return res, err
	}
	defer row.Close()

	q := query + filter + offset + limit
	rows, err := ar.db.NamedQuery(q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var aut models.Author
		err = rows.Scan(
			&aut.ID,
			&aut.Firstname,
			&aut.Secondname,
			&aut.Email,
			&aut.Age,
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

func (ar *authorPg) UpdateAuthor(details models.UpdateAuthor, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id

	query := `UPDATE author SET `

	if len(details.Firstname) > 0 {
		params["firstname"] = details.Firstname
		query += `firstname = :firstname,`
	}

	if len(details.Secondname) > 0 {
		params["secondname"] = details.Secondname
		query += `secondname = :secondname,`
	}

	if len(details.Email) > 0 {
		params["email"] = details.Email
		query += `email = :email,`
	}

	if details.Age > 0 {
		params["firstname"] = details.Firstname
		query += `firstname = :firstname,`
	}

	query += `updated_at = now() WHERE id =:id`

	result, err := ar.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (ar *authorPg) DeleteAuthor(id string) (int64, error) {
	query := `DELETE FROM author WHERE id = $1`

	result, err := ar.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
