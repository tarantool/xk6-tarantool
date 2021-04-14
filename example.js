import tarantool from "k6/x/tarantool";

const conn = tarantool.connect("localhost:3301");

export const setup = () => {
  tarantool.insert(conn, "cars", [1, "cadillac"]);
};

export default () => {
  console.log(tarantool.select(conn, "cars", "pk", 0, 1, 0, [1]));
};

export const teardown = () => {
  tarantool.delete(conn, "pk", [1]);
};
