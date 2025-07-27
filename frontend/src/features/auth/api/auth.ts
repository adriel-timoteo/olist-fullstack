import { fetchClient } from "../../../api/client";
import type { LoginRequest, LoginResponse } from "../types";


export async function loginApi(req: LoginRequest): Promise<LoginResponse> {
  return fetchClient<LoginResponse>("/login", {
    method: "POST",
    body: JSON.stringify(req),
  });
}