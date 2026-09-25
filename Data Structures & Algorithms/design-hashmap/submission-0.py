class MyHashMap:

    def __init__(self):
        self.size = 1000
        self.buckets = [None] * self.size

    def put(self, key: int, value: int) -> None:
        node = Node(key, value)
        index = key % self.size

        if self.buckets[index] is None:
            self.buckets[index] = node
            return
        
        head = self.buckets[index]

        while head is not None:
            if head.key == key:
                head.value = value
                return

            if head.next is None:
                break

            head = head.next

        head.next = node


    def get(self, key: int) -> int:
        index = key % self.size
        
        if self.buckets[index] is None:
            return -1

        curr = self.buckets[index]

        while curr is not None:
            if curr.key == key:
                return curr.value
            
            curr = curr.next
        
        
        return -1

    def remove(self, key: int) -> None:
        index = key % self.size

        if self.buckets[index] is None:
            return

        prev = None
        current = self.buckets[index]

        while current is not None:
            if current.key == key:

                if prev is None:
                    self.buckets[index] = current.next
                else:
                    prev.next = current.next

                return

            prev = current
            current = current.next


class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.next = None

# Your MyHashMap object will be instantiated and called as such:
# obj = MyHashMap()
# obj.put(key,value)
# param_2 = obj.get(key)
# obj.remove(key)