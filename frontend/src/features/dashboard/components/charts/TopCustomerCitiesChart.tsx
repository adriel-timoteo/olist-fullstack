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
    autoFit: true,
    height: 300,
  };

  return <Bar {...config} />;
};

export default TopCustomerCitiesChart;
