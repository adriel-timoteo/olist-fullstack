import { Card, Typography } from "antd";
import { type ReactNode, useState } from "react";
import dayjs, { Dayjs } from "dayjs";

const { Text } = Typography;

interface ChartCardProps {
  title: string;
  children: (props: { dateRange: [Dayjs, Dayjs] }) => ReactNode;
  filters?: (
    range: [Dayjs, Dayjs],
    setRange: (range: [Dayjs, Dayjs]) => void
  ) => ReactNode;
  defaultRange?: [Dayjs, Dayjs];
  height?: number;
}

const ChartCard = ({
  title,
  children,
  filters,
  defaultRange = [dayjs().subtract(7, "day"), dayjs()],
  height = 350,
}: ChartCardProps) => {
  const [dateRange, setDateRange] = useState<[Dayjs, Dayjs]>(defaultRange);

  return (
    <Card style={{ height, display: "flex", flexDirection: "column" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: 8,
        }}
      >
        <Text strong>{title}</Text>
        {filters?.(dateRange, setDateRange)}
      </div>
      <div style={{ flex: 1 }}>{children({ dateRange })}</div>
    </Card>
  );
};

export default ChartCard;
