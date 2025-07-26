import { Card } from "antd";

interface DashboardCardProps {
  title: string;
  children: React.ReactNode;
  height?: string | number;
}

const DashboardCard = ({
  title,
  children,
  height = 300,
}: DashboardCardProps) => {
  return (
    <Card
      title={<span className="text-lg font-semibold">{title}</span>}
      bordered={false}
      style={{
        height,
        borderRadius: "16px",
        boxShadow: "0 2px 8px rgba(0, 0, 0, 0.06)",
        overflow: "hidden",
      }}
    >
      <div style={{ height: "100%" }}>{children}</div>
    </Card>
  );
};

export default DashboardCard;
