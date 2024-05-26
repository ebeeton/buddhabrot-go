import PlotForm from "./plotform";
import PlotSummaryList from "./plotsummarylist";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-row p-24">
      <div className="w-1/4">
        <PlotForm>
        </PlotForm>
      </div>
      <div className="w-3/4">
        <PlotSummaryList>
        </PlotSummaryList>
      </div>
    </main>
  );
}
