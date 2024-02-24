// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// AssetPrice is the predicate function for assetprice builders.
type AssetPrice func(*sql.Selector)

// AssetPriceOrErr calls the predicate only if the error is not nit.
func AssetPriceOrErr(p AssetPrice, err error) AssetPrice {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// EventLog is the predicate function for eventlog builders.
type EventLog func(*sql.Selector)

// Signer is the predicate function for signer builders.
type Signer func(*sql.Selector)
