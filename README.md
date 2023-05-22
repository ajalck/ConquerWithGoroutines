# ConquerWithGoroutines
***

> ## Concurrency and Parallelism
===
>> *Concurrency* :- **Dealing** multiple tasks at a time    
>> *Parallelism* :- **Doing** multiple tasks at a time

> ## Goroutines
===
>>Goroutines are actually designed to perform multiple tasks or methods _concurrently_ in a go program .

>> ### Working of goroutines 
- - 
* If we are not using more goroutines in our program , go will only work as a single threaded programing language . The step by step      execution will only be done and it would be by the main goroutine. 
* So goroutines are actually destributing tasks to multiple threads.
* Other go routines are executed by the main goroutine . So if the execution of main goroutine get completed, it will terminate. And it will not wait for other goruotines to execute. 
```
func main(){
    go say("Hello")
    go say("Hi")
}
```
in this snippet , main goroutine will be terminated without executing anything.
* To avoid this we need a wait in main goroutine. 
>Two methods to achieve this 
>>1. call `time.Sleep()` and add a time to sleep the main goroutine for that time. And within this time , other goroutines will be executed. But its not recomented because, we can't assume the time that need for the execution of each goroutine. So its not efficient.
>>2. Synchronization mechanism.

>> ### Synchronization mechanism

* We can synchronize the working of goroutines by **go channels** and **waitGroups**.
>* `sync.WaitGroups{}` allows us to complete the execution of a goroutine or a group of goroutines before the termination of the main goroutine.
>* channels enable a controlled communication and data sharing between goroutines.

**_By using synchronization mechanism we can use the groutines efficiently_**

>> ## Race Condition

* Race Condition occures when multiple goroutines come to use one memory. 
>>> ![Alt text](./images/goConcurrency.readme.png)
* So to avoid this we can use `sync.Mutex{}`
* Mutex is **locking** the memory space for a goroutine during its execution if its trying to write into that memory.And **unclocks** the memory after writing. 
>> In other words Mutex allows only one thread to use one memory space at a time.
* `sync.RWmutex{}` --> Its not a widely using method . Eventhough , its concept is ,
>> Its allowing multiple threads to perform read operation on a memory but locking the read operation when a thread come to write on that memory._But there is a misunderstanding that RWmutex{} will block a memory to read for a goroutine if multiple goroutines comes at a time.As well as it only blocks the *same memory* which locked for writing from the reading.
