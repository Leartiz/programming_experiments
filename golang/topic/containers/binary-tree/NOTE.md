# Binary tree & traversals

Бинарное дерево: у узла не больше двух детей (`Left`, `Right`).

## BT
```
        1
      /   \
     2     3
    / \
   4   5
```
> inorder: 4 2 5 1 3

## BST
```
        4
      /   \
     2     5
    / \
   1   3
```
> inorder: 1 2 3 4 5

## Обходы (DFS / BFS)

> DFS - Depth-First Search (поиск в глубину)
> BFS - Breadth-First Search (поиск в ширину)

| Имя | Порядок | Как |
|-----|---------|-----|
| preorder | N -> L -> R | рекурсия / стек |
| inorder | L -> N -> R | рекурсия / стек |
| postorder | L -> R -> N | рекурсия / стек |
| level-order | по уровням | очередь (BFS) |

> Для BST **inorder** даёт отсортированную последовательность.

## Пример дерева в main

```
        1
      /   \
     2     3
    / \
   4   5
```

- preorder:  `1 2 4 5 3`
- inorder:   `4 2 5 1 3`
- postorder: `4 5 2 3 1`
- level:     `1 2 3 4 5`
