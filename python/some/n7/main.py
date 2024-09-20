import clang.cindex
import sys

def print_function_info(node):
    """Prints information about a function definition."""

    print(f"{node.kind}: {node.spelling}")
    for arg in node.get_arguments():
        print(f"    {arg.type.spelling} {arg.spelling}")

def analyze_code(filename):
    """Parses a C++ file and analyzes its contents."""
    index = clang.cindex.Index.create()
    translation_unit = index.parse(filename, args=['-x', 'c++', '-std=c++17']) 

    for node in translation_unit.cursor.walk_preorder():
        print_function_info(node)

if __name__ == "__main__":
    analyze_code("main.cpp")