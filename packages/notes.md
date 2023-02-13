1. Evry Go program is made up of packages.
2. Programs start running in package main.
3. By convention the package name is the same as the last eleemnt of the import path.
4. import code groups the imports into a parenthesized,"factored" import statement.
5. Can write multiple import statements too.
6. In go, a name is exported if it begins with a capital letter, wehn importing a package we can only refer to only to its exported names. Any "unexported names" are not accessible from outside the package.
7. `