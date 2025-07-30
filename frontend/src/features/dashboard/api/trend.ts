import { fetchClient } from "../../../api/client";
import type { DeliveredTrendResponse } from "../types";

export async function deliveredTrendApi(
  interval: string,
  start: string,
  end: string,
): Promise<DeliveredTrendResponse> {
  const query = new URLSearchParams({ start, end, interval }).toString();
  console.log(query)
  return fetchClient<DeliveredTrendResponse>(`/products/trends/delivered?${query}`, {
    method: "GET",
  });
}