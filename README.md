# windows_lock

A small application that sits in a Windows tray and locks screen after the system is not used for some time.
In some rare occasions Windows won't lock by built in tool because of something resets the system idle timer.
In my case once it was a running in the background OneDrive, and in another case I couldn't Identify what is doing that.
All USB devices were disconnected and most of the office software weren't running but something still resets the idle timer.
And Windows wouldn't lock in 6 out of 10 times.
It was faster to create the tool then figthing the problem.

The tool skips few idle resets in 5 second period and will lock the screen anyway.
If more resets happens the timer resets and and the screen won't be locked.
