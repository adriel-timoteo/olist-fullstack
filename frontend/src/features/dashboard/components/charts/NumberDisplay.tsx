import { Card, Typography } from "antd";

const { Title, Text } = Typography;

interface NumberDisplayProps {
  title: string;
  value: number | string;
  prefix?: string;
  suffix?: string;
  precision?: number;
  color?: string;
  height?: number;
}

const NumberDisplay = ({
  title,
  value,
  prefix = "",
  suffix = "",
  precision = 0,
  color = "#1677ff",
  height = 120,
}: NumberDisplayProps) => {
  const formattedValue =
    typeof value === "number" ? value.toFixed(precision) : value;

  return (
    <Card
      style={{
        height,
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
      }}
    >
      <Text type="secondary">{title}</Text>
      <Title level={2} style={{ color, margin: 0 }}>
        {prefix}
        {formattedValue}
        {suffix}
      </Title>
    </Card>
  );
};

export default NumberDisplay;
