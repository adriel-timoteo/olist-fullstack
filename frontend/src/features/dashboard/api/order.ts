import { fetchClient } from "../../../api/client";
import type { CountResponse, DeliveredStatusResponse, DeliveredTrendResponse, OrderByHourResponse, RateResponse } from "../types";

export async function deliveredTrendApi(
  interval: string,
  start: string,
  end: string,
): Promise<DeliveredTrendResponse> {
  const query = new URLSearchParams({ start, end, interval }).toString();
  return fetchClient<DeliveredTrendResponse>(`/order/trends/delivered?${query}`, {
    method: "GET",
  });
}

export async function onTimeDeliveryRateApi(): Promise<RateResponse> {
  return fetchClient<RateResponse>("/order/delivery/ontime-rate", {
    method: "GET",
  });
}

export async function deliveredStatusApi(
  start: string,
  end: string,
): Promise<DeliveredStatusResponse> {
  const query = new URLSearchParams({ start, end }).toString();
  return fetchClient<DeliveredStatusResponse>(`/order/trends/status?${query}`, {
    method: "GET",
  });
}

export async function totalRevenue(): Promise<CountResponse> {
  return fetchClient<CountResponse>("/order/total-revenue", {
    method: "GET",
  });
}

export async function averageOrderValue(): Promise<CountResponse> {
  return fetchClient<CountResponse>("/order/aov", {
    method: "GET",
  });
}

export async function averageDeliveryTime(): Promise<CountResponse> {
  return fetchClient<CountResponse>("/order/delivery/speed", {
    method: "GET",
  });
}

export async function averageOrderByHour(): Promise<OrderByHourResponse> {
  return fetchClient<OrderByHourResponse>("/order/hourly", {
    method: "GET",
  });
}