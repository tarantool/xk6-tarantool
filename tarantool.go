package tarantool

import (
	"fmt"

	"github.com/loadimpact/k6/js/modules"
	"github.com/tarantool/go-tarantool"
)

func init() {
	modules.Register("k6/x/tarantool", new(TARANTOOL))
}

// TARANTOOL is the k6 Tarantool extension
type TARANTOOL struct{}

// Connect creates a new Tarantool connection
func Connect(addr string, opts tarantool.Opts) *tarantool.Connection {
	if addr == "" {
		addr = "localhost:3301"
	}
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		return nil
	}
	return conn
}

// Select performs select to box.space
func (TARANTOOL) Select(conn *tarantool.Connection, space, index interface{}, offset, limit, iterator uint32, key interface{}) *tarantool.Response {
	resp, err := conn.Select(space, index, offset, limit, iterator, key)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Insert performs insertion to box.space
func (TARANTOOL) Insert(conn *tarantool.Connection, space, data interface{}) *tarantool.Response {
	resp, err := conn.Insert(space, data)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Replace performs "insert or replace" action to box.space
func (TARANTOOL) Replace(conn *tarantool.Connection, space, data interface{}) *tarantool.Response {
	resp, err := conn.Replace(space, data)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Delete performs deletion of a tuple by key
func (TARANTOOL) Delete(conn *tarantool.Connection, space, index, key interface{}) *tarantool.Response {
	resp, err := conn.Delete(space, index, key)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Update performs update of a tuple by key
func (TARANTOOL) Update(conn *tarantool.Connection, space, index, key, ops interface{}) *tarantool.Response {
	resp, err := conn.Update(space, index, key, ops)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Upsert performs "update or insert" action of a tuple by key
func (TARANTOOL) Upsert(conn *tarantool.Connection, space, tuple, ops interface{}) *tarantool.Response {
	resp, err := conn.Upsert(space, tuple, ops)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Call calls registered tarantool function
func (TARANTOOL) Call(conn *tarantool.Connection, fnName string, args interface{}) *tarantool.Response {
	resp, err := conn.Call(fnName, args)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// Eval passes lua expression for evaluation
func (TARANTOOL) Eval(conn *tarantool.Connection, expr string, args interface{}) *tarantool.Response {
	resp, err := conn.Eval(expr, args)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
