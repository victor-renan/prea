package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"prea/internal/common"
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
	rows, err := Conn.Query(Ctx, sql)

	var items []T

	for rows.Next() {
		var item T
		err := rows.Scan(ModelToDest(&item)...)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}

	return items, err
}

func (dbg DBGeneric[T]) GetById(id int64) (T, error) {
	sql := "select * from " + dbg.Model.Table() + " where id=$1"

	var item T
	err := Conn.QueryRow(Ctx, sql, id).Scan(ModelToDest(&item)...)

	return item, err
}

func (dbg DBGeneric[T]) Create(data T) (T, error) {
	keys, values := ModelToInsert(data)

	sql := "insert into " + dbg.Model.Table() + keys + " values " + values

	_, err := Conn.Exec(Ctx, sql)

	if err != nil {
		return dbg.Model, err
	}

	last, err := dbg.GetLast()

	return last, err
}

func (dbg DBGeneric[T]) Update(id int64, partial T) (T, error) {
	states := ModelToUpdate(partial)

	sql := "update " + dbg.Model.Table() + " set " + states + " where id=$1"
	_, err := Conn.Exec(Ctx, sql, id)

	if err != nil {
		return dbg.Model, err
	}

	updated, err := dbg.GetById(id)

	return updated, err
}

func (dbg DBGeneric[T]) Delete(id int64) error {
	_, err := dbg.GetById(id)

	if err != nil {
		return err
	}

	sql := "delete from " + dbg.Model.Table() + " where id=$1"
	_, err = Conn.Exec(Ctx, sql, id)

	return err
}

func (dbg DBGeneric[T]) GetLast() (T, error) {
	sql := "select * from " + dbg.Model.Table() + " order by id desc limit 1"
	rows, err := Conn.Query(Ctx, sql)

	var item T
	for rows.Next() {
		err := rows.Scan(ModelToDest(&item)...)

		if err != nil {
			return item, err
		}
	}

	return item, err
}
