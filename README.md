# GoCache

This repository contains an implementation of an LRU (Least Recently Used) cache in Go. An LRU cache is a data structure that caches a limited number of items and discards the least recently used item when the cache is full and a new item is added.

## Getting Started

To get started, you'll need to have Go installed on your machine. Once you have Go installed, clone the repository and navigate to the project directory:

```
$ git clone https://github.com/rajatmjain/GoCache.git
$ cd GoCache
```

You can then run the code using the following command:

```
$ go run main.go
```

This will run the LRU cache implementation and output the results to the console.

## Usage

The main.go file contains an example of how to use the LRU cache implementation. Here's a brief overview of the code:

```go
cache := NewCache()
initialSample := []string{"Real Madrid", "Bayern Munich", "Manchester City", "Barcelona", "PSG"}
testSample := []string{"Liverpool", "Manchester City", "Bayern Munich", "Real Madrid", "Chelsea", "Barcelona", "PSG"}

for _, word := range initialSample {
	cache.Check(word)
	cache.Display()
}

for _, word := range testSample {
	cache.Check(word)
	cache.Display()
}

cache.Flush()
cache.Display()

```

The code creates a new cache instance using the `NewCache()` function. It then populates the cache with an initial sample of items and tests the cache with a test sample of items.

The `Check()` method is used to add an item to the cache or retrieve an existing item from the cache. If the item is already in the cache, it is moved to the front of the cache (i.e., the most recently used position). If the cache is full, the least recently used item is removed from the cache.

The `Display()` method is used to display the current state of the cache.

The `Flush()` method is used to clear the cache.

The `Len()` method returns the number of items in the cache.

The `Contains()` method checks if an item is in the cache.

The `Items()` method returns a copy of the hash map used by the cache.

The `Get()` method retrieves an item from the cache by its key. If the item is found, it is moved to the front of the cache and its value is returned. If the item is not found, an error is returned.

The `SetCacheSize()` method is used to set the maximum size of the cache. If the new size is smaller than the current size, the least recently used items are removed from the back of the queue until the queue is the new size. If the new size is larger than the current size, the cache size is updated but the queue is not modified.

## Implementation Details

The LRU cache implementation is based on a doubly linked list and a hash table. The doubly linked list is used to maintain the order of the items in the cache, with the most recently used item at the front of the list and the least recently used item at the back of the list. The hash table is used to provide fast access to items in the cache.

The Queue struct represents the doubly linked list. It contains a Head and a Tail node, which are used as sentinels to simplify the code. The Length field is used to keep track of the number of items in the list.

The Cache struct contains a Queue and a Hash. The Hash is a map that maps strings (i.e., the cache keys) to nodes in the doubly linked list.

The `Check()` method first checks if the item is already in the cache by looking it up in the Hash. If the item is already in the cache, it is moved to the front of the list by calling the `Remove()` method to remove the item from its current position and the `Add()` method to add the item to the front of the list. If the item is not in the cache, a new node is created and added to the front of the list using the `Add()` method. If the cache is full, the least recently used item is removed from the cache by calling the `Remove()` method on the node at the back of the list.

The `Add()` method adds a new node to the front of the list by updating the pointers of the adjacent nodes.


