package firebird

import (
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

// custom type definitions for firebird come in here
func FirebirdConnectString(username string, password string, hostname string, database string) string {
	return fmt.Sprintf("%s:%s@%s/%s", username, password, hostname, database)
}
