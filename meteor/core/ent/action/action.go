// Code generated by entc, DO NOT EDIT.

package action

const (
	// Label holds the string label denoting the action type in the database.
	Label = "action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// FieldQueued holds the string denoting the queued field in the database.
	FieldQueued = "queued"
	// FieldResponded holds the string denoting the responded field in the database.
	FieldResponded = "responded"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"

	// EdgeTargeting holds the string denoting the targeting edge name in mutations.
	EdgeTargeting = "targeting"

	// Table holds the table name of the action in the database.
	Table = "actions"
	// TargetingTable is the table the holds the targeting relation/edge. The primary key declared below.
	TargetingTable = "host_actions"
	// TargetingInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	TargetingInverseTable = "hosts"
)

// Columns holds all SQL columns for action fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldMode,
	FieldArgs,
	FieldQueued,
	FieldResponded,
	FieldResult,
}

var (
	// TargetingPrimaryKey and TargetingColumn2 are the table columns denoting the
	// primary key for the targeting relation (M2M).
	TargetingPrimaryKey = []string{"host_id", "action_id"}
)
