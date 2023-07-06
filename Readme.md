Tree
=======================

Library for work with trees.

You can create a tree and use a list of functions to work with it.

- [Tree](#tree)
    * [Tree functions](#tree-functions)
        + [1. Tree creation example](#1-tree-creation-example)
        + [2. Tree traversal](#2-tree-traversal)
        + [3. Search element](#3-search-element)
        + [4. Min tree element](#4-min-tree-element)
        + [5. Max tree element](#5-max-tree-element)
        + [6. PreOrder Successor](#6-preorder-successor)
        + [7. PostOrder Successor](#7-postorder-successor)

## Tree functions
### 1. Tree creation example

```
intTree := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}

stringTree := tree[string]{
    value: "root value",
    left:  nil,
    right: nil,
}
```

### 2. Tree traversal
you can make tree traversal:
```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)

resultAsc := t.InOrderTreeWalk(Asc)   // [4 8 20 22]
resultDesc := t.InOrderTreeWalk(Desc) // [22 20 8 4]
```

### 3. Search element

```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

resultNil := t.Search(15) //nil
result := t.Search(12)    // tree with root 12
```

### 4. Min tree element
```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.Min() // tree with root 4
```
### 5. Max tree element
```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.Max() // tree with root 22
```

### 6. PreOrder Successor
```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.preOrderSuccessor(t.left) // tree with root 4
```

### 7. PostOrder Successor
```
t := tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.postOrderSuccessor(t.left) // tree with root 10
```