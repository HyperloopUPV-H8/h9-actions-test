import { BrokerStructure } from "./BrokerStructure";

export type Callback<T extends Topic> = (msg: Response<T>) => unknown;
export type BackendMessage<
    T extends keyof BrokerStructure = keyof BrokerStructure
> = {
    topic: string;
    msg: Response<T>;
};

export type Topic = keyof BrokerStructure & string;

export type Request<T extends keyof BrokerStructure> =
    BrokerStructure[T]["request"];
export type Response<T extends keyof BrokerStructure> =
    BrokerStructure[T]["response"];
