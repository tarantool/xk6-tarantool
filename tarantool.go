package tarantool

import (
	"github.com/tarantool/go-tarantool"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/tarantool", new(Tarantool))
}

var (
	chCallFutures = make(chan *tarantool.Future, 4096)
)

// Tarantool is the k6 Tarantool extension
type Tarantool struct{}

func (Tarantool) ResolveCallFutures() {
	go func() {
		for fut := range chCallFutures {
			if _, err := fut.Get(); err != nil {
				panic(err)
			}
		}
	}()
}

// Connect creates a new Tarantool connection
func (Tarantool) Connect(addr string, opts tarantool.Opts) (*tarantool.Connection, error) {
	if addr == "" {
		addr = "localhost:3301"
	}
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Select performs select to box.space
func (Tarantool) Select(conn *tarantool.Connection, space, index interface{}, offset, limit, iterator uint32, key interface{}) (*tarantool.Response, error) {
	resp, err := conn.Select(space, index, offset, limit, iterator, key)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Insert performs insertion to box.space
func (Tarantool) Insert(conn *tarantool.Connection, space, data interface{}) (*tarantool.Response, error) {
	resp, err := conn.Insert(space, data)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Replace performs "insert or replace" action to box.space
func (Tarantool) Replace(conn *tarantool.Connection, space, data interface{}) (*tarantool.Response, error) {
	resp, err := conn.Replace(space, data)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Delete performs deletion of a tuple by key
func (Tarantool) Delete(conn *tarantool.Connection, space, index, key interface{}) (*tarantool.Response, error) {
	resp, err := conn.Delete(space, index, key)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Update performs update of a tuple by key
func (Tarantool) Update(conn *tarantool.Connection, space, index, key, ops interface{}) (*tarantool.Response, error) {
	resp, err := conn.Update(space, index, key, ops)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Upsert performs "update or insert" action of a tuple by key
func (Tarantool) Upsert(conn *tarantool.Connection, space, tuple, ops interface{}) (*tarantool.Response, error) {
	resp, err := conn.Upsert(space, tuple, ops)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Call calls registered tarantool function
func (Tarantool) Call(conn *tarantool.Connection, fnName string, args interface{}) (*tarantool.Response, error) {
	resp, err := conn.Call(fnName, args)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (Tarantool) CallAsyncNoReturn(conn *tarantool.Connection, fnName string, args interface{}) {
	chCallFutures <- conn.CallAsync(fnName, args)
}

// Call17 calls registered tarantool function
func (Tarantool) Call17(conn *tarantool.Connection, fnName string, args interface{}) (*tarantool.Response, error) {
	resp, err := conn.Call17(fnName, args)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Eval passes lua expression for evaluation
func (Tarantool) Eval(conn *tarantool.Connection, expr string, args interface{}) (*tarantool.Response, error) {
	resp, err := conn.Eval(expr, args)
	if err != nil {
		return nil, err
	}
	return resp, err
}
