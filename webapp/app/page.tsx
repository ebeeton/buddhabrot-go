import PlotForm from "./plotform";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <PlotForm apiUrl={process.env.PLOTTER_API as string}>
      </PlotForm>      
    </main>
  );
}
