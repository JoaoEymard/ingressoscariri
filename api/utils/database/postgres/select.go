package postgres

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"strings"

	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// SetParams Seta os Params para executar a SQL
func SetParams(params url.Values, filtros map[string]string) (string, string, string, string, error) {

	var (
		filter, order, limit, offset string
		filters                      []string
	)

	if params.Get("filtro") != "" {
		var jsonFilters map[string]interface{}

		err := json.Unmarshal([]byte(params.Get("filtro")), &jsonFilters)
		if err != nil {
			return "", "", "", "", err
		}

		if jsonFilters != nil {
			for key, value := range jsonFilters {
				if filtros[key] != "" {
					filters = append(filters, fmt.Sprintf(filtros[key], value))
				}
			}

			if len(filters) > 1 {
				filter = fmt.Sprintf("WHERE %v", strings.Join(filters[0:len(filters)], " AND "))
			} else if filters != nil {
				filter = fmt.Sprintf("WHERE %v", strings.Join(filters, ""))
			}
		}
	}

	if params.Get("ordene") != "" && params.Get("tipo") != "" {
		if filtros[params.Get("ordene")] != "" {
			order = fmt.Sprintf("ORDER BY %v %v", params.Get("ordene"), strings.ToUpper(params.Get("tipo")))
		}
	}

	if params.Get("limite") != "" {
		_, err := strconv.Atoi(params.Get("limite"))
		if err == nil && params.Get("limite") != "0" {
			limit = fmt.Sprintf("LIMIT %v", params.Get("limite"))
		}
	} else {
		limit = fmt.Sprintf("LIMIT 15")
	}

	if params.Get("desloque") != "" {
		_, err := strconv.Atoi(params.Get("desloque"))
		if err == nil {
			offset = fmt.Sprintf("OFFSET %v", params.Get("desloque"))
		}
	}

	return filter, order, limit, offset, nil
}

// SelectOne Coleta e retorna um map com o nome das colunas e o valor recebido do banco
func SelectOne(query string) (map[string]interface{}, error) {

	var (
		values map[string]interface{}
	)

	rows, err := postgres.Query(query)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()

	for rows.Next() {

		var (
			rowsValues = make(map[string]interface{}, len(columns))
			refs       = make([]interface{}, 0, len(columns))
		)

		for _, column := range columns {
			var ref interface{}
			rowsValues[column] = &ref
			refs = append(refs, &ref)
		}

		rows.Scan(refs...)

		values = rowsValues

	}

	return values, nil
}

// Select Coleta e retorna um array map com o nome das colunas e os valores recebido do banco
func Select(query string) ([]map[string]interface{}, error) {

	var (
		values []map[string]interface{}
	)

	rows, err := postgres.Query(query)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()

	for rows.Next() {

		var (
			rowsValues = make(map[string]interface{}, len(columns))
			refs       = make([]interface{}, 0, len(columns))
		)
		// var id, nome, ultimoAcesso, ativo, cpf, dataNascimento, sexo, nivel, contatos interface{}

		for _, column := range columns {
			var ref interface{}
			refs = append(refs, &ref)
			rowsValues[column] = &ref
		}

		rows.Scan(refs...)
		// rows.Scan(&id, &nome, &ultimoAcesso, &ativo, &cpf, &dataNascimento, &sexo, &nivel, &contatos)

		fmt.Printf("%T\n", refs[8])

		values = append(values, rowsValues)

	}

	return values, nil
}
