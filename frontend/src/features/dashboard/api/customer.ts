import { fetchClient } from "../../../api/client";
import type { CountResponse, RateResponse, TopCitiesResponse } from "../types";

export async function topCitiesApi(
  limit: string,
): Promise<TopCitiesResponse> {
  const query = new URLSearchParams({ limit }).toString();
  console.log(query)
  return fetchClient<TopCitiesResponse>(`/customer/top-cities?${query}`, {
    method: "GET",
  });
}


export async function totalUniqueCustApi(): Promise<CountResponse> {
  return fetchClient<CountResponse>("/customer/total", {
    method: "GET",
  });
}

export async function repeatPurchaseRateApi(): Promise<RateResponse> {
  return fetchClient<RateResponse>("/customer/repeat-rate", {
    method: "GET",
  });
}
