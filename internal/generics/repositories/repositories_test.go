package repositories_test

import (
	"fmt"
	"prea/internal/generics/repositories"
	"testing"
	"time"
)

type ModelTest struct {
	Id       string    `db:"id" db_type:"uuid"`
	Name     string    `db:"name" db_type:"varchar"`
	Integ    int       `db:"integ" db_type:"int"`
	DateTime time.Time `db:"datetime" db_type:"timestamp"`
	Date     time.Time `db:"date" db_type:"date"`
}

func (ModelTest) Table() string {
	return "modeltest"
}

func (ModelTest) Pk() string {
	return "Id"
}

func compareDs(t *testing.T, ds []string, exp_ds []string) {
	act := fmt.Sprintf("Comparing...\na:%v\nb:%v\n", ds, exp_ds)

	if len(ds) != len(exp_ds) {
		t.Fatalf("%v expected %v, got %v", act, exp_ds, ds)
	}

	for i := range ds {
		if ds[i] != exp_ds[i] {
			t.Fatalf("%v expected %v, got %v", act, exp_ds[i], ds[i])
		}
	}
}

func TestModelToKV(t *testing.T) {
	datetime := time.Now()
	date := time.Now()

	model := ModelTest{
		Id:       "111-222-333",
		Name:     "Renan",
		Integ:    1,
		DateTime: datetime,
		Date:     date,
	}

	ks, vs := repositories.ModelToKV(model)

	exp_ks := []string{
		`"name"`,
		`"integ"`,
		`"datetime"`,
		`"date"`,
	}

	exp_vs := []string{
		"'Renan'",
		"1",
		fmt.Sprintf("'%v'", datetime.Format(time.DateTime)),
		fmt.Sprintf("'%v'", date.Format(time.DateOnly)),
	}

	compareDs(t, ks, exp_ks)
	compareDs(t, vs, exp_vs)
}

func TestModelToInsert(t *testing.T) {
	datetime := time.Now()
	date := time.Now()
	model := ModelTest{
		Name:     "Renan",
		Integ:    1,
		DateTime: datetime,
		Date:     date,
	}

	table, values := repositories.ModelToInsert(model)

	exp_table := `("name","integ","datetime","date")`
	exp_values := fmt.Sprintf(
		`('Renan',1,'%v','%v')`,
		datetime.Format(time.DateTime),
		datetime.Format(time.DateOnly),
	)

	if table != exp_table {
		t.Fatalf("table %v not equals %v", table, exp_table)
	}

	if values != exp_values {
		t.Fatalf("values %v not equals %v", values, exp_values)
	}
}

func TestModelToUpdate(t *testing.T) {
	datetime := time.Now()
	date := time.Now()
	model := ModelTest{
		Name:     "Renan",
		Integ:    1,
		DateTime: datetime,
		Date:     date,
	}

	test := repositories.ModelToUpdate(model)

	exp_test := fmt.Sprintf(
		`"name"='Renan',"integ"=1,"datetime"='%v',"date"='%v'`,
		datetime.Format(time.DateTime),
		date.Format(time.DateOnly),
	)

	if test != exp_test {
		t.Fatalf("table %v not equals %v", test, exp_test)
	}
}
