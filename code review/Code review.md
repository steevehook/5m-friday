# Code Reviews (Talk)

## Correctness (40%)

- Clone branch on local machine
- Understand the requirements
- Make sure it works
- Make sure all the tests & code style check pass

## Scope limited (+15%)

- Don't do more than required
- avoid mixing refactoring with implementations
- Designate refactoring in separate PRs
- Do not suggest out of scope changes
- Leave TODO(s) for future implementations

## Simplicity (+15%)

- Keep things simple (KISS)
- Avoid reinventing the wheel, use trusted solutions
- Avoid over engineering (implementing patterns for the sake of patterns)
- Avoid using large scale technologies if not needed
- avoid anticipating every possible future situation
- Do not fix what's not broken
- Limit the team to a certain feature set
- Avoid using complicated/sophisticated language constructs

## Readability & code style (+10%)

- Stick to team code style & conventions
- Avoid LoC optimization
- Name things properly
- Keep things small & concise with a single responsibility

## Consistency (+5%)

- Stick to the current project structure
- Stick to code style & conventions
- new changes or adoptions must be discussed team wide

## Testability (+10%)

- Write code which is easily testable
- Always cover code with unit & integration tests
- Don't tie the code base to tests

## Good documentation (+5%)

- Always write descriptive docs (API docs, code docs)
- Keep docs up to date
- 

## Consider

- use an adequate non-offensive language - sometimes it does not matter what you say but how you say it
- talk things through
- ask for advice before implementing
- critique the idea not the person
- be persistent and ask for review till it's merged
- be detailed about the suggested change
- admit yor mistakes and don't be too proud even if you're a senior

## Avoid

- long code reviews (comments) or long PRs (LoC)
- enforcing personal preferences
- having biased or personal opinions in arguments
- micro optimizations, unless needed
- comments about naming, separation, spaces, reordering and e.t.c
- trading in delivery speed for code quality

### Question:
Does your comment benefit your team mate, or it's just a pointless statement to show how smart you are?
