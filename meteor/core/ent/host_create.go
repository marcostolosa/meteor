// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/degenerat3/meteor/meteor/core/ent/action"
	"github.com/degenerat3/meteor/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/meteor/core/ent/group"
	"github.com/degenerat3/meteor/meteor/core/ent/host"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// HostCreate is the builder for creating a Host entity.
type HostCreate struct {
	config
	mutation *HostMutation
	hooks    []Hook
}

// SetHostname sets the hostname field.
func (hc *HostCreate) SetHostname(s string) *HostCreate {
	hc.mutation.SetHostname(s)
	return hc
}

// SetInterface sets the interface field.
func (hc *HostCreate) SetInterface(s string) *HostCreate {
	hc.mutation.SetInterface(s)
	return hc
}

// SetLastSeen sets the lastSeen field.
func (hc *HostCreate) SetLastSeen(i int) *HostCreate {
	hc.mutation.SetLastSeen(i)
	return hc
}

// AddBotIDs adds the bots edge to Bot by ids.
func (hc *HostCreate) AddBotIDs(ids ...int) *HostCreate {
	hc.mutation.AddBotIDs(ids...)
	return hc
}

// AddBots adds the bots edges to Bot.
func (hc *HostCreate) AddBots(b ...*Bot) *HostCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return hc.AddBotIDs(ids...)
}

// AddActionIDs adds the actions edge to Action by ids.
func (hc *HostCreate) AddActionIDs(ids ...int) *HostCreate {
	hc.mutation.AddActionIDs(ids...)
	return hc
}

// AddActions adds the actions edges to Action.
func (hc *HostCreate) AddActions(a ...*Action) *HostCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hc.AddActionIDs(ids...)
}

// AddMemberIDs adds the member edge to Group by ids.
func (hc *HostCreate) AddMemberIDs(ids ...int) *HostCreate {
	hc.mutation.AddMemberIDs(ids...)
	return hc
}

// AddMember adds the member edges to Group.
func (hc *HostCreate) AddMember(g ...*Group) *HostCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return hc.AddMemberIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (hc *HostCreate) Mutation() *HostMutation {
	return hc.mutation
}

// Save creates the Host in the database.
func (hc *HostCreate) Save(ctx context.Context) (*Host, error) {
	if err := hc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Host
	)
	if len(hc.hooks) == 0 {
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hc.mutation = mutation
			node, err = hc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HostCreate) SaveX(ctx context.Context) *Host {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hc *HostCreate) preSave() error {
	if _, ok := hc.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New("ent: missing required field \"hostname\"")}
	}
	if _, ok := hc.mutation.Interface(); !ok {
		return &ValidationError{Name: "interface", err: errors.New("ent: missing required field \"interface\"")}
	}
	if _, ok := hc.mutation.LastSeen(); !ok {
		return &ValidationError{Name: "lastSeen", err: errors.New("ent: missing required field \"lastSeen\"")}
	}
	if v, ok := hc.mutation.LastSeen(); ok {
		if err := host.LastSeenValidator(v); err != nil {
			return &ValidationError{Name: "lastSeen", err: fmt.Errorf("ent: validator failed for field \"lastSeen\": %w", err)}
		}
	}
	return nil
}

func (hc *HostCreate) sqlSave(ctx context.Context) (*Host, error) {
	h, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	h.ID = int(id)
	return h, nil
}

func (hc *HostCreate) createSpec() (*Host, *sqlgraph.CreateSpec) {
	var (
		h     = &Host{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: host.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: host.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Hostname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldHostname,
		})
		h.Hostname = value
	}
	if value, ok := hc.mutation.Interface(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldInterface,
		})
		h.Interface = value
	}
	if value, ok := hc.mutation.LastSeen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastSeen,
		})
		h.LastSeen = value
	}
	if nodes := hc.mutation.BotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   host.BotsTable,
			Columns: host.BotsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.ActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   host.ActionsTable,
			Columns: host.ActionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   host.MemberTable,
			Columns: host.MemberPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return h, _spec
}

// HostCreateBulk is the builder for creating a bulk of Host entities.
type HostCreateBulk struct {
	config
	builders []*HostCreate
}

// Save creates the Host entities in the database.
func (hcb *HostCreateBulk) Save(ctx context.Context) ([]*Host, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Host, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*HostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (hcb *HostCreateBulk) SaveX(ctx context.Context) []*Host {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
