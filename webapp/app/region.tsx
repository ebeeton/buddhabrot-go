// Plot area of the complex plane.
export default interface Region {
    readonly minReal: number,
    readonly maxReal: number,
    readonly minImag: number,
    readonly maxImag: number
};