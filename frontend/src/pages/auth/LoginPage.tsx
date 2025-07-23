import { Col, Row } from "antd";
import LoginForm from "../../features/auth/components/LoginForm";

const LoginPage = () => {
  return (
    <Row className="min-h-screen">
      <Col
        xs={0}
        md={12}
        style={{
          backgroundColor: "#f5f5f5",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          padding: 0,
        }}
      >
        <img
          src="https://images.pexels.com/photos/5696170/pexels-photo-5696170.jpeg"
          alt="Login"
          style={{
            width: "100%",
            height: "100vh",
            objectFit: "cover",
          }}
        />
      </Col>

      <Col
        xs={24}
        md={12}
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          padding: "2rem",
        }}
      >
        <LoginForm />
      </Col>
    </Row>
  );
};

export default LoginPage;
