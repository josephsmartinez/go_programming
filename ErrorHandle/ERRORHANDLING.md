# Error Handling

When should panic be used?
One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.