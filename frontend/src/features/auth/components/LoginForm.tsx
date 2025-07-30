import { LockOutlined, MailOutlined } from "@ant-design/icons";
import { Button, Card, Form, Input, Spin, Typography } from "antd";
import { useLogin } from "../hooks/useLogin";
import type { AuthRequest } from "../types";
import { Link } from "react-router";

const { Title, Text } = Typography;

const LoginForm = () => {
  const { mutate: login, isPending, contextHolder } = useLogin();

  const onFinish = (values: AuthRequest) => {
    login(values);
    console.log("Login values:", values);
  };

  return (
    <>
      {contextHolder}
      <Card style={{ maxWidth: 400, width: "100%" }}>
        <Title level={2} className="text-center mb-4">
          Login
        </Title>
        <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          layout="vertical"
        >
          <Form.Item
            name="email"
            rules={[{ required: true, message: "Please input your email!" }]}
          >
            <Input prefix={<MailOutlined />} placeholder="Email" size="large" />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[{ required: true, message: "Please input your password!" }]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Password"
              size="large"
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" block size="large">
              {isPending ? <Spin /> : "Log In"}
            </Button>
          </Form.Item>

          <Text type="secondary" className="text-center block">
            Don’t have an account? <Link to="/register">Register here</Link>
          </Text>
        </Form>
      </Card>
    </>
  );
};

export default LoginForm;
