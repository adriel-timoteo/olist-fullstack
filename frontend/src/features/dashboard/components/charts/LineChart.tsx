import { Line, type LineConfig } from "@ant-design/plots";
import { Card, Typography } from "antd";

const { Text } = Typography;

type LineChartProps = Partial<LineConfig> & {
  title: string;
  data: Record<string, unknown>[];
  xField: string;
  yField: string;
  color?: string;
  height?: number;
};

const LineChart = ({
  title,
  data,
  xField,
  yField,
  ...rest
}: LineChartProps) => {
  const config: LineConfig = {
    data,
    xField,
    yField,
    autoFit: true,
    smooth: true,
    point: { size: 5, shape: "diamond" },
    ...rest,
  };

  return (
    <Card>
      <Text strong>{title}</Text>
      <Line {...config} />
    </Card>
  );
};

export default LineChart;
