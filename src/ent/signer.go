// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KenshiTech/unchained/ent/signer"
)

// Signer is the model entity for the Signer schema.
type Signer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Key holds the value of the "key" field.
	Key []byte `json:"key,omitempty"`
	// Points holds the value of the "points" field.
	Points int64 `json:"points,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SignerQuery when eager-loading is set.
	Edges        SignerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SignerEdges holds the relations/edges for other nodes in the graph.
type SignerEdges struct {
	// AssetPrice holds the value of the AssetPrice edge.
	AssetPrice []*AssetPrice `json:"AssetPrice,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssetPriceOrErr returns the AssetPrice value or an error if the edge
// was not loaded in eager-loading.
func (e SignerEdges) AssetPriceOrErr() ([]*AssetPrice, error) {
	if e.loadedTypes[0] {
		return e.AssetPrice, nil
	}
	return nil, &NotLoadedError{edge: "AssetPrice"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Signer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case signer.FieldKey:
			values[i] = new([]byte)
		case signer.FieldID, signer.FieldPoints:
			values[i] = new(sql.NullInt64)
		case signer.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Signer fields.
func (s *Signer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case signer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case signer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case signer.FieldKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value != nil {
				s.Key = *value
			}
		case signer.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				s.Points = value.Int64
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Signer.
// This includes values selected through modifiers, order, etc.
func (s *Signer) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryAssetPrice queries the "AssetPrice" edge of the Signer entity.
func (s *Signer) QueryAssetPrice() *AssetPriceQuery {
	return NewSignerClient(s.config).QueryAssetPrice(s)
}

// Update returns a builder for updating this Signer.
// Note that you need to call Signer.Unwrap() before calling this method if this Signer
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Signer) Update() *SignerUpdateOne {
	return NewSignerClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Signer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Signer) Unwrap() *Signer {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Signer is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Signer) String() string {
	var builder strings.Builder
	builder.WriteString("Signer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(fmt.Sprintf("%v", s.Key))
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", s.Points))
	builder.WriteByte(')')
	return builder.String()
}

// Signers is a parsable slice of Signer.
type Signers []*Signer
