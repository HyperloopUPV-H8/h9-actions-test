import styles from "./Connections.module.scss";
import { ConnectionView } from "./ConnectionView/ConnectionView";
import { useConnections } from "./useConnections";

export const Connections = () => {
    const connections = useConnections();

    return (
        <div className={styles.connectionsWrapper}>
            <ConnectionView connection={connections.backend} />
            <div className={styles.boards}>
                {Object.values(connections.boards).map((conn) => {
                    return (
                        <ConnectionView
                            key={conn.name}
                            connection={conn}
                        />
                    );
                })}
            </div>
        </div>
    );
};
