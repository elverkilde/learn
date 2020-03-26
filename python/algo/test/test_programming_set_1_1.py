#!/usr/bin/python -tt
# -*- coding: utf-8 -*-

"""
Download the text file here. Use IntegerArray.txt

This file contains all of the 100,000 integers between 1 and 100,000
(inclusive) in some order, with no integer repeated.

Your task is to compute the number of inversions in the file given,
where the ith row of the file indicates the ith entry of an array.
Because of the large size of this array, you should implement the
fast divide-and-conquer algorithm covered in the video lectures.
The numeric answer for the given input file should be typed in the
space below.

So if your answer is 1198233847, then just type 1198233847 in the
space provided without any space / commas / any other punctuation
marks. You can make up to 5 attempts, and we'll use the best one for grading.
(We do not require you to submit your code, so feel free to use any
programming language you want --- just type the final numeric answer
in the following space.)

[TIP: before submitting, first test the correctness of your program
on some small test files or your own devising. Then post your best
test cases to the discussion forums to help your fellow students!]
"""

#import os
#import unittest
#
#from src.counting_inversions import sort_and_count_inversions
#
#
#class ProblemSet1Test(unittest.TestCase):
#
#    def test_count_inversions_for_large_dataset(self):
#        numbers = []
#        with open('{base}/test/IntegerArray.txt'.format(base=os.getcwd()), 'r') as f:
#            for line in f:
#                numbers.append(int(line))
#
#        sorted_arr, num_inversions = sort_and_count_inversions(numbers)
#        self.assertEqual(num_inversions, 2407905288, 'should return the correct value')


## Brute force approach.
#num_inversions = 0
#n = len(numbers)
#for i in xrange(n):
#    for j in xrange(i, n):
#        if numbers[i] > numbers[j]:
#            num_inversions += 1
#print num_inversions
