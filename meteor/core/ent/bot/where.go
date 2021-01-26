// Code generated by entc, DO NOT EDIT.

package bot

import (
	"github.com/degenerat3/meteor/meteor/core/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Interval applies equality check predicate on the "interval" field. It's identical to IntervalEQ.
func Interval(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterval), v))
	})
}

// Delta applies equality check predicate on the "delta" field. It's identical to DeltaEQ.
func Delta(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDelta), v))
	})
}

// LastSeen applies equality check predicate on the "lastSeen" field. It's identical to LastSeenEQ.
func LastSeen(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeen), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// IntervalEQ applies the EQ predicate on the "interval" field.
func IntervalEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterval), v))
	})
}

// IntervalNEQ applies the NEQ predicate on the "interval" field.
func IntervalNEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterval), v))
	})
}

// IntervalIn applies the In predicate on the "interval" field.
func IntervalIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInterval), v...))
	})
}

// IntervalNotIn applies the NotIn predicate on the "interval" field.
func IntervalNotIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInterval), v...))
	})
}

// IntervalGT applies the GT predicate on the "interval" field.
func IntervalGT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterval), v))
	})
}

// IntervalGTE applies the GTE predicate on the "interval" field.
func IntervalGTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterval), v))
	})
}

// IntervalLT applies the LT predicate on the "interval" field.
func IntervalLT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterval), v))
	})
}

// IntervalLTE applies the LTE predicate on the "interval" field.
func IntervalLTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterval), v))
	})
}

// DeltaEQ applies the EQ predicate on the "delta" field.
func DeltaEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDelta), v))
	})
}

// DeltaNEQ applies the NEQ predicate on the "delta" field.
func DeltaNEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDelta), v))
	})
}

// DeltaIn applies the In predicate on the "delta" field.
func DeltaIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDelta), v...))
	})
}

// DeltaNotIn applies the NotIn predicate on the "delta" field.
func DeltaNotIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDelta), v...))
	})
}

// DeltaGT applies the GT predicate on the "delta" field.
func DeltaGT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDelta), v))
	})
}

// DeltaGTE applies the GTE predicate on the "delta" field.
func DeltaGTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDelta), v))
	})
}

// DeltaLT applies the LT predicate on the "delta" field.
func DeltaLT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDelta), v))
	})
}

// DeltaLTE applies the LTE predicate on the "delta" field.
func DeltaLTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDelta), v))
	})
}

// LastSeenEQ applies the EQ predicate on the "lastSeen" field.
func LastSeenEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeen), v))
	})
}

// LastSeenNEQ applies the NEQ predicate on the "lastSeen" field.
func LastSeenNEQ(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSeen), v))
	})
}

// LastSeenIn applies the In predicate on the "lastSeen" field.
func LastSeenIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastSeen), v...))
	})
}

// LastSeenNotIn applies the NotIn predicate on the "lastSeen" field.
func LastSeenNotIn(vs ...int) predicate.Bot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastSeen), v...))
	})
}

// LastSeenGT applies the GT predicate on the "lastSeen" field.
func LastSeenGT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSeen), v))
	})
}

// LastSeenGTE applies the GTE predicate on the "lastSeen" field.
func LastSeenGTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSeen), v))
	})
}

// LastSeenLT applies the LT predicate on the "lastSeen" field.
func LastSeenLT(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSeen), v))
	})
}

// LastSeenLTE applies the LTE predicate on the "lastSeen" field.
func LastSeenLTE(v int) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSeen), v))
	})
}

// HasInfecting applies the HasEdge predicate on the "infecting" edge.
func HasInfecting() predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InfectingTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InfectingTable, InfectingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInfectingWith applies the HasEdge predicate on the "infecting" edge with a given conditions (other predicates).
func HasInfectingWith(preds ...predicate.Host) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InfectingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InfectingTable, InfectingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bot) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bot) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bot) predicate.Bot {
	return predicate.Bot(func(s *sql.Selector) {
		p(s.Not())
	})
}
