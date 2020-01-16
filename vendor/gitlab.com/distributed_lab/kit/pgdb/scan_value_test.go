package pgdb_test

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/stretchr/testify/require"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type JSONColumn struct {
	I int
}

func (s *JSONColumn) Scan(src interface{}) error {
	return pgdb.JSONScan(src, s)
}

func (s JSONColumn) Value() (driver.Value, error) {
	return pgdb.JSONValue(s)
}

func TestJSONScanValue(t *testing.T) {
	db := getDB(t)

	cases := []struct {
		tpe string
	}{
		{"jsonb"},
		{"json"},
		{"text"},
		{"bytea"},
	}

	for _, tc := range cases {
		t.Run(tc.tpe, func(t *testing.T) {
			require.NoError(t, db.ExecRaw(fmt.Sprintf(`create temporary table t (c %s)`, tc.tpe)))

			t.Run("scan non-null to value", func(t *testing.T) {
				require.NoError(t, db.ExecRaw(`insert into t (c) values ('{"I": 1}')`))
				var got struct {
					C JSONColumn
				}
				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.Equal(t, 1, got.C.I)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			t.Run("scan non-null to pointer", func(t *testing.T) {
				require.NoError(t, db.ExecRaw(`insert into t (c) values ('{"I": 1}')`))
				var got struct {
					C *JSONColumn
				}
				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.NotNil(t, got.C)
				require.Equal(t, 1, got.C.I)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			// DISCUSS: actually I'm not totally sure what should be behaviour here.
			t.Run("scan null to value", func(t *testing.T) {
				require.NoError(t, db.ExecRaw(`insert into t (c) values (null)`))
				var got struct {
					C JSONColumn
				}
				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.Zero(t, got.C)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			t.Run("scan null to pointer", func(t *testing.T) {
				require.NoError(t, db.ExecRaw(`insert into t (c) values (null)`))
				var got struct {
					C *JSONColumn
				}
				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.Nil(t, got.C)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			t.Run("value value", func(t *testing.T) {
				expected := JSONColumn{
					I: 1337,
				}

				require.NoError(t, db.Exec(squirrel.Insert("t").SetMap(map[string]interface{}{
					"c": expected,
				})))

				var got struct {
					C JSONColumn
				}

				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.Equal(t, expected, got.C)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			t.Run("value nil", func(t *testing.T) {
				require.NoError(t, db.Exec(squirrel.Insert("t").SetMap(map[string]interface{}{
					"c": (*JSONColumn)(nil),
				})))

				var got struct {
					C *JSONColumn
				}

				require.NoError(t, db.GetRaw(&got, `select * from t`))
				require.Nil(t, got.C)
				require.NoError(t, db.ExecRaw(`truncate table t`))
			})

			// TODO: add invalid type tests

			require.NoError(t, db.ExecRaw(`drop table t`))
		})
	}

}
