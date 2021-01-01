# 深入理解以太坊存储（三）：Patricia Trie

这是本系列的第三篇文章。第一篇文介绍了数据结构[trie](https://en.wikipedia.org/wiki/Trie)，第二篇文章介绍了trie的空间优化版：[radix trie](https://en.wikipedia.org/wiki/Radix_tree)，这篇文章将介绍radix trie的特例：[PATRICIA trie](https://en.wikipedia.org/wiki/Radix_tree#PATRICIA)。（记录历史时刻：写这篇文章的时候，BTC的价格又创新高，突破了$29,000一枚！😲）



## Patricia Trie介绍

PATRICIA是"Practical Algorithm to Retrieve Information Coded in Alphanumeric"的首字母缩写词。这么长看着就吓人，所以这里就不仔细介绍了，感兴趣的读者可以在网上搜索相关资料。不过，如果你已经读过前两篇文章的话，那么不理解这个词的含义也并不影响这篇文章的阅读。

由前一篇文章可知，“radix trie”的radix（后文简称`r`）和我们如何比较键密切相关。如果比较键时每次对比`n`个比特，那么`r = 2^n`。例如，在前一篇文章中，我们约定键是ASCII字符串，在比较键时每次对比8个比特，因此在前一篇文章中实现的radix trie的`r = 2^8`。当`n = 1`时（此时`r = 2`），我们就得到了一棵patricia trie。

关于patricia trie的理论知识就先介绍到这里，接下来我们将在前一章radix trie（`r = 2^8`）代码的基础上，实现`r = 2^4`和`r = 2^1`的radix trie，后者就是本文的主角patricia trie（本文并不考运行虑效率，只考虑开发效率，所以直接用了最省事儿的方式来实现这两个变种树😊）。第一篇文章用了一个简单的关联数组作为例子，本文直接拿过来再次使用：

```json
{
  "foo": 123,
  "bar": 456,
  "baz": 789,
}
```

这一系列文章的配套代码可以在[这里](https://github.com/zxh0/blog/code/go/mympt)找到。



## r=2^4

我们循序渐进，先把键的对比单位从8个比特（一个字节）改成4个比特（也叫一个[nibble](https://en.wikipedia.org/wiki/Nibble)）。如前文所述，可以偷个懒，直接在前一篇文章`RadixTrieNode`的基础之上构建即可。我们把新的结构体叫做`Radix16TrieNode`（实在是没想到更好听的名字😅），下面是它的定义，以及一个“构造函数”：

```go
type Radix16Trie struct {
	Root *RadixTrieNode
}

func NewRadix16Trie() SearchTree {
	return &Radix16Trie{Root: &RadixTrieNode{}}
}
```

插入、查找、遍历操作都非常简单，加工一下键，然后转发给`RadixTrieNode`的相应方法即可，代码如下所示：

```go
func (node *Radix16Trie) Put(key string, val uint) {
	node.Root.Put(utils.ToHex(key), val)
}
func (node *Radix16Trie) Get(key string) uint {
	return node.Root.Get(utils.ToHex(key))
}
func (node *Radix16Trie) ForEach(cb func(string, uint)) {
	node.Root.ForEach(func(k string, v uint) {
		cb(utils.FromHex(k), v)
	})
}
```

插入和查找时，我们把键转换成nibble（用16进制数字表示）串。遍历时，我们把nibble串再转换回ASCII字符串。这两种转换由辅助函数`ToHex()`和`FromHex()`完成，下面是这两个函数的代码：

```go
func ToHex(s string) string {
	return hex.EncodeToString([]byte(s))
}
func FromHex(s string) string {
	bytes, _ := hex.DecodeString(s)
	return string(bytes)
}
```

这样就实现好了，就这么简单。照搬以前的测试，稍微改一改，跑一下：

```go
func TestRadix16Trie(t *testing.T) {
	trie := NewRadix16Trie()
	trie.Put("foo", 123) // 666f6f
	trie.Put("bar", 456) // 626172
	trie.Put("baz", 789) // 62617a
	require.Equal(t, uint(123), trie.Get("foo"))
	require.Equal(t, uint(456), trie.Get("bar"))
	require.Equal(t, uint(789), trie.Get("baz"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}
```

目测没啥问题（如果读者发现代码有bug，欢迎指出，我会修改bug并更新文章），整棵树的内部状态打印出来是下面这样：

```json
{
  "Root": {
    "prefix": "6",
    "kids": {
      "2": {
        "prefix": "617",
        "kids": {
          "2": { "val": 456 },
          "a": { "val": 789 }
        }
      },
      "6": {
        "prefix": "6f6f",
        "val": 123
      }
    }
  }
}
```



## r=2^1

现在让我们把键的对比单位改成一个比特，并把这个结构体叫做`PatriciaTrie`。下面是结构体的定义，以及一个“构造函数”：

```go
type PatriciaTrie struct {
	Root *RadixTrieNode
}

func NewPatriciaTrie() SearchTree {
	return &PatriciaTrie{Root: &RadixTrieNode{}}
}
```

插入、查找、遍历操作同样非常简单，加工一下键，然后转发给`RadixTrieNode`的相应方法即可，代码如下所示：

```go
func (node *PatriciaTrie) Put(key string, val uint) {
	node.Root.Put(utils.ToBin(key), val)
}
func (node *PatriciaTrie) Get(key string) uint {
	return node.Root.Get(utils.ToBin(key))
}
func (node *PatriciaTrie) ForEach(cb func(string, uint)) {
	node.Root.ForEach(func(k string, v uint) {
		cb(utils.FromBin(k), v)
	})
}
```

插入和查找时，我们把键转换成比特（用二进制数字`0`和`1`表示）串。遍历时，我们把比特串再转换回ASCII字符串。这两种转换由辅助函数`ToBin()`和`FromBin()`完成。Go语言没有提供现成的库，得我们自己来实现了。下面是这两个函数的代码：

```go
func ToBin(s string) string {
	bs := make([]byte, 0, len(s)*8)
	for _, b := range s {
		for i := 0; i < 8; i++ {
			if (b<<i)&0x80 > 0 {
				bs = append(bs, byte('1'))
			} else {
				bs = append(bs, byte('0'))
			}
		}
	}
	return string(bs)
}

func FromBin(bs string) string {
	s := make([]byte, len(bs)/8)
	for i := 0; i < len(bs); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			b <<= 1
			if bs[i+j] == '1' {
				b |= 1
			}
		}
		s[i/8] = b
	}
	return string(s)
}
```

还是照搬刚才的测试，稍微改一改，跑一下：

```go
func TestPatriciaTrie(t *testing.T) {
	trie := NewPatriciaTrie()
	trie.Put("foo", 123) // 011001100110111101101111
	trie.Put("bar", 456) // 011000100110000101110010
	trie.Put("baz", 789) // 011000100110000101111010
	require.Equal(t, uint(123), trie.Get("foo"))
	require.Equal(t, uint(456), trie.Get("bar"))
	require.Equal(t, uint(789), trie.Get("baz"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}
```

同样是目测没啥问题，整棵树的内部状态打印出来是下面这样：

```json
{
  "Root": {
    "prefix": "01100",
    "kids": {
      "0": {
        "prefix": "10011000010111",
        "kids": {
          "0": { "prefix": "010", "val": 456 },
          "1": { "prefix": "010", "val": 789 }
        }
      },
      "1": {
        "prefix": "100110111101101111",
        "val": 123
      }
    }
  }
}
```



## 总结/广告

本文介绍了radix trie的两个特例：`r=2^4`和`r=2^1`（后者有自己专门的名字：Patricia Trie），并且利用`r=2^8`的radix trie的实现，给出了这两种特例的实现。到这里为止，我们花了3篇文章分别介绍了trie、radix trie和patricia trie。Trie的基础知识介绍就到此为止了，在下一篇文章中，我们将讨论另外一个有趣的数据结构：merkle tree。

如果你喜欢这种循序渐进、文字+可执行代码结合的写作风格的话，那么作者写的三本书一定不要错过。感兴趣的读者请关注[《自己动手写Java虚拟机》](https://e.jd.com/30231161.html?ebook=1)、[《自己动手实现Lua》](https://item.jd.com/12458778.html)和[《WebAssembly原理与核心技术》](https://item.jd.com/12998058.html)。

<img src="https://pic4.zhimg.com/80/v2-523aa681b72cdf2f8004c583735c5cbf_1440w.jpg" alt="book3" style="zoom:50%;" />

