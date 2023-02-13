1. var stmt declares a list of variables as in function argument lists,the type is last.
2. it can be at the package or function level.
3. A var declaration can include initializers,one per variable.
4. If an initializer is present , the type can be omitted, the varibale will take the type of the initiailizer.
5. Inside a funtion the :=(short assignment) statements can be used in place of a var declaation with implicit types.
6. Outside a function ,every statement begins with a  keywordand so the := construct is not available.
7. Go's basic types are:
   1. bool
   2. string 
   3. int int8,int16,int32(rune),int64 
   4. uint,uint8(byte),uint61,uint32,uint64,uintptr
   5. float32,float64
   6. rune -> represents a Unicode code point 
   7. complex64 complex128
8. Varibales declared without an explicit initial value are given their 0 value.
9. 0 values are:
   1.  0 for numeric types
   2.  false for boolean type 
   3.  "" for strings 
10. The expression T(v) converts the value V to the type T.
11. When declaring a variable without specifying an explicit type(using := syntax or var = expresion syntax), the aribales type is inferred from the value on the right hand side.
12. If the right hand size contains a untyped numeric constantn, the new varibale may be a int, float64 or complex128 .
13. Constants are declared like varibales, but with the const keyword. 
14. Constants can be character, string,boolean or numeric values.
15. Constants cannot be delcared using the := syntax.
16. Numeric constants are high-precision values. An untyped constant takes the type needed by its context.
17. 
