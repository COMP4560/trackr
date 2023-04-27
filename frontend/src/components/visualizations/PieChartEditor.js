import { useState, forwardRef, useImperativeHandle } from "react";
import PieChart from "./PieChart";

const PieChartEditor = forwardRef(({ metadata, setError }, ref) => {
  const [sort] = useState(metadata?.sort || "");

  useImperativeHandle(ref, () => ({
    submit() {
      return PieChart.serialize(sort);
    },
  }));

  return <></>;
});

export default PieChartEditor;
