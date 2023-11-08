The Problem:
Given a schedule of n processes with a start and end time, choose the minimum number of processes from that set of n
to serve as supervisors over the processes with which their start and end times overlap.  Every process must have at least 
one overlapping process.

Solution:
supervisors is a stack and that is the returned stack of supervisors

1. Order the set of n processes by start time ascending and store it in a stack schedule with smallest time on top.
2. If _schedule_ is empty return _supervisors_
3. Pop _schedule_ and mark it as a supervisor by placing it in stack _supervisors_
3.   Pop the next process in _schedule_ and call it _i_.
   4. If i's end time is greater than last element in _supervisors_
      5. Pop _supervisors_
      6. Push i onto  _supervisors_
   6. goto step 2 -- use recursion of course


func makeSchedule(schedule Stack[Process], supervisors Stack[Processes]) Stack[Process] {
      if len(schedule) == 0 return supervisors
      
      a = Pop(schedule)
      supervisors = Push(supervisors, a)
      i = Pop(schedule)
      if i.endTime > Peek(supervisors).endTime {
            Pop(supervisors)
            supervisors = Push(supervisors, i)
      }
      makeSchedule(schedule, supervisors)
}
      
            

}



7. 
6. 
3. 
You have a bunch of two-step processes that need to be completed by a supercomputer for the first part and then a PC for the 
second part.  The supercomputer can only accept one of the two-tep processes at a time. After the first step you have 
an unlimited number of PCs that can do the secpomd step in parallel. 

Devise a schedule for a list of n two-step processes that has the minimum completion time for the whole list. 

Solution:
Order the two-step process list by PC time minus Supercomputer time in descending order. The heuristic is to order the list of processes by maximum 
idle time for the supercomputer.  This is the fastest way to order the n processes.