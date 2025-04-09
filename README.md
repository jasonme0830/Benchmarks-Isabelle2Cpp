# Benchmarks-Isabelle2Cpp
To benchmark C++ codes that generated from Isabelle2Cpp.

The files c_peter, imp, sml, llvm of bs are download 
from https://github.com/lammich/isabelle_llvm.git,
which created by peter lammich.

Need: 
clang version 10.0.0-4ubuntu1
go version go1.18.10 linux/amd64

All the operations are displayed in the makefile.

benchmark-Isabelle2Cpp
    |--function name(e.g addListHead, bs)

        |--llvm  (generated by Isabelle-llvm of peter lammich)

        |--cpp_before  (generated by Isabelle2Cpp)

        |--cpp_after  (generated by optimized Isabelle2Cpp)

        |--c_peter  (copy from benchmark of peter lammich)
    
        |--cpp_slist  (by Isabelle2Cpp but not use std type)

        |--imp  (copy from benchmark of peter lammich)

        |--sml_peter  (copy from benchmark of peter lammich)

        |--go  (generated by Isabelle-go)

        |--go_rule  (by Isabelle-go, but use generic type)


This message is for second reclone code.

vscode settings
{
    "configurations": [
        {
            "name": "Linux",
            "includePath": [
                "${workspaceFolder}/**"
            ],
            "defines": [],
            "compilerPath": "/usr/bin/g++-11",
            "cStandard": "c17",
            "cppStandard": "c++17",
            "intelliSenseMode": "linux-gcc-x64"
        }
    ],
    "version": 4
}