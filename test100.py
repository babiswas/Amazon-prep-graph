def longest_valid_substring(str1):
    l=list(str1)
    temp=list()
    l=l[::-1]
    max=0
    count=0
    for i in range(len(str1)):
       if l[i]==')':
          temp.append(l[i])
       elif l[i]=='(':
          if temp:
             temp.pop()
             count=count+2
          else:
             if max<count:
                max=count
             count=0
    if count>max:
       max=count
    return max


if __name__=="__main__":
   print(longest_valid_substring("()(()))))"))
   print(longest_valid_substring(")()())"))
   print(longest_valid_substring("((()()()()(((())"))
   print(longest_valid_substring("((((((()"))
   print(longest_valid_substring(")))))))))("))
   print(longest_valid_substring(")("))
   print(longest_valid_substring("()(()))))"))