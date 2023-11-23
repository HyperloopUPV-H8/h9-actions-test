import { InfoMessage } from "../../../../..";
import styles from "./InfoMessageView.module.scss";

type Props = {
    message: InfoMessage;
    className: string;
};

export const InfoMessageView = ({ message, className }: Props) => {
    return (
        <div className={`${styles.infoMessageView} ${className}`}>
            <div className={styles.board}>{message.board}</div>
            <div>{message.payload}</div>
        </div>
    );
};
