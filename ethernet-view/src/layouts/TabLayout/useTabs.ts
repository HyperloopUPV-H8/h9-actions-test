import { TabItem } from "layouts/TabLayout/TabItem";
import { useState } from "react";

export function useTabs(tabItems: TabItem[]) {
    //FIXME: fails if item length == 0
    const [visibleTab, setVisibleTab] = useState(tabItems[0]);

    return [tabItems, visibleTab, setVisibleTab] as const;
}
