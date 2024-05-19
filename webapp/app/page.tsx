import PlotForm from "./plotform";
import PlotSummaryList from "./plotsummarylist";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <PlotForm>
      </PlotForm>
      <PlotSummaryList>
      </PlotSummaryList>
    </main>
  );
}
