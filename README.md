learning-goroutines
===================

These files are step-by-step files to teach new Go users how goroutines work.

1. send to channel that isn't receiving  
2. go send to channel through function - still not receiving  
3. go call function from goroutine, works properly  
4. go add recursion: main pulls first thing in channel and quits  
5. go make separate channels in recursion: sending before channels hooked up  
6. go connect channels with relevant goroutine  
7. go range: channels never closed  8 minutes ago  
8. go defer outside of gofunc: closes channel immediately after first recursion level and quits  
9. go defer inside gofunc: works properly with range  
