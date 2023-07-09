tree
=======================

Library for work with Binary trees.

You can create a Binary tree and use a list of functions to work with it. 

## tree functions
  - [Tree creation example](#tree-creation-example)
  - [Tree traversal](#tree-traversal)
  - [Search element](#search-element)
  - [Min tree element](#min-tree-element)
  - [Max tree element](#max-tree-element)
  - [PreOrder Successor](#preorder-successor)
  - [PostOrder Successor](#postorder-successor)
  - [Find parent](#find-parent)
  - [Delete node from tree](#delete-node-from-tree)


### Tree creation example

```
t := tree.CreateNode[int](15) //int node
t := tree.CreateNode[string]("abc") //string node
```

### Tree traversal
you can make tree traversal:
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)

resultAsc := t.InOrderTreeWalk(Asc)   // [4 8 20 22]
resultDesc := t.InOrderTreeWalk(Desc) // [22 20 8 4]
```

### Search element

```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

resultNil := t.Search(15) //nil
result := t.Search(12)    // tree with root 12
```

### Min tree element
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.Min() // tree with root 4
```
### Max tree element
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.Max() // tree with root 22
```

### PreOrder Successor
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.PreOrderSuccessor(t.left) // tree with root 4
```

### PostOrder Successor
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

result := t.PostOrderSuccessor(t.left) // tree with root 10
```

### Find parent
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

p := t.Parent(8) //20
```

### Delete node from tree
```
t := tree.CreateNode[int](20)
t.Insert(22)
t.Insert(8)
t.Insert(4)
t.Insert(12)
t.Insert(10)
t.Insert(14)

err := t.Delete(22) // without err
```