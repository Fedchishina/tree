tree
=======================

Library for work with Binary trees.

You can create a Binary node and use a list of functions to work with it. 

## Tree functions
  - [Empty tree's creation example](#empty-trees-creation-example)
  - [Tree's creation with one element example](#trees-creation-with-one-element-example)
  - [Insert element to tree](#insert-element-to-tree)
  - [Tree traversal](#tree-traversal)
  - [Exists element](#exists-element)
  - [Get value by key element](#get-value-by-key-element)
  - [Min tree element](#min-tree-element)
  - [Max tree element](#max-tree-element)
  - [PreOrder Successor](#preorder-successor)
  - [PostOrder Successor](#postorder-successor)
  - [Delete node from node](#delete-node-from-node)


### Empty tree's creation example

```
t := tree.New[int]() // empty int tree
t := tree.New[string]() // empty string tree
```

### Tree's creation with one element example

```
t := tree.NewWithElement[int](1,1) // int tree creation with one element
t := tree.NewWithElement[string]("key", "value") // string tree creation with one element
```

### Insert element to tree
```
t := tree.New[int]() // empty int tree
t.Insert(22, 22) // insert to tree element: key=22, value=22
t.Insert(8, 8) // insert to tree element: key=8, value=8
t.Insert(4, 4) // insert to tree element: key=4, value=4

// or
t.InsertWithoutRecursion(4, 4) // insert to tree element: key=4, value=4
```

### Tree traversal
you can make tree traversal by two methods:
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultAsc := t.InOrderTreeWalk(tree.Asc)   // [4, 8, 22]
resultDesc := t.InOrderTreeWalk(tree.Desc)   // [22, 8, 4]

// or
resultAsc := t.InOrderTreeWalkWithStack(tree.Asc)   // [4, 8, 22]
resultDesc := t.InOrderTreeWalkWithStack(tree.Desc)   // [22, 8, 4]

```

### Exists element

```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil := t.Exists(15) // false
result    := t.Exists(8)  // true
```

### Get value by key element

```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil, err := t.GetValue(15) // nil, err
result, err    := t.GetValue(8)  // 8, nil
```

### Min tree element
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Min() // 4
```
### Max tree element
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Max() // 22
```

### PreOrder Successor

```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil, err := t.PreOrderSuccessor(22) // nil, err
result, err    := t.PreOrderSuccessor(8)  // 22, nil
```

### PostOrder Successor
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil, err := t.PostOrderSuccessor(4) // nil, err
result, err    := t.PostOrderSuccessor(8)  // 4, nil
```

### Delete element by key from tree
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

err := t.Delete(22) // without err
```