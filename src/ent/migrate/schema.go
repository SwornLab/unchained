// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssetPricesColumns holds the columns for the "asset_prices" table.
	AssetPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "block", Type: field.TypeUint64},
		{Name: "signers_count", Type: field.TypeUint64, Nullable: true},
		{Name: "price", Type: field.TypeUint, SchemaType: map[string]string{"postgres": "numeric(78, 0)", "sqlite3": "numeric(78, 0)"}},
		{Name: "signature", Type: field.TypeBytes, Size: 96},
		{Name: "asset", Type: field.TypeString, Nullable: true},
		{Name: "chain", Type: field.TypeString, Nullable: true},
		{Name: "pair", Type: field.TypeString, Nullable: true},
	}
	// AssetPricesTable holds the schema information for the "asset_prices" table.
	AssetPricesTable = &schema.Table{
		Name:       "asset_prices",
		Columns:    AssetPricesColumns,
		PrimaryKey: []*schema.Column{AssetPricesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "assetprice_block_chain_asset_pair",
				Unique:  true,
				Columns: []*schema.Column{AssetPricesColumns[1], AssetPricesColumns[6], AssetPricesColumns[5], AssetPricesColumns[7]},
			},
		},
	}
	// CorrectnessReportsColumns holds the columns for the "correctness_reports" table.
	CorrectnessReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "signers_count", Type: field.TypeUint64},
		{Name: "timestamp", Type: field.TypeUint64},
		{Name: "signature", Type: field.TypeBytes, Size: 48},
		{Name: "hash", Type: field.TypeBytes, Size: 64},
		{Name: "topic", Type: field.TypeBytes, Size: 64},
		{Name: "correct", Type: field.TypeBool},
	}
	// CorrectnessReportsTable holds the schema information for the "correctness_reports" table.
	CorrectnessReportsTable = &schema.Table{
		Name:       "correctness_reports",
		Columns:    CorrectnessReportsColumns,
		PrimaryKey: []*schema.Column{CorrectnessReportsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "correctnessreport_topic_hash",
				Unique:  true,
				Columns: []*schema.Column{CorrectnessReportsColumns[5], CorrectnessReportsColumns[4]},
			},
			{
				Name:    "correctnessreport_topic_timestamp_hash",
				Unique:  false,
				Columns: []*schema.Column{CorrectnessReportsColumns[5], CorrectnessReportsColumns[2], CorrectnessReportsColumns[4]},
			},
		},
	}
	// EventLogsColumns holds the columns for the "event_logs" table.
	EventLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "block", Type: field.TypeUint64},
		{Name: "signers_count", Type: field.TypeUint64},
		{Name: "signature", Type: field.TypeBytes, Size: 96},
		{Name: "address", Type: field.TypeString},
		{Name: "chain", Type: field.TypeString},
		{Name: "index", Type: field.TypeUint64},
		{Name: "event", Type: field.TypeString},
		{Name: "transaction", Type: field.TypeBytes, Size: 32},
		{Name: "args", Type: field.TypeJSON},
	}
	// EventLogsTable holds the schema information for the "event_logs" table.
	EventLogsTable = &schema.Table{
		Name:       "event_logs",
		Columns:    EventLogsColumns,
		PrimaryKey: []*schema.Column{EventLogsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "eventlog_block_transaction_index",
				Unique:  true,
				Columns: []*schema.Column{EventLogsColumns[1], EventLogsColumns[8], EventLogsColumns[6]},
			},
			{
				Name:    "eventlog_block_address_event",
				Unique:  false,
				Columns: []*schema.Column{EventLogsColumns[1], EventLogsColumns[4], EventLogsColumns[7]},
			},
		},
	}
	// SignersColumns holds the columns for the "signers" table.
	SignersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "evm", Type: field.TypeString, Nullable: true},
		{Name: "key", Type: field.TypeBytes, Unique: true, Size: 96},
		{Name: "shortkey", Type: field.TypeBytes, Unique: true, Size: 96},
		{Name: "points", Type: field.TypeInt64},
		{Name: "correctness_report_signers", Type: field.TypeInt, Nullable: true},
	}
	// SignersTable holds the schema information for the "signers" table.
	SignersTable = &schema.Table{
		Name:       "signers",
		Columns:    SignersColumns,
		PrimaryKey: []*schema.Column{SignersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "signers_correctness_reports_signers",
				Columns:    []*schema.Column{SignersColumns[6]},
				RefColumns: []*schema.Column{CorrectnessReportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "signer_key",
				Unique:  true,
				Columns: []*schema.Column{SignersColumns[3]},
			},
			{
				Name:    "signer_shortkey",
				Unique:  true,
				Columns: []*schema.Column{SignersColumns[4]},
			},
		},
	}
	// AssetPriceSignersColumns holds the columns for the "asset_price_signers" table.
	AssetPriceSignersColumns = []*schema.Column{
		{Name: "asset_price_id", Type: field.TypeInt},
		{Name: "signer_id", Type: field.TypeInt},
	}
	// AssetPriceSignersTable holds the schema information for the "asset_price_signers" table.
	AssetPriceSignersTable = &schema.Table{
		Name:       "asset_price_signers",
		Columns:    AssetPriceSignersColumns,
		PrimaryKey: []*schema.Column{AssetPriceSignersColumns[0], AssetPriceSignersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "asset_price_signers_asset_price_id",
				Columns:    []*schema.Column{AssetPriceSignersColumns[0]},
				RefColumns: []*schema.Column{AssetPricesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "asset_price_signers_signer_id",
				Columns:    []*schema.Column{AssetPriceSignersColumns[1]},
				RefColumns: []*schema.Column{SignersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventLogSignersColumns holds the columns for the "event_log_signers" table.
	EventLogSignersColumns = []*schema.Column{
		{Name: "event_log_id", Type: field.TypeInt},
		{Name: "signer_id", Type: field.TypeInt},
	}
	// EventLogSignersTable holds the schema information for the "event_log_signers" table.
	EventLogSignersTable = &schema.Table{
		Name:       "event_log_signers",
		Columns:    EventLogSignersColumns,
		PrimaryKey: []*schema.Column{EventLogSignersColumns[0], EventLogSignersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "event_log_signers_event_log_id",
				Columns:    []*schema.Column{EventLogSignersColumns[0]},
				RefColumns: []*schema.Column{EventLogsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "event_log_signers_signer_id",
				Columns:    []*schema.Column{EventLogSignersColumns[1]},
				RefColumns: []*schema.Column{SignersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetPricesTable,
		CorrectnessReportsTable,
		EventLogsTable,
		SignersTable,
		AssetPriceSignersTable,
		EventLogSignersTable,
	}
)

func init() {
	SignersTable.ForeignKeys[0].RefTable = CorrectnessReportsTable
	AssetPriceSignersTable.ForeignKeys[0].RefTable = AssetPricesTable
	AssetPriceSignersTable.ForeignKeys[1].RefTable = SignersTable
	EventLogSignersTable.ForeignKeys[0].RefTable = EventLogsTable
	EventLogSignersTable.ForeignKeys[1].RefTable = SignersTable
}
