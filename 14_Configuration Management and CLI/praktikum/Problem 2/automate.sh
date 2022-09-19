#!/bin/sh

timestamp=$(date +%H.%M.%S)

mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)
mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me
mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/personal
mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/professional
mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_friends
mkdir ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info
touch ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/personal/facebook.txt
touch ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/professional/linkedin.txt
touch ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_friends/list_of_my_friends.txt
touch ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/about_this_laptop.txt
touch ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/internet_connection.txt

echo "https://www.facebook.com/$2" > ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/personal/facebook.txt
echo "https://www.linkedin.com/in/$3" > ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/about_me/professional/linkedin.txt

echo "$(curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt)" > ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_friends/list_of_my_friends.txt

echo "My username: $(uname)" > ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/about_this_laptop.txt
echo "My hostname: $(hostname)" >> ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/about_this_laptop.txt
echo "With host: $(uname -a)" >> ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/about_this_laptop.txt

echo "Connection to google:" > ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/internet_connection.txt
echo "$(ping google.com)" >> ./$1\ at\ $(date +%a)\ $(date +%b)\ $(date +%d)\ $timestamp\ WIB\ $(date +%Y)/my_system_info/internet_connection.txt

echo "Done"