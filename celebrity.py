def celebrity_problem(M):
   indegree=[0]*len(M)
   outdegree=[0]*len(M)
   for i in range(len(M)):
     for j in range(len(M)):
        if M[i][j]==1:
           indegree[j]+=1
           outdegree[i]+=1
   for i in range(len(M)):
     if indegree[i]==len(M)-1 and outdegree[i]==0:
        return i
   return -1

if __name__=="__main__":
   print(celebrity_problem([[0,0,1,0],[0,0,1,0],[0,0,0,0],[0,0,1,0]]))