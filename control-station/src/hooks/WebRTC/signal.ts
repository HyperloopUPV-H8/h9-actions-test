export type Signal<Kind extends SignalKinds> = {
    name: Kind;
    payload: SignalPayloadMap[Kind];
};

export type SignalHandles = {
    [kind in SignalKinds]: (signal: Signal<kind>) => void;
};

export type SignalKinds = "offer" | "answer" | "candidate" | "close" | "poll";
export type SignalPayloadMap = {
    offer: RTCSessionDescriptionInit;
    answer: RTCSessionDescriptionInit;
    candidate: RTCIceCandidateInit;
    close: null;
    poll: null;
};
