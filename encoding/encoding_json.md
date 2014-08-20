#json包

import "encoding/json"

---

##简介

json和xml之争一直存在着。

##概览
json实现了JSON对象（RFC4627中定义）的编码和解码。JSON和Go值之间的映射在Marshal和Unmarshal函数的文档中描述。可以参考文章` http://golang.org/doc/articles/json_and_go.html `理解json包。


###func Compact
```go
func Compact(dst *bytes.Buffer, src []byte) error
```
Compact向`dst`中增加JSON-encoded `src`，同时消除那些微不足道的空格。

###func HTMLEscape
```go
func HTMLEscape(dst *bytes.Buffer, src []byte)
```
HTMLEscape向`dst`中增加JSON-encoded `src`，同时字符串中的<, >, &, U+2028 and U+2029 字符转换为 \u003c, \u003e, \u0026, \u2028, \u2029，这样JSON可以安全地嵌入在HTML <脚本>标记。因为历史的原因，web浏览器不会尊重包括`<script>`标签的标准的HTML转义，所以会使用供选择的JSON编码。


###func Indent
```go
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
```
Indent向`dst`中增加带有缩进形式的JSON-encoded `src`。一个JSON对象中的每个元素或者一个JSON array中的每个元素开始是一个新的缩进的行，该行以一个前缀开始，接下来是一个或者更多缩进的复本（根据缩进嵌套）。附加在`dst`后面的数据不会以前缀或者缩进开始，而且没有trailing 换行，这样使得它很容易的嵌入到其他格式化了的JSON数据中。

###func Marshal
```go
func Marshal(v interface{}) ([]byte, error)
```
Marshal返回`v·的JSON编码结果。

Marshal递归遍历`v`的值。如果遇到的值实现了Marshaler 接口而且不是一个空指针，Marshal调用它的MarshalJSON 方法来产生JSON。空指针异常并不是必须的，除了与之类似的mimics（在UnmarshalJSON中的必须的异常）之外。

否则，Marshal 使用下面的类型独立的默认编码：

布尔值编码为JSON booleans。

Floating point、integer和 Number 编码为JSON numbers。

字符串编码为JSON strings。如果遇到一个无效的UTF-8序列，InvalidUTF8Error 会返回。“<”和“>”被转为"\u003c"和"\u003e"来防止一些浏览器误将JSON解析为HTML。因为同样的原因，“&”也被用“\u0026”替换。

数组和切片编码为JSON数组。例外的是[]byte编码为base-64编码的字符串和空切片编码为null JSON对象。

结构编码为JSON对象。每一个导出的结构成员成为了对象的成员除了
```go
- the field's tag is "-", or
- the field is empty and its tag specifies the "omitempty" option.
```
空值为flase、0、任意空指针或者接口值，还有数组、切片、map或者零长度的字符串。对象的默认key字符串是结构成员名字但是在结构的成员tag值中可以具体指定。结构成员tag值中的“json” key是key名字，接下来是可选的逗号和选择项。比如：

```
// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`
```
"string"选项指示了成员是存储为JSON，它在JSON编码的字符串内部。它只应用于sting成员、float point或者integer类型。当与javascript程序交互时，这个编码的额外的层次
有时使用。

```go
Int64String int64 `json:",string"`
```
如果它是一个非空字符串（只包含Unicode字母、数字、美元符号、百分号、连字符、下划线、斜杠），key名字会被使用。

匿名结构成员经常被编码（marshaled ），如果它们的内部导出变量是外部结构的变量的话。这服从于常见的Go的可见规则，修订的规则在下一段话。匿名结构成员（带有一个名字，其名字在它的JSON tag中给出）会被视为有那个名字，而不是匿名。

关于结构成员的Go的可见规则为JSON而修订，当决定哪个成员被marshal 或被unmarshal。如果有很多成员在同一个层次，和那个层次在最小的嵌套层次（还有将会因此成为被常见的Go规则嵌套层次）；以下的额外规则为：

1)在这些成员中，如果有些是JSON-tagged，只有tagged变量才会被考虑，尽管有多种会引起冲突的untagged成员。
2）如果恰好有一个成员（tagged或者不按照第一个规则），它会被选中。
3）除非有多种成员，并且所有的都被忽略；没有错误产生。

处理匿名结构成员是在Go1.1加入的。在Go 1.1之前，匿名结构成员会被忽略。在当前和早先版本中，为了强制忽略的匿名结构成员，会给成员一个JSON tag “-”。

MAP值编码为JSON对象。MAP的key必须是字符串，对象的key直接作为map keys。

指针值编码为指针指向的值。空指针编码为null JSON对象。

接口值编码为在接口中存储的值。空接口编码为null JSON对象。

Channel， complex 以及函数不能被编码为JSON。尝试编码这样的值会导致Marshal 返回UnsupportedTypeError。

JSON不可以代表循环数据结构，Marshal不处理它们。向Marshal传入循环结构会导致无限的递归。

###func MarshalIndent
```go
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```
MarshalIndent和Marshal很像，但是使用了缩进`indent`来格式化输出。

###func Unmarshal
```go
func Unmarshal(data []byte, v interface{}) error
```
Unmarshal解析JSON编码的数据，然后将结果存储在`v`指向的值。

Unmarshal使用Marshal使用的编码的相反行为，按照需要分配map、切片和指针，用以下的附加规则：

为了解码（unmarshal） JSON为一个指针，Unmarshal先处理JSON是JSON literal null的情形。在那种情形下，Unmarshal设置指针为nil。否则，Unmarshal 函数unmarshal JSON为指针所指向的值。如果指针为nil，Unmarshal为指针分配新的值。

为了unmarshal JSON为一个结构，Unmarshal匹配得到的对象keys为Marshal使用的keys（结构成员名或者它的tag），不仅倾向于精确的匹配而且接受大小写敏感的匹配。

为了unmarshal JSON为一个接口值，Unmarshal存储以下的一个于接口值中：

```go
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
```

如果JSON值不适合一个给定的目标类型或者如果JSON标号超出目标类型，Unmarshal跳过那个成员然后尽力完成unmarshalling。如果没有遇到更多严重的错误，Unmarshal返回UnmarshalTypeError，它描述了最早的这种错误。

通过设置GO值为nil，JSON空值unmarshal为接口、map、指针或者切片。因为在JSON中使用null意味着“not present”， unmarshal一个JSON null为任意其他GO类型不会影响值和产生错误。

当 unmarshal引用的字符串，无效的UTF8或者无效的UTF-16代理对不会按照错误对待。然而，它们用Unicode replacement 字符U+FFFD来代替。

>Unmarshal是如何定义存放解码的数据的呢？对于一个给定的 JSON key"Foo"，Unmarshal会查询结构体的域来寻找（in order of preference）：
- 一个带有标签"Foo" 的可导出域（更多关于结构体标签见Go spec）
- 一个名为"Foo" 的可导出域，或者
- 一个名为"FOO" 或者 "FoO 或者其他大小写的匹配"Foo"的可导出域

###type Decoder struct
```go
type Decoder struct {
}
```
Decoder类型读取并解码输入流中的JSON对象。

###func NewDecoder
```go
func NewDecoder(r io.Reader) *Decoder
```
NewDecoder返回一个新的dedocer，这个decoder从`r`读取。

这个decoder引入它自己的缓冲，从`r`读取的数据可能超过JSON值的要求。

###func (*Decoder) Buffered
```go
func (dec *Decoder) Buffered() io.Reader
```
Buffered返回一个reader，这个reader的数据残留在Decoder的缓冲中。reader是有效的，直到下次调用Decode。

###func (*Decoder) Decode
```go
func (dec *Decoder) Decode(v interface{}) error
```
Decode读取下一个JSON编码的值，然后存储到`v`所指的值中。

想了解更多关于JONS向Go值转换的信息，请看Unmarshal 函数的文档。

###func (*Decoder) UseNumber
```go
func (dec *Decoder) UseNumber()
```
UseNumber引起Decoder来将一个数字（作为一个Number而不是一个float64数值）解码成一个接口（interface{}）。

###type Encoder struct
```go
type Encoder struct {
}
```
Encoder向输出流写入JSON对象。

###func NewEncoder
```go
func NewEncoder(w io.Writer) *Encoder
```
NewEncoder返回一个写入`w`的新的encoder。

###func (*Encoder) Encode
```go
func (enc *Encoder) Encode(v interface{}) error
```
Encode将JSON编码的`v`写入到流中，接下来是换行符。

想了解更多关于Go值向JSON转换的信息，请看Marshal 函数的文档。

###type InvalidUTF8Error struct
```go
type InvalidUTF8Error struct {
    S string // the whole string value that caused the error
}
```
在GO1.2之前，当Marshal 尝试编码带有无效UTF-8序列的字符串时，会返回一个InvalidUTF8Error。在Go 1.2中，Marshal通过用 rune类型的U+FFFD（Unicode replacement）来替换非法的字节这种方法来强迫字符串转为UTF-8。这种错误不会产生但是为了向后兼容而保留下来。

###func (*InvalidUTF8Error) Error
```go
func (e *InvalidUTF8Error) Error() string
```

###type InvalidUnmarshalError struct
```go
type InvalidUnmarshalError struct {
    Type reflect.Type
}
```
InvalidUnmarshalError 描述了一个无效的传给Unmarshal的参数（传给Unmarshal的参数必须是一个非空指针）。

###func (*InvalidUnmarshalError) Error
```go
func (e *InvalidUnmarshalError) Error() string
```

###type Marshaler interface
```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```
Marshaler 是一个接口。实现它的类型可以将它们自己编码为一个有效的JSON。

###type MarshalerError struct
```go
type MarshalerError struct {
    Type reflect.Type
    Err  error
}
```

###func (*MarshalerError) Error
```go
func (e *MarshalerError) Error() string
```

###type Number
```go
type Number string
```
一个Number代表了JSON数字。

###func (Number) Float64
```go
func (n Number) Float64() (float64, error)
```
Float64返回float64的数字。

###func (Number) Int64
```go
func (n Number) Int64() (int64, error)
```
返回int64的数字。

###func (Number) String
```go
func (n Number) String() string
```
返回数字的字符串形式。

###type RawMessage
```go
type RawMessage []byte
```
RawMessage是一个原始的编码的JSON对象。它实现了Marshaler 和Unmarshaler ，并可以用来延迟JSON解析或者JSON编码预计算。

###func (*RawMessage) MarshalJSON
```go
func (m *RawMessage) MarshalJSON() ([]byte, error)
```
返回`*m`作为`m`的JSON编码。

###func (*RawMessage) UnmarshalJSON
```go
func (m *RawMessage) UnmarshalJSON(data []byte) error
```
UnmarshalJSON 设置`*m`为data的一份拷贝。

###type SyntaxError struct
```go
type SyntaxError struct {
    Offset int64 // error occurred after reading Offset bytes
}
```
SyntaxError 描述了JSON的语法错误。

###func (*SyntaxError) Error
```go
func (e *SyntaxError) Error() string
```

###type UnmarshalFieldError struct
```go
type UnmarshalFieldError struct {
    Key   string
    Type  reflect.Type
    Field reflect.StructField
}
```
UnmarshalFieldError 描述了一个JSON对象key，这个key引起了一个没有导出的（因此为不可写的）的结构成员（ struct field）（不再使用，为了兼容性而保留）。

###func (*UnmarshalFieldError) Error
```go
func (e *UnmarshalFieldError) Error() string
```

###type UnmarshalTypeError struct
```go
type UnmarshalTypeError struct {
    Value string       // description of JSON value - "bool", "array", "number -5"
    Type  reflect.Type // type of Go value it could not be assigned to
}
```
UnmarshalTypeError描述了这样的一个JSON值，它不适合特定GO类型的值。

###func (*UnmarshalTypeError) Error
```go
func (e *UnmarshalTypeError) Error() string
```
Unmarshaler 是一个接口，实现它的接口可以unmarshal 一个它们自己的JSON描述。输入可以假设为一个JSON值的有效的编码。

UnmarshalJSON 如果想返回后仍然保留数据，它必须复制JSON数据。

###type Unmarshaler interface
```go
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```
当Marshal 尝试编码一个不支持的值类型，将返回UnsupportedTypeError 。

###type UnsupportedTypeError struct
```go
type UnsupportedTypeError struct {
    Type reflect.Type
}
```

###func (*UnsupportedTypeError) Error
```go
func (e *UnsupportedTypeError) Error() string
```

###type UnsupportedValueError struct
```go
type UnsupportedValueError struct {
    Value reflect.Value
    Str   string
}
```

###func (*UnsupportedValueError) Error
```go
func (e *UnsupportedValueError) Error() string
```

