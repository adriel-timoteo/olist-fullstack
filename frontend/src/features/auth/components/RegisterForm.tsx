import { LockOutlined, MailOutlined } from "@ant-design/icons";
import { Button, Card, Form, Input, Spin, Typography } from "antd";
import { useRegister } from "../hooks/useRegister"; // Adjust this hook as needed
import type { AuthRequest } from "../types";

const { Title } = Typography;

const RegisterForm = () => {
  const { mutate: register, isPending, contextHolder } = useRegister();

  const [form] = Form.useForm();

  const onFinish = (values: AuthRequest) => {
    register(values);
    console.log("Register values:", values);
  };

  return (
    <>
      {contextHolder}
      <Card style={{ maxWidth: 400, width: "100%" }}>
        <Title level={2} className="text-center mb-4">
          Register
        </Title>
        <Form form={form} name="register" onFinish={onFinish} layout="vertical">
          <Form.Item
            name="email"
            rules={[
              { required: true, message: "Please input your email!" },
              { type: "email", message: "Please enter a valid email!" },
            ]}
          >
            <Input prefix={<MailOutlined />} placeholder="Email" size="large" />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[
              { required: true, message: "Please input your password!" },
              { min: 6, message: "Password must be at least 6 characters" },
            ]}
            hasFeedback
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Password"
              size="large"
            />
          </Form.Item>

          <Form.Item
            name="confirmPassword"
            dependencies={["password"]}
            hasFeedback
            rules={[
              { required: true, message: "Please confirm your password!" },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue("password") === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("Passwords do not match!"));
                },
              }),
            ]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Confirm Password"
              size="large"
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" block size="large">
              {isPending ? <Spin /> : "Register"}
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </>
  );
};

export default RegisterForm;
