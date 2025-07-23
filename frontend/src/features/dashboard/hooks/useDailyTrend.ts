// src/hooks/useDailyTrend.ts
import { useQuery } from "@tanstack/react-query";
import { getDailyTrend, type DailyTrendFilter } from "../api/dailyTrend";

export function useDailyTrend(filter?: DailyTrendFilter) {
  return useQuery({
    queryKey: ["dailyTrend", filter],
    queryFn: () => getDailyTrend(filter),
    enabled: true,
  });
}
