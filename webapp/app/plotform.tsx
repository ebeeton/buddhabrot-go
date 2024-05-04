'use client'
export default function PlotForm() {    

    function plot(formData: FormData) {
        // Append a default gradient until an editor can be built.
        var defaultGradient = [{
            "color":"#000000",
            "position":0.0
        }, {
           "color":"#ff8000",
           "position":0.5
        }, {
           "color":"#ffff00",
           "position":0.75
        }, {
            "color":"#FFFFFF",
            "position":1.0
        }];
        // TODO:: send this to the API.
        var plotParams = { formData, gradient: defaultGradient };
        console.log(plotParams);
    }

    return (
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
            <div className="mb-4">
                <fieldset>
                    <legend className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">Plotting Region</legend>
                    <div className="mb-4 pl-4">
                        <label htmlFor="minReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                            Minimum real
                        </label>
                        <input name="minReal" defaultValue={-2.0}
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                    </div>
                    <div className="mb-4 pl-4">
                        <label htmlFor="maxReal" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                            Maximum real
                        </label>
                        <input name="maxReal" defaultValue={1.6}
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                    </div>
                    <div className="mb-4 pl-4">
                        <label htmlFor="minImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                            Minimum imaginary
                        </label>
                        <input name="minImag" defaultValue={-2}
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                    </div>
                    <div className="mb-4 pl-4">
                        <label htmlFor="maxImag" className="block uppercase tracking-wide text-gray-700 text-xs font-bold mr-2">
                            Maximum imaginary
                        </label>
                        <input name="maxImag" defaultValue={2}
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" />
                    </div>
                </fieldset>
            </div>
            <div>
                <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Plot</button>
            </div>
        </form>
    );
}