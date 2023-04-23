import { useToggle } from "hooks/useToggle";
import { useEffect } from "react";
import style from "./ToggleSwitch.module.scss";

type Props = {
    onToggle: (isOn: boolean) => void;
    isOn: boolean;
    flip: () => void;
};

export function ToggleSwitch({ onToggle, isOn, flip }: Props) {
    useEffect(() => {
        onToggle(isOn); //TODO:onToggle?.(isOn); estaba así¿?
    }, [isOn]);

    return (
        <label
            className={
                isOn
                    ? style.toggleSwitchWrapperOn
                    : style.toggleSwitchWrapperOff
            }
            onClick={flip}
        >
            <div></div>
        </label>
    );
}
