import Region from "./region";
import Stop from "./stop";

// Parameters used to plot the 'brot.
export default interface PlotParams {
    readonly sampleSize: number,
    readonly maxIterations: number,
    readonly region: Region,
    readonly width: number,
    readonly height: number,
    readonly gradient: Stop[]
};