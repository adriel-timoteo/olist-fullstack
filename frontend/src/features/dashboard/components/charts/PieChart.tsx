import { Pie, type PieConfig } from "@ant-design/plots";
import { Card, Typography } from "antd";

const { Text } = Typography;

type PieChartProps = Partial<PieConfig> & {
  title: string;
  data: Record<string, unknown>[];
  angleField: string;
  colorField: string;
  height?: number;
};

const PieChart = ({
  title,
  data,
  angleField,
  colorField,
  ...rest
}: PieChartProps) => {
  const config: PieConfig = {
    data,
    angleField,
    colorField,
    radius: 1,
    autoFit: true,
    ...rest,
  };

  return (
    <Card>
      <Text strong>{title}</Text>
      <Pie {...config} />
    </Card>
  );
};

export default PieChart;
