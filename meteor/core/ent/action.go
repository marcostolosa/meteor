// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/degenerat3/meteor/meteor/core/ent/action"
	"github.com/facebook/ent/dialect/sql"
)

// Action is the model entity for the Action schema.
type Action struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Mode holds the value of the "mode" field.
	Mode string `json:"mode,omitempty"`
	// Args holds the value of the "args" field.
	Args string `json:"args,omitempty"`
	// Queued holds the value of the "queued" field.
	Queued bool `json:"queued,omitempty"`
	// Responded holds the value of the "responded" field.
	Responded bool `json:"responded,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActionQuery when eager-loading is set.
	Edges ActionEdges `json:"edges"`
}

// ActionEdges holds the relations/edges for other nodes in the graph.
type ActionEdges struct {
	// Targeting holds the value of the targeting edge.
	Targeting []*Host
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TargetingOrErr returns the Targeting value or an error if the edge
// was not loaded in eager-loading.
func (e ActionEdges) TargetingOrErr() ([]*Host, error) {
	if e.loadedTypes[0] {
		return e.Targeting, nil
	}
	return nil, &NotLoadedError{edge: "targeting"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Action) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // uuid
		&sql.NullString{}, // mode
		&sql.NullString{}, // args
		&sql.NullBool{},   // queued
		&sql.NullBool{},   // responded
		&sql.NullString{}, // result
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Action fields.
func (a *Action) assignValues(values ...interface{}) error {
	if m, n := len(values), len(action.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uuid", values[0])
	} else if value.Valid {
		a.UUID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mode", values[1])
	} else if value.Valid {
		a.Mode = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field args", values[2])
	} else if value.Valid {
		a.Args = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field queued", values[3])
	} else if value.Valid {
		a.Queued = value.Bool
	}
	if value, ok := values[4].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field responded", values[4])
	} else if value.Valid {
		a.Responded = value.Bool
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field result", values[5])
	} else if value.Valid {
		a.Result = value.String
	}
	return nil
}

// QueryTargeting queries the targeting edge of the Action.
func (a *Action) QueryTargeting() *HostQuery {
	return (&ActionClient{config: a.config}).QueryTargeting(a)
}

// Update returns a builder for updating this Action.
// Note that, you need to call Action.Unwrap() before calling this method, if this Action
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Action) Update() *ActionUpdateOne {
	return (&ActionClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Action) Unwrap() *Action {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Action is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Action) String() string {
	var builder strings.Builder
	builder.WriteString("Action(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(a.UUID)
	builder.WriteString(", mode=")
	builder.WriteString(a.Mode)
	builder.WriteString(", args=")
	builder.WriteString(a.Args)
	builder.WriteString(", queued=")
	builder.WriteString(fmt.Sprintf("%v", a.Queued))
	builder.WriteString(", responded=")
	builder.WriteString(fmt.Sprintf("%v", a.Responded))
	builder.WriteString(", result=")
	builder.WriteString(a.Result)
	builder.WriteByte(')')
	return builder.String()
}

// Actions is a parsable slice of Action.
type Actions []*Action

func (a Actions) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
