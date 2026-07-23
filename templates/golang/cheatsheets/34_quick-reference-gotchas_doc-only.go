// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 34. Quick Reference and Go-Specific Gotchas
// ===================================================================

/*
Quick reference: choose the operation before recalling punctuation.

Terms used below:
- Lower bound: first sorted index whose value is at least the target.
- Min-heap: priority queue that removes the smallest item first.
- heap.Interface: the five Len/Less/Swap/Push/Pop methods required by
  container/heap.
- Backtracking path: the current partial choices in recursive search.

Task                         Go 1.21+ form
Empty slice                  values := []T{} or make([]T, 0, capacity)
Append / pop stack           s = append(s, x) / s = s[:len(s)-1]
Independent slice copy       clone := slices.Clone(s)
Map lookup                   value, exists := m[key]
Readable set membership      set := map[T]bool{}; if set[value] { ... }
Key-only set                 set := map[T]struct{}{}; _, ok := set[value]
Sorted copy                  clone := slices.Clone(s); slices.Sort(clone)
Lower bound                  sort.Search(len(s), func(i int) bool { return s[i] >= x })
Queue                        queue + head index; mark visited before enqueue
Min-heap                     implement heap.Interface; heap.Push/heap.Pop use any
Mutable string assembly      strings.Builder or []byte
Unicode-safe characters      runes := []rune(text)
Integer min / max            min(a, b) / max(a, b)
Clear bit                    value &^= 1 << bit
Save backtracking path       answer = append(answer, slices.Clone(path))

Go-specific gotchas: symptom -> interview-safe default.

len(text) counts bytes        -> use []rune or utf8 helpers for Unicode
copy := original slice       -> shares a backing array; use slices.Clone
append(slice, x) ignored     -> assign the returned slice header
missing map key looks valid  -> use comma-ok when zero and missing differ
write to nil map             -> initialize with make or a literal first
range value mutation         -> mutate collection[index], not the copied value
map output changes order     -> collect and sort keys when order matters
math.Max on integers         -> use built-in max on Go 1.21+
~value for bitwise NOT       -> Go spells unary complement ^value
close channel from receiver  -> the sending side normally owns close
interface holding typed nil  -> dynamic type makes the interface non-nil
heap.Pop result stays any     -> value := heap.Pop(&h).(YourConcreteType)
*/
