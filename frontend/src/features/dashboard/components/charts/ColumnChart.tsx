import { Column, type ColumnConfig } from "@ant-design/plots";
import { Card, Typography } from "antd";

const { Text } = Typography;

type ColumnChartProps = Partial<ColumnConfig> & {
  title: string;
  data: Record<string, unknown>[];
  xField: string;
  yField: string;
  height?: number;
};

const ColumnChart = ({
  title,
  data,
  xField,
  yField,
  ...rest
}: ColumnChartProps) => {
  const config: ColumnConfig = {
    data,
    xField,
    yField,
    autoFit: true,
    ...rest,
  };

  return (
    <Card>
      <Text strong>{title}</Text>
      <Column {...config} />
    </Card>
  );
};

export default ColumnChart;
