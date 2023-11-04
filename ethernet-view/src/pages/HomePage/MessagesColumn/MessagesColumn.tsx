import styles from "pages/HomePage/MessagesColumn/MessagesColumn.module.scss";
import { SplitLayout } from "layouts/SplitLayout/SplitLayout";
import { Orientation } from "hooks/useSplit/Orientation";
import { TabLayout } from "layouts/TabLayout/TabLayout";
import { BiLineChart } from "react-icons/bi";
import { nanoid } from "nanoid";
import { MessagesContainer } from "components/MessagesContainer/MessagesContainer";
import { Logger } from "components/Logger/Logger";
import { useRef } from "react";
import { Connections } from "common";
import { BootloaderContainer } from "components/BootloaderContainer/BootloaderContainer";
import { Island } from "components/Island/Island";

export const MessagesColumn = () => {
    const messagesTabItems = useRef([
        {
            id: nanoid(),
            name: "Messages",
            icon: <BiLineChart />,

            component: <MessagesContainer />,
        },
    ]);

    return (
        <div className={styles.messageColumnWrapper}>
            <SplitLayout
                components={[
                    <TabLayout items={messagesTabItems.current}></TabLayout>,
                    <Island>
                        <Connections />
                    </Island>,
                ]}
                orientation={Orientation.VERTICAL}
            ></SplitLayout>
            <Logger />
            <BootloaderContainer />
        </div>
    );
};
