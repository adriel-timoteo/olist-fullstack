import { fetchClient } from "../../../api/client";
import type { AuthRequest, LoginResponse, RegisterResponse } from "../types";


export async function loginApi(req: AuthRequest): Promise<LoginResponse> {
  return fetchClient<LoginResponse>("/login", {
    method: "POST",
    body: JSON.stringify(req),
  });
}

export async function registerApi(req: AuthRequest): Promise<RegisterResponse> {
  return fetchClient<RegisterResponse>("/register", {
    method: "POST",
    body: JSON.stringify(req),
  });
}