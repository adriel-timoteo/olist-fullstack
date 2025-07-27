import { notification, type NotificationArgsProps } from "antd";

type NotificationType = "success" | "info" | "warning" | "error";

export function useAppNotification() {
  const [api, contextHolder] = notification.useNotification();

  const notify = (
    type: NotificationType,
    title: string,
    message: string,
    config?: Partial<NotificationArgsProps>
  ) => {
    api[type]({
      message: title,
      description: message,
      showProgress: true,
      pauseOnHover: true,
      ...config, // Optional override
    });
  };

  return { notify, contextHolder };
}
