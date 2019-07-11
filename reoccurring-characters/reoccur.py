def reoccur(s):
    occur = set()
    for i in range(len(s)):
        if s[i] in occur:
            return s[i]
        else:
            occur.add(s[i])
    return False

def reoccur_no_mem(s):
    if len(s) == 0:
        return False
    
    last = 0
    for i in range(len(s)-1):
        for j in range(i+1, len(s)):
            if s[i] == s[j]:
                if last == 0 or j < last:
                    last = j
    
    if last == 0:
        return False
    else:
        return s[last]
