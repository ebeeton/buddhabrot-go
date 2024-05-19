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
            .then(body => setPlots(body))
            .catch(console.error);
    });

    return (
        <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
            {plots?.map(p => (<PlotSummary plot={p} key={p.ID} />))}
        </div>
    );
}