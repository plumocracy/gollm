### Goals/Scope of this project

The current offerings for working with LLM's in go are pretty limited and as of writing the only other 
library that I can find that seems sensible seems to have fallen out of maintenance. My goal with this
package is to write a simple, sensible API for working with various different LLM calls primarily on the CLI.

### Scope
The current scope I'm targeting is something like this: 
1: Connect and query various models, including openrouter for an open model support.
2: Store responses either locally with SQLite or to any remote database.
3: Query Multiple Models simultaneously and get their results.
4: Built in memory system/structure.
5: SSE streaming.
