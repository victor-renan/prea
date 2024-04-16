package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBGeneric[T IModelInjectable] struct {
	Conn  *pgxpool.Pool
	Ctx   context.Context
	model T
}

func (dbg DBGeneric[T]) GetAll() ([]T, error) {
	sql := "select * from " + dbg.model.Table() + " order by id"
	rows, _ := dbg.Conn.Query(dbg.Ctx, sql)

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])

	return items, err
}

func (dbg DBGeneric[T]) Where(data any) ([]T, error) {
	ks, vs := ModelToWhere[T](data)

	sql := "select * from " + dbg.model.Table() + " where " + ks + " order by id"
	rows, _ := dbg.Conn.Query(dbg.Ctx, sql, vs...)

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])

	return items, err
}

func (dbg DBGeneric[T]) GetById(id string) (T, error) {
	sql := "select * from " + dbg.model.Table() + " where id=$1"
	rows, _ := dbg.Conn.Query(dbg.Ctx, sql, id)

	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return item, err
}

func (dbg DBGeneric[T]) GetFirst(data any) (T, error) {
	ks, vs := ModelToWhere[T](data)

	sql := "select * from " + dbg.model.Table() + " where " + ks + " order by id limit 1"
	rows, _ := dbg.Conn.Query(dbg.Ctx, sql, vs...)

	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return item, err
}

func (dbg DBGeneric[T]) Create(data any) (T, error) {
	k, v, vals := ModelToInsert[T](data)

	sql := "insert into " + dbg.model.Table() + k + " values " + v + "returning *"

	rows, err := dbg.Conn.Query(dbg.Ctx, sql, vals...)

	if err != nil {
		return dbg.model, err
	}

	last, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return last, err
}

func (dbg DBGeneric[T]) Update(id string, partial any) (T, error) {
	ks, vals := ModelToWhere[T](partial)

	sql := "update " + dbg.model.Table() + " set " + ks + " where id=$" + fmt.Sprint(1+len(vals)) + " returning *"

	rows, err := dbg.Conn.Query(dbg.Ctx, sql, append(vals, id)...)

	if err != nil {
		return dbg.model, err
	}

	item, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])

	return item, err
}

func (dbg DBGeneric[T]) Delete(id string) error {
	_, err := dbg.GetById(id)

	if err != nil {
		return err
	}

	sql := "delete from " + dbg.model.Table() + " where id=$1"
	_, err = dbg.Conn.Exec(dbg.Ctx, sql, id)

	return err
}
