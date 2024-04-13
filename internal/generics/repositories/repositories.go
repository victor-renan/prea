package repositories

import (
	"fmt"
	"log"
	"prea/internal/common"
	"reflect"
	"strings"
	"time"
)

const (
	UUID      = "uuid"
	SERIAL    = "serial"
	VARCHAR   = "varchar"
	TEXT      = "text"
	INT       = "int"
	FLOAT     = "float"
	DECIMAL   = "decimal"
	NULL      = "null"
	DATE      = "date"
	TIMESTAMP = "timestamp"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("repositories")
}

type IModelInjectable interface {
	Table() string
	Pk() string
}

type IGenericRepository[T IModelInjectable] interface {
	GetAll() ([]T, error)
	GetById(id string) (T, error)
	Create(data T) (T, error)
	Update(id string, partial T) (T, error)
	Delete(id string) error
	GetLast() (T, error)
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

func ModelToKV(data IModelInjectable) (keys []string, vals []string) {
	r := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := range r.NumField() {
		if r.Field(i).Name != data.Pk() && !v.Field(i).IsZero() {
			dbField := r.Field(i).Tag.Get("db")
			dbType := r.Field(i).Tag.Get("db_type")

			var parsedKey string

			parsedKey = `"` + strings.ToLower(r.Field(i).Name) + `"`

			if dbField != "" {
				parsedKey = `"` + dbField + `"`
			}

			keys = append(keys, parsedKey)

			var parsedValue string

			switch dbType {
			case UUID, VARCHAR, TEXT:
				parsedValue = "'" + fmt.Sprint(v.Field(i).Interface()) + "'"
			case DATE, TIMESTAMP:
				format := time.DateTime
				if dbType == DATE {
					format = time.DateOnly
				}
				parsedValue = "'" + v.Field(i).Interface().(time.Time).Format(format) + "'"
			default:
				parsedValue = fmt.Sprint(v.Field(i).Interface())
			}

			vals = append(vals, parsedValue)
		}
	}

	return
}

func ModelToInsert[T IModelInjectable](data T) (string, string) {
	ks, vs := ModelToKV(data)

	table := strings.Join(ks, ",")
	values := strings.Join(vs, ",")

	fmt.Print(values)

	return "(" + table + ")", "(" + values + ")"
}

func ModelToUpdate[T IModelInjectable](data T) string {
	ks, vs := ModelToKV(data)

	var final string

	for i := range ks {
		final += ks[i] + "=" + vs[i] + ","
	}

	return strings.TrimSuffix(final, ",")
}
