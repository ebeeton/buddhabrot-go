'use client'

import { useState } from "react";
import { APIURL, PLOTROUTE } from "./apiRoutes";
import PlotParams from "./plotParams";

export default function PlotForm() {
    // Data returned by the plot request.
    interface PlotResponse {
        readonly id: number
    }

    const MinReal = -2,
        MaxReal = 1.6,
        MinImag = -2,
        MaxImag = 2;

    const [minReal, setMinReal] = useState(MinReal);
    const [maxReal, setMaxReal] = useState(MaxReal);
    const [minImag, setMinImag] = useState(MinImag);
    const [maxImag, setMaxImag] = useState(MaxImag);
    const [width, setWidth] = useState(512);
    const [height, setHeight] = useState(512);

    async function plot(formData: FormData) {
        // Copy the plot parameters from the submitted form.
        const plotParams: PlotParams = {
            sampleSize: +(formData.get("sampleSize") as string),
            maxIterations: +(formData.get("maxIterations") as string),
            region: {
                minReal: +(formData.get("minReal") as string),
                maxReal: +(formData.get("maxReal") as string),
                minImag: +(formData.get("minImag") as string),
                maxImag: +(formData.get("maxImag") as string)
            },
            width: +(formData.get("width") as string),
            height: +(formData.get("height") as string),
            // Default gradient until an editor can be built.
            gradient: [{
                "color": "#000000",
                "position": 0.0
            }, {
                "color": "#ff8000",
                "position": 0.5
            }, {
                "color": "#ffff00",
                "position": 0.75
            }, {
                "color": "#FFFFFF",
                "position": 1.0
            }]
        };

        let plotResponse = await fetch(new URL(PLOTROUTE, APIURL), {
            method: "POST",
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify(plotParams)
        })
            .then(response => response.json())
            .then(body => body as PlotResponse)
            .catch(console.error);

        // TODO:: Add a message indicating the plot was submitted successfully.
        console.log(plotResponse);
    }

    function random(min: number, max: number): number {
        return Math.random() * (max - min) + min;
    }

    function randomizeRegion() {
        // Get two points on the real axis.
        const reals = [ random(MinReal, MaxReal), random(MinReal, MaxReal)];
        const minR = Math.min(...reals),
              maxR = Math.max(...reals);
        
        // Figure out the imaginary height while maintaining the image aspect ratio.
        const aspectRatio = height / width;
        const widthR = maxR - minR;
        const heightI = widthR * aspectRatio;
        console.debug(`Aspect ratio: ${aspectRatio} Real width: ${widthR} Imaginary height: ${heightI}`);

        // Get a minimum imaginary number, and offset the maximum from it.
        const minI = random(MinImag, MaxImag - heightI);
        const maxI = minI + heightI;
        
        setMinReal(Number(minR.toFixed(3)));
        setMaxReal(Number(maxR.toFixed(3)));
        setMinImag(Number(minI.toFixed(3)));
        setMaxImag(Number(maxI.toFixed(3)));
    }

    return (
        <form action={plot} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
            <div className="mb-4">
                <label htmlFor="width" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Width
                </label>
                <input name="width" defaultValue={width} type="number" min="1" onChange={e => setWidth(Number(e.target.value))}
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="height" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Height
                </label>
                <input name="height" defaultValue={height} type="number" min="1" onChange={e => setHeight(Number(e.target.value))}
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="sampleSize" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Sample size
                </label>
                <input name="sampleSize" defaultValue={100000000} type="number" min="1"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="maxIterations" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Maximum iterations
                </label>
                <input name="maxIterations" defaultValue={5000} type="number" min="1"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2 mb-4">Plotting Region
                <hr />
            </div>
            <div className="mb-4">
                <label htmlFor="minReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Minimum real
                </label>
                <input name="minReal" defaultValue={minReal} type="number" min="-2.0" max="1.6" step="0.0001"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="maxReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Maximum real
                </label>
                <input name="maxReal" defaultValue={maxReal} type="number" min="-2.0" max="1.6" step="0.0001"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="minImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Minimum imaginary
                </label>
                <input name="minImag" defaultValue={minImag} type="number" min="-2.0" max="2.0" step="0.0001"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div className="mb-4">
                <label htmlFor="maxImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                    Maximum imaginary
                </label>
                <input name="maxImag" defaultValue={maxImag} type="number" min="-2.0" max="2.0" step="0.0001"
                    className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
            </div>
            <div>
                <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Plot</button>
                <button type="button" onClick={randomizeRegion} className="ml-4 outline hover:text-white hover:bg-blue-700 outline-blue-500 text-blue-500 font-bold py-2 px-4 rounded outline-1">Random Region</button>
            </div>
        </form>
    );
}