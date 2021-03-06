// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
Package column contains the object to reference particular parts of a column.

See: http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/Column.html

Column is the object that defines parts of a column in Gremlin.
These values are such as:
Keys - Define the keys associated with the columns.
Values - Define the values associated with the columns.

A note about Column:

This object implements the Parameter interfaces used by graph traversals.
*/
package column

// http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/Column.html

// Column references a particular type of column in a complex
// data structure such as a Map, a Map.Entry, or a Path.
type Column string

const (
	// Keys associated with the data structure.
	Keys Column = "keys"
	// Values associated with the data structure.
	Values Column = "values"
)

func (c Column) String() string {
	return string(c)
}
