#!/bin/bash

go run . badexample00.txt | cat -e
echo

go run . badexample01.txt | cat -e
echo

go run . badexample02.txt | cat -e
echo

go run . badexample03.txt | cat -e
echo

go run . badexample04.txt | cat -e
echo

go run . badexample.txt | cat -e
echo

go run . goodexample00.txt | cat -e
echo

go run . goodexample01.txt | cat -e
echo

go run . goodexample02.txt | cat -e
echo

go run . goodexample03.txt | cat -e
echo

go run . hardexam.txt | cat -e
echo
