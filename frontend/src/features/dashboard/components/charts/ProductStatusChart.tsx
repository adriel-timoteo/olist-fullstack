import { Pie } from "@ant-design/plots";

const ProductStatusChart = () => {
  const data = [
    { type: "Delivered", value: 27 },
    { type: "In Transit", value: 25 },
    { type: "Pending", value: 18 },
    { type: "Cancelled", value: 10 },
  ];

  const config = {
    data,
    angleField: "value",
    colorField: "type",
    radius: 1,
    autoFit: true,
    height: 300,
  };

  return <Pie {...config} />;
};

export default ProductStatusChart;
