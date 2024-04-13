package repositories

import (
	"context"
	"fmt"
	"prea/internal/common"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	EnvConn string = "PGCONN"
)

var Ctx context.Context
var Conn *pgxpool.Pool

type DBGeneric[T IModelInjectable] struct {
	Model T
}

func init() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, common.GetEnv(EnvConn))

	common.ThrowException(err, GetLogger())

	Conn = pool
	Ctx = ctx
}

func (dbg DBGeneric[T]) GetAll() ([]T, error) {
	sql := "select * from " + dbg.Model.Table() + " order by id"
	rows, _ := Conn.Query(Ctx, sql)

	items, err := pgx.CollectRows(rows, pgx.RowToStructByPos[T])

	return items, err
}

func (dbg DBGeneric[T]) GetById(id string) (T, error) {
	sql := "select * from " + dbg.Model.Table() + " where id=$1"
	rows, _ := Conn.Query(Ctx, sql, id)

	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByPos[T])

	return item, err
}

func (dbg DBGeneric[T]) Create(data any) (T, error) {
	k, v, vals := ModelToInsert[T](data)

	sql := "insert into " + dbg.Model.Table() + k + " values " + v + "returning *"

	rows, err := Conn.Query(Ctx, sql, vals...)

	if err != nil {
		return dbg.Model, err
	}

	last, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return last, err
}

func (dbg DBGeneric[T]) Update(id string, partial any) (T, error) {
	ks, vals := ModelToUpdate[T](partial)

	sql := "update " + dbg.Model.Table() + " set " + ks + " where id=$" + fmt.Sprint(1 + len(vals)) + " returning *"

	rows, err := Conn.Query(Ctx, sql, append(vals, id)...)

	if err != nil {
		return dbg.Model, err
	}

	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return item, err
}

func (dbg DBGeneric[T]) Delete(id string) error {
	_, err := dbg.GetById(id)

	if err != nil {
		return err
	}

	sql := "delete from " + dbg.Model.Table() + " where id=$1"
	_, err = Conn.Exec(Ctx, sql, id)

	return err
}
