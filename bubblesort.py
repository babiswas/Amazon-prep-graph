def bubble_sort(arr):
   for i in range(len(arr)):
     for j in range(len(arr)-i-1):
        if arr[j]>arr[j+1]:
           arr[j],arr[j+1]=arr[j+1],arr[j]

if __name__=="__main__":
   arr=[12,1,3,4,13,7,0]
   bubble_sort(arr)
   print(arr)
       
    