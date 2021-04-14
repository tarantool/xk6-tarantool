import tarantool from "k6/x/tarantool";

const conn = tarantool.connect("localhost:3301");

export const setup = () => {
  tarantool.insert(conn, "cars", [1, "cadillac"]);
};

export let options = {
  scenarios: {
    constant_request_rate: {
      executor: "constant-arrival-rate",
      rate: 200000,
      timeUnit: "1s",
      duration: "30s",
      preAllocatedVUs: 100,
      maxVUs: 1000,
    },
  },
};

export default () => {
  tarantool.call(conn, "box.space.cars:select", [1]);
};

export const teardown = () => {
  tarantool.delete(conn, "cars", "pk", [1]);
};
