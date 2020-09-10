// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/degenerat3/meteor/meteor/core/ent/host"
	"github.com/facebook/ent/dialect/sql"
)

// Host is the model entity for the Host schema.
type Host struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// Interface holds the value of the "interface" field.
	Interface string `json:"interface,omitempty"`
	// LastSeen holds the value of the "lastSeen" field.
	LastSeen int `json:"lastSeen,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostQuery when eager-loading is set.
	Edges HostEdges `json:"edges"`
}

// HostEdges holds the relations/edges for other nodes in the graph.
type HostEdges struct {
	// Bots holds the value of the bots edge.
	Bots []*Bot
	// Actions holds the value of the actions edge.
	Actions []*Action
	// Member holds the value of the member edge.
	Member []*Group
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BotsOrErr returns the Bots value or an error if the edge
// was not loaded in eager-loading.
func (e HostEdges) BotsOrErr() ([]*Bot, error) {
	if e.loadedTypes[0] {
		return e.Bots, nil
	}
	return nil, &NotLoadedError{edge: "bots"}
}

// ActionsOrErr returns the Actions value or an error if the edge
// was not loaded in eager-loading.
func (e HostEdges) ActionsOrErr() ([]*Action, error) {
	if e.loadedTypes[1] {
		return e.Actions, nil
	}
	return nil, &NotLoadedError{edge: "actions"}
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading.
func (e HostEdges) MemberOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "member"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Host) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // hostname
		&sql.NullString{}, // interface
		&sql.NullInt64{},  // lastSeen
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Host fields.
func (h *Host) assignValues(values ...interface{}) error {
	if m, n := len(values), len(host.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	h.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hostname", values[0])
	} else if value.Valid {
		h.Hostname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field interface", values[1])
	} else if value.Valid {
		h.Interface = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field lastSeen", values[2])
	} else if value.Valid {
		h.LastSeen = int(value.Int64)
	}
	return nil
}

// QueryBots queries the bots edge of the Host.
func (h *Host) QueryBots() *BotQuery {
	return (&HostClient{config: h.config}).QueryBots(h)
}

// QueryActions queries the actions edge of the Host.
func (h *Host) QueryActions() *ActionQuery {
	return (&HostClient{config: h.config}).QueryActions(h)
}

// QueryMember queries the member edge of the Host.
func (h *Host) QueryMember() *GroupQuery {
	return (&HostClient{config: h.config}).QueryMember(h)
}

// Update returns a builder for updating this Host.
// Note that, you need to call Host.Unwrap() before calling this method, if this Host
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Host) Update() *HostUpdateOne {
	return (&HostClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *Host) Unwrap() *Host {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Host is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Host) String() string {
	var builder strings.Builder
	builder.WriteString("Host(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", hostname=")
	builder.WriteString(h.Hostname)
	builder.WriteString(", interface=")
	builder.WriteString(h.Interface)
	builder.WriteString(", lastSeen=")
	builder.WriteString(fmt.Sprintf("%v", h.LastSeen))
	builder.WriteByte(')')
	return builder.String()
}

// Hosts is a parsable slice of Host.
type Hosts []*Host

func (h Hosts) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
