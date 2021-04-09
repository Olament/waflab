import * as Setting from "../Setting";

export function getResult(testsetId, testcaseId, type) {
  return fetch(`${Setting.ServerUrl}/api/get-docker-health`, {
    method: "GET",
    credentials: "include"
  }).then(res => {
    if (!res.json()["health"]) {
      throw new Error("Backend docker is not healthy")
    }
  })
    .then(() => fetch(`${Setting.ServerUrl}/api/get-result?testsetId=${testsetId}&testcaseId=${testcaseId}&type=${type}`, {
      method: "GET",
      credentials: "include"
    })).then(res => res.json())
}
