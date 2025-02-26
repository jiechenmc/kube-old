import http from "k6/http";

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 1,
  // A string specifying the total duration of the test run.
  duration: "30s",
};

const baseUrl = "http://localhost";

export function setup() {
  // EMPTY
}

export default function (cookiesForURL) {
  http.get(`${baseUrl}/`);
}

export function teardown(data) {
  // 4. teardown code
}