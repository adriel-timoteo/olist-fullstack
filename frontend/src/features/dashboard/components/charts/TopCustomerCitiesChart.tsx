import { Bar } from "@ant-design/plots";

const TopCustomerCitiesChart = () => {
  const data = [
    { city: "Jakarta", customers: 200 },
    { city: "Surabaya", customers: 150 },
    { city: "Bandung", customers: 100 },
    { city: "Medan", customers: 80 },
  ];

  const config = {
    data,
    xField: "customers",
    yField: "city",
    seriesField: "city",
    color: "#faad14",
    legend: false,
  };

  return <Bar {...config} />;
};

export default TopCustomerCitiesChart;
