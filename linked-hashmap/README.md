# Linked Hashmap

The interview question here is to write a linked hashmap. A linked hashmap is blend of a linked list and a hash map. As a hash map, it is a store of key-value pairs. The keys will look up the values in constant time.

In addition, the keys also available as an ordered list. The are in the order of the most recently used and descending. In some hash map implementations, the keys are stored in an array. However, it is costly to re-order the array with each and every use. So the optimization is to store the keys in a linked list (doubly linked). This way, when a key is used, it is removed and inserted to the end of the list in constant time.
