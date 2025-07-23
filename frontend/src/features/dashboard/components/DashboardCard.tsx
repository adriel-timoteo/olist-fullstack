import React from "react";
import { Card } from "antd";

type Props = {
  title: string;
  children: React.ReactNode;
};

const DashboardCard = ({ title, children }: Props) => (
  <Card style={{ marginBottom: 16 }} title={title} bordered={false}>
    {children}
  </Card>
);

export default DashboardCard;
