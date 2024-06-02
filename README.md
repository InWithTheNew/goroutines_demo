# Go Routines Example

This very, very simple app is just an attempt at showing the power of go routines / concurrent processing.

It calls 2 funcs, ``DoBigQuery()`` and ``DoBigQueryInARoutine()``

As you'll see, it runs ``DoBigQuery()`` only 20 times, and ``DoBigQueryInARoutine()`` 50,000 times. They both call ``DoComplicatedComputeThingy()`` for the number that's passed into it. 

``DoComplicatedComputeThingy()`` simulated some kind of compute / transaction that'll take .2 seconds.

## How to run

Clone the repo, ensure you've got go installed.
cd into the directory and run ``go run .`` 

The output will be all of the calculations it's run from ``DoComplicatedComputeThingy()`` and then print the total time for sQ (serial query (no go routines)), and cQ (concurrent query (go routines)).