# Singly Linked Lists

class SinglyNode:
    def __init__(self,val,next=None):
        self.val = val
        self.next= next
    def __str__(self):
        return str(self.val)



Head = SinglyNode(1)
A= SinglyNode(2)
B= SinglyNode(3)
C= SinglyNode(4)

Head.next = A
A.next = B
B.next = C

print(Head)



# Traverse the List - O(n)
curr = Head
while curr:
    print(curr)
    curr=curr.next


# Display Linked List - O(n)
def display(head):
    curr=head
    elements=[]
    while curr:
        elements.append(str(curr.val))
        curr = curr.next
    print('->'.join(elements))

display(Head)

# Search for node value - o(n)
def search(head,val):
    curr=head
    while curr:
        if val==curr.val:
            return True
        curr=curr.next
    return False

search(Head,7)


# Doubly Linked Lists

class DoublyNode:
    def __init__(self,val,next=None,prev=None):
        self.val= val
        self.next = next
        self.prev = prev
    def __str__(self):
        return str(self.val)
    

head=tail=DoublyNode(1)
print(tail)


# Display - O(n)
def display(head):
    curr = head
    elements =[]
    while curr:
        elements.append(str(curr.val))
        curr=curr.next
    print('<->'.join(elements))

display(head)

def insert_at_beginning(head,tail,val):
    new_node = DoublyNode(val,next=head) # val = value , next=head means next of new node is head
    head.prev = new_node  # as doubly linked list , prev also needs to be set 
    return new_node , tail

def insert_at_end(head,tail,val):
    new_node= DoublyNode(val,prev=tail)
    tail.next = new_node
    return head, new_node

