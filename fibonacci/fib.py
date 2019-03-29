def fib(x):
    if x < 0:
        print("ERR: a non-positive integer was passed", x)
        return x
    
    if x == 1 or x == 2 or x == 0:
        return 1

    return fib(x-1) + fib(x-2)

def fib_test(x):
    print("computing fib of x:", x, fib(x))

if __name__ == "__main__":
    for x in [0,1,2,3,4,5,6,7,8,10, -5]:
        fib_test(x)

