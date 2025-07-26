import { Column } from "@ant-design/plots";

const MonthlyTrendChart = () => {
  const data = [
    { month: "Jan", value: 300 },
    { month: "Feb", value: 450 },
    { month: "Mar", value: 320 },
    { month: "Apr", value: 510 },
  ];

  const config = {
    data,
    xField: "month",
    yField: "value",
    autoFit: true,
    height: 300,
  };

  return <Column {...config} />;
};

export default MonthlyTrendChart;
