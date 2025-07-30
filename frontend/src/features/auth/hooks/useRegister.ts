import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { useAppNotification } from "../../../hooks/useAppNotification";
import type { ApiError } from "../../../types/api";
import { registerApi } from "../api/auth";

export function useRegister() {
  const navigate = useNavigate()
  const { notify, contextHolder } = useAppNotification();

  return {
    ...useMutation({
      mutationFn: registerApi,
      onSuccess: () => {
        notify("success", "Register Successful", "Please login using your credentials.")
        navigate("/login");
      },
      onError: (error: ApiError) => {
        let userMessage = "Something went wrong. Please try again.";

        // Map backend error codes to user-friendly messages
        switch (error.code) {
          case "VALIDATION_ERROR":
            userMessage = "Invalid email or password. Please try again with a different one.";
            break;
          case "CONFLICT":
            userMessage = "Email already exists. Please try a different email or login";
            break;
          case "TIMEOUT_ERROR":
            userMessage = "The server took too long to respond. Please try again.";
            break;
          default:
            userMessage = error.message || userMessage;
        }

        notify("error", "Register Failed", userMessage)
      },
    }),
    contextHolder,
  };
}
