# Read code and the terminal even when AI writes it

AI writes code faster than you ever will. That does not mean you can stop reading it.

The developers who get burned by AI — who ship bugs, introduce security holes, or break production — are almost always the ones who stopped reading the output and started just running it. The developers who thrive are the ones who use AI to generate and then use their own eyes to verify. That gap is entirely determined by whether you can read code and understand what a terminal is telling you.

## AI output is a first draft, not a final answer

AI is confident. It will produce clean-looking, well-formatted code that does the wrong thing without hesitation. It will suggest a shell command that deletes data, misconfigure a service, or introduce a subtle logic error that only shows up in edge cases. It does not know it is wrong. It will not warn you. The only defense you have is reading what it produced before you run it.

This is not a failure of AI. It is the nature of probabilistic generation. The model is guessing based on patterns, not reasoning from ground truth. Your job is to be the reasoning layer on top.

## Reading code is different from writing it

You do not need to be able to write production-quality code from scratch. You need to be able to read a function and understand what it does, trace a loop, follow a conditional, spot a variable that is never initialized, recognize when a command will do something destructive. That is a much lower bar than writing, and it is the bar that actually matters when you are working with AI.

Reading code is a skill you build by reading code. Start with the output AI gives you. Read every line before you run it. Ask yourself what each part does. Use AI itself to explain sections you do not understand — but read the explanation too.

## The terminal is not optional

AI can suggest commands. It can run them for you. It can interpret the output. But if you do not understand what the terminal is showing you, you cannot catch errors, you cannot debug, and you cannot tell when something went wrong silently. A process that exits with code 1 and no output is a failure. A permission denied error buried in fifty lines of log output is a problem. A command that completes successfully but writes to the wrong location is a disaster. None of these are obvious unless you are reading.

Linux terminal output follows patterns you can learn. Exit codes, error prefixes, file paths, process states, permission strings — these are a vocabulary. You do not need to memorize all of it. You need enough to know when something looks wrong and enough curiosity to look it up when it does.

## Why this matters more in the AI era, not less

When AI is doing more of the work, the surface area of things that could go wrong expands. An agent running a multi-step task will make a dozen decisions you never reviewed. Each one is an opportunity for a mistake to compound. The only way to catch them is to check the output, read the diffs, understand what changed and why.

The developers who will be most valuable are not the ones who can write the most code — AI handles that now. They are the ones who can direct AI effectively, catch its mistakes early, and understand the system well enough to know when something is wrong. That requires reading code and reading the terminal. There is no shortcut around it.
