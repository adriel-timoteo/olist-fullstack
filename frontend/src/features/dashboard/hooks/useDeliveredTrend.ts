import { useQuery } from "@tanstack/react-query";
import type { DeliveredTrendResponse } from "../types";
import { deliveredTrendApi } from "../api/trend";

export function useDeliveredTrend(start: string, end: string, interval: string) {
  return useQuery<DeliveredTrendResponse, Error>({
    queryKey: ["deliveredTrend", start, end],
    queryFn: () => deliveredTrendApi(start, end, interval),
    enabled: !!start && !!end, // Only run when both dates are provided
  });
}
