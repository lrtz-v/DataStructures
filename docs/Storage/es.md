# ES

## Lucene 索引

- 倒排索引
  - Term index
  - Team Dictionary
  - Posting List

## 索引构建及存储

### 正向索引

- 记录了 DocId 到 FieldValue 的映射关系，通过 DocID 就能直接获取字段值的能力

### 倒排索引

#### Postings

- Terms 列表：记录 Field 中出现的所有的 Terms
- Postings：记录 Term 所对应的所出现哪些文档的文档号，出现次数，位置信息等

#### TermVectors
