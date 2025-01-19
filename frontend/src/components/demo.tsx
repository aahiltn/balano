import React from "react";

interface DemoProps {
  title?: string;
}

const DemoReact: React.FC<DemoProps> = ({ title = "Demo Component" }) => {
  return (
    <div className="demo-container">
      <h1>{title}</h1>
      <p>This is a simple demo component.</p>
    </div>
  );
};

export default DemoReact;
