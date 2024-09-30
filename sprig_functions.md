
# `sprig` functions

**1. String Manipulation**

- `trim`, `trimSuffix`, `trimPrefix`:  Trimming whitespace or specific characters from strings
- `replace`, `replaceAll`: Replacing substrings within a string
- `title`, `upper`, `lower`: Case conversions
- `split`, `join`: Splitting and joining strings based on a delimiter
- `substr`, `hasPrefix`, `hasSuffix`: Substring extraction and checking

**2. Math and Numeric Operations**

- `add`, `sub`, `mul`, `div`:  Basic arithmetic operations
- `addf`, `subf`, `mulf`, `divf`: Arithmetic operations for floats
- `max`, `min`, `abs`: Finding maximum, minimum, and absolute values
- `randInt`, `randFloat`: Generating random integers and floats

**3. Date and Time Handling**

- `now`: Get the current time
- `date`: Format a date according to a specified layout
- `dateModify`: Add or subtract durations from a date
- `htmlDate`, `htmlDateInZone`: Format dates for HTML output

**4. Collections and Lists**

- `list`: Create a list from multiple values
- `first`, `last`, `index`: Access elements within a list
- `slice`: Extract a portion of a list
- `append`, `prepend`: Add elements to the beginning or end of a list
- `sort`, `reverse`: Sorting and reversing lists

**5. Dictionaries and Maps**

- `dict`: Create a dictionary/map from key-value pairs
- `get`, `set`: Access and modify values in a dictionary
- `hasKey`, `keys`, `values`: Check for key existence, get keys or values

**6. Type Conversions**

- `atoi`, `int`, `int64`: Convert strings to integers
- `toString`: Convert values to strings
- `toBool`: Convert values to booleans

**7. Control Flow and Logic**

- `if`, `else`: Conditional logic within templates
- `and`, `or`, `not`: Logical operations
- `eq`, `ne`, `lt`, `gt`: Comparison operators

**8. Other Useful Functions**

- `default`: Provide a default value if a variable is empty
- `coalesce`: Return the first non-empty value from a list
- `ternary`: Implement a ternary operator (condition ? valueIfTrue : valueIfFalse)
- `uuidv4`: Generate a UUID (Universally Unique Identifier)
- `b64enc`, `b64dec`: Base64 encoding and decoding
