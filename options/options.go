package options

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type CursorType int8

type FindOne struct {
	AllowPartialResults bool          // If true, allows partial results to be returned if some shards are down.
	BatchSize           int32         // Specifies the number of documents to return in every batch.
	Collation           *Collation    // Specifies a collation to be used
	Comment             string        // Specifies a string to help trace the operation through the database.
	CursorType          *CursorType   // Specifies the type of cursor to use
	Hint                interface{}   // Specifies the index to use.
	Max                 interface{}   // Sets an exclusive upper bound for a specific index
	MaxAwaitTime        time.Duration // Specifies the maximum amount of time for the server to wait on new documents.
	MaxTime             time.Duration // Specifies the maximum amount of time to allow the query to run.
	Min                 interface{}   // Specifies the inclusive lower bound for a specific index.
	NoCursorTimeout     bool          // If true, prevents cursors from timing out after an inactivity period.
	OplogReplay         bool          // Adds an option for internal use only and should not be set.
	Projection          interface{}   // Limits the fields returned for all documents.
	ReturnKey           bool          // If true, only returns index keys for all result documents.
	ShowRecordID        bool          // If true, a $recordId field with the record identifier will be added to the returned documents.
	Skip                int64         // Specifies the number of documents to skip before returning
	Snapshot            bool          // If true, prevents the cursor from returning a document more than once because of an intervening write operation.
	Sort                interface{}   // Specifies the order in which to return results.
}

type Update struct {
	Upsert bool
}

type Client struct {
	options.ClientOptions
}

type Collation struct {
	Locale          string `bson:",omitempty"` // The locale
	CaseLevel       bool   `bson:",omitempty"` // The case level
	CaseFirst       string `bson:",omitempty"` // The case ordering
	Strength        int    `bson:",omitempty"` // The number of comparision levels to use
	NumericOrdering bool   `bson:",omitempty"` // Whether to order numbers based on numerical order and not collation order
	Alternate       string `bson:",omitempty"` // Whether spaces and punctuation are considered base characters
	MaxVariable     string `bson:",omitempty"` // Which characters are affected by alternate: "shifted"
	Backwards       bool   `bson:",omitempty"` // Causes secondary differences to be considered in reverse order, as it is done in the French language
}

type InsertOne struct {
}

type Find struct {
	Page  int64
	Limit int64
}
