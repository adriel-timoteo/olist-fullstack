import { ApiError, type ApiResponse } from "../types/api";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function fetchClient<T>(
  input: string,
  init?: RequestInit
): Promise<T> {
  const token = localStorage.getItem("token");
  const url = input.startsWith("http") ? input : `${BASE_URL}${input}`;

  const response = await fetch(url, {
    ...init,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(init?.headers || {}),
    },
  });

  const json: ApiResponse<T> = await response
    .json()
    .catch(() => ({} as ApiResponse<T>));

  if (!response.ok || json.success === false) {
    const backendError = json.error;
    throw new ApiError(
      backendError?.message || "API request failed",
      backendError?.code || String(response.status),
      backendError?.fields
    );
  }

  return json as unknown as T;
}
