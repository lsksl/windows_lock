# windows_lock
It is an educational project. 

A small application that sits in a Windows tray and locks a screen after a system is not used for some time.
In some rare occasions Windows won't lock by the built in tool because of something resets the system idle timer.
In my case once it was a running in the background OneDrive, and in another case I couldn't Identify the cause.
All USB devices were disconnected and most of the office software weren't running but something still reset the idle timer.
And Windows wouldn't lock in 6 out of 10 times.
It was faster to create the tool then figthing the problem.

The tool skips few idle resets in 5 second period and will lock the screen anyway.
If more resets happen the timer resets and and the screen won't be locked.

Need to add a functionality that pauses the app when a screen is locked, but it works so far w/o issues.

To run the app download everything from binary directory and start *windows_lock.exe*.
For auto start create a shortcut to *windows_lock.exe* and place it in an autorun folder on your system.
To run in a debug mode open CMD, change current directory to where *windows_lock.exe* is located and execute: **windows_lock.exe -debug**
