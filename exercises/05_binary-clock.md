# Exercise 05

[Go Back](README.md)

## Description

Create a binary clock that prints the binary representation of hours, minutes, and seconds — one line per second.

Each component is split into two groups: **tens** and **units**, matching how a real digital clock works (e.g. 23:59:59 → tens=2 units=3, tens=5 units=9, tens=5 units=9). Each group is represented in binary with a fixed bit width:

- **Hours**: 6 bits total — 2 bits for tens (max 2), 4 bits for units (max 9)
- **Minutes**: 7 bits total — 3 bits for tens (max 5), 4 bits for units (max 9)
- **Seconds**: 7 bits total — 3 bits for tens (max 5), 4 bits for units (max 9)

Unused bit positions are filled with underscores (`X`). Format each line as:

```
HHTTT UUUU - MMTTT UUUU - SSTTT UUUU
```

Where `TTT`/`TT` are the binary tens bits and `UUUU` are the binary units bits.

You can count seconds from 00:00:00 or use any other starting time.

## Expected Output

First 10 seconds starting from 00:00:00:

```bash
XX00 0000 - X000 0000 - X000 0000
XX00 0000 - X000 0000 - X000 0001
XX00 0000 - X000 0000 - X000 0010
XX00 0000 - X000 0000 - X000 0011
XX00 0000 - X000 0000 - X000 0100
XX00 0000 - X000 0000 - X000 0101
XX00 0000 - X000 0000 - X000 0110
XX00 0000 - X000 0000 - X000 0111
XX00 0000 - X000 0000 - X000 1000
XX00 0000 - X000 0000 - X000 1001
```
