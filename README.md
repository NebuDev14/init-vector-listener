# Initialzation Vector Listener

## What is this?
This is the source code for the listener that powers [Initilization Vector](https://github.com/NebuDev14/initialization-vector), the app that handles 2023 MIT BWSI's lab submissions. It will listen for incoming TCP connections for BWSI students to submit their challenge flags, and update their completion status when appropriate.

## Setup and installation

Download the source code to your computer:

    $ https://github.com/NebuDev14/init-vector-listener

Install dependencies:

    $ go install

Run the server:

    $ go run main.go