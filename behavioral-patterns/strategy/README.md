# Strategy design pattern

## Description

The Strategy pattern uses different algorithms to achieve some specific functionality. These algorithms are hidden behind an interface and, of course, they must be interchangeable. All algorithms achieve the same functionality in a different way. For example, we could have a Sort interface and few sorting algorithms. The result is the same, some list is sorted, but we could have used quick sort, merge sort, and so on.

## Objectives
The objectives of the Strategy pattern are really clear. The pattern should do the following:

* Provide a few algorithms to achieve some specific functionality
* All types achieve the same functionality in a different way but the client of the strategy isn't affected
* 
The problem is that this definition covers a huge spectrum of possibilities. This is because Strategy pattern is actually used for a variety of scenarios and many software engineering solutions come with some kind of strategy within. Therefore it's better to see it in action with a real example.

## Structure

![alt text](../../assets/strategy-structure.png)

## Applicability

 Use the Strategy pattern when you want to use different variants of an algorithm within an object and be able to switch from one algorithm to another during runtime.

 The Strategy pattern allows you to indirectly alter the behavior of the object during runtime by associating it with different subobjects that can perform specific subtasks in different ways.

 Use the Strategy pattern when you have many similar classes that differ only in the way they perform certain behavior.

 The Strategy pattern allows you to extract the variant behavior to put it into a separate class hierarchy and combine the original classes into one, thereby reducing duplicate code.

 Use the pattern to isolate the business logic of a class from the implementation details of algorithms that may not be as important in the context of that logic.

 The Strategy pattern allows you to isolate the code, internal data and dependencies of various algorithms from the rest of the code. The various clients get a simple interface to execute the algorithms and change them during runtime.

## Pros 

* You can swap algorithms used inside an object at runtime.
* You can isolate the implementation details of an algorithm from the code that uses it.
* You can replace inheritance with composition.
Open/Closed Principle. You can introduce new strategies without having to change the context.

## Cons


* If you only have a couple of algorithms and they rarely change, there’s no real reason to overcomplicate the program with new classes and interfaces that come along with the pattern.
 Clients must be aware of the differences between strategies to be able to select a proper one.

* A lot of modern programming languages have functional type support that lets you implement different versions of an algorithm inside a set of anonymous functions. Then you could use these functions exactly as you’d have used the strategy objects, but without bloating your code with extra classes and interfaces.