import { Button, Card, Form, Input, Typography } from "antd";
import { LockOutlined, UserOutlined } from "@ant-design/icons";

const { Title } = Typography;

const onFinish = (values: unknown) => {
  console.log("Login values:", values);
};

const LoginForm = () => (
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
        name="username"
        rules={[{ required: true, message: "Please input your username!" }]}
      >
        <Input prefix={<UserOutlined />} placeholder="Username" size="large" />
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
          Log In
        </Button>
      </Form.Item>
    </Form>
  </Card>
);

export default LoginForm;
