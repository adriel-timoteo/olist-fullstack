import { fetchClient } from "../../../api/client";
import type { RateResponse } from "../types";

export async function repeatPurchaseRateApi(): Promise<RateResponse> {
  return fetchClient<RateResponse>("/customer/repeat-rate", {
    method: "GET",
  });
}