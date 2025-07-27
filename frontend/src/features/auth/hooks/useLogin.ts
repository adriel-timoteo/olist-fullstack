import { useMutation } from "@tanstack/react-query";
import { loginApi } from "../api/auth";
import { useNavigate } from "react-router";
import type { ApiError } from "../../../types/api";
import { useAppNotification } from "../../../hooks/useAppNotification";

export function useLogin() {
  const navigate = useNavigate()
  const { notify, contextHolder } = useAppNotification();

  return {
    ...useMutation({
      mutationFn: loginApi,
      onSuccess: (data) => {
        localStorage.setItem("token", data.data.token);
        navigate("/");
        console.log("Login successful:", data);
      },
      onError: (error: ApiError) => {
        let userMessage = "Something went wrong. Please try again.";

        // Map backend error codes to user-friendly messages
        switch (error.code) {
          case "VALIDATION_ERROR":
            userMessage = "Invalid email or password. Please try again.";
            break;
          case "TIMEOUT_ERROR":
            userMessage = "The server took too long to respond. Please try again.";
            break;
          default:
            userMessage = error.message || userMessage;
        }

        notify("error", "Login Failed", userMessage)

        console.error("Login failed:", error);
      },
    }),
    contextHolder,
  };
}
