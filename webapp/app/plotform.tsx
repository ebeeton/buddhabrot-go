'use client'

import { redirect } from "next/navigation";

export default function PlotForm({ apiUrl }: { apiUrl: string }) {
    // Plot area of the complex plane.
    interface Region {
        readonly minReal: number,
        readonly maxReal: number,
        readonly minImag: number,
        readonly maxImag: number
    };

    // A hex color and its position in a gradient.
    interface Stop {
        readonly color: string,
        readonly position: number
    }

    // Parameters used to plot the 'brot.
    interface PlotParams {
        readonly sampleSize: number,
        readonly maxIterations: number,
        readonly region: Region,
        readonly width: number,
        readonly height: number,
        readonly gradient: Stop[]
    };

    // Data returned by the plot request.
    interface PlotResponse {
        readonly id: number
    }

    async function plot(formData: FormData) {
        const PlotRoute = "/api/Plots";
        console.log(`API URL: ${new URL(PlotRoute, apiUrl)}`);
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

        let plotResponse = await fetch(new URL(PlotRoute, apiUrl), {
            method: "POST",
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify(plotParams)
        })
            .then(response => response.json())
            .then(body => body as PlotResponse)
            .catch(console.error);

        // TODO:: Redirect to... something?
        console.log(plotResponse);
    }

    return (
        <div>
            <form action={plot} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
                <div className="mb-4">
                    <label htmlFor="width" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Width
                    </label>
                    <input name="width" defaultValue={512}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="height" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Height
                    </label>
                    <input name="height" defaultValue={512}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="sampleSize" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Sample size
                    </label>
                    <input name="sampleSize" defaultValue={100000000}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="maxIterations" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Maximum iterations
                    </label>
                    <input name="maxIterations" defaultValue={5000}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2 mb-4">Plotting Region
                    <hr />
                </div>
                <div className="mb-4">
                    <label htmlFor="minReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Minimum real
                    </label>
                    <input name="minReal" defaultValue={-2.0}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="maxReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Maximum real
                    </label>
                    <input name="maxReal" defaultValue={1.6}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="minImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Minimum imaginary
                    </label>
                    <input name="minImag" defaultValue={-2}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div className="mb-4">
                    <label htmlFor="maxImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                        Maximum imaginary
                    </label>
                    <input name="maxImag" defaultValue={2}
                        className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                </div>
                <div>
                    <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Plot</button>
                </div>
            </form>
        </div>
    );
}