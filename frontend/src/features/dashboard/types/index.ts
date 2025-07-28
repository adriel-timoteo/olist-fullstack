import type { ApiResponse } from "../../../types/api";

export type DeliveredTrendResponse = ApiResponse<{time: Date, count: number}[]>
