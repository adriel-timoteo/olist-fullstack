import { useEffect, useState } from "react";
import { Spin } from "antd";
import { Line, type LineConfig } from "@ant-design/plots";
import { Dayjs } from "dayjs";

type LineChartProps = Partial<LineConfig> & {
  xField: string;
  yField: string;
  fetchData?: (
    start: string,
    end: string
  ) => Promise<Record<string, unknown>[]>;
  dateRange?: [Dayjs, Dayjs];
  data?: Record<string, unknown>[];
  height?: number;
};

const LineChart = ({
  xField,
  yField,
  fetchData,
  dateRange,
  data = [],
  height = 300,
  ...rest
}: LineChartProps) => {
  const [chartData, setChartData] = useState(data);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (fetchData && dateRange) {
      const [start, end] = dateRange;
      setLoading(true);
      fetchData(start.toISOString(), end.toISOString())
        .then((res) => setChartData(res ?? []))
        .finally(() => setLoading(false));
    }
  }, [fetchData, dateRange]);

  const config: LineConfig = {
    data: chartData,
    xField,
    yField,
    autoFit: true,
    smooth: true,
    point: { size: 5, shape: "diamond" },
    height,
    ...rest,
  };

  return loading ? (
    <Spin />
  ) : chartData.length > 0 ? (
    <Line {...config} />
  ) : (
    <div>No data available</div>
  );
};

export default LineChart;
