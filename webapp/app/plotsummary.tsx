import { APIURL, PLOTROUTE } from "./apiRoutes";
import Plot from "./plot";

export default function PlotSummary({plot}: {plot: Plot}) {
    return (
        <tr key={plot.ID}>
            <td>{plot.Filename ? (<a href={`${APIURL}${PLOTROUTE}/${plot.ID}`} target="_blank" rel="noreferrer">Image</a>) : "Not Plotted Yet"}</td>
            <td>{new Date(plot.CreatedAt).toLocaleString()}</td>
            <td>{new Date(plot.UpdatedAt).toLocaleString()}</td>
        </tr>
    );
}