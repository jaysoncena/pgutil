package pg

import (
	"strings"

	"github.com/sjansen/pgutil/internal/ddl"
)

// ListParameters describes database configuration parameters
func (c *Conn) ListParameters() (*ddl.Parameters, error) {
	c.log.Infow("listing parameters")

	params := &ddl.Parameters{}

	var tmp string
	c.log.Debugw("SHOW search_path")
	row := c.conn.QueryRow("SHOW search_path")
	err := row.Scan(&tmp)
	if err != nil {
		return nil, err
	}
	params.SearchPath = strings.Split(tmp, ",")

	c.log.Debugw("SHOW timezone")
	row = c.conn.QueryRow("SHOW timezone")
	err = row.Scan(&params.Timezone)
	if err != nil {
		return nil, err
	}

	return params, nil
}
