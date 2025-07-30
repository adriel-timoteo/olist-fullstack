import { fetchClient } from "../../../api/client";
import type { CountResponse } from "../types";

export async function totalUniqueCustApi(): Promise<CountResponse> {
  return fetchClient<CountResponse>("/customer/total", {
    method: "GET",
  });
}