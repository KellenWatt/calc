# calc
A simple, robust command line stack calculator repl.

Strictly designed for easy use as a calculator. I realized I didn't actually have a trivial 
calculator command on my computer (I know `bc` and `dc` exist, but I don't want to deal with them).
Using Ruby or Python works, but that's a lot of overhead for a basic calculator, and I usually don't 
need that. `calc` is designed to be simple, relatively easy to extend when necessary, and not weighted 
down by extra syntax. Basically, it's designed to be as painless as possible, while still being useful.

# Operation deferral
Whenever you have a stack with exactly one element on it, and you input a binary operator, the repl waits 
for the next number to be entered, then performs the requested operation as though the number just entered 
had been pushed on the stack before the operator. This helps facilitate straightforward operation chaining,
while making all intermediate  rsults cleanly available, as necessary.

If you have a pending operation, you can clear it with the keyword `cancel`. This stops the operation and 
leaves the stack as is.
