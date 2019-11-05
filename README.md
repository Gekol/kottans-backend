# kottans-backend

## Git Basics

I have watched the Udacity Git Course and it was awesome!!! I've already had experience with Git but it was great to have an opportunity to learn some new commands.

## Unix Shell

I have passed the Linux Survival Course and I had many things to learn there. The knowledge I got will be necessary for me in the future so I'm glad to have got it.

## Git Collaboration

I have watched the Udacity Git Collaboration Course. It was more difficult than I thought. However, now I know how to use Git and GitHub for teamwork. That knowledge is essential for any IT specialist so it's excellent I have it!!!

## Go basics 1
https://github.com/Gekol/kottans-backend/blob/master/roman-digits/main.go

## Memory Management
What's going to happen if program reaches maximum limit of stack?
Answer: there will be stack overflow error and the program will receive Segmentation Fault. But if the stack is dynamic we will be able to access to an unmapped memory region.

What's going to happen if program requests a big (more then 128KB) memory allocation on heap?
Answer: the heap will be enlarged via the brk() system call (implementation) to make room for the requested block.

What's the difference between Text and Data memory segments?
Answer: Data Segment stores the contents for static variables initialized in source code. The text segment is read-only and stores all of your code in addition to tidbits like string literals.
```
55d28a57c000-55d28a57f000 r--p 00000000 08:01 925055                     /usr/lib/gnome-settings-daemon/gsd-sound
55d28a57f000-55d28a581000 r-xp 00003000 08:01 925055                     /usr/lib/gnome-settings-daemon/gsd-sound
55d28a581000-55d28a582000 r--p 00005000 08:01 925055                     /usr/lib/gnome-settings-daemon/gsd-sound
55d28a583000-55d28a584000 r--p 00006000 08:01 925055                     /usr/lib/gnome-settings-daemon/gsd-sound
55d28a584000-55d28a585000 rw-p 00007000 08:01 925055                     /usr/lib/gnome-settings-daemon/gsd-sound
55d28b7ad000-55d28b811000 rw-p 00000000 00:00 0                          [heap]
7ffb70000000-7ffb70021000 rw-p 00000000 00:00 0 
7ffb70021000-7ffb74000000 ---p 00000000 00:00 0 
7ffb74000000-7ffb74021000 rw-p 00000000 00:00 0 
7ffb74021000-7ffb78000000 ---p 00000000 00:00 0 
7ffb78000000-7ffb78021000 rw-p 00000000 00:00 0 
7ffb78021000-7ffb7c000000 ---p 00000000 00:00 0 
7ffb7e05e000-7ffb7e070000 r--p 00000000 08:01 134863                     /usr/lib/x86_64-linux-gnu/gvfs/libgvfscommon.so
7ffb7e070000-7ffb7e089000 r-xp 00012000 08:01 134863                     /usr/lib/x86_64-linux-gnu/gvfs/libgvfscommon.so
7ffb7e089000-7ffb7e096000 r--p 0002b000 08:01 134863                     /usr/lib/x86_64-linux-gnu/gvfs/libgvfscommon.so
7ffb7e096000-7ffb7e09c000 r--p 00037000 08:01 134863                     /usr/lib/x86_64-linux-gnu/gvfs/libgvfscommon.so
7ffb7e09c000-7ffb7e09d000 rw-p 0003d000 08:01 134863                     /usr/lib/x86_64-linux-gnu/gvfs/libgvfscommon.so
7ffb7e09d000-7ffb7e0a7000 r--p 00000000 08:01 134821                     /usr/lib/x86_64-linux-gnu/gio/modules/libgvfsdbus.so
7ffb7e0a7000-7ffb7e0c4000 r-xp 0000a000 08:01 134821                     /usr/lib/x86_64-linux-gnu/gio/modules/libgvfsdbus.so
7ffb7e0c4000-7ffb7e0cf000 r--p 00027000 08:01 134821                     /usr/lib/x86_64-linux-gnu/gio/modules/libgvfsdbus.so
7ffb7e0cf000-7ffb7e0d1000 r--p 00031000 08:01 134821                     /usr/lib/x86_64-linux-gnu/gio/modules/libgvfsdbus.so
7ffb7e0d1000-7ffb7e0d2000 rw-p 00033000 08:01 134821                     /usr/lib/x86_64-linux-gnu/gio/modules/libgvfsdbus.so
7ffb7e0d2000-7ffb7e0d3000 ---p 00000000 00:00 0 
7ffb7e0d3000-7ffb7e8d3000 rw-p 00000000 00:00 0 
7ffb7e8d3000-7ffb7e903000 r--p 00000000 08:01 524328                     /usr/share/glib-2.0/schemas/gschemas.compiled
7ffb7e903000-7ffb7e904000 ---p 00000000 00:00 0 
7ffb7e904000-7ffb7f104000 rw-p 00000000 00:00 0 
7ffb7f104000-7ffb7f105000 ---p 00000000 00:00 0 
7ffb7f105000-7ffb7f905000 rw-p 00000000 00:00 0 
7ffb7f905000-7ffb806e0000 r--p 00000000 08:01 925836                     /usr/lib/locale/locale-archive
7ffb806e0000-7ffb806e7000 rw-p 00000000 00:00 0 
7ffb806e7000-7ffb806eb000 r--p 00000000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb806eb000-7ffb806fe000 r-xp 00004000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb806fe000-7ffb80707000 r--p 00017000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb80707000-7ffb80708000 ---p 00020000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb80708000-7ffb80709000 r--p 00020000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb80709000-7ffb8070a000 rw-p 00021000 08:01 926922                     /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.26.1
7ffb8070a000-7ffb8070e000 r--p 00000000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb8070e000-7ffb8071d000 r-xp 00004000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb8071d000-7ffb80720000 r--p 00013000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb80720000-7ffb80721000 ---p 00016000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb80721000-7ffb80722000 r--p 00016000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb80722000-7ffb80723000 rw-p 00017000 08:01 926551                     /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1
7ffb80723000-7ffb80724000 rw-p 00000000 00:00 0 
7ffb80724000-7ffb80738000 r--p 00000000 08:01 927758                     /usr/lib/x86_64-linux-gnu/libvorbisenc.so.2.0.11
7ffb80738000-7ffb8073b000 r-xp 00014000 08:01 927758                     /usr/lib/x86_64-linux-gnu/libvorbisenc.so.2.0.11
7ffb8073b000-7ffb807b2000 r--p 00017000 08:01 927758                     /usr/lib/x86_64-linux-gnu/libvorbisenc.so.2.0.11
7ffb807b2000-7ffb807ce000 r--p 0008d000 08:01 927758                     /usr/lib/x86_64-linux-gnu/libvorbisenc.so.2.0.11
7ffb807ce000-7ffb807cf000 rw-p 000a9000 08:01 927758                     /usr/lib/x86_64-linux-gnu/libvorbisenc.so.2.0.11
7ffb807cf000-7ffb807d2000 r--p 00000000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807d2000-7ffb807e8000 r-xp 00003000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807e8000-7ffb807f9000 r--p 00019000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807f9000-7ffb807fa000 ---p 0002a000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807fa000-7ffb807fb000 r--p 0002a000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807fb000-7ffb807fc000 rw-p 0002b000 08:01 927756                     /usr/lib/x86_64-linux-gnu/libvorbis.so.0.4.8
7ffb807fc000-7ffb80803000 r-xp 00000000 08:01 927348                     /usr/lib/x86_64-linux-gnu/libogg.so.0.8.2
7ffb80803000-7ffb80a03000 ---p 00007000 08:01 927348                     /usr/lib/x86_64-linux-gnu/libogg.so.0.8.2
7ffb80a03000-7ffb80a04000 r--p 00007000 08:01 927348                     /usr/lib/x86_64-linux-gnu/libogg.so.0.8.2
7ffb80a04000-7ffb80a05000 rw-p 00008000 08:01 927348                     /usr/lib/x86_64-linux-gnu/libogg.so.0.8.2
7ffb80a05000-7ffb80a07000 rw-p 00000000 00:00 0 
7ffb80a07000-7ffb80a7c000 r-xp 00000000 08:01 926339                     /usr/lib/x86_64-linux-gnu/libFLAC.so.8.3.0
7ffb80a7c000-7ffb80c7c000 ---p 00075000 08:01 926339                     /usr/lib/x86_64-linux-gnu/libFLAC.so.8.3.0
7ffb80c7c000-7ffb80c7d000 r--p 00075000 08:01 926339                     /usr/lib/x86_64-linux-gnu/libFLAC.so.8.3.0
7ffb80c7d000-7ffb80c7e000 rw-p 00076000 08:01 926339                     /usr/lib/x86_64-linux-gnu/libFLAC.so.8.3.0
7ffb80c7e000-7ffb80c83000 r--p 00000000 08:01 927307                     /usr/lib/x86_64-linux-gnu/libnsl-2.29.so
7ffb80c83000-7ffb80c91000 r-xp 00005000 08:01 927307                     /usr/lib/x86_64-linux-gnu/libnsl-2.29.so
7ffb80c91000-7ffb80c95000 r--p 00013000 08:01 927307                     /usr/lib/x86_64-linux-gnu/libnsl-2.29.so
7ffb80c95000-7ffb80c96000 r--p 00016000 08:01 927307                     /usr/lib/x86_64-linux-gnu/libnsl-2.29.so
7ffb80c96000-7ffb80c97000 rw-p 00017000 08:01 927307                     /usr/lib/x86_64-linux-gnu/libnsl-2.29.so
7ffb80c97000-7ffb80c99000 rw-p 00000000 00:00 0 
7ffb80c99000-7ffb80ca5000 r--p 00000000 08:01 926851                     /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.4
7ffb80ca5000-7ffb80d72000 r-xp 0000c000 08:01 926851                     /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.4
7ffb80d72000-7ffb80daf000 r--p 000d9000 08:01 926851                     /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.4
7ffb80daf000-7ffb80db1000 r--p 00115000 08:01 926851                     /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.4
7ffb80db1000-7ffb80db6000 rw-p 00117000 08:01 926851                     /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.4
7ffb80db6000-7ffb80db8000 r--p 00000000 08:01 927185                     /usr/lib/x86_64-linux-gnu/liblz4.so.1.8.3
7ffb80db8000-7ffb80de0000 r-xp 00002000 08:01 927185                     /usr/lib/x86_64-linux-gnu/liblz4.so.1.8.3
7ffb80de0000-7ffb80de3000 r--p 0002a000 08:01 927185                     /usr/lib/x86_64-linux-gnu/liblz4.so.1.8.3
7ffb80de3000-7ffb80de4000 r--p 0002c000 08:01 927185                     /usr/lib/x86_64-linux-gnu/liblz4.so.1.8.3
7ffb80de4000-7ffb80de5000 rw-p 0002d000 08:01 927185                     /usr/lib/x86_64-linux-gnu/liblz4.so.1.8.3
7ffb80de5000-7ffb80de8000 r--p 00000000 08:01 927187                     /usr/lib/x86_64-linux-gnu/liblzma.so.5.2.4
7ffb80de8000-7ffb80dff000 r-xp 00003000 08:01 927187                     /usr/lib/x86_64-linux-gnu/liblzma.so.5.2.4
7ffb80dff000-7ffb80e0a000 r--p 0001a000 08:01 927187                     /usr/lib/x86_64-linux-gnu/liblzma.so.5.2.4
7ffb80e0a000-7ffb80e0b000 r--p 00024000 08:01 927187                     /usr/lib/x86_64-linux-gnu/liblzma.so.5.2.4
7ffb80e0b000-7ffb80e0c000 rw-p 00025000 08:01 927187                     /usr/lib/x86_64-linux-gnu/liblzma.so.5.2.4
7ffb80e0c000-7ffb80e11000 r-xp 00000000 08:01 926387                     /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0
7ffb80e11000-7ffb81010000 ---p 00005000 08:01 926387                     /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0
7ffb81010000-7ffb81011000 r--p 00004000 08:01 926387                     /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0
7ffb81011000-7ffb81012000 rw-p 00005000 08:01 926387                     /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0
7ffb81012000-7ffb81014000 r-xp 00000000 08:01 926376                     /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0
7ffb81014000-7ffb81214000 ---p 00002000 08:01 926376                     /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0
7ffb81214000-7ffb81215000 r--p 00002000 08:01 926376                     /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0
7ffb81215000-7ffb81216000 rw-p 00003000 08:01 926376                     /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0
7ffb81216000-7ffb81218000 rw-p 00000000 00:00 0 
7ffb81218000-7ffb8121a000 r--p 00000000 08:01 927740                     /usr/lib/x86_64-linux-gnu/libuuid.so.1.3.0
7ffb8121a000-7ffb8121e000 r-xp 00002000 08:01 927740                     /usr/lib/x86_64-linux-gnu/libuuid.so.1.3.0
7ffb8121e000-7ffb8121f000 r--p 00006000 08:01 927740                     /usr/lib/x86_64-linux-gnu/libuuid.so.1.3.0
7ffb8121f000-7ffb81220000 r--p 00006000 08:01 927740                     /usr/lib/x86_64-linux-gnu/libuuid.so.1.3.0
7ffb81220000-7ffb81221000 rw-p 00007000 08:01 927740                     /usr/lib/x86_64-linux-gnu/libuuid.so.1.3.0
7ffb81221000-7ffb81226000 r-xp 00000000 08:01 926471                     /usr/lib/x86_64-linux-gnu/libasyncns.so.0.3.1
7ffb81226000-7ffb81425000 ---p 00005000 08:01 926471                     /usr/lib/x86_64-linux-gnu/libasyncns.so.0.3.1
7ffb81425000-7ffb81426000 r--p 00004000 08:01 926471                     /usr/lib/x86_64-linux-gnu/libasyncns.so.0.3.1
7ffb81426000-7ffb81427000 rw-p 00005000 08:01 926471                     /usr/lib/x86_64-linux-gnu/libasyncns.so.0.3.1
7ffb81427000-7ffb8142e000 r--p 00000000 08:01 927595                     /usr/lib/x86_64-linux-gnu/libsndfile.so.1.0.28
7ffb8142e000-7ffb81480000 r-xp 00007000 08:01 927595                     /usr/lib/x86_64-linux-gnu/libsndfile.so.1.0.28
7ffb81480000-7ffb8149d000 r--p 00059000 08:01 927595                     /usr/lib/x86_64-linux-gnu/libsndfile.so.1.0.28
7ffb8149d000-7ffb814a0000 r--p 00075000 08:01 927595                     /usr/lib/x86_64-linux-gnu/libsndfile.so.1.0.28
7ffb814a0000-7ffb814a1000 rw-p 00078000 08:01 927595                     /usr/lib/x86_64-linux-gnu/libsndfile.so.1.0.28
7ffb814a1000-7ffb814a3000 rw-p 00000000 00:00 0 
7ffb814a3000-7ffb814a6000 r--p 00000000 08:01 927827                     /usr/lib/x86_64-linux-gnu/libwrap.so.0.7.6
7ffb814a6000-7ffb814ab000 r-xp 00003000 08:01 927827                     /usr/lib/x86_64-linux-gnu/libwrap.so.0.7.6
7ffb814ab000-7ffb814ad000 r--p 00008000 08:01 927827                     /usr/lib/x86_64-linux-gnu/libwrap.so.0.7.6
7ffb814ad000-7ffb814ae000 r--p 00009000 08:01 927827                     /usr/lib/x86_64-linux-gnu/libwrap.so.0.7.6
7ffb814ae000-7ffb814af000 rw-p 0000a000 08:01 927827                     /usr/lib/x86_64-linux-gnu/libwrap.so.0.7.6
7ffb814af000-7ffb814be000 r--p 00000000 08:01 917629                     /usr/lib/x86_64-linux-gnu/libsystemd.so.0.24.0
7ffb814be000-7ffb81527000 r-xp 0000f000 08:01 917629                     /usr/lib/x86_64-linux-gnu/libsystemd.so.0.24.0
7ffb81527000-7ffb81548000 r--p 00078000 08:01 917629                     /usr/lib/x86_64-linux-gnu/libsystemd.so.0.24.0
7ffb81548000-7ffb8154b000 r--p 00098000 08:01 917629                     /usr/lib/x86_64-linux-gnu/libsystemd.so.0.24.0
7ffb8154b000-7ffb8154c000 rw-p 0009b000 08:01 917629                     /usr/lib/x86_64-linux-gnu/libsystemd.so.0.24.0
7ffb8154c000-7ffb8154f000 rw-p 00000000 00:00 0 
7ffb8154f000-7ffb8155a000 r--p 00000000 08:01 927867                     /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0
7ffb8155a000-7ffb8156d000 r-xp 0000b000 08:01 927867                     /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0
7ffb8156d000-7ffb81576000 r--p 0001e000 08:01 927867                     /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0
7ffb81576000-7ffb81577000 r--p 00026000 08:01 927867                     /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0
7ffb81577000-7ffb81578000 rw-p 00027000 08:01 927867                     /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0
7ffb81578000-7ffb81582000 r--p 00000000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb81582000-7ffb815b6000 r-xp 0000a000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb815b6000-7ffb815c6000 r--p 0003e000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb815c6000-7ffb815c7000 ---p 0004e000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb815c7000-7ffb815cc000 r--p 0004e000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb815cc000-7ffb815cd000 rw-p 00053000 08:01 926530                     /usr/lib/x86_64-linux-gnu/libblkid.so.1.1.0
7ffb815cd000-7ffb815d0000 r--p 00000000 08:01 927530                     /usr/lib/x86_64-linux-gnu/librt-2.29.so
7ffb815d0000-7ffb815d4000 r-xp 00003000 08:01 927530                     /usr/lib/x86_64-linux-gnu/librt-2.29.so
7ffb815d4000-7ffb815d6000 r--p 00007000 08:01 927530                     /usr/lib/x86_64-linux-gnu/librt-2.29.so
7ffb815d6000-7ffb815d7000 r--p 00008000 08:01 927530                     /usr/lib/x86_64-linux-gnu/librt-2.29.so
7ffb815d7000-7ffb815d8000 rw-p 00009000 08:01 927530                     /usr/lib/x86_64-linux-gnu/librt-2.29.so
7ffb815d8000-7ffb815e7000 r--p 00000000 08:01 927191                     /usr/lib/x86_64-linux-gnu/libm-2.29.so
7ffb815e7000-7ffb8168d000 r-xp 0000f000 08:01 927191                     /usr/lib/x86_64-linux-gnu/libm-2.29.so
7ffb8168d000-7ffb81724000 r--p 000b5000 08:01 927191                     /usr/lib/x86_64-linux-gnu/libm-2.29.so
7ffb81724000-7ffb81725000 r--p 0014b000 08:01 927191                     /usr/lib/x86_64-linux-gnu/libm-2.29.so
7ffb81725000-7ffb81726000 rw-p 0014c000 08:01 927191                     /usr/lib/x86_64-linux-gnu/libm-2.29.so
7ffb81726000-7ffb81731000 r--p 00000000 08:01 917739                     /usr/lib/x86_64-linux-gnu/libdbus-1.so.3.19.9
7ffb81731000-7ffb8175e000 r-xp 0000b000 08:01 917739                     /usr/lib/x86_64-linux-gnu/libdbus-1.so.3.19.9
7ffb8175e000-7ffb81773000 r--p 00038000 08:01 917739                     /usr/lib/x86_64-linux-gnu/libdbus-1.so.3.19.9
7ffb81773000-7ffb81774000 r--p 0004c000 08:01 917739                     /usr/lib/x86_64-linux-gnu/libdbus-1.so.3.19.9
7ffb81774000-7ffb81775000 rw-p 0004d000 08:01 917739                     /usr/lib/x86_64-linux-gnu/libdbus-1.so.3.19.9
7ffb81775000-7ffb81777000 rw-p 00000000 00:00 0 
7ffb81777000-7ffb81788000 r--p 00000000 08:01 137869                     /usr/lib/x86_64-linux-gnu/pulseaudio/libpulsecommon-12.2.so
7ffb81788000-7ffb817ce000 r-xp 00011000 08:01 137869                     /usr/lib/x86_64-linux-gnu/pulseaudio/libpulsecommon-12.2.so
7ffb817ce000-7ffb817f4000 r--p 00057000 08:01 137869                     /usr/lib/x86_64-linux-gnu/pulseaudio/libpulsecommon-12.2.so
7ffb817f4000-7ffb817f6000 r--p 0007c000 08:01 137869                     /usr/lib/x86_64-linux-gnu/pulseaudio/libpulsecommon-12.2.so
7ffb817f6000-7ffb817f7000 rw-p 0007e000 08:01 137869                     /usr/lib/x86_64-linux-gnu/pulseaudio/libpulsecommon-12.2.so
7ffb817f7000-7ffb817f9000 r--p 00000000 08:01 927402                     /usr/lib/x86_64-linux-gnu/libpcre.so.3.13.3
7ffb817f9000-7ffb8184b000 r-xp 00002000 08:01 927402                     /usr/lib/x86_64-linux-gnu/libpcre.so.3.13.3
7ffb8184b000-7ffb81869000 r--p 00054000 08:01 927402                     /usr/lib/x86_64-linux-gnu/libpcre.so.3.13.3
7ffb81869000-7ffb8186a000 r--p 00071000 08:01 927402                     /usr/lib/x86_64-linux-gnu/libpcre.so.3.13.3
7ffb8186a000-7ffb8186b000 rw-p 00072000 08:01 927402                     /usr/lib/x86_64-linux-gnu/libpcre.so.3.13.3
7ffb8186b000-7ffb8186d000 r--p 00000000 08:01 926796                     /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7ffb8186d000-7ffb81872000 r-xp 00002000 08:01 926796                     /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7ffb81872000-7ffb81873000 r--p 00007000 08:01 926796                     /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7ffb81873000-7ffb81874000 r--p 00007000 08:01 926796                     /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7ffb81874000-7ffb81875000 rw-p 00008000 08:01 926796                     /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7ffb81875000-7ffb8187c000 r--p 00000000 08:01 927462                     /usr/lib/x86_64-linux-gnu/libpthread-2.29.so
7ffb8187c000-7ffb8188b000 r-xp 00007000 08:01 927462                     /usr/lib/x86_64-linux-gnu/libpthread-2.29.so
7ffb8188b000-7ffb81890000 r--p 00016000 08:01 927462                     /usr/lib/x86_64-linux-gnu/libpthread-2.29.so
7ffb81890000-7ffb81891000 r--p 0001a000 08:01 927462                     /usr/lib/x86_64-linux-gnu/libpthread-2.29.so
7ffb81891000-7ffb81892000 rw-p 0001b000 08:01 927462                     /usr/lib/x86_64-linux-gnu/libpthread-2.29.so
7ffb81892000-7ffb81896000 rw-p 00000000 00:00 0 
7ffb81896000-7ffb8189a000 r--p 00000000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb8189a000-7ffb818a9000 r-xp 00004000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb818a9000-7ffb818ac000 r--p 00013000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb818ac000-7ffb818ad000 ---p 00016000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb818ad000-7ffb818ae000 r--p 00016000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb818ae000-7ffb818af000 rw-p 00017000 08:01 927508                     /usr/lib/x86_64-linux-gnu/libresolv-2.29.so
7ffb818af000-7ffb818b3000 rw-p 00000000 00:00 0 
7ffb818b3000-7ffb818b9000 r--p 00000000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818b9000-7ffb818d1000 r-xp 00006000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818d1000-7ffb818d8000 r--p 0001e000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818d8000-7ffb818d9000 ---p 00025000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818d9000-7ffb818da000 r--p 00025000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818da000-7ffb818db000 rw-p 00026000 08:01 927560                     /usr/lib/x86_64-linux-gnu/libselinux.so.1
7ffb818db000-7ffb818dd000 rw-p 00000000 00:00 0 
7ffb818dd000-7ffb818e8000 r--p 00000000 08:01 927217                     /usr/lib/x86_64-linux-gnu/libmount.so.1.1.0
7ffb818e8000-7ffb81923000 r-xp 0000b000 08:01 927217                     /usr/lib/x86_64-linux-gnu/libmount.so.1.1.0
7ffb81923000-7ffb81935000 r--p 00046000 08:01 927217                     /usr/lib/x86_64-linux-gnu/libmount.so.1.1.0
7ffb81935000-7ffb81938000 r--p 00057000 08:01 927217                     /usr/lib/x86_64-linux-gnu/libmount.so.1.1.0
7ffb81938000-7ffb81939000 rw-p 0005a000 08:01 927217                     /usr/lib/x86_64-linux-gnu/libmount.so.1.1.0
7ffb81939000-7ffb8193a000 r--p 00000000 08:01 926695                     /usr/lib/x86_64-linux-gnu/libdl-2.29.so
7ffb8193a000-7ffb8193c000 r-xp 00001000 08:01 926695                     /usr/lib/x86_64-linux-gnu/libdl-2.29.so
7ffb8193c000-7ffb8193d000 r--p 00003000 08:01 926695                     /usr/lib/x86_64-linux-gnu/libdl-2.29.so
7ffb8193d000-7ffb8193e000 r--p 00003000 08:01 926695                     /usr/lib/x86_64-linux-gnu/libdl-2.29.so
7ffb8193e000-7ffb8193f000 rw-p 00004000 08:01 926695                     /usr/lib/x86_64-linux-gnu/libdl-2.29.so
7ffb8193f000-7ffb81941000 r--p 00000000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb81941000-7ffb81952000 r-xp 00002000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb81952000-7ffb81958000 r--p 00013000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb81958000-7ffb81959000 ---p 00019000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb81959000-7ffb8195a000 r--p 00019000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb8195a000-7ffb8195b000 rw-p 0001a000 08:01 927897                     /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
7ffb8195b000-7ffb8195c000 r--p 00000000 08:01 917967                     /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.6000.0
7ffb8195c000-7ffb8195e000 r-xp 00001000 08:01 917967                     /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.6000.0
7ffb8195e000-7ffb8195f000 r--p 00003000 08:01 917967                     /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.6000.0
7ffb8195f000-7ffb81960000 r--p 00003000 08:01 917967                     /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.6000.0
7ffb81960000-7ffb81961000 rw-p 00004000 08:01 917967                     /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.6000.0
7ffb81961000-7ffb81966000 r--p 00000000 08:01 927780                     /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7ffb81966000-7ffb8196c000 r-xp 00005000 08:01 927780                     /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7ffb8196c000-7ffb8196f000 r--p 0000b000 08:01 927780                     /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7ffb8196f000-7ffb81971000 r--p 0000d000 08:01 927780                     /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7ffb81971000-7ffb81972000 rw-p 0000f000 08:01 927780                     /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7ffb81972000-7ffb81974000 rw-p 00000000 00:00 0 
7ffb81974000-7ffb81999000 r--p 00000000 08:01 926555                     /usr/lib/x86_64-linux-gnu/libc-2.29.so
7ffb81999000-7ffb81b0c000 r-xp 00025000 08:01 926555                     /usr/lib/x86_64-linux-gnu/libc-2.29.so
7ffb81b0c000-7ffb81b55000 r--p 00198000 08:01 926555                     /usr/lib/x86_64-linux-gnu/libc-2.29.so
7ffb81b55000-7ffb81b58000 r--p 001e0000 08:01 926555                     /usr/lib/x86_64-linux-gnu/libc-2.29.so
7ffb81b58000-7ffb81b5b000 rw-p 001e3000 08:01 926555                     /usr/lib/x86_64-linux-gnu/libc-2.29.so
7ffb81b5b000-7ffb81b5f000 rw-p 00000000 00:00 0 
7ffb81b5f000-7ffb81b6a000 r--p 00000000 08:01 927471                     /usr/lib/x86_64-linux-gnu/libpulse.so.0.20.3
7ffb81b6a000-7ffb81b98000 r-xp 0000b000 08:01 927471                     /usr/lib/x86_64-linux-gnu/libpulse.so.0.20.3
7ffb81b98000-7ffb81bad000 r--p 00039000 08:01 927471                     /usr/lib/x86_64-linux-gnu/libpulse.so.0.20.3
7ffb81bad000-7ffb81baf000 r--p 0004d000 08:01 927471                     /usr/lib/x86_64-linux-gnu/libpulse.so.0.20.3
7ffb81baf000-7ffb81bb0000 rw-p 0004f000 08:01 927471                     /usr/lib/x86_64-linux-gnu/libpulse.so.0.20.3
7ffb81bb0000-7ffb81bcb000 r--p 00000000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81bcb000-7ffb81c4b000 r-xp 0001b000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81c4b000-7ffb81ccd000 r--p 0009b000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81ccd000-7ffb81cce000 ---p 0011d000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81cce000-7ffb81ccf000 r--p 0011d000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81ccf000-7ffb81cd0000 rw-p 0011e000 08:01 917943                     /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.6000.0
7ffb81cd0000-7ffb81cd1000 rw-p 00000000 00:00 0 
7ffb81cd1000-7ffb81cdf000 r--p 00000000 08:01 917968                     /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.6000.0
7ffb81cdf000-7ffb81d11000 r-xp 0000e000 08:01 917968                     /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.6000.0
7ffb81d11000-7ffb81d29000 r--p 00040000 08:01 917968                     /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.6000.0
7ffb81d29000-7ffb81d2c000 r--p 00057000 08:01 917968                     /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.6000.0
7ffb81d2c000-7ffb81d2d000 rw-p 0005a000 08:01 917968                     /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.6000.0
7ffb81d2d000-7ffb81d64000 r--p 00000000 08:01 917920                     /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.6000.0
7ffb81d64000-7ffb81e65000 r-xp 00037000 08:01 917920                     /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.6000.0
7ffb81e65000-7ffb81eeb000 r--p 00138000 08:01 917920                     /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.6000.0
7ffb81eeb000-7ffb81ef3000 r--p 001bd000 08:01 917920                     /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.6000.0
7ffb81ef3000-7ffb81ef4000 rw-p 001c5000 08:01 917920                     /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.6000.0
7ffb81ef4000-7ffb81ef6000 rw-p 00000000 00:00 0 
7ffb81ef6000-7ffb81ef7000 r--p 00000000 08:01 917571                     /home/gekol/.config/dconf/user (deleted)
7ffb81ef7000-7ffb81ef8000 r--s 00000000 00:31 22                         /run/user/1000/dconf/user (deleted)
7ffb81ef8000-7ffb81efc000 r--p 00000000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81efc000-7ffb81f03000 r-xp 00004000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81f03000-7ffb81f06000 r--p 0000b000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81f06000-7ffb81f07000 ---p 0000e000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81f07000-7ffb81f08000 r--p 0000e000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81f08000-7ffb81f09000 rw-p 0000f000 08:01 136252                     /usr/lib/x86_64-linux-gnu/gio/modules/libdconfsettings.so
7ffb81f09000-7ffb81f11000 r--p 00000000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f11000-7ffb81f1a000 r-xp 00008000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f1a000-7ffb81f1f000 r--p 00011000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f1f000-7ffb81f20000 ---p 00016000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f20000-7ffb81f22000 r--p 00016000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f22000-7ffb81f23000 rw-p 00018000 08:01 925061                     /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7ffb81f23000-7ffb81f25000 rw-p 00000000 00:00 0 
7ffb81f25000-7ffb81f26000 r--p 00000000 08:01 926328                     /usr/lib/x86_64-linux-gnu/ld-2.29.so
7ffb81f26000-7ffb81f47000 r-xp 00001000 08:01 926328                     /usr/lib/x86_64-linux-gnu/ld-2.29.so
7ffb81f47000-7ffb81f4f000 r--p 00022000 08:01 926328                     /usr/lib/x86_64-linux-gnu/ld-2.29.so
7ffb81f4f000-7ffb81f50000 r--p 00029000 08:01 926328                     /usr/lib/x86_64-linux-gnu/ld-2.29.so
7ffb81f50000-7ffb81f51000 rw-p 0002a000 08:01 926328                     /usr/lib/x86_64-linux-gnu/ld-2.29.so
7ffb81f51000-7ffb81f52000 rw-p 00000000 00:00 0 
7ffd1ade0000-7ffd1ae01000 rw-p 00000000 00:00 0                          [stack]
7ffd1afb1000-7ffd1afb4000 r--p 00000000 00:00 0                          [vvar]
7ffd1afb4000-7ffd1afb5000 r-xp 00000000 00:00 0                          [vdso]
ffffffffff600000-ffffffffff601000 r-xp 00000000 00:00 0                  [vsyscall]
```
After reading the articles I felt that I have become much more sure in my knowledge about Memory Management. I am especially thankful to authors of the article "Anatomy of a Program in Memory."

## TCP. UDP. Network

https://github.com/Gekol/kottans-backend/blob/master/task_networks/Internet_101.jpg \
https://github.com/Gekol/kottans-backend/blob/master/task_networks/Networking_for_Web_Developers.png
The courses were really interesting and I've learned a lot of new things. I've never thought that there are wires under the Ocean. I used to think that it was the "power" of wireless Internet. My new knowledge is priceless for me.

## HTTP & HTTPS
1. Name at least three possible negative consequences of not using https.
    1. Other people will be able to eavesdrop my connection.
    2. Other people will be able to intercept and update messages I send to server from my browser.
    3. Server and my browser won't be able to recognize each other. Thus, I may get on absolutely wrong site.
2. Public key is used in assymetric key algorithm. Public key is used to encrypt messages while the private key is used to decrypt them. My browser uses the server's public key and encrypt my message for server with it and server uses its private key to decrypt my message and vice versa.
3.  1. POST
    2. GET
    3. PUT
    4. POST
    5. PUT
    6. POST
    
My curl requests:
curl -i https://api.github.com/users/Gekol
curl  https://api.github.com/gists/starred
curl --user "Gekol:BLABLABLA" https://api.github.com/gists/starred
curl --user "Gekol:************" https://api.github.com/gists/starred
curl --user "Gekol" https://api.github.com/gists/starred

## Patterns
https://github.com/Gekol/kottans-backend/blob/master/task_patterns/Lessons1.png \
https://github.com/Gekol/kottans-backend/blob/master/task_patterns/Lessons2.png \
https://github.com/Gekol/kottans-backend/blob/master/task_patterns/Lessons3.png \
https://github.com/Gekol/kottans-backend/blob/master/task_patterns/Lessons4.png \
https://github.com/Gekol/kottans-backend/blob/master/task_patterns/Lessons5.png

It was rather interesting but the course turned to be really big. It took me much more time than I expected
