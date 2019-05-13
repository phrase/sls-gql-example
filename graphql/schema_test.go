package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestSchema(t *testing.T) {
	schema, err := parseSchema()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("hello", func(t *testing.T) {
		res := schema.Exec(context.TODO(), "query { hello }", "", nil)
		b, err := json.Marshal(res)
		if err != nil {
			t.Fatal(err)
		}
		if has, want := string(b), `{"data":{"hello":"world"}}`; has != want {
			t.Errorf("has=%q want=%q", has, want)
		}
	})
	t.Run("createUser", func(t *testing.T) {
		variables := map[string]interface{}{
			"name": "Tobias Schwab",
			"age":  int32(39),
		}
		res := schema.Exec(context.TODO(), "mutation create($name:String!,$age:Int!) { createUser(name:$name,age:$age) }", "", variables)

		if len(res.Errors) > 0 {
			t.Fatalf("expected no errors, was %v", res.Errors)
		}
	})
}

func TestResolver(t *testing.T) {
	r := &resolver{}
	msg := "custom"
	res, err := r.Hello(context.TODO(), struct{ Msg *string }{Msg: &msg})
	if err != nil {
		t.Fatal(err)
	}
	if has, want := res, "custom"; has != want {
		t.Errorf("has=%q want=%q", has, want)
	}
}
