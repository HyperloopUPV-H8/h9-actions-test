import { useReducer } from "react";
import { ChartId, ChartInfo, MeasurementId } from "./types";
import { ChartActions } from "./ChartActions";

function reducer(state: ChartInfo[], action: ChartActions): ChartInfo[] {
    switch (action.type) {
        case "add_chart":
            const {chartId, measurementId} = action.payload;
            return [
                ...state,
                {
                    chartId: chartId,
                    measurementId: measurementId, 
                },
            ];
        case "remove_chart": {
            return state.filter((e) => (e.chartId !== action.payload));
        }
        default: {
            return state;
        }
    }
}

// Custom hook to manage the charts showed in the ChartMenu component.
// It stores each chart with an id and a list of different measurements to show in it.
export function useCharts() {
    const [charts, dispatch] = useReducer(reducer, []);

    return {
        charts,
        addChart: (chartId: ChartId, measurementId: MeasurementId) =>
            dispatch({ type: "add_chart", payload: { chartId: chartId, measurementId: measurementId } }),
        removeChart: (chartId: ChartId) =>
            dispatch({ type: "remove_chart", payload: chartId }),
    };
}
