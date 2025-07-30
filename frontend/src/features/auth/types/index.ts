import type { ApiResponse } from "../../../types/api";

export type AuthRequest = {
  email: string;
  password: string;
}


export type LoginResponse = ApiResponse<{token: string}>

export type RegisterResponse = ApiResponse<{user_id: number, email: string}>