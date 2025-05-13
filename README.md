# HTesTP

**HTesTP** is a **typed JSON endpoint testing library**.

Its purpose is to create **structured sequences of HTTP requests**, with **constraints applied to each request** that must be satisfied before proceeding.

The sequence is modeled as a **tree of HTTP nodes**.

---

## Node Types

Each node represents an HTTP request. There are two types:

- **Static Node**  
  A node that always performs the same operation with the same request body.

- **Dynamic Node**  
  A node that can be assigned callbacks to build a URL query and/or request body.  
  It uses a **global context** to fetch stored variables.

---

## Constraints

Constraints define conditions that must apply to the response data before proceeding to the next node.

- **Important:**  
  All constraints attached to a node will be **executed**, even if one of them fails.

Constraint types:

- **Exist Constraint**  
  Takes a JSON field and a data type, and checks if that field exists in the response.

- **Match Constraint**  
  Takes a JSON field, a data type, and an expected value, then checks if the field has the expected value.

- **Store Variant**  
  *All constraints have a `<store_constraint>` variant*, which checks if the constraint is satisfied and **stores the found value into the global context**.

- **Not Exist Constraint**  
  `~~TODO: future work~~`

---

## Test Construction

Tests are constructed using a **TestBuilder**, which provides:

- `Add<node_type>()`
- `Add<node_type>Branch()`
- `Add<constraint_type>()`

---

##  Example

```go
//start of the program
builder := test_builder.CreateNewBuilder()
builder.
  AddStaticNode(
    "https://httpbin.org/json",
    models.GET,
    nil,
  ).
  AddMatchStoreConstraint(
    "slideshow.author",
    "Yours Truly",
    models.TypeString,
    "authorName",
  ).
  AddDynamicNode("https://openlibrary.org/search.json", models.GET,
    func(m *map[string]models.TypedVariable) map[string]string {
      key := (*m)["authorName"]
      Map := map[string]string{}
      Map["q"] = key.Value.(string)
      return Map
    }, nil).
  AddExistConstraint("docs[2].author_key[0]", models.TypeString)

//	each operation builds a new branch, with the parent being the previous builder's current branch
//	the builder doesn't proceed to any branch and stays at current
branch1 := builder.AddStaticNodeBranch("http://httpbin.org/b1", models.GET, nil)
branch2 := builder.AddStaticNodeBranch("http://httpbin.org/b2", models.GET, nil)
branch3 := builder.AddStaticNodeBranch("http://httpbin.org/b3", models.GET, nil)

//	branch objects are also test builders, and can be used in the same manner
branch1.AddStaticNode("http://httpbin.org/b11", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)
branch2.AddStaticNode("http://httpbin.org/b22", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)
branch3.AddStaticNode("http://httpbin.org/b33", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)

//	WARNING: when running the test, you must always start from the root builder
status := builder.Run()
fmt.Printf("Test Passed: %v", status)

//	option to print the contents of the test
//	each branch will print only the contents of its branch

branch1.PrintList()
branch2.PrintList()
branch3.PrintList()

//	example of the unsafe function SetBranchTo(), in most cases it shouldn't be used, as it can lead to unintentional unallocation of nodes
builder.SetBranchTo(branch3)

//program end
```

[![Go](https://github.com/ZiyadHQ/H_TesT_P/actions/workflows/go.yml/badge.svg?branch=main&event=check_suite)](https://github.com/ZiyadHQ/H_TesT_P/actions/workflows/go.yml)

#cloc output:
See [cloc.md](./cloc.md) updated line counts after every merge into main
