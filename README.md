HTesTP is a typed json end-point testing library.
its purpose is to create structured sequence of http request, with constraints applied to each request that must be satisfied before proceeding.
the sequence is a tree of http nodes.

each node represents an http request, there are two types:
    -   static node: a node that always performs the same operation with the same request body.
    -   dynamic node: a node that can be assigned callbacks that build a url query and/or a request body, it
        uses a global context to fetch stored variables.

constraints define a condition that must apply to the response data before proceeding to the next node, although its
to note that all constraints that apply to a node will be executed even if one of them fails, the constraints have several
types, which are:
    *all constraints have a <store_constraint> variant, it checks if the constraint is satisfied and stores the found variable in the global context*
    -   exist constraint: it takes a json field and a data type, and checks if that field exists in the response.
    -   match constraint: it takes a json field, a data type and an expected value, then checks if that field has the expected value.
    -   ~~not_exist constraint: TODO: future work~~

tests are constructed through a builder <test_builder>, which has an Add<node_type>() and Add<node_type>Branch(), also an Add<constraint_type>().

