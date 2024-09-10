
export type NodeTS<T> = {
    value: T,
    prev?: NodeTS<T>,
    next?: NodeTS<T>,
};

