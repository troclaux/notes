import { TreeNode } from "./TreeNode";

export function walk(curr: TreeNode<number> | null, path: number[]): number[] {
    if (!curr) {
        return path;
    }

    path.push(curr.value);

    walk(curr.left, path);
    walk(curr.right, path);

    return path;
}
