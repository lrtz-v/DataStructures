# 数据结构及算法

- [SOURCE](https://github.com/jwasham/google-interview-university.git)

## 基本数据结构

### 数组（Arrays）

- 实现一个可自动调整大小的动态数组。
- [ ] 介绍：
  - [数组（视频）](https://www.coursera.org/learn/data-structures/lecture/OsBSF/arrays)
  - [数组的基础知识（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Basic-arrays/149042/177104-4.html)
  - [多维数组（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Multidimensional-arrays/149042/177105-4.html)
  - [动态数组（视频）](https://www.coursera.org/learn/data-structures/lecture/EwbnV/dynamic-arrays)
  - [不规则数组（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Jagged-arrays/149042/177106-4.html)
  - [调整数组的大小（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Resizable-arrays/149042/177108-4.html)
- [ ] 实现一个动态数组（可自动调整大小的可变数组）：
  - [ ] 练习使用数组和指针去编码，并且指针是通过计算去跳转而不是使用索引
  - [ ] 通过分配内存来新建一个原生数据型数组
    - 可以使用 int 类型的数组，但不能使用其语法特性
    - 从大小为 16 或更大的数（使用 2 的倍数 —— 16、32、64、128）开始编写
  - [ ] size() —— 数组元素的个数
  - [ ] capacity() —— 可容纳元素的个数
  - [ ] is_empty()
  - [ ] at(index) —— 返回对应索引的元素，且若索引越界则愤然报错
  - [ ] push(item)
  - [ ] insert(index, item) —— 在指定索引中插入元素，并把后面的元素依次后移
  - [ ] prepend(item) —— 可以使用上面的 insert 函数，传参 index 为 0
  - [ ] pop() —— 删除在数组末端的元素，并返回其值
  - [ ] delete(index) —— 删除指定索引的元素，并把后面的元素依次前移
  - [ ] remove(item) —— 删除指定值的元素，并返回其索引（即使有多个元素）
  - [ ] find(item) —— 寻找指定值的元素并返回其中第一个出现的元素其索引，若未找到则返回 -1
  - [ ] resize(new_capacity) // 私有函数
    - 若数组的大小到达其容积，则变大一倍
    - 获取元素后，若数组大小为其容积的 1/4，则缩小一半
- [ ] 时间复杂度
  - 在数组末端增加/删除、定位、更新元素，只允许占 O(1) 的时间复杂度（平摊（amortized）去分配内存以获取更多空间）
  - 在数组任何地方插入/移除元素，只允许 O(n) 的时间复杂度
- [ ] 空间复杂度

  - 因为在内存中分配的空间邻近，所以有助于提高性能
  - 空间需求 = （大于或等于 n 的数组容积）\* 元素的大小。即便空间需求为 2n，其空间复杂度仍然是 O(n)

### Linked Lists

- [ ] Description:
  - [ ] [Singly Linked Lists (video)](https://www.coursera.org/lecture/data-structures/singly-linked-lists-kHhgK)
  - [ ] [CS 61B - Linked Lists 1 (video)](https://archive.org/details/ucberkeley_webcast_htzJdKoEmO0)
  - [ ] [CS 61B - Linked Lists 2 (video)](https://archive.org/details/ucberkeley_webcast_-c4I3gFYe3w)
- [ ] [C Code (video)](https://www.youtube.com/watch?v=QN6FPiD0Gzo) - not the whole video, just portions about Node struct and memory allocation
- [ ] Linked List vs Arrays:
  - [Core Linked Lists Vs Arrays (video)](https://www.coursera.org/lecture/data-structures-optimizing-performance/core-linked-lists-vs-arrays-rjBs9)
  - [In The Real World Linked Lists Vs Arrays (video)](https://www.coursera.org/lecture/data-structures-optimizing-performance/in-the-real-world-lists-vs-arrays-QUaUd)
- [ ] [why you should avoid linked lists (video)](https://www.youtube.com/watch?v=YQs6IC-vgmo)
- [ ] Gotcha: you need pointer to pointer knowledge:
      (for when you pass a pointer to a function that may change the address where that pointer points)
      This page is just to get a grasp on ptr to ptr. I don't recommend this list traversal style. Readability and maintainability suffer due to cleverness.
  - [Pointers to Pointers](https://www.eskimo.com/~scs/cclass/int/sx8.html)
- [ ] Implement (I did with tail pointer & without):
  - [ ] size() - returns number of data elements in list
  - [ ] empty() - bool returns true if empty
  - [ ] value_at(index) - returns the value of the nth item (starting at 0 for first)
  - [ ] push_front(value) - adds an item to the front of the list
  - [ ] pop_front() - remove front item and return its value
  - [ ] push_back(value) - adds an item at the end
  - [ ] pop_back() - removes end item and returns its value
  - [ ] front() - get value of front item
  - [ ] back() - get value of end item
  - [ ] insert(index, value) - insert value at index, so current item at that index is pointed to by new item at index
  - [ ] erase(index) - removes node at given index
  - [ ] value_n_from_end(n) - returns the value of the node at nth position from the end of the list
  - [ ] reverse() - reverses the list
  - [ ] remove_value(value) - removes the first item in the list with this value
- [ ] Doubly-linked List
  - [Description (video)](https://www.coursera.org/lecture/data-structures/doubly-linked-lists-jpGKD)
  - No need to implement

### 字符串

- 字符串匹配
  - BF(Brute Force)
  - RK(Rabin-Karp)
  - [KMP](https://wiki.jikexueyuan.com/project/kmp-algorithm/)
  - [boyer_moore](http://www.cs.jhu.edu/~langmea/resources/lecture_notes/boyer_moore.pdf)
  - AC 自动机

## 逻辑数据结构

### 堆栈（Stack）

- [ ] [堆栈（视频）](https://www.coursera.org/learn/data-structures/lecture/UdKzQ/stacks)
- [ ] [使用堆栈 —— 后进先出（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Using-stacks-last-first-out/149042/177120-4.html)
- [ ] 可以不实现，因为使用数组来实现并不重要

### 队列（Queue）

- [ ] [使用队列 —— 先进先出（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Using-queues-first-first-out/149042/177122-4.html)
- [ ] [队列（视频）](https://www.coursera.org/learn/data-structures/lecture/EShpq/queue)
- [ ] [原型队列/先进先出（FIFO）](https://en.wikipedia.org/wiki/Circular_buffer)
- [ ] [优先级队列（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Priority-queues-deques/149042/177123-4.html)
- [ ] 使用含有尾部指针的链表来实现:
  - enqueue(value) —— 在尾部添加值
  - dequeue() —— 删除最早添加的元素并返回其值（首部元素）
  - empty()
- [ ] 使用固定大小的数组实现：
  - enqueue(value) —— 在可容的情况下添加元素到尾部
  - dequeue() —— 删除最早添加的元素并返回其值
  - empty()
  - full()
- [ ] 花销：
  - 在糟糕的实现情况下，使用链表所实现的队列，其入列和出列的时间复杂度将会是 O(n)。因为，你需要找到下一个元素，以致循环整个队列
  - enqueue：O(1)（平摊（amortized）、链表和数组 [探测（probing）]）
  - dequeue：O(1)（链表和数组）
  - empty：O(1)（链表和数组）

### 哈希表（Hash table）

- [ ] 视频：

  - [ ] [链式哈希表（视频）](https://www.youtube.com/watch?v=0M_kIqhwbFo&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=8)
  - [ ] [Table Doubling 和 Karp-Rabin（视频）](https://www.youtube.com/watch?v=BRO7mVIFt08&index=9&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [Open Addressing 和密码型哈希（Cryptographic Hashing）（视频）](https://www.youtube.com/watch?v=rvdJDijO2Ro&index=10&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [PyCon 2010：The Mighty Dictionary（视频）](https://www.youtube.com/watch?v=C4Kc8xzcA68)
  - [ ] [（进阶）随机取样（Randomization）：全域哈希（Universal Hashing）& 完美哈希（Perfect Hashing）（视频）](https://www.youtube.com/watch?v=z0lJ2k0sl1g&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=11)
  - [ ] [（进阶）完美哈希（Perfect hashing）（视频）](https://www.youtube.com/watch?v=N0COwN14gt0&list=PL2B4EEwhKD-NbwZ4ezj7gyc_3yNrojKM9&index=4)

- [ ] 在线课程：

  - [ ] [哈希函数的掌握（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Understanding-hash-functions/149042/177126-4.html)
  - [ ] [使用哈希表（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Using-hash-tables/149042/177127-4.html)
  - [ ] [哈希表的支持（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Supporting-hashing/149042/177128-4.html)
  - [ ] [哈希表的语言支持（视频）](https://www.lynda.com/Developer-Programming-Foundations-tutorials/Language-support-hash-tables/149042/177129-4.html)
  - [ ] [基本哈希表（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/m7UuP/core-hash-tables)
  - [ ] [数据结构（视频）](https://www.coursera.org/learn/data-structures/home/week/3)
  - [ ] [电话薄问题（Phone Book Problem）（视频）](https://www.coursera.org/learn/data-structures/lecture/NYZZP/phone-book-problem)
  - [ ] 分布式哈希表：
    - [Dropbox 中的瞬时上传及存储优化（视频）](https://www.coursera.org/learn/data-structures/lecture/DvaIb/instant-uploads-and-storage-optimization-in-dropbox)
    - [分布式哈希表（视频）](https://www.coursera.org/learn/data-structures/lecture/tvH8H/distributed-hash-tables)

- [ ] 使用线性探测的数组去实现
  - hash(k, m) —— m 是哈希表的大小
  - add(key, value) —— 如果 key 已存在则更新值
  - exists(key)
  - get(key)
  - remove(key)

## 树（Trees）

- ### 树 —— 笔记 & 背景

  - [ ] [系列：基本树（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/ovovP/core-trees)
  - [ ] [系列：树（视频）](https://www.coursera.org/learn/data-structures/lecture/95qda/trees)
  - 基本的树形结构
  - 遍历
  - 操作算法
  - BFS（广度优先检索，breadth-first search）
    - [MIT（视频）](https://www.youtube.com/watch?v=s-CYnVz-uh4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=13)
    - 层序遍历（使用队列的 BFS 算法）
      - 时间复杂度： O(n)
      - 空间复杂度：
        - 最好情况： O(1)
        - 最坏情况：O(n/2)=O(n)
  - DFS（深度优先检索，depth-first search）
    - [MIT（视频）](https://www.youtube.com/watch?v=AfSk24UTFS8&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=14)
    - 笔记：
      - 时间复杂度：O(n)
      - 空间复杂度：
        - 最好情况：O(log n) - 树的平均高度
        - 最坏情况：O(n)
    - 中序遍历（DFS：左、节点本身、右）
    - 后序遍历（DFS：左、右、节点本身）
    - 先序遍历（DFS：节点本身、左、右）

- ### 二叉查找树（Binary search trees）：BSTs

  - [ ] [二叉查找树概览（视频）](https://www.youtube.com/watch?v=x6At0nzX92o&index=1&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
  - [ ] [系列（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/p82sw/core-introduction-to-binary-search-trees)
    - 从符号表开始到 BST 程序
  - [ ] [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/E7cXP/introduction)
  - [ ] [MIT（视频）](https://www.youtube.com/watch?v=9Jry5-82I68)
  - C/C++:
    - [ ] [二叉查找树 —— 在 C/C++ 中实现（视频）](https://www.youtube.com/watch?v=COZK7NATh4k&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=28)
    - [ ] [BST 的实现 —— 在堆栈和堆中的内存分配（视频）](https://www.youtube.com/watch?v=hWokyBoo0aI&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=29)
    - [ ] [在二叉查找树中找到最小和最大的元素（视频）](https://www.youtube.com/watch?v=Ut90klNN264&index=30&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [寻找二叉树的高度（视频）](https://www.youtube.com/watch?v=_pnqMz5nrRs&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=31)
    - [ ] [二叉树的遍历 —— 广度优先和深度优先策略（视频）](https://www.youtube.com/watch?v=9RHO6jU--GU&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=32)
    - [ ] [二叉树：层序遍历（视频）](https://www.youtube.com/watch?v=86g8jAQug04&index=33&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [二叉树的遍历：先序、中序、后序（视频）](https://www.youtube.com/watch?v=gm8DUJJhmY4&index=34&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [判断一棵二叉树是否为二叉查找树（视频）](https://www.youtube.com/watch?v=yEwSGhSsT0U&index=35&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [从二叉查找树中删除一个节点（视频）](https://www.youtube.com/watch?v=gcULXE7ViZw&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=36)
    - [ ] [二叉查找树中序遍历的后继者（视频）](https://www.youtube.com/watch?v=5cPbNCrdotA&index=37&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
  - [ ] 实现：
    - [ ] insert // 往树上插值
    - [ ] get_node_count // 查找树上的节点数
    - [ ] print_values // 从小到大打印树中节点的值 (inorder)
    - [ ] get_successor // 返回给定值的后继者，若没有则返回-1
    - [ ] delete_tree
    - [ ] delete_value
    - [ ] is_in_tree // 如果值存在于树中则返回 true
    - [ ] get_min // 返回树上的最小值
    - [ ] get_max // 返回树上的最大值
    - [ ] get_height // 返回节点所在的高度（如果只有一个节点，那么高度则为 1）

- ### 堆（Heap） / 优先级队列（Priority Queue） / 二叉堆（Binary Heap）

  - 可视化是一棵树，但通常是以线性的形式存储（数组、链表）
  - [ ] [堆](<https://en.wikipedia.org/wiki/Heap_(data_structure)>)
  - [ ] [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/2OpTs/introduction)
  - [ ] [无知的实现（视频）](https://www.coursera.org/learn/data-structures/lecture/z3l9N/naive-implementations)
  - [ ] [二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/GRV2q/binary-trees)
  - [ ] [关于树高的讨论（视频）](https://www.coursera.org/learn/data-structures/supplement/S5xxz/tree-height-remark)
  - [ ] [基本操作（视频）](https://www.coursera.org/learn/data-structures/lecture/0g1dl/basic-operations)
  - [ ] [完全二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/gl5Ni/complete-binary-trees)
  - [ ] [伪代码（视频）](https://www.coursera.org/learn/data-structures/lecture/HxQo9/pseudocode)
  - [ ] [堆排序 —— 跳到起点（视频）](https://youtu.be/odNJmw5TOEE?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3291)
  - [ ] [堆排序（视频）](https://www.coursera.org/learn/data-structures/lecture/hSzMO/heap-sort)
  - [ ] [构建一个堆（视频）](https://www.coursera.org/learn/data-structures/lecture/dwrOS/building-a-heap)
  - [ ] [MIT：堆与堆排序（视频）](https://www.youtube.com/watch?v=B7hVxCmfPtM&index=4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [CS 61B Lecture 24：优先级队列（视频）](https://www.youtube.com/watch?v=yIUFT6AKBGE&index=24&list=PL4BBB74C7D2A1049C)
  - [ ] [构建线性时间复杂度的堆（大顶堆）](https://www.youtube.com/watch?v=MiyLo8adrWw)
  - [ ] 实现一个大顶堆：
    - [ ] insert
    - [ ] sift_up —— 用于插入元素
    - [ ] get_max —— 返回最大值但不移除元素
    - [ ] get_size() —— 返回存储的元素数量
    - [ ] is_empty() —— 若堆为空则返回 true
    - [ ] extract_max —— 返回最大值并移除
    - [ ] sift_down —— 用于获取最大值元素
    - [ ] remove(i) —— 删除指定索引的元素
    - [ ] heapify —— 构建堆，用于堆排序
    - [ ] heap_sort() —— 拿到一个未排序的数组，然后使用大顶堆进行就地排序
      - 注意：若用小顶堆可节省操作，但导致空间复杂度加倍。（无法做到就地）

- ### 字典树（Tries）

  - 需要注意的是，字典树各式各样。有些有前缀，而有些则没有。有些使用字符串而不使用比特位来追踪路径。
  - 阅读代码，但不实现。
  - [ ] [数据结构笔记及编程技术](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#Tries)
  - [ ] 短课程视频：
    - [ ] [对字典树的介绍（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/08Xyf/core-introduction-to-tries)
    - [ ] [字典树的性能（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/PvlZW/core-performance-of-tries)
    - [ ] [实现一棵字典树（视频）](https://www.coursera.org/learn/data-structures-optimizing-performance/lecture/DFvd3/core-implementing-a-trie)
  - [ ] [字典树：一个被忽略的数据结构](https://www.toptal.com/java/the-trie-a-neglected-data-structure)
  - [ ] [高级编程 —— 使用字典树](https://www.topcoder.com/community/data-science/data-science-tutorials/using-tries/)
  - [ ] [标准教程（现实中的用例）（视频）](https://www.youtube.com/watch?v=TJ8SkcUSdbU)
  - [ ] [MIT，高阶数据结构，使用字符串追踪路径（可事半功倍）](https://www.youtube.com/watch?v=NinWEPPrkDQ&index=16&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf)

- ### 平衡查找树（Balanced search trees）

  - 掌握至少一种平衡查找树（并懂得如何实现）：
  - “在各种平衡查找树当中，AVL 树和 2-3 树已经成为了过去，而红黑树（red-black trees）看似变得越来越受人青睐。这种令人特别感兴趣的数据结构，亦称伸展树（splay tree）。它可以自我管理，且会使用轮换来移除任何访问过根节点的 key。” —— Skiena
  - 因此，在各种各样的平衡查找树当中，我选择了伸展树来实现。虽然，通过我的阅读，我发现在 Google 的面试中并不会被要求实现一棵平衡查找树。但是，为了胜人一筹，我们还是应该看看如何去实现。在阅读了大量关于红黑树的代码后，我才发现伸展树的实现确实会使得各方面更为高效。
    - 伸展树：插入、查找、删除函数的实现，而如果你最终实现了红黑树，那么请尝试一下：
    - 跳过删除函数，直接实现搜索和插入功能
  - 我希望能阅读到更多关于 B 树的资料，因为它也被广泛地应用到大型的数据库当中。
  - [ ] [自平衡二叉查找树](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)

  - [ ] **AVL 树**

    - 实际中：我能告诉你的是，该种树并无太多的用途，但我能看到有用的地方在哪里：
      AVL 树是另一种平衡查找树结构。其可支持时间复杂度为 O(log n) 的查询、插入及删除。
      它比红黑树严格意义上更为平衡，从而导致插入和删除更慢，但遍历却更快。正因如此，才彰显其结构的魅力。
      只需要构建一次，就可以在不重新构造的情况下读取，适合于实现诸如语言字典（或程序字典，如一个汇编程序或解释程序的操作码）。
    - [ ] [MIT AVL 树 / AVL 树的排序（视频）](https://www.youtube.com/watch?v=FNeL18KsWPc&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=6)
    - [ ] [AVL 树（视频）](https://www.coursera.org/learn/data-structures/lecture/Qq5E0/avl-trees)
    - [ ] [AVL 树的实现（视频）](https://www.coursera.org/learn/data-structures/lecture/PKEBC/avl-tree-implementation)
    - [ ] [分离与合并](https://www.coursera.org/learn/data-structures/lecture/22BgE/split-and-merge)

  - [ ] **伸展树**

    - 实际中：伸展树一般用于缓存、内存分配者、路由器、垃圾回收者、数据压缩、ropes（字符串的一种替代品，用于存储长串的文本字符）、Windows NT（虚拟内存、网络及文件系统）等的实现。
    - [ ] [CS 61B：伸展树（Splay trees）（视频）](https://www.youtube.com/watch?v=Najzh1rYQTo&index=23&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd)
    - [ ] MIT 教程：伸展树（Splay trees）：
      - 该教程会过于学术，但请观看到最后的 10 分钟以确保掌握。
      - [视频](https://www.youtube.com/watch?v=QnPl_Y6EqMo)

  - [ ] **2-3 查找树**

    - 实际中：2-3 树的元素插入非常快速，但却有着查询慢的代价（因为相比较 AVL 树来说，其高度更高）。
    - 你会很少用到 2-3 树。这是因为，其实现过程中涉及到不同类型的节点。因此，人们更多地会选择红黑树。
    - [ ] [2-3 树的直感与定义（视频）](https://www.youtube.com/watch?v=C3SsdUqasD4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=2)
    - [ ] [2-3 树的二元观点](https://www.youtube.com/watch?v=iYvBtGKsqSg&index=3&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
    - [ ] [2-3 树（学生叙述）（视频）](https://www.youtube.com/watch?v=TOb1tuEZ2X4&index=5&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)

  - [ ] **2-3-4 树 (亦称 2-4 树)**

    - 实际中：对于每一棵 2-4 树，都有着对应的红黑树来存储同样顺序的数据元素。在 2-4 树上进行插入及删除操作等同于在红黑树上进行颜色翻转及轮换。这使得 2-4 树成为一种用于掌握红黑树背后逻辑的重要工具。这就是为什么许多算法引导文章都会在介绍红黑树之前，先介绍 2-4 树，尽管**2-4 树在实际中并不经常使用**。
    - [ ] [CS 61B Lecture 26：平衡查找树（视频）](https://www.youtube.com/watch?v=zqrqYXkth6Q&index=26&list=PL4BBB74C7D2A1049C)
    - [ ] [自底向上的 2-4 树（视频）](https://www.youtube.com/watch?v=DQdMYevEyE4&index=4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
    - [ ] [自顶向下的 2-4 树（视频）](https://www.youtube.com/watch?v=2679VQ26Fp4&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=5)

  - [ ] **B 树**

    - 有趣的是：为啥叫 B 仍然是一个神秘。因为 B 可代表波音（Boeing）、平衡（Balanced）或 Bayer（联合创造者）
    - 实际中：B 树会被广泛适用于数据库中，而现代大多数的文件系统都会使用到这种树（或变种)。除了运用在数据库中，B 树也会被用于文件系统以快速访问一个文件的任意块。但存在着一个基本的问题，那就是如何将文件块 i 转换成一个硬盘块（或一个柱面-磁头-扇区）上的地址。
    - [ ] [B 树](https://en.wikipedia.org/wiki/B-tree)
    - [ ] [B 树的介绍（视频）](https://www.youtube.com/watch?v=I22wEC1tTGo&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6&index=6)
    - [ ] [B 树的定义及其插入操作（视频）](https://www.youtube.com/watch?v=s3bCdZGrgpA&index=7&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
    - [ ] [B 树的删除操作（视频）](https://www.youtube.com/watch?v=svfnVhJOfMc&index=8&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
    - [ ] [MIT 6.851 —— 内存层次模块（Memory Hierarchy Models）（视频）](https://www.youtube.com/watch?v=V3omVLzI0WE&index=7&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf)
      - 覆盖有高速缓存参数无关型（cache-oblivious）B 树和非常有趣的数据结构
      - 头 37 分钟讲述的很专业，或许可以跳过（B 指块的大小、即缓存行的大小）

  - [ ] **红黑树**

    - 实际中：红黑树提供了在最坏情况下插入操作、删除操作和查找操作的时间保证。这些时间值的保障不仅对时间敏感型应用有用，例如实时应用，还对在其他数据结构中块的构建非常有用，而这些数据结构都提供了最坏情况下的保障；例如，许多用于计算几何学的数据结构都可以基于红黑树，而目前 Linux 系统所采用的完全公平调度器（the Completely Fair Scheduler）也使用到了该种树。在 Java 8 中，红黑树也被用于存储哈希列表集合中相同的数据，而不是使用链表及哈希码。
    - [ ] [Aduni —— 算法 —— 课程 4（该链接直接跳到开始部分）（视频）](https://youtu.be/1W3x0f_RmUo?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3871)
    - [ ] [Aduni —— 算法 —— 课程 5（视频）](https://www.youtube.com/watch?v=hm2GHwyKF1o&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=5)
    - [ ] [黑树（Black Tree）](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
    - [ ] [二分查找及红黑树的介绍](https://www.topcoder.com/community/data-science/data-science-tutorials/an-introduction-to-binary-search-and-red-black-trees/)

  - [ ] N 叉树（K 叉树、M 叉树）
  - 注意：N 或 K 指的是分支系数（即树的最大分支数）：
    - 二叉树是一种分支系数为 2 的树
    - 2-3 树是一种分支系数为 3 的树
  - [ ] [K 叉树](https://en.wikipedia.org/wiki/K-ary_tree)

## 图（Graphs）

图论能解决计算机科学里的很多问题，所以这一节会比较长，像树和排序的部分一样。

- Yegge 的笔记:

  - 有 3 种基本方式在内存里表示一个图:
    - 对象和指针
    - 矩阵
    - 邻接表
  - 熟悉以上每一种图的表示法，并了解各自的优缺点
  - 宽度优先搜索和深度优先搜索 - 知道它们的计算复杂度和设计上的权衡以及如何用代码实现它们
  - 遇到一个问题时，首先尝试基于图的解决方案，如果没有再去尝试其他的。

- [ ] Skiena 教授的课程 - 很不错的介绍:

  - [ ] [CSE373 2012 - 课程 11 - 图的数据结构 (video)](https://www.youtube.com/watch?v=OiXxhDrFruw&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=11)
  - [ ] [CSE373 2012 - 课程 12 - 广度优先搜索 (video)](https://www.youtube.com/watch?v=g5vF8jscteo&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=12)
  - [ ] [CSE373 2012 - 课程 13 - 图的算法 (video)](https://www.youtube.com/watch?v=S23W6eTcqdY&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=13)
  - [ ] [CSE373 2012 - 课程 14 - 图的算法 (1) (video)](https://www.youtube.com/watch?v=WitPBKGV0HY&index=14&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [CSE373 2012 - 课程 15 - 图的算法 (2) (video)](https://www.youtube.com/watch?v=ia1L30l7OIg&index=15&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [CSE373 2012 - 课程 16 - 图的算法 (3) (video)](https://www.youtube.com/watch?v=jgDOQq6iWy8&index=16&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)

- [ ] 图 (复习和其他):

  - [ ] [6.006 单源最短路径问题 (video)](https://www.youtube.com/watch?v=Aa2sqUhIn-E&index=15&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [6.006 Dijkstra 算法 (video)](https://www.youtube.com/watch?v=2E7MmKv0Y24&index=16&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [6.006 Bellman-Ford 算法(video)](https://www.youtube.com/watch?v=ozsuci5pIso&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=17)
  - [ ] [6.006 Dijkstra 效率优化 (video)](https://www.youtube.com/watch?v=CHvQ3q_gJ7E&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=18)
  - [ ] [Aduni: 图的算法 I - 拓扑排序, 最小生成树, Prim 算法 - 第六课 (video)](https://www.youtube.com/watch?v=i_AQT_XfvD8&index=6&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
  - [ ] [Aduni: 图的算法 II - 深度优先搜索, 广度优先搜索, Kruskal 算法, 并查集数据结构 - 第七课 (video)](https://www.youtube.com/watch?v=ufj5_bppBsA&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=7)
  - [ ] [Aduni: 图的算法 III: 最短路径 - 第八课 (video)](https://www.youtube.com/watch?v=DiedsPsMKXc&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=8)
  - [ ] [Aduni: 图的算法. IV: 几何算法介绍 - 第九课 (video)](https://www.youtube.com/watch?v=XIAQRlNkJAw&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=9)
  - [ ] [CS 61B 2014 (从 58:09 开始) (video)](https://youtu.be/dgjX4HdMI-Q?list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&t=3489)
  - [ ] [CS 61B 2014: 加权图 (video)](https://www.youtube.com/watch?v=aJjlQCFwylA&list=PL-XXv-cvA_iAlnI-BQr9hjqADPBtujFJd&index=19)
  - [ ] [贪心算法: 最小生成树 (video)](https://www.youtube.com/watch?v=tKwnms5iRBU&index=16&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
  - [ ] [图的算法之强连通分量 Kosaraju 算法 (video)](https://www.youtube.com/watch?v=RpgcYiky7uw)

- 完整的 Coursera 课程:

  - [ ] [图的算法 (video)](https://www.coursera.org/learn/algorithms-on-graphs/home/welcome)

- Yegge: 如果有机会，可以试试研究更酷炫的算法:

  - [ ] Dijkstra 算法 - 上文 - 6.006
  - [ ] A\* 算法
    - [ ] [A\* 算法](https://en.wikipedia.org/wiki/A*_search_algorithm)
    - [ ] [A\* 寻路教程 (video)](https://www.youtube.com/watch?v=KNXfSOx4eEE)
    - [ ] [A\* 寻路 (E01: 算法解释) (video)](https://www.youtube.com/watch?v=-L-WgKMFuhE)

- 我会实现:
  - [ ] DFS 邻接表 (递归)
  - [ ] DFS 邻接表 (栈迭代)
  - [ ] DFS 邻接矩阵 (递归)
  - [ ] DFS 邻接矩阵 (栈迭代)
  - [ ] BFS 邻接表
  - [ ] BFS 邻接矩阵
  - [ ] 单源最短路径问题 (Dijkstra)
  - [ ] 最小生成树
  - 基于 DFS 的算法 (根据上文 Aduni 的视频):
    - [ ] 检查环 (我们会先检查是否有环存在以便做拓扑排序)
    - [ ] 拓扑排序
    - [ ] 计算图中的连通分支
    - [ ] 列出强连通分量
    - [ ] 检查双向图

## 算法

### 排序

#### 初级排序

- 选择排序
- 插入排序
- 希尔排序

#### 归并排序

#### 快速排序

#### 堆排序

- [ ] 笔记:

  - 实现各种排序 & 知道每种排序的最坏、最好和平均的复杂度分别是什么场景:
    - 不要用冒泡排序 - 大多数情况下效率感人 - 时间复杂度 O(n^2), 除非 n <= 16
  - [ ] 排序算法的稳定性 ("快排是稳定的么?")
    - [排序算法的稳定性](https://en.wikipedia.org/wiki/Sorting_algorithm#Stability)
    - [排序算法的稳定性](http://stackoverflow.com/questions/1517793/stability-in-sorting-algorithms)
    - [排序算法的稳定性](http://stackoverflow.com/questions/1517793/stability-in-sorting-algorithms)
    - [排序算法的稳定性](http://www.geeksforgeeks.org/stability-in-sorting-algorithms/)
    - [排序算法 - 稳定性](http://homepages.math.uic.edu/~leon/cs-mcs401-s08/handouts/stability.pdf)
  - [ ] 哪种排序算法可以用链表？哪种用数组？哪种两者都可？
    - 并不推荐对一个链表排序，但归并排序是可行的.
    - [链表的归并排序](http://www.geeksforgeeks.org/merge-sort-for-linked-list/)

- 关于堆排序，请查看前文堆的数据结构部分。堆排序很强大，不过是非稳定排序。

- [ ] [冒泡排序 (video)](https://www.youtube.com/watch?v=P00xJgWzz2c&index=1&list=PL89B61F78B552C1AB)
- [ ] [冒泡排序分析 (video)](https://www.youtube.com/watch?v=ni_zk257Nqo&index=7&list=PL89B61F78B552C1AB)
- [ ] [插入排序 & 归并排序 (video)](https://www.youtube.com/watch?v=Kg4bqzAqRBM&index=3&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
- [ ] [插入排序 (video)](https://www.youtube.com/watch?v=c4BRHC7kTaQ&index=2&list=PL89B61F78B552C1AB)
- [ ] [归并排序 (video)](https://www.youtube.com/watch?v=GCae1WNvnZM&index=3&list=PL89B61F78B552C1AB)
- [ ] [快排 (video)](https://www.youtube.com/watch?v=y_G9BkAm6B8&index=4&list=PL89B61F78B552C1AB)
- [ ] [选择排序 (video)](https://www.youtube.com/watch?v=6nDMgr0-Yyo&index=8&list=PL89B61F78B552C1AB)

- [ ] 斯坦福大学关于排序算法的视频:

  - [ ] [课程 15 | 编程抽象 (video)](https://www.youtube.com/watch?v=ENp00xylP7c&index=15&list=PLFE6E58F856038C69)
  - [ ] [课程 16 | 编程抽象 (video)](https://www.youtube.com/watch?v=y4M9IVgrVKo&index=16&list=PLFE6E58F856038C69)

- [ ] Shai Simonson 视频, [Aduni.org](http://www.aduni.org/):

  - [ ] [算法 - 排序 - 第二讲 (video)](https://www.youtube.com/watch?v=odNJmw5TOEE&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=2)
  - [ ] [算法 - 排序 2 - 第三讲 (video)](https://www.youtube.com/watch?v=hj8YKFTFKEE&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=3)

- [ ] Steven Skiena 关于排序的视频:

  - [ ] [课程从 26:46 开始 (video)](https://youtu.be/ute-pmMkyuk?list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&t=1600)
  - [ ] [课程从 27:40 开始 (video)](https://www.youtube.com/watch?v=yLvp-pB8mak&index=8&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [课程从 35:00 开始 (video)](https://www.youtube.com/watch?v=q7K9otnzlfE&index=9&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [课程从 23:50 开始 (video)](https://www.youtube.com/watch?v=TvqIGu9Iupw&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=10)

- [ ] 加州大学伯克利分校（UC Berkeley） 大学课程:

  - [ ] [CS 61B 课程 29: 排序 I (video)](https://www.youtube.com/watch?v=EiUvYS2DT6I&list=PL4BBB74C7D2A1049C&index=29)
  - [ ] [CS 61B 课程 30: 排序 II (video)](https://www.youtube.com/watch?v=2hTY3t80Qsk&list=PL4BBB74C7D2A1049C&index=30)
  - [ ] [CS 61B 课程 32: 排序 III (video)](https://www.youtube.com/watch?v=Y6LOLpxg6Dc&index=32&list=PL4BBB74C7D2A1049C)
  - [ ] [CS 61B 课程 33: 排序 V (video)](https://www.youtube.com/watch?v=qNMQ4ly43p4&index=33&list=PL4BBB74C7D2A1049C)

- [ ] - 归并排序:
  - [ ] [使用外部数组](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/sorting/mergesort.c)
  - [ ] [对原数组直接排序](https://github.com/jwasham/practice-cpp/blob/master/merge_sort/merge_sort.cc)
- [ ] - 快速排序:
  - [ ] [实现](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/randomization/quick.c)
  - [ ] [实现](https://github.com/jwasham/practice-c/blob/master/quick_sort/quick_sort.c)

- [ ] 实现:

  - [ ] 归并：平均和最差情况的时间复杂度为 O(n log n)。
  - [ ] 快排：平均时间复杂度为 O(n log n)。
  - 选择排序和插入排序的最坏、平均时间复杂度都是 O(n^2)。
  - 关于堆排序，请查看前文堆的数据结构部分。

- [ ] 有兴趣的话，还有一些补充 - 但并不是必须的:
  - [ ] [基数排序](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#radixSort)
  - [ ] [基数排序 (video)](https://www.youtube.com/watch?v=xhr26ia4k38)
  - [ ] [基数排序, 计数排序 (线性时间内) (video)](https://www.youtube.com/watch?v=Nz1KZXbghj8&index=7&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [随机算法: 矩阵相乘, 快排, Freivalds' 算法 (video)](https://www.youtube.com/watch?v=cNB2lADK3_s&index=8&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
  - [ ] [线性时间内的排序 (video)](https://www.youtube.com/watch?v=pOKy3RZbSws&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf&index=14)
