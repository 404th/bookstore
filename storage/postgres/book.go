package postgres

import (
	"errors"

	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type bookPg struct {
	db *sqlx.DB
}

//////////////
func (bk *bookPg) CreateBook(details models.Book) (string, error) {
	var res string
	var count_cat int
	var count_auth int

	q1 := `SELECT count(1) FROM book_category WHERE id=$1;`
	row1 := bk.db.QueryRow(q1, details.CategoryID)
	if err := row1.Scan(&count_cat); err != nil {
		return res, err
	}

	if count_cat < 1 {
		return res, errors.New("there is no category_name with the given id")
	}

	q2 := `SELECT count(1) FROM author WHERE id=$1;`
	row2 := bk.db.QueryRow(q2, details.AuthorID)
	if err := row2.Scan(&count_auth); err != nil {
		return res, err
	}
	if count_auth < 1 {
		return res, errors.New("there is no author with the given id")
	}

	query := `INSERT INTO book (
		id,
		category_id,
		author_id,
		name,
		price,
		definition,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	) RETURNING id;`

	row := bk.db.QueryRow(query,
		details.ID,
		details.CategoryID,
		details.AuthorID,
		details.Name,
		details.Price,
		details.Definition,
		details.CreatedAt,
		details.UpdatedAt,
	)

	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

func (bk *bookPg) GetBook(id string) (models.Book, error) {
	var res models.Book
	query := `
				SELECT
					id,
					category_id,
					author_id,
					name,
					price,
					definition,
					created_at,
					updated_at
				FROM
					book
				WHERE id=$1;
			`

	row := bk.db.QueryRow(query, id)
	if err := row.Scan(
		&res.ID,
		&res.CategoryID,
		&res.AuthorID,
		&res.Name,
		&res.Price,
		&res.Definition,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}

func (bk *bookPg) GetAllBooks(queryParam models.ApplicationQueryParamModel) ([]models.Book, error) {
	var res []models.Book = []models.Book{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		category_id,
		author_id,
		name,
		price,
		definition,
		created_at,
		updated_at
	FROM
		book`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (name ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM book" + filter
	row, err := bk.db.NamedQuery(cQ, params)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	q := query + filter + offset + limit
	rows, err := bk.db.NamedQuery(q, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var aut models.Book
		err = rows.Scan(
			&aut.ID,
			&aut.CategoryID,
			&aut.AuthorID,
			&aut.Name,
			&aut.Price,
			&aut.Definition,
			&aut.CreatedAt,
			&aut.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		res = append(res, aut)
	}

	return res, nil
}

func (bk *bookPg) UpdateBook(details models.UpdateBook, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id
	params["updated_at"] = details.UpdatedAt

	query := `UPDATE book SET `

	if len(details.CategoryID) > 0 {
		params["category_id"] = details.CategoryID
		query += `category_id = :category_id,`
	}

	if len(details.AuthorID) > 0 {
		params["author_id"] = details.AuthorID
		query += `author_id = :author_id,`
	}

	if len(details.Name) > 0 {
		params["name"] = details.Name
		query += `name = :name,`
	}
	if details.Price != 0 {
		params["price"] = details.Price
		query += `price = :price,`
	}
	if len(details.Definition) > 0 {
		params["definition"] = details.AuthorID
		query += `definition = :definition,`
	}

	query += `updated_at = :updated_at WHERE id =:id`

	result, err := bk.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (bk *bookPg) DeleteBook(id string) (int64, error) {
	query := `DELETE FROM book WHERE id = $1`

	result, err := bk.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
