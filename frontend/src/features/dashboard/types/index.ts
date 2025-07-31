import type { ApiResponse } from "../../../types/api";

export type DeliveredTrendResponse = ApiResponse<{time: Date, count: number}[]>
export type DeliveredStatusResponse = ApiResponse<{time: Date, status: string, count: number}[]>

export type TopCitiesResponse = ApiResponse<{city: string, count: number}[]>
export type TopCategoriesResponse = ApiResponse<{category: string, count: number}[]>

export type CountResponse = ApiResponse<{count: number}>
export type RateResponse = ApiResponse<{rate: number}>