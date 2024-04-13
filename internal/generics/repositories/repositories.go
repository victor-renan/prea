package repositories

import (
	"fmt"
	"log"
	"prea/internal/common"
	"reflect"
	"strings"
)

const (
	PKName = "Id"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("repositories")
}

type IModelInjectable interface {
	Table() string
}

type IGenericRepository[T IModelInjectable] interface {
	GetAll() ([]T, error)
	GetById(id string) (T, error)
	Create(data any) (T, error)
	Update(id string, partial any) (T, error)
	Delete(id string) error
}

func ModelToDest[T any](data *T) []any {
	addrs := reflect.ValueOf(data)

	var vals []any

	elem := addrs.Elem()
	for i := range elem.NumField() {
		vals = append(vals, elem.Field(i).Addr().Interface())
	}

	return vals
}

func ModelToKV[T IModelInjectable](data any) (keys []string, vals []any) {
	r := reflect.TypeFor[T]()
	v := reflect.ValueOf(data)

	for i := range v.NumField() {
		part, find := r.FieldByName(v.Type().Field(i).Name);
		if find {
			if part.Type == v.Field(i).Type() && part.Name != PKName && !v.Field(i).IsZero() {
				dbField := r.Field(i).Tag.Get("db")

				var parsedKey string

				parsedKey = `"` + strings.ToLower(part.Name) + `"`

				if dbField != "" {
					parsedKey = `"` + dbField + `"`
				}

				keys = append(keys, parsedKey)
				vals = append(vals, v.Field(i).Interface())
			}
		}
	}

	return
}

func ModelToInsert[T IModelInjectable](data any) (string, string, []any) {
	ks, vs := ModelToKV[T](data)

	table := strings.Join(ks, ",")
	vals := ""

	for i := range ks {
		vals += fmt.Sprintf("$%v", i+1) + ","
	}

	return "(" + table + ")", "(" + strings.TrimSuffix(vals, ",") + ")", vs
}

func ModelToUpdate[T IModelInjectable](data any) (string, []any) {
	ks, vs := ModelToKV[T](data)

	var final string

	for i := range ks {
		final += ks[i] + "=$" + fmt.Sprint(i+1) + ","
	}

	return strings.TrimSuffix(final, ","), vs
}
