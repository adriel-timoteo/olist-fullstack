export type ApiResponse<T> = {
  success: boolean;
  message: string;
  error?: BackendError;
  data: T;
}

export type BasicResponse<> = {
  success: boolean;
  message: string;
  error?: BackendError;
}

export interface BackendError {
  code: string;
  message: string;
  fields?: FieldError[];
}

export interface FieldError {
  field: string;
  message: string;
}

export class ApiError extends Error {
  code: string;
  details?: FieldError[];

  constructor(message: string, code : string, details?: FieldError[]) {
    super(message);
    this.name = "ApiError";
    this.code = code;
    this.details = details;
  }
}