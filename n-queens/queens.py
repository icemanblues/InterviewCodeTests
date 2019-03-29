# q is a list of queens. the element index is the column, the number is the row
# i is the index (column) we are trying to solve
def queens(q, i):
    if i >= len(q):
        return True

    for j in range(len(q)):
        # insert the trial value. if good, keep on recursing
        # in the recurse cannot complete, then remove this and move on
        q[i] = j

        # check rows and diagonals
        valid = valid_queens(q,i,j)
        if not valid:
            continue

        # the recursion
        ok = queens(q, i+1)
        if ok:
            return True
    
    # need to return something. if we tried everyone and we couldn't solve it
    return False


def valid_queens(q,i,j):
    for k in range(i):
        # this is the row check
        if q[k] == j:
            return False
        
        # this is the diagonal check
        if abs(i-k) == abs(j-q[k]):
            return False
    
    return True


if __name__ == "__main__":
    print("solving queens")
    q = [0,0,0,0,0,0,0,0]
    #q = 1
    ok = queens(q, 0)
    if ok:
        print("we solved it.", q)
    else:
        print("no solution found")
