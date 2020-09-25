# Redis

## Redis数据结构

- String
  - int 8字节长整型
  - SDS
    - embstr 小于44个字节
    - raw 大44个字节
- List
  - zipList
  - linkedList
- Hash
  - zipList
  - hashTable
- Set
  - int set
  - hashTable
- Sorted Set
  - zipList
  - dict存储member对应的score + skipList存储根据score对member的定位结构
- GEO
  - geohash
- HyperLogLog 海量数据统计
- Bitmap
  - bit 数组，统计海量数据

