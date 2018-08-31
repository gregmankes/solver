# Systems of equations evaluator

This is a command line client to evaluate system of linear equations using top down recursive memoization.

The input must be in a file, and be of the following format:
```
# (LHS = RHS)

offset = 4 + random + 1
location = 1 + origin + offset
origin = 3 + 5
random = 2
```

## Assumptions
The dag will eventually resolve to a single value.

## Dependencies

```
Docker
Make
```
That's it!

## Running the tool
The file that you wish to input to the evaluation tool must be in the same directory as the main.go executable. 

You may run the tool with the following:
```
make run filename=<your_filename>
```

## Libraries used
None! Everything used in the source code is built in to golang.