export type TreeNode<T> = {
    value: T,
    left: TreeNode<T> | null,
    right: TreeNode<T> | null,
};
