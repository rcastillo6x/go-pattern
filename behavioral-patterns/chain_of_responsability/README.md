# Chain of responsability pattern

## Description

Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

The Chain of Responsibility relies on transforming particular behaviors into stand-alone objects called handlers. In our case, each check should be extracted to its own class with a single method that performs the check. The request, along with its data, is passed to this method as an argument.

The pattern suggests that you link these handlers into a chain. Each linked handler has a field for storing a reference to the next handler in the chain. In addition to processing a request, handlers pass the request further along the chain. The request travels along the chain until all handlers have had a chance to process it.

## Structure

![alt text](../../assets/cor-structure.png)

## Applicability

 Use the Chain of Responsibility pattern when your program is expected to process different kinds of requests in various ways, but the exact types of requests and their sequences are unknown beforehand.

 The pattern lets you link several handlers into one chain and, upon receiving a request, “ask” each handler whether it can process it. This way all handlers get a chance to process the request.

 Use the pattern when it’s essential to execute several handlers in a particular order.

 Since you can link the handlers in the chain in any order, all requests will get through the chain exactly as you planned.

 Use the CoR pattern when the set of handlers and their order are supposed to change at runtime.

 If you provide setters for a reference field inside the handler classes, you’ll be able to insert, remove or reorder handlers dynamically

## Pros 

* You can control the order of request handling.
  
* Single Responsibility Principle. You can decouple classes that invoke operations from classes that perform operations.
  
* Open/Closed Principle. You can introduce new handlers into the app without breaking the existing client code.

## Cons

* Some requests may end up unhandled