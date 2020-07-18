package all

// Importing this package will import all supported backends
// and register them in Repository registry
import (
	_ "github.com/theycallmemac/odin/odin-engine/pkg/repository/nosql"
)
