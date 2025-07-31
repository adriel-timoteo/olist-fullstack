import { fetchClient } from "../../../api/client";
import type { TopCategoriesResponse } from "../types";

export async function topCategoriesApi(
  limit: string,
): Promise<TopCategoriesResponse> {
  const query = new URLSearchParams({ limit }).toString();
  console.log(query)
  return fetchClient<TopCategoriesResponse>(`/product/top-categories?${query}`, {
    method: "GET",
  });
}