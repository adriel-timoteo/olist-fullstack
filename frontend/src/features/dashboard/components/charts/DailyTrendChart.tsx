import { Line } from "@ant-design/plots";

const DailyTrendChart = () => {
  const data = [
    { date: "2025-07-01", value: 120 },
    { date: "2025-07-02", value: 200 },
    { date: "2025-07-03", value: 150 },
    { date: "2025-07-04", value: 80 },
    { date: "2025-07-05", value: 170 },
  ];

  const config = {
    data,
    xField: "date",
    yField: "value",
    autoFit: true,
    height: 300,
    point: {
      size: 5,
      shape: "diamond",
    },
    smooth: true,
    color: "#1677ff",
  };

  return <Line {...config} />;
};

export default DailyTrendChart;
