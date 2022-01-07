"""
rec2 = [1,1,3,3]
rec1 = [0,0,2,2]
"""
# better memory
def intersect(rec1, rec2):
    if rec2[1] >= rec1[3] or rec2[3] <= rec1[1]:
        return False
    if rec2[0] >= rec1[2] or rec2[2] <= rec1[0]:
        return False
    return True
#  better runtime
def intersect(rec1, rec2):
    a,b,c,d=rc1
    e,f,g,h=rc2
    return min(d,h)>max(b,f)and min(c,g)>max(a,e)


