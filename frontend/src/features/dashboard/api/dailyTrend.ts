import { fetchClient } from "../../../api/client";

export interface DailyTrend {
  date: string;
  value: number;
}

export interface DailyTrendFilter {
  startDate?: string;
  endDate?: string;
}

export function getDailyTrend(filter?: DailyTrendFilter): Promise<DailyTrend[]> {
  const params = new URLSearchParams();
  if (filter?.startDate) params.append("startDate", filter.startDate);
  if (filter?.endDate) params.append("endDate", filter.endDate);

  const url = `/api/daily-trend${params.toString() ? `?${params.toString()}` : ""}`;
  return fetchClient<DailyTrend[]>(url);
}
