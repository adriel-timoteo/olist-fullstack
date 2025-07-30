import { Card, Typography, Spin } from "antd";
import { useEffect, useState } from "react";

const { Title, Text } = Typography;

interface NumberDisplayProps {
  title: string;
  fetchData?: () => Promise<number>;
  prefix?: string;
  suffix?: string;
  precision?: number;
  color?: string;
  height?: number;
}

const NumberDisplay = ({
  title,
  fetchData,
  prefix = "",
  suffix = "",
  precision = 0,
  color = "#1677ff",
  height = 120,
}: NumberDisplayProps) => {
  const [displayValue, setDisplayValue] = useState(0);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (fetchData) {
      setLoading(true);
      fetchData()
        .then((res) => {
          setDisplayValue(res);
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [fetchData]);

  const formattedValue = displayValue.toFixed(precision);

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
      {loading ? (
        <Spin />
      ) : (
        <Title level={2} style={{ color, margin: 0 }}>
          {prefix}
          {formattedValue}
          {suffix}
        </Title>
      )}
    </Card>
  );
};

export default NumberDisplay;
