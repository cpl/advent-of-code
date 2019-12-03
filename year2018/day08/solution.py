import sys


class Node():

    def __init__(self, children_count, metadata_count):
        self.children = [None for _ in range(children_count)]
        self.metadata = [None for _ in range(metadata_count)]

    def get_value(self):
        if len(self.children) == 0:
            return sum(self.metadata)
        else:
            value = 0
            for data in self.metadata:
                if data > 0 and data <= len(self.children):
                    value += self.children[data-1].get_value()
            return value

    def __str__(self):
        return "(C:%s MD:%s)" % (str(self.children), self.metadata)

    def __repr__(self):
        return self.__str__()


total_metadata = 0


def create_node(digits):
    global total_metadata

    if digits == []:
        return

    n = Node(digits[0], digits[1])
    print("NEW NODE", digits[0], digits[1])

    digits.pop(0)
    digits.pop(0)


    if len(n.children) != 0:
        for c in range(len(n.children)):
            print("NEW CHILDREN", c)
            n.children[c] = create_node(digits)

    if len(n.metadata) != 0:
        print("NEW METADATA", digits[:len(n.metadata)])
        for m in range(len(n.metadata)):
            n.metadata[m] = digits[0]
            total_metadata += digits[0]
            digits.pop(0)


    return n


digits = []
with open(sys.argv[1], "r") as fp:
    c = fp.read()
    digits = [int(d) for d in c.split(' ')]


root = create_node(digits)



print("TOTAL METADATA:", total_metadata)
print("ROOT VALUE:", root.get_value())
