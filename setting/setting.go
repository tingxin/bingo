package setting

import "fmt"

var (
	// DataDBConnStr used to conn data source db
	DataDBConnStr = fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "Tech@915", "10.128.234.165:4000", "iris_test")

	// MetaDBConnStr used to conn meta source db
	MetaDBConnStr = fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "Tech@915", "10.128.234.165:4000", "iris_meta")

	// AuthKey used to
	AuthKey = "token"
)
