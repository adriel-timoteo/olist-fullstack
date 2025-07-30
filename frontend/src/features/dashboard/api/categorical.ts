import { fetchClient } from "../../../api/client";
import type { TopCategoriesResponse, TopCitiesResponse } from "../types";

export async function topCitiesApi(
  limit: string,
): Promise<TopCitiesResponse> {
  const query = new URLSearchParams({ limit }).toString();
  console.log(query)
  return fetchClient<TopCitiesResponse>(`/customer/top-cities?${query}`, {
    method: "GET",
  });
}

export async function topCategoriesApi(
  limit: string,
): Promise<TopCategoriesResponse> {
  const query = new URLSearchParams({ limit }).toString();
  console.log(query)
  return fetchClient<TopCategoriesResponse>(`/customer/top-categories?${query}`, {
    method: "GET",
  });
}