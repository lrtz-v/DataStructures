# Mysql

## 索引

- 最左前缀
  - 检索数据时从联合索引的最左边开始匹配
  - 配合索引下推

## 锁

- InnoDB 的锁

  - 库锁：对全库加锁，让整个库处于只读状态
    - 使用场景
      - 全库逻辑备份
  - 表锁：限制表的读写
    - 在还没有出现更细粒度的锁的时候，表锁是最常用的处理并发的方式
  - 表元数据锁（MDL）
    - 访问表时，加 MDL 读锁
    - 更改表结构时，加 MDL 写锁， 读写互斥，写写互斥
  - 行读写锁：在需要的时候才加上的，等到事务结束时释放
    - 死锁检测
  - 间隙锁
    - 锁的就是两个值之间的空隙
    - 跟间隙锁存在冲突关系的，是“往这个间隙中插入一个记录”这个操作。间隙锁之间都不存在冲突关系
    - 影响了并发度
  - next-key 锁：间隙锁和行锁合称 next-key lock， 前开后闭区间

- 乐观锁与悲观锁

  - 乐观锁： 适合读取操作比较频繁的场景
    - 使用条件限制
    - 数据表加 version 字段
  - 悲观锁：适合写入操作比较频繁
    - 共享锁
      - 读锁，允许事务读一行数据
    - 排他锁
      - 写锁，允许事务删除或更新一行数据

- binlog

  - 使用场景

    - 主从复制
    - 数据恢复：借助 mysqlbinlog 工具

  - 刷盘时机：sync_binlog 控制

    - 0：不去强制要求，由系统自行判断何时写入磁盘
    - 1：每次 commit 的时候都要将 binlog 写入磁盘
    - N：每 N 个事务，才会将 binlog 写入磁盘

  - 格式
    - 基于行的日志：记录成每一行数据被修改的形式
      - 优点：
        - 不记录执行的 sql 语句的上下文相关的信息，仅需要记录那一条记录被修改成什么了
      - 缺点：
        - 所有的执行的语句当记录到日志中的时候，都将以每行记录的修改来记录，这样可能会产生大量的日志内容
    - 基于 SQL 的日志：记录每一条会修改数据的 SQL
      - 优点
        - 不记录每一行数据的变化，文件较小
        - 可以用于实时的还原，而不仅仅用于复制
        - 主从版本可以不一样，从服务器版本可以比主服务器版本高
      - 缺点：
        - 不是所有的 UPDATE 语句都能被复制，尤其是包含不确定操作的时候
        - 为了让这些语句在 slave 端也能正确执行，还必须记录每条语句在执行的时候的一些相关信息，也就是上下文信息，如部分函数
    - Mixed：根据执行的每一条具体的 SQL 语句来区分对待记录的日志形式，在 statement 和 row 之间选择一种

- redo log

  - 物理日志，记录的是“在某个数据页上做了什么修改”
  - 使用场景
    - 用于保证 crash-safe 能力
  - 刷盘时间：根据 innodb_flush_log_at_trx_commit 设置
    - 设置为 0 的时候，表示每次事务提交时都只是把 redo log 留在 redo log buffer 中
    - 设置为 1 的时候，表示每次事务提交时都将 redo log 直接持久化到磁盘
    - 设置为 2 的时候，表示每次事务提交时都只是把 redo log 写到 page cache
    - 定时写入
      - 1.InnoDB 后台线程，每隔 1 秒，就会把 redo log buffer 中的日志，调用 write 写到文件系统的 page cache，然后调用 fsync 持久化到磁盘
      - 2.redo log buffer 占用的空间即将达到 innodb_log_buffer_size 一半的时候，后台线程会主动写 page cache
  - 为什么两阶段提交可以保证数据的一致性

    - 写 redo log 之前崩溃
      - 事务还没有开始提交，所以奔溃恢复跟该事务没有关系
    - redolog 处于 prepare 状态，写 binlog 之前崩溃
      - 重启恢复：redo log 的事务不完整，回滚
      - 备份恢复：没有 binlog，数据一致。
    - commit redo log 的时候奔溃
      - 重启恢复：虽没有 commit，但满足 prepare 和 binlog 完整，所以重启后会自动 commit
      - 备份恢复：有 binlog，数据一致。

  - 重启恢复流程
    - 读取 redo log， 从 checkpoint 对日志重放
    - 检测哪些事务是完整的并且处于 prepare 状态
    - 根据事务 ID（XID），对照 binlog， 如果找不到，则回滚

- undo log

  - 使用场景
    - 事务回滚
    - MVCC
  - 写入时机
    - DML 操作修改聚簇索引前，记录 undo log
    - 二级索引记录的修改，不记录 undo log

- relay log

  - slave 将 master 的 binary log events 追加到它的中继日志（relay log）
  - 解析 relay-log 的内容，执行

- crash safe

  - 事务提交过程中任何阶段，MySQL 突然奔溃，重启后都能保证事务的完整性，已提交的数据不会丢失，未提交完整的数据会自动进行回滚

## MyISAM 和 InnoDB 的区别

- 存储结构

  - MyISAM：表在磁盘上存储称为三个文件，表结构、数据、索引
  - InnoDB：数据文件按主键聚集

- 事务支持

  - MyISAM：强调性能，每次查询具有原子性，不支持事务
  - InnoDB：支持事务

- 锁差异

  - MyISAM：只支持表锁
  - InnoDB：支持行锁

- 外键

  - MyISAM：不支持
  - InnoDB：支持

- 行数

  - MyISAM：用一个变量保存了整个表的行数
  - InnoDB：不保存表的具体行数

- 可移植性、备份与恢复

  - MyISAM：数据以文件的形式存储，数据转移会方便；在备份和恢复时可以针对某个表单独操作
  - InnoDB：binlog, redo log

- AUTO_INCREMENT
  - MyISAM：自增值保存在数据文件中
  - InnoDB：自增值保存在了内存里

## Mysql AUTO_INCREMENT

- 自增值持久化

  - 5.7 及以前版本
    - 自增值保存在内存，重启后，取表中的 max(id)，然后将 max(id) + 1 作为自增值
  - 8.0 版本及以后
    - 自增值变更记录在 redo log 中，依靠 redo log 恢复

- 自增值过程

  - 当插入数据行的 id 字段指定为 0、null 或未指定时，使用 AUTO_INCREMEN 值
  - 如果 插入值 < 自增值，那么这个表的自增值不变
  - 如果 插入值 ≥ 自增值，就需要把当前自增值修改为新的自增值

- 自增值计算

  - 从 auto_increment_offset 开始，以 auto_increment_increment 为步长，持续叠加，直到找到第一个大于 X 的值

- 自增值不连续

  - 唯一键冲突
  - 事务回滚
  - 批量插入数据，批量申请 id，会造成 id 浪费

- 自增锁：innodb_autoinc_lock_mode 参数设置
  - innodb_autoinc_lock_mode 为 0
    - 语句执行后才释放锁（会影响并发度）
  - innodb_autoinc_lock_mode 为 1
    - 普通 insert 语句，自增锁在申请主键的动作之后就马上释放
    - 批量插入数据的语句，自增锁还是要等语句结束后才被释放（批量申请策略）
  - innodb_autoinc_lock_mode 为 2
    - 所有的申请自增主键的动作都是申请后就释放锁

## 索引失效

- 选择索引是优化器的工作，目的是找到最优的执行方案，考虑因素：扫描行数，临时表、排序等
- 1.扫描行数预估
  - 根据统计信息（索引的“区分度”）来估算记录数；一个索引上不同的值的个数，我们称之为“基数”；基数越大，索引的区分度越好
  - 基数统计-采样统计
    - InnoDB 默认会选择 N 个数据页，统计这些页面上的不同值，得到一个平均值，然后乘以这个索引的页面数，就得到了这个索引的基数
    - 当变更的数据行数超过 1/M 的时候，会自动触发重新做一次索引统计，M 根据 innodb_stats_persistent 的模式设置而不同
- 2.回表代价
  - 解决办法
    - force index
    - 修改语句，引导 MySQL 使用我们期望的索引
      - 利用 limit、order by
    - 统计数据修正
      - analyze table t 重新统计索引数据
    - 创建更合适的索引，删除旧索引

## 临时表

- UNION 操作
  - 使用临时表暂存数据
  - UNION ALL 则不需要，结果不需要去重，直接作为结果集返回
- group by
  - 使用临时表暂存数据并排序
  - 优化
    - 使用 generated column，构造 group by 使用的列
    - 尽量让 group by 过程用上表的索引
    - 如果 group by 需要统计的数据量不大，尽量只使用内存临时表；也可以通过适当调大 tmp_table_size 参数，来避免用到磁盘临时表
    - 如果数据量实在太大，使用 SQL_BIG_RESULT 这个提示，来告诉优化器直接使用排序算法得到 group by 的结果

## 分布式事务
