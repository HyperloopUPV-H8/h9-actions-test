import { NumericValue } from "../models/PodData/Measurement";

export type PacketUpdate = {
    readonly id: number;
    readonly hexValue: string;
    readonly cycleTime: number;
    readonly count: number;
    readonly measurementUpdates: {
        readonly [name: string]: NumericValue | boolean | string;
    };
};
