# -*- coding: utf-8 -*-

class Graph:
    """ Base data structure to hold graph data using adjacency lists.
    Uses the adjacency list paradigm to save on space.
    """

    def __init__(self, directed=False):
        # Marks whether or not the graph is directed. Defaults to False.
        self.directed = directed

        # Dict structure to hold the graph. Format {tail: {head: value}}
        self.table = {}

        # Dict structure to hold the vertices values.
        self.values = {}

    def split_edge(self, edge):
        tail = edge[0]
        head = edge[1]
        if len(edge) == 2:
            value = True
        else:
            value = edge[2]
        return (tail, head, value)

    def add_vertex(self, vertex):
        if vertex not in self.table:
            self.table[vertex] = {}

    def add_edge(self, edge):
        (tail, head, value) = self.split_edge(edge)

        self.add_vertex(head)
        self.add_vertex(tail)

        if head == tail:
            return

        self.table[tail][head] = value
        if self.directed == False:
            self.table[head][tail] = value

    def get_vertices(self):
        return self.table.keys()

    def get_edges(self):
        output = []
        for tail, edges in self.table.iteritems():
            for head, value in edges.iteritems():
                if self.directed == False and (head, tail, value) in output:
                    continue
                output.append((tail, head, value))
        return output

    def adjacent(self, tail, head):
        if tail in self.table:
            if head in self.table[tail]:
                return True
        return False

    def neighbours(self, vertex):
        if vertex not in self.table:
            return []
        else:
            return self.table[vertex].keys()

    def get_edge_value(self, edge):
        (tail, head, __) = self.split_edge(edge)
        if self.adjacent(tail, head):
            return self.table[tail][head]
        else:
            return None

    def set_edge_value(self, edge, newValue):
        (tail, head, value) = self.split_edge(edge)

        if self.adjacent(tail, head):
            self.table[tail][head] = newValue
        if self.adjacent(head, tail):
            self.table[head][tail] = newValue

    def get_vertex_value(self, vertex):
        if vertex in self.values:
            return self.values[vertex]
        return None

    def set_vertex_value(self, vertex, newValue):
        self.values[vertex] = newValue

    def remove_edge(self, edge):
        (tail, head, value) = self.split_edge(edge)
        if tail in self.table:
            if head in self.table[tail]:
                del self.table[tail][head]
        if self.directed == False:
            if head in self.table:
                if tail in self.table[head]:
                    del self.table[head][tail]

    def remove_vertex(self, vertex):
        if vertex not in self.table:
            return
        del self.table[vertex]

        for tail, edges in self.table.iteritems():
            if vertex in edges:
                del self.table[tail][vertex]

    def rename_vertex(self, old, new):
        if old not in self.table:
            return

        if new in self.table:
            self.table[new] = dict(self.table[old].items() +self.table[new].items())
        else:
            self.table[new] = self.table[old]
        del self.table[old]

        for tail, edges in self.table.iteritems():
            if tail == new and new in edges:
                del self.table[tail][new]
            if old in edges:
                if tail != new:
                    self.table[tail][new] = self.table[tail][old]
                del self.table[tail][old]

    @staticmethod
    def build(vertices = [], edges = [], directed = False):
        g = Graph(directed)
        for vertex in vertices:
            g.add_vertex(vertex)
        for edge in edges:
            g.add_edge(edge)
        return g