package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToProductGroup(rows *sql.Rows) (*ProductGroup, error) {
	defer rows.Close()
	productGroup := ProductGroup{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&productGroup.ProductGroup,
			&productGroup.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productGroup, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productGroup, nil
	}

	return &productGroup, nil
}
