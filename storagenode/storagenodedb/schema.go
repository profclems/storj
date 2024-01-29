//lint:file-ignore * generated file
// AUTOGENERATED BY storj.io/storj/storagenode/storagenodedb/schemagen.go
// DO NOT EDIT

package storagenodedb

import "storj.io/common/dbutil/dbschema"

func Schema() map[string]*dbschema.Schema {
	return map[string]*dbschema.Schema{
		"bandwidth": {
			Tables: []*dbschema.Table{
				{
					Name: "bandwidth_usage",
					Columns: []*dbschema.Column{
						{
							Name:       "action",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "amount",
							Type:       "BIGINT",
							IsNullable: false,
						},
						{
							Name:       "created_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
				{
					Name:       "bandwidth_usage_rollups",
					PrimaryKey: []string{"action", "interval_start", "satellite_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "action",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "amount",
							Type:       "BIGINT",
							IsNullable: false,
						},
						{
							Name:       "interval_start",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				{Name: "idx_bandwidth_usage_created", Table: "bandwidth_usage", Columns: []string{"created_at"}, Unique: false, Partial: ""},
				{Name: "idx_bandwidth_usage_satellite", Table: "bandwidth_usage", Columns: []string{"satellite_id"}, Unique: false, Partial: ""},
			},
		},
		"heldamount": {
			Tables: []*dbschema.Table{
				{
					Name:       "payments",
					PrimaryKey: []string{"id"},
					Columns: []*dbschema.Column{
						{
							Name:       "amount",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "created_at",
							Type:       "timestamp",
							IsNullable: false,
						},
						{
							Name:       "id",
							Type:       "bigserial",
							IsNullable: false,
						},
						{
							Name:       "notes",
							Type:       "TEXT",
							IsNullable: true,
						},
						{
							Name:       "period",
							Type:       "TEXT",
							IsNullable: true,
						},
						{
							Name:       "receipt",
							Type:       "TEXT",
							IsNullable: true,
						},
						{
							Name:       "satellite_id",
							Type:       "bytea",
							IsNullable: false,
						},
					},
				},
				{
					Name:       "paystubs",
					PrimaryKey: []string{"period", "satellite_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "codes",
							Type:       "TEXT",
							IsNullable: false,
						},
						{
							Name:       "comp_at_rest",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "comp_get",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "comp_get_audit",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "comp_get_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "comp_put",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "comp_put_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "created_at",
							Type:       "timestamp",
							IsNullable: false,
						},
						{
							Name:       "disposed",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "distributed",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "held",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "owed",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "paid",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "period",
							Type:       "TEXT",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "bytea",
							IsNullable: false,
						},
						{
							Name:       "surge_percent",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "usage_at_rest",
							Type:       "double precision",
							IsNullable: false,
						},
						{
							Name:       "usage_get",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "usage_get_audit",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "usage_get_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "usage_put",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "usage_put_repair",
							Type:       "bigint",
							IsNullable: false,
						},
					},
				},
			},
		},
		"info": {},
		"notifications": {
			Tables: []*dbschema.Table{
				{
					Name:       "notifications",
					PrimaryKey: []string{"id"},
					Columns: []*dbschema.Column{
						{
							Name:       "created_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "message",
							Type:       "TEXT",
							IsNullable: false,
						},
						{
							Name:       "read_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "sender_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "title",
							Type:       "TEXT",
							IsNullable: false,
						},
						{
							Name:       "type",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
		},
		"orders": {
			Tables: []*dbschema.Table{
				{
					Name: "order_archive_",
					Columns: []*dbschema.Column{
						{
							Name:       "archived_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "order_limit_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "order_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "serial_number",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "status",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
					},
				},
				{
					Name: "unsent_order",
					Columns: []*dbschema.Column{
						{
							Name:       "order_limit_expiration",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "order_limit_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "order_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "serial_number",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				{Name: "idx_order_archived_at", Table: "order_archive_", Columns: []string{"archived_at"}, Unique: false, Partial: ""},
				{Name: "idx_orders", Table: "unsent_order", Columns: []string{"satellite_id", "serial_number"}, Unique: true, Partial: ""},
			},
		},
		"piece_expiration": {
			Tables: []*dbschema.Table{
				{
					Name:       "piece_expirations",
					PrimaryKey: []string{"piece_id", "satellite_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "deletion_failed_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "piece_expiration",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "piece_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "trash",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				{Name: "idx_piece_expirations_deletion_failed_at", Table: "piece_expirations", Columns: []string{"deletion_failed_at"}, Unique: false, Partial: ""},
				{Name: "idx_piece_expirations_piece_expiration", Table: "piece_expirations", Columns: []string{"piece_expiration"}, Unique: false, Partial: ""},
				{Name: "idx_piece_expirations_trashed", Table: "piece_expirations", Columns: []string{"satellite_id", "trash"}, Unique: false, Partial: "trash = 1"},
			},
		},
		"piece_spaced_used": {
			Tables: []*dbschema.Table{
				{
					Name: "piece_space_used",
					Columns: []*dbschema.Column{
						{
							Name:       "content_size",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: true,
						},
						{
							Name:       "total",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				{Name: "idx_piece_space_used_satellite_id", Table: "piece_space_used", Columns: []string{"satellite_id"}, Unique: true, Partial: ""},
			},
		},
		"pieceinfo": {
			Tables: []*dbschema.Table{
				{
					Name: "pieceinfo_",
					Columns: []*dbschema.Column{
						{
							Name:       "deletion_failed_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "order_limit",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "piece_creation",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "piece_expiration",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "piece_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "piece_size",
							Type:       "BIGINT",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
						{
							Name:       "uplink_piece_hash",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				{Name: "idx_pieceinfo__expiration", Table: "pieceinfo_", Columns: []string{"piece_expiration"}, Unique: false, Partial: "piece_expiration IS NOT NULL"},
				{Name: "pk_pieceinfo_", Table: "pieceinfo_", Columns: []string{"satellite_id", "piece_id"}, Unique: true, Partial: ""},
			},
		},
		"pricing": {
			Tables: []*dbschema.Table{
				{
					Name:       "pricing",
					PrimaryKey: []string{"satellite_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "audit_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "disk_space_price",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "egress_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "repair_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
		},
		"reputation": {
			Tables: []*dbschema.Table{
				{
					Name:       "reputation",
					PrimaryKey: []string{"satellite_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "audit_history",
							Type:       "BLOB",
							IsNullable: true,
						},
						{
							Name:       "audit_reputation_alpha",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "audit_reputation_beta",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "audit_reputation_score",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "audit_success_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "audit_total_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "audit_unknown_reputation_alpha",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "audit_unknown_reputation_beta",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "audit_unknown_reputation_score",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "disqualified_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "joined_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "offline_suspended_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "offline_under_review_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "online_score",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "suspended_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "updated_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "vetted_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
					},
				},
			},
		},
		"satellites": {
			Tables: []*dbschema.Table{
				{
					Name: "satellite_exit_progress",
					Columns: []*dbschema.Column{
						{
							Name:       "bytes_deleted",
							Type:       "INTEGER",
							IsNullable: false,
						},
						{
							Name:       "completion_receipt",
							Type:       "BLOB",
							IsNullable: true,
						},
						{
							Name:       "finished_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "initiated_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "satellites", Column: "node_id", OnDelete: "", OnUpdate: ""},
						},
						{
							Name:       "starting_disk_usage",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
				{
					Name:       "satellites",
					PrimaryKey: []string{"node_id"},
					Columns: []*dbschema.Column{
						{
							Name:       "added_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "address",
							Type:       "TEXT",
							IsNullable: true,
						},
						{
							Name:       "node_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "status",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
		},
		"secret": {
			Tables: []*dbschema.Table{
				{
					Name:       "secret",
					PrimaryKey: []string{"token"},
					Columns: []*dbschema.Column{
						{
							Name:       "created_at",
							Type:       "timestamp with time zone",
							IsNullable: false,
						},
						{
							Name:       "token",
							Type:       "bytea",
							IsNullable: false,
						},
					},
				},
			},
		},
		"storage_usage": {
			Tables: []*dbschema.Table{
				{
					Name:       "storage_usage",
					PrimaryKey: []string{"satellite_id", "timestamp"},
					Columns: []*dbschema.Column{
						{
							Name:       "at_rest_total",
							Type:       "REAL",
							IsNullable: false,
						},
						{
							Name:       "interval_end_time",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						{
							Name:       "timestamp",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
					},
				},
			},
		},
		"used_serial": {},
	}
}
