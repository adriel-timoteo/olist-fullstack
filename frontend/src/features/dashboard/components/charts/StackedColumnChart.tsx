import { Column, type ColumnConfig } from "@ant-design/plots";
import { Spin } from "antd";
import type { Dayjs } from "dayjs";
import { useEffect, useState } from "react";

type StackedColumnChartProps = Partial<ColumnConfig> & {
  data?: Record<string, unknown>[];
  xField: string;
  yField: string;
  colorField: string;
  fetchData?: (
    start: string,
    end: string
  ) => Promise<Record<string, unknown>[]>;
  dateRange?: [Dayjs, Dayjs];
  height?: number;
};

const StackedColumnChart = ({
  data = [],
  xField,
  yField,
  colorField,
  fetchData,
  dateRange,
  height = 300,
  ...rest
}: StackedColumnChartProps) => {
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

  const config: ColumnConfig = {
    data: chartData,
    xField,
    yField,
    colorField,
    stack: true,
    autoFit: true,
    height,
    ...rest,
  };

  return loading ? (
    <Spin />
  ) : chartData.length > 0 ? (
    <Column {...config} />
  ) : (
    <div>No data available</div>
  );
};

export default StackedColumnChart;
