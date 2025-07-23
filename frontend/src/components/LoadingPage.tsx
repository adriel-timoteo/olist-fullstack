import { Spin } from "antd";

const LoadingPage = () => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
        background: "#f0f2f5",
        flexDirection: "column",
      }}
    >
      <Spin size="large" />
      <p style={{ marginTop: "16px", fontSize: "16px", color: "#555" }}>
        Loading, please wait...
      </p>
    </div>
  );
};

export default LoadingPage;
