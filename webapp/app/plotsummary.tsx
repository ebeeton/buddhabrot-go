import { APIURL, PLOTROUTE } from "./apiRoutes";
import Plot from "./plot";

export default function PlotSummary({plot}: {plot: Plot}) {
    return (
        <div key={plot.ID}>
            <p>Created: {plot.CreatedAt}</p>
            <p>Updated: {plot.UpdatedAt}</p>
            <p>{plot.Filename ? (<a href={`${APIURL}${PLOTROUTE}/${plot.ID}`} target="_blank">Image</a>) : "Not Plotted Yet."}</p>
        </div>
    );
}