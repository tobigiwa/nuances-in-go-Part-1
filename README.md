# Nuances in Go: Exploring the Tiny Details. Part-1.

![](/mascot.jpeg)

I would like to shed light on the lesser-explored aspects of Go programming, those subtle behaviors that may have gone unnoticed, occasionally puzzled you, or even caused you to stumble upon solutions without fully grasping why they work as they do. Some of you might have even intuited these behaviors, finding them peculiar. Let's dive right into these nuances.

## Arguments evaluation in range loops
We're all familiar with the range loop syntax in Go, which necessitates an expression, such as *_i, v:= range exp_*, where **exp** represents the expression. The important question here is how is this expression evaluated. Let me prove that you know it or otherwise, try answering the code question below, try it for real, it would make the reading a bit more interesting (I fear I might be a terrible writer); 

```go
    arrayOfInt := []int{5, 6, 7}
    for range arrayOfInt {
    arrayOfInt = append(arrayOfInt, 10)
    }
    fmt.Println(arrayOfInt) // what is our output?
```

You've got your output?, Excellent. Now, consider this similar piece of code:

```go
    arrayOfInt := []int{5, 6, 7}
    for i := 0; i < len(arrayOfInt); i++ {
    arrayOfInt = append(arrayOfInt, 10)
    }
    fmt.Println(arrayOfInt) // What do you expect the output to be?
```
What do you think will be different between these two code snippets? I sincerely hope you haven't run them on your system yet (especially the last one).

The respective outputs are:
> {5, 6, 7, 10, 10, 10}

and

> (No output; an infinite loop)

Isn't it intriguing that two seemingly similar tasks behave so differently? Which of these codes do you find more unexpected or peculiar? Let's explore these nuances further, especially the **range**. This is the jinx of the range loop;

> It's important to note that when utilizing a range loop, the specified expression is evaluated only once, before the loop's initiation. In this context, "evaluation" refers to the process of duplicating the provided expression into a temporary variable, which is then iterated upon by the range loop. The range loop uses this temporary variable.

The temporary slice used by range remains a three-length slice. Hence, the loop completes after three iterations. Returning to the topic of the range operator, it’s crucial to understand that the behavior we’ve discussed — where the expression is evaluated just once — applies to Channels and Arrays as well.

In the scenario of the classic for-loop, the loop never terminates. This is because the `len(arrayOfInt)` expression is re-evaluated during each iteration, and since we continually add elements, we never reach a termination condition.

I’ll say it is pretty crucial to be aware of this distinction to employ Go loops effectively, Period.

Before I go, I’d like to suggest a book for your next reading choice, [Domain-Driven Design with Golang](https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450) by [Matthew Boyle](https://twitter.com/MattJamesBoyle). It’s a great read, and I’m starting it too. Stay tuned, as I’ll be sharing insights from it with you soon.


