'use client'
import { useState, useEffect } from "react";
import Plot from "./plot";
import PlotSummary from "./plotsummary";
import { APIURL, PLOTROUTE } from "./apiRoutes";

export default function PlotSummaryList() {
    let [plots, setPlots] = useState<Plot[]>();
    const refreshMS = 3000;
    // Refresh on an interval from https://stackoverflow.com/a/64144607/2382333
    useEffect(() => {
        const fetchData = async() => {
            fetch(new URL(PLOTROUTE, APIURL))
            .then(response => response.json())
            .then(body => {
                setPlots(body);
                console.debug("Plot summary list updated.");
            })
            .catch(console.error);
        };

        fetchData();

        const id = setInterval(() => {
            fetchData();
        }, refreshMS);

        return () => clearInterval(id);
    }, []);

    return (
        <div className="bg-white shadow-md rounded ml-8 px-8 pt-6">
            <table className="w-full">
                <caption>
                    Plot History
                </caption>
                <tbody>
                    <tr className="text-left">
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