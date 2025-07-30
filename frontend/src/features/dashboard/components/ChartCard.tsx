import { Card, Typography } from "antd";
import { type ReactNode, useState } from "react";

const { Text } = Typography;

interface ChartCardProps<T> {
  title: string;
  children: (props: { filters: T }) => ReactNode;
  filters?: (filters: T, setFilters: (filters: T) => void) => ReactNode;
  defaultFilters: T;
  height?: number;
}

const ChartCard = <T,>({
  title,
  children,
  filters,
  defaultFilters,
  height = 350,
}: ChartCardProps<T>) => {
  const [currentFilters, setCurrentFilters] = useState<T>(defaultFilters);

  return (
    <Card style={{ height, display: "flex", flexDirection: "column" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <Text strong>{title}</Text>
        {filters?.(currentFilters, setCurrentFilters)}
      </div>
      <div style={{ flex: 1 }}>{children({ filters: currentFilters })}</div>
    </Card>
  );
};

export default ChartCard;
