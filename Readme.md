tree
=======================

Library for work with Binary trees.

You can create a Binary node and use a list of functions to work with it. 

## Tree functions
  - [Empty tree's creation example](#empty-trees-creation-example)
  - [Tree's creation with one element example](#trees-creation-with-one-element-example)
  - [Insert element to tree](#insert-element-to-tree)
  - [Tree traversal](#tree-traversal)
  - [Search element](#search-element)
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
```

### Tree traversal
you can make tree traversal:
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultAsc := t.InOrderTreeWalk(tree.Asc)   // [Element{4, 4}, Element{8, 8}, Element{22, 22}]
resultDesc := t.InOrderTreeWalk(tree.Desc)   // [Element{22, 22}, Element{8, 8}, Element{4, 4}]
```

### Search element

```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil := t.Search(15) // nil
result    := t.Search(8)  // Element {key:8, value: 8}
```

### Min tree element
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Min() // Element {key:4, value: 4}
```
### Max tree element
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Max() // Element {key:22, value: 22}
```

### PreOrder Successor
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil := t.PreOrderSuccessor(22) // nil
result    := t.PreOrderSuccessor(8)  // Element {key:22, value: 22}
```

### PostOrder Successor
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

resultNil := t.PostOrderSuccessor(4) // nil
result    := t.PostOrderSuccessor(8)  // Element {key:4, value: 4}
```

### Delete element by key from tree
```
t := tree.New[int]()
t.Insert(22, 22) 
t.Insert(8, 8)
t.Insert(4, 4)

err := t.Delete(22) // without err
```