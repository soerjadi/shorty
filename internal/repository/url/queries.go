package url

import (
	"context"
	"strconv"
	"strings"
)

func buildQueries(ctx context.Context, req buildQueryParam) buildQueryResult {
	var (
		query     strings.Builder
		result    buildQueryResult
		arguments []interface{}
		where     []string
	)

	query.WriteString(req.query)

	if req.req.Query != nil {
		arguments = append(arguments, req.req.Query)

		argumentLength := "$" + strconv.Itoa(len(arguments))

		where = append(where, "domain = %s OR domain_ext = %s OR short_url = %s ", argumentLength, argumentLength, argumentLength)
	}

	if req.req.SSL != nil {
		arguments = append(arguments, req.req.SSL)

		where = append(where, "is_ssl = %s ", "$"+strconv.Itoa(len(arguments)))
	}

	if len(where) > 0 {
		query.WriteString(" WHERE " + strings.Join(where, " AND "))
	}

	query.WriteString("ORDER BY created_at DESC")

	result.args = arguments
	result.query = query.String()

	return result
}
