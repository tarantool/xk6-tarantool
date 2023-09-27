# xk6-tarantool

This is a [k6](https://github.com/loadimpact/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Clone `xk6`:
  ```shell
  git clone https://github.com/k6io/xk6.git
  cd xk6
  ```

2. Build the binary:
  ```shell
  CGO_ENABLED=1 go run ./cmd/xk6/main.go build master \
    --with github.com/tarantool/xk6-tarantool
  ```

## Example

Make sure you configured a space in Tarantool console first:

```lua
box.schema.space.create("cars")
box.space.cars:format({
    {name = 'id', type = 'unsigned'},
    {name = 'model', type = 'string'}
})
box.space.cars:create_index('pk', {
    type = 'hash',
    parts = {'id'}
})
```

Test script:

```javascript
import tarantool from "k6/x/tarantool";

const conn = tarantool.connect("localhost:3301");

export const setup = () => {
  tarantool.insert(conn, "cars", [1, "cadillac"]);
};

export default () => {
  console.log(tarantool.call(conn, "box.space.cars:select", [1]));
};

export const teardown = () => {
  tarantool.delete(conn, "cars", "pk", [1]);
};
```

Result output:

```
$ ./k6 run test.js 

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: test.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] <3 OK [[1 cadillac]]>                         source=console

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

     █ setup

     █ teardown

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=1.46ms min=1.39ms med=1.48ms max=1.51ms p(90)=1.5ms p(95)=1.51ms
     iterations...........: 1   62.668421/s
```
