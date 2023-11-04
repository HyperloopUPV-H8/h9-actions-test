import styles from "./BoardsPage.module.scss";
import { useMeasurements } from "../useMeasurements";
import { LCU } from "../Boards/LCU/LCU";
import { PCU } from "../Boards/PCU/PCU";
import { BMSL } from "../Boards/BMSL/BMSL";
import { OBCCU } from "../Boards/OBCCU/OBCCU";
import { VCU } from "../Boards/VCU/VCU";
import {
    Connections,
    selectBmslMeasurements,
    selectLcuMeasurements,
    selectObccuMeasurements,
    selectPcuMeasurements,
    selectVcuMeasurements,
} from "common";

export const BoardsPage = () => {
    const measurements = useMeasurements();

    return (
        <div className={styles.boardsPage}>
            <OBCCU {...selectObccuMeasurements(measurements)} />
            <div className={styles.column}>
                <VCU {...selectVcuMeasurements(measurements)} />
                <PCU {...selectPcuMeasurements(measurements)} />
            </div>
            <div className={styles.gridColumn}>
                <Connections />
                <LCU {...selectLcuMeasurements(measurements)} />
                <BMSL {...selectBmslMeasurements(measurements)} />
            </div>
        </div>
    );
};
