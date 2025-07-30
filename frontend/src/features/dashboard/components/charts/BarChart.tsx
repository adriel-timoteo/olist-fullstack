import { Bar, type BarConfig } from "@ant-design/plots";
import { Spin } from "antd";
import { useEffect, useState } from "react";

type BarChartProps = Partial<BarConfig> & {
  data?: Record<string, unknown>[];
  xField: string;
  yField: string;
  seriesField?: string;
  fetchData?: (limit: number) => Promise<Record<string, unknown>[]>;
  limit?: number;
  height?: number;
};

const BarChart = ({
  data = [],
  xField,
  yField,
  seriesField,
  fetchData,
  limit = 10,
  height = 300,
  ...rest
}: BarChartProps) => {
  const [chartData, setChartData] = useState(data);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (fetchData) {
      setLoading(true);
      fetchData(limit)
        .then((res) => setChartData(res ?? []))
        .finally(() => setLoading(false));
    }
  }, [fetchData, limit]);

  const config: BarConfig = {
    data: chartData,
    xField,
    yField,
    seriesField,
    autoFit: true,
    height,
    ...rest,
  };

  return loading ? (
    <Spin />
  ) : chartData.length > 0 ? (
    <Bar {...config} />
  ) : (
    <div>No data available</div>
  );
};

export default BarChart;
