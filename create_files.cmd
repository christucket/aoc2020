@echo off
if not exist ./inputs/day%1.sample break > ./inputs/day%1.sample
if not exist ./inputs/day%1.input break > ./inputs/day%1.input
if not exist ./day%1.go break > ./day%1.go