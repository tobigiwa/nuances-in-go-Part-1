# Nuances in Go: Exploring the Tiny Details

I want to shed light on the lesser-explored aspects of Go programming, those subtle behaviors that may have gone unnoticed, occasionally puzzled you, or even caused you to stumble upon solutions without fully grasping why they work as they do. Some of you might have even intuited these behaviors, finding them peculiar. Let's dive right into these nuances.

## Arguments evaluation in range loops
We all know that the range loop syntax requires an expression; *_i, v := range exp_*, **exp** is the expression. The important question here is how is this expression evaluated. Let me prove that you know it or otherwise, try answer the code question below, try it for real, it would make the reading a bit more interesting (i fear i might be a terrible writer); 

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
What do you think will be different between these two code snippets? I sincerely hope you haven't run them on your system yet.

The respective output are:
> {5, 6, 7, 10, 10, 10}

> (No output; an infinite loop)

Isn't it intriguing that two seemingly similar tasks behave so differently? Which of these codes do you find more unexpected or peculiar? Let's explore these nuances further, especuially the **range**. This is the jinjx of the range loop;

> It's important to note that when utilizing a range loop, the specified expression is evaluated only once, before the loop's initiation. In this context, "evaluation" refers to the process of duplicating the provided expression into a temporary variable, which is then iterated upon by the range loop. The range loop uses this temporary variable.

The temporary slice used by range remains a three-length slice. Hence, the loop completes after three iterations. Returning to the topic of the range operator, it's crucial to understand that the behavior we've discussed—where the expression is evaluated just once—applies to channels and array as well.

In this scenario of the classic for-loop, the loop never terminates. This is because the len(s) expression is re-evaluated during each iteration, and since we continually add elements, we never reach a termination condition. 

I'll say it pretty crucial to be aware of this distinction in order to employ Go loops effectively.Period.



