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
        "http://httpbin.org/json",
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
        },
        nil).AddExistConstraint("docs[2].author_key[0]", models.TypeString)

status := builder.Run()
fmt.Printf("Test Passed: %v", status)

//	option to print the contents of the test
builder.PrintList()

//program end
```