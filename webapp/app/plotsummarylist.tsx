'use client'
import { useState, useEffect } from "react";
import Plot from "./plot";
import PlotSummary from "./plotsummary";
import { APIURL, PLOTROUTE } from "./apiRoutes";

export default function PlotSummaryList() {
    let [plots, setPlots] = useState<Plot[]>();
    useEffect(() => {
        fetch(new URL(PLOTROUTE, APIURL))
            .then(response => response.json())
            .then(body => {
                setPlots(body);
                console.log("Plot summary list updated.");
            })
            .catch(console.error);
    }, []);

    return (
        <div className="bg-white shadow-md rounded ml-8 px-8 pt-6">
            <table>
                <tbody>
                    <tr>
                        <th>
                            Image
                        </th>
                        <th>
                            Created
                        </th>
                        <th>
                            Updated
                        </th>
                    </tr>
                    {plots?.map(p => (<PlotSummary plot={p} key={p.ID} />))}
                </tbody>
            </table>
        </div>
    );
}