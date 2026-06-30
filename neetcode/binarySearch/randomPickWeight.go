package binarysearch

import "math/rand"

type randomPickWeight struct {
	prefix []int
}

func newRandomPickWeight(w []int) randomPickWeight {
	prefix := []int{0}
	for _, weight := range w {
		prefix = append(prefix, prefix[len(prefix)-1]+weight)
	}

	return randomPickWeight{prefix: prefix}
}

func (rp *randomPickWeight) pickIndex() int {
	target := float64(rp.prefix[len(rp.prefix)-1]) * rand.Float64()
	l, r := 1, len(rp.prefix)

	for l < r {
		mid := (l + r) >> 1
		if float64(rp.prefix[mid]) <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l - 1
}

/*
Let's slow way down and build the intuition from scratch, no code yet.

## The goal

You have items with different weights, and you want to randomly pick an item such that
**heavier items get picked more often**. Like a weighted lottery.

`w = [1, 3, 2]` means:
- item 0 has weight 1
- item 1 has weight 3
- item 2 has weight 2

Total weight = 6. So intuitively, item 1 should be picked 3/6 = 50% of the time,
item 2 should be picked 2/6 ≈ 33% of the time, and item 0 should be picked 1/6 ≈ 17% of the time.

## The "raffle tickets" analogy

Imagine a raffle. Each item gets a number of raffle tickets equal to its weight:

- item 0 gets **1** ticket
- item 1 gets **3** tickets
- item 2 gets **2** tickets

Total tickets in the bag = 6.

Now imagine numbering all the tickets 1 through 6 and writing on each one which item it belongs to:

```
Ticket #:   1     2     3     4     5     6
Owner:    item0  item1 item1 item1 item2 item2
```

If you draw **one random ticket number** out of 6,
the chance you draw an item-1 ticket is 3 out of 6 — exactly matching its weight.
That's the entire idea. **More tickets = more chances to be drawn = higher probability**,
and the number of tickets is just the weight.

## Now connect this to "buckets" on a number line

Instead of discrete tickets numbered 1-6, the code uses a continuous number line from 0 to 6,
and instead of giving each item individual numbered tickets,
it gives each item a **contiguous chunk (bucket)** of that line, with chunk size = its weight:

```
0        1              4        6
|--------|--------------|--------|
 item 0     item 1          item 2
 (width 1)  (width 3)       (width 2)
```

- item 0 owns the stretch `[0, 1)` — 1 unit wide
- item 1 owns the stretch `[1, 4)` — 3 units wide
- item 2 owns the stretch `[4, 6)` — 2 units wide

This is exactly the same as the raffle analogy,
just using a continuous line instead of discrete numbered tickets.
The **width of each bucket equals its weight**.

## Why throwing a dart works

Now, instead of drawing a random ticket,
imagine you close your eyes and throw a dart at this number line,
anywhere between 0 and 6, completely uniformly (every point equally likely to be hit).

Since item 1's bucket takes up half the total line length (3 out of 6),
your dart has a 50% chance of landing in item 1's territory —
even though you didn't pick "item 1" directly,
the *size of its bucket* makes it more likely to be hit.
Same logic as before: bigger bucket = more likely to catch the dart = higher probability of being picked.

This dart throw is exactly what

```go
target := total * rand.Float64()
```

does — it generates a uniformly random point somewhere on that 0-to-6 line. That point is the "dart."

## The binary search is just "which bucket did the dart land in?"

Once you have the dart's landing spot (`target`),
you need to figure out *whose* bucket it's in.
That's all the binary search is doing — checking the bucket boundaries
(`prefix = [0, 1, 4, 6]`) to find which range `target` falls between.

If `target = 3.6`, you can see by eye it's between `1` and `4`, so it's in item 1's bucket.
The binary search just does this boundary-finding efficiently instead of you eyeballing it.

## Summary of the mental model

| Concept | What it means |
|---|---|
| Weight | How many "raffle tickets" / how wide a slice of the number line an item owns |
| Bucket | The contiguous range on the number line that belongs to one item, with width = its weight |
| `target` | A uniformly random "dart throw" landing somewhere on the full number line |
| Binary search | Finds which bucket the dart landed in |
| Why it's weighted correctly | Bigger weight → wider bucket → higher chance the random dart lands inside it |

*/
