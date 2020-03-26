# -*- coding: utf-8 -*-

from breadth_first_search import bfs


def bfs_connected_components(graph):
    """ Finds all the connected subgraphs in a given undirected graph.

    It does so by iterating over all non-visited elements in a graph
    and running bfs over it.

    Complexity: O(n+m)

    Args:
        graph: instance of src.graph.Graph class which encapsulates all
            edges, vertices as well as values for edges and vertices.

    Returns:
        list, of lists of vertexes which are connected. Format [[vertex,..]]
    """
    subgraphs = []
    explored_vertices = []

    for vertex in graph.get_vertices():
        if vertex in explored_vertices:
            continue
        visited_vertices = bfs(graph, vertex)

        explored_vertices.extend(visited_vertices)
        subgraphs.append(visited_vertices)

    return subgraphs
