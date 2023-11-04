export type NumericInputData = {
    readonly value: number | null;
    readonly placeholder?: string;
    readonly isValid?: boolean;
};

export type BooleanInputData = {
    readonly value: boolean;
};

export type EnumInputData = {
    readonly options: string[];
    readonly value?: string;
};

export type ExpandablePairsInputData = {
    value: [number | null, number | null][];
};

export type Inputs = {
    numeric: NumericInputData;
    boolean: BooleanInputData;
    enum: EnumInputData;
    expandablePairs: ExpandablePairsInputData;
};

export type InputEvent = {
    [K in keyof Inputs]: { type: K; value: Inputs[K]["value"] };
}[keyof Inputs];
