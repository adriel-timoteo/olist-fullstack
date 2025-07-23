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
    color: "#52c41a",
    columnWidthRatio: 0.5,
  };

  return <Column {...config} />;
};

export default MonthlyTrendChart;
