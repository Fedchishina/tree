Tree
=======================

Library for work with trees.

You can create a Tree and use a list of functions to work with it.

- [Tree functions](#Tree-functions)
    * [Tree creation example](#1-Tree-creation-example)
    * [Tree traversal](#2-Tree-traversal)
    * [Search element](#3-search-element)
    * [Min Tree element](#4-min-Tree-element)
    * [Max Tree element](#5-max-Tree-element)
    * [PreOrder Successor](#6-preorder-successor)
    * [PostOrder Successor](#7-postorder-successor)

## Tree functions
### Tree creation example

```
intTree := Tree[int]{
    value: 20,
    left:  nil,
    right: nil,
}

stringTree := Tree[string]{
    value: "root value",
    left:  nil,
    right: nil,
}
```

### Tree traversal
you can make Tree traversal:
```
t := Tree[int]{
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

### Search element

```
t := Tree[int]{
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
result := t.Search(12)    // Tree with root 12
```

### Min Tree element
```
t := Tree[int]{
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

result := t.Min() // Tree with root 4
```
### Max Tree element
```
t := Tree[int]{
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

result := t.Max() // Tree with root 22
```

### PreOrder Successor
```
t := Tree[int]{
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

result := t.PreOrderSuccessor(t.left) // Tree with root 4
```

### PostOrder Successor
```
t := Tree[int]{
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

result := t.PostOrderSuccessor(t.left) // Tree with root 10
```