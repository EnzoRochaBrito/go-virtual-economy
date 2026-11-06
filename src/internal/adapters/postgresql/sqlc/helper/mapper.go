package helper

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Int4(v int) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(v),
		Valid: true,
	}
}

func Date(time time.Time) pgtype.Date {
	return pgtype.Date{
		Time: time,
		InfinityModifier: pgtype.Finite,
		Valid: true,
	}
}