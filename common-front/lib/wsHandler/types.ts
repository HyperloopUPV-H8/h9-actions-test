import { HandlerMessages } from "./HandlerMessages";

export type WsMessage = {
    topic: string;
    payload: any;
};

export type PostRequest<Req, Res> = {
    type: "post";
    request: Req;
    response: Res;
};

export type Subscription<T> = {
    type: "subscription";
    subscribe: boolean;
    id: string;
    response: T;
};

export type Exchange<Req, Res> = {
    type: "exchange";
    request: Req;
    response: Res;
};

export type PostTopic = TopicFromValueType<PostRequest<any, any>>;
export type SubscriptionTopic = TopicFromValueType<Subscription<any>>;
export type ExchangeTopic = TopicFromValueType<Exchange<any, any>>;

type TopicFromValueType<T> = {
    [K in keyof HandlerMessages]: HandlerMessages[K] extends T ? K : never;
}[keyof HandlerMessages];
