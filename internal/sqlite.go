package internal

var (
	migrations = []string{
		`CREATE TABLE IF NOT EXISTS choomba (
    url TEXT PRIMARY KEY,
    mediatype TEXT NOT NULL,
    tags TEXT NOT NULL,
    auxlinks TEXT,
    summary TEXT NOT NULL,
    savedAt INTEGER NOT NULL
)`,
	}
)

func Migrations() []string {
	return migrations
}

// ConnectionName
// The name may be a filename, e.g., "/tmp/mydata.sqlite", or a URI, in which case it may include a '?' followed by one or more query parameters.
// For example, "file:///tmp/mydata.sqlite?_pragma=foreign_keys(1)&_time_format=sqlite".
// The supported query parameters are:
// _pragma: Each value will be run as a "PRAGMA ..." statement (with the PRAGMA keyword added for you).
// May be specified more than once, '&'-separated.
// For more information on supported PRAGMAs see: https://www.sqlite.org/pragma.html
// _time_format: The name of a format to use when writing time values to the database.
// Currently the only supported value is "sqlite",
// which corresponds to format 7 from https://www.sqlite.org/lang_datefunc.html#time_values, including the timezone specifier.
// If this parameter is not specified, then the default String() format will be used.
// _txlock: The locking behavior to use when beginning a transaction.
// May be "deferred" (the default), "immediate", or "exclusive" (case insensitive).
// See: https://www.sqlite.org/lang_transaction.html#deferred_immediate_and_exclusive_transactions
func ConnectionName() string {
	return "/tmp/choomba.sqlite"
}
