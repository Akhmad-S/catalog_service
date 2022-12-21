package postgres

import (
	"errors"

	ecom "github.com/uacademy/e_commerce/catalog_service/proto-gen/e_commerce"
)

func (stg Postgres) CreateCategory(id string, input *ecom.CreateCategoryRequest) error {
	_, err := stg.db.Exec(`INSERT INTO "category" (id, category_title) VALUES ($1, $2)`, id, input.CategoryTitle)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) GetCategoryList(offset, limit int, search string) (resp *ecom.GetCategoryListResponse, err error) {
	resp = &ecom.GetCategoryListResponse{
		Categories: make([]*ecom.Category, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	category_title,
	created_at,
	updated_at
	FROM "category" WHERE deleted_at IS NULL AND (category_title ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		c := &ecom.Category{}
		var updatedAt *string

		err := rows.Scan(
			&c.Id,
			&c.CategoryTitle,
			&c.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil {
			c.UpdatedAt = *updatedAt
		}

		resp.Categories = append(resp.Categories, c)
	}

	return resp, err
}

func (stg Postgres) GetCategoryById(id string) (*ecom.GetCategoryByIdResponse, error) {
	res := &ecom.GetCategoryByIdResponse{}
	var updatedAt *string

	err := stg.db.QueryRow(`SELECT id, category_title, created_at, updated_at FROM category WHERE id=$1 AND deleted_at IS NULL`, id).Scan(
		&res.Id, &res.CategoryTitle, &res.CreatedAt, &updatedAt,
	)

	if err != nil {
		return nil, errors.New("category not found")
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	return res, nil
}

func (stg Postgres) UpdateCategory(input *ecom.UpdateCategoryRequest) error {
	res, err := stg.db.NamedExec(`UPDATE "category"  SET category_title=:ct, updated_at=now() WHERE deleted_at IS NULL AND id=:id`, map[string]interface{}{
		"id": input.Id,
		"ct": input.CategoryTitle,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("category not found")
}

func (stg Postgres) DeleteCategory(id string) error {
	res, err := stg.db.Exec(`UPDATE "category"  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}
	return errors.New("categpry not found")
}
