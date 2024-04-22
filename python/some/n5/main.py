def null_decorator(func):
    return func

@null_decorator
def greet():
    return 'Hello!'

print(greet())

# ------------------------------------------------------------------------

def proxy(func):
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper
