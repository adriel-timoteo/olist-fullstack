import { Column, type ColumnConfig } from "@ant-design/plots";
import { Spin } from "antd";
import { useEffect, useState } from "react";

type ColumnChartProps = Partial<ColumnConfig> & {
  data?: Record<string, unknown>[];
  xField: string;
  yField: string;
  height?: number;
  fetchData?: () => Promise<Record<string, unknown>[]>;
};

const ColumnChart = ({
  data = [],
  fetchData,
  height = 300,
  xField,
  yField,
  ...rest
}: ColumnChartProps) => {
  const [chartData, setChartData] = useState(data);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (fetchData) {
      setLoading(true);
      fetchData()
        .then((res) => setChartData(res ?? []))
        .finally(() => setLoading(false));
    }
  }, [fetchData]);

  const config: ColumnConfig = {
    data: chartData,
    xField,
    yField,
    autoFit: true,
    height,
    ...rest,
  };

  console.log(chartData);

  return loading ? (
    <Spin />
  ) : chartData.length > 0 ? (
    <Column {...config} />
  ) : (
    <div>No data available</div>
  );
};

export default ColumnChart;
