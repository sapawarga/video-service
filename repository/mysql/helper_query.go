package mysql

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/model"
)

func selectQuery(ctx context.Context, query bytes.Buffer, params *model.GetListVideoRepoRequest) (newQuery bytes.Buffer, queryParams []interface{}) {
	newQuery.Reset()
	if params.RegencyID != nil {
		newQuery.WriteString(andWhere(ctx, query, "kabkota_id", "="))
		queryParams = append(queryParams, converter.GetInt64FromPointer(params.RegencyID))
	}
	if params.Title != nil {
		newQuery.WriteString(andWhere(ctx, query, "title", "LIKE"))
		queryParams = append(queryParams, "%"+converter.GetStringFromPointer(params.Title)+"%")
	}
	if params.CategoryID != nil {
		newQuery.WriteString(andWhere(ctx, query, "category_id", "="))
		queryParams = append(queryParams, converter.GetInt64FromPointer(params.CategoryID))
	}
	return newQuery, queryParams
}

func andWhere(ctx context.Context, query bytes.Buffer, field string, action string) string {
	var newQuery string
	if strings.Contains(query.String(), "WHERE") {
		newQuery = fmt.Sprintf(" AND %s %s ? ", field, action)
	} else {
		newQuery = fmt.Sprintf(" WHERE %s %s ? ", field, action)
	}
	return newQuery
}
func updateQuery(ctx context.Context, fields ...string) string {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(" %s = :%s ", fields[0], fields[0]))
	if len(fields) > 1 {
		for i := 1; i < len(fields); i++ {
			query.WriteString(fmt.Sprintf(" , %s = :%s ", fields[i], fields[i]))
		}
	}
	return query.String()
}
