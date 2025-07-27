import { Bar, type BarConfig } from "@ant-design/plots";
import { Card, Typography } from "antd";

const { Text } = Typography;

type BarChartProps = Partial<BarConfig> & {
  data: Record<string, unknown>[];
  xField: string;
  yField: string;
  seriesField?: string;
  height?: number;
  title: string;
};

const BarChart = ({
  data,
  title,
  xField,
  yField,
  seriesField,
  ...rest
}: BarChartProps) => {
  const config: BarConfig = {
    data,
    xField,
    yField,
    seriesField,
    autoFit: true,
    ...rest,
  };

  return (
    <Card>
      <Text strong>{title}</Text>
      <Bar {...config} />
    </Card>
  );
};

export default BarChart;
