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
